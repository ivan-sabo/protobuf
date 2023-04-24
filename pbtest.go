package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ivan-sabo/protobuf/examples/go/apiv1"
	"google.golang.org/protobuf/proto"
)

var fileName string = "pb_test"

func main() {
	writeFile()
	readFile()
}

func readFile() {
	finf, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("an error occured: %e\n", err)
	}

	reader := bufio.NewReader(finf)
	data := make([]byte, 1000)
	_, err = reader.Read(data)
	if err != nil && err != io.EOF {
		log.Fatalf("an error occured: %e\n", err)
	}

	book := &apiv1.AddressBook{}
	err = proto.Unmarshal(data, book)
	if err != nil {
		log.Fatalf("an error occured: %e\n", err)
	}

	fmt.Println(book)
}

func writeFile() {
	mark := &apiv1.Person{
		Name:  "Mark",
		Id:    4192,
		Email: "mark@gmail.com",
		Phones: []*apiv1.Person_PhoneNumber{
			{
				Number: "+4915222333",
				Type:   apiv1.Person_MOBILE,
			},
		},
	}

	book := apiv1.AddressBook{
		People: []*apiv1.Person{
			mark,
		},
	}

	out, err := proto.Marshal(&book)
	if err != nil {
		log.Fatalf("an error occured: %e\n", err)
	}

	finfo, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("an error occured: %e\n", err)
	}

	buf := bufio.NewWriter(finfo)
	_, err = buf.Write(out)
	if err != nil {
		log.Fatalf("an error occured: %e\n", err)
	}
}
