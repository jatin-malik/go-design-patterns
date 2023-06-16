package behavioral

import "fmt"

// Iterator is a behavioral design pattern that allows sequential traversal through a complex data structure
// without exposing its internal details.

type Iterator interface {
	hasNext() bool
	getNext() any
}

type Collection interface {
	createIterator() Iterator
}

type User struct {
	Name string
	Age  int
}

type UserCollection struct {
	users []*User
}

type UserIterator struct {
	idx   int     // cur element
	users []*User // has its own copy of the elements
}

func (u *UserIterator) hasNext() bool {
	if u.idx < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() any {
	elem := (*User)(nil)
	if u.hasNext() {
		elem = u.users[u.idx]
		u.idx++
	}
	return elem
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{users: u.users}
}

func RunIteratorDemo() {

	user1 := &User{Name: "Elliot", Age: 25}
	user2 := &User{Name: "Darlene", Age: 24}

	// make a users collection
	c := UserCollection{
		users: []*User{user1, user2},
	}

	uci := c.createIterator()

	for uci.hasNext() {
		fmt.Println(uci.getNext().(*User)) // type assertion
	}
}
