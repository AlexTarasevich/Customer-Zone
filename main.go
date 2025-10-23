package main

import (
	"script_cz/cookie"
	"script_cz/credentials"
	"script_cz/download"
	"script_cz/path"

	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Constants

func main() {
	// 	if len(os.Args) < 3 || len(os.Args) > 4 {
	// 		fmt.Println(`Usage: ./customer_zone <command> <sdk_path> [<dest>]
	// Example: ./customer_zone download craft-sdk/macos/aarch64/3.3/sdk.tar.gz
	// Example: ./customer_zone download craft-sdk/macos/aarch64/3.3/sdk.tar.gz files/
	// Example: ./customer_zone download craft-sdk/macos/aarch64/3.3/sdk.tar.gz files/archive.tar.gz`)
	// 		os.Exit(1)
	// 	}

	command := os.Args[1]
	filePath := strings.TrimLeft(os.Args[2], "/")
	destPath := ""
	if len(os.Args) == 4 {
		destPath = os.Args[3]
	}

	// Only support "download"
	if command != "download" {
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Determine output path
	outputFile, err := path.ResolveOutputPath(filePath, destPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Download SDK:", filePath)
	fmt.Println("Will save to:", outputFile)

	// Get credentials
	username, password, err := credentials.GetCredentials()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Get session cookie
	sessionCookie, err := cookie.GetSessionCookie(username, password)
	if err != nil {
		fmt.Println("Error getting session:", err)
		os.Exit(1)
	}

	// Download file
	if err := download.DownloadFile(filePath, outputFile, sessionCookie); err != nil {
		fmt.Println("Error downloading file:", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Successfully downloaded %s → %s\n", filepath.Base(filePath), outputFile)
}
