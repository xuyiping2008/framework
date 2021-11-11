package ping

import (
	"strings"
)

type node struct {
	Pattern  string  `json:"pattern"`
	Part     string  `json:"part"`
	Children []*node `json:"children"`
	IsWild   bool    `json:"is_wild"`
}

func (n *node) mathChild(part string) *node {
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)

	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// /
// &{pattern:/ part: children:[] isWild:false}
// pattern = '/hello'  parts = [hello]
// &{pattern:/ part: children:[&node{pattern: /hello part: part,children [] isWild:false] isWild:false}
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}

	part := parts[height]
	child := n.mathChild(part)
	if child == nil {
		child = &node{Part: part, IsWild: part[0] == ':' || part[0] == '*'}
		n.Children = append(n.Children, child)
	}

	child.insert(pattern, parts, height+1)
}

// [] 0
// [hello] 0
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
