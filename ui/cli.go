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
		option{"Search Posts And Publications", c.searchPosts},
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

// addPublication allows user to add a new publication
func (c *CLI) addPublication() {
	fmt.Print("Enter Title: ")
	title, _ := c.reader.ReadString('\n')

	fmt.Print("Enter Feed URL: ")
	url, _ := c.reader.ReadString('\n')

	c.PublicationService.Add(entities.Publication{Title: title[:len(title)-1], URL: url[:len(url)-1]})
}

// listAllPublications lists all publications
func (c *CLI) listAllPublications() {
	pubs := c.PublicationService.ListAll()
	if len(pubs) == 0 {
		fmt.Println("No Publications Yet")
	} else {
		fmt.Println("\nPublications:")
	}

	for _, pub := range pubs {
		fmt.Println(&pub)
	}
}

// listLatestPosts lists the posts ordered by time (latest first)
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

	c.pagedList("Latest Posts", posts, 0, 10)
}

// searchPosts allows users to search for posts
func (c *CLI) searchPosts() {
	fmt.Print("\nEnter search query: ")
	userInput, _ := c.reader.ReadString('\n')
	userInput = userInput[:len(userInput)-1]

	if userInput == "" {
		return
	}

	if results := c.PublicationService.Search("Publication", userInput); len(results) > 0 {
		fmt.Println("")
		fmt.Println("Publications:")
		for _, res := range results {
			if pub, err := c.PublicationService.Find(res.ID); err == nil {
				fmt.Println(&pub)
			}
		}
	}

	if results := c.PublicationService.Search("Post", userInput); len(results) > 0 {
		ids := make([]string, 0)
		for _, res := range results {
			ids = append(ids, res.ID)
		}
		posts := c.PublicationService.GetPosts(ids)
		c.pagedList("Posts", posts, 0, 10)
	} else {
		fmt.Println("")
		fmt.Println("No Posts Found")
	}
}

// pagedList is a method to break up posts into multuple pages
func (c *CLI) pagedList(title string, list []entities.Post, page int, size int) {
	start := page * size
	end := (page + 1) * size

	if start > len(list) || page < 0 || size < 1 {
		return
	} else if end > len(list) {
		end = len(list)
	}

	fmt.Printf("\n%s [%d - %d of %d]:\n", title, start, end, len(list))
	for index, post := range list[start:end] {
		fmt.Printf("[%d] %s\n", index, &post)
	}

	fmt.Print("\nEnter [0-9] for details, [p] for previous, [n] for next, or any other key to go back: ")
	userInput, _ := c.reader.ReadString('\n')
	userInput = userInput[:len(userInput)-1]

	if input, err := strconv.Atoi(userInput); err == nil && input >= 0 && input+start < len(list) {
		c.viewPost(list[start+input])
		c.pagedList(title, list, page, size)
	} else if userInput == "n" {
		c.pagedList(title, list, page+1, size)
	} else if userInput == "p" {
		c.pagedList(title, list, page-1, size)
	}

	return
}

// viewPost allows users to view more details about a post
func (c *CLI) viewPost(post entities.Post) {
	fmt.Println("")
	fmt.Println(post.Title)
	fmt.Println(post.URL)
	fmt.Printf("Posted at %s", post.CreatedAt.Format("2006-01-02 03:04PM"))
	fmt.Println("")

	if results := c.PublicationService.Search("Post", post.Title); len(results) > 1 {
		fmt.Println("")
		fmt.Println("Similar Posts:")
		for _, res := range results[1:] {
			fmt.Println(&res)
		}
		fmt.Println("")
	}
}

// fetchAll fetches all the feeds from publications
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

// feedFetchWorker is the concurrency mechanism to allow multiple feed fetcher workers
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
