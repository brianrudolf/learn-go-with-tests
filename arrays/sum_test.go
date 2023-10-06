package arrays

import "testing"

func assertSum(t testing.TB, got, expected int, numbers []int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %d, expected %d, given %v", got, expected, numbers)
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
