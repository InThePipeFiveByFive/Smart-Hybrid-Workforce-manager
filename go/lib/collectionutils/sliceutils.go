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

// This method calculates the intersection between two generic slices
func SliceIntersection[T comparable](slice1, slice2 []T) []T {
	intersectionMap := make(map[T]bool)
	for _, i := range slice1 {
		intersectionMap[i] = false
	}
	intersection := []T{}
	for _, i := range slice2 {
		if _, ok := intersectionMap[i]; ok {
			intersection = append(intersection, i)
		}
	}
	return intersection
}

// Removes the ONLY THE FIRST matching element from the slice without preserving ordering
func RemoveElementNoOrder[T comparable](slice []T, element T) []T {
	for i := range slice {
		if slice[i] == element {
			slice[i] = slice[len(slice)-1] // replace with last element
			return slice[:len(slice)-1]
		}
	}
	return slice // not found
}

// Removes the element at the specified indexm without preserving order. If the element does not exist, nothing is done
func RemElemenAtI[T any](slice []T, index int) []T {
	if index >= len(slice) || index < 0 {
		return slice
	}
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Contains slice
// Linear search (O(n) complexity)
func Contains[T comparable](s []T, e T) bool {
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

// Partitions a 1D into a 2D array, according to the sizes in the sizez array
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

func Copy1DArr[T any](arr []T) []T {
	copied := make([]T, len(arr))
	copy(copied, arr)
	return copied
}

func Copy2DArr[T any](arr [][]T) [][]T {
	copied := make([][]T, len(arr))
	for i := 0; i < len(arr); i++ {
		copied[i] = Copy1DArr(arr[i])
	}
	return copied
}
