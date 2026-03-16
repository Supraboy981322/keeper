package golang

//no dependencies

/*
 * misc helpers
 */

//helper to flip a boolean
//   (why isn't this a built-in?)
//  NOTE: this being the expected way is stupid
//    'boolean = !boolean'
func Flip(thing *bool) {
	*thing = !*thing
}

//helper to panic if bool is false
//   (why isn't this a built-in?)
func Assert(condition bool) {
	if !condition { panic("assertion failed") }
}

//helper to assert a series of conditions
func AssertMany(conditions []bool) {
	RunForMany(Assert, conditions)
}

func AssertFunc(f func(arg any) bool, args any) {
	Assert(
		f(args),
	)
}

func AssertManyFunc[T any](f func(arg T) bool, things []T) {
	RunForMany(
		func (thing T) {
			Assert(f(thing))
		},
		things,
	)
}
