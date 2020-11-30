package main

import (
	"fmt"
	"sync"
	"time"
)

func wrong() {

	function := func(i int) {

		time.Sleep(3)
		fmt.Println(i)

	}
	for i := 0; i < 10; i++ {
		go function(i) // Like thread
	}

	/*
		time ./syncWait
		0.00user 0.00system 0:00.00elapsed 133%CPU (0avgtext+0avgdata 1904maxresident)k
		0inputs+0outputs (0major+199minor)pagefaults 0swaps

		Mean go doesn't wait goruntines's finish :D (*)
	*/

}

func toward() {
	var gp sync.WaitGroup
	function := func(i int) {
		time.Sleep(3)
		fmt.Println(i)
		gp.Done()

	}
	for i := 0; i < 10; i++ {
		gp.Add(1)
		go function(i)
	}
	gp.Wait() // Wait all gp.Add 's Done()

	/*
		time ./syncWait
		2
		9
		4
		0
		5
		3
		6
		7
		1
		8
		0.00user 0.00system 0:00.00elapsed 150%CPU (0avgtext+0avgdata 2104maxresident)k
		0inputs+0outputs (0major+257minor)pagefaults 0swaps
	*/
}

func main() {
	//wrong()
	toward()

}
