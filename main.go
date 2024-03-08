package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func isKeyWork(word string) bool {
	keywords := []string{"add", "delete", "done", "-d"}
	for _, i := range keywords {
		if word == i {
			return true
		}
	}
	return false
}

func getKeywordsMap(words []string) map[string]int {
	res := map[string]int{"add": -1, "delete": -1, "done": -1, "-d": -1}
	for i, v := range words {
		if isKeyWork(v) {
			res[v] = i
		}
	}
	return res
}

// func getFlag(inputs []string) (string, int) {
// 	for ind, i := range inputs {
// 		if strings.Contains(i, "-") {
// 			return i, ind
// 		}
// 	}
// 	return "", -1
// }

func processInput(inputs []string, db *gorm.DB) {
	keywordsMap := getKeywordsMap(inputs)
	switch inputs[1] {
	case "add":
		title := strings.Join(inputs[2:], " ")
		desc := ""
		if keywordsMap["-d"] > -1 {
			desc = strings.Join(inputs[keywordsMap["-d"]+1:], " ")
			title = strings.Join(inputs[2:keywordsMap["-d"]], " ")
		}
		create(db, Todo{Title: title, Description: &desc})
	case "done":
		id, err := strconv.Atoi(inputs[2])
		if err != nil {
			return
		}
		item := get(db, id)
		item.IsCompleted = !item.IsCompleted
		update(db, item)
	}

}

func displayItems(todos []Todo) {
	for _, i := range todos {
		done := "TODO"
		if i.IsCompleted {
			done = "DONE"
		}
		if *i.Description != "" {
			fmt.Printf("%d) %s (%s) - %s\n", i.ID, i.Title, *i.Description, done)
		} else {
			fmt.Printf("%d) %s - %s\n", i.ID, i.Title, done)
		}
	}
}

func main() {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 2 {
		processInput(os.Args, db)
	} else if len(os.Args) == 2 {
		panic("Wrong number arguments")

	}
	todos := getAll(db)
	displayItems(todos)
}
