package parse

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

var (
	ErrTypeNotFound = errors.New("Cannot find type in the source code.")
	ErrTypeInvalid  = errors.New("Cannot convert type to a SQL type.")
)

func Parse(path, name string) (*Node, error) {

	var fset = token.NewFileSet()
	var file, err = parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		spec, ok := gen.Specs[0].(*ast.TypeSpec)
		if !ok {
			continue
		}
		if spec.Name.String() != name {
			continue
		}

		var node = new(Node)
		node.Name = spec.Name.String()
		node.Type = spec.Name.String()
		node.Pkg = file.Name.Name
		err = buildNodes(node, spec)
		return node, err
	}

	return nil, ErrTypeNotFound
}

func buildNodes(parent *Node, spec *ast.TypeSpec) error {
	ident, ok := spec.Type.(*ast.StructType)
	if !ok {
		return ErrTypeInvalid
	}

	for _, field := range ident.Fields.List {
		var tag string
		if field.Tag != nil {
			tag = field.Tag.Value
		}
		buildNode(parent, field.Type, field.Names[0].Name, tag)
	}
	return nil
}

func buildNode(parent *Node, expr ast.Expr, name, tag string) error {
	var err error

	switch ident := expr.(type) {
	case *ast.Ident:
		if ident.Obj == nil {
			node := &Node{
				Name: name,
				Type: ident.Name,
				Kind: Types[ident.Name],
			}
			node.Tags, err = parseTag(tag)
			if err != nil {
				return err
			}
			parent.append(node)
			return nil
		}
		spec, ok := ident.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			goto invalidType
		}
		node := &Node{
			Name: name,
			Type: ident.Name,
			Kind: Struct,
		}
		node.Tags, err = parseTag(tag)
		if err != nil {
			return err
		}
		parent.append(node)
		return buildNodes(node, spec)

	case *ast.ArrayType:
		if ident.Len != nil {
			goto invalidType
		}
		node := &Node{
			Name: name,
			Kind: Slice,
			Type: fmt.Sprintf("[]%s", ident.Elt),
		}
		node.Tags, err = parseTag(tag)
		if err != nil {
			return err
		}
		if node.Type == "[]byte" {
			node.Kind = Bytes
		}
		parent.append(node)
		return nil

	case *ast.MapType:
		type_ := fmt.Sprintf("map[%s]%s", ident.Key, ident.Value)
		node := &Node{Name: name, Type: type_, Kind: Map}
		node.Tags, err = parseTag(tag)
		if err != nil {
			return err
		}
		parent.append(node)
		return nil

	case *ast.StarExpr:
		innerIdent, ok := ident.X.(*ast.Ident)
		if !ok {
			goto invalidType
		}
		if innerIdent.Obj == nil || innerIdent.Obj.Decl == nil {
			goto invalidType
		}
		spec, ok := innerIdent.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			goto invalidType
		}
		node := &Node{Name: name, Type: innerIdent.Name, Kind: Ptr}
		node.Tags, err = parseTag(tag)
		if err != nil {
			return err
		}
		if node.Tags.Skip {
			return nil
		}
		parent.append(node)
		return buildNodes(node, spec)
	}

invalidType:
	return fmt.Errorf("%s is not a valid type", name)
}
