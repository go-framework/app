package configurer

type Marshaler interface {
	Marshal(map[string]interface{}) ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal([]byte) (map[string]interface{}, error)
}