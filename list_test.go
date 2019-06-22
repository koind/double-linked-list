package double_linked_list

import (
	"testing"
)

var testCases = map[int]struct{}{
	2: struct{}{},
	4: struct{}{},
	5: struct{}{},
}

func TestСheck(t *testing.T) {
	list := NewList()

	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)

	list.PushBack(4)
	list.PushBack(5)
	list.PushBack(6)

	list.Remove(1)
	list.Remove(3)
	list.Remove(6)

	item := list.First()
	for {
		val := item.Value().(int)
		if _, has := testCases[val]; !has {
			t.Errorf("value should be in the list: %v - %v", val, testCases)
		}

		item = item.Next()
		if item == nil {
			break
		}
	}
}
