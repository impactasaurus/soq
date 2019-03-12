package search

import (
	"github.com/blevesearch/bleve"
	"github.com/impactasaurus/soq"
)

type Engine struct {
	i bleve.Index
	c Cache
}

type Cache interface {
	Questionnaire(id string) (soq.Questionnaire, error)
	QuestionnaireIDs() []string
}

func New(c Cache) (*Engine, error) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		return nil, err
	}

	for _, qID := range c.QuestionnaireIDs() {
		q, err := c.Questionnaire(qID)
		if err != nil {
			return nil, err
		}
		err = index.Index(q.ID, q)
		if err != nil {
			return nil, err
		}
	}

	return &Engine{
		i: index,
		c: c,
	}, nil
}

func (e *Engine) Search(query string, page, limit int) (soq.QuestionnaireList, error) {
	rr, err := e.i.Search(bleve.NewSearchRequest(bleve.NewMatchQuery(query)))
	if err != nil {
		return soq.QuestionnaireList{}, err
	}
	pi := soq.PageInfo{
		Limit:       limit,
		HasNextPage: false,
		Page:        page,
	}
	minIdx := page * limit
	exclusiveMaxIdx := minIdx + limit
	if rr.Hits.Len() <= minIdx {
		return soq.QuestionnaireList{
			PageInfo: pi,
		}, nil
	}
	if rr.Hits.Len() > exclusiveMaxIdx {
		pi.HasNextPage = true
	}
	if rr.Hits.Len() < exclusiveMaxIdx {
		exclusiveMaxIdx = rr.Hits.Len()
	}
	results := make([]*soq.Questionnaire, 0, limit)
	for _, h := range rr.Hits[minIdx:exclusiveMaxIdx] {
		q, err := e.c.Questionnaire(h.ID)
		if err != nil {
			return soq.QuestionnaireList{}, err
		}
		results = append(results, &q)
	}
	return soq.QuestionnaireList{
		PageInfo:       pi,
		Questionnaires: results,
	}, nil
}
