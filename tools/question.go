package main

import (
	"log"

	"github.com/impactasaurus/soq"
	"github.com/manifoldco/promptui"
)

func adjustScorings(id string, scorings []*soq.Scoring) {
	if len(scorings) == 0 {
		return
	}
	ss := make([]string, len(scorings)+1)
	ss[0] = "None"
	for i := range scorings {
		ss[i+1] = scorings[i].Name
	}
	idx, _, err := (&promptui.Select{
		Label: "Select Scoring",
		Items: ss,
	}).Run()
	if err != nil {
		log.Fatalf("Creating questionnaire failed: %s", err.Error())
	}
	if idx == 0 {
		return
	}
	scorings[idx-1].Questions = append(scorings[idx-1].Questions, &id)
}

func promptQuestion(defaultScale []*soq.Point, scorings []*soq.Scoring) soq.LikertQuestion {

	id := uuid()

	name := mustRun(promptui.Prompt{
		Label:    "Question",
		Validate: mustBeString,
	})
	short := mustRunSP(promptui.Prompt{
		Label: "Short",
	})
	desc := mustRunSP(promptui.Prompt{
		Label: "Description",
	})

	i := mustBool("Is Inverse")
	s := defaultScale
	if i {
		s = invert(defaultScale)
	}

	adjustScorings(id, scorings)

	return soq.LikertQuestion{
		ID:          id,
		Question:    name,
		Short:       short,
		Description: desc,
		Scale:       s,
	}
}
