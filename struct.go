package main

import "fmt"

func main() {
	myEnemy()
}

type School struct {
	number int //init struct (class)
	name string
	friends []string
}

type Enemy struct {
	School //get objects from School, inheritance
	hisName string `required max:"50"`
}

func myFriends() {
	aSchool := School{ //create object from school class named aSchool
		number: 2, //number delete juga gapapa tp tiati not good
		name: "sam",
		friends: []string {
			"lol",
			"imagine",
		},
	}

	fmt.Println(aSchool.name)
}

func myEnemy() {
	aSchool := Enemy{ //create object from Enemy class named aSchool
		School: School{name: "Sam", friends: []string {"lol", "lolz"}}, // input inheritace enemy from School
		hisName: "sheesh",
	}

	fmt.Println(aSchool)
}