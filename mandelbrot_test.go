package main

import "testing"

func TestEscape(t *testing.T) {
	data := []struct {
		c complex128
		e int
	}{
		{-0.5-0.5i, 63},
		{-0.75+0.1i, 32},
		{0.75-0.25i, 2},
		{-2.01, 0},
		{1.99+1.99i, 0},
		{2.2i, 0},
	}

	for _, d := range data {
		r := escape(d.c)
		if d.e != r {
			t.Errorf("TestEscape: unexpected result for %g: %d != %d", d.c, d.e, r)
		}
	}
}

func BenchmarkEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		escape(0.35+0.21i)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < 10; i++ {
		center := 1.5+0i + 0.1+0.1i * complex(float64(i), 2*float64(i));
		generate(320, 200, center, 0.5)
	}
}

// Below is an example of an alternative "escape" implementation.  This might work faster on
// some devices (switching to 64 floating point would possibly help more).
// switch this function out for the original escape and rerun the benchmark to test.
/*
func escape(c complex128) int {
	z := c
	for i := 0; i < _MaxEscape-1; i++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			return i
		}
		z = z*z + c
	}
	return _MaxEscape - 1
}
*/
