package router

import (
	"github.com/gin-gonic/gin"
	"text/template"
)

func Router(runChan chan<- struct{}, terminalChan chan<- struct{}) *gin.Engine {
	g := gin.New()
	g.GET("/", func(c *gin.Context) {
		t1, err := template.ParseFiles("./router/index.html")
		if err != nil {
			panic(err)
		}
		t1.Execute(c.Writer, nil)
	})

	g.GET("/run", func(c *gin.Context) {
		runChan<- struct{}{}
		t1, err := template.ParseFiles("./router/index.html")
		if err != nil {
			panic(err)
		}
		t1.Execute(c.Writer, nil)
	})

	g.GET("/terminal", func(c *gin.Context) {
		close(runChan)
		terminalChan<- struct{}{}
		t1, err := template.ParseFiles("./router/index.html")
		if err != nil {
			panic(err)
		}
		t1.Execute(c.Writer, nil)
	})

	return g
}
