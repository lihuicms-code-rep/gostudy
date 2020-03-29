package main

import (
	"fmt"
	"time"
)

//Task类描述
//具体要被处理的任务
type Task struct {
	f func() error   //具体的业务就是f
}

func NewTask(f func() error) *Task {
	return &Task{
		f,
	}
}

//任务执行
func (t *Task) ExecTask() {
	t.f()
}

//协程池的类描述
type GoPool struct {
	//任务进入的入口
	entryChan chan *Task

	//内部任务队列
	jobsChan chan *Task

	//goroutine数量
	maxNum int
}

func NewGoPool(num int) *GoPool {
	return &GoPool{
		entryChan:make(chan *Task),
		jobsChan:make(chan *Task),
		maxNum:num,
	}
}

//创建一个worker并工作
func (p *GoPool) worker(id int) {
	//1.永久从jobsChan中取任务
	for task := range p.jobsChan {
		//2.如果取出任务,则做事
		task.ExecTask()
		fmt.Printf("worker id:%d exec task finished\n", id)
	}
}

//协程池工作
func (p *GoPool) run() {
	//1.根据数量,创建goroutine worker
	for i := 0; i < p.maxNum; i++ {
		go p.worker(i)
	}

	//2.从entry中进入的任务进入jobs中
	for task := range p.entryChan {
		p.jobsChan <- task
	}
}

func business() error {
	fmt.Println(time.Now())
	return nil
}


func main() {
	//1.创建一些具体任务
	t := NewTask(business)

	//2.创建协程池
	p := NewGoPool(4)

	//3.将任务交给协程池做
	go func() {
		for {
			p.entryChan <- t
		}
	}()

	//4.启动pool
	p.run()

}


