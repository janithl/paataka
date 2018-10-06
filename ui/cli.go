package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/janithl/paataka/entities"
	"github.com/janithl/paataka/usecases"
)

// CLI struct implements the CLI
type CLI struct {
	PublicationService usecases.PublicationService
	reader             *bufio.Reader
}

type option struct {
	name     string
	function func()
}

// GetInput gets user input via CLI
func (c *CLI) GetInput() {
	c.reader = bufio.NewReader(os.Stdin)

	options := []option{
		option{"Add Publication", c.addPublication},
		option{"List All Publications", c.listAll},
		option{"Fetch All Posts", c.fetchAll},
	}

	for {
		fmt.Println("")
		for index, opt := range options {
			fmt.Printf("[%1d] %s\n", index+1, opt.name)
		}
		fmt.Println("[q] Quit")

		fmt.Print("\nEnter Command: ")

		userInput, _ := c.reader.ReadString('\n')
		userInput = userInput[:len(userInput)-1]

		if userInput == "q" {
			return
		}

		if input, err := strconv.Atoi(userInput); err != nil || input < 1 || input > len(options) {
			fmt.Println("Incorrect Option Selected")
		} else {
			options[input-1].function()
		}
	}
}

func (c *CLI) addPublication() {
	fmt.Print("Enter Title: ")
	title, _ := c.reader.ReadString('\n')

	fmt.Print("Enter URL: ")
	url, _ := c.reader.ReadString('\n')

	c.PublicationService.Add(entities.Publication{Title: title[:len(title)-1], URL: url[:len(url)-1]})
}

func (c *CLI) listAll() {
	pubs := c.PublicationService.ListAll()
	if len(pubs) == 0 {
		fmt.Println("No Publications Yet")
	} else {
		fmt.Println("\nPublications:")
	}

	for _, pub := range pubs {
		fmt.Printf("%s (%s)\n", pub.Title, pub.URL)
	}
}

func (c *CLI) fetchAll() {
	pubs := c.PublicationService.ListAll()
	for _, pub := range pubs {
		fmt.Printf("\nFetching %s (%s)\n", pub.Title, pub.URL)
		err := c.PublicationService.FetchPublicationPosts(pub)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
