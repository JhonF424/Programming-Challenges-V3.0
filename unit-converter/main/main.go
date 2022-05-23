package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("---------------")
	fmt.Println("Unit Converter")
	fmt.Println("---------------")
	for {
		fmt.Println("Choose a category")
		fmt.Println("0. Weight")
		fmt.Println("1. Area")
		fmt.Println("2. Lenght")
		fmt.Println("99. Exit")
		fmt.Print("> ")
		key, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		key = strings.TrimSpace(key)
		switch key {
		case "0":
			fmt.Println("---------------")
			fmt.Println("Witght Converter")
			fmt.Println("---------------")
			weight()
			fmt.Println("---------------")
		case "1":
			fmt.Println("---------------")
			fmt.Println("Area Converter")
			fmt.Println("---------------")
		case "2":
			fmt.Println("---------------")
			fmt.Println("Lenght Converter")
			fmt.Println("---------------")
			lenght()
			fmt.Println("---------------")
		case "99":
			fmt.Println("Good Bye!")
			os.Exit(0)
		default:
			fmt.Println("This is not an option!")
		}
	}
}

func weight() {
	fmt.Println("0. Pounds -> Kg")
	fmt.Println("1. Kg -> Pounds")
	fmt.Println("2. Grams -> Kg")
	fmt.Println("3. Kg -> Grams")
	fmt.Println("4. Kg -> Tons")
	fmt.Println("5. Tons -> Kg")
}

func lenght() {
	fmt.Println("0. Metters -> Kilometters")
	fmt.Println("1. Kilometters -> Metters")
	fmt.Println("2. cm -> mm")
	fmt.Println("3. mm -> cm")
	fmt.Println("4. cm -> Metters")
	fmt.Println("5. Metters -> cm")
	fmt.Println("6. Miles -> Kilometters")
	fmt.Println("7. Km -> Miles")
	fmt.Println("8. Yards -> Km")
	fmt.Println("9. Km -> Yards")
}
