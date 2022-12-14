// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteAllTokensForUser(ctx context.Context, arg DeleteAllTokensForUserParams) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserEmail(ctx context.Context, email string) (User, error)
	GetUserFromToken(ctx context.Context, arg GetUserFromTokenParams) (User, error)
	InsertSession(ctx context.Context, arg InsertSessionParams) (Session, error)
	InsertToken(ctx context.Context, arg InsertTokenParams) error
}

var _ Querier = (*Queries)(nil)
