package main

// template to create a constant variable.
var sConst = `
const %s = %s
`

// template to wrap a string in multi-line quotes.
var sQuote = "`\n%s\n`"

// template to declare the package name.
var sPackage = `
package %s

// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.
`

// template to delcare the package imports.
var sImport = `
import (
	%s
)
`

// function template to scan a single row.
const sScanRow = `
func Scan%s(row *sql.Row) (*%s, error) {
	%s

	err := row.Scan(
		%s
	)
	if err != nil {
		return nil, err
	}

	v := &%s{}
	%s

	return v, nil
}
`

// function template to scan multiple rows.
const sScanRows = `
func Scan%s(rows *sql.Rows) ([]*%s, error) {
	var err error
	var vv []*%s

	%s
	for rows.Next() {
		err = rows.Scan(
			%s
		)
		if err != nil {
			return vv, err
		}

		v := &%s{}
		%s
		vv = append(vv, v)
	}
	return vv, rows.Err()
}
`

const sSliceRow = `
func Slice%s(v *%s) []interface{} {
	%s
	%s

	return []interface{}{
		%s
	}
}
`

const sSelectRow = `
func Select%s(db *sql.DB, query string, args ...interface{}) (*%s, error) {
	row := db.QueryRow(query, args...)
	return Scan%s(row)
}
`

// function template to select multiple rows.
const sSelectRows = `
func Select%s(db *sql.DB, query string, args ...interface{}) ([]*%s, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return Scan%s(rows)
}
`

// function template to insert a single row.
const sInsertAndGetLastId = `
func Insert%s(db *sql.DB, query string, v *%s) error {
	res, err := db.Exec(query, Slice%s(v)[1:]...)
	if err != nil {
		return err
	}

	v.ID, err = res.LastInsertId()
	return err
}
`

const sInsert = `
func Insert%s(db *sql.DB, query string, v *%s) error {
	_, err := db.Exec(query, Slice%s(v)...)
	return err
}
`

// function template to update a single row.
const sUpdate = `
func Update%s(db *sql.DB, query string, v *%s) error {

	args := Slice%s(v)[1:]
	args = append(args, v.ID)
	_, err := db.Exec(query, args...)
	return err 
}
`
