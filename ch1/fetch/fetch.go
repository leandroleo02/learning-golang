package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fetch(urls []string) {
	for _, url := range urls {
		resp, err := http.Get(ensureUrlFormat(url))
		fmt.Printf("Response status code: %d\n", resp.StatusCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		w, err := io.Copy(os.Stdout, resp.Body)
		fmt.Printf("%d", w)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func ensureUrlFormat(url string) string {
	if !strings.HasPrefix(url, "https://") {
		return "https://" + url
	}
	return url
}
