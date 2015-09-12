package demo

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
	var v5 bool
	var v6 string
	var v7 string
	var v8 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
	)
	if err != nil {
		return nil, err
	}

	v := &User{}
	v.ID = v0
	v.Login = v1
	v.Email = v2
	v.Avatar = v3
	v.Active = v4
	v.Admin = v5
	v.token = v6
	v.secret = v7
	v.hash = v8

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
	var v5 bool
	var v6 string
	var v7 string
	var v8 string

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,
			&v7,
			&v8,
		)
		if err != nil {
			return vv, err
		}

		v := &User{}
		v.ID = v0
		v.Login = v1
		v.Email = v2
		v.Avatar = v3
		v.Active = v4
		v.Admin = v5
		v.token = v6
		v.secret = v7
		v.hash = v8

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
	var v5 bool
	var v6 string
	var v7 string
	var v8 string

	v0 = v.ID
	v1 = v.Login
	v2 = v.Email
	v3 = v.Avatar
	v4 = v.Active
	v5 = v.Admin
	v6 = v.token
	v7 = v.secret
	v8 = v.hash

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,
		v8,
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
 user_id     INTEGER PRIMARY KEY AUTOINCREMENT
,user_login  TEXT
,user_email  TEXT
,user_avatar TEXT
,user_active BOOLEAN
,user_admin  BOOLEAN
,user_token  TEXT
,user_secret TEXT
,user_hash   TEXT
);
`

const InsertUserStmt = `
INSERT INTO users (
 user_login
,user_email
,user_avatar
,user_active
,user_admin
,user_token
,user_secret
,user_hash
) VALUES (?,?,?,?,?,?,?,?)
`

const SelectUserStmt = `
SELECT 
 user_id
,user_login
,user_email
,user_avatar
,user_active
,user_admin
,user_token
,user_secret
,user_hash
FROM users 
`

const SelectUserRangeStmt = `
SELECT 
 user_id
,user_login
,user_email
,user_avatar
,user_active
,user_admin
,user_token
,user_secret
,user_hash
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
,user_login
,user_email
,user_avatar
,user_active
,user_admin
,user_token
,user_secret
,user_hash
FROM users 
WHERE user_id=?
`

const UpdateUserPkeyStmt = `
UPDATE users SET 
 user_id=?
,user_login=?
,user_email=?
,user_avatar=?
,user_active=?
,user_admin=?
,user_token=?
,user_secret=?
,user_hash=? 
WHERE user_id=?
`

const DeleteUserPkeyStmt = `
DELETE FROM users 
WHERE user_id=?
`

const CreateUserLoginStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_login ON users (user_login)
`

const SelectUserLoginStmt = `
SELECT 
 user_id
,user_login
,user_email
,user_avatar
,user_active
,user_admin
,user_token
,user_secret
,user_hash
FROM users 
WHERE user_login=?
`

const UpdateUserLoginStmt = `
UPDATE users SET 
 user_id=?
,user_login=?
,user_email=?
,user_avatar=?
,user_active=?
,user_admin=?
,user_token=?
,user_secret=?
,user_hash=? 
WHERE user_login=?
`

const DeleteUserLoginStmt = `
DELETE FROM users 
WHERE user_login=?
`

const CreateUserEmailStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS user_email ON users (user_email)
`

const SelectUserEmailStmt = `
SELECT 
 user_id
,user_login
,user_email
,user_avatar
,user_active
,user_admin
,user_token
,user_secret
,user_hash
FROM users 
WHERE user_email=?
`

const UpdateUserEmailStmt = `
UPDATE users SET 
 user_id=?
,user_login=?
,user_email=?
,user_avatar=?
,user_active=?
,user_admin=?
,user_token=?
,user_secret=?
,user_hash=? 
WHERE user_email=?
`

const DeleteUserEmailStmt = `
DELETE FROM users 
WHERE user_email=?
`
