package main

import (
	"fmt"
	"math/rand"
	"slices"
)

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
	Name   string
	Artist string
	Year   int
	Genres []string
	Mood   Mood
}

var music []Music = []Music{
	{"808s & Heartbreak", "Kanye West", 2008, []string{"электропоп", "синти-поп", "RnB"}, sad},
	{"Платина", "Платина", 2024, []string{"трэп", "рейдж", "хип-хоп"}, happy},
	{"Pink Floyd", "Wish You Were Here", 1975, []string{"рок", "арт-рок", "экспериментальный рок"}, calm},
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

type MerchID int

type Merch struct {
	Type  MerchType
	Price float64
}

var merchCatalog map[MerchID]Merch = map[MerchID]Merch{
	1: {Type: sweater, Price: 40.99},
	2: {Type: trousers, Price: 50.33},
	3: {Type: tShirt, Price: 20.49},
	4: {Type: cap, Price: 10.09},
}

var merchCart map[MerchID]int = make(map[MerchID]int, len(merchCatalog))

func login() string {
	fmt.Print("Привет, как тебя зовут?\n>>> ")
	var name string
	fmt.Scanln(&name)
	return name
}

func menu() {
	fmt.Println("\nМеню\n1 - фильм по жанру\n2 - фильм по нстроению\n3 - музыка по жанру\n4 - музыка по нстроению\n5 - анекдот\n6 - каталог мерча\n7 - заказать мерч\n0 - выход")
	fmt.Print(">>> ")
}

func tellJoke() {
	rndIndex := rand.Intn(len(jokes))
	fmt.Printf("\n%s\n", jokes[rndIndex])
}

func printMovieInfo(movie Movie) {
	fmt.Printf("\n%s (%d) %.2f\nЖанры: ", movie.Name, movie.Year, movie.Rating)
	for i, g := range movie.Genres {
		if i == len(movie.Genres)-1 {
			fmt.Printf("%s\n", g)
		} else {
			fmt.Printf("%s, ", g)
		}
	}
	fmt.Printf("Настроение: %s\n", movie.Mood)
}

func printMusicInfo(music Music) {
	fmt.Printf("\n\"%s\" (%s) %d\nЖанры: ", music.Name, music.Artist, music.Year)
	for i, g := range music.Genres {
		if i == len(music.Genres)-1 {
			fmt.Printf("%s\n", g)
		} else {
			fmt.Printf("%s, ", g)
		}
	}
	fmt.Printf("Настроение: %s\n", music.Mood)
}

func recommendMovieByGenre() {
	fmt.Print("\nВведите жанр фильма\n>>> ")
	var genre string
	fmt.Scanln(&genre)
	filtered := make([]Movie, 0, len(movies))
	for _, m := range movies {
		if slices.Contains(m.Genres, genre) {
			filtered = append(filtered, m)
		}
	}
	fLen := len(filtered)
	if fLen == 0 {
		fmt.Printf("Фильмов в жанре \"%s\" не найдено\n\n", genre)
		return
	}
	rndIndex := rand.Intn(fLen)
	printMovieInfo(filtered[rndIndex])
}

func recommendMovieByMood() {
	fmt.Print("\nВведите ваше настроение\n>>> ")
	var mood Mood
	fmt.Scanln(&mood)
	filtered := make([]Movie, 0, len(movies))
	for _, m := range movies {
		if m.Mood == mood {
			filtered = append(filtered, m)
		}
	}
	fLen := len(filtered)
	if fLen == 0 {
		fmt.Printf("Фильмов под настроение \"%s\" не найдено\n\n", mood)
		return
	}
	rndIndex := rand.Intn(fLen)
	printMovieInfo(filtered[rndIndex])
}

func recommendMusicByGenre() {
	fmt.Print("\nВведите жанр музыки\n>>> ")
	var genre string
	fmt.Scanln(&genre)
	filtered := make([]Music, 0, len(music))
	for _, m := range music {
		if slices.Contains(m.Genres, genre) {
			filtered = append(filtered, m)
		}
	}
	fLen := len(filtered)
	if fLen == 0 {
		fmt.Printf("Музыки в жанре \"%s\" не найдено\n\n", genre)
		return
	}
	rndIndex := rand.Intn(fLen)
	printMusicInfo(filtered[rndIndex])
}

func recommendMusicByMood() {
	fmt.Print("\nВведите ваше настроение\n>>> ")
	var mood Mood
	fmt.Scanln(&mood)
	filtered := make([]Music, 0, len(music))
	for _, m := range music {
		if m.Mood == mood {
			filtered = append(filtered, m)
		}
	}
	fLen := len(filtered)
	if fLen == 0 {
		fmt.Printf("Музыки под настроение \"%s\" не найдено\n\n", mood)
		return
	}
	rndIndex := rand.Intn(fLen)
	printMusicInfo(filtered[rndIndex])
}

func showMerchCatalog() {
	fmt.Println("\nКаталог")
	for id, m := range merchCatalog {
		fmt.Printf("id: %d | тип: %s | цена: %.2f\n", id, m.Type, m.Price)
	}
}

func showCart() {
	var total float64 = 0
	fmt.Println("\nКорзина")
	for id, n := range merchCart {
		m := merchCatalog[id]
		cost := m.Price * float64(n)
		total += cost
		fmt.Printf("id: %d | тип: %s | цена: %.2f у.е. | кол-во: %d | стоимость: %.2f\n", id, m.Type, m.Price, n, cost)
	}
	fmt.Printf("Итого: %.2f\n", total)
}

func addMerchToCart() {
	fmt.Print("\nВведите ID мерча\n>>> ")
	var id MerchID
	fmt.Scanln(&id)
	if merch, ok := merchCatalog[id]; ok {
		fmt.Printf("Введите количество \"%s\"\n>>> ", merch.Type)
		var n int
		fmt.Scanln(&n)
		merchCart[id] += n
	} else {
		fmt.Printf("Мерч с ID \"%d\" не найден\n", id)
	}
}

func orderMerch() {
	showMerchCatalog()
menu:
	for {
		fmt.Print("\nМеню\n1 - добавить мерч в корзину\n2 - оформить заказ\n0 - выйти\n>>> ")
		var opt string
		fmt.Scanln(&opt)
		switch opt {
		case "1":
			addMerchToCart()
			showCart()
		case "2":
			showCart()
			accept := "ПОДТВЕРДИТЬ"
			fmt.Printf("\nВведите \"%s\"\n>>> ", accept)
			fmt.Scanln(&opt)
			if opt == accept {
				fmt.Println("\nЗаказ успешно оформлен")
				emptyCart()
				break menu
			}
		case "0":
			fmt.Println("Заказ отменен")
			emptyCart()
			break menu
		default:
			fmt.Println("Нет такой опции")
		}
	}
}

func emptyCart() {
	merchCart = make(map[MerchID]int, len(merchCatalog))
	fmt.Println("Корзина очищена")
}

func menuSwitcher(opt string) {
}

func main() {
	name := login()
	fmt.Printf("Welcome, %s!\n", name)
	var opt string
menu:
	for {
		menu()
		fmt.Scanln(&opt)
		switch opt {
		case "1":
			recommendMovieByGenre()
		case "2":
			recommendMovieByMood()
		case "3":
			recommendMusicByGenre()
		case "4":
			recommendMusicByMood()
		case "5":
			tellJoke()
		case "6":
			showMerchCatalog()
		case "7":
			orderMerch()
		case "0":
			break menu
		default:
			fmt.Println("Нет такой опции")
		}
	}
}

// это просто огромная куча дублированного кода (говна), да.
