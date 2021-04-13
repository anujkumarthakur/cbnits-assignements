package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"book/graph/generated"
	"book/graph/logic"
	"book/graph/model"
	"context"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	newbook, err := logic.CreateNewBook(input)
	if err != nil {
		return nil, err
	}
	if newbook != nil {
		return newbook, nil
	}
	return nil, nil
}

func (r *mutationResolver) GetSingleBook(ctx context.Context, input model.BookID) (*model.Book, error) {
	book, err := logic.GetBook(input)
	if err != nil {
		return nil, err
	}
	if book != nil {
		return book, nil
	}
	return nil, err
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	//panic(fmt.Errorf("not implemented"))
	books, err := logic.GetBooks()
	if err != nil {
		return nil, err
	}
	if books != nil {
		return books, nil
	}
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
