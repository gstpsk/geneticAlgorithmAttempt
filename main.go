package main

import (
	"fmt"
	"math/rand"
)

var POP_SIZE int = 69

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

// Selection
func createPool(population []Organism, target []byte, maxFitness float64) (pool []Organism) {
	pool = make([]Organism, 0)
	// create a pool for next generation
	for i := 0; i < len(population); i++ {
		population[i].calcFitness(target)
		num := int((population[i].Fitness / maxFitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

func main() {
	fmt.Println("testje")
}
