package main

import "log"

func main() {

	var data []int
	data = append(data, 3)
	data = append(data, 4)
	data = append(data, 3)
	data = append(data, 1)
	data = append(data, 9)
	data = append(data, 27)
	data = append(data, 2)
	data = append(data, 10)


	log.Println("Input Array - ", data)

	for i := 1; i < len(data) ; i++ {
		tmp := data[i]
		log.Println("Checking - ", tmp )

		for j := i-1; j>=0; j-- {
			if data[j] > data[j+1] {
				data[j+1] = data[j]
				data[j] = tmp
			}
			log.Println("Array sorted upto i=", i, data)
		}
	}

	log.Println("Sorted array - ", data)
}
