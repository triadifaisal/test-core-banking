package helper

func ToPointer[T int | int32 | int64 | string | bool](item T) *T {
	return &item
}
