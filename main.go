package main

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

// func getFlag(inputs []string) string {
// 	flags := []string{}
// 	for _, i := range inputs {
// 		if strings.Contains(i, "-") {
// 			flags = append(flags, i)
// 		}
// 	}
// 	if len(flags) > 0 {
// 		return flags[0]
// 	}
// 	return ""
// }

func processInput(inputs []string, db *gorm.DB) {

	switch inputs[1] {
	case "add":
		create(db, Todo{Title: inputs[2], Description: &inputs[2]})

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
	} else if len(os.Args) < 2 {
		todos := getAll(db)
		for _, i := range todos {
			done := ""
			if i.IsCompleted {
				done = "DONE"
			}
			fmt.Printf("%d) %s - %s", i.ID, i.Title, done)
		}
	} else {
		panic("Wrong number arguments")
	}
}
