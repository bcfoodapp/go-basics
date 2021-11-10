package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	api := NewAPI()
	router.GET("/", root)
	router.GET("/vendors/", api.vendors)
	router.GET("/vendors/:id", api.vendor)

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

type API struct {
	Vendors []Vendor
}

func NewAPI() *API {
	vendors := make([]Vendor, 1)
	vendors[0] = Vendor{
		ID:   "cycledogs",
		Name: "Cycle Dogs",
	}
	return &API{vendors}
}

func (b *API) vendors(c *gin.Context) {
	c.JSON(http.StatusOK, &b.Vendors)
}

func (b *API) vendor(c *gin.Context) {
	id := c.Param("id")
	for _, vendor := range b.Vendors {
		if vendor.ID == id {
			c.JSON(http.StatusOK, &vendor)
			return
		}
	}
	c.Status(http.StatusNotFound)
}
