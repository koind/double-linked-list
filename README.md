# double-linked-list

Implement the doubly linked list on Go.


## Installation

Run the following command from you terminal:


 ```bash
 go get github.com/koind/double-linked-list
 ```

## Usage

Package usage example.

```go
package main

import (
	"fmt"
	"github.com/koind/double-linked-list"
)

func main() {
	list := double_linked_list.NewList()
	
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	
	item := list.Last()
	fmt.Println(item.Value()) // 3
}
```

## Available Methods

The following methods are available:

##### koind/double-linked-list

```go
NewList() *List
Len() int
First() *Item
Last() *Item
PushFront(v interface{})
PushBack(v interface{})
Remove(v interface{}) bool
Value() interface{}
Next() *Item
Prev() *Item
```

## Tests

Run the following command from you terminal:

```
go test -v .
```