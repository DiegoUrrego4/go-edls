package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag pattern
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "all files including hidden files")
	flagNumberRecords := flag.Int("n", 0, "numbers of records")

	// order flags
	hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	hasOrderBySize := flag.Bool("s", false, "sort by size, smallest first")
	hasOrderReverse := flag.Bool("r", false, "reverse order while sorting")

	flag.Parse() // Mapea cada uno de los flags y almacena en variables

	fmt.Println("pattern:", *flagPattern)
	fmt.Println("all:", *flagAll)
	fmt.Println("flagNumberRecords:", *flagNumberRecords)
	fmt.Println("hasOrderByTime:", *hasOrderByTime)
	fmt.Println("hasOrderBySize:", *hasOrderBySize)
	fmt.Println("hasOrderReverse:", *hasOrderReverse)
}
