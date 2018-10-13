package cache

import (
	"errors"

	"sort"

	soq "github.com/impactasaurus/soq-api"
)

type Cache struct {
	byID   map[string]soq.Questionnaire
	byName []*soq.Questionnaire
}

func New(qq []soq.Questionnaire) (*Cache, error) {
	byName := make([]*soq.Questionnaire, len(qq))
	byID := map[string]soq.Questionnaire{}
	for idx := range qq {
		byID[qq[idx].ID] = qq[idx]
		byName[idx] = &qq[idx]
	}
	sort.Slice(byName, func(i, j int) bool {
		return byName[i].Name < byName[j].Name
	})
	return &Cache{
		byID:   byID,
		byName: byName,
	}, nil
}

func (c *Cache) Questionnaire(id string) (soq.Questionnaire, error) {
	q, ok := c.byID[id]
	if !ok {
		return soq.Questionnaire{}, errors.New("could not find questionnaire")
	}
	return q, nil
}

func (c *Cache) Questionnaires(page, limit int) (soq.QuestionnaireList, error) {
	offset := page * limit
	if offset >= len(c.byName) {
		return soq.QuestionnaireList{
			PageInfo: soq.PageInfo{
				Limit:       limit,
				HasNextPage: false,
				Page:        page,
			},
			Questionnaires: []*soq.Questionnaire{},
		}, nil
	}
	end := offset + limit
	if end > len(c.byName) {
		end = len(c.byName)
	}
	more := end != len(c.byName)
	return soq.QuestionnaireList{
		PageInfo: soq.PageInfo{
			Limit:       limit,
			HasNextPage: more,
			Page:        page,
		},
		Questionnaires: c.byName[offset:end],
	}, nil
}
