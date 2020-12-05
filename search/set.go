package set

import "fmt"

var (
	NotFound       = fmt.Errorf("could not find entry")
	NotImplemented = fmt.Errorf("method not implemented")
)

type Set interface {
	Empty() bool
	Set(key int, val int)
	Get(key int) (int, error)
	Del(key int) (int, error)
}
