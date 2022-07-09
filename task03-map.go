package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {

	// new slice for keys of input map
	keys := make([]int, 0, len(input))
	// populate the slice by keys
	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	// populate the result slice by map values in the ascending key order
	for _, k := range keys {
		result = append(result, input[k])
	}

	return
}
