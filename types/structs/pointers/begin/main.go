// types/structs/pointers/begin/main.go
package main

type author struct {
	first string
	last  string
}

// fullName returns the full name of the author
func (a author) fullName() string {
	return a.first + " " + a.last
}

// changeName changes the first and last name of the author
//

func main() {
	// a := author{
	// 	first: "Mar1",
	// 	last:  "Twain",
	// }

	// fmt.Println(a.fullName())

	// call changeName to update name of author
	//

	// fmt.Println(a.fullName())
}
