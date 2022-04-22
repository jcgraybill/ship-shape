package planet

func (p *Planet) Center() (int, int) {
	return p.x, p.y
}

func (p *Planet) Name() string {
	return p.name
}

func (p *Planet) Resources() map[int]uint8 {
	return p.resources
}
