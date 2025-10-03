package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://api.sse-server.com/api/events?token=JWT_TOKEN")
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println("Received:", scanner.Text())
	}
}
