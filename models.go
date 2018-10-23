// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package soq_api

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
)

// Band provides context to scores within a certain band of values
type Band struct {
	Min         *float64 `json:"min"`
	Max         *float64 `json:"max"`
	Label       string   `json:"label"`
	Description *string  `json:"description"`
}

// Changes details the changes made in a given version
type Changes struct {
	Version   string   `json:"version"`
	Changes   []string `json:"changes"`
	Timestamp string   `json:"timestamp"`
}

// LikertQuestion is a likert scale question
type LikertQuestion struct {
	ID          string   `json:"id"`
	Question    string   `json:"question"`
	Description *string  `json:"description"`
	Short       *string  `json:"short"`
	Scale       []*Point `json:"scale"`
}

func (LikertQuestion) IsQuestion() {}

// Link captures a URL with some textual context
type Link struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	URL         string  `json:"url"`
}

// PageInfo details the current page and whether pages are available either side
type PageInfo struct {
	HasNextPage bool `json:"hasNextPage"`
	Page        int  `json:"page"`
	Limit       int  `json:"limit"`
}

// Point is a value and a label, used for scale questions
type Point struct {
	Value float64 `json:"value"`
	Label *string `json:"label"`
}

// Question is the basic fields required by questions
type Question interface {
	IsQuestion()
}

// Questionnaire is a collection of questions and scorings
type Questionnaire struct {
	ID           string      `json:"id"`
	Logo         *string     `json:"logo"`
	Name         string      `json:"name"`
	Short        *string     `json:"short"`
	Version      string      `json:"version"`
	Changelog    []Changes   `json:"changelog"`
	Description  *string     `json:"description"`
	Attribution  *string     `json:"attribution"`
	Instructions *string     `json:"instructions"`
	Links        []*Link     `json:"links"`
	Questions    []*Question `json:"questions"`
	Scorings     []*Scoring  `json:"scorings"`
}

// QuestionnaireList is a list of questionnaires along with page information
type QuestionnaireList struct {
	Questionnaires []*Questionnaire `json:"questionnaires"`
	PageInfo       PageInfo         `json:"pageInfo"`
}

// Scoring is an aggregation of questions into a score which can be tracked and evaluated
type Scoring struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description"`
	Questions   []*string   `json:"questions"`
	Aggregation Aggregation `json:"aggregation"`
	Bands       []*Band     `json:"bands"`
}

// Aggregation details the mathematical function used to aggregate questions
type Aggregation string

const (
	// Mean takes the average of the question values
	AggregationMean Aggregation = "MEAN"
	// Sum totals all the questions values
	AggregationSum Aggregation = "SUM"
)

func (e Aggregation) IsValid() bool {
	switch e {
	case AggregationMean, AggregationSum:
		return true
	}
	return false
}

func (e Aggregation) String() string {
	return string(e)
}

func (e *Aggregation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Aggregation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Aggregation", str)
	}
	return nil
}

func (e Aggregation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
