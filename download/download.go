package download

// Функция которая скачивает файл по URL, сформированную из downloadURL + filePath
// передает cookie сессии в запросе и записывает все в outputFile
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"strings"
)

const downloadURL = "https://www.tarantool.io/en/accounts/customer_zone/packages/"

func DownloadFile(filePath, outputFile, sessionCookie string) error {
	url := downloadURL + filePath
	fmt.Println("Downloading:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "tt")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", sessionCookie)

	client := &http.Client{}
	// 	Timeout: 60 * time.Second,
	// }
	fmt.Println("==> Creating request:", url)
	resp, err := client.Do(req)
	fmt.Println("==> Got response status:", resp.StatusCode)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP error %d while downloading file", resp.StatusCode)
	}

	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// Check if downloaded file is HTML (error page)
	checkFile, _ := os.Open(outputFile)
	defer checkFile.Close()
	reader := bufio.NewReader(checkFile)
	for i := 0; i < 50; i++ {
		line, err := reader.ReadString('\n')
		if strings.Contains(line, "<!DOCTYPE html>") {
			return errors.New("server returned HTML page instead of file (check path or credentials)")
		}
		if err != nil {
			break
		}
	}
	return nil
}
