package helper

type Waiter[T any] struct {
	stopChan chan struct{}
	resChan  chan T
}

func NewWaiter[T any]() *Waiter[T] {
	return &Waiter[T]{
		stopChan: make(chan struct{}),
		resChan:  make(chan T),
	}
}

func (w Waiter[T]) Run(action func() T) {
	go func() {
		select {
		case <-w.stopChan:
			return
		default:
		}

		res := action()

		select {
		case <-w.stopChan:
		case w.resChan <- res:
		}
	}()
}

func (w Waiter[T]) Result() <-chan T {
	return w.resChan
}

func (w Waiter[T]) Close() {
	close(w.stopChan)
}
