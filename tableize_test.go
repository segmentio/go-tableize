package tableize_test

import (
	"encoding/json"
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

func BenchmarkSmall(b *testing.B) {
	str := `{
	  "anonymousId": "b2e9efda-4fc1-4bfa-8dc8-95ce56d8f53e",
	  "projectId": "gnv5tty0m6",
	  "properties": {
	    "path": "/_generated_background_page.html",
	    "referrer": "",
	    "search": "",
	    "title": "",
	    "url": "chrome-extension://djkmaiagjpalbehljdggfaebgmgioobl/_generated_background_page.html"
	  },
	  "receivedAt": "2014-05-13T20:28:52.803Z",
	  "requestId": "e894944b-9afe-49c0-a782-1e0cfc68fe48",
	  "timestamp": "2014-05-13T20:28:50.540Z",
	  "type": "page",
	  "userId": "6fac5180-b4d5-4305-a210-a1674bb3af4b",
	  "version": 2
	}`

	event := make(map[string]interface{})
	check(json.Unmarshal([]byte(str), &event))

	for i := 0; i < b.N; i++ {
		Tableize(&Input{Value: event})
	}
}

func BenchmarkMedium(b *testing.B) {
	str := `{
	  "anonymousId": "b2e9efda-4fc1-4bfa-8dc8-95ce56d8f53e",
	  "channel": "client",
	  "context": {
	    "ip": "67.208.188.98",
	    "library": {
	      "name": "analytics.js",
	      "version": "unknown"
	    },
	    "userAgent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.131 Safari/537.36"
	  },
	  "projectId": "gnv5tty0m6",
	  "properties": {
	    "path": "/_generated_background_page.html",
	    "referrer": "",
	    "search": "",
	    "title": "",
	    "url": "chrome-extension://djkmaiagjpalbehljdggfaebgmgioobl/_generated_background_page.html"
	  },
	  "receivedAt": "2014-05-13T20:28:52.803Z",
	  "requestId": "e894944b-9afe-49c0-a782-1e0cfc68fe48",
	  "timestamp": "2014-05-13T20:28:50.540Z",
	  "type": "page",
	  "userId": "6fac5180-b4d5-4305-a210-a1674bb3af4b",
	  "version": 2
	}`

	event := make(map[string]interface{})
	check(json.Unmarshal([]byte(str), &event))

	for i := 0; i < b.N; i++ {
		Tableize(&Input{Value: event})
	}
}

func BenchmarkLarge(b *testing.B) {
	str := `{
	  "anonymousId": "b2e9efda-4fc1-4bfa-8dc8-95ce56d8f53e",
	  "channel": "client",
	  "context": {
	    "ip": "67.208.188.98",
	    "library": {
	      "name": "analytics.js",
	      "version": "unknown"
	    },
	    "userAgent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.131 Safari/537.36"
	  },
	  "projectId": "gnv5tty0m6",
	  "properties": {
	    "path": "/_generated_background_page.html",
	    "referrer": "",
	    "search": "",
	    "title": "",
	    "url": "chrome-extension://djkmaiagjpalbehljdggfaebgmgioobl/_generated_background_page.html",
	    "anonymousId": "b2e9efda-4fc1-4bfa-8dc8-95ce56d8f53e",
		  "channel": "client",
		  "context": {
		    "ip": "67.208.188.98",
		    "library": {
		      "name": "analytics.js",
		      "version": "unknown"
		    },
		    "userAgent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.131 Safari/537.36"
		  },
		  "projectId": "gnv5tty0m6",
		  "properties": {
		    "path": "/_generated_background_page.html",
		    "referrer": "",
		    "search": "",
		    "title": "",
		    "url": "chrome-extension://djkmaiagjpalbehljdggfaebgmgioobl/_generated_background_page.html",
		    "anonymousId": "b2e9efda-4fc1-4bfa-8dc8-95ce56d8f53e",
			  "channel": "client",
			  "context": {
			    "ip": "67.208.188.98",
			    "library": {
			      "name": "analytics.js",
			      "version": "unknown"
			    },
			    "userAgent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.131 Safari/537.36"
			  },
			  "projectId": "gnv5tty0m6",
			  "properties": {
			    "path": "/_generated_background_page.html",
			    "referrer": "",
			    "search": "",
			    "title": "",
			    "url": "chrome-extension://djkmaiagjpalbehljdggfaebgmgioobl/_generated_background_page.html"
			  },
			  "receivedAt": "2014-05-13T20:28:52.803Z",
			  "requestId": "e894944b-9afe-49c0-a782-1e0cfc68fe48",
			  "timestamp": "2014-05-13T20:28:50.540Z",
			  "type": "page",
			  "userId": "6fac5180-b4d5-4305-a210-a1674bb3af4b",
			  "version": 2
		  },
		  "receivedAt": "2014-05-13T20:28:52.803Z",
		  "requestId": "e894944b-9afe-49c0-a782-1e0cfc68fe48",
		  "timestamp": "2014-05-13T20:28:50.540Z",
		  "type": "page",
		  "userId": "6fac5180-b4d5-4305-a210-a1674bb3af4b",
		  "version": 2
	  },
	  "receivedAt": "2014-05-13T20:28:52.803Z",
	  "requestId": "e894944b-9afe-49c0-a782-1e0cfc68fe48",
	  "timestamp": "2014-05-13T20:28:50.540Z",
	  "type": "page",
	  "userId": "6fac5180-b4d5-4305-a210-a1674bb3af4b",
	  "version": 2
	}`

	event := make(map[string]interface{})
	check(json.Unmarshal([]byte(str), &event))

	for i := 0; i < b.N; i++ {
		Tableize(&Input{Value: event})
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
