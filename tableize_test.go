package tableize_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestTableizeArray(t *testing.T) {
	str := "{    \"context_client\": null,\n    \"context_latitude\": null,\n    \"context_location\": null,\n    \"context_longitude\": null,\n    \"data\": \"{\\\"via_zendesk\\\":true}\",\n    \"event_type\": \"FacebookComment\",\n    \"graph_object_id\": \"null\",\n    \"program\": \"source-runner\",\n    \"ticket_event_id\": 17820988764,\n    \"ticket_event_via\": \"Mail\",\n    \"ticket_id\": \"357276\",\n    \"timestamp\": \"2014-07-18T04:12:42.000Z\",\n    \"trusted\": true,\n    \"updater_id\": \"473752564\",\n    \"version\": \"8bc54c3\",\n    \"via\": {\n      \"channel\": \"email\",\n      \"source\": {\n        \"from\": {\n          \"address\": null,\n          \"name\": \"a b\",\n          \"original_recipients\": [\n            \"a@b.com\",\n            \"service@v.com.au\"\n          ]\n        },\n        \"rel\": null,\n        \"to\": {\n          \"address\": null,\n          \"name\": \"amaysim\"\n        }\n      }\n    }\n}"
	prop := make(map[string]interface{})
	var properties []byte
	properties = []byte(str)
	if len(properties) > 0 {
		dec := json.NewDecoder(bytes.NewReader(properties))
		dec.UseNumber()
		if err := dec.Decode(&prop); err != nil {
			fmt.Println(err.Error())
		}
	}
	//flatMap := Tableize(&Input{Value: prop})
	var arr []interface{}
	arr = append(arr,"a@b.com", "service@v.com.au")
	marshal, _ := json.Marshal(arr)//"[]interface {}[\"a@b.com\",\"service@v.com.au\"]"
	arrStr := string(marshal)
	spec := []struct {
		desc     string
		input    Input
		expected map[string]interface{}
	}{
		{
			desc: "array inside",
			input: Input{Value: prop/*flatMap*/, StringifyArr: true},
			expected: map[string]interface{}{
				"context_latitude" : nil,
				"context_location" : nil,
				"event_type" : "FacebookComment",
				"graph_object_id" : "null",
				"program" : "source-runner",
				"ticket_event_id" : json.Number("17820988764"),
				"ticket_event_via" : "Mail",
				"ticket_id" : "357276",
				"trusted" : true,
				"via_source_rel" : nil,
				"via_source_to_address" : nil,
				"context_client" : nil,
				"context_longitude" : nil,
				"data" : "{\"via_zendesk\":true}",
				"via_source_from_address" : nil,
				"via_source_from_name" : "a b",
				"via_source_from_original_recipients" : arrStr,
				"timestamp" : "2014-07-18T04:12:42.000Z",
				"updater_id" : "473752564",
				"version" : "8bc54c3",
				"via_channel" : "email",
				"via_source_to_name" : "amaysim",
			},
		},
		{
			desc: "simple arr StringifyArr=true",
			input: Input{Value: map[string]interface{}{
				"colors": []interface{}{"red", "blue"},
			}, StringifyArr: true,
			},
			expected: map[string]interface{}{
				"colors": "[\"red\",\"blue\"]",
			},
		},
		{
			desc: "simple arr StringifyArr=false",
			input: Input{Value: map[string]interface{}{
				"colors": []interface{}{"red", "blue"},
			}, StringifyArr: false,
			},
			expected: map[string]interface{}{
				"colors": []interface{}{"red", "blue"},
			},
		},
	}                                                                                          

	for _, test := range spec {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.expected, Tableize(&test.input))
		})
	}
}
