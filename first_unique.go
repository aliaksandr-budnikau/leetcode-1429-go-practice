package firstunique

// FirstUnique tracks the first integer that occurs exactly once.
//
// TODO: Implement the methods below.
type FirstUnique struct {
	counterMap map[int]int
	queue      []int
}

func Constructor(nums []int) FirstUnique {
	counterMap := make(map[int]int, len(nums))
	queue := make([]int, 0, len(nums))

	for _, n := range nums {
		counterMap[n]++
	}

	for _, n := range nums {
		if counterMap[n] == 1 {
			queue = append(queue, n)
		}
	}

	return FirstUnique{
		counterMap: counterMap,
		queue:      queue,
	}
}

func (fu *FirstUnique) ShowFirstUnique() int {
	for _, n := range fu.queue {
		if fu.counterMap[n] == 1 {
			return n
		}
		fu.queue = fu.queue[1:]
	}
	return -1
}

// Add inserts value into the data structure.
func (fu *FirstUnique) Add(n int) {
	if fu.counterMap[n] == 0 {
		fu.queue = append(fu.queue, n)
	}
	fu.counterMap[n]++
}
