package godash

func max(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

func ChunkInt64(array []int64, size int32) [][]int64 {
	if size == 0 {
		size = 1
	} else {
		size = max(size, 1)
	}
	var length int32
	if len(array) > 0 {
		length = int32(len(array))
	} else {
		length = 0
	}
	var newArray [][]int64
	if length == 0 || size < 1 {
		return newArray
	}
	var index int32 = 0
	var resIndex int32 = 0
	var newArrayLength = ((length / size) + 1)
	result := make([][]int64, newArrayLength)
	for index < length {
		tempArray := array[index:(index + size)]
		result = append(result, tempArray)
		index = index + size
		resIndex = resIndex + 1
	}
	return result
}
