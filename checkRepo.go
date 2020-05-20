package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: checkRepo <owner>/<repoName>")
		os.Exit(2)
	}
	repoName := os.Args[1]
	token, found := os.LookupEnv("API_TOKEN")
	if !found {
		fmt.Println("Must set API_TOKEN environment variable")
		os.Exit(2)
	}

	url := "https://gitea.pfdev.de/api/v1/repos/" + repoName
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "token "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		os.Exit(1)
	}
}
