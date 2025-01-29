package main

import (
	"fmt"
	"myProg/array_dt"
)

func main() {
	var a []array_dt.Array_DT
	val1 := &array_dt.Value{}
	val1.Add(1)
	val2 := &array_dt.Value{}
	val2.Add(1.3)
	val3 := &array_dt.Value{}
	val3.Add("asd")
	val4 := &array_dt.Value{}
	val4.Add(true)

	a = append(a, val1)
	a = append(a, val2)
	a = append(a, val3)
	a = append(a, val4)

	for _, v := range a {
		fmt.Println(v.Get())
	}
}
