package parse

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	EncodeGzip = "gzip"
	EncodeJson = "json"
)

// Tag stores the parsed data from the tag string in
// a struct field.
type Tag struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Primary bool   `yaml:"pk"`
	Auto    bool   `yaml:"auto"`
	Index   string `yaml:"index"`
	Unique  string `yaml:"unique"`
	Size    int    `yaml:"size"`
	Skip    bool   `yaml:"skip"`
	Encode  string `yaml:"encode"`
}

// parseTag parses a tag string from the struct
// field and unmarshals into a Tag struct.
func parseTag(raw string) (*Tag, error) {
	var tag = new(Tag)

	raw = strings.Replace(raw, "`", "", -1)
	raw = reflect.StructTag(raw).Get("sql")

	// if the tag indicates the field should
	// be skipped we can exit right away.
	if strings.TrimSpace(raw) == "-" {
		tag.Skip = true
		return tag, nil
	}

	// otherwise wrap the string in curly braces
	// so that we can use the Yaml parser.
	raw = fmt.Sprintf("{ %s }", raw)

	// unmarshals the Yaml formatted string into
	// the Tag structure.
	var err = yaml.Unmarshal([]byte(raw), tag)
	return tag, err
}
