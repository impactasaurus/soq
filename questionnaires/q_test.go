package questionnaires_test

import (
	"testing"

	"net/http"

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
			if q.Attribution == nil {
				t.Logf("should questionnaire %s have attribution?", q.Name)
			}
			if q.License == "" {
				t.Errorf("questionnaire %s missing license information", q.Name)
			}
			testQuestions(t, q)
			testScorings(t, q)
			testLinks(t, q)
		})
	}
}

func testScorings(t *testing.T, qs soq.Questionnaire) {
	seen := map[string]bool{}
	for _, s := range qs.Scorings {
		if _, ok := seen[s.ID]; ok {
			t.Errorf("duplicate scoring ID %s", s.ID)
		}
		seen[s.ID] = true
		if len(s.Questions) == 0 {
			t.Errorf("no questions in scoring %s", s.ID)
		}
		if s.Aggregation != "sum" && s.Aggregation != "mean" {
			t.Errorf("unknown aggregation %s on scoring %s", s.Aggregation, s.ID)
		}
		if len(s.Bands) > 0 {
			testBands(t, s)
		}
	}
}

func testBands(t *testing.T, s *soq.Scoring) {
	seen := map[float64]bool{}
	for _, b := range s.Bands {
		if b.Min != nil && b.Max != nil && *b.Min < *b.Max {
			t.Errorf("min is more than max for band %s of scoring %s", b.Label, s.ID)
		}
		if b.Min != nil {
			if _, ok := seen[*b.Min]; ok {
				t.Errorf("two bands feature value %f for scoring %s", *b.Min, s.ID)
			}
			seen[*b.Min] = true
		}
		if b.Max != nil {
			if _, ok := seen[*b.Max]; ok {
				t.Errorf("two bands feature value %f for scoring %s", *b.Min, s.ID)
			}
			seen[*b.Max] = true
		}
	}
}

func testQuestions(t *testing.T, qs soq.Questionnaire) {
	if len(qs.Questions) == 0 {
		t.Errorf("no questions")
	}
	seen := map[string]bool{}
	questions := map[string]bool{}
	for _, q := range qs.Questions {
		aq := *q
		switch c := aq.(type) {
		case *soq.LikertQuestion:
			if _, ok := seen[c.ID]; ok {
				t.Errorf("duplicate question IDs: %s", c.ID)
			}
			seen[c.ID] = true
			if _, ok := questions[c.Question]; ok {
				t.Errorf("duplicate question: %s", c.Question)
			}
			seen[c.Question] = true
		default:
			t.Errorf("unknown question type")
		}
	}
}

func testLinks(t *testing.T, qs soq.Questionnaire) {
	for _, l := range qs.Links {
		resp, err := http.Get(l.URL)
		if err != nil {
			t.Fatalf("error encountered confirming link %s: %s", l.URL, err.Error())
		}
		resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			t.Errorf("unexpected status code confirming link %s: %d", l.URL, resp.StatusCode)
		}
	}
}
