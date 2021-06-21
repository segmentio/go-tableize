package tableize

import (
	"log"
	"sort"

	"github.com/segmentio/encoding/json"
	snakecase "github.com/segmentio/go-snakecase"
)

// Input contains all information to be tableized
type Input struct {
	Value map[string]interface{}

	// Optional
	HintSize        int
	Substitutions   map[string]string
	StringifyArrays bool
}

// Tableize the given map by flattening and normalizing all
// of the key/value pairs recursively.
func Tableize(in *Input) map[string]interface{} {
	if in.HintSize == 0 {
		in.HintSize = len(in.Value)
	}

	ret := make(map[string]interface{}, in.HintSize)
	visit(ret, in.Value, "", in.Substitutions, in.StringifyArrays)
	return ret
}

// Visit map recursively and populate `ret`.
// We have to lowercase for now so that properties
// are mapped correctly to the schema fetched from
// redshift, as RS _always_ lowercases the column
// name in information_schema.columns.
func visit(ret map[string]interface{}, m map[string]interface{}, prefix string, substitutions map[string]string, stringifyArrays bool) {
	var val interface{}
	var renamed string
	var ok bool
	keys := getSortedKeys(m)

	for _, key := range keys {
		val = m[key]
		if len(substitutions) > 0 {
			if renamed, ok = substitutions[prefix+key]; ok {
				key = renamed
			}
		}
		key = prefix + snakecase.Snakecase(key)

		switch t := val.(type) {
		case map[string]interface{}:
			visit(ret, t, key+"_", substitutions, stringifyArrays)
		case []interface{}:
			if stringifyArrays {
				valByteArr, err := json.Marshal(val)
				if err != nil {
					log.Printf("go-tableize: dropping array value %+v that could not be converted to string: %s\n", val, err)
				} else {
					ret[key] = string(valByteArr)
				}
			} else {
				ret[key] = val
			}
		default:
			ret[key] = val
		}
	}
}

func getSortedKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
