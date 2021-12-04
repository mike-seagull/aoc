package main
import (
	"os"
	"fmt"
	req "main/pkg"
)



func main() {
	token := os.Args[1:][0]
	req.Token = token
	input := req.GetTodaysInput()
	fmt.Printf(input)
}