package main

import (
	"fmt"
	"sort"
)

type task struct {
	start int
	end   int
}

type tasks []task

func (t tasks) Len() int           { return len(t) }
func (t tasks) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t tasks) Less(i, j int) bool { return t[i].end < t[j].end }

func newTasks(t []task) tasks {
	return t
}

func (t tasks) calc() int {
	sort.Sort(&t)

	fmt.Printf("%#v\n", t)
	var count int
	var progress int
	for _, val := range t {
		if val.start > progress {
			count++
			progress = val.end
		}
	}

	return count
}
