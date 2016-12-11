/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package gen

import "math/rand"

//New returns generator of random number [0,2)
func New(s rand.Source, ct int64) <-chan int {
	ch := make(chan int)
	r := rand.New(s)

	go func(ct int64) {
		for i := int64(0); i < ct; i++ {
			ch <- r.Intn(2)
		}
		close(ch)
	}(ct)

	return ch
}
