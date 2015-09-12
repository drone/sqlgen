package main

import (
	"fmt"
	"io"
	"strings"

	"bitbucket.org/pkg/inflect"
	"github.com/drone/sqlgen/schema"
)

// writeSchema writes SQL statements to CREATE, INSERT,
// UPDATE and DELETE values from Table t.
func writeSchema(w io.Writer, d schema.Dialect, t *schema.Table) {

	writeConst(w,
		d.Table(t),
		"create", inflect.Singularize(t.Name),
	)

	writeConst(w,
		d.Insert(t),
		"insert", inflect.Singularize(t.Name),
	)

	writeConst(w,
		d.Select(t, nil),
		"select", "all", t.Name,
	)

	writeConst(w,
		d.SelectRange(t, nil),
		"select", inflect.Singularize(t.Name), "range",
	)

	writeConst(w,
		d.SelectCount(t, nil),
		"select", inflect.Singularize(t.Name), "count",
	)

	if len(t.Primary) != 0 {
		writeConst(w,
			d.Select(t, t.Primary),
			"select", inflect.Singularize(t.Name), "primary", "key",
		)

		writeConst(w,
			d.Update(t, t.Primary),
			"update", inflect.Singularize(t.Name), "primary", "key",
		)

		writeConst(w,
			d.Delete(t, t.Primary),
			"delete", inflect.Singularize(t.Name), "primary", "key",
		)
	}

	for _, ix := range t.Index {

		writeConst(w,
			d.Index(t, ix),
			"create", ix.Name,
		)

		writeConst(w,
			d.Select(t, ix.Fields),
			"select", ix.Name,
		)

		if !ix.Unique {

			writeConst(w,
				d.SelectRange(t, ix.Fields),
				"select", ix.Name, "range",
			)

			writeConst(w,
				d.SelectCount(t, ix.Fields),
				"select", ix.Name, "count",
			)

		} else {

			writeConst(w,
				d.Update(t, ix.Fields),
				"update", ix.Name,
			)

			writeConst(w,
				d.Delete(t, ix.Fields),
				"delete", ix.Name,
			)
		}
	}
}

// WritePackage writes the Go package header to
// writer w with the given package name.
func writePackage(w io.Writer, name string) {
	fmt.Fprintf(w, sPackage, name)
}

// writeConst is a helper function that writes the
// body string to a Go const variable.
func writeConst(w io.Writer, body string, label ...string) {
	// create a snake case variable name from
	// the specified labels. Then convert the
	// variable name to a quoted, camel case string.
	name := strings.Join(label, "_")
	name = inflect.Typeify(name)

	// quote the body using multi-line quotes
	body = fmt.Sprintf(sQuote, body)

	fmt.Fprintf(w, sConst, name, body)
}
