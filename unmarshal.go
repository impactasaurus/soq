package soq_api

import (
	"encoding/json"
	"fmt"
)

func unmarshalQuestion(raw json.RawMessage) (Question, error) {
	typ := struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(raw, &typ); err != nil {
		return nil, err
	}
	switch typ.Type {
	case "likert":
		l := LikertQuestion{}
		if err := json.Unmarshal(raw, &l); err != nil {
			return nil, err
		}
		return &l, nil
	default:
		return nil, fmt.Errorf("unknown question type %s", typ.Type)
	}
}

func (q *Questionnaire) UnmarshalJSON(data []byte) error {
	type Alias Questionnaire
	aux := &struct {
		Questions []json.RawMessage `json:"questions"`
		*Alias
	}{
		Alias: (*Alias)(q),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	outQuestions := make([]*Question, len(aux.Questions))
	for idx, rawQ := range aux.Questions {
		q, err := unmarshalQuestion(rawQ)
		if err != nil {
			return err
		}
		outQuestions[idx] = &q
	}
	q.Questions = outQuestions
	return nil
}
