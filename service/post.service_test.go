package service

import (
	"testing"

	"github.com/renatospaka/golang-rest-api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func(mock *mockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func(mock *mockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "B"}
	testService := NewPostService(nil)

	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error())
}

func TestValidateEmptyPostText(t *testing.T) {
	post := entity.Post{ID: 1, Title: "A", Text: ""}
	testService := NewPostService(nil)

	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "The post text is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	var id int64 = 1 
	mockRepo := new(mockRepository)
	post := entity.Post{ID: id, Title: "A", Text: "B"}

	//setup the expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)
	testService := NewPostService(mockRepo)

	//mock assertion: Behavioral
	result, _ := testService.FindAll()
	mockRepo.AssertExpectations(t)

	//data assertion
	assert.Equal(t, id, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mockRepository)
	post := entity.Post{Title: "A", Text: "B"}

	//setup the expectations
	mockRepo.On("Save").Return(&post, nil)
	testService := NewPostService(mockRepo)

	//mock assertion: Behavioral
	result, err := testService.Create(&post)
	mockRepo.AssertExpectations(t)

	//data assertion
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}
