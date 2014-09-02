package magistrate

import (
	"strconv"

	"github.com/cloudfoundry-incubator/guillotine/inquisitor"
	"github.com/cloudfoundry-incubator/guillotine/headsman"

)

type Magistrate struct {
	iq inquisitor.Inquisitor
	h headsman.Headsman
}

func NewMagistrate(inq inquisitor.Inquisitor, h headsman.Headsman) *Magistrate {
	return &Magistrate{
		iq : inq,
		h : h,
	}
}

func (g *Magistrate) DeliberateAndExecute() {
	result, err := g.iq.ReadCsv()

	if err != nil {
		g.h.Chop()
	} else {
		connectionArray := g.iq.ConnectionArrayFromCsv(result)

		var count = 0
		for _,element := range connectionArray {
			int_element, _ := strconv.Atoi(element)
			if int_element > 0 {
				count++
			}
		}

		if count > 1 {
			g.h.Chop()
		}
	}
}

