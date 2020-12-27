package cmd

import (
	"fmt"
	"os"
)

func errExit(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
