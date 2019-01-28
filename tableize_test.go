package tableize_test

import (
	"testing"

	. "github.com/segmentio/go-tableize"
	"github.com/stretchr/testify/assert"
)

func TestTableize(t *testing.T) {
	spec := []struct {
		desc     string
		input    Input
		expected map[string]interface{}
	}{
		{
			desc: "simple",
			input: Input{Value: map[string]interface{}{
				"species": "ferret",
			}},
			expected: map[string]interface{}{
				"species": "ferret",
			},
		},
		{
			desc: "nested",
			input: Input{
				Value: map[string]interface{}{
					"name": map[string]interface{}{
						"first": "tobi",
					},
				},
			},
			expected: map[string]interface{}{"name_first": "tobi"},
		},
		{
			desc: "normalize keys",
			input: Input{
				Value: map[string]interface{}{
					"first   name  ": "tobi",
					"Last-Name":      "holowaychuk",
					"NickName":       "shupa",
					"$some_thing":    "tobi",
				},
			},
			expected: map[string]interface{}{
				"first_name": "tobi",
				"last_name":  "holowaychuk",
				"nick_name":  "shupa",
				"some_thing": "tobi",
			},
		},
		{
			desc: "conflicting keys",
			input: Input{
				Value: map[string]interface{}{
					"firstName":   "first_1",
					"first_name":  "first_2",
					"lastName":    "last_1",
					"last_name":   "last_2",
					"middleName":  "middle_1",
					"middle_name": "middle_2",
					"first":       map[string]interface{}{"name": "first_3"},
					"last":        map[string]interface{}{"name": "last_3"},
				},
			},
			expected: map[string]interface{}{
				"first_name":  "first_2",
				"last_name":   "last_2",
				"middle_name": "middle_2",
			},
		},
		{
			desc: "substitutions",
			input: Input{
				Value: map[string]interface{}{
					"name": map[string]interface{}{
						"first   name  ": "tobi",
						"Last-Name":      "holowaychuk",
						"NickName":       "shupa",
						"$some_thing":    "tobi",
					},
					"_mid":    "value",
					"species": "ferret",
				},
				Substitutions: map[string]string{
					"species":          "r_species",
					"name_$some_thing": "just_some_thing",
					"_mid":             "u_mid",
				},
			},
			expected: map[string]interface{}{
				"name_first_name":      "tobi",
				"name_last_name":       "holowaychuk",
				"r_species":            "ferret",
				"name_nick_name":       "shupa",
				"name_just_some_thing": "tobi",
				"u_mid":                "value",
			},
		},
	}

	for _, test := range spec {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.expected, Tableize(&test.input))
		})
	}
}
