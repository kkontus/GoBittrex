package talgo

func Average(l []float64, h []float64) []float64 {
	return avgSlices(l, h)
}

func avgSlices(l []float64, h []float64) []float64 {
	avg := []float64{}
	if l != nil && h != nil && len(l) == len(h) {
		for i, _ := range l {
			m := (l[i] + h[i]) / 2
			avg = append(avg, m)
		}
	}
	return avg
}
