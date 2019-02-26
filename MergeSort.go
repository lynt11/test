package main

import "fmt"

type MergeSort struct {
}

func (object MergeSort) Sort(values []int) {
	mergeSort(values)
}
func mergeSort(values []int) {
	length := len(values)
	//gap计数
	gap := 1
	for gap < length {
		fmt.Println("--merge by gap", gap)
		mergeByGap(values, gap)
		gap = gap * 2
		fmt.Println("--the tuple is", values)
	}
}
func mergeByGap(values []int, gap int) {
	length := len(values)
	for i := 0; i < length; i += 2 * gap {
		length := len(values)
		for i := 0; i < length; i += 2 * gap {
			mergeTwoValues(values, i, gap)
		}
	}
}
func mergeTwoValues(values []int, start int, gap int) {
	length := len(values)
	if start+gap >= length {
		return
	}
	slice := make([]int, 2*gap)
	lpos, rpos, slicepos := start, start+gap, 0
	for lpos < start+gap && (rpos < start+2*gap && rpos < length) {
		if values[lpos] <= values[rpos] {
			slice[slicepos] = values[lpos]
			lpos++
		} else {
			slice[slicepos] = values[rpos]
			rpos++
		}
		slicepos++
	}
	if lpos != start+gap {
		for lpos < start+gap {
			slice[slicepos] = values[lpos]
			lpos++
			slicepos++
		}
	} else if rpos != start+2*gap && rpos != length {
		for rpos < start+2*gap && rpos < length {
			slice[slicepos] = values[rpos]
			rpos++
			slicepos++
		}
	}
	for i := 0; i < slicepos; i++ {
		values[i+start] = slice[i]
	}
	fmt.Println("    ----  merge result is ", slice[:slicepos])
	// fmt.Println("    ----  merge result is ", values[start:start+2*gap])
}
func main() {

}
