package demo

//go:generate ../sqlgen -file user.go -type User -pkg demo -o user_sql.go

type User struct {
	ID     int64  `sql:"pk: true, auto: true"`
	Login  string `sql:"unique: user_login"`
	Email  string `sql:"unique: user_email"`
	Avatar string
	Active bool
	Admin  bool

	// oauth token and secret
	token  string
	secret string

	// randomly generated hash used to sign user
	// session and application tokens.
	hash string
}
