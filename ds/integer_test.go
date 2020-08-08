package ds

import "testing"

func TestCompare(t *testing.T) {
	i := Integer{1}
	j := Integer{10}
	res := i.Compare(j)
	if res != -1 {
		t.Errorf("%v must be lower than %v, but got %v", i, j, res)
	}

	res = j.Compare(i)
	if res != 1 {
		t.Errorf("%v must be higher than %v, but got %v", j, i, res)
	}

	i = Integer{2}
	j = Integer{2}
	res = i.Compare(j)
	if res != 0 {
		t.Errorf("%v must be equal to %v, but got %v", i, j, res)
	}

	i = Integer{23}
	j = Integer{23}
	if !i.HigherEqual(j) {
		t.Errorf("Unpexpected %v.HigherEqual(%v) == false",
		    i,j)
	}

	i = Integer{23}
	j = Integer{42}
	if i.HigherEqual(j) {
		t.Errorf("Unpexpected %v.HigherEqual(%v) == true",
		    i,j)
	}
	if !j.HigherEqual(i) {
		t.Errorf("Unpexpected %v.HigherEqual(%v) == false",
		    j,i)
	}
}
