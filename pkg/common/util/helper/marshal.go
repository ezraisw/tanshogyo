package helper

func M(marshaller func(any) ([]byte, error), v any) ([]byte, error) {
	return marshaller(v)
}

func UM[T any](unmarshaler func([]byte, any) error, data []byte) (v T, err error) {
	err = unmarshaler(data, &v)
	return
}
