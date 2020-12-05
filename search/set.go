package search

import "fmt"

var (
	NotFound       = fmt.Errorf("could not find entry")
	NotImplemented = fmt.Errorf("method not implemented")
)

type Put interface {
	Empty() bool
	Put(key int, val int)
	Get(key int) (int, error)
	Del(key int) (int, error)
}
