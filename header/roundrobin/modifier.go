package rollover

import (
	"github.com/farukterzioglu/martian-components/header/roundrobin/modifier"
	"github.com/google/martian/parse"
)

func init() {
	parse.Register("header.RoundRobin", modifier.FromJSON)
}
