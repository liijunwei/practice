package main

import (
	"bufio"
	"os"
)

func readSamples() []string {
	f, err := os.Open("./words.txt")
	boom(err)
	defer f.Close()

	words := make([]string, 0, 466550)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		boom(err)
	}

	return words
}

// steal from: https://gitlab.com/tsoding/trie/-/blob/main/fruits.h?ref_type=heads
var samples = []string{
	"Apple",
	"Apricot",
	"Avocado",
	"Banana",
	"Bilberry",
	"Blackberry",
	"Blackcurrant",
	"Blueberry",
	"Boysenberry",
	"Currant",
	"Cherry",
	"Cherimoya",
	"Chico fruit",
	"Cloudberry",
	"Coconut",
	"Cranberry",
	"Cucumber",
	"Custard apple",
	"Damson",
	"Date",
	"Dragonfruit",
	"Durian",
	"Elderberry",
	"Feijoa",
	"Fig",
	"Goji berry",
	"Gooseberry",
	"Grape",
	"Raisin",
	"Grapefruit",
	"Guava",
	"Honeyberry",
	"Huckleberry",
	"Jabuticaba",
	"Jackfruit",
	"Jambul",
	"Jujube",
	"Juniper berry",
	"Kiwano",
	"Kiwifruit",
	"Kumquat",
	"Lemon",
	"Lime",
	"Loquat",
	"Longan",
	"Lychee",
	"Mango",
	"Mangosteen",
	"Marionberry",
	"Melon",
	"Cantaloupe",
	"Honeydew",
	"Watermelon",
	"Miracle fruit",
	"Mulberry",
	"Nectarine",
	"Nance",
	"Olive",
	"Orange",
	"Blood orange",
	"Clementine",
	"Mandarine",
	"Tangerine",
	"Papaya",
	"Passionfruit",
	"Peach",
	"Pear",
	"Persimmon",
	"Physalis",
	"Plantain",
	"Plum",
	"Prune",
	"Pineapple",
	"Plumcot",
	"Pomegranate",
	"Pomelo",
	"Purple mangosteen",
	"Quince",
	"Raspberry",
	"Salmonberry",
	"Rambutan",
	"Redcurrant",
	"Salal berry",
	"Salak",
	"Satsuma",
	"Soursop",
	"Star fruit",
	"Solanum quitoense",
	"Strawberry",
	"Tamarillo",
	"Tamarind",
	"Ugli fruit",
	"Yuzu",
}
