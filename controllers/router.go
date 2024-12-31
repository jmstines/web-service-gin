package controllers

import "github.com/gin-gonic/gin"

type Router interface {
	GET(relativePath string, handlers ...gin.HandlerFunc)
	POST(relativePath string, handlers ...gin.HandlerFunc)
	PATCH(relativePath string, handlers ...gin.HandlerFunc)
	DELETE(relativePath string, handlers ...gin.HandlerFunc)
}
