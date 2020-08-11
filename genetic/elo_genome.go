package genetic

import (
	"math/rand"

	"github.com/CArnoud/go-rebbl-elo/database/models"
)

// EloPredictor encapsulates the interaction between EloGenome and a generic EloPredictor.
type EloPredictor interface {
	ProcessMatch(*models.Match)
	Evaluate() (float64)
}

// EloGenome represents an individual parameter set for an Elo predictor.
type EloGenome struct {
	predictor EloPredictor
	kFactor int
	deviation int
	norm int
}

// NewEloGenome instantiates an EloGenome.
func NewEloGenome(kFactor int, deviation int, norm int) *EloGenome {
	return &EloGenome{
		kFactor: kFactor,
		deviation: deviation,
		norm: norm,
	}
}

// Evaluate calculates the fitness of the genome.
func (eg *EloGenome) Evaluate() (float64, error) {
	return 0.0, nil
}

// Mutate causes a random change in the genome.
func (eg *EloGenome) Mutate(rng *rand.Rand) {

}

// Crossover combines two individuals.
func (eg *EloGenome) Crossover(eg2 EloGenome, rng *rand.Rand) {

}

// Clone creates a copy of the current individual.
func (eg *EloGenome) Clone() EloGenome {
	return *NewEloGenome(eg.kFactor, eg.deviation, eg.norm)
}
