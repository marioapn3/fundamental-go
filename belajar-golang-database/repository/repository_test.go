package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	db := belajar_golang_database.GetConnection()
	defer db.Close()

	commentRepo := NewCommentRepository(db)

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "ctx@gmail.com",
		Comment: "ini adalah komentar",
	}

	result, err := commentRepo.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("Comment ID:", result.Id)
}

func TestFindById(t *testing.T) {
	db := belajar_golang_database.GetConnection()
	defer db.Close()

	commentRepo := NewCommentRepository(db)

	ctx := context.Background()
	comment, err := commentRepo.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Comment:")
	fmt.Println("ID:", comment.Id)
	fmt.Println("Email:", comment.Email)
	fmt.Println("Comment:", comment.Comment)
}

func TestFindAll(t *testing.T) {
	db := belajar_golang_database.GetConnection()
	defer db.Close()

	commentRepo := NewCommentRepository(db)

	ctx := context.Background()
	comments, err := commentRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println("Comment:")
		fmt.Println("ID:", comment.Id)
		fmt.Println("Email:", comment.Email)
		fmt.Println("Comment:", comment.Comment)
	}
}
