package pointers

type TestPointer struct {
	Ptr *string `json:"ptr"` // want "nullable field 'Ptr' in struct 'TestPointer' must include 'omitempty' in its json tag to avoid marshaling as null"
}
