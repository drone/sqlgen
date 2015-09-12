package parse

type Node struct {
	Pkg  string // source code package.
	Name string // source code name.
	Kind uint8  // source code kind.
	Type string // source code type.
	Tags *Tag

	Parent *Node
	Nodes  []*Node
}

func (n *Node) append(node *Node) {
	node.Parent = n
	n.Nodes = append(n.Nodes, node)
}

// Walk traverses the node tree, invoking the callback
// function for each node that is traversed.
func (n *Node) Walk(fn func(*Node)) {
	for _, node := range n.Nodes {
		fn(node)
		node.Walk(fn)
	}
}

// WalkRev traverses the tree in reverse order, invoking
// the callback function for each parent node until
// the root node is reached.
func (n *Node) WalkRev(fn func(*Node)) {
	if n.Parent != nil {
		n.Parent.WalkRev(fn)
	}
	fn(n) // this was previously inside the if block
}

// Edges returns a flattened list of all edge
// nodes in the Tree.
func (n *Node) Edges() []*Node {
	var nodes []*Node
	n.Walk(func(node *Node) {
		if len(node.Nodes) == 0 {
			nodes = append(nodes, node)
		}
	})
	return nodes
}

// Path returns the absolute path of the node
// in the Tree.
func (n *Node) Path() []*Node {
	var nodes []*Node
	n.WalkRev(func(node *Node) {
		nodes = append(nodes, node)
	})
	return nodes
}
