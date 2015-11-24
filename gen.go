package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/drone/sqlgen/parse"
	"github.com/drone/sqlgen/schema"
)

var (
	input      = flag.String("file", "", "input file name; required")
	output     = flag.String("o", "", "output file name; required")
	pkgName    = flag.String("pkg", "main", "output package name; required")
	typeName   = flag.String("type", "", "type to generate; required")
	database   = flag.String("db", "sqlite", "sql dialect; required")
	genSchema  = flag.Bool("schema", true, "generate sql schema and queries")
	genFuncs   = flag.Bool("funcs", true, "generate sql helper functions")
	extraFuncs = flag.Bool("extras", true, "generate extra sql helper functions")
	doDI       = flag.Bool("doDI", false, "generate an interface for DI")
	diName     = flag.String("diName", "sql.DB", "the concrete class to attach methods to")
)

func main() {
	flag.Parse()

	// parses the syntax tree into something a bit
	// easier to work with.
	tree, err := parse.Parse(*input, *typeName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// if the code is generated in a different folder
	// that the struct we need to import the struct
	if tree.Pkg != *pkgName && *pkgName != "main" {
		// TODO
	}

	// load the Tree into a schema Object
	table := schema.Load(tree)
	dialect := schema.New(schema.Dialects[*database])

	var buf bytes.Buffer

	if *genFuncs {
		writePackage(&buf, *pkgName)
		writeImports(&buf, tree, "database/sql")
		writeRowFunc(&buf, tree)
		writeRowsFunc(&buf, tree)
		writeSliceFunc(&buf, tree)

		if *extraFuncs {
			writeSelectRow(&buf, tree, *doDI, *diName)
			writeSelectRows(&buf, tree, *doDI, *diName)
			writeInsertFunc(&buf, tree, *doDI, *diName)
			writeUpdateFunc(&buf, tree, *doDI, *diName)
		}
		if *doDI {
			writeInterface(&buf, tree)
		}
	} else {
		writePackage(&buf, *pkgName)
	}

	// write the sql functions
	if *genSchema {
		writeSchema(&buf, dialect, table)
	}

	// formats the generated file using gofmt
	pretty, err := format(&buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	// create output source for file. defaults to
	// stdout but may be file.
	var out io.WriteCloser = os.Stdout
	if *output != "" {
		out, err = os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		defer out.Close()
	}

	io.Copy(out, pretty)
}
