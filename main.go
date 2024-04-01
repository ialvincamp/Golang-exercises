package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	lowerLetterCase = "abcdefghijklmnopqrstuvwxyz"
	upperLettercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialLetter   = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	number          = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generatePassword(length int, useUpperLetters bool, useLowerLetters bool, useSpecial bool, useNum bool) string {
	var validChars string
	if useUpperLetters {
		validChars += lowerLetterCase
	}
	if useLowerLetters {
		validChars += lowerLetterCase
	}
	if useSpecial {
		validChars += specialLetter
	}
	if useNum {
		validChars += number
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = validChars[rand.Intn(len(validChars))]
	}
	return string(b)
}

func main() {
	var password string
	var inputString string
	var upperLetter bool
	var lowerLetter bool
	var specialLetter bool
	var number bool
	fmt.Println("Type Y or y, if you want to include this character / letter to your password")

	fmt.Println("Do you want to include this UpperCase Letter to your password?")
	fmt.Scan(&inputString)
	if strings.ToLower(inputString) == "y" {
		upperLetter = true
	}
	inputString = ""

	fmt.Println("Do you want to include this LowerCase Letter to your password?")
	fmt.Scan(&inputString)
	if strings.ToLower(inputString) == "y" {
		lowerLetter = true
	}
	inputString = ""

	fmt.Println("Do you want to include this Special Letter to your password?")
	fmt.Scan(&inputString)
	if strings.ToLower(inputString) == "y" {
		specialLetter = true
	}
	inputString = ""

	fmt.Println("Do you want to include this Number to your password?")
	fmt.Scan(&inputString)
	if strings.ToLower(inputString) == "y" {
		number = true
	}

	password = generatePassword(16, upperLetter, lowerLetter, specialLetter, number)
	fmt.Println("Random password:", password)
}
