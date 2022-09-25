package main

import (
	"fmt"
	"runtime"
)

func main() {
	// refactor with switch statement
	switch os := runtime.GOOS; os {
	case "linux", "darwin", "unix":
		fmt.Println("*nix variant")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s.\n", os)
	}
}
