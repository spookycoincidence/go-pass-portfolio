package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sethvargo/go-password/password"
)

func main() {
	length := flag.Int("length", 16, "length of the password")
	numDigits := flag.Int("digits", 4, "number of digits in the password")
	numSymbols := flag.Int("symbols", 2, "number of symbols in the password")
	noUpper := flag.Bool("no-upper", false, "disable uppercase letters")
	allowRepeat := flag.Bool("repeat", false, "allow characters to repeat")
	count := flag.Int("count", 1, "number of passwords to generate")

	flag.Parse()

	for i := 0; i < *count; i++ {
		res, err := password.Generate(*length, *numDigits, *numSymbols, *noUpper, *allowRepeat)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Password %d: %s\n", i+1, res)
	}
}
