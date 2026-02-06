package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"example.com/m/internal/services"
	"github.com/gin-gonic/gin"
)

type LinksHandler struct {
	service *services.LinksService
	host    string
	port    string
}

type CreateLinkResponse struct {
	Status       string `json:"status"`
	ShortLink    string `json:"short_link"`
	OriginalLink string `json:"original_link"`
	ExpiresIn    string `json:"expires_in"`
}

func NewLinksHandler(service *services.LinksService, host string, port string) *LinksHandler {
	return &LinksHandler{service: service, host: host, port: port}
}

func (h *LinksHandler) CreateLink(c *gin.Context) {
	link := c.Param("link")
	link = strings.TrimPrefix(link, "/")
	if c.Request.URL.RawQuery != "" {
		link += "?" + c.Request.URL.RawQuery
	}

	link, err := NormalizeURL(link)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed", "message": "invalid link"})
		return
	}

	id, err := h.service.CreateLink(link)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed", "message": "failed to create link"})
		return
	}

	address := fmt.Sprintf("http://%v:%s/r/%v", h.host, h.port, id)

	response := CreateLinkResponse{
		Status:       "success",
		ShortLink:    address,
		OriginalLink: link,
		ExpiresIn:    "30d",
	}

	c.JSON(http.StatusOK, response)
}

func (h *LinksHandler) Redirect(c *gin.Context) {
	id := c.Param("id")
	original, err := h.service.GetLink(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed", "message": "link not found"})
		return
	}

	c.Redirect(http.StatusFound, original)
}

func NormalizeURL(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", errors.New("empty URL")
	}

	if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
		raw = "https://" + raw
	}

	u, err := url.Parse(raw)
	if err != nil {
		return "", errors.New("invalid URL")
	}

	if u.Host == "" {
		return "", errors.New("invalid host in URL")
	}

	return u.String(), nil
}
