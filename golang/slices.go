package golang


//no dependencies


/*
 * slice related helpers
 */

//remove last item from slice (returns item)
//   NOTE: having to do this is moronic:
//   'slice = slice[:len(slice)-1]'
func Pop[S ~[]T, T any](buf *S) T {
	var item T
	if len(*buf) > 0 {
		item = (*buf)[len(*buf)-1]
		*buf = (*buf)[:len(*buf)-1]
	}
	return item
}

//remove first item from slice (returns item)
//   NOTE: see 'pop(...)'
func Shift[S ~[]T, T any](buf *S) T {
	var item T
	if len(*buf) > 0 {
		item = (*buf)[0]
		*buf = (*buf)[1:]
	}
	return item
}

//append to a slice (doesn't return new slice)
//  NOTE: same comment applies to this
//   'slice = append(slice, thing)'
func Add[S ~[]T, T any](buf *S, thing ...T) {
	*buf = append(*buf, thing...)
}

//empties a slice and returns slice of old items
//  (original slice becomes empty, useful using a slice a buffer)
//  NOTE: why isn't this a built-in?
func Drain[S ~[]T, T any](buf *S) []T {
	var res []T
	for len(*buf) > 0 {
		Add(&res, Shift(buf))
	}
	return res
}

//filters item 'thing' from slice (doesn't return new slice)
func Filter[S ~[]T, T comparable](buf *S, thing T) {
	l := len(*buf)
	for _, a := range *buf {
		if a != thing { Add(buf, a) }
	}
	for range l { Shift(buf) }
}
