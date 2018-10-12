//go:generate gorunpkg github.com/99designs/gqlgen

package api

import (
	"context"

	soq "github.com/impactasaurus/soq-api"
)

type resolver struct {
	fetcher QuestionnaireFetcher
}

func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *resolver }

func (r *queryResolver) Questionnaires(ctx context.Context, page *int, limit *int) (soq.QuestionnaireList, error) {
	p := 0
	l := 10
	if page != nil {
		p = *page
	}
	if limit != nil {
		l = *limit
	}
	return r.fetcher.Questionnaires(p, l)
}
func (r *queryResolver) Questionnaire(ctx context.Context, id string) (soq.Questionnaire, error) {
	return r.fetcher.Questionnaire(id)
}
