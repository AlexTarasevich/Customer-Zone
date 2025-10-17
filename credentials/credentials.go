package credentials

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// getCredentials loads username/password from env or credentials file
func GetCredentials() (string, string, error) {
	file, err := os.Open("credentials.txt")

	if err != nil {
		return "", "", fmt.Errorf("cannot open credentials.txt: %w", err)
	}
	defer file.Close()

	var username, password string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "user=") {
			username = strings.TrimPrefix(line, "user=")
		}
		if strings.HasPrefix(line, "password=") {
			password = strings.TrimPrefix(line, "password=")
		}
	}

	if username == "" || password == "" {
		return "", "", fmt.Errorf("invalid credentials file format: expected 'user=' and 'password=' lines")
	}
	return username, password, nil
}
