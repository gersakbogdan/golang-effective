/*
	Go doesn't provide automatic support for getters and setters.
	There's nothing wrong with providing getters and setters yourself, and it's often appropriate to do so,
	but it's neither idiomatic nor necessary to put Get into the getter's name. If you have a field called
	owner (lower case, unexported), the getter method should be called Owner (upper case, exported),
	not GetOwner. The use of upper-case names for export provides the hook to discriminate the field from
	the method. A setter function, if needed, will likely be called SetOwner.
*/
package main

import "fmt"

type Person struct {
	name string
}

func (p Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func main() {
	me := Person{"Bogdan"}
	fmt.Println("My name is: ", me.Name())

	me.SetName("Adrian")
	fmt.Println("My new name is: ", me.Name())
}
