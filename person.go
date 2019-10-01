package person

import "fmt"

// Person struct represents a person
// Name = person's name
// GettingReadyTime = time needed for the person to get ready
// Ready = this channel will close when person is ready to go
type Person struct {
	Name             string
	GettingReadyTime int
	Ready            chan struct{}
}

// Person behaviours

// GrabGlasses simulates a person grabbing their glasses
func (p Person) GrabGlasses() {
	fmt.Println(p.Name, " grabs the glasses")

	dur := waitRandomTime(5, 10)
	p.GettingReadyTime += dur

	fmt.Printf("It took %s %d seconds\n\n", p.Name, dur)

}

// TideShoes simulates a person tiding their shoes
func (p Person) TideShoes() {
	fmt.Println(p.Name, " tides the shoes")

	dur := waitRandomTime(20, 35)
	p.GettingReadyTime += dur

	fmt.Printf("It took %s %d seconds\n\n", p.Name, dur)

}

// CloseWindow simulates a person closing the windows
func (p Person) CloseWindow() {
	fmt.Println(p.Name, " closes the window")

	dur := waitRandomTime(2, 5)
	p.GettingReadyTime += dur

	fmt.Printf("It took %s %d seconds\n\n", p.Name, dur)

}

// TurnOffTheFan simulates a person turning off the ceiling fan
func (p Person) TurnOffTheFan() {
	fmt.Println(p.Name, " turns off the fan")

	dur := waitRandomTime(3, 6)
	p.GettingReadyTime += dur

	fmt.Printf("It took %s %d seconds\n\n", p.Name, dur)

}

// PocketBelongings simulates a person pocketing their belongings
func (p Person) PocketBelongings() {
	fmt.Println(p.Name, " pockets the belongings")

	dur := waitRandomTime(5, 40)
	p.GettingReadyTime += dur

	fmt.Printf("It took %s %d seconds\n\n", p.Name, dur)

}
