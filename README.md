**sqlgen** generates SQL statements and database helper functions from your Go structs. It can be used in place of a simple ORM or hand-written SQL. See the [demo](https://github.com/drone/sqlgen/tree/master/demo) directory for examples.

### Install

Install or upgrade with this command:

```
go get -u github.com/drone/sqlgen
```

### Usage

```
Usage of sqlgen:
  -type string
    	type to generate; required
  -file string
    	input file name; required
  -o string
    	output file name
  -pkg string
    	output package name
  -db string
    	sql dialect; sqlite, postgres, mysql
  -schema
    	generate sql schema and queries; default true
  -funcs
    	generate sql helper functions; default true
```

### Tutorial

First, let's start with a simple `User` struct in `user.go`:

```Go
type User struct {
	ID     int64
	Login  string
	Email  string
}
```

We can run the following command:

```
sqlgen -file user.go -type User -pkg demo
```

The tool outputs the following generated code:

```Go
func ScanUser(row *sql.Row) (*User, error) {
	var v0 int64
	var v1 string
	var v2 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
	)
	if err != nil {
		return nil, err
	}

	v := &User{}
	v.ID = v0
	v.Login = v1
	v.Email = v2

	return v, nil
}

const CreateUserStmt = `
CREATE TABLE IF NOT EXISTS users (
 user_id     INTEGER
,user_login  TEXT
,user_email  TEXT
);
`

const SelectUserStmt = `
SELECT 
 user_id
,user_login
,user_email
FROM users 
`

const SelectUserRangeStmt = `
SELECT 
 user_id
,user_login
,user_email
FROM users 
LIMIT ? OFFSET ?
`


// more functions and sql statements not displayed
```

This is a great start, but what if we want to specify primary keys, column sizes and more? This may be acheived by annotating your code using Go tags. For example, we can tag the `ID` field to indicate it is a primary key and will auto increment:

```Go
type User struct {
    ID      int64  `sql:"pk: true, auto: true"`
    Login   string
    Email   string
}
```

This information allows the tool to generate smarter SQL statements:

```diff
CREATE TABLE IF NOT EXISTS users (
-user_id     INTEGER PRIMARY KEY AUTOINCREMENT
+user_id     INTEGER
,user_login  TEXT
,user_email  TEXT
);
```

Including SQL statements to select, insert, update and delete data using the primary key:

```Go
const SelectUserPkeyStmt = `
SELECT 
 user_id
,user_login
,user_email
WHERE user_id=?
`

const UpdateUserPkeyStmt = `
UPDATE users SET 
 user_id=?
,user_login=?
,user_email=?
WHERE user_id=?
`

const DeleteUserPkeyStmt = `
DELETE FROM users 
WHERE user_id=?
`
```

We can take this one step further and annotate indexes. In our example, we probably want to make sure the `user_login` field has a unique index:

```Go
type User struct {
    ID      int64  `sql:"pk: true, auto: true"`
    Login   string `sql:"unique: user_login"`
    Email   string
}
```

This information instructs the tool to generate the following:


```Go
const CreateUserLogin = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON users (user_login)
```

The tool also assumes that we probably indend to fetch data from the database using this index. The tool will therefore automatically generate the following queries:

```Go
const SelectUserLoginStmt = `
SELECT 
 user_id
,user_login
,user_email
WHERE user_login=?
`

const UpdateUserLoginStmt = `
UPDATE users SET 
 user_id=?
,user_login=?
,user_email=?
WHERE user_login=?
`

const DeleteUserLoginStmt = `
DELETE FROM users 
WHERE user_login=?
`
```

### Nesting

Nested Go structures can be flattened into a single database table. As an example, we have a `User` and `Address` with a one-to-one relationship. It may not always make sense to normalize our data across tables.

```Go
type User struct {
	ID     int64  `sql:"pk: true"`
	Login  string
	Email  string
	Addr   *Address
}

type Address struct {
	City   string
	State  string
	Zip    string `sql:"index: user_zip"`
}
```

The above relationship is flattened into a single table (see below). When the data is retrieved from the database the nested structure is restored.

```sql
CREATE TALBE IF NOT EXISTS users (
 user_id         INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login      TEXT
,user_email      TEXT
,user_addr_city  TEXT
,user_addr_state TEXT
,user_addr_zip   TEXT
);
```

### Dialects

You may specify one of the following SQL dialects when generating your code: `postgres`, `mysql` and `sqlite`. The default value is `sqlite`.

```
sqlgen -file user.go -type User -pkg demo -db postgres
```


### Go Generate

Example use with `go:generate`:

```Go
package demo

//go:generate sqlgen -file user.go -type User -pkg demo -o user_sql.go

type User struct {
	ID     int64  `sql:"pk: true, auto: true"`
	Login  string `sql:"unique: user_login"`
	Email  string `sql:"size: 1024"`
	Avatar string
}
```

### Benchmarks

This tool demonstrates performance gains, albeit small, over light-weight ORM packages such as `sqlx` and `meddler`. Over time I plan to expand the benchmarks to include additional ORM packages.

To run the project benchmarks:

```
go get ./...
go generate ./...
go build
cd bench
go test -bench=Bench
```

Example selecing a single row:

```
BenchmarkMeddlerRow-4      30000        42773 ns/op
BenchmarkSqlxRow-4         30000        41554 ns/op
BenchmarkSqlgenRow-4       50000        39664 ns/op

```

Selecting multiple rows:

```
BenchmarkMeddlerRows-4      2000      1025218 ns/op
BenchmarkSqlxRows-4         2000       807213 ns/op
BenchmarkSqlgenRows-4       2000       700673 ns/op
```


#### Credits

This tool was inspired by [scaneo](https://github.com/variadico/scaneo).
