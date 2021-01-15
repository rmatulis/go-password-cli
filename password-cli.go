package main

import (
	"flag"
	"fmt"
	"github.com/sethvargo/go-password/password"
	"math/rand"
)

func main() {

	numbPtr := flag.Int("length", 20, "Length of password.")

	// Use symbols or not to
	symbols := flag.Bool("symbols", false, "If used will set symbols in password")

	// allow repeat charters
	repeat := flag.Bool("repeat", false, "If used will allow repeat charters and numbers")
	flag.Parse()

	var symbolGen int
	// if symbol is true then set the number that should be generated
	if *symbols == true {
		symMax := 10
		symMin := 3
		symbolGen = rand.Intn(symMax-symMin) + symMin
	}

	res, err := password.Generate(*numbPtr, 10, symbolGen, false, *repeat)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(res)
}
