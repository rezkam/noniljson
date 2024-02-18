package maps

type TestMap struct {
	Map map[string]int `json:"map"` // want "nullable field 'Map' in struct 'TestMap' must include 'omitempty' in its json tag to avoid marshaling as null"
}
