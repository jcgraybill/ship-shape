package planet

func (p *Planet) Center() (int, int) {
	return p.Bounds.Min.X + p.Bounds.Dx()/2, p.Bounds.Min.Y + p.Bounds.Dy()/2
}

func (p *Planet) Name() string {
	return p.name
}

func (p *Planet) Resources() map[int]uint8 {
	return p.resources
}
