package bench

//go:generate ../sqlgen -file type.go -type User -pkg bench -o type_sql.go

type User struct {
	ID      int64  `sql:"pk: true, auto: true"   meddler:"user_id,pk"   db:"user_id"`
	Name    string `sql:"unique: user_name"      meddler:"user_name"    db:"user_name"`
	Pass    string `sql:""                       meddler:"user_pass"    db:"user_pass"`
	Email   string `sql:"unique: user_email"     meddler:"user_email"   db:"user_email"`
	Active  bool   `sql:"index:  user_active"    meddler:"user_active"  db:"user_active"`
	Created int64  `sql:""                       meddler:"user_created" db:"user_created"`
	Updated int64  `sql:""                       meddler:"user_updated" db:"user_updated"`
}
