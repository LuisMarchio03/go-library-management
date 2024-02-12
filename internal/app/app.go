// internal/app/app.go
package app

import (
    "fmt"
    "library-api/internal/pkg/book"
)

// Run is the entry point of the application
func Run() {
    //TODO: 01. Importar autores do CSV
    err := importAuthors("authors.csv")
    if err != nil {
        fmt.Println("Erro ao importar autores:", err)
        return
    }

    //TODO: 02. Inicializar o repositório de livros
    bookRepository := &book.InMemoryBookRepository{}

    //TODO: 03. Criar, ler, atualizar, excluir livros usando bookRepository...
}
