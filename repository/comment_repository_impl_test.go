package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	learn_golang_database_mysql "learn-golang-database-mysql"
	"learn-golang-database-mysql/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(learn_golang_database_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repositorytest@example.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(learn_golang_database_mysql.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 47)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(learn_golang_database_mysql.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
