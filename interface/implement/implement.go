package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "" {
		talk = ""
	} else {
		talk = think
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "yes"
	fmt.Println(peo.Speak(think))
}
