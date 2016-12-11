/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package gen

import (
	"fmt"
	"math/rand"

	"github.com/seehuhn/mt19937"
)

//RNGs is kind of RNG
type RNGs int

//Kind of RNG
const (
	NULL RNGs = iota + 1
	GO
	MT
)

func (r RNGs) String() string {
	switch r {
	case MT:
		return "MT"
	default:
		return "default"
	}
}

//NewRndSource returns Source of random numbers
func NewRndSource(rng string, seed int64) (rand.Source, error) {
	r, err := rngType(rng)
	switch r {
	case MT:
		mt := mt19937.New()
		mt.Seed(seed)
		return mt, err
	default:
		return rand.NewSource(seed), err
	}
}

//RngType returns RNGs from string.
func rngType(s string) (RNGs, error) {
	switch s {
	case "MT":
		return MT, nil
	case "GO":
		return GO, nil
	default:
		return NULL, fmt.Errorf("invalid -rsource parameter: %s\n", s)
	}
}
