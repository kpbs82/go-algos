package longest_word

import (
	"log"
	"strings"
)

var WordsInGrid []string
var Grid [][]rune

func BuildGrid() {
	Grid = [][]rune{
		{'Q', 'W', 'E', 'R', 'T', 'N', 'U', 'I'},
		{'O', 'P', 'A', 'A', 'D', 'F', 'G', 'H'} ,
		{'T', 'K', 'L', 'Z', 'X', 'C', 'V', 'B'} ,
		{'N', 'M', 'R', 'W', 'F', 'R', 'T', 'Y'} ,
		{'U', 'I', 'O', 'P', 'A', 'S', 'D', 'F'} ,
		{'G', 'H', 'J', 'O', 'L', 'Z', 'X', 'C'} ,
		{'V', 'B', 'N', 'M', 'Q', 'W', 'E', 'R'} ,
		{'T', 'Y', 'U', 'I', 'O', 'P', 'A', 'S'} ,
	}

	log.Println("Grid - ", Grid)
}


func SearchWordsinGrid()  {

	//rows := 8
	//columns := 8

	//i:=3
	//j:=4
	// starting at each element in grid
	//for  i := 0; i < rows; i++ {
	//	for j := 0; j < columns; j++ {
	//		start := strings.ToLower(string(Grid[i][j]))
			start := "f"

			log.Println("Looking for all words starting with - ", start)

			if Dict.traverse(start) {
				log.Println("Found words starting with - ", start)
				//walkLikeKnight(start, i+2, j+1)
			} else {
				//continue
			}
		//}
	//}

}


func findInPrefixDict(word string) bool {
	exists := Dict.PrefixExists(word)
	if exists {
		return true

	} else {
		return false
	}
}

func findInDict(word string) bool {
	exists := Dict.Exists(word)
	if exists {
		return true

	} else {
		return false
	}
}

func walkLikeKnight(prefix string, i int, j int) bool {

	rows := 8
	columns := 8

	if i >= rows || j >= columns || i < 0 || j < 0 {
		return false
	}

	word := prefix + strings.ToLower(string(Grid[i][j]))
	log.Println("Looking for - ", word)


	if findInDict(word) {
		log.Println("Found - ", word)
	} else {
		//return true
	}

	//if walkLikeKnight(prefix, i+2, j+1) {
	//	log.Println("Looking for -1  - ", word)
	//	return true
	//}
	if walkLikeKnight(word, i-2, j+1) {
		log.Println("Looking for 0 - ", word)
		return true
	}
	if walkLikeKnight(word, i-2, j-1) {
		log.Println("Looking for 1 - ", word)
		return true
	}
	if walkLikeKnight(word, i+1, j+2) {
		log.Println("Looking for 2 - ", word)
		return true
	}
	if walkLikeKnight(word, i+1, j-2) {
		log.Println("Looking for 3 - ", word)
		return true
	}
	if walkLikeKnight(word, i-1, j+2) {
		log.Println("Looking for 4 - ", word)
		return true
	}
	if walkLikeKnight(word, i-1, j-2) {
		log.Println("Looking for 5 - ", word)
		return true
	}

	return false

}


//func walkLikeKnight(prefix string, i int, j int) {
//
//	rows := 8
//	columns := 8
//
//	log.Println("Find all words starting with - ", prefix)
//
//	found := findInDict(prefix, Dict)
//	if found {
//		r := i + 1
//		c := j + 2
//		if r < rows && c < columns {
//			word := prefix + string(Grid[r][c])
//			walkLikeKnight(word, r, c)
//		}
//	}
//
//
//	//word = append(word, grid[i+1][j+2])
//	//word = append(word, grid[i+1][j-2])
//	//word = append(word, grid[i+2][j+1])
//	//word = append(word, grid[i+2][j-1])
//	//word = append(word, grid[i-1][j+2])
//	//word = append(word, grid[i-1][j-2])
//	//word = append(word, grid[i-2][j+1])
//	//word = append(word, grid[i-2][j-1])
//
//}

//switch  {
//case 1:
//r := i+1
//c := j+2
//if r<rows && c<columns {
//word := prefix + string(Grid[r][c])
//walkLikeKnight(word, r, c)
//}
//
//case 2:
//r := i+1
//c := j-2
//if r<rows && c<columns {
//word := prefix + string(Grid[r][c])
//walkLikeKnight(word, r, c)
//}
//
//case 3:
//r := i+2
//c := j+1
//if r<rows && c<columns {
//word := prefix + string(Grid[r][c])
//walkLikeKnight(word, r, c)
//}
//case 4:
//r := i+2
//c := j-1
//if r<rows && c<columns {
//word := prefix + string(Grid[r][c])
//walkLikeKnight(word, r, c)
//}
//case 5:
//r := i-1
//c := j+2
//if r<rows && c<columns {
//word := prefix + string(Grid[r][c])
//walkLikeKnight(word, r, c)
//}
//case 6:
//r := i-1
//c := j-2
//if r<rows && c<columns {
//word := prefix + string(Grid[r][c])
//walkLikeKnight(word, r, c)
//}
//case 7:
//r := i-2
//c := j+1
//if r<rows && c<columns {
//word := prefix + string(Grid[r][c])
//walkLikeKnight(word, r, c)
//}
//case 8:
//r := i-2
//c := j-1
//if r<rows && c<columns {
//word := prefix + string(Grid[r][c])
//walkLikeKnight(word, r, c)
//}
//}