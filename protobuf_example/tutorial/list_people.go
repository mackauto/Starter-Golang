package main

import (
	"fmt"
	pb "github.com/JPMike/Starter-Golang/protobuf_example/tutorial/go_proto"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func writePerson(w io.Writer, p *pb.Person) {
	_, _ = fmt.Fprintln(w, "Person ID:", p.Id)
	_, _ = fmt.Fprintln(w, "Name:", p.Name)
	if p.Email != "" {
		_, _ = fmt.Fprintln(w, "E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.Person_MOBILE:
			_, _ = fmt.Fprint(w, "Mobile phone #:")
		case pb.Person_HOME:
			_, _ = fmt.Fprint(w, "Home phone #:")
		case pb.Person_WORK:
			_, _ = fmt.Fprint(w, "Work phone #:")
		}
		_, _ = fmt.Fprintln(w, pn.Number)
	}
}

func listPeople(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book", err)
	}
	listPeople(os.Stdout, book)
}
