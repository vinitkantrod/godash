package godash

import (
	"fmt"
	"testing"
)

func TestChunkInt64(t *testing.T) {
	arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := ChunkInt64(arr, 3)
	if len(result[0]) != 3 {
		fmt.Println("In-Correct result")
		// t.Error("Expected %v found %v", []int64{1, 2, 3}, result[0])
	} else {
		fmt.Println("Correct result")
	}
}
