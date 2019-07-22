package main

// ref: https://draveness.me/golang/basic/golang-interface.html

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c *Cat) Quack() {
	println(c.Name + " meow")
}

func main() {
	var c Duck = &Cat{Name: "grooming"}
	c.Quack()        // Duck type, dynamic patching
	c.(*Cat).Quack() // determine during compile time
}
