/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 - 2022 Ralph Caraveo (deckarep@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package mapset

import (
	"testing"
)

func makeSetInt(ints []int) Set[int] {
	s := NewSet[int]()
	for _, i := range ints {
		s.Add(i)
	}
	return s
}

func makeUnsafeSetInt(ints []int) Set[int] {
	s := NewThreadUnsafeSet[int]()
	for _, i := range ints {
		s.Add(i)
	}
	return s
}

func makeSetIntWithAppend(ints ...int) Set[int] {
	s := NewSet[int]()
	s.Append(ints...)
	return s
}

func makeUnsafeSetIntWithAppend(ints ...int) Set[int] {
	s := NewThreadUnsafeSet[int]()
	s.Append(ints...)
	return s
}

func assertEqual[T comparable](a, b Set[T], t *testing.T) {
	if !a.Equal(b) {
		t.Errorf("%v != %v\n", a, b)
	}
}

func Test_NewSet(t *testing.T) {
	a := NewSet[int]()
	if a.Cardinality() != 0 {
		t.Error("NewSet should start out as an empty set")
	}

	assertEqual(NewSet([]int{}...), NewSet[int](), t)
	assertEqual(NewSet([]int{1}...), NewSet(1), t)
	assertEqual(NewSet([]int{1, 2}...), NewSet(1, 2), t)
	assertEqual(NewSet([]string{"a"}...), NewSet("a"), t)
	assertEqual(NewSet([]string{"a", "b"}...), NewSet("a", "b"), t)
}

func Test_NewUnsafeSet(t *testing.T) {
	a := NewThreadUnsafeSet[int]()

	if a.Cardinality() != 0 {
		t.Error("NewSet should start out as an empty set")
	}
}

func Test_AddSet(t *testing.T) {
	a := makeSetInt([]int{1, 2, 3})

	if a.Cardinality() != 3 {
		t.Error("AddSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func Test_AddUnsafeSet(t *testing.T) {
	a := makeUnsafeSetInt([]int{1, 2, 3})

	if a.Cardinality() != 3 {
		t.Error("AddSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func Test_AppendSet(t *testing.T) {
	a := makeSetIntWithAppend(1, 2, 3)

	if a.Cardinality() != 3 {
		t.Error("AppendSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func Test_AppendUnsafeSet(t *testing.T) {
	a := makeUnsafeSetIntWithAppend(1, 2, 3)

	if a.Cardinality() != 3 {
		t.Error("AppendSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func Test_AddSetNoDuplicate(t *testing.T) {
	a := makeSetInt([]int{7, 5, 3, 7})

	if a.Cardinality() != 3 {
		t.Error("AddSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AddSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func Test_AddUnsafeSetNoDuplicate(t *testing.T) {
	a := makeUnsafeSetInt([]int{7, 5, 3, 7})

	if a.Cardinality() != 3 {
		t.Error("AddSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AddSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func Test_AppendSetNoDuplicate(t *testing.T) {
	a := makeSetIntWithAppend(7, 5, 3, 7)

	if a.Cardinality() != 3 {
		t.Error("AppendSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AppendSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func Test_AppendUnsafeSetNoDuplicate(t *testing.T) {
	a := makeUnsafeSetIntWithAppend(7, 5, 3, 7)

	if a.Cardinality() != 3 {
		t.Error("AppendSetNoDuplicate set should have 3 elements since 7 is a duplicate")
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("AppendSetNoDuplicate set should have a 7, 5, and 3 in it.")
	}
}

func Test_RemoveSet(t *testing.T) {
	a := makeSetInt([]int{6, 3, 1})

	a.Remove(3)

	if a.Cardinality() != 2 {
		t.Error("RemoveSet should only have 2 items in the set")
	}

	if !(a.Contains(6) && a.Contains(1)) {
		t.Error("RemoveSet should have only items 6 and 1 in the set")
	}

	a.Remove(6)
	a.Remove(1)

	if a.Cardinality() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func Test_RemoveAllSet(t *testing.T) {
	a := makeSetInt([]int{6, 3, 1, 8, 9})

	a.RemoveAll(3, 1)

	if a.Cardinality() != 3 {
		t.Error("RemoveAll should only have 2 items in the set")
	}

	if !a.Contains(6, 8, 9) {
		t.Error("RemoveAll should have only items (6,8,9) in the set")
	}

	a.RemoveAll(6, 8, 9)

	if a.Cardinality() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func Test_RemoveUnsafeSet(t *testing.T) {
	a := makeUnsafeSetInt([]int{6, 3, 1})

	a.Remove(3)

	if a.Cardinality() != 2 {
		t.Error("RemoveSet should only have 2 items in the set")
	}

	if !(a.Contains(6) && a.Contains(1)) {
		t.Error("RemoveSet should have only items 6 and 1 in the set")
	}

	a.Remove(6)
	a.Remove(1)

	if a.Cardinality() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func Test_RemoveAllUnsafeSet(t *testing.T) {
	a := makeUnsafeSetInt([]int{6, 3, 1, 8, 9})

	a.RemoveAll(3, 1)

	if a.Cardinality() != 3 {
		t.Error("RemoveAll should only have 2 items in the set")
	}

	if !a.Contains(6, 8, 9) {
		t.Error("RemoveAll should have only items (6,8,9) in the set")
	}

	a.RemoveAll(6, 8, 9)

	if a.Cardinality() != 0 {
		t.Error("RemoveSet should be an empty set after removing 6 and 1")
	}
}

func Test_ContainsSet(t *testing.T) {
	a := NewSet[int]()

	a.Add(71)

	if !a.Contains(71) {
		t.Error("ContainsSet should contain 71")
	}

	a.Remove(71)

	if a.Contains(71) {
		t.Error("ContainsSet should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("ContainsSet should contain 13, 7, 1")
	}
}

func Test_ContainsUnsafeSet(t *testing.T) {
	a := NewThreadUnsafeSet[int]()

	a.Add(71)

	if !a.Contains(71) {
		t.Error("ContainsSet should contain 71")
	}

	a.Remove(71)

	if a.Contains(71) {
		t.Error("ContainsSet should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("ContainsSet should contain 13, 7, 1")
	}
}

func Test_ContainsMultipleSet(t *testing.T) {
	a := makeSetInt([]int{8, 6, 7, 5, 3, 0, 9})

	if !a.Contains(8, 6, 7, 5, 3, 0, 9) {
		t.Error("ContainsAll should contain Jenny's phone number")
	}

	if a.Contains(8, 6, 11, 5, 3, 0, 9) {
		t.Error("ContainsAll should not have all of these numbers")
	}
}

func Test_ContainsMultipleUnsafeSet(t *testing.T) {
	a := makeUnsafeSetInt([]int{8, 6, 7, 5, 3, 0, 9})

	if !a.Contains(8, 6, 7, 5, 3, 0, 9) {
		t.Error("ContainsAll should contain Jenny's phone number")
	}

	if a.Contains(8, 6, 11, 5, 3, 0, 9) {
		t.Error("ContainsAll should not have all of these numbers")
	}
}

func Test_ContainsOneSet(t *testing.T) {
	a := NewSet[int]()

	a.Add(71)

	if !a.ContainsOne(71) {
		t.Error("ContainsSet should contain 71")
	}

	a.Remove(71)

	if a.ContainsOne(71) {
		t.Error("ContainsSet should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.ContainsOne(13) && a.ContainsOne(7) && a.ContainsOne(1)) {
		t.Error("ContainsSet should contain 13, 7, 1")
	}
}

func Test_ContainsOneUnsafeSet(t *testing.T) {
	a := NewThreadUnsafeSet[int]()

	a.Add(71)

	if !a.ContainsOne(71) {
		t.Error("ContainsSet should contain 71")
	}

	a.Remove(71)

	if a.ContainsOne(71) {
		t.Error("ContainsSet should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.ContainsOne(13) && a.ContainsOne(7) && a.ContainsOne(1)) {
		t.Error("ContainsSet should contain 13, 7, 1")
	}
}

func Test_ContainsAnySet(t *testing.T) {
	a := NewSet[int]()

	a.Add(71)

	if !a.ContainsAny(71) {
		t.Error("ContainsSet should contain 71")
	}

	if !a.ContainsAny(71, 10) {
		t.Error("ContainsSet should contain 71 or 10")
	}

	a.Remove(71)

	if a.ContainsAny(71) {
		t.Error("ContainsSet should not contain 71")
	}

	if a.ContainsAny(71, 10) {
		t.Error("ContainsSet should not contain 71 or 10")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.ContainsAny(13, 17, 10)) {
		t.Error("ContainsSet should contain 13, 17, or 10")
	}
}

func Test_ContainsAnyElement(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(3)
	a.Add(5)

	b := NewSet[int]()
	a.Add(2)
	a.Add(4)
	a.Add(6)

	if ret := a.ContainsAnyElement(b); ret {
		t.Errorf("set a not contain any element in set b")
	}

	a.Add(10)

	if ret := a.ContainsAnyElement(b); ret {
		t.Errorf("set a not contain any element in set b")
	}

	b.Add(10)

	if ret := a.ContainsAnyElement(b); !ret {
		t.Errorf("set a contain 10")
	}
}
func Test_ClearSet(t *testing.T) {
	a := makeSetInt([]int{2, 5, 9, 10})

	a.Clear()

	if a.Cardinality() != 0 {
		t.Error("ClearSet should be an empty set")
	}
}

func Test_ClearUnsafeSet(t *testing.T) {
	a := makeUnsafeSetInt([]int{2, 5, 9, 10})

	a.Clear()

	if a.Cardinality() != 0 {
		t.Error("ClearSet should be an empty set")
	}
}

func Test_CardinalitySet(t *testing.T) {
	a := NewSet[int]()

	if a.Cardinality() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(1)

	if a.Cardinality() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Remove(1)

	if a.Cardinality() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(9)

	if a.Cardinality() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Clear()

	if a.Cardinality() != 0 {
		t.Error("set should have a size of 1")
	}
}

func Test_CardinalityUnsafeSet(t *testing.T) {
	a := NewThreadUnsafeSet[int]()

	if a.Cardinality() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(1)

	if a.Cardinality() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Remove(1)

	if a.Cardinality() != 0 {
		t.Error("set should be an empty set")
	}

	a.Add(9)

	if a.Cardinality() != 1 {
		t.Error("set should have a size of 1")
	}

	a.Clear()

	if a.Cardinality() != 0 {
		t.Error("set should have a size of 1")
	}
}

func Test_SetIsSubset(t *testing.T) {
	a := makeSetInt([]int{1, 2, 3, 5, 7})

	b := NewSet[int]()
	b.Add(3)
	b.Add(5)
	b.Add(7)

	if !b.IsSubset(a) {
		t.Error("set b should be a subset of set a")
	}

	b.Add(72)

	if b.IsSubset(a) {
		t.Error("set b should not be a subset of set a because it contains 72 which is not in the set of a")
	}
}

func Test_SetIsProperSubset(t *testing.T) {
	a := makeSetInt([]int{1, 2, 3, 5, 7})
	b := makeSetInt([]int{7, 5, 3, 2, 1})

	if !a.IsSubset(b) {
		t.Error("set a should be a subset of set b")
	}
	if a.IsProperSubset(b) {
		t.Error("set a should not be a proper subset of set b (they're equal)")
	}

	b.Add(72)

	if !a.IsSubset(b) {
		t.Error("set a should be a subset of set b")
	}
	if !a.IsProperSubset(b) {
		t.Error("set a should be a proper subset of set b")
	}
}

func Test_UnsafeSetIsSubset(t *testing.T) {
	a := makeUnsafeSetInt([]int{1, 2, 3, 5, 7})

	b := NewThreadUnsafeSet[int]()
	b.Add(3)
	b.Add(5)
	b.Add(7)

	if !b.IsSubset(a) {
		t.Error("set b should be a subset of set a")
	}

	b.Add(72)

	if b.IsSubset(a) {
		t.Error("set b should not be a subset of set a because it contains 72 which is not in the set of a")
	}
}

func Test_UnsafeSetIsProperSubset(t *testing.T) {
	a := makeUnsafeSetInt([]int{1, 2, 3, 5, 7})
	b := NewThreadUnsafeSet[int]()
	b.Add(7)
	b.Add(1)
	b.Add(5)
	b.Add(3)
	b.Add(2)

	if !a.IsSubset(b) {
		t.Error("set a should be a subset of set b")
	}
	if a.IsProperSubset(b) {
		t.Error("set a should not be a proper subset of set b (they're equal)")
	}

	b.Add(72)

	if !a.IsSubset(b) {
		t.Error("set a should be a subset of set b")
	}
	if !a.IsProperSubset(b) {
		t.Error("set a should be a proper subset of set b because set b has 72")
	}
}

func Test_SetIsSuperset(t *testing.T) {
	a := NewSet[int]()
	a.Add(9)
	a.Add(5)
	a.Add(2)
	a.Add(1)
	a.Add(11)

	b := NewSet[int]()
	b.Add(5)
	b.Add(2)
	b.Add(11)

	if !a.IsSuperset(b) {
		t.Error("set a should be a superset of set b")
	}

	b.Add(42)

	if a.IsSuperset(b) {
		t.Error("set a should not be a superset of set b because set b has a 42")
	}
}

func Test_SetIsProperSuperset(t *testing.T) {
	a := NewSet[int]()
	a.Add(5)
	a.Add(2)
	a.Add(11)

	b := NewSet[int]()
	b.Add(2)
	b.Add(5)
	b.Add(11)

	if !a.IsSuperset(b) {
		t.Error("set a should be a superset of set b")
	}
	if a.IsProperSuperset(b) {
		t.Error("set a should not be a proper superset of set b (they're equal)")
	}

	a.Add(9)

	if !a.IsSuperset(b) {
		t.Error("set a should be a superset of set b")
	}
	if !a.IsProperSuperset(b) {
		t.Error("set a not be a proper superset of set b because set a has a 9")
	}

	b.Add(42)

	if a.IsSuperset(b) {
		t.Error("set a should not be a superset of set b because set b has a 42")
	}
	if a.IsProperSuperset(b) {
		t.Error("set a should not be a proper superset of set b because set b has a 42")
	}
}

func Test_UnsafeSetIsSuperset(t *testing.T) {
	a := NewThreadUnsafeSet[int]()
	a.Add(9)
	a.Add(5)
	a.Add(2)
	a.Add(1)
	a.Add(11)

	b := NewThreadUnsafeSet[int]()
	b.Add(5)
	b.Add(2)
	b.Add(11)

	if !a.IsSuperset(b) {
		t.Error("set a should be a superset of set b")
	}

	b.Add(42)

	if a.IsSuperset(b) {
		t.Error("set a should not be a superset of set b because set a has a 42")
	}
}

func Test_UnsafeSetIsProperSuperset(t *testing.T) {
	a := NewThreadUnsafeSet[int]()
	a.Add(5)
	a.Add(2)
	a.Add(11)

	b := NewThreadUnsafeSet[int]()
	b.Add(2)
	b.Add(5)
	b.Add(11)

	if !a.IsSuperset(b) {
		t.Error("set a should be a superset of set b")
	}
	if a.IsProperSuperset(b) {
		t.Error("set a should not be a proper superset of set b (they're equal)")
	}

	a.Add(9)

	if !a.IsSuperset(b) {
		t.Error("set a should be a superset of set b")
	}
	if !a.IsProperSuperset(b) {
		t.Error("set a not be a proper superset of set b because set a has a 9")
	}

	b.Add(42)

	if a.IsSuperset(b) {
		t.Error("set a should not be a superset of set b because set b has a 42")
	}
	if a.IsProperSuperset(b) {
		t.Error("set a should not be a proper superset of set b because set b has a 42")
	}
}

func Test_SetUnion(t *testing.T) {
	a := NewSet[int]()

	b := NewSet[int]()
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	c := a.Union(b)

	if c.Cardinality() != 5 {
		t.Error("set c is unioned with an empty set and therefore should have 5 elements in it")
	}

	d := NewSet[int]()
	d.Add(10)
	d.Add(14)
	d.Add(0)

	e := c.Union(d)
	if e.Cardinality() != 8 {
		t.Error("set e should have 8 elements in it after being unioned with set c to d")
	}

	f := NewSet[int]()
	f.Add(14)
	f.Add(3)

	g := f.Union(e)
	if g.Cardinality() != 8 {
		t.Error("set g should still have 8 elements in it after being unioned with set f that has duplicates")
	}
}

func Test_UnsafeSetUnion(t *testing.T) {
	a := NewThreadUnsafeSet[int]()

	b := NewThreadUnsafeSet[int]()
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	c := a.Union(b)

	if c.Cardinality() != 5 {
		t.Error("set c is unioned with an empty set and therefore should have 5 elements in it")
	}

	d := NewThreadUnsafeSet[int]()
	d.Add(10)
	d.Add(14)
	d.Add(0)

	e := c.Union(d)
	if e.Cardinality() != 8 {
		t.Error("set e should have 8 elements in it after being unioned with set c to d")
	}

	f := NewThreadUnsafeSet[int]()
	f.Add(14)
	f.Add(3)

	g := f.Union(e)
	if g.Cardinality() != 8 {
		t.Error("set g should still have 8 elements in it after being unioned with set f that has duplicates")
	}
}

func Test_SetIntersect(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(3)
	a.Add(5)

	b := NewSet[int]()
	a.Add(2)
	a.Add(4)
	a.Add(6)

	c := a.Intersect(b)

	if c.Cardinality() != 0 {
		t.Error("set c should be the empty set because there is no common items to intersect")
	}

	a.Add(10)
	b.Add(10)

	d := a.Intersect(b)

	if !(d.Cardinality() == 1 && d.Contains(10)) {
		t.Error("set d should have a size of 1 and contain the item 10")
	}
}

func Test_UnsafeSetIntersect(t *testing.T) {
	a := NewThreadUnsafeSet[int]()
	a.Add(1)
	a.Add(3)
	a.Add(5)

	b := NewThreadUnsafeSet[int]()
	a.Add(2)
	a.Add(4)
	a.Add(6)

	c := a.Intersect(b)

	if c.Cardinality() != 0 {
		t.Error("set c should be the empty set because there is no common items to intersect")
	}

	a.Add(10)
	b.Add(10)

	d := a.Intersect(b)

	if !(d.Cardinality() == 1 && d.Contains(10)) {
		t.Error("set d should have a size of 1 and contain the item 10")
	}
}

func Test_SetDifference(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewSet[int]()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	c := a.Difference(b)

	if !(c.Cardinality() == 1 && c.Contains(2)) {
		t.Error("the difference of set a to b is the set of 1 item: 2")
	}
}

func Test_UnsafeSetDifference(t *testing.T) {
	a := NewThreadUnsafeSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewThreadUnsafeSet[int]()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	c := a.Difference(b)

	if !(c.Cardinality() == 1 && c.Contains(2)) {
		t.Error("the difference of set a to b is the set of 1 item: 2")
	}
}

func Test_SetSymmetricDifference(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(45)

	b := NewSet[int]()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	c := a.SymmetricDifference(b)

	if !(c.Cardinality() == 6 && c.Contains(2) && c.Contains(45) && c.Contains(4) && c.Contains(5) && c.Contains(6) && c.Contains(99)) {
		t.Error("the symmetric difference of set a to b is the set of 6 items: 2, 45, 4, 5, 6, 99")
	}
}

func Test_UnsafeSetSymmetricDifference(t *testing.T) {
	a := NewThreadUnsafeSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(45)

	b := NewThreadUnsafeSet[int]()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	c := a.SymmetricDifference(b)

	if !(c.Cardinality() == 6 && c.Contains(2) && c.Contains(45) && c.Contains(4) && c.Contains(5) && c.Contains(6) && c.Contains(99)) {
		t.Error("the symmetric difference of set a to b is the set of 6 items: 2, 45, 4, 5, 6, 99")
	}
}

func Test_SetEqual(t *testing.T) {
	a := NewSet[int]()
	b := NewSet[int]()

	if !a.Equal(b) {
		t.Error("Both a and b are empty sets, and should be equal")
	}

	a.Add(10)

	if a.Equal(b) {
		t.Error("a should not be equal to b because b is empty and a has item 1 in it")
	}

	b.Add(10)

	if !a.Equal(b) {
		t.Error("a is now equal again to b because both have the item 10 in them")
	}

	b.Add(8)
	b.Add(3)
	b.Add(47)

	if a.Equal(b) {
		t.Error("b has 3 more elements in it so therefore should not be equal to a")
	}

	a.Add(8)
	a.Add(3)
	a.Add(47)

	if !a.Equal(b) {
		t.Error("a and b should be equal with the same number of elements")
	}
}

func Test_UnsafeSetEqual(t *testing.T) {
	a := NewThreadUnsafeSet[int]()
	b := NewThreadUnsafeSet[int]()

	if !a.Equal(b) {
		t.Error("Both a and b are empty sets, and should be equal")
	}

	a.Add(10)

	if a.Equal(b) {
		t.Error("a should not be equal to b because b is empty and a has item 1 in it")
	}

	b.Add(10)

	if !a.Equal(b) {
		t.Error("a is now equal again to b because both have the item 10 in them")
	}

	b.Add(8)
	b.Add(3)
	b.Add(47)

	if a.Equal(b) {
		t.Error("b has 3 more elements in it so therefore should not be equal to a")
	}

	a.Add(8)
	a.Add(3)
	a.Add(47)

	if !a.Equal(b) {
		t.Error("a and b should be equal with the same number of elements")
	}
}

func Test_SetClone(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)

	b := a.Clone()

	if !a.Equal(b) {
		t.Error("Clones should be equal")
	}

	a.Add(3)
	if a.Equal(b) {
		t.Error("a contains one more element, they should not be equal")
	}

	c := a.Clone()
	c.Remove(1)

	if a.Equal(c) {
		t.Error("C contains one element less, they should not be equal")
	}
}

func Test_UnsafeSetClone(t *testing.T) {
	a := NewThreadUnsafeSet[int]()
	a.Add(1)
	a.Add(2)

	b := a.Clone()

	if !a.Equal(b) {
		t.Error("Clones should be equal")
	}

	a.Add(3)
	if a.Equal(b) {
		t.Error("a contains one more element, they should not be equal")
	}

	c := a.Clone()
	c.Remove(1)

	if a.Equal(c) {
		t.Error("C contains one element less, they should not be equal")
	}
}

func Test_Each(t *testing.T) {
	a := NewSet[string]()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewSet[string]()
	a.Each(func(elem string) bool {
		b.Add(elem)
		return false
	})

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Each) through the first set")
	}

	var count int
	a.Each(func(elem string) bool {
		if count == 2 {
			return true
		}
		count++
		return false
	})
	if count != 2 {
		t.Error("Iteration should stop on the way")
	}
}

func Test_Iter(t *testing.T) {
	a := NewSet[string]()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewSet[string]()
	for val := range a.Iter() {
		b.Add(val)
	}

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Iter) through the first set")
	}
}

func Test_UnsafeIter(t *testing.T) {
	a := NewThreadUnsafeSet[string]()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewThreadUnsafeSet[string]()
	for val := range a.Iter() {
		b.Add(val)
	}

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Iter) through the first set")
	}
}

func Test_Iterator(t *testing.T) {
	a := NewSet[string]()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewSet[string]()
	for val := range a.Iterator().C {
		b.Add(val)
	}

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Iterator) through the first set")
	}
}

func Test_UnsafeIterator(t *testing.T) {
	a := NewThreadUnsafeSet[string]()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewThreadUnsafeSet[string]()
	for val := range a.Iterator().C {
		b.Add(val)
	}

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Iterator) through the first set")
	}
}

func Test_IteratorStop(t *testing.T) {
	a := NewSet[string]()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	it := a.Iterator()
	it.Stop()
	for range it.C {
		t.Error("The iterating (Iterator) did not stop after Stop() has been called")
	}
}

func Test_PopSafe(t *testing.T) {
	a := NewSet[string]()

	a.Add("a")
	a.Add("b")
	a.Add("c")
	a.Add("d")

	aPop := func() (v string) {
		v, _ = a.Pop()
		return
	}

	captureSet := NewSet[string]()
	captureSet.Add(aPop())
	captureSet.Add(aPop())
	captureSet.Add(aPop())
	captureSet.Add(aPop())
	finalNil := aPop()

	if captureSet.Cardinality() != 4 {
		t.Error("unexpected captureSet cardinality; should be 4")
	}

	if a.Cardinality() != 0 {
		t.Error("unepxected a cardinality; should be zero")
	}

	if !captureSet.Contains("c", "a", "d", "b") {
		t.Error("unexpected result set; should be a,b,c,d (any order is fine")
	}

	if finalNil != "" {
		t.Error("when original set is empty, further pops should result in nil")
	}
}

func Test_PopUnsafe(t *testing.T) {
	a := NewThreadUnsafeSet[string]()

	a.Add("a")
	a.Add("b")
	a.Add("c")
	a.Add("d")

	aPop := func() (v string) {
		v, _ = a.Pop()
		return
	}

	captureSet := NewThreadUnsafeSet[string]()
	captureSet.Add(aPop())
	captureSet.Add(aPop())
	captureSet.Add(aPop())
	captureSet.Add(aPop())
	finalNil := aPop()

	if captureSet.Cardinality() != 4 {
		t.Error("unexpected captureSet cardinality; should be 4")
	}

	if a.Cardinality() != 0 {
		t.Error("unepxected a cardinality; should be zero")
	}

	if !captureSet.Contains("c", "a", "d", "b") {
		t.Error("unexpected result set; should be a,b,c,d (any order is fine")
	}

	if finalNil != "" {
		t.Error("when original set is empty, further pops should result in nil")
	}
}

func Test_PopNSafe(t *testing.T) {
	a := NewSet[string]()
	a.Add("a")
	a.Add("b")
	a.Add("c")
	a.Add("d")

	// Test pop with n <= 0
	items, count := a.PopN(0)
	if count != 0 {
		t.Errorf("expected 0 items popped, got %d", count)
	}
	items, count = a.PopN(-1)
	if count != 0 {
		t.Errorf("expected 0 items popped, got %d", count)
	}

	captureSet := NewSet[string]()

	// Test pop 2 items
	items, count = a.PopN(2)
	if count != 2 {
		t.Errorf("expected 2 items popped, got %d", count)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items in slice, got %d", len(items))
	}
	if a.Cardinality() != 2 {
		t.Errorf("expected 2 items remaining, got %d", a.Cardinality())
	}
	captureSet.Append(items...)

	// Test pop more than remaining
	items, count = a.PopN(3)
	if count != 2 {
		t.Errorf("expected 2 items popped, got %d", count)
	}
	if a.Cardinality() != 0 {
		t.Errorf("expected 0 items remaining, got %d", a.Cardinality())
	}
	captureSet.Append(items...)

	// Test pop from empty set
	items, count = a.PopN(1)
	if count != 0 {
		t.Errorf("expected 0 items popped, got %d", count)
	}
	if len(items) != 0 {
		t.Errorf("expected empty slice, got %d items", len(items))
	}

	if !captureSet.Contains("c", "a", "d", "b") {
		t.Error("unexpected result set; should be a,b,c,d (any order is fine")
	}
}

func Test_PopNUnsafe(t *testing.T) {
	a := NewThreadUnsafeSet[string]()
	a.Add("a")
	a.Add("b")
	a.Add("c")
	a.Add("d")

	// Test pop with n <= 0
	items, count := a.PopN(0)
	if count != 0 {
		t.Errorf("expected 0 items popped, got %d", count)
	}
	items, count = a.PopN(-1)
	if count != 0 {
		t.Errorf("expected 0 items popped, got %d", count)
	}

	captureSet := NewThreadUnsafeSet[string]()

	// Test pop 2 items
	items, count = a.PopN(2)
	if count != 2 {
		t.Errorf("expected 2 items popped, got %d", count)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items in slice, got %d", len(items))
	}
	if a.Cardinality() != 2 {
		t.Errorf("expected 2 items remaining, got %d", a.Cardinality())
	}
	captureSet.Append(items...)

	// Test pop more than remaining
	items, count = a.PopN(3)
	if count != 2 {
		t.Errorf("expected 2 items popped, got %d", count)
	}
	if a.Cardinality() != 0 {
		t.Errorf("expected 0 items remaining, got %d", a.Cardinality())
	}
	captureSet.Append(items...)

	// Test pop from empty set
	items, count = a.PopN(1)
	if count != 0 {
		t.Errorf("expected 0 items popped, got %d", count)
	}
	if len(items) != 0 {
		t.Errorf("expected empty slice, got %d items", len(items))
	}

	if !captureSet.Contains("c", "a", "d", "b") {
		t.Error("unexpected result set; should be a,b,c,d (any order is fine")
	}
}

func Test_EmptySetProperties(t *testing.T) {
	empty := NewSet[string]()

	a := NewSet[string]()
	a.Add("1")
	a.Add("foo")
	a.Add("bar")

	b := NewSet[string]()
	b.Add("one")
	b.Add("two")
	b.Add("3")
	b.Add("4")

	if !empty.IsSubset(a) || !empty.IsSubset(b) {
		t.Error("The empty set is supposed to be a subset of all sets")
	}

	if !a.IsSuperset(empty) || !b.IsSuperset(empty) {
		t.Error("All sets are supposed to be a superset of the empty set")
	}

	if !empty.IsSubset(empty) || !empty.IsSuperset(empty) {
		t.Error("The empty set is supposed to be a subset and a superset of itself")
	}

	c := a.Union(empty)
	if !c.Equal(a) {
		t.Error("The union of any set with the empty set is supposed to be equal to itself")
	}

	c = a.Intersect(empty)
	if !c.Equal(empty) {
		t.Error("The intersection of any set with the empty set is supposed to be the empty set")
	}

	if empty.Cardinality() != 0 {
		t.Error("Cardinality of the empty set is supposed to be zero")
	}
}

func Test_ToSliceUnthreadsafe(t *testing.T) {
	s := makeUnsafeSetInt([]int{1, 2, 3})
	setAsSlice := s.ToSlice()
	if len(setAsSlice) != s.Cardinality() {
		t.Errorf("Set length is incorrect: %v", len(setAsSlice))
	}

	for _, i := range setAsSlice {
		if !s.Contains(i) {
			t.Errorf("Set is missing element: %v", i)
		}
	}
}

func Test_NewSetFromMapKey_Ints(t *testing.T) {
	m := map[int]int{
		5: 5,
		2: 3,
	}

	s := NewSetFromMapKeys(m)

	if len(m) != s.Cardinality() {
		t.Errorf("Length of Set is not the same as the map. Expected: %d. Actual: %d", len(m), s.Cardinality())
	}

	for k := range m {
		if !s.Contains(k) {
			t.Errorf("Element %d not found in map: %v", k, m)
		}
	}
}

func Test_NewSetFromMapKey_Strings(t *testing.T) {
	m := map[int]int{
		5: 5,
		2: 3,
	}

	s := NewSetFromMapKeys(m)

	if len(m) != s.Cardinality() {
		t.Errorf("Length of Set is not the same as the map. Expected: %d. Actual: %d", len(m), s.Cardinality())
	}

	for k := range m {
		if !s.Contains(k) {
			t.Errorf("Element %q not found in map: %v", k, m)
		}
	}
}

func Test_NewThreadUnsafeSetFromMapKey_Ints(t *testing.T) {
	m := map[int]int{
		5: 5,
		2: 3,
	}

	s := NewThreadUnsafeSetFromMapKeys(m)

	if len(m) != s.Cardinality() {
		t.Errorf("Length of Set is not the same as the map. Expected: %d. Actual: %d", len(m), s.Cardinality())
	}

	for k := range m {
		if !s.Contains(k) {
			t.Errorf("Element %d not found in map: %v", k, m)
		}
	}
}

func Test_NewThreadUnsafeSetFromMapKey_Strings(t *testing.T) {
	m := map[int]int{
		5: 5,
		2: 3,
	}

	s := NewThreadUnsafeSetFromMapKeys(m)

	if len(m) != s.Cardinality() {
		t.Errorf("Length of Set is not the same as the map. Expected: %d. Actual: %d", len(m), s.Cardinality())
	}

	for k := range m {
		if !s.Contains(k) {
			t.Errorf("Element %q not found in map: %v", k, m)
		}
	}
}

func Test_Elements(t *testing.T) {
	a := NewSet[string]()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewSet[string]()
	Elements(a)(func(elem string) bool {
		b.Add(elem)
		return true
	})

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Each) through the first set")
	}

	var count int
	Elements(a)(func(elem string) bool {
		if count == 2 {
			return false
		}
		count++
		return true
	})

	if count != 2 {
		t.Error("Iteration should stop on the way")
	}
}

func Test_Example(t *testing.T) {
	/*
	   requiredClasses := NewSet()
	   requiredClasses.Add("Cooking")
	   requiredClasses.Add("English")
	   requiredClasses.Add("Math")
	   requiredClasses.Add("Biology")

	   scienceSlice := []interface{}{"Biology", "Chemistry"}
	   scienceClasses := NewSet(scienceSlice)

	   electiveClasses := NewSet()
	   electiveClasses.Add("Welding")
	   electiveClasses.Add("Music")
	   electiveClasses.Add("Automotive")

	   bonusClasses := NewSet()
	   bonusClasses.Add("Go Programming")
	   bonusClasses.Add("Python Programming")

	   //Show me all the available classes I can take
	   allClasses := requiredClasses.Union(scienceClasses).Union(electiveClasses).Union(bonusClasses)
	   fmt.Println(allClasses) //Set{English, Chemistry, Automotive, Cooking, Math, Biology, Welding, Music, Go Programming}

	   //Is cooking considered a science class?
	   fmt.Println(scienceClasses.Contains("Cooking")) //false

	   //Show me all classes that are not science classes, since I hate science.
	   fmt.Println(allClasses.Difference(scienceClasses)) //Set{English, Automotive, Cooking, Math, Welding, Music, Go Programming}

	   //Which science classes are also required classes?
	   fmt.Println(scienceClasses.Intersect(requiredClasses)) //Set{Biology}

	   //How many bonus classes do you offer?
	   fmt.Println(bonusClasses.Cardinality()) //2

	   //Do you have the following classes? Welding, Automotive and English?
	   fmt.Println(allClasses.ContainsAll("Welding", "Automotive", "English"))
	*/
}
