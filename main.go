package main

import "fmt"
import C "github.com/zugarzeeker/go-config/config"

func main() {
	// fmt.Println("[key1]", C.GetEnv("key1"))
	fmt.Println("[key1]", C.GetEnv("key", "helloooo"))
}
