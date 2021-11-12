package ping

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const banner = `
 _ __ (_)_ __   __ _ 
| '_ \| | '_ \ / _" |
| |_) | | | | | (_| |
| .__/|_|_| |_|\__, |
|_|            |___/ v%s
High performance, minimalist Go web framework
`

const (
	version = "1.0.0"
)

type HandlerFunc func(c *Context)

type Engine struct {
	Router        *router `json:"router"`
	*RouterGroup  `json:"router_group"`
	Groups        []*RouterGroup `json:"groups"`
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

type RouterGroup struct {
	Prefix      string        `json:"prefix"`
	Middlewares []HandlerFunc `json:"middlewares"`
	Engine      *Engine       `json:"engine"`
}

func Default() *Engine {
	e := &Engine{
		Router: NewRouter(),
	}
	e.RouterGroup = &RouterGroup{Engine: e}
	e.Groups = []*RouterGroup{e.RouterGroup}
	e.Use(Logger(), Recovery())
	return e
}

func (e *Engine) Run(addr string) error {
	fmt.Fprintf(os.Stdout, banner, version)
	return http.ListenAndServe(addr, e)
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHtmlGlob(pattern string) {
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	e := group.Engine
	newGroup := &RouterGroup{
		Prefix: prefix,
		Engine: e,
	}
	e.Groups = append(e.Groups, newGroup)
	return newGroup
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.Prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (group *RouterGroup) Static(relativePath, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPath := path.Join(relativePath, "/*filepath")
	group.GET(urlPath, handler)
}

func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Middlewares = append(group.Middlewares, middleware...)
}

func (group *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	patter := group.Prefix + comp
	log.Printf("Route %4s - %s", method, patter)
	group.Engine.Router.addRoute(method, patter, handler)
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.Router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.Groups {
		if strings.HasPrefix(req.URL.Path, group.Prefix) {
			middlewares = append(middlewares, group.Middlewares...)
		}
	}
	c := NewContext(w, req)
	c.Handlers = middlewares
	c.engine = e
	e.Router.handle(c)
	c.Next()
}
