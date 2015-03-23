package collector

type System struct {
}

func (SystemPtr *System) Collect() <-chan *System {
	out := make(chan *System)

	go func() {

		close(out)
	}()

	return out
}
