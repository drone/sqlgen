package parse

import (
	"reflect"
	"testing"
)

var tagTests = []struct {
	raw string
	tag *Tag
}{
	{
		`sql:"-"`,
		&Tag{Skip: true},
	},
	{
		`sql:"pk: true, auto: true"`,
		&Tag{Primary: true, Auto: true},
	},
	{
		`sql:"name: foo"`,
		&Tag{Name: "foo"},
	},
	{
		`sql:"type: varchar"`,
		&Tag{Type: "varchar"},
	},
	{
		`sql:"size: 2048"`,
		&Tag{Size: 2048},
	},
	{
		`sql:"index: fake_index"`,
		&Tag{Index: "fake_index"},
	},
	{
		`sql:"unique: fake_unique_index"`,
		&Tag{Unique: "fake_unique_index"},
	},
}

func TestParseTag(t *testing.T) {
	for _, test := range tagTests {

		var want = test.tag
		var got, err = parseTag(test.raw)

		if err != nil {
			t.Errorf("Got Error parsing Tag %s. %s", test.raw, err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted Tag %+v, got Tag %+v", want, got)
		}
	}
}
