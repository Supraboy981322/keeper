package golang

//no dependencies

/*
 * helpers for functions
 */

func RunManyWith[T any](funcs []func(T), param T) {
	for _, f := range funcs { f(param) }
}

func RunForMany[T any](f func(T), params []T) {
	for _, p := range params { f(p) }
}

//if you need this, GET HELP
func RunManyForMany[T any](funcs []func(T), params []T) {
	for _, f := range funcs {
		for _, p := range params {
			f(p)
		}
	}
}
