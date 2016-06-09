package tornago

// Dispatcher is an algorithm that resolves the set of contact generated at
// every step.
type Dispatcher interface {
	ResolveContacts([]Contact, float32)
}
