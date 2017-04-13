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

	for i:=0; i<len(data); i++ {
		for j:=0; j<len(data) -1; j++ {
			if data[j] > data[j+1] {
				log.Println("Swapping - ", data[j], data[j+1])
				tmp := data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
			}
		}
		log.Println(" Array - ", data)
	}
	log.Println("Sorted Array - ", data)
}
