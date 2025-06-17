package handlers

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"byfood-task/models"

	"github.com/gin-gonic/gin"
)

type URLHandler struct{}

func NewURLHandler() *URLHandler {
	return &URLHandler{}
}

func (h *URLHandler) ProcessURL(c *gin.Context) {
	var req models.ProcessURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("ERROR: Invalid input for ProcessURL: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedURL, err := url.Parse(req.URL)
	if err != nil {
		log.Printf("ERROR: Failed to parse URL '%s': %v", req.URL, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	switch req.Operation {
	case "canonical":
		applyCanonicalRules(parsedURL)
	case "redirection":
		applyRedirectionRules(parsedURL)
	case "all":
		applyCanonicalRules(parsedURL)
		applyRedirectionRules(parsedURL)
	}

	processedURLString := parsedURL.String()
	if req.Operation == "redirection" || req.Operation == "all" {
		processedURLString = strings.ToLower(processedURLString)
	}

	resp := models.ProcessURLResponse{
		ProcessedURL: processedURLString,
	}

	log.Printf("INFO: Processed URL. Original: '%s', Operation: '%s', Result: '%s'", req.URL, req.Operation, resp.ProcessedURL)
	c.JSON(http.StatusOK, resp)
}

func applyCanonicalRules(u *url.URL) {
	u.RawQuery = ""
	u.Path = strings.TrimSuffix(u.Path, "/")
}


func applyRedirectionRules(u *url.URL) {
	u.Host = "www.byfood.com"
}