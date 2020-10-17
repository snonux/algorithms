package set

import "fmt"

var (
	NotFound       = fmt.Errorf("Could not find entry")
	NotImplemented = fmt.Errorf("Method not implemented")
)

type Set interface {
	Empty() bool
	Set(key int, val int)
	Get(key int) (int, error)
	Del(key int) (int, error)
}
