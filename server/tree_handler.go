package server

import "strings"

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	handler handleFunc
}

func (n *HandlerBasedOnTree) ServerHTTP(c *Context) {
	panic("123123123")
}

func (n *HandlerBasedOnTree) Route(method string, pattern string, handleFunc handleFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := n.root

	for index, path := range paths {
		matchChild, ok := n.findMatchChild(cur, path)

		if ok {
			cur = matchChild
		} else {
			n.createSubTree(cur, paths[index:], handleFunc)
			return
		}
	}
}

func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handleFunc handleFunc) {

	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}

	cur.handler = handleFunc
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 4),
	}
}

func (h *HandlerBasedOnTree) findMatchChild(root *node, path string) (*node, bool) {
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}
