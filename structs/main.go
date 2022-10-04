package main

import "fmt"

type MyStruct struct {
	name    string
	address string
	bal     float64
}

type Rectangle struct {
	length, height float64
}

type Contact struct {
	phone string
}

type Business struct {
	name    string
	address string
	Contact
}

func UsingMyStruct() MyStruct {
	var ts MyStruct
	ts.name = "Tom Smith"
	ts.address = "5 main st"
	ts.bal = 244.4
	fmt.Println("ts: ", ts) // ts:  {Tom Smith 5 main st 244.4}

	return ts
}

func GetInfo(ms MyStruct) {
	fmt.Printf("%s owes us %.2f\n", ms.name, ms.bal) // Tom Smith owes us 244.40
}

func EditAddress(ms *MyStruct, address string) {
	ms.address = address
}

func (r Rectangle) Area() float64 {
	return r.length * r.height
}

func (b Business) Info() {
	fmt.Printf("Contact at %s is %s", b.name, b.Contact.phone)
}

func main() {
	ts := UsingMyStruct()

	GetInfo(ts)
	EditAddress(&ts, "12 main st")
	fmt.Println("Address: ", ts.address) // Address:  12 main st

	ms := MyStruct{"New Person", "43 main st", 123}
	fmt.Println("Name: ", ms.name) // Name:  New Person

	newRectangle := Rectangle{10.0, 15.0}
	fmt.Println("newRectangle.Area(): ", newRectangle.Area()) // newRectangle.Area():  150

	// Composition is struct inside other struct:
	contact := Contact{"98123-0923"}
	business := Business{"Company", "12 main st", contact}
	business.Info() // Contact at Company is 98123-0923
}
