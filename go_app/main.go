// Created by: Jakub Malhotra
// Created on: April 2023
//
// This program calculates the volume of a pyramid

package main

import "fmt"

func main() {
	// This function calculates the volume of a pyramid
	var length int
	var width int
	var height int
	var volume int

	fmt.Println("This program finds the volume of a pyramid.")
	fmt.Println()
	fmt.Print("Enter the length (cm): ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width (cm): ")
	fmt.Scanln(&width)
	fmt.Print("Enter the height (cm): ")
	fmt.Scanln(&height)

	volume = (length * height * width) / 3

	fmt.Println()
	fmt.Println("The volume is:", volume, "cm³.")
	fmt.Println("\nDone.")
}
