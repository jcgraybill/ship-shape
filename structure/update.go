package structure

// FIXME planet with zero habitability should not produce population

func (s *Structure) Update(count int) bool {
	if s.data.Produces.Rate > 0 {
		if s.storage.Resource == s.data.Produces.Resource && s.storage.Amount < s.storage.Capacity {
			productionRate := 255 - (int(s.data.Produces.Rate) * (int(s.Planet().Resources()[s.data.Produces.Requires]) / 255))
			if count%productionRate == 0 {
				s.storage.Amount += 1
				return true
			}
		}
	}

	return false
}
