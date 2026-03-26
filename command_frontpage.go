package main

import (
	"fmt"
	"log"
	"github.com/Pankajkumar2608/cli_go/hackerapi"
)

func callbackMap() error {
	hacknewapiClient := hackerapi.NewClient()
	resp, err := hacknewapiClient.FirstPageReq()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Front Page News:")
	for _,hits := range resp.Hits{
		fmt.Printf(" - %s\n", hits.Title)
		fmt.Printf(" - %s\n", hits.Author)
		fmt.Printf(" - %s\n", hits.StoryText)
		fmt.Printf(" - %s\n", hits.CreatedAt)
		fmt.Printf(" - %s\n", hits.URL)
	}
	return nil
}