package sort

import (
	"algorithms/ds"
)

func Quick(a ds.ArrayList) ds.ArrayList {
    a = Shuffle(a)
    quick(a, 0, len(a)-1)
    return a
}

func quick(a ds.ArrayList, lo, hi int) {
    if hi <= lo {
        return
    }

    j := quickPartition(a, lo, hi)
    quick(a, lo, j-1)
    quick(a, j+1, hi)
}

func quickPartition(a ds.ArrayList, lo, hi int) int {
    i := lo // Left scan index
    j := hi+1 // Right scan index
    v := a[lo] // Partitioning item

    for {
        // Scan left
        for {
            i++
            if i == hi {
                break
            }
            if a[i].Lower(v) {
                continue
            }
            break
        }

        // Scan right
        for {
            j--
            if j == lo {
                break
            }
            if v.Lower(a[j]) {
                continue
            }
            break
        }

        // Check scan complete
        if i >= j {
            break
        }

        a.Swap(i, j)
    }

    // Put partitioning item v into a[j]
    a.Swap(lo, j)

    // whith a[lo..j-1 <= a[j] <= a[j+1..hi]
    return j
}
