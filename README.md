**sqlgen** generates SQL statements and helper functions from your Go structs. It can be used in place of a simple ORM or hand-written SQL. See the [demo](https://github.com/drone/sqlgen/tree/master/demo) directory for examples.

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

```
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

This will output the following generated code:

```
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

const SelectUsers = `
SELECT 
 user_id
,user_login
,user_email
FROM users 
`

const SelectUserRange = `
SELECT 
 user_id
,user_login
,user_email
FROM users 
LIMIT ? OFFSET ?
`


// more functions and sql statements not displayed
```

### Tags

You may annotate your fields with the following tags:

* `auto` the field is auto-incremented
* `pk` the field is a primary key
* `size` the field size
* `type` the field type (TODO)
* `index` the field uses the specified index
* `unique` the field uses the specified unique index
* `-` ignores the field

For example:

```
type User struct {
    ID      int64  `sql:"pk: true, auto: true"` // primary key, increment
    Login   string `sql:"unique: user_login"    // creates unique index
    Email   string `sql:"size:255"`
    Company string `sql:"index: user_company"`  // creates index
    Temp    string `sql:"-"`                    // skip this field
}
```

Adding `unique` and `index` tags will generate `create index` statements, as well as `select`, `select count`, `select range`, `update` and `delete` statements using the indexed fields. For example:

```
const CreateUserLogin = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON users (user_login)
`

const SelectUserLogin = `
SELECT 
 user_id
,user_login
,user_email
,user_company
FROM users 
WHERE user_login=?
`


// more sql statements not displayed
```

### Nesting

Nested Go structures can be flattened into a single database table. As an example, we have a `User` and `Address` with a one-to-one relationship. It may not always make sense to normalize our data across tables.

```
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

```
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

```
package demo

//go:generate sqlgen -file user.go -type User -pkg demo -o user_sql.go

type User struct {
	ID     int64  `sql:"pk: true, auto: true"`
	Login  string `sql:"unique: user_login"`
	Email  string `sql:"unique: user_email"`
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
