package main

import "fmt"

func show(names []string, prices []int) {
	fmt.Println("---Menu---")
	for i := range names {
		fmt.Print(i + 1, ". ")
		fmt.Print(names[i], ": ")
		fmt.Println(prices[i])
	}
	fmt.Println("---------")
}

func add(names []string, prices []int, name string, price int) ([]string, []int) {
	names = append(names, name)
	prices = append(prices, price)
	fmt.Println("Added successfully", name)
	return names, prices
}

func edit(names []string, prices []int, index int, name string, price int) ([]string, []int)  {
	names[index] = name
	prices[index] = price
	fmt.Println("Updated successfully", name)
	return names, prices
}

func delete(names []string, prices []int, index int) ([]string, []int) {
	deleteName := names[index]
	names = append(names[:index], names[index+1:]...)
	prices = append(prices[:index], prices[index+1:]...)
	fmt.Println("Deleted successfully", deleteName)
	return names, prices
}