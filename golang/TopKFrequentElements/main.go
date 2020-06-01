package main

type item struct {
	value int
	key   int
}

type MinHeap struct {
	size    int
	maxsize int
	heap    []item
}

func (m *MinHeap) add(val item) {
	if m.size < m.maxsize {
		m.heap = append(m.heap, val)
		m.size++
		m.heapfyUp(m.size - 1)
	} else {
		if val.value > m.heap[0].value {
			m.heap[0] = val
			m.heapfyDown(0)
		}
	}
}

func (m *MinHeap) pop() int {
	ret := m.heap[0].key
	m.heap[0] = m.heap[m.size-1]
	m.size--
	m.maxsize--
	m.heapfyDown(0)
	return ret
}

func (m *MinHeap) swap(a, b int) {
	m.heap[a], m.heap[b] = m.heap[b], m.heap[a]
}

func (m *MinHeap) heapfyUp(index int) {
	for index > 0 {
		p := (index - 1) / 2
		if m.heap[p].value > m.heap[index].value {
			m.swap(p, index)
			index = p
		} else {
			break
		}
	}
}

func (m *MinHeap) heapfyDown(index int) {
	for 2*index+1 < m.size {
		s := index
		if m.heap[2*index+1].value < m.heap[index].value {
			s = 2*index + 1
		}
		if 2*index+2 < m.size && (m.heap[2*index+2].value < m.heap[s].value) {
			s = 2*index + 2
		}
		if s != index {
			m.swap(s, index)
			index = s
		} else {
			break
		}
	}
}

func main() {

}

func topKFrequent(nums []int, k int) []int {
	obj := &MinHeap{0, k, []item{}}
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++

	}
	for key, val := range hash {
		i := item{val, key}
		obj.add(i)
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		feek := obj.pop()
		result[k-i-1] = feek
	}
	return result
}
