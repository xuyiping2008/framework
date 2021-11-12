package ping

import (
	"net/http"
	"strings"
)

type router struct {
	Roots    map[string]*node       `json:"roots"`
	Handlers map[string]HandlerFunc `json:"handlers"`
}

func NewRouter() *router {
	return &router{
		Handlers: make(map[string]HandlerFunc),
		Roots:    make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	if _, ok := r.Roots[method]; !ok {
		r.Roots[method] = &node{}
	}

	r.Roots[method].insert(pattern, parts, 0)
	if r.Roots[method] != nil {
		PrintChildren(r.Roots[method].Children)
	}
	r.Handlers[key] = handler
}

func PrintChildren(n []*node) {
	if n == nil {
		return
	}
	for _, child := range n {
		PrintChildren(child.Children)
	}
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {

	root, ok := r.Roots[method]
	if !ok {
		return nil, nil
	}
	parts := parsePattern(path)
	params := make(map[string]string)

	n := root.search(parts, 0)
	if n != nil {
		newParts := parsePattern(n.Pattern)
		for index, part := range newParts {
			if part[0] == ':' {
				params[part[1:]] = parts[index]
			}

			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(parts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.Pattern
		c.Handlers = append(c.Handlers, r.Handlers[key])
	} else {
		c.Handlers = append(c.Handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
		})
	}
}
