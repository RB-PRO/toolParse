package main

type cost struct {
	citilink int
}
type link struct {
	citilink string
}
type product struct {
	articl string // Артикл
	title  string // Название
	costs  cost
	links  link
}

func main() {
	sitilinkParse("%D1%82%D0%B5%D0%BB%D0%B5%D1%84%D0%BE%D0%BD")
}
