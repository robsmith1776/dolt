package main

import (
	"github.com/attic-labs/noms/go/types/ase"
	"log"
	"time"
)

type ASEBench struct {
	ase *ase.AsyncSortedEdits
}

func NewASEBench(slSize, asyncConcurrency, sortConcurrency int) *ASEBench {
	return &ASEBench{ase.NewAsyncSortedEdits(slSize, asyncConcurrency, sortConcurrency)}
}

func (msb *ASEBench) GetName() string {
	return "async sorted edits"
}

func (msb *ASEBench) AddEdits(nextEdit NextEdit) {
	k, v := nextEdit()

	for k != nil {
		msb.ase.Set(k, v)
		k, v = nextEdit()
	}

	startFinish := time.Now()
	msb.ase.FinishedEditing()
	endFinish := time.Now()
	finishDelta := endFinish.Sub(startFinish)

	log.Println("finish took", finishDelta.Seconds(), "seconds")
}

func (msb *ASEBench) SortEdits() {
	msb.ase.Sort()

	itr := msb.ase.Iterator()
	numItems, inOrder := ase.IsInOrder(itr)
	log.Println("in order:", inOrder, "- num items:", numItems)
}

func (msb *ASEBench) Map() {
}
