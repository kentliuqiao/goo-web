package main

import (
	"fmt"
	"goo"
	"html/template"
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
	r := goo.New()
	r.Use(goo.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Kent", Age: 18}
	stu2 := &student{Name: "Jack", Age: 20}
	r.GET("/", func(c *goo.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *goo.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", goo.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *goo.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", goo.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
