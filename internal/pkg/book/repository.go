// internal/pkg/book/repository.go
package book

// BookRepository is the interface that wraps the basic book repository operations
type BookRepository interface {
    Create(book *Book) error
		ReadAll() ([]*Book, error)
    Read(id int) (*Book, error)
    Update(book *Book) error
    Delete(id int) error
}
