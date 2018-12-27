package main

import (
	"fmt"
	"github.com/8lall0/liblsdj"
	"os"
)

func main() {
	f, err := os.Open("A.Day.With.You.lsdsng")
	if err != nil {
		fmt.Println("Errore")
		return
	}
	_ = liblsdj.ReadLsdsng(f)

}
