package util

func ToSliceAny[T any](objs []T) []any {
	var datas []any
	for _, obj := range objs {
		datas = append(datas, obj)
	}
	return datas
}
