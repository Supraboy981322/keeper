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
	var res S
	DrainInto(&res, buf)
	return res
}

//drains into an existing slice (appending to it)
func DrainInto[S ~[]T, T any](into *S, from *S) {
	for len(*from) > 0 {
		Add(into, Shift(from))
	}
}
func DrainAdd[S ~[]T, T any](append_to *[]S, from *S) {
	*append_to = append(*append_to, []T{})
	for len(*from) > 0 {
		item := (*from)[0]
		(*from) =(*from)[1:]
		(*append_to)[len(*append_to)-1] = append((*append_to)[len(*append_to)-1], item)
	}
}

//filters item 'thing' from slice (doesn't return new slice)
func Filter[S ~[]T, T comparable](buf *S, thing T) {
	l := len(*buf)
	for _, a := range *buf {
		if a != thing { Add(buf, a) }
	}
	for range l { Shift(buf) }
}
