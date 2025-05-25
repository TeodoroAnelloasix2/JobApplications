package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	canal := make(chan string)

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
		"https://roma.it",
	}

	for _, api := range apis {
		go CheckApi(api, canal)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-canal)
	}

	alapsed := time.Since(start)

	fmt.Printf("Listo! Tomò %v segundos!", alapsed.Seconds())
}

func CheckApi(api string, ch chan string) {

	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Error: %s estàa caido!\n", api)
		return
	}
	ch <- fmt.Sprintf("Succes: %s està en funcionamiento!\n", api)
}
