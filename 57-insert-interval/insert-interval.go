func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	index := len(intervals)
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > newInterval[0] {
			index = i
			break
		}
	}
	prev := newInterval
	intervals = append(intervals, newInterval)
	for i := index; i < len(intervals); i++ {
		temp := intervals[i]
		intervals[i] = prev
		prev = temp

	}
	prev = intervals[0]
	res := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if i == len(intervals)-1 {
			if intervals[i][0] <= prev[1] {
				prev[1] = max(prev[1], intervals[i][1])
				prev[0] = min(prev[0], intervals[i][0])
				res = append(res, prev)
			} else {
				res = append(res, prev)
				res = append(res, intervals[i])
			}
			break
		}
		if intervals[i][0] <= prev[1] {
			prev[1] = max(prev[1], intervals[i][1])
			prev[0] = min(prev[0], intervals[i][0])
			// [1,6]
		} else {
			res = append(res, prev)
			prev = intervals[i]
		}
	}
	return res
}