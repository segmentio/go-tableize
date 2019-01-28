package tableize_test

import (
	"encoding/json"
	"testing"

	. "github.com/segmentio/go-tableize"
)

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
