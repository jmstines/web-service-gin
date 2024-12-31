package controllers

import "github.com/gin-gonic/gin"

type GinRouter struct {
	engine *gin.Engine
}

func NewGinRouter(engine *gin.Engine) *GinRouter {
	return &GinRouter{engine: engine}
}

func (r *GinRouter) GET(relativePath string, handlers ...gin.HandlerFunc) {
	r.engine.GET(relativePath, handlers...)
}

func (r *GinRouter) POST(relativePath string, handlers ...gin.HandlerFunc) {
	r.engine.POST(relativePath, handlers...)
}

func (r *GinRouter) PATCH(relativePath string, handlers ...gin.HandlerFunc) {
	r.engine.POST(relativePath, handlers...)
}

func (r *GinRouter) DELETE(relativePath string, handlers ...gin.HandlerFunc) {
	r.engine.POST(relativePath, handlers...)
}
