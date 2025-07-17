package utils

func ArrMap[I, O any](arr []I, cb func(I) O) []O {
	res := make([]O, len(arr))
	for index, item := range arr {
		res[index] = cb(item)
	}
	return res
}
