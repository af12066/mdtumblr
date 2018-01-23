package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/dghubble/oauth1"
	"gopkg.in/russross/blackfriday.v2"
)

// Tumblr access token (token credential) requests on behalf of a user
func main() {
	var (
		blogIDFlag string
		titleFlag  string
		stateFlag  string
	)
	flag.StringVar(&blogIDFlag, "u", "", "The blog hostname")
	flag.StringVar(&titleFlag, "t", "", "The title of the post")
	flag.StringVar(&stateFlag, "s", "published", "The state of the post. Specify one of the following:  published, draft, queue, private")
	flag.Parse()

	if blogIDFlag == "" {
		fmt.Print("Enter the blog hostname: ")
		blogIDFlag, _ = scanFlag()
	}
	if titleFlag == "" {
		fmt.Print("Enter the title of the post: ")
		titleFlag, _ = scanFlag()
	}

	// read credentials from environment variables
	consumerKey := os.Getenv("TUMBLR_CLIENT_ID")
	consumerSecret := os.Getenv("TUMBLR_CLIENT_SECRET")
	accessToken := os.Getenv("TUMBLR_ACCESS_TOKEN")
	accessSecret := os.Getenv("TUMBLR_ACCESS_SECRET")
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Missing required environment variable")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// httpClient will automatically authorize http.Request's
	httpClient := config.Client(oauth1.NoContext, token)

	// Read markdown file
	fname := flag.Arg(0)
	fmt.Printf("fname: %s\n", fname)
	md, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	mdconvert := blackfriday.Run([]byte(md), blackfriday.WithNoExtensions())

	// get information about the current authenticated user
	path := fmt.Sprintf("https://api.tumblr.com/v2/blog/%s/post", blogIDFlag)
	values := url.Values{}
	values.Add("type", "text")
	values.Add("state", stateFlag)
	values.Add("format", "html")
	values.Add("title", titleFlag)
	values.Add("body", string(mdconvert))
	resp, err := httpClient.PostForm(path, values)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
}

func scanFlag() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return text, nil
}
