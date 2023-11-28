package slices

// To 转换为其他类型
func To[T, R any](list []T, f func(T) R) []R {
	rs := make([]R, 0, len(list))
	for _, v := range list {
		rs = append(rs, f(v))
	}
	return rs
}

// Index 查找元素的索引
func Index[T comparable](list []T, v T) int {
	for i, item := range list {
		if item == v {
			return i
		}
	}
	return -1
}

// IndexOf 查找元素的索引
func IndexOf[T any](list []T, f func(T) bool) int {
	for i, item := range list {
		if f(item) {
			return i
		}
	}
	return -1
}

// Contains 是否包含
func Contains[T comparable](list []T, v T) bool {
	return Index(list, v) != -1
}

// ContainsOf 是否包含
func ContainsOf[T any](list []T, f func(T) bool) bool {
	return IndexOf(list, f) != -1
}
