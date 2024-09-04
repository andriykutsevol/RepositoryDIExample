package router

import (
	"github.com/gin-gonic/gin"
	"mavianceutililtyservice/internal/presentation/http/handler"
)

type Router interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type router struct {
	administrativeregionHandler handler.AdminRegion
}

func NewRouter(
	administrativeregionHandler handler.AdminRegion,
) Router {
	return &router{
		administrativeregionHandler: administrativeregionHandler,
	}
}

func (r *router) Register(app *gin.Engine) error {
	r.RegisterAPI(app)
	return nil
}

func (r *router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
