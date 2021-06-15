The Go Playground
 Imports

Hello, playground
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}
type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	jim := Person{
		firstName: "jim",
		lastName:  "party",
		contact: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 110096,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("qsstech")
	jim.print()
}
func (p Person) print() {
	fmt.Println("%+v", p)
}
func (pointerToPerson *Person) updateName(NewFirstName string) {
	fmt.Println("InSide update :", NewFirstName)
	(*pointerToPerson).firstName = NewFirstName
}

