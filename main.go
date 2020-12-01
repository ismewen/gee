package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/index", func(o *gee.Context) {
		o.String(http.StatusOK, "index")
	})

	v1 := r.Group("/v1")
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
	r.Run(":9999")
}
