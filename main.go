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
	{"808s & Heartbreak", "Kanye West", 2008, []string{"электропоп"}, sad},
	{"Платина", "Платина", 2024, []string{"трэп"}, happy},
	{"Pink Floyd", "Wish You Were Here", 1975, []string{"рок"}, calm},
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

var catalog []MerchType = []MerchType{sweater, trousers, tShirt, cap}
var cart map[Merch]int = make(map[Merch]int)

func login() string {
	fmt.Print("Привет, как тебя зовут?\n>>> ")
	var name string
	fmt.Scanln(&name)
	return name
}

func menu() {
	fmt.Println("Меню\n1 - рекомендация фильма по жанру\n2 - рекомендация фильма по нстроению\n3 - рекомендация музыки по жанру\n4 - рекомендация музыки по нстроению\n5 - рассказать анекдот\n6 - вывести каталог мерча\n7 - заказать мерч\n0 - выход")
	fmt.Print(">>> ")
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
	fmt.Printf("Настроение: %s\n\n", movie.Mood)
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
	fmt.Printf("Настроение: %s\n\n", music.Mood)
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

func switcher(opt string) {
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
		break
	case "6":
		break
	case "7":
		break
	case "0":
		break
	}
}

func main() {
	name := login()
	fmt.Printf("Welcome, %s!\n", name)
	var opt string
	for {
		menu()
		fmt.Scanln(&opt)
		switcher(opt)
	}
}
