package tableize

import snakecase "github.com/segmentio/go-snakecase"

// Tableize the given map by flattening and normalizing all
// of the key/value pairs recursively.
func Tableize(m map[string]interface{}, hintSize ...int) map[string]interface{} {
	hint := len(m)
	if len(hintSize) > 0 {
		hint = hintSize[0]
	}
	ret := make(map[string]interface{}, hint)
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
		key = prefix + snakecase.Snakecase(key)
		if _, ok := val.(map[string]interface{}); ok {
			visit(ret, val.(map[string]interface{}), key+"_")
		} else {
			ret[key] = val
		}
	}
}
