package tableize

import snakecase "github.com/segmentio/go-snakecase"

type Input struct {
	Value map[string]interface{}

	// Optional
	HintSize      int
	Substitutions map[string]string
}

// Tableize the given map by flattening and normalizing all
// of the key/value pairs recursively.
func Tableize(in *Input) map[string]interface{} {
	if in.HintSize == 0 {
		in.HintSize = len(in.Value)
	}

	ret := make(map[string]interface{}, in.HintSize)
	visit(ret, in.Value, "", in.Substitutions)
	return ret
}

// Visit map recursively and populate `ret`.
// We have to lowercase for now so that properties
// are mapped correctly to the schema fetched from
// redshift, as RS _always_ lowercases the column
// name in information_schema.columns.
func visit(ret map[string]interface{}, m map[string]interface{}, prefix string, substitutions map[string]string) {
	for key, val := range m {
		if renamed, ok := substitutions[prefix+key]; ok {
			key = renamed
		}
		key = prefix + snakecase.Snakecase(key)
		if _, ok := val.(map[string]interface{}); ok {
			visit(ret, val.(map[string]interface{}), key+"_", substitutions)
		} else {
			ret[key] = val
		}
	}
}
