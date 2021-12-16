package main

import (
	"fmt"
	"time"
)

type Corrutina struct {
	ch chan int
}

func (c Corrutina) Hold() {
	<-c.ch
}

func (c Corrutina) Active() {
	time.Sleep(1000 * time.Millisecond)
	c.ch <- 0
}

func Visualiza() {
	fmt.Println("Funcion Visualiza, Antes del Hold")
	c.Active()
	c.Hold()
	fmt.Println("Funcion Visualiza, Despues del Hold")
	c.Active()
}

var channel = make(chan int)
var c = Corrutina{channel}

func main() {
	fmt.Println("En funcion main")
	go Visualiza()
	c.Hold()
	fmt.Println("Regreso al main")
	c.Active()
	c.Hold()
	fmt.Println("Regreso al main")
	close(channel)
}
