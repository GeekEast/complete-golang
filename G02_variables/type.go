package G02_variables

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func InitializeType() {
	c1 := Celsius(-273.50) // short variable
	println(c1)            // will print the underlying type
	var c2 Celsius
	println(c2) // initialized as the zero value of underlying type
	const c3 Celsius = 23.2
	println(c3)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func ConvertTypes() (func(c Celsius) Fahrenheit, func(f Fahrenheit) Celsius) {
	CToF := func(c Celsius) Fahrenheit {
		return Fahrenheit(c*9/5 + 32)
	}

	FToC := func(f Fahrenheit) Celsius {
		return Celsius((f - 32) * 5 / 9)
	}

	return CToF, FToC
}

func CompareTypes() {
	var c Celsius
	var f Fahrenheit
	res := f - CToF(c)
	println(res == Fahrenheit(30))
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func AttachFuncToType() {
	c := Celsius(30)
	println(c) // will call the String() method automatically
}
