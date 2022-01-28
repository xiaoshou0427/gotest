package main

import (
	"fmt"
)

func init() {
	fmt.Println("I am init function")
}

func init() {
	fmt.Println("test2")
}
func main() {
	score1, score2, score3, score4, score5 := getScoresOfstudent("Oscar")
	fmt.Println(score1, score2, score3, score4, score5)
}

func getScoresOfstudent(stdName string) (chinese int, match int, english int, physic int, nature int) {
	fmt.Println("Below are your scores: ")
	chinese, match, english, physic, nature = 0, 0, 0, 0, 0
	return
}

func init() {
	fmt.Println("test3")
}
