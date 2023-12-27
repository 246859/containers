# containers

![Static Badge](https://img.shields.io/badge/go-%3E%3D1.21-blue)
![GitHub License](https://img.shields.io/github/license/246859/containers)

base data structure and algorithm implement by go genericity



## install
```bash
$ go get -u github.com/246859/containers@latest
```



## usage

Import the package and use the structure. Here is a simple example about ArrayList.

```go
import (
    "fmt"
    "github.com/246859/containers/lists"
)

func main() {
    arrayList := lists.NewArrayList[int](10)
    arrayList.Add(1)
    arrayList.Add(2)
    arrayList.Add(3, 4, 5)

    ele, has := arrayList.Get(0)
    if !has {
       panic("element not found")
    }
    fmt.Println(ele)
}
```

Here is a simple example about PriorityQueue.

```go
package main

import (
    "cmp"
    "fmt"
    "github.com/246859/containers/queues"
)

func main() {
    priorityQueue := queues.NewPriorityQueue[int](20, cmp.Compare[int])
    priorityQueue.Push(2)
    priorityQueue.Push(1)
    priorityQueue.Push([]int{0, 6, 3, 2, 5}...)

    top, has := priorityQueue.Peek()
    fmt.Println(top, has)
}
```

know more information to see  [ToDo List](TODO.md).



## contribute

1. fork this repository
2. checkout your feature branch
3. commit changes, know more about [commit-messages-guide](https://github.com/RomuloOliveira/commit-messages-guide)
4. create a pull request to this repository

