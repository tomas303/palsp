package discover

import (
	"fmt"
	"net/url"
	"runtime"
	"strings"
)

type pathElements struct {
	dir  string
	name string
	ext  string
}

func (p pathElements) Path() string {
	return p.dir + p.name + p.ext
}

func (p pathElements) Dir() string {
	return p.dir
}

func (p pathElements) Name() string {
	return p.name
}

func (p pathElements) Ext() string {
	return p.ext
}

func DecodePath(path string) pathElements {
	var normPath string
	if uri, err := url.Parse(path); err == nil && uri.Scheme == "file" {
		normPath = uri.Path
		// On Windows, remove leading slash if present
		if runtime.GOOS == "windows" && strings.HasPrefix(normPath, "/") {
			normPath = normPath[1:]
		}
	} else {
		normPath = path
	}

	// Split the path into directory, filename, and extension
	lastSlash := strings.LastIndex(normPath, "/")
	dir := ""
	fileWithExt := normPath
	if lastSlash >= 0 {
		dir = normPath[:lastSlash+1]
		fileWithExt = normPath[lastSlash+1:]
	}

	lastDot := strings.LastIndex(fileWithExt, ".")
	name := fileWithExt
	ext := ""
	if lastDot >= 0 {
		name = fileWithExt[:lastDot]
		ext = fileWithExt[lastDot:]
	}

	return pathElements{
		dir:  dir,
		name: name,
		ext:  ext,
	}
}

// SplitQualifiedName splits a dotted name (e.g., "a.b.c") into:
// - a slice of prefix parts (e.g., ["a", "b"])
// - the last part as a name (e.g., "c")
func SplitQualifiedName(qualifiedName string) ([]string, string) {
	// Split the name by dots
	parts := strings.Split(qualifiedName, ".")

	if len(parts) <= 1 {
		// No dots in the name, return empty prefix and the original name
		return []string{}, qualifiedName
	}

	// Get all parts except the last one as prefix
	prefix := parts[:len(parts)-1]
	// Get the last part as the name
	name := parts[len(parts)-1]

	return prefix, name
}

// stack is a generic stack that holds values of any type.
type stack[T any] struct {
	data []T
}

// Push adds an element to the top of the stack.
func (s *stack[T]) push(v T) {
	s.data = append(s.data, v)
}

// Pop removes and returns the top element of the stack; returns the zero value if empty.
func (s *stack[T]) pop() T {
	var zero T
	if len(s.data) == 0 {
		return zero
	}
	index := len(s.data) - 1
	elem := s.data[index]
	s.data = s.data[:index]
	return elem
}

// Peek returns the top element of the stack without removing it; returns the zero value if empty.
func (s *stack[T]) peek() T {
	var zero T
	if s.isEmpty() {
		return zero
	}
	return s.data[len(s.data)-1]
}

// Get returns the element at the specified index; returns the zero value if out of bounds.
func (s *stack[T]) get(index int) T {
	var zero T
	if index < 0 || index >= len(s.data) {
		return zero
	}
	return s.data[index]
}

// Enumerate returns a slice of all elements in the stack.
func (s *stack[T]) all() []T {
	return s.data
}

// Reverse returns a slice of all elements in the stack in reverse order.
func (s *stack[T]) reverse() []T {
	result := make([]T, len(s.data))
	for i := 0; i < len(s.data); i++ {
		result[i] = s.data[len(s.data)-1-i]
	}
	return result
}

func (s *stack[T]) joinByDot() string {
	parts := []string{}
	for _, v := range s.data {
		parts = append(parts, fmt.Sprint(v))
	}
	return strings.Join(parts, ".")
}

// IsEmpty returns true if the stack is empty.
func (s *stack[T]) isEmpty() bool {
	return len(s.data) == 0
}

// Length returns the number of elements in the stack.
func (s *stack[T]) length() int {
	return len(s.data)
}

// newStack creates and returns a new empty stack.
func newStack[T any]() *stack[T] {
	return &stack[T]{
		data: make([]T, 0),
	}
}
