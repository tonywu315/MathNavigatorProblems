package main

import (
	"fmt"
	"github.com/tonywu315/MathNavigatorProblems/golang"
)

func main() {
	fmt.Println(golang.FindSubstring("googlygoggles", "googly"))
	fmt.Println(golang.Split([]int{1, 2, 3, 4},[]int{1, 3, 5}))
	golang.PrintCandidateInfo([]string{"A", "B", "B", "B", "C", "C", "D"})
}
