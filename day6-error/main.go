package main

import (
	"fmt"
	"goo"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := goo.Default()
	r.GET("/", func(c *goo.Context) {
		c.String(http.StatusOK, "Hello Kent\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *goo.Context) {
		names := []string{"kent"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
