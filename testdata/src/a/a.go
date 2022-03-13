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

	_ = hoge() //want "blank error"

	log.Println(hoge())

	i, err := hoge2()
	if err != nil {
		log.Println(i, err)
	}

	hoge2() // want "NG"

	log.Println(hoge2())

	fmt.Println("Hello")

	i2, _ := hoge2() // want "blank error"
	log.Println(i2)

	_, _ = hoge2() // want "blank error"

	_, _ = hoge3() // want "blank error"

	_, _ = hoge(), hoge() // want "blank error"

	_, _ = hoge(), hoge4() // want "blank error"
}

func hoge() error {
	return nil
}

func hoge2() (int, error) {
	return 0, nil
}

func hoge3() (error, int) {
	return nil, 0
}

func hoge4() int {
	return 0
}
