package bootstrap

import (
	"database/sql"
	"mellow/repositories"
	"mellow/repositories/repoimpl"
)

type Repositories struct {
	UserRepository  repositories.UserRepository
	AuthRepository  repositories.AuthRepository
	PostRepository  repositories.PostRepository
	GroupRepository repositories.GroupRepository
	CommentRepository repositories.CommentRepository
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository:  repoimpl.NewUserRepository(db),
		AuthRepository:  repoimpl.NewAuthRepository(db),
		PostRepository:  repoimpl.NewPostRepository(db),
		GroupRepository: repoimpl.NewGroupRepository(db),
		CommentRepository: repoimpl.NewCommentRepository(db),
	}
}
