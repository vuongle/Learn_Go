package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type ExecutionFn func(ctx context.Context, args interface{}) (interface{}, error)

type JobDescriptor struct {
	JobID    string
	JobType  string
	Metadata map[string]interface{}
}

type Result struct {
	Value      interface{}
	Err        error
	Descriptor JobDescriptor
}

type Job struct {
	Descriptor JobDescriptor
	ExecFn     ExecutionFn
	Args       interface{}
}

func (j *Job) execute(ctx context.Context) Result {
	fmt.Printf("Job %s is running\n", j.Descriptor.JobID)
	value, err := j.ExecFn(ctx, j.Args)
	if err != nil {
		return Result{
			Err:        err,
			Descriptor: j.Descriptor,
		}
	}

	return Result{
		Value:      value,
		Descriptor: j.Descriptor,
	}
}

// define a worker function
func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	// use infinite loop to listen on jobs
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			// execute the job and send result to "results" channel
			results <- job.execute(ctx)
		case <-ctx.Done():
			fmt.Printf("cancelled worker. Error detail: %v\n", ctx.Err())
			results <- Result{
				Err: ctx.Err(),
			}
			return
		}
	}
}

// define a worker pool
type WorkerPool struct {
	workersCount int
	jobs         chan Job
	results      chan Result
	Done         chan struct{}
}

func NewWorkerPool(workersCount int) WorkerPool {
	return WorkerPool{
		workersCount: workersCount,
		jobs:         make(chan Job, workersCount),
		results:      make(chan Result, workersCount),
		Done:         make(chan struct{}),
	}
}

func (p *WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < p.workersCount; i++ {
		// add each worker to wait group
		wg.Add(1)

		// fan out worker goroutines
		//reading from jobs channel and
		//pushing calcs into results channel
		go worker(ctx, &wg, p.jobs, p.results)
	}

	// wait for all workers to finish
	wg.Wait()

	// once all workers are done, close both channels
	close(p.Done)
	close(p.results)
}

func (p *WorkerPool) Results() <-chan Result {
	return p.results
}

func (p *WorkerPool) GenerateFrom(jobsBulk []Job) {
	for i := range jobsBulk {
		j := jobsBulk[i]
		fmt.Printf("Stream %s to the jobs channel\n", j.Descriptor.JobID)
		p.jobs <- j
	}

	close(p.jobs)
}

const (
	jobsCount   = 10
	workerCount = 2
)

var (
	// A custom execution function of a job. the logic is just to multiply the argument by 2.
	execFn = func(ctx context.Context, args interface{}) (interface{}, error) {
		argVal, ok := args.(int)
		if !ok {
			return nil, errors.New("wrong argument type")
		}

		return argVal * 2, nil
	}
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	// create a worker pool
	wp := NewWorkerPool(workerCount)

	// transfer jobs to the jobs channel in the worker pool by using generator pattern
	go wp.GenerateFrom(testJobs())

	// start the worker pool
	go wp.Run(ctx)

	// get results from the results channel
	for {
		select {
		case r, ok := <-wp.Results():

			// continue reading the results channel even there is any error
			if !ok {
				continue
			}

			fmt.Printf("Result: %v\n", r.Value)
		case <-wp.Done:
			return
		default:
		}
	}

}

func testJobs() []Job {
	jobs := make([]Job, jobsCount)
	for i := 0; i < jobsCount; i++ {
		jobs[i] = Job{
			Descriptor: JobDescriptor{
				JobID:    fmt.Sprintf("Job#%v", i),
				JobType:  "anyType",
				Metadata: nil,
			},
			ExecFn: execFn,
			Args:   i,
		}
	}
	return jobs
}
