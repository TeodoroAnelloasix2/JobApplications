package book

import "fmt"

type Printeable interface {
	PrintInfo()
}

func Print(p Printeable) {
	p.PrintInfo()
}

// Estructura con letras mayuscula en sus caracteristicas es publico
type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}

// Simular constructor
func NewBook(title, author string, pages int) *Book {

	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

//Estructura con letras  minusculas en sus caracteristicas es privado

type libroPrivado struct {
	titulo  string
	autor   string
	paginas int
}

func NewPrivateBook(t, a string, p int) *libroPrivado {

	return &libroPrivado{
		titulo:  t,
		autor:   a,
		paginas: p,
	}
}

// Set method
func (pb *libroPrivado) SetTitle(title string) {
	pb.titulo = title
}

// Get method
func (pb *libroPrivado) GetTitle() string {
	return pb.titulo
}

//Composicion para simular herencia

type TexBook struct {
	B         Book
	editorial string
	level     string
}

func NewTexBook(t, a, e, l string, pages int) *TexBook {
	return &TexBook{
		B:         *NewBook(t, a, pages),
		editorial: e,
		level:     l,
	}
}
func (b *TexBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nLevel: %s", b.B.Title, b.B.Author, b.B.Pages, b.editorial, b.level)
}
