package main

import (
    "fmt"
)

type contactInfo struct {
    email string
    zipCode int
}

type person struct {
    firstName string
    lastName string
    contactInfo // struct within struct
    // the above is the same as contactInfo contactInfo
}


func main() {
    //// relying on ordering of fields for assignment
    //alex := person{"Alex", "Anderson"}
    //
    //// assignment by field
    //john := person{
    //    firstName: "John",
    //    lastName: "Doe",
    //}
    //
    //fmt.Println(alex)
    //fmt.Println(john)

    // multiple line struct definition. All lines ends with a comma.
    jim := person{
        firstName: "Jim",
        lastName: "Party",
        contactInfo: contactInfo{
            email: "jim@gmail.com",
            zipCode: 91234,
        },
    }

    jim.print()

    jimPointer := &jim
    jimPointer.updateName("Foo")

    // Alternatively, we can just do
    // jim.updateName("Foo") <--- this is a go shortcut

    jim.print()
}


func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
}


func (p person) print() {
    // this is the syntax to print k-v
    fmt.Printf("%+v", p)
}