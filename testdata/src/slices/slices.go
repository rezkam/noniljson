package slices

type TestSlice struct {
	Slice []string `json:"slice"` // want "nullable field 'Slice' in struct 'TestSlice' must include 'omitempty' in its json tag to avoid marshaling as null"
}
