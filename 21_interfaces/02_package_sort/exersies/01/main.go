/*
type people []string
studyGroup := people{"Zeno", "John", "Al", "Jenny"}
*/

package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	//fmt.Println(studyGroup)
	//sort.Strings(studyGroup)
	//fmt.Println(studyGroup)
	//sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	//fmt.Println(studyGroup)

	fmt.Println(studyGroup)
	sort.Sort(people(studyGroup))
	fmt.Println(studyGroup)
}
