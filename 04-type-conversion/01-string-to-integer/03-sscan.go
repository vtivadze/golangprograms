package main

import (
	"fmt"
	"reflect"
)

func main() {
	strVar := "100"
	intVal := 0
	_, err := fmt.Sscan(strVar, &intVal)
	fmt.Println(intVal, err, reflect.TypeOf(intVal))
}
