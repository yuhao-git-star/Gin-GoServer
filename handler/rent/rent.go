package rentapi

import (
	"log"
	"net/http"

	"github.com/cheekybits/genny/generic"
	"github.com/gin-gonic/gin"

	rent "yasuoyuhao-591-api/services"
)

type DataContent generic.Type

type ResponseData struct {
	success bool
	code    int
	content []DataContent
}

// NewGenericQueue makes a new empty Generic queue.
func NewGenericQueue() *ResponseData {
	return &ResponseData{content: make([]DataContent, 0)}
}

// Searchrentby591 shows `OK` as the ping-pong result.
func Searchrentby591(c *gin.Context) {

	options := rent.NewOptions()
	url, err := rent.GenerateURL(options)
	if err != nil {
		log.Fatalf("\x1b[91;1m%s\x1b[0m", err)
	}

	f := rent.NewFiveN1(url)
	total := f.GetTotalPage()

	if err := f.Scrape(total); err != nil {
		log.Fatal(err)
	}

	data := NewGenericQueue()

	data.success = true
	data.code = 200
	data.content = append(data.content, f.RentList)

	c.JSON(http.StatusOK, gin.H{
		"success": data.success,
		"code":    data.code,
		"content": data.content,
	})
}
