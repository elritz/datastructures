package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	// Binary Search
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	// index := binarySearch(a, 11)
	// fmt.Println("index =>", index)

	// new  - only allocates memory - no initialization of memory
	// make - allocates and initializes memory
	// This means non zero storage

	// Maps

	// score := make(map[string]int)

	// score["paul"] = 22
	// score["christian"] = 25
	// score["vicky"] = 23
	// score["sam"] = 27

	// fmt.Println(score)
	// getChiritasScore := score["christian"]

	// fmt.Println(getChiritasScore)
	// // provide the map - key to delete
	// delete(score, "christian")

	// for k, v := range score {
	// 	// fmt.Println(k, v)
	// 	fmt.Printf("Score of %v is %v \n", k, v)
	// }

	// Structs
	// dont have a concept of inheritance

	// type User struct {
	// 	Name  string // - how is default value -
	// 	Email string
	// 	Age   int
	// }

	// rob := User{Name: "Rob"}

	// fmt.Printf("%v\n", rob)
	// fmt.Printf("%+v\n", rob)

	// // When you are using new - you are allocating memory and this is a pointer a referfence to the memory
	// var sam = new(User)
	// sam.Name = "Sam"
	// sam.Email = "same@gmail.com"
	// sam.Age = 23

	// fmt.Printf("%+v\n", sam)

	// var tobby = &User{Name: "Tobby", Email: "tobby@gmail.com", Age: 23}
	// fmt.Printf("%+v\n", tobby)

	// Loops

	// start := 1
	// things := []string{"ma", "pa", "ta", "da"}

	// fmt.Println(things)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i + start)
	// }

	// // this is a loop
	// for i := 0; i < len(things); i++ {
	// 	fmt.Println(things[i])
	// }

	// // this is a loop
	// for i, thing := range things {

	// 	fmt.Println(i, thing)

	// }

	// for start < 100 {
	// 	start += start
	// 	if start == 32 {
	// 		goto lcolabel
	// 	}

	// lcolabel:
	// 	fmt.Println("GOne to the label: ", start)
	// 	break
	// 	// continue
	// 	// fmt.Println("Start is now: ", start)
	// }

	// Rater and timestamping

	var name string
	var userRating string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your full name: ")

	// read the string till we hit a new line
	name, _ = reader.ReadString('\n')

	// take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Dosa center between 1 and 5: ")

	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	fmt.Printf("Hello %v, \n Thanks for rating our dosa center with %v star rating. \n\n Your rating was recorded in our system at %v", name, mynumRating, time.Now().Format(time.Stamp))

	// Frontend

	// Backend

}

func binarySearch(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (high + low) / 2

		if value == a[mid] {
			return mid
		} else if value < a[mid] {
			high = mid - 1
		} else if value > a[mid] {
			low = mid + 1
		}
	}

	return -1
}
