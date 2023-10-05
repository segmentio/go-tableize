
# tablelize

> **Note**  
> Segment has paused maintenance on this project, but may return it to an active status in the future. Issues and pull requests from external contributors are not being considered, although internal contributions may appear from time to time. The project remains available under its open source license for anyone to use.

 Go tableize the given map by recursively walking the map and normalizing
 its keys to produce a flat SQL-friendly map.

## CLI

  ```bash
  $ go get github.com/segmentio/go-tableize/cmd/tableize
  $ echo '{"user": { "id": 1 }}' | tableize
  { "user_id": 1 }
  ```

## Example

```go
event := map[string]interface{}{
  "name": map[string]interface{}{
    "first   name  ": "tobi",
    "last-name":      "holowaychuk",
  },
  "species": "ferret",
}

flat := Tableize(event)
assert(t, flat["name_first_name"] == "tobi")
assert(t, flat["name_last_name"] == "holowaychuk")
assert(t, flat["species"] == "ferret")
```
