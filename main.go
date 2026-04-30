package main

type Mood string

const (
	happy Mood = "веселое"
	calm  Mood = "спокойное"
	sad   Mood = "грустное"
)

type Movie struct {
	Name   string
	Year   int
	Rating float64
	Genres []string
	Mood   Mood
}

var movies []Movie = []Movie{
	{"Шрэк", 2001, 7.6, []string{"мультфильм", "фэнтези", "мелодрама", "комедия", "приключения", "семейный"}, happy},
	{"Зеленая миля", 1999, 8.0, []string{"драма", "фэнтези", "криминал"}, sad},
	{"Унесенные призраками", 2001, 7.7, []string{"аниме", "мультфильм", "фэнтези", "приключения", "семейный"}, calm},
}

type Music struct {
	Name string
	Mood Mood
}

var music []Music = []Music{
	{"808s & Heartbreak (Kanye West)", sad},
	{"Платина (Платина)", happy},
	{"Circles (Mac Miller)", calm},
}

var jokes []string = []string{
	"Анекдот про Штрилица 1",
	"Анекдот про Штрилица 2",
	"Анекдот про Штрилица 3",
}

type MerchType string

const (
	sweater  MerchType = "кофта"
	trousers MerchType = "брюки"
	tShirt   MerchType = "футболка"
	cap      MerchType = "кепка"
)

type Merch struct {
	Type  MerchType
	Price int64
}

var cart map[Merch]int = make(map[Merch]int)

func main() {
}
