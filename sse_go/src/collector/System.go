package collector

type System struct {
}

func (SystemPtr *System) Collect() *System {

	return SystemPtr
}
