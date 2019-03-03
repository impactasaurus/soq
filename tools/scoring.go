package main

import (
	"log"

	"github.com/impactasaurus/soq"
	"github.com/manifoldco/promptui"
)

func promptCategory() *soq.Scoring {

	name := mustRun(promptui.Prompt{
		Label:    "Scoring Name",
		Validate: mustBeString,
	})

	description := mustRunSP(promptui.Prompt{
		Label: "Scoring Description",
	})

	aa := []string{string(soq.AggregationMean), string(soq.AggregationSum)}
	_, a, err := (&promptui.Select{
		Label: "Select Aggregation",
		Items: aa,
	}).Run()
	if err != nil {
		log.Fatalf("Creating questionnaire failed: %s", err.Error())
	}

	return &soq.Scoring{
		ID:          uuid(),
		Name:        name,
		Description: description,
		Aggregation: soq.Aggregation(a),
	}
}
