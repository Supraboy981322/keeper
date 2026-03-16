package main

import (
	"fmt"
	"slices"
	"errors"
	keeper "github.com/Supraboy981322/keeper/golang"
)

func main() {
	{
		//sanity check
		fmt.Println("keeper.Assert(true)\nkeeper.AssertMany([]bool{ true, true })")
		keeper.Assert(true)
		keeper.AssertMany([]bool{ true, true })
	}
	{
		//flip a bool
		fmt.Println("keeper.Flip(&should_be_true)")
		should_be_true := false
		keeper.Flip(&should_be_true)
		keeper.Assert(should_be_true)
	}
	{
		//filter empty strings (for example)
		fmt.Println(`keeper.Filter(&[]string{ "foo", "bar", "", "baz" })`)
		foo := []string{ "foo", "bar", "", "baz" }
		keeper.Filter(&foo, "")
		for _, thing := range foo {
			keeper.Assert(len(thing) > 0)
		}
	}
	{
		//take (get and remove) the first item from a slice
		fmt.Println("keeper.Shift(&[]int32{ 1, 2, 3, 4 })")
		foo := []int32{ 1, 2, 3, 4 }
		keeper.Assert(keeper.Shift(&foo) == 1)
		keeper.Assert(len(foo) == 3)
	}
	{
		//take (get and remove) the last item from a slice 
		fmt.Println("keeper.Pop(&[][]rune{ ... })")
		foo := [][]rune{
			{ 'f', 'o', 'o' },
			{ 'b', 'a', 'r' },
			{ 'b', 'a', 'z' },
			{ 'q', 'u', 'x' },
		}
		popped := keeper.Pop(&foo)
		keeper.Assert(slices.Equal(popped, []rune{ 'q', 'u', 'x' }))
		keeper.Assert(len(foo) == 3)
	}
	{
		//append (without returning a new slice) to a slice
		fmt.Println("keeper.Add(&[]float32{ 0.123, 0.456, 0.789 })")
		foo := []float32{ 0.123, 0.456, 0.789 }
		keeper.Add(&foo, 0.00)
		keeper.Assert(len(foo) == 4)
		keeper.Assert(foo[3] == 0.00)
	}
	{
		//drain (empty into a new slice) a slice
		fmt.Println("keeper.Drain(&[]byte{ 't', 'h', 'u', 'd' })")
		foo := []byte{ 't', 'h', 'u', 'd' }
		old_len := len(foo)
		old := foo
		bar := keeper.Drain(&foo)
		keeper.AssertMany([]bool{
			len(foo) == 0,
			len(bar) == old_len,
			slices.Equal(old, bar),
		})
	}
	{
		//drain into a different (pre-existing) slice, appending to it
		fmt.Println("keeper.DrainInto(&[]string{ ... }, &[]string{ ... })")
		into := []string{ "some", "existing", "stuff" }
		into_before := into
		into_len_before := len(into)

		from := []string{ "foo", "bar", "baz" }
		from_before := from
		from_len_before := len(from)

		keeper.DrainInto(&into, &from)

		keeper.AssertMany([]bool{
			len(from) == 0,
			len(into) == into_len_before + from_len_before,
			slices.Equal(from_before, into[into_len_before:]),
			slices.Equal(into_before, into[:into_len_before]),
		})
	}
	{
		fmt.Println("keeper.RunForMany(..., []*[]error)")
		foo := []*[]error{
			{ nil, errors.New("foo"), nil, nil },
			{ errors.New("bar"), errors.New("baz"), nil, errors.New("quz") },
		}
		keeper.RunForMany(
			func(thing *[]error) {
				keeper.Filter(thing, nil)
			},
			foo,
		)
		keeper.AssertManyFunc(
			func (errs *[]error) bool {
				return !slices.Contains(*errs, nil)
			},
			foo,
		)
	}
}
