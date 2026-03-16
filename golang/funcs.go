package golang


//no dependencies


/*
 * helpers for functions
 */


//run many functions with one param
func RunManyWith[T any](funcs []func(T), param T) {
	for _, f := range funcs { f(param) }
}

//run a function for each param
func RunForMany[T any](f func(T), params []T) {
	for _, p := range params { f(p) }
}

//run many functions with many params
//  WARN: if you need this, GET HELP
func RunManyForMany[T any](funcs []func(T), params []T) {
	for _, f := range funcs {
		for _, p := range params {
			f(p)
		}
	}
}
