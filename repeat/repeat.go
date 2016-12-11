/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package repeat

import (
	"fmt"
	"io"
	"math"
	"math/rand"

	"github.com/spiegel-im-spiegel/cointoss/toss"
	"github.com/spiegel-im-spiegel/gocli"
)

//Context is context for plot package.
type Context struct {
	ui          *gocli.UI
	errorWriter io.Writer
	src         rand.Source
	tossCount   int64
	repeatCount int64
}

//NewContext returns Context instance
func NewContext(w, e io.Writer, rs rand.Source, ct, cr int64) *Context {
	return &Context{ui: gocli.NewUI(nil, w, e), errorWriter: e, src: rs, tossCount: ct, repeatCount: cr}
}

//Execute output coin toss.
func Execute(cxt *Context) error {
	if cxt.repeatCount <= 0 {
		return fmt.Errorf("invalid argument \"%v\" for --rcount option", cxt.repeatCount)
	}
	rcf := float64(cxt.repeatCount)

	min := int64(cxt.tossCount + 1)
	max := int64(0)
	sum := int64(0)
	sum2 := int64(0)
	for i := int64(0); i < cxt.repeatCount; i++ {
		fronts, err := toss.Execute(toss.NewContext(nil, cxt.errorWriter, cxt.src, cxt.tossCount))
		if err != nil {
			return err
		}
		cxt.ui.Outputln(fronts)
		if fronts < min {
			min = fronts
		}
		if fronts > max {
			max = fronts
		}
		sum += fronts
		sum2 += fronts * fronts
	}
	//statistics
	cxt.ui.OutputErrln(fmt.Sprintf("minimum value: %v", min))
	cxt.ui.OutputErrln(fmt.Sprintf("maximum value: %v", max))
	ave := float64(sum) / rcf
	cxt.ui.OutputErrln(fmt.Sprintf("average value: %7.5f", ave))
	cxt.ui.OutputErrln(fmt.Sprintf("standard deviation: %7.5f", math.Sqrt(float64(sum2)/rcf-ave*ave)))

	return nil
}
