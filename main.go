package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Middleware V2 Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.GET("/index", func(o *gee.Context) {
		o.String(http.StatusOK, "index")
	})

	v1 := r.Group("/v1")
	v1.Use(gee.Logger())
	{
		v1.GET("/hello", func(o *gee.Context) {
			o.String(http.StatusOK, "v1 hello")
		})

		v1.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello"+c.Params["name"])
		})
	}

	v1.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "v1 index")
	})

	v2 := r.Group("/v2")
	{
		v2.GET("/", func(c *gee.Context) {
			c.String(http.StatusOK, "")
		})
	}
	v2.Use(onlyForV2())
	r.Run(":9999")
}
