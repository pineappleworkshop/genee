package cmd

import (
	"fmt"
	"genee/services"
	"os"
)

func errExit(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func errExitClean(msg interface{}, destination string) {
	services.DeleteRoot(destination)
	fmt.Println("Error:", msg)
	os.Exit(1)
}
