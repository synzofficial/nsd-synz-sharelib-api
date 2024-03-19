package typeconvertutil

func ToPtr[T any](in T) *T {
	return &in
}

func PtrValOrDefault[T any](in *T, defaultVal T) *T {
	if in != nil {
		return in
	}
	return &defaultVal
}
