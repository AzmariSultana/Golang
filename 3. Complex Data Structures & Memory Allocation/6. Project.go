package main

import "fmt"

// a struct groups related fields into one type
type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact // all contacts, in the order they were added
var contactIndexByName map[string]int // name -> position in contactList, for fast lookup
var nextID = 1

// init runs automatically before main
func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int) // a map must be made before you can add keys
}

func addContact(name, email, phone string) { // name, email, phone all have type string
	if _, exists := contactIndexByName[name]; exists { // names must be unique
		fmt.Printf("Contact already exists: %v\n", name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1 // the new contact is the last element
	fmt.Printf("Contact added: %v\n", name)
}

// returns a pointer to the contact, or nil if there is no contact with that name
func findContact(name string) *Contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index] // address of the element inside the slice, not a copy
	}
	return nil
}

func ListContacts() {
	fmt.Println("--- Listing Contacts ---")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for i, contact := range contactList { // contact is a copy of each element
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", i+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}

	fmt.Println("")
}

func main() {

	addContact("Azmari Wonderland", "azmari@example.com", "111-2222")
	addContact("Sultana The Builder", "sultana@example.com", "333-4444")
	addContact("Tushmi Brown", "tushmi@example.com", "555-6666")
	addContact("Azmari Wonderland", "azmari.new@example.com", "777-8888") // Attempt to add duplicate

	ListContacts()

	sultana := findContact("Sultana") // the key must match exactly; the saved name is "Sultana The Builder"
	if sultana == nil {
		fmt.Println("No Sultana contact found.")
	} else {
		fmt.Println("Sultana contact found.", sultana.Name) // Go reads sultana.Name through the pointer automatically
	}

}
