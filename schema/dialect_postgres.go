package schema

import (
	"fmt"
)

type postgres struct {
	base
}

func newPostgres() Dialect {
	d := &postgres{}
	d.base.Dialect = d
	return d
}

func (d *postgres) Column(f *Field) (_ string) {
	// postgres uses a special column type
	// to autoincrementing keys.
	if f.Auto {
		return "SERIAL"
	}

	switch f.Type {
	case INTEGER:
		return "INTEGER"
	case BOOLEAN:
		return "BOOLEAN"
	case BLOB:
		return "BYTEA"
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

func (d *postgres) Token(v int) (_ string) {
	switch v {
	case AUTO_INCREMENT:
		// postgres does not support the
		// auto-increment keyword.
		return
	case PRIMARY_KEY:
		return "PRIMARY KEY"
	default:
		return
	}
}

func (d *postgres) Param(i int) string {
	return fmt.Sprintf("$%d", i+1)
}
