package input

import (
	"bufio"
	"fmt"
	"os"
)

func BackMenu() {
	fmt.Println()
	fmt.Print("Press any key to back... ")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
}
