package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var lang string
	if val, ok := os.LookupEnv("weather_lang"); ok {
		lang = val + "."
	}

	url := fmt.Sprintf("http://%swttr.in/%s", lang, strings.Join(os.Args[1:], "+"))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// pretend to be curl to get the ASCII representation of the weather report
	req.Header.Set("User-Agent", "curl")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
