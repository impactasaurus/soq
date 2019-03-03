package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/impactasaurus/soq"
	. "github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
)

func main() {
	long := mustRun(promptui.Prompt{
		Label: "Questionnaire Name",
	})

	short := mustRunSP(promptui.Prompt{
		Label: "Questionnaire Short Name",
	})

	out := soq.Questionnaire{
		ID:      uuid(),
		Version: "1.0.0",
		Name:    long,
		Short:   short,
		Changelog: []soq.Changes{{
			Version:   "1.0.0",
			Timestamp: time.Now().Format(time.RFC3339),
			Changes:   []string{"initial version"},
		}},
	}

	noCats := mustRunI(promptui.Prompt{
		Label:    "Number of scorings",
		Validate: mustBeNumber,
	})
	out.Scorings = make([]*soq.Scoring, noCats)
	for i := range out.Scorings {
		fmt.Println(Bold(Cyan(fmt.Sprintf("Scoring %d", i+1))))
		out.Scorings[i] = promptCategory()
	}

	fmt.Println(Bold(Blue("Default Scale")))
	defaultScale := promptScale()

	noQuestions := mustRunI(promptui.Prompt{
		Label:    "Number of questions",
		Validate: mustBeNumber,
	})
	out.Questions = make([]*soq.Question, noQuestions)
	for i := range out.Questions {
		fmt.Println(Bold(Magenta(fmt.Sprintf("Question %d", i+1))))
		var q soq.Question
		q = promptQuestion(defaultScale, out.Scorings)
		out.Questions[i] = &q
	}

	j, _ := json.MarshalIndent(out, "", "  ")

	file := fmt.Sprintf("%s.json", *short)
	err := ioutil.WriteFile(file, j, 0644)
	if err != nil {
		log.Fatalf("Failed to save file: %s", err.Error())
	}
	fmt.Println(fmt.Sprintf("Saved to %s", file))
}
