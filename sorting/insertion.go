package sorting

import (
	"cmp"
)

func Insertion[S ~[]T, T cmp.Ordered](slice S) S {
	for i := 1; i < len(slice); i++ {
		for j:= i ; j > 0 && slice[j-1] > slice[j]; j-- {
		slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}
	return slice
}
