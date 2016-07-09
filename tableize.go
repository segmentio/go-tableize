package tableize

import (
	"strings"

	. "github.com/segmentio/go-snakecase"
)

// Tableize the given map by flattening and normalizing all
// of the key/value pairs recursively.
func Tableize(m map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	visit(ret, m, "")
	return ret
}

// Visit map recursively and populate `ret`.
// We have to lowercase for now so that properties
// are mapped correctly to the schema fetched from
// redshift, as RS _always_ lowercases the column
// name in information_schema.columns.
func visit(ret map[string]interface{}, m map[string]interface{}, prefix string) {
	for key, val := range m {
		nKey := prefix + Snakecase(key)

		// Keys that start or end with underscores may result in weird behaviors
		// combined with the Snakecase function
		if strings.HasPrefix(key, "_") || strings.HasSuffix(key, "_") {
			if ret[nKey] != nil || m[nKey] != nil {
				continue
			}
		}

		if _, ok := val.(map[string]interface{}); ok {
			visit(ret, val.(map[string]interface{}), nKey+"_")
		} else {
			ret[nKey] = val
		}
	}
}
