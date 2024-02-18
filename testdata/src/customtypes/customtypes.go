package customtypes

type MyCustomSlice []int

type TestCustomType struct {
	Custom MyCustomSlice `json:"custom"` // want "nullable field 'Custom' in struct 'TestCustomType' must include 'omitempty' in its json tag to avoid marshaling as null"
}
