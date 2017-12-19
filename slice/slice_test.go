package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEach(t *testing.T) {
	{
		array := []int{1, 2, 3, 4, 5}
		result := 15

		acc := 0
		Each(array, func(i int) {
			acc += array[i]
		})

		if acc != result {
			t.Fatalf("Each: accumulate %v, want %v, get %v", array, result, acc)
		}
	}

	{
		array := []string{"a", "b", "c", "d"}
		result := "abcd"

		acc := ""
		Each(array, func(i int) {
			acc += array[i]
		})

		if acc != result {
			t.Fatalf("Each: accumulate %v, want %v, get %v", array, result, acc)
		}
	}
}

func TestFilter(t *testing.T) {
	{
		array := []int{1, 2, 3, 4, 5}
		expect := []int{1, 3, 5}
		pred := func(i int) bool { return array[i]%2 == 1 }

		result := Filter(array, pred)
		if reflect.DeepEqual(result, expect) == false {
			t.Fatalf("Filter: want %v, get %v", expect, result)
		}
	}

	{
		array := []string{"a", "ab", "ac", "bd", "e"}
		expect := []string{"ab", "ac", "bd"}
		pred := func(i int) bool { return len(array[i]) > 1 }

		result := Filter(array, pred)
		if reflect.DeepEqual(result, expect) == false {
			t.Fatalf("Filter: want %v, get %v", expect, result)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		array  interface{}
		value  interface{}
		result int
	}{
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]string{"a", "ab", "ba", "bb", "b"}, "b", 4},
		{[]string{"a", "ab", "ba", "bb", "b"}, "c", -1},
	}

	for _, test := range tests {
		index := Find(test.array, test.value)
		if index != test.result {
			t.Fatalf("Find(%v, %v), want %v, get %v",
				test.array, test.value, test.result, index)
		}
	}
}

func TestFindIf(t *testing.T) {
	{
		array := []int{1, 2, 3, 4, 5}
		value := 5
		result := 4

		index := FindIf(array, func(i int) bool {
			return array[i] == value
		})

		if index != result {
			t.Fatalf("FindIf(%v, %v), want %v, get %v", array, value, result, index)
		}
	}

	{
		array := []string{"a", "b", "c", "d"}
		value := "d"
		result := 3

		index := FindIf(array, func(i int) bool {
			return array[i] == value
		})

		if index != result {
			t.Fatalf("FindIf(%v, %v), want %v, get %v", array, value, result, index)
		}
	}
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		array  interface{}
		value  interface{}
		result bool
	}{
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 6, false},
		{[]string{"a", "ab", "ba", "bb", "b"}, "b", true},
		{[]string{"a", "ab", "ba", "bb", "b"}, "c", false},
	}

	for _, test := range tests {
		res := Includes(test.array, test.value)
		if res != test.result {
			t.Fatalf("Includes(%v, %v), want %v, get %v",
				test.array, test.value, test.result, res)
		}
	}
}

func TestMap(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	expect := []string{"1", "2", "3", "4", "5"}

	result := Map(array, func(i int) string {
		return fmt.Sprintf("%d", array[i])
	}).([]string)

	if reflect.DeepEqual(result, expect) == false {
		t.Fatalf("Map want %v, get %v", expect, result)
	}
}
