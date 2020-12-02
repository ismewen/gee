package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("route %4s-%s", method, handler)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewwares []HandlerFunc
	for _, group := range engine.groups {
		print("serverhttp****")
		log.Println(req.URL.Path, group.prefix)
		log.Println(strings.HasPrefix(req.URL.Path, group.prefix))
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			log.Println("add middlewwares", group, group.middlewares)
			middlewwares = append(middlewwares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewwares
	engine.router.handle(c)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ShowParts() {
	print("**********")
	for _, root := range engine.router.roots {
		root.showAllPath()
	}
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
