package a

import (
	"fmt"
	"log"
)

func f() {
	// The pattern can be written in regular expression.

	err := hoge()
	if err != nil {
		log.Println(err)
	}

	hoge() // want "NG"

	_ = hoge() //want "NG"

	log.Println(hoge())

	i, err := hoge2()
	if err != nil {
		log.Println(i, err)
	}

	hoge2() // want "NG"

	log.Println(hoge2())

	fmt.Println("Hello")

	i2, _ := hoge2() // want "NG"
	log.Println(i2)
}

func hoge() error {
	return nil
}

func hoge2() (int, error) {
	return 0, nil
}
