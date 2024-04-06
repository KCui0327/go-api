package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetCredentials() (string, string, string) {
	file, err := os.Open("utils/config.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", ""
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	username := scanner.Text()
	scanner.Scan()
	password := scanner.Text()
	scanner.Scan()
	ip := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return "", "", ""
	}

	return username, password, ip
}
