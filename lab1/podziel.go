package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

// main kolejno:
// + czyta wyrazy z pliku 'słowa.txt'
// + na wszystkie możliwe sposobu dzieli każdy wyraz na 2 części
// + jeśli obie częsci są wyrazam, przyznaje tym częściom po 1 punkcie
// + sortuje te części wg malejącej kolejności punktów
// + wypisuje takie części, które mają więcej niż 0 punktów i ich punkty
// + wypisuje takie części, które mają tyle samo punktów, w porządku leksykograficznym
func main() {
	wordList := ReadLines("slowa.txt")
	wordSet := map[string]bool{}
	for _, w := range wordList {
		wordSet[w] = true
	}
	wordCounter := map[string]int{}
	for _, w := range wordList {
		for _, parts := range Split(w) {
			IncrementIfBothIn(parts, wordSet, &wordCounter)
		}
	}
	pairs := Sort(wordCounter)
	for _, p := range pairs {
		fmt.Printf("%d %s\n", p.points, p.word)
	}
}

// Read wczytuje wiersze z pliku o nazwie 'filename' i zwraca je w
// wycinku tablicy łańcuchów
func ReadLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

// Split na wszystie możliwe sposoby dzieli wyraz 'word' na 2
// niepuste łańcuchy i zwraca te 2 łańcuchy w wycinku tablicy
// wycinków tablicy łańcuchów
func Split(word string) [][]string {
	r := [][]string{}
	for i := 1; i < len(word); i++ {
		r = append(r, []string{word[:i], word[i:]})
	}
	return r
}

// IncrementIfBothIn zwiększa licznik 'counter' przy tych łańcuchach z
// wycinka 'parts', które należą do zbioru łańcuchów 'set'
func IncrementIfBothIn(parts []string, set map[string]bool, counter *map[string]int) {
	if set[parts[0]] && set[parts[1]] {
		(*counter)[parts[0]]++
		(*counter)[parts[1]]++
	}
}

type Pair struct {
	word   string
	points int
}

// Sort sortuje wycinek par (word, points). Gdy 2 pary mają różną
// liczbę punktów, ta para, która ma więcej punktów, poprzedza tę
// parę, który ma mniej punktów. Gdy 2 pary mają tyle samo punktów, ta
// para, której pole 'word' jest wcześniej w porządku leksykograficznym, poprzedza
// drugą parę
func Sort(counter map[string]int) []Pair {
	r := []Pair{}
	for w, p := range counter {
		r = append(r, Pair{w, p})
	}
	slices.SortFunc(r, func(a, b Pair) int {
		if n := b.points - a.points; n != 0 {
			return n
		}
		return cmp.Compare(a.word, b.word)
	})
	return r
}