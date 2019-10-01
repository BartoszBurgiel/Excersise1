package main

import (
	"dailywalk/person"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Concurrency Excersise 1")
	fmt.Println("")

	var gettingReadyWG sync.WaitGroup

	// Bob person setup
	bob := person.NewPerson("Bob")

	// Alice person setup
	alice := person.NewPerson("Alice")

	// Channels to control events only occuring once
	// (if the window is already closed it does not to be closed again)
	windowClosedC := make(chan bool)
	fanOffC := make(chan bool)

	gettingReadyWG.Add(1)
	// Set windowClosedC and fanOffC to false by default
	go func() {
		windowClosedC <- false
		fanOffC <- false

		gettingReadyWG.Done()
	}()

	gettingReadyWG.Add(2)
	// Goroutine for Bob
	go morningRoutine(bob, windowClosedC, fanOffC, &gettingReadyWG)

	// Goroutine for Alice
	go morningRoutine(alice, windowClosedC, fanOffC, &gettingReadyWG)

	// "Wait" untill both persons are ready
	<-bob.Ready
	<-alice.Ready

	// After all persons got ready start the alarm part
	gettingReadyWG.Wait()

	done := make(chan struct{})

	go func() {

		// Arm the alarm
		alarm := time.NewTicker(time.Millisecond)

		// Controll the execution of tiding shoes by the persons
		var shoesWg sync.WaitGroup

		ignitionTime := time.Now()

		fmt.Println("Arming alarm")

		shoesWg.Add(2)
		go bob.TideShoes(&shoesWg)
		go alice.TideShoes(&shoesWg)

		// Whitespace
		fmt.Println("")

		shoesWg.Wait()

		for {

			select {
			case <-alarm.C:

				// Alarm is armed after 500ms
				if time.Since(ignitionTime).Round(time.Millisecond) > 500 {

					alarm.Stop()
					fmt.Println("Alarm is armed")

					done <- struct{}{}

				}
			}
		}
	}()

	<-done
}

// morning routine describes a morning routine for each person
func morningRoutine(p *person.Person, w, f chan bool, wg *sync.WaitGroup) {
	fmt.Printf("%s starts getting ready\n\n", p.Name)

	p.GrabGlasses()
	p.TightenBelt()

	p.CloseWindow(w)
	p.TurnOffTheFan(f)

	p.PocketBelongings()

	fmt.Printf("%s is ready to go!\n %s spent %d seconds on getting ready\n\n", p.Name, p.Name, p.GettingReadyTime)
	p.Ready <- struct{}{}
	wg.Done()
}
