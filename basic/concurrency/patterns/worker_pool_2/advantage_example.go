package main

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	numberOfWorkers int
	jobs            []*Job
	jobQueue        chan *Job
	wg              sync.WaitGroup
}

// Create a worker pool
func NewPool(jobs []*Job, numberOfWorkers int) *Pool {
	return &Pool{
		jobs:            jobs,
		numberOfWorkers: numberOfWorkers,
		jobQueue:        make(chan *Job),
	}
}

// Define a func to:
// 1. Create workers
// 2. Send jobs to job queue
func (p *Pool) Run() {

	// 1. Create workers
	for i := 0; i < p.numberOfWorkers; i++ {
		go p.work()
	}

	// 2. Listen and receive(read) jobs from a queue
	// 2.1 add wait group for each job
	p.wg.Add(len(p.jobs))
	// 2.2 transer jobs from slice to queue(channel)
	for _, job := range p.jobs {
		p.jobQueue <- job
	}

	// 2.3 close the job queue
	close(p.jobQueue)

	// 2.4 wait for all workers to finish
	p.wg.Wait()
}

func (p *Pool) work() {

	// listen and receive(read) jobs from a queue
	for job := range p.jobQueue {

		// do some work
		job.Run(&p.wg)
	}
}

type Job struct {
	Id  int
	Err error

	f func() error
}

func NewJob(id int, f func() error) *Job {
	return &Job{Id: id, f: f}
}

func (j *Job) Run(wg *sync.WaitGroup) {
	fmt.Printf("Job %d is running\n", j.Id)

	time.Sleep(time.Second)
	j.Err = j.f()

	fmt.Printf("Job %d done\n", j.Id)
	wg.Done()

}

func main() {

	// create a list of jobs
	jobs := []*Job{
		NewJob(1, func() error { return nil }),
		NewJob(2, func() error { return nil }),
		NewJob(3, func() error { return nil }),
		NewJob(4, func() error { return nil }),
		NewJob(5, func() error { return nil }),
		NewJob(6, func() error { return nil }),
		NewJob(7, func() error { return nil }),
		NewJob(8, func() error { return nil }),
		NewJob(9, func() error { return nil }),
		NewJob(10, func() error { return nil }),
		NewJob(11, func() error { return nil }),
		NewJob(12, func() error { return nil }),
		NewJob(13, func() error { return nil }),
		NewJob(14, func() error { return nil }),
		NewJob(15, func() error { return nil }),
		NewJob(16, func() error { return nil }),
		NewJob(17, func() error { return nil }),
		NewJob(18, func() error { return nil }),
	}

	// create a worker pool and start the workers
	p := NewPool(jobs, 5)
	p.Run()
}
