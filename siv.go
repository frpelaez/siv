package siv

type Handle struct {
	Index      uint32
	Generation uint32
}

type slot[T any] struct {
	Value      T
	Generation uint32
	NextFree   int32
	IsActive   bool
}

type Siv[T any] struct {
	items    []slot[T]
	freeHead int32
	count    int
}

func New[T any](capacity int) *Siv[T] {
	return &Siv[T]{
		items:    make([]slot[T], 0, capacity),
		freeHead: -1,
		count:    0,
	}
}

func (s *Siv[T]) Add(value T) Handle {
	var idx int
	if s.freeHead != -1 {
		idx = int(s.freeHead)
		s.freeHead = s.items[idx].NextFree
		s.items[idx].Value = value
		s.items[idx].IsActive = true
	} else {
		idx = int(len(s.items))
		slot := slot[T]{
			Value:      value,
			Generation: 0,
			NextFree:   -1,
			IsActive:   true,
		}
		s.items = append(s.items, slot)
	}

	s.count++

	return Handle{
		Index:      uint32(idx),
		Generation: s.items[idx].Generation,
	}
}

func (s *Siv[T]) Get(handle Handle) (*T, bool) {
	if int(handle.Index) >= len(s.items) {
		return nil, false
	}

	slot := &s.items[handle.Index]

	if slot.Generation != handle.Generation {
		return nil, false
	}

	return &slot.Value, true
}

func (s *Siv[T]) Remove(handle Handle) bool {
	if int(handle.Index) >= len(s.items) {
		return false
	}

	slot := &s.items[handle.Index]

	if slot.Generation != handle.Generation {
		return false
	}

	slot.Generation++

	var zero T
	slot.Value = zero
	slot.IsActive = false

	slot.NextFree = s.freeHead
	s.freeHead = int32(handle.Index)

	return true
}

func (s *Siv[T]) ForEach(fn func(h Handle, v *T) bool) {
	for i := 0; i < len(s.items); i++ {
		slot := &s.items[i]

		if !slot.IsActive {
			continue
		}

		h := Handle{
			Index:      uint32(i),
			Generation: slot.Generation,
		}

		if !fn(h, &slot.Value) {
			return
		}
	}
}
