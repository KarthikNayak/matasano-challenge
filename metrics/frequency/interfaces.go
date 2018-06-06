package frequency

type Frequency interface {
	GetFrequency(b []byte) float64
}
