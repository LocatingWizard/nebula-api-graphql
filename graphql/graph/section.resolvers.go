package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"

	"github.com/LocatingWizard/nebula_api_graphql/graph/model"
)

// CourseReference is the resolver for the course_reference field.
func (r *sectionResolver) CourseReference(ctx context.Context, obj *model.Section) (*model.Course, error) {
	id := obj.CourseReference
	course, err := r.Query().CourseByID(ctx, id.Hex())
	if err != nil {
		return nil, err
	}
	return course, nil
}

// Professors is the resolver for the professors field.
func (r *sectionResolver) Professors(ctx context.Context, obj *model.Section) ([]*model.Professor, error) {
	var out []*model.Professor
	prof_ids := obj.Professors
	for _, id := range prof_ids {
		prof, err := r.Query().ProfessorByID(ctx, id.Hex())
		if err != nil {
			return nil, err
		}
		out = append(out, prof)
	}
	return out, nil
}

// Section returns SectionResolver implementation.
func (r *Resolver) Section() SectionResolver { return &sectionResolver{r} }

type sectionResolver struct{ *Resolver }