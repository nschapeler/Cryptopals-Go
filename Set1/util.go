package Set1

import "container/heap"

// Heap implementation to get the bottom x values out of our map (Used for finding keysize in Kasiski Method)
func getBottomXValues(m map[uint32]float32, x uint32) []kv {
	h := makeHeap(m)
	res := make([]kv, x)

	for i := range res {
		res[i] = heap.Pop(h).(kv)
	}

	return res

}

// Based on https://stackoverflow.com/questions/52514821/how-to-extract-x-top-int-values-from-a-map-in-golang
func makeHeap(m map[uint32]float32) *KeyDistanceHeap {
	h := &KeyDistanceHeap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, kv{k, v})
	}
	return h
}

// See https://golang.org/pkg/container/heap/

type kv struct {
	key   uint32
	value float32
}
type KeyDistanceHeap []kv

func (h KeyDistanceHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h KeyDistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h KeyDistanceHeap) Len() int           { return len(h) }

func (h *KeyDistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(kv))
}

func (h *KeyDistanceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
