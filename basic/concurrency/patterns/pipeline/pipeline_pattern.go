package main

import "fmt"

// bai toan: tinh tong binh phuong cua 1 day so
// cach giai quyet: su dung cac pattern duoc goi la "pipeline", fan-in, fan-out
// trong lap trinh concurrency cua golang de tan dung da luong cua CPU
func main() {

	//----------------------------------------------------------------------------------------------------
	// Cach tinh thong thuong:
	// Voi cach tinh nay thi chuong trinh chay trong 1 process (khong tan dung duoc tinh da luong cua CPU)
	//----------------------------------------------------------------------------------------------------
	// numbers := []int{}
	// for i := 1; i <= 10000; i++ {
	// 	numbers = append(numbers, i)
	// }

	// sum := 0
	// for i := 0; i < 10000; i++ {
	// 	sum += numbers[i] * numbers[i]
	// }

	// fmt.Printf("Sum: %v", sum)

	//----------------------------------------------------------------------------------------------------
	// Cach su dung goroutine + channel:
	// Tan dung tinh da luong cua CPU
	//----------------------------------------------------------------------------------------------------
	fmt.Println("Fan In - Fan Out")
	numbers := []int{}
	for i := 1; i <= 10000; i++ {
		numbers = append(numbers, i)
	}

	// generate pipeline
	inputChan := generatePipeline(numbers)

	// su dung pattern duoc goi la 'fan-out': nghia la lay cac gia tri trong "inputChan"
	// chia ra cho cac channel con. trong vi du nay la chia cho 4 channel con
	c1 := fanOut(inputChan)
	c2 := fanOut(inputChan)
	c3 := fanOut(inputChan)
	c4 := fanOut(inputChan)

	// su dung pattern duoc goi la 'fan-in': nghia la gom cac channel con thanh 1 channel
	c := fanIn(c1, c2, c3, c4)

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += <-c
	}

	fmt.Printf("Sum: %v", sum)
}

// tao 1 func
// tham so: la 1 day so can tinh
// tra ve: la 1 channel. Vi func nay send tung value cua day so vao channel
// nen gia tri tra ve la "<-chan"
func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range numbers {
			//send value of "n" to the channel "out"
			out <- n
		}
		close(out)
	}()

	return out
}

func fanOut(in <-chan int) <-chan int {
	outSub := make(chan int)

	go func() {
		// lay tung gia tri ma generatePipeline send vao channel "in"
		// va send no vo channel "outSub"
		for n := range in {
			outSub <- (n * n)
		}
		close(outSub)
	}()

	return outSub
}

func fanIn(inputChannels ...<-chan int) <-chan int {
	in := make(chan int)

	go func() {

		// duyet qua danh sach cac channel
		for _, c := range inputChannels {
			// duyet qua cac gia tri trong tung channel dang loop
			for n := range c {
				in <- n
			}
		}
	}()

	return in
}
