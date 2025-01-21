package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com/",
		"https://dev.azure.com/",
		"https://api.github.com/",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com/",
	}

	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	// fmt.Println(<-ch)

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	// time.Sleep(5 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Tiempo de ejecución: %v\n", elapsed.Seconds())

}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		// fmt.Printf("Error al revisar la API %s\n", api)
		ch <- fmt.Sprintf("Error al revisar la API %s\n", api)
		return
	}

	// fmt.Printf("API %s revisada correctamente\n", api)
	ch <- fmt.Sprintf("API %s revisada correctamente\n", api)
}
