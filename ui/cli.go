package ui

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		option{"List All Publications", c.listAllPublications},
		option{"List Latest Posts", c.listLatestPosts},
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

	fmt.Print("Enter Feed URL: ")
	url, _ := c.reader.ReadString('\n')

	c.PublicationService.Add(entities.Publication{Title: title[:len(title)-1], URL: url[:len(url)-1]})
}

func (c *CLI) listAllPublications() {
	pubs := c.PublicationService.ListAll()
	if len(pubs) == 0 {
		fmt.Println("No Publications Yet")
	} else {
		fmt.Println("\nPublications:")
	}

	for _, pub := range pubs {
		fmt.Printf("%-20s %-48s %4d posts\n", pub.Title, fmt.Sprintf("%.46s", pub.URL), len(pub.Posts))
	}
}

func (c *CLI) listLatestPosts() {
	pubs := c.PublicationService.ListAll()
	posts := []entities.Post{}

	for _, pub := range pubs {
		for _, post := range pub.Posts {
			posts = append(posts, post)
		}
	}

	if len(posts) == 0 {
		fmt.Println("No Posts Yet")
		return
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})

	c.pagedList(posts, 0, 10)
}

func (c *CLI) pagedList(list []entities.Post, page int, size int) {
	start := page * size
	end := (page + 1) * size

	if start > len(list) || page < 0 || size < 1 {
		return
	} else if end > len(list) {
		end = len(list)
	}

	fmt.Printf("\nLatest Posts [%d - %d of %d]:\n", start, end, len(list))
	for _, post := range list[start:end] {
		fmt.Printf("%-60s %19s\n", fmt.Sprintf("%.58s", post.Title), post.CreatedAt.Format("2006-01-02 03:04PM"))
	}

	fmt.Print("\nEnter [p] for previous, [n] for next, or any other key to go back: ")
	userInput, _ := c.reader.ReadString('\n')
	userInput = userInput[:len(userInput)-1]

	if userInput == "n" {
		c.pagedList(list, page+1, size)
	} else if userInput == "p" {
		c.pagedList(list, page-1, size)
	}

	return
}

func (c *CLI) fetchAll() {
	pubs := c.PublicationService.ListAll()
	jobs := make(chan entities.Publication, 6)
	results := make(chan string, 6)

	for w := 1; w <= 3; w++ {
		go c.feedFetchWorker(w, jobs, results)
	}

	for _, pub := range pubs {
		jobs <- pub
	}
	close(jobs)

	for a := 1; a <= len(pubs); a++ {
		result := <-results
		fmt.Println(result)
	}
}

func (c *CLI) feedFetchWorker(id int, jobs <-chan entities.Publication, results chan<- string) {
	for pub := range jobs {
		err := c.PublicationService.FetchPublicationPosts(pub)
		if err != nil {
			results <- fmt.Sprintf("[Worker %d] Error fetching feed %q", id, pub.Title)
		} else {
			results <- fmt.Sprintf("[Worker %d] Fetched feed %q", id, pub.Title)
		}
	}
}
