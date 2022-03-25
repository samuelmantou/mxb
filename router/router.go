package router

import (
	"github.com/gin-gonic/gin"
	"text/template"
)

type Console struct {
	done chan struct{}
}

func NewConsole() *Console {
	c := Console{
		done: make(chan struct{}),
	}

	return &c
}

func (receiver *Console) lazyInit()  {
	if receiver.done == nil {
		receiver.done = make(chan struct{})
	}
}

func (receiver *Console) Done() <-chan struct{} {
	receiver.lazyInit()
	return receiver.done
}

func (receiver *Console) Handler() *gin.Engine {
	g := gin.New()
	g.GET("/", func(c *gin.Context) {
		t1, err := template.ParseFiles("./router/index.html")
		if err != nil {
			panic(err)
		}
		t1.Execute(c.Writer, nil)
	})

	g.GET("/run", func(c *gin.Context) {
		t1, err := template.ParseFiles("./router/index.html")
		if err != nil {
			panic(err)
		}
		t1.Execute(c.Writer, nil)
	})

	g.GET("/terminal", func(c *gin.Context) {
		receiver.lazyInit()
		receiver.done<- struct{}{}
		t1, err := template.ParseFiles("./router/index.html")
		if err != nil {
			panic(err)
		}
		t1.Execute(c.Writer, nil)
	})

	return g
}
