package phil

type Observer[T any] interface {
	Notify(event T)
	Close()
}

type Subject interface {
	Notify(interface{})
	register(chan<- interface{})
	deregister(chan<- interface{})
}

type subject struct {
	observers []chan<- interface{}
}

func (s *subject) register(observer chan<- interface{}) {
	s.observers = append(s.observers, observer)
}

func (s *subject) deregister(observer chan<- interface{}) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *subject) Notify(event interface{}) {
	for _, o := range s.observers {
		o <- event
	}
}

func RegisterObserver[T any](subject Subject, channel chan interface{}, fnNotify func(T)) {
	subject.register(channel)
	go func() {
		for {
			read, more := <-channel
			if more {
				if event, ok := read.(T); ok {
					fnNotify(event)
				}
			} else {
				subject.deregister(channel)
				println("deregistered observer")
				return
			}
		}
	}()
}
