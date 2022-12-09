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
	name := "Двигатель BRAIT BR465P"

	// https://www.ozon.ru/search/?text=%D0%94%D0%B2%D0%B8%D0%B3%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%20BRAIT%20BR465P&from_global=true
	// %D0%94%D0%B2%D0%B8%D0%B3%D0%B0%D1%82%D0%B5%D0%BB%D1%8C+BRAIT+BR465P

	//sitilinkParse("%D1%82%D0%B5%D0%BB%D0%B5%D1%84%D0%BE%D0%BD")
	ozonParse(name)
}
