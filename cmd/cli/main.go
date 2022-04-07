package main

import (
	"flag"
	"fmt"

	"github.com/shkwsk/algo/pkg/myalgo"
	"github.com/sirupsen/logrus"
)

func main() {
	var input int
	flag.IntVar(&input, "i", -1, "input int value")
	flag.Parse()
	if input < 0 {
		logrus.Fatal("Please input 0 or more.")
	}

	fmt.Println("Input: ", input)
	fmt.Println("IsPrime: ", myalgo.IsPrime(input))
	fmt.Println("DivisorEnumeration: ", myalgo.DivisorEnumeration(input))
	fmt.Println("PrimeFactrization: ", myalgo.PrimeFactrization(input))
}
