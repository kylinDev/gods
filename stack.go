package main

import (
	"errors"
	"fmt"
	"sync"
)

type CustomStack struct {
	Container []string
	Size int
	lock sync.RWMutex
}

func NewCustomStack() *CustomStack {
	return &CustomStack{}
}

func (c *CustomStack) Seek() string {
	if c.Size == 0 {
		return ""
	}
	val := c.Container[c.Size-1]
	return val
}

func (c *CustomStack) Pop() (string,error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.Size == 0 {
		return "",errors.New("stack is empty")
	}
	val := c.Container[c.Size-1]
	c.Container = c.Container[:c.Size-1]
	c.Size --
	return val,nil
}

func (c *CustomStack) Push(e string)  {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.Container = append(c.Container,e)
	c.Size++
}
func main() {
	//test()

	stack := NewCustomStack()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	for _,s := range stack.Container {
		fmt.Printf("s:%s,size=%d\n",s,stack.Size)
	}
	val,err := stack.Pop()
	if err != nil {
		fmt.Printf("pop err:%+v\n",err)
	}
	fmt.Printf("%s is pop,current size is %d \n",val,stack.Size)
}


