package main

import (
    "container/heap"
    "fmt"
)

// stores [freq, num] pairs
type MinHeap [][2]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] } // min by freq
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func topKFrequentMinHeap(nums []int, k int) []int {
    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }

    h := &MinHeap{}
    heap.Init(h)

    for num, freq := range count {
        heap.Push(h, [2]int{freq, num})
        if h.Len() > k {
            heap.Pop(h)  // remove lowest freq
        }
    }

    res := make([]int, k)
    for i := k - 1; i >= 0; i-- {
        res[i] = heap.Pop(h).([2]int)[1]
    }
    return res
}

func main() {
    nums := []int{1, 2, 2, 3, 3, 3}
    k := 2
    fmt.Println(topKFrequentMinHeap(nums, k))
}
