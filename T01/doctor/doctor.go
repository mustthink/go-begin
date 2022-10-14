package main

import (
	"bufio"
	"fmt"
	"os"
)

type answer struct {
	awer   string
	nextid int
}

type question struct {
	id   int
	qion string
	anss []answer
}

//Усі питання мають однакову структуру, тому їх можна легко доповнювати й збільшувати ланцюжки діалогу або доповнювати відповіді до вже існуючих питань

var quest = []question{{
	qion: "Good afternoon, how are you feeling?",
	anss: []answer{{"Good", 1}, {"I am sick", 2}},
}, {
	qion: "All right, take care :)",
	anss: []answer{{"", 999}},
}, {
	qion: "Tell us about your illness",
	anss: []answer{{"", 3}},
},
	{
		qion: "We will contact you",
		anss: []answer{{"", 999}},
	},
}

func getQ(id int) string {
	if id < len(quest) {
		return quest[id].qion
	} else {
		return "Error question"
	}
}

func getNextID(q question, ans string) int {
	for _, v := range q.anss {
		if v.awer == ans || v.awer == "" {
			return v.nextid
		}
	}
	return 999
}

func asking() {
	id := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("-", getQ(id))
	for scanner.Scan() {
		id = getNextID(quest[id], scanner.Text())
		if id != 999 {
			fmt.Println("-", getQ(id))
		} else {
			break
		}
	}
}

func main() {
	asking()
	fmt.Println("The end!")
}
