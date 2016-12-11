/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package toss

import (
	"fmt"
	"io"
	"math/rand"

	"github.com/spiegel-im-spiegel/cointoss/gen"
	"github.com/spiegel-im-spiegel/gocli"
)

//Context is context for plot package.
type Context struct {
	ui        *gocli.UI
	src       rand.Source
	tossCount int64
}

//NewContext returns Context instance
func NewContext(w, e io.Writer, rs rand.Source, ct int64) *Context {
	return &Context{ui: gocli.NewUI(nil, w, e), src: rs, tossCount: ct}
}

//Execute output coin toss.
func Execute(cxt *Context) (int64, error) {
	if cxt.tossCount <= 0 {
		return 0, fmt.Errorf("invalid argument \"%v\" for --toss option", cxt.tossCount)
	}

	ch := gen.New(cxt.src, cxt.tossCount)
	fronts := int64(0)
	for t := range ch {
		cxt.ui.Outputln(t)
		if t != 0 {
			fronts++
		}
	}

	return fronts, nil
}
