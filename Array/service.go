package main

import "fmt"

func show(names [100]string, prices [100]int, length int) {
	fmt.Println("---Menu---")
	for i := range length {
		fmt.Print(i + 1, ". ")
		fmt.Print(names[i], ": ")
		fmt.Println(prices[i])
	}
	fmt.Println("---------")
}

func add(names [100]string, prices [100]int, length int, name string, price int) ([100]string, [100]int, int) {
	names[length] = name
	prices[length] = price
	length++
	fmt.Println("Added successfuly", name)
	return names, prices, length
}

func edit(names [100]string, prices [100]int, index int, name string, price int) ([100]string, [100]int) {
	names[index] = name
	prices[index] = price
	index++
	fmt.Println("Updated successfuly", name)
	return names, prices
}

func delete(names [100]string, prices [100]int, length int, index int) ([100]string, [100]int, int) {
	deleteName := names[index]
	for i := index; i < length-1; i++ {
		names[i] = names[i+1]
		prices[i] = prices[i+1]
	}
	length--
	fmt.Println("Deleted successfully", deleteName)
	return names, prices, length
}