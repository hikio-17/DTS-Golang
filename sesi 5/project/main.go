package main

import (
	"errors"
	"fmt"
)

// func main() {
// 	var number int
// 	var err error

// 	number, err = strconv.Atoi("123agh")

// 	if err == nil {
// 		fmt.Println(number)
// 	} else {
// 		fmt.Println(err.Error())
// 	}

// 	number, err = strconv.Atoi("123")

// 	if err == nil {
// 		fmt.Println(number)
// 	} else {
// 		fmt.Println(err.Error())
// 	}
// }

func main() {
	defer catchErr()
	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfect")
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "validPassword", nil
}
