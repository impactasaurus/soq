//go:generate gorunpkg github.com/99designs/gqlgen

package api

import (
	"context"

	"github.com/impactasaurus/soq"
)

type resolver struct {
	fetcher QuestionnaireFetcher
}

func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *resolver }

func getPagination(page *int, limit *int) (int, int) {
	p := 0
	l := 10
	if page != nil {
		p = *page
	}
	if limit != nil {
		l = *limit
	}
	return p, l
}

func (r *queryResolver) Questionnaires(ctx context.Context, page *int, limit *int) (soq.QuestionnaireList, error) {
	p, l := getPagination(page, limit)
	return r.fetcher.Questionnaires(p, l)
}
func (r *queryResolver) Questionnaire(ctx context.Context, id string) (soq.Questionnaire, error) {
	return r.fetcher.Questionnaire(id)
}
func (r *queryResolver) Search(ctx context.Context, query string, page *int, limit *int) (soq.QuestionnaireList, error) {
	p, l := getPagination(page, limit)
	return r.fetcher.Search(query, p, l)
}
