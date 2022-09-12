package collectionutils

// SequentialSequence returns a sequence of sequential integers in the range [start, end)
// no numbers returned if start >= end
func SequentialSequence(start, end int) []int {
	if start >= end {
		return []int{}
	}
	seq := make([]int, end-start)
	for i := range seq {
		seq[i] = i + start
	}
	return seq
}

// This method finds the intersection between two integer slices
func IntSliceIntersection(slice1, slice2 []int) []int {
	// Gets the intersection of two integer slices
	intersectionMap := make(map[int]bool)
	for _, i := range slice1 {
		intersectionMap[i] = false
	}
	intersection := []int{}
	for _, i := range slice2 {
		if _, ok := intersectionMap[i]; ok {
			intersection = append(intersection, i)
		}
	}
	return intersection
}

// Removes the ONLY THE FIRST matching element from the slice without preserving ordering
func RemoveElementNoOrder(slice []int, element int) []int {
	for i := range slice {
		if slice[i] == element {
			slice[i] = slice[len(slice)-1] // replace with last element
			return slice[:len(slice)-1]
		}
	}
	return slice // not found
}

// Removes the element at the specified index. If the element does not exist, nothing is done
func RemElemenAtI(slice []int, index int) []int {
	if index >= len(slice) {
		return slice
	}
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Contains slice
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Flattens a 2D array
// Does not copy elements
func Flatten2DArr[T any](arr [][]T) []T {
	totalLen := 0
	for i := 0; i < len(arr); i++ {
		totalLen += len(arr[i])
	}
	flattenedArr := make([]T, totalLen)
	flatIndex := 0
	for _, subArr := range arr {
		for _, el := range subArr {
			flattenedArr[flatIndex] = el
			flatIndex++
		}
	}
	return flattenedArr
}

func PartitionArray[T any](arr []T, sizes []int) [][]T {
	partitionedArray := make([][]T, len(sizes))
	flatIndex := 0
	for i := 0; i < len(sizes); i++ {
		partitionedArray[i] = make([]T, sizes[i])
		for j := 0; j < sizes[i]; j++ {
			partitionedArray[i][j] = arr[flatIndex]
			flatIndex++
		}
	}
	return partitionedArray
}
