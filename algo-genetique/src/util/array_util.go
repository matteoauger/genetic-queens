package util


func Find(source []int, x int) int {
	for i, n := range source {
			if x == n {
					return i
			}
	}
	return -1
}
