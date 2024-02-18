package nonnullables

type TestNonNullTypes struct {
	Int    int    `json:"int"`
	String string `json:"string"`
	Bool   bool   `json:"bool"`
	// Demonstrating both with and without `omitempty` for non-nullable types
	Float                  float64 `json:"float,omitempty"`
	NullAbleWithoutJsonTag []string
}
