package main
import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// A slice of sample websites
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	for _, url := range urls {
		go checkUrl2(url)
	}

	time.Sleep(time.Second * 6)
}

//checks and prints a message if a website is up or down
func checkUrl2(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}