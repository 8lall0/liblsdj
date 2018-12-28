package main

import (
	"fmt"
	"github.com/8lall0/liblsdj"
	"os"
)

func main() {
	f, err := os.Open("test.lsdsng")
	if err != nil {
		fmt.Println("Errore")
		return
	}
	pro := liblsdj.ReadLsdsng(f)
	_ = f.Close()

	wr, err := os.Create("test2.lsdsng")
	if err != nil {
		fmt.Println("Errore")
		return
	}
	liblsdj.WriteLsdsng(wr, pro)
	_ = wr.Close()
}
