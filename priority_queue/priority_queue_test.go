/*
 * @Author: Swking
 * @Date: 2022-05-29 20:20:43
 * @LastEditTime: 2022-05-30 14:42:49
 * @LastEditors: Swking
 * @FilePath: \util\queue\priority_queue_test.go
 * @Description:
 */

package priority_queue

import (
	"fmt"
	"testing"
)

// greater 小顶堆
func less(a, b Item) bool { return a.value < b.value }

// greater 大顶堆
func greater(a, b Item) bool { return a.value > b.value }

type Item struct {
	value int
	name  string
}

func TestPriorityQueue(t *testing.T) {

	var err error

	// pq1 := PriorityQueue[int]{
	// 	Cmp: CmpMinTop,
	// }
	// pq1.New()
	// pq1.Push(1)
	// fmt.Println("Arr", pq1.ToArray())
	// fmt.Println()
	// var item *int
	// item, err = pq1.Pop()
	// fmt.Println("Pop", item, err)
	// fmt.Println("Arr", pq1.ToArray())
	// fmt.Println("==================")

	pq := PriorityQueue[Item]{
		Cmp: greater,
	}
	var arr []Item = []Item{
		{value: 1, name: "a"},
		{value: -2, name: "b"},
		{value: 3, name: "c"},
		{value: -4, name: "d"},
		{value: 6, name: "e"},
		{value: 5, name: "f"},
	}
	pq.NewByArray(arr)
	fmt.Println("Arr", pq.ToArray())
	fmt.Println()

	var PopItem *Item
	var PushItem Item

	PopItem, err = pq.Pop()
	fmt.Println("Pop", PopItem, err)
	fmt.Println("Arr", pq.ToArray())
	fmt.Println()

	PopItem, err = pq.Pop()
	fmt.Println("Pop", PopItem, err)
	fmt.Println("Arr", pq.ToArray())
	fmt.Println()

	PushItem = Item{value: 4, name: "g"}
	fmt.Println("Push", PushItem)
	pq.Push(PushItem)
	fmt.Println("Arr", pq.ToArray())
	fmt.Println()

	PushItem = Item{value: 8, name: "f"}
	fmt.Println("Push", PushItem)
	pq.Push(PushItem)
	fmt.Println("Arr", pq.ToArray())
	fmt.Println()

	n := pq.Len()
	for i := 0; i < n; i++ {
		PopItem, err = pq.Pop()
		fmt.Println("Pop", PopItem, err)
		fmt.Println("Arr", pq.ToArray())
		fmt.Println()
	}

	PopItem, err = pq.Pop()
	fmt.Println("Pop", PopItem, err)
	fmt.Println("Arr", pq.ToArray())
	fmt.Println()

	PopItem, err = pq.Pop()
	fmt.Println("Pop", PopItem, err)
	fmt.Println("Arr", pq.ToArray())
	fmt.Println()

}
