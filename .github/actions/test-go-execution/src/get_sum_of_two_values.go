package main
import (
	"fmt"
	"os"
)
func main(){
	fmt.Println("Welcome")
	args := os.Args
	fmt.Printf("Given arguments: %v\n", stringArrayToHumanizedString(args))
	var sum = args[0] + args[1]
	print(sum)
}

func stringArrayToHumanizedString(input []string) string {
	return "[" + strings.Join(input, ", ") + "]"
}
