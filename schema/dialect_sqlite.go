package schema

type sqlite struct {
	base
}

func newSqlite() Dialect {
	d := &sqlite{}
	d.base.Dialect = d
	return d
}
