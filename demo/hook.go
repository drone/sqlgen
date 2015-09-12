package demo

//go:generate ../sqlgen -file hook.go -type Hook -pkg demo -o hook_sql.go -db mysql

type Hook struct {
	ID         int64 `sql:"pk: true, auto: true"`
	Sha        string
	After      string
	Before     string
	Created    bool
	Deleted    bool
	Forced     bool
	HeadCommit *Commit
}

type Commit struct {
	ID        string
	Message   string
	Timestamp string
	Author    *Author
	Committer *Author
}

type Author struct {
	Name     string
	Email    string
	Username string
}
