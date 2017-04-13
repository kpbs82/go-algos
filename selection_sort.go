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
		smallest := data[i]
		log.Println("Look for smaller than - ", smallest )
		for j:= i+1; j<len(data); j++ {
			if smallest > data[j] {
				smallest = data[j]
				data[j] = data[i]
			}
		}
		data[i] = smallest
		log.Println("Array now - ", data)
	}
	log.Println("Final Sorted array - ", data)
}

