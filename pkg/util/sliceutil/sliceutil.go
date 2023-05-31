package sliceutil

import (
	"strings"
)

// RemoveFromStringSlice removes the first occurrence of the specified element
// from this list, if it is present.
// If the list does not contain the element, it is unchanged.
func RemoveFromStringSlice(s []string, e string) []string {
	i := IndexOfStringSlice(s, e)
	if i < 0 {
		return s
	}

	sl := make([]string, len(s)-1)
	copy(sl, append(s[:i], s[i+1:]...))
	return sl
}

// RemoveDuplicateFromStringSlice removes duplicate elements
// of the string slice.
func RemoveDuplicateFromStringSlice(s []string) []string {
	m := make(map[string]struct{})
	sl := make([]string, 0)

	for _, e := range s {
		if _, ok := m[e]; !ok {
			sl = append(sl, e)
			m[e] = struct{}{}
		}
	}
	return sl
}

// RemoveDuplicateFromIntSlice removes duplicate elements
// of the integer slice. 去重
func RemoveDuplicateFromIntSlice(s []int) []int {
	m := make(map[int]struct{})
	sl := make([]int, 0)

	for _, e := range s {
		if _, ok := m[e]; !ok {
			sl = append(sl, e)
			m[e] = struct{}{}
		}
	}
	return sl
}

// IndexOfStringSlice returns the index of the first occurrence of the specified
// element in this list, or -1 if this list does not contain the element.
func IndexOfStringSlice(s []string, e string) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

func ContainsInStringSlice(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func ContainsInStringSliceIgnoreCase(s []string, e string) bool {
	for _, v := range s {
		if strings.ToLower(v) == strings.ToLower(e) {
			return true
		}
	}
	return false
}

func ContainsInIntSlice(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func ContainsInInt32Slice(s []int32, e int32) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// IntersectionIntSlice 用于获取两个 Int 类型切片的交集
func IntersectionIntSlice(i1 []int, i2 []int) []int {
	if len(i1) == 0 || len(i2) == 0 {
		return []int{}
	}

	var result []int
	for _, v1 := range i1 {
		canAdd := false
		for _, v2 := range i2 {
			if v1 == v2 {
				canAdd = true
			}
		}
		if canAdd {
			result = append(result, v1)
		}
	}
	return result
}

// 删除 int 类型切片中的元素
func RemoveValueFromIntSlice(a []int, rmVal int) []int {
	var result []int
	for _, value := range a {
		if rmVal != value {
			result = append(result, value)
		}
	}
	return result
}

// GetComplementaryAndIntersection 取 oldIds 不在 newIds 的id，取 newIds 不在 oldIds 的id，取交集
func GetComplementaryAndIntersection(oldIds []int, newIds []int) (needToDeleteIds, needToCreateIds, needToUpdateIds []int) {
	oldLevelIdMap := make(map[int]bool)
	for _, oldLevelId := range oldIds {
		oldLevelIdMap[oldLevelId] = true
	}
	inputLevelIdMap := make(map[int]bool)
	for _, inputLevelId := range newIds {
		inputLevelIdMap[inputLevelId] = true
	}
	for _, oldLevelId := range oldIds {
		if _, ok := inputLevelIdMap[oldLevelId]; !ok {
			needToDeleteIds = append(needToDeleteIds, oldLevelId)
		}
	}
	for _, inputLevelId := range newIds {
		if _, ok := oldLevelIdMap[inputLevelId]; !ok {
			needToCreateIds = append(needToCreateIds, inputLevelId)
		} else {
			needToUpdateIds = append(needToUpdateIds, inputLevelId)
		}
	}
	return
}

// SplitIntoChunk 会将 items 切片切割成块，每一块的大小上限是 chunkSize
func SplitIntoChunk[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
