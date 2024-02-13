package arrays

import (
	"reflect"
	"testing"
)

// Wrote when there were multiple tests, this is extra when we only have one test
func assertSum(t testing.TB, got, expected int, numbers []int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %d, expected %d, given %v", got, expected, numbers)
	}
}

func assertSumAll(t testing.TB, got, expected []int, slices ...[]int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %d, expected %d, given %v", got, expected, slices)
	}
}

func TestSum(t *testing.T) {

	t.Run("sum a collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		assertSum(t, got, expected, numbers)
	})

}

func TestSumAll(t *testing.T) {
	t.Run("sum a varying number of slices that contain a varying number of integers and returns a slice with the sums of each input slice",
		func(t *testing.T) {
			a := []int{1, 2, 3}
			b := []int{1, 2, 3, 4, 5}

			got := SumAll(a, b)
			want := []int{6, 15}

			assertSumAll(t, got, want, a, b)

		})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum slice tails (all numbers but the head)",
		func(t *testing.T) {
			a := []int{}
			b := []int{1, 2, 3, 4, 5}

			got := SumAllTails(a, b)
			want := []int{0, 14}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %d, want %d, input: %v, %v", got, want, a, b)
			}
		})
}
