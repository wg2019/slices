package slices

// IntValue 获取整型切片内容.
func IntValue(slice []int, index int) int {
	if index < 0 {
		return 0
	}
	count := len(slice)
	if index < count {
		return slice[index]
	}
	return 0
}

// StringValue 获取字符串切片内容.
func StringValue(slice []string, index int) string {
	if index < 0 {
		return ""
	}
	count := len(slice)
	if index < count {
		return slice[index]
	}
	return ""
}
