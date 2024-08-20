package cli

import (
	"fmt"
	"go-intensive/internal/service"
	"os"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI {
	return &BookCLI{service: service}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: books search <nome_do_livro>")
		return
	}

	bookName := os.Args[2]

	cli.searchBooks(bookName)
}

func (cli *BookCLI) searchBooks(name string) {
	books, err := cli.service.SearchBooksByName(name)
	if err != nil {
		fmt.Println("Erro ao buscar livros:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("Nenhum livro encontrado com o nome:", name)
		return
	}

	fmt.Printf("Encontrado(s) %d livro(s):\n", len(books))
	for _, book := range books {
		fmt.Printf("ID: %d, Título: %s, Autor: %s, Gênero: %s\n",
			book.ID, book.Title, book.Author, book.Genre)
	}
}
