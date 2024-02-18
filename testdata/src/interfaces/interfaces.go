package interfaces

type TestInterface struct {
	I interface{} `json:"i"` // want "nullable field 'I' in struct 'TestInterface' must include 'omitempty' in its json tag to avoid marshaling as null"
}
