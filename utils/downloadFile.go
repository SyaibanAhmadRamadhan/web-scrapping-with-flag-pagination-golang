package utils

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func BuildFileName(fullUrlFile string) *os.File {
	fileUrl, err := url.Parse(fullUrlFile)
	if err != nil {
		panic(err)
	}

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	file, err := os.Create("images/" + segments[len(segments)-1])
	if err != nil {
		panic(err)
	}
	return file
}

func PutFile(file *os.File, fullUrlFile string) {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fullUrlFile)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)

	defer file.Close()

	if err != nil {
		panic(err)
	}

	// log.Printf("download file : %v | %d", file.Name(), size)
}
