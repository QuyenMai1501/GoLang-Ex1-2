package main

func main() {
	menu := map[string]int{"Coffee": 30_000, "Tea": 12_000, "Juice": 4_000}
	show(menu)

	for i := range menu {
		print(i, " ")
		menu = add(menu, "Milk", 15_000)
	}

	menu = edit(menu, "Water", 5_000)
	menu = edit(menu, "Matcha Latte", 25_000)
	menu = remove(menu, "Milk")
	show(menu)
}
