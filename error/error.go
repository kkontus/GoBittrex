package helpers

import (
	"fmt"
	"os"
)

func ShowError(err error) {
	fmt.Printf("%s", "Error message: ")
	fmt.Printf("%s\n", err)
	os.Exit(1)
}
