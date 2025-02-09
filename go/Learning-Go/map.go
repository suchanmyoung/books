package main

import "fmt"

func maps() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1

	totalWins2 := make(map[string]int)
	totalWins2["Orcas"] = 2
	v, ok := totalWins2["Orcas"]

	fmt.Println(totalWins, totalWins2)
	fmt.Println(v, ok)
}

func main() {
	maps()
}
