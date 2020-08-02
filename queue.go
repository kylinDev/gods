package main

import (
	"errors"
	"fmt"
	"sync"
)

type CustomQueue struct {
	Container []string
	Size int
	lock sync.RWMutex
}

func (c *CustomQueue) Pop()(string,error)  {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.Size == 0 {
		return "",errors.New("queue is empty")
	}
	e := c.Container[c.Size-1]
	c.Size--
	slice := make([]string,0)
	for i:=0;i<c.Size;i++ {
		slice = append(slice,c.Container[i])
	}
	c.Container = slice
	return e,nil
}

func (c *CustomQueue) PushFront(e string)  {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.Size == 0 {
		c.Container = append(c.Container,e)
		return
	}
	slice := make([]string,0)
	slice = append(slice,e)
	for _,ele := range c.Container {
		slice = append(slice,ele)
	}
	c.Container = slice
	c.Size++
}


func (c *CustomQueue) PushBack(e string)  {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.Container = append(c.Container,e)
	c.Size++
}

func (c *CustomQueue) QueueSize()  int{
	return c.Size
}

func main()  {
	q := new(CustomQueue)
	q.PushBack("b")
	q.PushBack("c")
	for _,e := range q.Container {
		fmt.Printf("e:%s\n",e)
	}
	fmt.Printf("push front\n")
	q.PushFront("a")
	for _,e := range q.Container {
		fmt.Printf("e:%s\n",e)
	}

	fmt.Printf("pop \n")
	e,_ := q.Pop()
	fmt.Printf("%s is pop\n",e)
	for _,e := range q.Container {
		fmt.Printf("e:%s\n",e)
	}
}