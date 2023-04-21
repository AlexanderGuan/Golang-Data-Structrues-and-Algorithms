package slice

// 迭代遍历数组
func traverse(arr []int) {
	for idx, val := range arr {
		println(idx, val)
	}
}

// 递归遍历数组
func traverseRecursive(arr []int, i int) {
	if i == len(arr) {
		return
	}
	// 前序位置
	traverseRecursive(arr, i+1)
	// 后序位置
}
