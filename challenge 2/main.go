package main

import "fmt"

func main(){
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i) // Nila i = 0,1,2,3,4
	}

	var j int = 0
	for j < 5 {
		fmt.Println("Nilai j = ", j) // Nilai j = 0,1,2,3,4
		j++
	}

	for pos, char := range "САШАРВО" {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos) //САШАРВО = 0,2,4,6,8,10,12
	}

	for {
		if j++; j <= (10) {
			fmt.Println("Nilai j = ", j) //Nilai j = 6,7,8,9,10
		} else {
			break
		}
	}
}
