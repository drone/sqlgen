package demo

//go:generate ../sqlgen -file issue.go -type Issue -pkg demo -o issue_sql.go -db postgres

type Issue struct {
	ID       int64 `sql:"pk: true, auto: true"`
	Number   int
	Title    string   `sql:"size: 512"`
	Body     string   `sql:"size: 2048"`
	Assignee string   `sql:"index: issue_assignee"`
	State    string   `sql:"size: 50"`
	Labels   []string `sql:"encode: json"`

	locked bool `sql:"-"`
}
