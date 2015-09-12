package bench

// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

import (
	"database/sql"
)

func ScanUser(row *sql.Row) (*User, error) {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 int64
	var v6 int64

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
	)
	if err != nil {
		return nil, err
	}

	v := &User{}
	v.ID = v0
	v.Name = v1
	v.Pass = v2
	v.Email = v3
	v.Active = v4
	v.Created = v5
	v.Updated = v6

	return v, nil
}

func ScanUsers(rows *sql.Rows) ([]*User, error) {
	var err error
	var vv []*User

	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 int64
	var v6 int64

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,
		)
		if err != nil {
			return vv, err
		}

		v := &User{}
		v.ID = v0
		v.Name = v1
		v.Pass = v2
		v.Email = v3
		v.Active = v4
		v.Created = v5
		v.Updated = v6

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceUser(v *User) []interface{} {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 int64
	var v6 int64

	v0 = v.ID
	v1 = v.Name
	v2 = v.Pass
	v3 = v.Email
	v4 = v.Active
	v5 = v.Created
	v6 = v.Updated

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
	}
}

func SelectUser(db *sql.DB, query string, args ...interface{}) (*User, error) {
	row := db.QueryRow(query, args...)
	return ScanUser(row)
}

func SelectUsers(db *sql.DB, query string, args ...interface{}) ([]*User, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanUsers(rows)
}

func InsertUser(db *sql.DB, query string, v *User) error {

	res, err := db.Exec(query, SliceUser(v)[1:]...)
	if err != nil {
		return err
	}

	v.ID, err = res.LastInsertId()
	return err
}

func UpdateUser(db *sql.DB, query string, v *User) error {

	args := SliceUser(v)[1:]
	args = append(args, v.ID)
	_, err := db.Exec(query, args...)
	return err
}

const CreateUserStmt = `
CREATE TABLE IF NOT EXISTS users (
 user_id      INTEGER PRIMARY KEY AUTOINCREMENT
,user_name    TEXT
,user_pass    TEXT
,user_email   TEXT
,user_active  BOOLEAN
,user_created INTEGER
,user_updated INTEGER
);
`

const InsertUserStmt = `
INSERT INTO users (
 user_name
,user_pass
,user_email
,user_active
,user_created
,user_updated
) VALUES (?,?,?,?,?,?)
`

const SelectUserStmt = `
SELECT 
 user_id
,user_name
,user_pass
,user_email
,user_active
,user_created
,user_updated
FROM users 
`

const SelectUserRangeStmt = `
SELECT 
 user_id
,user_name
,user_pass
,user_email
,user_active
,user_created
,user_updated
FROM users 
LIMIT ? OFFSET ?
`

const SelectUserCountStmt = `
SELECT count(1)
FROM users 
`

const SelectUserPkeyStmt = `
SELECT 
 user_id
,user_name
,user_pass
,user_email
,user_active
,user_created
,user_updated
FROM users 
WHERE user_id=?
`

const UpdateUserPkeyStmt = `
UPDATE users SET 
 user_id=?
,user_name=?
,user_pass=?
,user_email=?
,user_active=?
,user_created=?
,user_updated=? 
WHERE user_id=?
`

const DeleteUserPkeyStmt = `
DELETE FROM users 
WHERE user_id=?
`

const CreateUserNameStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_name ON users (user_name)
`

const SelectUserNameStmt = `
SELECT 
 user_id
,user_name
,user_pass
,user_email
,user_active
,user_created
,user_updated
FROM users 
WHERE user_name=?
`

const UpdateUserNameStmt = `
UPDATE users SET 
 user_id=?
,user_name=?
,user_pass=?
,user_email=?
,user_active=?
,user_created=?
,user_updated=? 
WHERE user_name=?
`

const DeleteUserNameStmt = `
DELETE FROM users 
WHERE user_name=?
`

const CreateUserEmailStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON users (user_email)
`

const SelectUserEmailStmt = `
SELECT 
 user_id
,user_name
,user_pass
,user_email
,user_active
,user_created
,user_updated
FROM users 
WHERE user_email=?
`

const UpdateUserEmailStmt = `
UPDATE users SET 
 user_id=?
,user_name=?
,user_pass=?
,user_email=?
,user_active=?
,user_created=?
,user_updated=? 
WHERE user_email=?
`

const DeleteUserEmailStmt = `
DELETE FROM users 
WHERE user_email=?
`
