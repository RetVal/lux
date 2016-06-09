package tornago

// Constraint represents all types of constraint between rigid bodies.
type Constraint interface {
	// GenerateContacts is given a slice of contacts of size at least 1. Do not
	// increase the slice of the slice. The reason that the size is limited is
	// to better control memory allocation and time spent resolving contacts.
	// Returns how many contacts we're generated.
	GenerateContacts([]Contact) int
}
