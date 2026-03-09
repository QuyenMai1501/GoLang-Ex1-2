package main

func main() {
	var n [3]string = [3]string{"Coke", "Pepsi", "Fanta"}
	var p [3]int = [3]int{10_000, 12_000, 8_000}

	var names = n[0:3]
	var prices = p[0:3]

	show(names, prices)

	for i := range 4 {
		print(i, " ")
		names, prices = add(names, prices, "Sting", 11_000)
	}

	names, prices = edit(names, prices, 3, "Sprite", 13_000)
	names, prices = edit(names, prices, 3, "C2", 9_000)
	names, prices = delete(names, prices, 6)
	show(names, prices)
}
