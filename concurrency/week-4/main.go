package main

import (
	"fmt"
	"sync"
	"time"
)

// Host who decide who can eat
type Host struct {
	stopper chan int
}

// Chopstick represets a single chopstick to eat
type Chopstick struct {
	sync.Mutex
}

// Philosopher A ver normal human
type Philosopher struct {
	id                            int
	leftChopstick, rightChopstick *Chopstick
}

// GiveMePermission Allow only 2 phillosophers eating
func (host Host) GiveMePermission() {
	host.stopper <- 1
}

// Done Release chopsticks
func (host Host) Done() {
	<-host.stopper
}

// Eat eat, no more.
func (philosopher *Philosopher) Eat(wg *sync.WaitGroup, host Host) {
	for i := 0; i < 3; i++ {
		host.GiveMePermission()
		philosopher.leftChopstick.Lock()
		philosopher.rightChopstick.Lock()
		fmt.Printf("Starting to eat %d\n", philosopher.id)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Finishing eating %d. (time: %d)\n", philosopher.id, i+1)
		philosopher.rightChopstick.Unlock()
		philosopher.leftChopstick.Unlock()
		host.Done()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	host := Host{
		stopper: make(chan int, 2),
	}

	chopsticks := make([]*Chopstick, 0)
	philosophers := make([]*Philosopher, 0)

	fmt.Println("Preparing the chopsticks")

	for i := 0; i < 5; i++ {
		chopsticks = append(chopsticks, &Chopstick{})
	}

	fmt.Println("Preparing the table")

	for i := 0; i < 5; i++ {
		philosopher := Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
		philosophers = append(philosophers, &philosopher)
	}

	fmt.Printf("Let's go\n\n")

	for j := 0; j < 5; j++ {
		go philosophers[j].Eat(&wg, host)
	}

	wg.Wait()
}
