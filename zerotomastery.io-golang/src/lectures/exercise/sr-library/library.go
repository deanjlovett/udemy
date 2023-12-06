//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
) 

type Member struct {
	// id int
	memName string
	books map[string] *Book
}
type MemberMap map[string] Member

type Book struct {
	// id int
	bookName string
	checkOuts [] CheckOut
	lastCheckOut *CheckOut
}
type BookMap map[string] Book

type CheckOut struct{
	memberPtr *Member
	checkOutTime time.Time
	checkInTime time.Time
}

type Library struct{
	members MemberMap
	books BookMap
	checkedOutBooks map[string] *Book
}

func checkOutBook( aLibraryPtr *Library, aBookPtr *Book, aMemberPtr *Member){
	fmt.Println("checkOutBook")
	if aLibraryPtr == nil || aBookPtr==nil || aMemberPtr==nil {
		return // complain
	}
	if aBookPtr.lastCheckOut != nil {
		return // book already checked out
	}
	var checkout CheckOut
	checkout.checkOutTime = time.Now()
	checkout.memberPtr =  aMemberPtr
	fmt.Println("  checkout:",checkout)

	fmt.Println("  ",aBookPtr.checkOuts)
	aBookPtr.checkOuts = append(aBookPtr.checkOuts,checkout)
	fmt.Println("  ",aBookPtr.checkOuts)
	aBookPtr.lastCheckOut = &(aBookPtr.checkOuts[len(aBookPtr.checkOuts)-1])

	aMemberPtr.books[aBookPtr.bookName] = aBookPtr

	aLibraryPtr.checkedOutBooks[aBookPtr.bookName] = aBookPtr
}
func checkInBook( aLibraryPtr *Library, aBookPtr *Book){
	if aLibraryPtr == nil || aBookPtr==nil {return}// complain}
	if aBookPtr.lastCheckOut==nil { return }

	aBookPtr.lastCheckOut.checkInTime = time.Now()
	memberPtr := aBookPtr.lastCheckOut.memberPtr
	if memberPtr == nil {
		return
	}
	delete( memberPtr.books, aBookPtr.bookName)
	delete( aLibraryPtr.checkedOutBooks, aBookPtr.bookName)
	aBookPtr.lastCheckOut = nil
}
func addMember(aLibraryPtr *Library, name string) Member{
	//if aLibraryPtr==nil {return}
	aLibraryPtr.members[name] = Member{
		memName:name,
		books:make(map[string] *Book),
	}
	return aLibraryPtr.members[name]
}
func addBook(aLibraryPtr *Library, name string) Book{
	//if aLibraryPtr==nil {return}
	aLibraryPtr.books[name] = Book{
		bookName:name,

	}
	return aLibraryPtr.books[name]
}
func printLibraryState(aLibraryPtr *Library){
	if aLibraryPtr==nil {return}
	fmt.Println("=====================")
	fmt.Println("library:")
	fmt.Println("  checked out books:")
	for key,elem := range aLibraryPtr.checkedOutBooks{
		fmt.Println("  ",key,":",elem)
	}
	fmt.Println()
	fmt.Println("members:")
	for key,elem := range aLibraryPtr.members {
		fmt.Println("  ",key,":",elem)
		if len(elem.books)>0 {
			fmt.Println("  books:")
			for bk,book := range elem.books {
				fmt.Println("    ",bk,book)
			}
		}
	}
	if len(aLibraryPtr.members )==0 {
		fmt.Println("(no members)")
	}
	fmt.Println()
	fmt.Println("books:")
	for key,elem := range aLibraryPtr.books {
		fmt.Println("  ",key,":", elem)
		if elem.lastCheckOut!=nil {
			fmt.Println("    last checkout:",elem.lastCheckOut.memberPtr.memName)
		}
	}
	if len(aLibraryPtr.books )==0 {
		fmt.Println("(no books)")
	}
}

func main() {
	var aLibrary Library
	aLibrary.members = make(MemberMap) 
	aLibrary.books = make(BookMap)
	aLibrary.checkedOutBooks = make(map[string] *Book)

	printLibraryState(&aLibrary)

	addMember(&aLibrary,"member_A")
	addMember(&aLibrary,"member_B")
	aMember := addMember(&aLibrary,"member_C")

	printLibraryState(&aLibrary)

	addBook(&aLibrary,"book_A")
	addBook(&aLibrary,"book_B")
	addBook(&aLibrary,"book_C")
	aBook := addBook(&aLibrary,"book_D")

	printLibraryState(&aLibrary)

	checkOutBook(&aLibrary, &aBook, &aMember)
	printLibraryState(&aLibrary)

	checkInBook(&aLibrary, &aBook)
	printLibraryState(&aLibrary)

}
