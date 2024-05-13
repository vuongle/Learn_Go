package main

import (
	"fmt"
	"time"
)

// There are 3 objects in job queue
// 1.Job
// 2.Queue
// 3.Worker

type Job interface {
	Process()
}

type Worker struct {
	WorkerId   int
	Done       chan bool // used to infirm that Worker finishes a job
	JobRunning chan Job  // this channel receives a job to do
}

func NewWorker(workerId int, jobChannel chan Job) *Worker {
	return &Worker{
		WorkerId:   workerId,
		Done:       make(chan bool),
		JobRunning: jobChannel,
	}
}

func (w *Worker) Run() {
	fmt.Println("Run worker id: ", w.WorkerId)
	// worker runs concurrency
	go func() {
		// use a infinity loop to listen when Job is done by worker
		// when the Job is done -> exit the loop
		for {
			select {
			case job := <-w.JobRunning:
				fmt.Println("Job running by worker id: ", w.WorkerId)
				job.Process()
			case <-w.Done:
				fmt.Println("Job is done by workder: ", w.WorkerId)
				return // exit the loop
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.Done <- true // send a value to a channel
}

type JobQueue struct {
	Workers    []*Worker // a list of workers that can take jobs and do the jobs
	JobRunning chan Job  // a job is sent to the queue
	Done       chan bool // used to infirm when there is no job in the queue
}

func NewJobQueue(numberOfWorkers int) JobQueue {
	workers := make([]*Worker, numberOfWorkers, numberOfWorkers) // 2nd param: length of slice,// 3rd param: capacity of slice
	jobRunning := make(chan Job)

	// because the "workers" slice is currently empty after using make()
	// therefore, init each element inside it
	for i := 0; i < numberOfWorkers; i++ {
		workers[i] = NewWorker(i, jobRunning)
	}

	return JobQueue{
		Workers:    workers,
		JobRunning: jobRunning,
		Done:       make(chan bool),
	}

}

func (q *JobQueue) Push(job Job) {
	q.JobRunning <- job // send a value to a channel (means: add a job to queue)
}

func (q *JobQueue) Start() {

	// use 2 goroutines to run concurrency

	go func() {
		// use a loop to run all workers that registered with JobQueue
		for i := 0; i < len(q.Workers); i++ {
			q.Workers[i].Run()
		}
	}()

	go func() {
		// use another loop to listen when the JobQueue has no no job
		// when all done -> stop workers
		for {
			select {
			case <-q.Done:
				for i := 0; i < len(q.Workers); i++ {
					q.Workers[i].Stop()
				}
				return
			}
		}

	}()
}

func (q *JobQueue) Stop() {
	q.Done <- true
}

// Implement Process() of the Job interface
// After implementing, the send is also a job
type Sender struct {
	Email string
}

func (s Sender) Process() {
	fmt.Println("Sent an email to: ", s.Email)
}

func main() {
	emails := []string{
		"user1@gmail.com",
		"user2@gmail.com",
		"user3@gmail.com",
		"user4@gmail.com",
		"user5@gmail.com",
	}

	// create queue
	jobQueue := NewJobQueue(4)
	jobQueue.Start()

	// init jobs (sender is a job because it implements Process() method)
	for _, email := range emails {
		sender := Sender{Email: email}
		jobQueue.Push(sender)
	}

	time.AfterFunc(time.Second*3, func() {
		jobQueue.Stop()
	})

	time.Sleep(time.Second * 30)
}
