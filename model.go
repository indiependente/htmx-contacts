package main

// Contact represents a contact in the address book.
type Contact struct {
	ID    int
	Name  string
	Email string
}

// newContact creates a new contact with the given name and email.
func newContact(name, email string) Contact {
	id++
	return Contact{
		Name:  name,
		Email: email,
		ID:    id,
	}
}

// Contacts is a map of contacts.
type Contacts map[int]Contact

// hasEmail returns true if the contacts contain the given email.
// It returns false otherwise.
func (c Contacts) hasEmail(email string) bool {
	for _, contact := range c {
		if contact.Email == email {
			return true
		}
	}
	return false
}

// Data represents the data of the application.
type Data struct {
	Contacts Contacts
}

// newData creates a new data with some initial contacts.
func newData() Data {
	return Data{
		Contacts: Contacts{
			1: newContact("Alice", "alice@email.com"),
			2: newContact("Bob", "bob@email.com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}
