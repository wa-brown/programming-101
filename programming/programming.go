package programming

import (
	"fmt"
	"strconv"
)

func VariableExample() {
	// Variable signature - var {name} {type}
	var myVar int

	// Variable types
	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32

	// How to use a variable
	myVar = 1

	fmt.Println(myVar)

	myString := "Warren's 101 is the best :)" // Go complains that I haven't used the variable

	// So i need to use this variable in my code, this is an IF statement.
	fmt.Println(myString)

}

// LoopExample is a toy example of using loops with slices and maps.
func LoopExample() {
	// More types
	// list/slice
	mySlice := []int64{1, 2, 3, 4, 5, 6, 7}
	// maps
	myMap := make(map[string]float64)

	// If
	if len(mySlice) > 0 {
		mySlice[0] = 8
	}

	// Range over slice/list
	for i, val := range mySlice {
		myMap[strconv.FormatInt(val, 10)] = float64(val)
		fmt.Printf("Index: %d -- Value: %v\n", i, val)
	}

	// Range over map
	for key, val := range myMap {
		fmt.Printf("Key: %v -- Value: %v\n", key, val)
	}
}

// FunctionExample is a toy example of why functions are used.
func FunctionExample() {
	// I want to count the even numbers in a list
	count1 := int64(0)
	slice1 := []int64{1, 2, 3, 4, 5, 6}

	for _, val := range slice1 {
		if val%2 == 0 {
			count1++
		}
	}

	fmt.Printf("Even count in list %v is %v\n", slice1, count1)

	count2 := int64(0)
	slice2 := []int64{7, 8, 9, 10, 11, 12}

	for _, val := range slice2 {
		if val%2 == 0 {
			count2++
		}
	}

	fmt.Printf("Even count in list %v is %v\n", slice2, count2)

	slice3 := []int64{13, 14, 15, 16, 17, 18}
	fmt.Printf("Even count using countEven function: %v\n", countEven(slice3))
}

// Notice that this function definition is not capitalised
func countEven(usrSlice []int64) int64 {
	count := int64(0)

	for _, val := range usrSlice {
		if val%2 == 0 {
			count++
		}
	}

	return count
}

// struct types hold data in a structured way usually for some kind of real world object like a person.
type person struct {
	name string
	age  int
}

// StructExample is a toy example of a person struct with a name and age.
func StructExample() person {
	warren := person{
		name: "Warren",
		age:  25,
	}

	fmt.Printf("My struct: %v\n", warren)
	fmt.Printf("My name: %s\n", warren.name)
	fmt.Printf("My age: %d\n", warren.age)

	return warren
}

// Methods are functions that can be called by our objects and usually interact with the struct.
func (p person) printAge() {
	fmt.Printf("My age is %d\n", p.age)
}

// Notice that this function reciever uses a '*' or pointer
func (p *person) haveBirthday() {
	p.age++
}

func MethodExample(myPerson person) {
	myPerson.printAge()
	myPerson.haveBirthday()
	myPerson.printAge()
}
