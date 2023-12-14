package utils

func Permutations(arr []interface{}) [][]interface{} {
	var helper func([]interface{}, int)
	var res [][]interface{}

	helper = func(arr []interface{}, n int) {
		if n == 1 {
			tmp := make([]interface{}, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
