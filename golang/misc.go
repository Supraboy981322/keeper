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
