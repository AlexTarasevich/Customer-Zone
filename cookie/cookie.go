package cookie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// Struct for API login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Query    string `json:"query"`
}

const apiURL = "https://www.tarantool.io/en/accounts/customer_zone/api"

// getSessionCookie logs in to API and extracts sessionid from Set-Cookie
func GetSessionCookie(username, password string) (string, error) {

	payload := LoginRequest{
		Username: username,
		Password: password,
		Query:    "query",
	}

	body, _ := json.Marshal(payload)
	fmt.Println("==> Preparing login request")
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "tt")

	client := &http.Client{}
	fmt.Println("==> Sending HTTP request to:", apiURL)
	resp, err := client.Do(req)
	fmt.Println("==> Got response or error:", err)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Extract sessionid from Set-Cookie
	for _, cookie := range resp.Header["Set-Cookie"] {
		if strings.HasPrefix(cookie, "sessionid=") {
			re := regexp.MustCompile(`sessionid=[^;]*`)
			match := re.FindString(cookie)
			if match != "" {
				return match, nil
			}
		}
	}
	return "", fmt.Errorf("no sessionid cookie found (status %d)", resp.StatusCode)
}
