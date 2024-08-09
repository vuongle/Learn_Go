Nói qua về Worker Pool Worker Pool qua tên gọi ta có thể hình dung được công dụng của nó là tạo ra một nơi chứa gọi là pool, để chứa các worker của chúng ta. Mục đích là để ta có thể quản lý các worker, quản lý việc phân phối task và đặc biệt là kiểm soát được những tài nguyên dùng chung giữa các worker. Ví dụ như các worker chạy đồng thời và cùng truy xuất vào 1 file hoặc dùng chung một API.

Workerpool in Golang which is also known as thread pool is the pattern used to achieve concurrency in golang. The idea behind a worker pool is to have a fixed number of goroutines running in the background, waiting for work to be assigned to them. When work is assigned to a worker, it is executed in the background, allowing the main goroutine to continue executing other code.

Problem it solves:

Assume you don't have a limitless resource on your machine; the minimum size of a goroutine object is 2 KB; creating too many goroutines will soon exhaust your machine's memory, and the CPU will continue executing the job until it hits the limit. We can minimize the burst of CPU and memory by utilizing a restricted pool of workers and keeping the job in the queue since the task will wait in the queue until the worker pulls it. In this scenario, a workerpool is very beneficial to use as it only limits task execution, not the number of tasks queued. Also, it never blocks the submitting tasks.

Working of the Worker Pool in Golang?
The worker pool is responsible for managing the lifecycle of the workerpool goroutines and distributing the tasks among them.
The basic structure of a workerpool in Golang consists of three main components:
task queue: is a channel that is used to store incoming tasks
worker pool: is a group of worker goroutines that are responsible for handling the tasks from the task queue
manager goroutine: is responsible for managing the worker pool, creating new worker goroutines as needed, and distributing the tasks among them.

Steps to Implement Worker Pool in Golang with Code
-Define the worker function:
The worker function should have a channel for receiving jobs and a channel for sending results. The worker function should listen for jobs on the jobs channel, perform the desired operations on each job, and send the output to the results channel.
-Create the channels:
In the main function, create the channels for jobs and results. The jobs channel should be used to send jobs to the worker pool, and the results channel should be used to receive output from the worker pool.
-Start the goroutines:
Use the go keyword to start a specified number of goroutines, each running the worker function with the jobs and results channels as arguments.
-Send jobs to the jobs channel:
Use a loop to send a specified number of jobs to the jobs channel.
-Close the jobs channel:
Once all the jobs have been sent, close the jobs channel to signal to the worker goroutines that there are no more jobs to be processed.
-Receive outputs from the results channel:
Use a loop to receive the outputs from the results channel.
-Wait for the goroutines to finish:
Use the sync package's WaitGroup to wait for the goroutines to finish(this is optional we can work without waitgroup also).
-Close the results channel:
Once all the outputs are received, close the results channel.

links
https://itnext.io/explain-to-me-go-concurrency-worker-pool-pattern-like-im-five-e5f1be71e2b0
https://github.com/godoylucase/workers-pool/blob/develop/wpool/job.go
