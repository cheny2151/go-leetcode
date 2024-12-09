package medium

import "testing"

// https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

func TestName(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	findKthLargest(nums, 20)
}

func findKthLargest(nums []int, k int) int {
	heap := make([]*int, k)
	heap[0] = &nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if i < k {
			// 初始化堆
			heap[i] = &num
			pidx := (i - 1) >> 1
			for {
				if pidx >= 0 && fixHeap(pidx, heap) {
					pidx = (pidx - 1) >> 1
				} else {
					break
				}
			}
		} else if num > *heap[0] {
			heap[0] = &num
			fixHeap(0, heap)
		}
	}

	return *heap[0]
}

func fixHeap(idx int, heap []*int) bool {
	minIdx := idx
	lidx := idx*2 + 1
	if lidx < len(heap) && (heap[lidx] != nil && *heap[lidx] < *heap[minIdx]) {
		minIdx = lidx
	}
	ridx := idx*2 + 2
	if ridx < len(heap) && (heap[ridx] != nil && *heap[ridx] < *heap[minIdx]) {
		minIdx = ridx
	}
	if minIdx != idx {
		swap(minIdx, idx, heap)
		fixHeap(minIdx, heap)
		return true
	}
	return false
}

func swap(a, b int, heap []*int) {
	bv := heap[b]
	heap[b] = heap[a]
	heap[a] = bv
}

// 官方题解
func findKthLargest2(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize/2 - 1; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
