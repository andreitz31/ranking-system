package utils

import (
	"sort"
)

//Kv is a key-value structure
type Kv struct {
	Key   int
	Value int
}

// SortMap returns a sorted key-value slice
func SortMap(someMap map[int]int) []Kv {
	var keyValueList []Kv
	for k, v := range someMap {
		keyValueList = append(keyValueList, Kv{k, v})
	}
	sort.Slice(keyValueList, func(i, j int) bool {
		return keyValueList[i].Value > keyValueList[j].Value
	})
	return keyValueList
}
