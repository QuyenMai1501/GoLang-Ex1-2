package main

func main() {
	var names [100]string = [100]string{"Coca-cola", "Pepsi", "Fanta"}
	var prices [100]int = [100]int{10_000, 12_000, 8_000}
	var length int = 3

	show(names, prices, length)
	names, prices, length = add(names, prices, length, "Sprite", 9_000)
	names, prices, length = add(names, prices, length, "Sprite", 9_000)
	names, prices, length = add(names, prices, length, "Sprite", 9_000)
	names, prices, length = add(names, prices, length, "Sprite", 9_000)

	names, prices = edit(names, prices, 3, "Sting", 11_000)
	names, prices = edit(names, prices, 4, "Iced Tea", 5_000)
	names, prices, length = delete(names, prices, length, 6)
	show(names, prices, length)
}