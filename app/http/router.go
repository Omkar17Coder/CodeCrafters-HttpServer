package http

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/types"
)

type TrieNode struct {
	chilren  map[string]*TrieNode
	handlers map[string]types.HandlerFunc // Store handler for Different methods.
}

type Router struct {
	root *TrieNode
}

func NewRouter() *Router {
	return &Router{
		root: &TrieNode{chilren: make(map[string]*TrieNode), handlers: make(map[string]types.HandlerFunc)},
	}
}

func (r *Router) Register(method, path string, handlerFunction types.HandlerFunc) {
	node := r.root
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" {
			continue // we can skip
		}
		if _, exits := node.chilren[part]; !exits {
			node.chilren[part] = &TrieNode{chilren: make(map[string]*TrieNode), handlers: make(map[string]types.HandlerFunc)}
		}
		node = node.chilren[part]
	}
	node.handlers[method] = handlerFunction

}

func (r *Router) Route(method, path string) (types.HandlerFunc, bool) {
	node := r.root
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" {
			continue // /api//hello -> /api/hello
		}
		if child, exist := node.chilren[part]; exist {
			node = child
		} else if child, exist := node.chilren["*"]; exist {
			node = child
		} else {
			return nil, false
		}
	}
	handler, exists := node.handlers[method]
	if !exists {
		fmt.Println("I dont exits")
		return nil, false
	}
	return handler, true
}
