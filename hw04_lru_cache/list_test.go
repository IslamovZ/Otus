package hw04lrucache

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())
		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		fmt.Println("30 is ==== ", l.Show()[4], l.Show()[4].Next)

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		fmt.Println("30 is ==== ", l.Show()[4], l.Show()[4].Next)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]

		fmt.Println("30 is ==== ", l.Show()[4], l.Show()[4].Next)

		l.MoveToFront(l.Back()) // [70, 80, 60, 40, 10, 30, 50]

		fmt.Println("30 is ==== ", l.Show()[5], l.Show()[5].Next)

		index := 0
		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			fmt.Println("element is ", *i)
			fmt.Println("element of ", l.Show()[index])

			elems = append(elems, i.Value.(int))
			index++
		}

		fmt.Println(l.Show())
		fmt.Println(elems)

		// require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
