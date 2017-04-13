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

	mergeSort(data)
}


func mergeSort (data []int) {

	log.Println("Input array - ", data)

	if len(data) < 2 {
		log.Println("Single element sorted array - ", data)
		return
	}

	// split the array
	mid := len(data) / 2

	// call merge recursively
	mergeSort(data[:mid])
	mergeSort(data[mid:])

	log.Println("Data to be worked on - ", data)

	if (data[mid-1] < data[mid]) {
		log.Println("Already Sorted - ", data)
		return
	}

	// copy the first half to a slice - tmp memory
	s := make([]int, mid)
	copied := copy(s, data[:mid])
	if copied < 1 {
		log.Println("Copy did not work, mid - ", mid)
		log.Println(" New slice - ", s)
	}

	// set left and right indices to move through
	l := 0
	r := mid
	for i:=0; ; i++ {
		if s[l] < data[r] {
			data[i] = s[l]
			l++
			if (l == mid) {
				break
			}
		} else {
			data[i] = data[r]
			r++
			if ( r == len(data)) {
				copy(data[i+1:], s[l:mid])
				break
			}
		}
	}
	log.Println("Sorted - ", data)
	return
}
