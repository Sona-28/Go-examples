package main

import "fmt"

func handlePanic() {
	if a := recover(); a != nil {
		fmt.Println("3. Recover from ", a)

	}
	fmt.Println("4.I am stopping here by closing all the connections")
}

func entry(language *string, name *string) {
	fmt.Println("1= Entry Regular")
	defer handlePanic()
	defer doCleanUp()
	if language == nil {
		panic("Error: Language can't be nil")
	}
	if name == nil {
		panic("Panic - because of name")

	}
	fmt.Println("language", language)
	fmt.Println("name", name)
}

func doCleanUp() {
	fmt.Println("2. defer-entry-cleanup")
}

func main() {
	defer fmt.Println("7. i am the first defer")
	defer fmt.Println("6.defer - in main")
	lang := "go lang"
	entry(&lang, nil)
	fmt.Println("5. main-normal", lang)
}
