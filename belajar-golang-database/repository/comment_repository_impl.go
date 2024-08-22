package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	var comment entity.Comment
	query := "SELECT id, email, comment FROM comments WHERE id = ?"
	err := repository.DB.QueryRowContext(ctx, query, id).Scan(&comment.Id, &comment.Email, &comment.Comment)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	var comments []entity.Comment
	query := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return comments, err
	}

	defer rows.Close()

	for rows.Next() {
		var comment entity.Comment
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comments, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
