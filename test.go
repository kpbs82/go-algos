package main

import (
	"log"
	"strings"
)

func main() {
	//result := bracket_match("(())")
	//log.Println("Returned - ", result)

	//almost_palindromes("abba")

	//ascii_deletion_distance("thought", "sloughs")

	four_letter_words("This is test")
}


func bracket_match(bracket_string string) int {

	if bracket_string[0] != '(' || bracket_string[len(bracket_string)-1] != ')' {
		return 0
	}

	open := 0
	close := 0
	for _,s := range bracket_string {

		log.Println("S - ", s)
		if s == '(' {
			open++
		} else if s == ')' {
			close ++
		}
	}

	if open == close {
		return 1
	}

	return 0
}

func almost_palindromes(str string) int {

	diff := 0
	start := 0
	end := len(str)-1

	for i:=start; i<len(str); i++ {

		if str[start] != str[end] {
			diff++
		}
		start++
		end--
	}

	log.Println("Diff - ", diff)
	return diff

}

func ascii_deletion_distance(str1 string, str2 string) int {

	total := 0

	//var newStr1, newStr2 string

	if str1 == str2 {
		return total
	}

	// search chars from str1 not in str2
	for i:=0; i<len(str1); i++ {
		if !strings.Contains(str2, string(str1[i])){
			log.Println("Not in str2 - ", string(str1[i]))
			total = total + int(str1[i])
		}
	}

	// search chars from str2 not in str1
	for i:=0; i<len(str2); i++ {
		if !strings.Contains(str1, string(str2[i])){
			total = total + int(str2[i])
			log.Println("Not in str1 - ", string(str2[i]))
		}
	}

	//log.Println("New str 1 - ", newStr1, newStr2)
	//
	//if newStr1 != newStr2 {
	//	for i:=0; i<len(str1); i++ {
	//		if newStr1[i] != newStr2[i] {
	//			total = total + int(newStr1[i])
	//		}
	//	}
	//}

	return total
}

func four_letter_words(sentence string) int {

	total := 0

	split_sentence := strings.Split(sentence, " ")

	for i:=0; i<len(split_sentence); i++ {
		log.Println("SPlit - ", split_sentence[i])
		if len(split_sentence[i]) == 4 {
			total = total + 1
		}
	}

	log.Println("Total - ", total)
	return total
}