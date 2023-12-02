package aoc_utils

type Channel[T any] chan T

func NewChannel[T any](buffer int) Channel[T] {
	return make(chan T, buffer)
}

func NewChannelFromSlice[T any](input []T, buffer int) Channel[T] {
	channel := NewChannel[T](buffer)
	go func() {
		for _, a := range input {
			channel <- a
		}
		close(channel)
	}()
	return channel
}

func (input Channel[T]) Filter(f func(T) bool) (output Channel[T]) {
	output = NewChannel[T](10)
	go func() {
		for a := range input {
			if f(a) {
				output <- a
			}

		}
		close(output)
	}()
	return
}

func (input Channel[T]) Map(f func(T) T) (output Channel[T]) {
	output = NewChannel[T](10)
	go func() {
		for a := range input {
			output <- f(a)
		}
		close(output)
	}()
	return
}

func (input Channel[T]) Reduce(f func(T, T) T) (a T) {
	more := false
	a, more = <-input
	if more {
		for b := range input {
			a = f(a, b)
		}
	}
	return
}

type Channels[T any] chan chan T

func NewChannels[T any](buffer int) Channels[T] {
	return make(chan chan T, buffer)
}

func (input Channel[T]) Split(f func(T) bool) Channels[T] {
	output := NewChannels[T](10)
	go func() {
		current := NewChannel[T](10)
		for a := range input {
			if f(a) {
				close(current)
				output <- current
				current = NewChannel[T](10)
			} else {
				current <- a
			}
		}
		close(current)
		output <- current
		close(output)
	}()
	return output
}
