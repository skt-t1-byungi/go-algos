package heap

type Heap []int

func NewHeap() *Heap {
	return new(Heap)
}

func (h *Heap) Push(n int) {
	*h = append(*h, n)
	arr := *h
	idx := len(arr) - 1

	if idx == 0 {
		return
	}

	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if arr[parentIdx] <= arr[idx] {
			break
		}
		arr[parentIdx], arr[idx] = arr[idx], arr[parentIdx]
		idx = parentIdx
	}
}

func (h *Heap) Pop() int {
	arr := *h
	root := arr[0]

	size := len(arr) - 1
	arr[0] = arr[size]
	idx := 0

	for idx < size {
		right := (idx + 1) * 2
		if right < size && arr[right] < arr[idx] {
			arr[right], arr[idx] = arr[idx], arr[right]
			idx = right
			continue
		}

		left := right - 1
		if left < size && arr[left] < arr[idx] {
			arr[left], arr[idx] = arr[idx], arr[left]
			idx = left
			continue
		}

		break
	}

	*h = arr[:size]

	return root
}
