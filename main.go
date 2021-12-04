package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	var counter int32

	if len(os.Args) != 3 {
		fmt.Println("It's mandatory provide 2 arguments")
		fmt.Println("Example: [executable] [url] [characther to find]")
		fmt.Println("Example: counter_http.exe https://es.wikipedia.org/wiki/PHP a")
		return
	}

	url := os.Args[1]
	ch := os.Args[2]

	fmt.Printf("url: %s\ncharacter: %s\n", url, ch)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	formatted := string(b)

	for _, s := range formatted {
		if string(s) == ch {
			counter++
		}
	}

	fmt.Printf("The number of \"%s\" in \"%s\" is: %d\n", ch, url, counter)
}
