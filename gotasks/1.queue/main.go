package main

import "fmt"

type Queue struct{
	front, rear, n int
}

var que[100] int

func (q *Queue) insert(){
	var val int 
	if q.rear == q.n - 1{
		fmt.Println("Queue overflow")
	}else{
		if q.front == -1{
			q.front = 0
		}
		fmt.Print("Element to be inserted: ")
		fmt.Scanln(&val)
		q.rear++
		que[q.rear] = val
	}
}

func (q *Queue) delete(){
	if q.front == - 1 && q.front > q.rear{
		fmt.Println("Queue underflow")
	}else{
		fmt.Println("Element deleted from queue is " ,que[q.front])
		q.front++
	}
}

func (q *Queue) display(){
	if q.front == -1{
		fmt.Println("Queue is empty")
	}else{
		fmt.Println("Queue elements are: ")
		for i:=q.front;i<=q.rear;i++{
			fmt.Printf("%d ",que[i])
		}
		fmt.Print("\n")
	}
}

func main(){
	q := &Queue{-1, -1, 100}
	var choice int
	fmt.Println("1. Insert element to queue")
	fmt.Println("2. Delete element from queue")
	fmt.Println("3. Display queue")
	for i:=1;i!=0;{
		fmt.Scanln(&choice)
		switch choice{
		case 1: q.insert()
		case 2: q.delete()
		case 3: q.display()
		default:
			i=0
			fmt.Println("Invalid choice")
		}
	}

}