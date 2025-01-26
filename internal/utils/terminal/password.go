package terminal

import (
	"fmt"
	"golang.org/x/term"
	"syscall"
)

func HidePassword() ([]byte, error) {
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}
	fmt.Println()
	return passwordBytes, nil
}
