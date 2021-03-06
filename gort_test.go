package gort_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fixpoint/gort"
)

func TestConcat(t *testing.T) {
	cases := []struct {
		E int
		V []int
	}{
		{0, []int{0, 0, 0}},

		{1, []int{1, 0, 0}},
		{1, []int{0, 1, 0}},
		{1, []int{0, 0, 1}},
		{1, []int{1, 1, 0}},
		{1, []int{0, 1, 1}},
		{1, []int{1, 0, 1}},
		{1, []int{1, 1, 1}},

		{-1, []int{-1, 0, 0}},
		{-1, []int{0, -1, 0}},
		{-1, []int{0, 0, -1}},
		{-1, []int{-1, -1, 0}},
		{-1, []int{0, -1, -1}},
		{-1, []int{-1, 0, -1}},
		{-1, []int{-1, -1, -1}},

		{1, []int{1, 1, 0}},
		{1, []int{1, -1, 0}},
		{-1, []int{-1, 1, 0}},
		{-1, []int{-1, -1, 0}},

		{1, []int{0, 1, 1}},
		{1, []int{0, 1, -1}},
		{-1, []int{0, -1, 1}},
		{-1, []int{0, -1, -1}},

		{1, []int{1, 0, 1}},
		{1, []int{1, 0, -1}},
		{-1, []int{-1, 0, 1}},
		{-1, []int{-1, 0, -1}},
	}
	for _, c := range cases {
		r := gort.Concat(c.V...)
		assert.Equalf(t, c.E, r, "Concat(%v)", c.V)
	}
}

func TestConcatToLess(t *testing.T) {
	cases := []struct {
		E bool
		V []int
	}{
		{false, []int{0, 0, 0}},

		{false, []int{1, 0, 0}},
		{false, []int{0, 1, 0}},
		{false, []int{0, 0, 1}},
		{false, []int{1, 1, 0}},
		{false, []int{0, 1, 1}},
		{false, []int{1, 0, 1}},
		{false, []int{1, 1, 1}},

		{true, []int{-1, 0, 0}},
		{true, []int{0, -1, 0}},
		{true, []int{0, 0, -1}},
		{true, []int{-1, -1, 0}},
		{true, []int{0, -1, -1}},
		{true, []int{-1, 0, -1}},
		{true, []int{-1, -1, -1}},

		{false, []int{1, 1, 0}},
		{false, []int{1, -1, 0}},
		{true, []int{-1, 1, 0}},
		{true, []int{-1, -1, 0}},

		{false, []int{0, 1, 1}},
		{false, []int{0, 1, -1}},
		{true, []int{0, -1, 1}},
		{true, []int{0, -1, -1}},

		{false, []int{1, 0, 1}},
		{false, []int{1, 0, -1}},
		{true, []int{-1, 0, 1}},
		{true, []int{-1, 0, -1}},
	}
	for _, c := range cases {
		r := gort.ConcatToLess(c.V...)
		assert.Equalf(t, c.E, r, "ConcatToLess(%v)", c.V)
	}

}

func TestConcatLazy(t *testing.T) {
	cases := []struct {
		E int
		V []int
	}{
		{0, []int{0, 0, 0}},

		{1, []int{1, 0, 0}},
		{1, []int{0, 1, 0}},
		{1, []int{0, 0, 1}},
		{1, []int{1, 1, 0}},
		{1, []int{0, 1, 1}},
		{1, []int{1, 0, 1}},
		{1, []int{1, 1, 1}},

		{-1, []int{-1, 0, 0}},
		{-1, []int{0, -1, 0}},
		{-1, []int{0, 0, -1}},
		{-1, []int{-1, -1, 0}},
		{-1, []int{0, -1, -1}},
		{-1, []int{-1, 0, -1}},
		{-1, []int{-1, -1, -1}},

		{1, []int{1, 1, 0}},
		{1, []int{1, -1, 0}},
		{-1, []int{-1, 1, 0}},
		{-1, []int{-1, -1, 0}},

		{1, []int{0, 1, 1}},
		{1, []int{0, 1, -1}},
		{-1, []int{0, -1, 1}},
		{-1, []int{0, -1, -1}},

		{1, []int{1, 0, 1}},
		{1, []int{1, 0, -1}},
		{-1, []int{-1, 0, 1}},
		{-1, []int{-1, 0, -1}},
	}
	for _, c := range cases {
		var expressions []func() int
		for _, v := range c.V {
			func(iv int) {
				expressions = append(expressions, func() int {
					return iv
				})
			}(v)
		}
		r := gort.ConcatLazy(expressions...)
		assert.Equalf(t, c.E, r, "ConcatLazy(%v)", c.V)
	}
}

func TestConcatToLessLazy(t *testing.T) {
	cases := []struct {
		E bool
		V []int
	}{
		{false, []int{0, 0, 0}},

		{false, []int{1, 0, 0}},
		{false, []int{0, 1, 0}},
		{false, []int{0, 0, 1}},
		{false, []int{1, 1, 0}},
		{false, []int{0, 1, 1}},
		{false, []int{1, 0, 1}},
		{false, []int{1, 1, 1}},

		{true, []int{-1, 0, 0}},
		{true, []int{0, -1, 0}},
		{true, []int{0, 0, -1}},
		{true, []int{-1, -1, 0}},
		{true, []int{0, -1, -1}},
		{true, []int{-1, 0, -1}},
		{true, []int{-1, -1, -1}},

		{false, []int{1, 1, 0}},
		{false, []int{1, -1, 0}},
		{true, []int{-1, 1, 0}},
		{true, []int{-1, -1, 0}},

		{false, []int{0, 1, 1}},
		{false, []int{0, 1, -1}},
		{true, []int{0, -1, 1}},
		{true, []int{0, -1, -1}},

		{false, []int{1, 0, 1}},
		{false, []int{1, 0, -1}},
		{true, []int{-1, 0, 1}},
		{true, []int{-1, 0, -1}},
	}
	for _, c := range cases {
		var expressions []func() int
		for _, v := range c.V {
			func(iv int) {
				expressions = append(expressions, func() int {
					return iv
				})
			}(v)
		}
		r := gort.ConcatToLessLazy(expressions...)
		assert.Equalf(t, c.E, r, "ConcatLazy(%v)", c.V)
	}
}

type Iface struct {
	Address string
	Ifname  string
	Ifindex int
}

func ExampleConcatToLess() {
	lessIfaceAsc := func(ifaces []Iface) func(int, int) bool {
		return func(i, j int) bool {
			a := ifaces[i]
			b := ifaces[j]
			return gort.ConcatToLess(
				compareString(a.Address, b.Address),
				compareString(a.Ifname, b.Ifname),
				compareInt(a.Ifindex, b.Ifindex),
			)
		}
	}

	ifaces := []Iface{
		{
			"192.168.1.1",
			"eth1",
			2,
		},
		{
			"192.168.1.1",
			"eth1",
			1,
		},
		{
			"192.168.1.1",
			"eth0",
			1,
		},
		{
			"192.168.1.0",
			"eth0",
			1,
		},
	}

	fmt.Println("Before")
	for _, iface := range ifaces {
		fmt.Println(iface)
	}

	sort.Slice(ifaces, lessIfaceAsc(ifaces))

	fmt.Println("After")
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
	// Output:
	// Before
	// {192.168.1.1 eth1 2}
	// {192.168.1.1 eth1 1}
	// {192.168.1.1 eth0 1}
	// {192.168.1.0 eth0 1}
	// After
	// {192.168.1.0 eth0 1}
	// {192.168.1.1 eth0 1}
	// {192.168.1.1 eth1 1}
	// {192.168.1.1 eth1 2}
}

func compareString(a, b string) int {
	return strings.Compare(a, b)
}

func compareInt(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
