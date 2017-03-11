package main

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type URLReader interface {
	PageTitle(url string) (string, error)
}

type DefaultURLReader struct{}

func (r *DefaultURLReader) PageTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.StartTagToken:
			tn, _ := z.TagName()
			if strings.ToLower(string(tn)) == "title" {
				z.Next()
				return string(z.Raw()), nil
			}
		case html.ErrorToken:
			err = z.Err()
			if err == io.EOF {
				return "", nil
			}
			return "", err
		}
	}

	return "", nil
}
