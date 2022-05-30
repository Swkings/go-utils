/**
 * @Author: Swking
 * @File:  priority_queue_test
 * @Date: 2022/5/29 20:00
 * @Description: 泛型版 优先队列  未做安全操作
 */

package priority_queue

import (
	"errors"
)

func CmpMinTop(a, b int) bool { return a < b }
func CmpMaxTop(a, b int) bool { return a > b }

type PriorityQueue[T any] struct {
	vector []T
	Cmp    func(a, b T) bool
}

// justify 调整heap
func (pq *PriorityQueue[T]) justify() {
	n := pq.Len()
	for i := n/2 - 1; i >= 0; i-- {
		pq.down(i, n)
	}
}

func (pq *PriorityQueue[T]) New() {
	pq.vector = make([]T, 0)
}

func (pq *PriorityQueue[T]) NewByArray(arr []T) {
	pq.vector = append(pq.vector, arr...)
	pq.justify()
}

func (pq *PriorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && pq.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !pq.Less(j, i) {
			break
		}
		pq.Swap(i, j)
		i = j
	}
	return i > i0
}

func (pq *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.Less(j, i) {
			break
		}
		pq.Swap(i, j)
		j = i
	}
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.vector)
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.vector[i], pq.vector[j] = pq.vector[j], pq.vector[i]
}

// Less 绑定Less方法，这里用的是小于号，生成的是小根堆
func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return pq.Cmp(pq.vector[i], pq.vector[j])
}

// Pop
func (pq *PriorityQueue[T]) Pop() (*T, error) {
	n := pq.Len() - 1
	if n < 0 {
		return nil, errors.New("PriorityQueue is empty")
	}
	pq.Swap(0, n)
	pq.down(0, n)
	item := pq.vector[n]
	pq.vector = pq.vector[0:n]
	return &item, nil
}

// Push 绑定push方法
func (pq *PriorityQueue[T]) Push(item T) {
	pq.vector = append(pq.vector, item)
	pq.up(pq.Len() - 1)
}

func (pq *PriorityQueue[T]) Front() (*T, error) {
	n := pq.Len() - 1
	if n < 0 {
		return nil, errors.New("PriorityQueue is empty")
	}
	return &(pq.vector[0]), nil
}

// ToArray 生成数组, 深拷贝
func (pq *PriorityQueue[T]) ToArray() []T {
	pqCopy := PriorityQueue[T]{
		Cmp: pq.Cmp,
	}
	pqCopy.NewByArray(pq.vector)
	n := pqCopy.Len()
	if n == 0 {
		return []T{}
	}
	sortArr := make([]T, n)
	var item *T
	for i := 0; i < n; i++ {
		item, _ = pqCopy.Pop()
		sortArr[i] = *item
	}
	return sortArr
}
