package programming

import (
	"fmt"
	"strconv"
)

func VariableExample() {
    // Variable signature
    var myVar int

    // Variable types
    // int8
    // int16
    // int32
    // int64
    // uint8
    // uint16
    // uint32
    // uint64
    // float32
    // float64

    // string

    // bool

    // byte

    // rune

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

    for _, val := range slice1{
        if val % 2 == 0 {
            count1++
        }
    }

    fmt.Printf("Even count in list %v is %v\n", slice1, count1)

    count2 := int64(0)
    slice2 := []int64{7, 8, 9, 10, 11, 12}

    for _, val := range slice2{
        if val % 2 == 0 {
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
        if val % 2 == 0 {
            count++
        }
    }

    return count
}

// struct types hold data in a structured way usually for some kind of real world object like a person.
type person struct {
    name string
    age int
}

// StructExample is a toy example of a person struct with a name and age.
func StructExample() person{
    warren := person{
        name: "Warren",
        age: 25,
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




