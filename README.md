**sqlgen** generates SQL statements and helper functions from your Go structures. It can be used in place of a simple ORM or hand-written SQL. See the **demo** directory for examples.

## Running Benchmarks

```sh
sqlgen -file bench/type.go -type User -pkg bench -o bench/type_sql.go
cd bench
go test -bench=Bench
```

## Benchmark Results

Select a single row:

```
BenchmarkMeddlerRow-4      30000        42773 ns/op
BenchmarkSqlxRow-4         30000        41554 ns/op
BenchmarkSqlRow-4          50000        39664 ns/op

```

Select multiple rows:

```
BenchmarkMeddlerRows-4      2000      1025218 ns/op
BenchmarkSqlxRows-4         2000       807213 ns/op
BenchmarkSqlRows-4          2000       700673 ns/op
```

## Tags

```
- `auto` the field is auto-incremented
- `pk` the field is a primary key
- `size` the field size
- `type` the field type
- `index` the field uses the specified index
- `unique` the field uses the specified unique index
- `-` ignores the field
```

For example:

```
type User struct {
    ID      int64  `sql:"pk: true, auto: true"    // primary key
    Name    string                                
    Phone   string `sql:"size:255"`               // override field type
    Email   string `sql:"unique: ux_user_login"`  // creates index
    Company string `sql:"index: ix_user_company"` // creates unique index
    Temp    string `sql:"-"`                      // skip this field
}
```
