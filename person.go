package person

// Person struct represents a person
// Name = person's name
// GettingReadyTime = time needed for the person to get ready
// Ready = this channel will close when person is ready to go
type Person struct {
	Name             string
	GettingReadyTime int
	Ready            chan struct{}
}
