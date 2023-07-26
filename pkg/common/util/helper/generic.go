package helper

func ZeroOf[T any]() (_ T) { return }

func Ptr[T any](v T) *T {
	return &v
}

func Coalesce[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}

func CoalesceZero[T any](v *T) T {
	return Coalesce(v, ZeroOf[T]())
}
