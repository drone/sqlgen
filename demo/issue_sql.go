package demo

// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

import (
	"database/sql"
	"encoding/json"
)

func ScanIssue(row *sql.Row) (*Issue, error) {
	var v0 int64
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

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

	v := &Issue{}
	v.ID = v0
	v.Number = v1
	v.Title = v2
	v.Body = v3
	v.Assignee = v4
	v.State = v5
	json.Unmarshal(v6, &v.Labels)

	return v, nil
}

func ScanIssues(rows *sql.Rows) ([]*Issue, error) {
	var err error
	var vv []*Issue

	var v0 int64
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

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

		v := &Issue{}
		v.ID = v0
		v.Number = v1
		v.Title = v2
		v.Body = v3
		v.Assignee = v4
		v.State = v5
		json.Unmarshal(v6, &v.Labels)

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceIssue(v *Issue) []interface{} {
	var v0 int64
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

	v0 = v.ID
	v1 = v.Number
	v2 = v.Title
	v3 = v.Body
	v4 = v.Assignee
	v5 = v.State
	v6, _ = json.Marshal(&v.Labels)

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

func SelectIssue(db *sql.DB, query string, args ...interface{}) (*Issue, error) {
	row := db.QueryRow(query, args...)
	return ScanIssue(row)
}

func SelectIssues(db *sql.DB, query string, args ...interface{}) ([]*Issue, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanIssues(rows)
}

func InsertIssue(db *sql.DB, query string, v *Issue) error {

	res, err := db.Exec(query, SliceIssue(v)[1:]...)
	if err != nil {
		return err
	}

	v.ID, err = res.LastInsertId()
	return err
}

func UpdateIssue(db *sql.DB, query string, v *Issue) error {

	args := SliceIssue(v)[1:]
	args = append(args, v.ID)
	_, err := db.Exec(query, args...)
	return err
}

const CreateIssueStmt = `
CREATE TABLE IF NOT EXISTS issues (
 issue_id       SERIAL PRIMARY KEY 
,issue_number   INTEGER
,issue_title    VARCHAR(512)
,issue_body     VARCHAR(2048)
,issue_assignee VARCHAR(512)
,issue_state    VARCHAR(50)
,issue_labels   BYTEA
);
`

const InsertIssueStmt = `
INSERT INTO issues (
 issue_number
,issue_title
,issue_body
,issue_assignee
,issue_state
,issue_labels
) VALUES ($1,$2,$3,$4,$5,$6)
`

const SelectIssueStmt = `
SELECT 
 issue_id
,issue_number
,issue_title
,issue_body
,issue_assignee
,issue_state
,issue_labels
FROM issues 
`

const SelectIssueRangeStmt = `
SELECT 
 issue_id
,issue_number
,issue_title
,issue_body
,issue_assignee
,issue_state
,issue_labels
FROM issues 
LIMIT $1 OFFSET $2
`

const SelectIssueCountStmt = `
SELECT count(1)
FROM issues 
`

const SelectIssuePkeyStmt = `
SELECT 
 issue_id
,issue_number
,issue_title
,issue_body
,issue_assignee
,issue_state
,issue_labels
FROM issues 
WHERE issue_id=$1
`

const UpdateIssuePkeyStmt = `
UPDATE issues SET 
 issue_id=$1
,issue_number=$2
,issue_title=$3
,issue_body=$4
,issue_assignee=$5
,issue_state=$6
,issue_labels=$7 
WHERE issue_id=$8
`

const DeleteIssuePkeyStmt = `
DELETE FROM issues 
WHERE issue_id=$1
`
