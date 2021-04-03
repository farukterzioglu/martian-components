package rollover

import (
	"github.com/farukterzioglu/martian-components/header/rollover/modifier"
	"github.com/google/martian/parse"
)

func init() {
	parse.Register("header.RollOver", modifier.FromJSON)
}
