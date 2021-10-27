package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	backend := NewBackend()
	router.GET("/", root)
	router.GET("/vendors/", backend.vendors)
	router.GET("/vendors/:id", backend.vendor)

	server := http.Server{
		Addr:         ":8000",
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func root(c *gin.Context) {
	c.String(http.StatusOK, "You've called the root route!")
}

type Vendor struct {
	ID   string
	Name string
}

type Backend struct {
	Vendors []Vendor
}

func NewBackend() *Backend {
	vendors := make([]Vendor, 1)
	vendors[0] = Vendor{
		ID:   "cycledogs",
		Name: "Cycle Dogs",
	}
	return &Backend{vendors}
}

func (b *Backend) vendors(c *gin.Context) {
	c.JSON(http.StatusOK, &b.Vendors)
}

func (b *Backend) vendor(c *gin.Context) {
	id := c.Param("id")
	for _, vendor := range b.Vendors {
		if vendor.ID == id {
			c.JSON(http.StatusOK, &vendor)
			return
		}
	}
	c.Status(http.StatusNotFound)
}
