package ordenamiento

import (
	"sort"
	"fmt"
) 

func Ordenamiento(s []int) []int {
	sort.Sort(sort.IntSlice(s))
    fmt.Println(s)

	return s
}