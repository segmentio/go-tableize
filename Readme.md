
# tablelize

 Go tableize the given map by recursively walking the map and normalizing
 its keys to produce a flat SQL-friendly map.

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