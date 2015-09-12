package bench

import (
	"database/sql"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/meddler"
)

var db *sql.DB
var dbx *sqlx.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db.Exec("DROP TABLE users;")
	dbx = sqlx.NewDb(db, "sqlite3")

	ddl := []string{CreateUserStmt}
	for _, stmt := range ddl {
		_, err = db.Exec(stmt)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < 100; i++ {
		user := &User{}
		user.Name = randomdata.FullName(randomdata.RandomGender)
		user.Email = randomdata.Email()
		user.Pass = "pa55word"
		user.Created = time.Now().Unix()
		user.Updated = time.Now().Unix()

		err := InsertUser(db, InsertUserStmt, user)
		if err != nil {
			panic(err)
		}
	}
}

var result *User
var results []*User

func BenchmarkMeddlerRow(b *testing.B) {
	var user *User
	var err error

	for n := 0; n < b.N; n++ {
		user = &User{}
		err = meddler.QueryRow(db, user, SelectUserPkeyStmt, 1)
		if err != nil {
			panic(err)
		}
	}
	result = user
}

func BenchmarkMeddlerRows(b *testing.B) {
	var users []*User
	var err error

	for n := 0; n < b.N; n++ {
		err = meddler.QueryAll(db, &users, SelectUserStmt)
		if err != nil {
			panic(err)
		}
	}
	results = users
}

func BenchmarkSqlxRow(b *testing.B) {
	var user *User
	var err error

	for n := 0; n < b.N; n++ {
		user = &User{}
		err = dbx.Get(user, SelectUserPkeyStmt, 1)
		if err != nil {
			panic(err)
		}
	}
	result = user
}

func BenchmarkSqlxRows(b *testing.B) {
	var users []*User
	var err error

	for n := 0; n < b.N; n++ {
		err = dbx.Select(&users, SelectUserStmt)
		if err != nil {
			panic(err)
		}
	}
	results = users
}

func BenchmarkSqlgenRow(b *testing.B) {
	var user *User
	var err error

	for n := 0; n < b.N; n++ {
		user, err = SelectUser(db, SelectUserPkeyStmt, 1)
		if err != nil {
			panic(err)
		}
	}
	result = user
}

func BenchmarkSqlgenRows(b *testing.B) {
	var users []*User
	var err error

	for n := 0; n < b.N; n++ {
		users, err = SelectUsers(db, SelectUserStmt)
		if err != nil {
			panic(err)
		}
	}
	results = users
}
