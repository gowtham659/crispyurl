package handler

import (
	"fmt"
	"log"
	"math/rand"
	"urlshortnerreactgo/model"

	"github.com/gofiber/fiber/v2"
)

type UrlDetails struct {
	OriginalUrl string `json:"original_url"`
	ShortenUrl  string `json:"shorten_url,omitempty"`
	ShortCode   string `json:"short_code,omitempty"`
	Message     string `json:"message"`
}

type RequestUrl struct {
	OrgUrl 		string `json:"OrgUrl"`
	BaseUrl		string `json:"BaseUrl"`
}

// Generate a random short code
func GenerateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 5
	shortCode := make([]byte, length)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortCode)
}

func ShortenUrl(c *fiber.Ctx) error {
	var urldetails UrlDetails
	var actualUrl RequestUrl
	fmt.Println("raw",string(c.Body()))
	//parsing request body by binding into struct
	if er := c.BodyParser(&actualUrl); er != nil {
		log.Println("parsing error: ", er)
	}

	fmt.Println("request url: ", actualUrl) 

	//returns shortcode
	scode := GenerateShortCode()

	//inserting original url, shortcode into db
	status := model.InsertUrl(actualUrl.OrgUrl, actualUrl.BaseUrl, scode)
	if status {
		shortUrl := actualUrl.BaseUrl + scode

		urldetails.OriginalUrl = actualUrl.OrgUrl
		urldetails.Message = "success"
		urldetails.ShortenUrl = shortUrl
		urldetails.ShortCode = scode

		fmt.Println(urldetails)
		return c.JSON(urldetails)
	}
	urldetails.Message = "failed"
	fmt.Println(urldetails)
	return c.JSON(urldetails)
}

// RedirectURL handles redirecting short codes to original URLs
func RedirectURL(c *fiber.Ctx) error {
	urldetails := UrlDetails{}

	//reading path param value
	shortCode := c.Params("shortcode")
	// Retrieve the original URL from the database
	originalURL, err := model.GetOriginalURL(shortCode)
	if err == nil && len(originalURL) > 0 {
		c.Redirect(originalURL)
	}
	urldetails.Message = "not found"
	return c.JSON(urldetails)
}
