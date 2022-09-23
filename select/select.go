package main

import (
	"fmt"
	"time"
)

// modelo de chamada com canal 
func main(){
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			canal1 <- "A ----"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal1 <- "B ***********************************"
		}
	}()

	// for{
	// 	mensagemCanal1 := <-canal1
	// 	fmt.Println(mensagemCanal1)

	// 	mensagemCanal2 := <-canal2
	// 	fmt.Println(mensagemCanal2)
	// }

	for{
		select{
			case mensagemCanal1 := <-canal1:
				fmt.Println(mensagemCanal1)

			case mensagemCanal2 := <-canal2:
				fmt.Println(mensagemCanal2)
		}
	}

}