package discover

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"palsp/internal/parser"
	"runtime"
	"strings"
	"sync"
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

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}

func getFileModTime(filepath string) (int64, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return fileInfo.ModTime().Unix(), nil
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

// KeyLock provides locking capabilities based on generic keys
// K must be comparable (usable as a map key)
type KeyLock[K comparable] struct {
	mu    sync.RWMutex
	locks map[K]*sync.Mutex
}

// NewKeyLock creates a new KeyLock instance
func NewKeyLock[K comparable]() *KeyLock[K] {
	return &KeyLock[K]{
		locks: make(map[K]*sync.Mutex),
	}
}

// Lock acquires a lock for the given key
// Multiple goroutines trying to Lock the same key will block until Unlock is called
// Different keys can be locked concurrently
func (k *KeyLock[K]) Lock(key K) {
	// First try to get the lock with a read lock (faster)
	k.mu.RLock()
	mutex, exists := k.locks[key]
	k.mu.RUnlock()

	if !exists {
		// Need to create a lock - requires write lock
		k.mu.Lock()
		// Check again in case another goroutine created it
		mutex, exists = k.locks[key]
		if !exists {
			mutex = &sync.Mutex{}
			k.locks[key] = mutex
		}
		k.mu.Unlock()
	}

	// Acquire the mutex for this key
	mutex.Lock()
}

// Unlock releases a lock for the given key
// Must be called after Lock with the same key
func (k *KeyLock[K]) Unlock(key K) {
	k.mu.RLock()
	mutex, exists := k.locks[key]
	k.mu.RUnlock()

	if exists {
		mutex.Unlock()
	}
}

// Example usage:
// func Process(unitName string) {
//     locker := NewKeyLock[string]()
//     locker.Lock(unitName)
//     defer locker.Unlock(unitName)
//
// }

// formatFileURI converts a filesystem path to a properly formatted URI
// Works for both Windows and Unix-like systems
func FormatFileURI(path string) string {
	// Normalize path separators to forward slashes (for Windows)
	normalizedPath := strings.ReplaceAll(path, "\\", "/")

	if runtime.GOOS == "windows" {
		// Check if path has a drive letter (e.g. C:)
		if len(normalizedPath) >= 2 && normalizedPath[1] == ':' {
			// Windows path with drive letter needs three slashes: file:///C:/path
			return "file:///" + normalizedPath
		}

		// Windows path without drive letter (e.g. network path)
		// Should start with file:// and not have a leading slash
		if strings.HasPrefix(normalizedPath, "/") {
			return "file://" + normalizedPath
		}
		return "file:///" + normalizedPath
	} else {
		// Unix paths always start with /
		if !strings.HasPrefix(normalizedPath, "/") {
			normalizedPath = "/" + normalizedPath
		}
		return "file://" + normalizedPath
	}
}

// todo : later make dictionary of symbolic names on startup
func findParserSymbolicNameID(target string) (int, bool) {
	for i, str := range parser.PascalLexerLexerStaticData.SymbolicNames {
		if str == target {
			return i, true
		}
	}
	return -1, false
}
