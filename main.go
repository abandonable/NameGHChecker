package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

func main() {
	filecontent, err := os.ReadFile("words.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	wordlist := strings.Split(string(filecontent), "\n")
	client := req.C()

	for _, words := range wordlist {
		checkign(strings.ReplaceAll(words, "\r", ""), client)
		time.Sleep(1 * time.Second)
	}
}

func checkign(username string, client *req.Client) {
	resp, err := client.R().SetHeader("Accept", "*/*").SetHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:145.0) Gecko/20100101 Firefox/145.0").Get(fmt.Sprintf("https://github.com/signup_check/username?value=%s", username))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch resp.StatusCode {
	case 200:
		fmt.Printf("[FREE] %s\n", username)
	case 422:
		/* For Debug Purposes
		Ignore this Line Otherwise */

		//fmt.Printf("[NOT FREE] %s\n", username)
		return
	case 429:
		fmt.Println("==== RATE LIMIT COOLDOWN... ====")
		time.Sleep(2 * time.Minute)
		fmt.Println("==== CONTINUE... ====")

	default:
		fmt.Println(resp.StatusCode)
		fmt.Println(resp)
		return
	}

}
