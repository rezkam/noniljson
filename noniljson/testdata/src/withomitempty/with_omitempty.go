package withomitempty

type TestOmitEmpty struct {
	Ptr   *string        `json:"ptr,omitempty"`
	Slice []int          `json:"slice,omitempty"`
	Map   map[string]int `json:"map,omitempty"`
	I     interface{}    `json:"i,omitempty"`
}
