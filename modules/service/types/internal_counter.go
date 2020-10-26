package types

type InternalCounter struct {
	count int64
}

func (vtc *InternalCounter) Count() int64 {
	return vtc.count
}

func (vtc *InternalCounter) Incr() {
	vtc.count++
}

func NewInternalCounter() *InternalCounter {
	return &InternalCounter{
		count: 0,
	}
}
