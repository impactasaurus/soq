package main

import (
	"fmt"

	"github.com/impactasaurus/soq"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
)

func promptPoint() *soq.Point {
	val := mustRunI(promptui.Prompt{
		Label:    "Point value",
		Validate: mustBeNumber,
	})
	label := mustRunSP(promptui.Prompt{
		Label: "Point Label",
	})
	return &soq.Point{
		Label: label,
		Value: float64(val),
	}
}

func promptScale() []*soq.Point {
	noPoints := mustRunI(promptui.Prompt{
		Label:    "Number of points",
		Validate: mustBeNumber,
	})
	out := make([]*soq.Point, noPoints)
	for i := range out {
		fmt.Println(aurora.Blue(fmt.Sprintf("Scale point %d", i+1)))
		out[i] = promptPoint()
	}
	return out
}

func invert(orig []*soq.Point) []*soq.Point {
	out := make([]*soq.Point, len(orig))
	for i := range orig {
		out[i] = &soq.Point{
			Label: orig[i].Label,
			Value: orig[len(orig)-(i+1)].Value,
		}
	}
	return out
}
