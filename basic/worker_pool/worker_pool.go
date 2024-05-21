package main

import (
	"fmt"
	"time"
)

func main() {
	const numberOfJobs = 5

	// Tao 2 buffered channel mo ta cho so luong cong viec va ket qua cua cac cong viec do
	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)

	// Tao 3 worker
	// Tai thoi diem nay, 3 worker deu bi block vi chua co job nao duoc send vo channel "jobs"
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Send 5 jobs vo channel
	for j := 1; j <= numberOfJobs; j++ {
		fmt.Printf("------send job #%d to jobs channel\n", j)
		jobs <- j
	}
	close(jobs)

	// Lay ket qua ma 3 worker da xong
	for a := 1; a <= numberOfJobs; a++ {
		<-results
	}
}

// Nhiem vu cua worker:
// 1. Nhan 1 job trong channel "jobs"
// 2. Thuc hien 1 nhiem vu
// 3. Send ket qua cua nhiem vu vo channel "results"
// channel "jobs": chi doc tu channel nen se la "<-chan"
// channel "results": chi write vo channel nen se la "chan<-" (huong cua dau "<-" chia vo "chan")
func worker(id int, jobs <-chan int, results chan<- int) {

	// Duyet qua tung job ( gia tri cua job chinh la so int duoc send trong main)
	for j := range jobs {

		// 1. Nhan 1 job trong channel "jobs"
		fmt.Printf("worker: #%d started job: #%v\n", id, j)
		time.Sleep(time.Second)

		// 2. Thuc hien 1 nhiem vu (j * 2)
		// 3. Send ket qua cua nhiem vu vo channel "results"
		fmt.Println("worker: #", id, "finished job: #", j)
		results <- (j * 2)
	}
}
