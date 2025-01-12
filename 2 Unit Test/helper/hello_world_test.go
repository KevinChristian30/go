package helper

import (
	"fmt"
	"main/entity"
	"main/repository"
	"main/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("This is run before the whole package is ran")
	m.Run()
	fmt.Println("This is run after the whole package is ran")
}

func TestHelloWorld(t *testing.T) {
	// If fail, it will execute t.fail()
	assert.Equal(t, HelloWorld("Kevin"), "Hello, Kevin")

	// If fail, it will execute t.failNow() (stops function execution)
	require.Equal(t, HelloWorld("Kevin"), "Hello, Kevin")
}

func TestHelloWorldSkip(t *testing.T) {
	t.Skip()

	require.Equal(t, HelloWorld("Kevin"), "Hello, Kevin")
}

func TestWithSubtest(t *testing.T) {
	t.Run("Name", func(t *testing.T) {
		assert.Equal(t, HelloWorld("Kevin"), "Hello, Kevin")
	})
}

func TestWithMockNil(t *testing.T) {
	categoryRepository := &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
	categoryService := service.CategoryService{Repository: categoryRepository}

	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestWithMockNotNil(t *testing.T) {
	categoryRepository := &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
	categoryService := service.CategoryService{Repository: categoryRepository}

	result := entity.Category{Id: "1", Name: "Gadget"}
	categoryRepository.Mock.On("FindById", "1").Return(result)

	category, err := categoryService.Get("1")
	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, result.Id, category.Id)
	assert.Equal(t, result.Name, category.Name)
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kevin")
	}
}
