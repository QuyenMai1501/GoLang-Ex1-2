package main

import "fmt"

func show(menu map[string]int) {
	fmt.Println("---Menu---")
	for name, price := range menu {
		fmt.Println("Name: ", name)
		fmt.Println("Price: ", price)
	}
	fmt.Println("---------")
}

func add(menu map[string]int, name string, price int) map[string]int  {
	menu[name] = price
	fmt.Println("Added successfully.", name)
	return menu
}

func edit(menu map[string]int, name string, price int) map[string]int {
	menu[name] = price
	fmt.Println("Edited successfully.", name)
	return menu
}

func remove(menu map[string]int, name string) map[string]int {
	delete(menu, name)
	fmt.Println("Deleted successfully.", name)
	return menu
}