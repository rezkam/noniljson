# NonNilJson Linter

The `noniljson` linter is a static analysis tool for Go that ensures nullable fields in structs used for JSON marshaling include the `omitempty` option or are handled correctly to avoid marshaling as `null` in the JSON output. This tool is designed to help Go developers maintain clean and error-free JSON output by enforcing best practices for JSON serialization of Go structures.

## Examples

Consider the following Go struct:

```go
type ExampleStruct struct {
    Name    *string `json:"name"`
    Address string  `json:"address,omitempty"`
}
```

Running `noniljson` on this code will report an issue with the `Name` field for not including the `omitempty` option, since it's a pointer to a string and thus nullable.

