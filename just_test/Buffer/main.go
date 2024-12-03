package main

import (
	"bytes"
	"fmt"
)

func main() {
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error%s", err)
	// }
	// defer f.Close()

	// s := []byte("Hello gophers!!!")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error%s", err)
	// }

	b := bytes.NewBufferString("Hello gophers,")
	fmt.Println(b.String())
	b.WriteString("Gophers!!")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("XXX")
	fmt.Println(b.String())

	b.Write([]byte("YYY"))
	fmt.Println(b.String())

}
