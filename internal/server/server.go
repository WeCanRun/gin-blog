package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterWarp struct {
	gr gin.IRouter
}

func NewRouter() *RouterWarp {
	return &RouterWarp{
		gr: &gin.RouterGroup{},
	}
}

func (w *RouterWarp) GR() gin.IRouter {
	return w.gr
}

func (w *RouterWarp) POST(path string, handler Handler) {
	w.gr.POST(path, HandlerWarp(handler))
}

func (w *RouterWarp) GET(path string, handler Handler) {
	w.gr.GET(path, HandlerWarp(handler))
}

func (w *RouterWarp) Use(handlers ...Handler) {
	w.gr.Use(HandlerWarp(handlers...))
}

func (w *RouterWarp) Handle(method, path string, handler Handler) {
	w.gr.Handle(method, path, HandlerWarp(handler))
}

func (w *RouterWarp) Any(path string, handler Handler) {
	w.gr.Any(path, HandlerWarp(handler))
}

func (w *RouterWarp) DELETE(path string, handler Handler) {
	w.gr.DELETE(path, HandlerWarp(handler))
}

func (w *RouterWarp) PATCH(path string, handler Handler) {
	w.gr.PATCH(path, HandlerWarp(handler))
}

func (w *RouterWarp) PUT(path string, handler Handler) {
	w.gr.PUT(path, HandlerWarp(handler))
}

func (w *RouterWarp) OPTIONS(path string, handler Handler) {
	w.gr.OPTIONS(path, HandlerWarp(handler))
}

func (w *RouterWarp) HEAD(path string, handler Handler) {
	w.gr.HEAD(path, HandlerWarp(handler))
}

func (w *RouterWarp) StaticFile(path, filePath string) {
	w.gr.StaticFile(path, filePath)
}

func (w *RouterWarp) Static(path, filePath string) {
	w.gr.Static(path, filePath)
}

func (w *RouterWarp) StaticFS(path string, fs http.FileSystem) {
	w.gr.StaticFS(path, fs)
}

func (w *RouterWarp) Group(path string, handlers ...Handler) *RouterWarp {
	return &RouterWarp{
		gr: w.gr.Group(path, HandlerWarp(handlers...)),
	}
}

type Handler func(*Context) error

func HandlerWarp(handler ...Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
