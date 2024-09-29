package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// Poem represents a poem with its content
type Poem struct {
	Content string `json:"content"`
}

// Song represents a song with its title
type Song struct {
	Title string `json:"title"`
}

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Load HTML templates
	router.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML { return template.HTML(s) },
	})
	router.LoadHTMLGlob("templates/*")

	// Routes
	router.GET("/", showIndex)
	router.GET("/poem", getPoem)
	router.GET("/songs", getSongs)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

// showIndex renders the homepage
func showIndex(c *gin.Context) {
	// Read poem content
	poemContent, err := ioutil.ReadFile("poems/poem1.txt")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading poem")
		return
	}

	// Get random songs (assuming you have this function)
	songs := getRandomSongs()

	// Render the template with the data
	c.HTML(http.StatusOK, "index.html", gin.H{
		"poemContent": string(poemContent),
		"songs":       songs,
	})
}

// getPoem handles fetching a random poem with caching
func getPoem(c *gin.Context) {
	poems, err := filepath.Glob("poems/*.txt")
	if err != nil || len(poems) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No poems found."})
		return
	}

	// Randomly select a poem
	rand.Seed(time.Now().UnixNano())
	selectedPoem := poems[rand.Intn(len(poems))]

	// Read the poem content
	// content := // whatever was assigned here

	// For simplicity, returning the filename as content
	c.JSON(http.StatusOK, Poem{Content: selectedPoem})
}

// getSongs handles fetching a list of 5 random songs
func getSongs(c *gin.Context) {
	songs, err := filepath.Glob("songs/*.txt")
	if err != nil || len(songs) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No songs found."})
		return
	}

	// Randomly select 5 songs
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(songs), func(i, j int) {
		songs[i], songs[j] = songs[j], songs[i]
	})
	selectedSongs := songs
	if len(songs) > 5 {
		selectedSongs = songs[:5]
	}

	// Convert to Song structs
	var songList []Song
	for _, song := range selectedSongs {
		title := filepath.Base(song)
		title = title[:len(title)-len(filepath.Ext(song))]
		songList = append(songList, Song{Title: title})
	}

	c.JSON(http.StatusOK, songList)
}

// Assuming you have a function like this to get random songs
func getRandomSongs() []string {
	return []string{"song1", "song2", "song3"}
}
