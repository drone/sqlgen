package schema

type sqlite struct {
	base
}

func newSQLite() Dialect {
	d := &sqlite{}
	d.base.Dialect = d
	return d
}
