package server

import (
	"errors"
	"net/http"
	"strings"
)

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	handler handleFunc
}

func (n *HandlerBasedOnTree) ServerHTTP(c *Context) {
	handler, found := n.findRouter(c.R.URL.Path)

	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("not found"))
	}

	handler(c)
}

func (n *HandlerBasedOnTree) findRouter(path string) (handleFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")

	cur := n.root

	for _, p := range paths {
		matchChild, found := n.findMatchChild(cur, p)

		if !found {
			return nil, false
		}
		cur = matchChild
	}

	if cur.handler == nil {
		return nil, false
	}

	return cur.handler, true
}

func (n *HandlerBasedOnTree) Route(method string, pattern string, handleFunc handleFunc)  error {

	err := n.validatePattern(pattern)

	if err != nil{
		return err
	}

	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := n.root

	for index, path := range paths {
		matchChild, ok := n.findMatchChild(cur, path)

		if ok {
			cur = matchChild
		} else {
			n.createSubTree(cur, paths[index:], handleFunc)
			return nil
		}
	}
	return nil
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

	var wildcardNode *node
	for _, child := range root.children {
		if child.path == path && child.path != "*" {
			return child, true
		}

		if child.path == "*" {
			wildcardNode = child
		}

	}
	return wildcardNode, wildcardNode != nil
}

func (n *HandlerBasedOnTree) validatePattern(pattern string) error {
	pos := strings.Index(pattern, "*")

	if pos > 0 {
		if pos != len(pattern)-1 {
			return errors.New("错误的路径")
		}

		if pattern[pos-1] != '/' {
			return errors.New("错误的路径")
		}
	}

	return nil
}
