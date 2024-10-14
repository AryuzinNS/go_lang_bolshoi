package main

import (
	"Go_lang/internal/pkg/storage"
	"fmt"
	"log"
)

func main() {
	st, err := storage.NewStorage()
	if err != nil {
		fmt.Println("Smth went wrooong wrooong")
		log.Fatal(err)
	}
	st.Set("1", "value1")
	rs1 := *st.Get("1")
	rs2 := st.GetKind("1")
	fmt.Println(rs1, rs2)
}
