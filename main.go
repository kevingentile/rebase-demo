package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevingentile/rebase-demo/animal"
)

func main() {
	fmt.Println("Hello World!")

	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	engine.GET("/human/:name", func(c *gin.Context) {
		name := c.Param("name")
		h := animal.NewHuman(name)

		c.JSON(http.StatusOK, walkerPayloadGenerator("human", h))
	})

	engine.GET("/shark/:species", func(c *gin.Context) {
		species := c.Param("species")
		s := animal.NewShark(species)
		c.JSON(http.StatusOK, swimmerPayloadGenerator("shark", s))
	})

	engine.Run()
}

func walkerPayloadGenerator(walkerType string, walker animal.Walker) gin.H {
	return gin.H{
		"description": animal.DescribeWalker(walker),
		walkerType:    walker,
	}
}

func swimmerPayloadGenerator(swimmerType string, swimmer animal.Swimmer) gin.H {
	return gin.H{
		"description": animal.DescribeSwimmer(swimmer),
		swimmerType:   swimmer,
	}
}
