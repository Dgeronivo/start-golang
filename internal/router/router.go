package router

import (
	"io"
	"log"
	questlisting "website/internal/quest-listing"

	"github.com/gin-gonic/gin"
)

func CreateNew() *gin.Engine {
	router := gin.Default()

	// router.GET("/static/:file", func(ctx *gin.Context) {

	// 	log.Println(ctx.Param("file"))
	// 	ctx.Writer.Write([]byte("hello"))
	// })
	router.Static("assets", "./web/assets")

	router.GET("/test", func(ctx *gin.Context) {
		// time.Sleep(3 * time.Second)
		ctx.IndentedJSON(200, map[string]string{"test": "test"})
	})

	router.GET("/quests", renderProcessor(questlisting.RenderListing))

	return router
}

func renderProcessor(responseContentWriteFunc func(writer io.Writer) error) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Writer.WriteHeader(200)
		err := responseContentWriteFunc(ctx.Writer)
		if err != nil {
			log.Fatalln(err)
			ctx.AbortWithStatus(500)
			return
		}
	}
}
