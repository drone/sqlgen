package demo

// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

import (
	"database/sql"
)

func ScanLegacyUser(row *sql.Row) (*LegacyUser, error) {
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

	v := &LegacyUser{}
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

func ScanLegacyUsers(rows *sql.Rows) ([]*LegacyUser, error) {
	var err error
	var vv []*LegacyUser

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

		v := &LegacyUser{}
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

func SliceLegacyUser(v *LegacyUser) []interface{} {
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

func SelectLegacyUser(db *sql.DB, query string, args ...interface{}) (*LegacyUser, error) {
	row := db.QueryRow(query, args...)
	return ScanLegacyUser(row)
}

func SelectLegacyUsers(db *sql.DB, query string, args ...interface{}) ([]*LegacyUser, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanLegacyUsers(rows)
}

func InsertLegacyUser(db *sql.DB, query string, v *LegacyUser) error {

	res, err := db.Exec(query, SliceLegacyUser(v)[1:]...)
	if err != nil {
		return err
	}

	v.ID, err = res.LastInsertId()
	return err
}

func UpdateLegacyUser(db *sql.DB, query string, v *LegacyUser) error {

	args := SliceLegacyUser(v)[1:]
	args = append(args, v.ID)
	_, err := db.Exec(query, args...)
	return err
}

const CreateUserStmt = `
CREATE TABLE IF NOT EXISTS users (
 id     INTEGER PRIMARY KEY AUTOINCREMENT
,login  TEXT
,email  TEXT
,avatar TEXT
,active BOOLEAN
,admin  BOOLEAN
,token  TEXT
,secret TEXT
,hash   TEXT
);
`

const InsertUserStmt = `
INSERT INTO users (
 login
,email
,avatar
,active
,admin
,token
,secret
,hash
) VALUES (?,?,?,?,?,?,?,?)
`

const SelectUserStmt = `
SELECT 
 id
,login
,email
,avatar
,active
,admin
,token
,secret
,hash
FROM users 
`

const SelectUserRangeStmt = `
SELECT 
 id
,login
,email
,avatar
,active
,admin
,token
,secret
,hash
FROM users 
LIMIT ? OFFSET ?
`

const SelectUserCountStmt = `
SELECT count(1)
FROM users 
`

const SelectUserPkeyStmt = `
SELECT 
 id
,login
,email
,avatar
,active
,admin
,token
,secret
,hash
FROM users 
WHERE id=?
`

const UpdateUserPkeyStmt = `
UPDATE users SET 
 id=?
,login=?
,email=?
,avatar=?
,active=?
,admin=?
,token=?
,secret=?
,hash=? 
WHERE id=?
`

const DeleteUserPkeyStmt = `
DELETE FROM users 
WHERE id=?
`

const CreateLoginStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS login ON users (login)
`

const SelectLoginStmt = `
SELECT 
 id
,login
,email
,avatar
,active
,admin
,token
,secret
,hash
FROM users 
WHERE login=?
`

const UpdateLoginStmt = `
UPDATE users SET 
 id=?
,login=?
,email=?
,avatar=?
,active=?
,admin=?
,token=?
,secret=?
,hash=? 
WHERE login=?
`

const DeleteLoginStmt = `
DELETE FROM users 
WHERE login=?
`

const CreateEmailStmt = `
CREATE UNIQUE INDEX IF NOT EXISTS email ON users (email)
`

const SelectEmailStmt = `
SELECT 
 id
,login
,email
,avatar
,active
,admin
,token
,secret
,hash
FROM users 
WHERE email=?
`

const UpdateEmailStmt = `
UPDATE users SET 
 id=?
,login=?
,email=?
,avatar=?
,active=?
,admin=?
,token=?
,secret=?
,hash=? 
WHERE email=?
`

const DeleteEmailStmt = `
DELETE FROM users 
WHERE email=?
`
