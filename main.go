package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var POP_SIZE int =
var MUTATION_RATE = 0.005 // .5% chance of mutating

type Organism struct {
	DNA     []byte
	Fitness float64
}

func (d *Organism) calcFitness(target []byte) {
	score := 0
	for i := 0; i < len(d.DNA); i++ {
		if d.DNA[i] == target[i] {
			score++
		}
	}
	d.Fitness = float64(score) / float64(len(d.DNA))
	return
}

func createOrganism(target []byte) (organism Organism) {
	ba := make([]byte, len(target))
	for i := 0; i < len(target); i++ {
		ba[i] = byte(rand.Intn(95) + 32) // introduces variety: this is black magic
	}
	// Form object
	organism = Organism{
		DNA:     ba,
		Fitness: 0,
	}
	organism.calcFitness(target)
	return
}

func createPopulation(target []byte) (population []Organism) {
	population = make([]Organism, POP_SIZE)
	for i := 0; i < POP_SIZE; i++ {
		population[i] = createOrganism(target)
	}
	return
}

// Create a pool of organisms
func createPool(population []Organism, target []byte, maxFitness float64) (pool []Organism) {
	pool = make([]Organism, 0)
	// create a pool for next generation
	for i := 0; i < len(population); i++ {
		population[i].calcFitness(target)
		num := int((population[i].Fitness / maxFitness) * 100)
		// The higher the fitness the more children it will create.
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

// Baby maker lollllzzz xdddd
func crossover(d1 Organism, d2 Organism) (child Organism) {
	child = Organism{
		DNA:    make([]byte, len(d1.DNA)),
		Fitness: 0,
	}
	// Generate random number with length of DNA and loop
	// over it to take DNA from both Organisms.
	mid := rand.Intn(len(d1.DNA))
	for i := 0; i < len(d1.DNA); i++ {
		if i > mid {
			child.DNA[i] = d1.DNA[i]
		} else {
			child.DNA[i] = d2.DNA[i]
		}	}
	return child
}

func (d *Organism) mutate() {
	// Loop over DNA and randomly mutate some bits
	// depending on mutation rate
	for i := 0; i < len(d.DNA); i++ {
		if rand.Float64() < MUTATION_RATE {
			d.DNA[i] = byte(rand.Intn(95) + 32)
		}
	}
}

func naturalSelection(pool []Organism, population []Organism, target []byte) []Organism {
	nextGen := make([]Organism, len(population))
	for i := 0; i < len(population); i++ {
		// Generate two random integers
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		// Select two random elements in previously created pool
		a := pool[r1]
		b := pool[r2]

		// Make baby ;))
		child := crossover(a, b)
		// Mutate the baby
		child.mutate()
		child.calcFitness(target)

		nextGen[i] = child
	}
	return nextGen
}

func getBest(population []Organism) (bestOrganism Organism) {
	bestOrganism = Organism{
		Fitness: 0,
	}
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > bestOrganism.Fitness {
			bestOrganism = population[i]
		}
	}
	return
}

func main() {
	fmt.Println("Highly crackhead genetic algorithm made by bert")
	var target []byte
	fmt.Printf("Enter a target string: ")
	reader := bufio.NewReader(os.Stdin)
	target, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	target = target[:len(target)-2]

	start := time.Now()
	rand.Seed(start.UTC().UnixNano())
	population := createPopulation(target)

	fCompleted := false
	generation := 0

	for !fCompleted { // why the fuck are there no while loops???
		generation++
		bestOrganism := getBest(population)
		fmt.Printf("\r generation: %d | %s | fitness: %2f", generation, string(bestOrganism.DNA), bestOrganism.Fitness)
		// If bytes are an exact match with the target
		if bytes.Compare(bestOrganism.DNA, target) == 0 {
			fCompleted = true
		} else { // Create a new generation
			maxFitness := bestOrganism.Fitness
			pool := createPool(population, target, maxFitness)
			population = naturalSelection(pool, population, target)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\nTook %f seconds\n", elapsed.Seconds())
}
