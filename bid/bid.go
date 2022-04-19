package bid

import "github.com/jcgraybill/ship-shape/structure"

type Bid struct {
	Structure *structure.Structure
	Resource  int
	Urgency   uint8
}

func New() *Bid {
	bid := Bid{
		Resource: 1,
		Urgency:  1,
	}
	return &bid
}
