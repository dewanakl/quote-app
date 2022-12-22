package routes

import (
	"quoteapp/controller"
	"quoteapp/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Route struct {
	g     *gin.Engine
	quote *controller.Quote
}

func NewRoute(db *gorm.DB, gi *gin.Engine) *gin.Engine {
	route := &Route{
		g: gi,
	}

	route.quote = controller.NewQuoteController(model.NewQuoteModel(db))

	route.routes()
	return gi
}

func (r *Route) routes() {
	r.g.GET("/fetch", r.quote.Fetch)
	r.g.GET("/select", r.quote.Show)
	r.g.GET("/count", r.quote.Count)
	r.g.POST("/add", r.quote.Store)
}
