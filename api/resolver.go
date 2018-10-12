//go:generate gorunpkg github.com/99designs/gqlgen

package api

import (
	context "context"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Questionnaires(ctx context.Context, page *int, limit *int) (QuestionnaireList, error) {
	return QuestionnaireList{}, nil
}
func (r *queryResolver) Questionnaire(ctx context.Context, id string) (Questionnaire, error) {
	return Questionnaire{}, nil
}
