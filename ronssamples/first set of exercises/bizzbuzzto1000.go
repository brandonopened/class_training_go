package main

import "fmt"

func main() {
    
    total := 0	

	for i := 0; i <= 1000; i++ {
		if i%3 == 0 {
			total = total +i
		} 
		if i%5 == 0 {
			total = total +i
		} 

	}
	fmt.Println(total)
}
