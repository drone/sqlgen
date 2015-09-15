package demo

//go:generate ../sqlgen -file user.go -type LegacyUser -pkg demo -o user_sql.go

type LegacyUser struct {
	SQLName string `sql:"name: users, skip: true"`
	ID      int64  `sql:"pk: true, auto: true"`
	Login   string `sql:"unique: login"`
	Email   string `sql:"unique: email"`
	Avatar  string
	Active  bool
	Admin   bool

	// oauth token and secret
	token  string
	secret string

	// randomly generated hash used to sign user
	// session and application tokens.
	hash string
}
