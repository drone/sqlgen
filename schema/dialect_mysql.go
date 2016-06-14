package schema

import (
	"fmt"
)

type mysql struct {
	base
}

func newMySQL() Dialect {
	d := &mysql{}
	d.base.Dialect = d
	return d
}

func (d *mysql) Column(f *Field) (_ string) {
	switch f.Type {
	case INTEGER:
		return "INTEGER"
	case BOOLEAN:
		return "BOOLEAN"
	case BLOB:
		return "MEDIUMBLOB"
	case VARCHAR:
		// assigns an arbitrary size if
		// none is provided.
		size := f.Size
		if size == 0 {
			size = 512
		}
		return fmt.Sprintf("VARCHAR(%d)", size)
	default:
		return
	}
}

func (d *mysql) Token(v int) (_ string) {
	switch v {
	case AUTO_INCREMENT:
		return "AUTO_INCREMENT"
	case PRIMARY_KEY:
		return "PRIMARY KEY"
	default:
		return
	}
}
