package questionnaires_test

import (
	"testing"

	soq "github.com/impactasaurus/soq-api"
	"github.com/impactasaurus/soq-api/questionnaires"
)

func TestQuestionnaires(t *testing.T) {
	qq, err := questionnaires.LoadDirectory(".")
	if err != nil {
		t.Fatal(err)
	}
	if len(qq) == 0 {
		t.Fatal("expecting questionnaires")
	}
	for _, q := range qq {
		t.Run(q.Name, func(t *testing.T) {
			if q.ID == "" {
				t.Errorf("ID not specified")
			}
			testQuestions(t, q)
		})
	}
}

func testQuestions(t *testing.T, qs soq.Questionnaire) {
	if len(qs.Questions) == 0 {
		t.Errorf("no questions")
	}
	seen := map[string]bool{}
	for _, q := range qs.Questions {
		aq := *q
		switch c := aq.(type) {
		case *soq.LikertQuestion:
			if _, ok := seen[c.ID]; ok {
				t.Errorf("duplicate question IDs: %s", c.ID)
			}
			seen[c.ID] = true
		default:
			t.Errorf("unknown question type")
		}
	}
}
