package schema

// List of basic types
const (
	INTEGER int = iota
	VARCHAR
	BOOLEAN
	REAL
	BLOB
)

// List of vendor-specific keywords
const (
	AUTO_INCREMENT = iota
	PRIMARY_KEY
)

type Table struct {
	Name         string
	LastInsertId bool

	Fields  []*Field
	Index   []*Index
	Primary []*Field
}

type Field struct {
	Name    string
	Type    int
	Primary bool
	Auto    bool
	Size    int
}

type Index struct {
	Name   string
	Unique bool

	Fields []*Field
}
