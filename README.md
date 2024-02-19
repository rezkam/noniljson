# NonNilJson Linter

The `noniljson` linter is a static analysis tool for Go that ensures nullable fields in structs used for JSON marshaling include the `omitempty` option or are handled correctly to avoid marshaling as `null` in the JSON output. This tool is designed to help Go developers maintain clean and error-free JSON output by enforcing best practices for JSON serialization of Go structures.

## Why `noniljson`?

In Go, when struct fields are marshaled into JSON, fields with `nil` values are included in the output as `null`. This behavior can lead to several issues:

- **Unexpected `null` values in JSON output**: Consumers of your JSON API may not expect or handle `null` values correctly, leading to errors or inconsistencies in data processing.
- **Data Sparsity**: Output JSON might include a lot of `null` values for optional fields, leading to larger payloads and inefficient data transmission, especially in microservices or APIs where bandwidth and payload size are concerns.
- **API Contract Violation**: If your API specification defines optional fields should be omitted if not present, including `nil` values as `null` in the output violates this contract.

The `noniljson` linter addresses these issues by enforcing the `omitempty` option for nullable struct fields. This option ensures that fields with `nil` values are omitted from the marshaled JSON, resulting in cleaner, more efficient, and spec-compliant JSON output.

By integrating `noniljson` into your Go projects, you can automatically enforce best practices for JSON serialization across your codebase, improving the quality and reliability of your APIs and data interfaces.

### Use Cases

- **API Development**: Ensuring your JSON endpoints output clean, compliant JSON.
- **Microservices**: Reducing payload size for efficient inter-service communication.
- **Data Storage**: Preparing structs for marshaling into JSON databases or files without spurious `null` values.

In essence, `noniljson` helps maintain high-quality code standards, ensuring that your Go applications produce optimal JSON output, enhancing compatibility and efficiency in data exchange.

## Examples

Consider the following Go struct:

```go
type ExampleStruct struct {
    Name    *string `json:"name"`
    Address string  `json:"address,omitempty"`
}
```

Running `noniljson` on this code will report an issue with the `Name` field for not including the `omitempty` option, since it's a pointer to a string and thus nullable.
