package array

import "testing"

func Sum(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}

func SumAll(array ...[]int) int {
	a := 0
	for _, nums := range array {
		for _, v := range nums {
			a += v
		} 
	}
	return a
}


/*func Testarray(t *testing.T) {
	got := sum([5]int{1,2,3,4,5})
	want := 15

	if got != want {
		t.Errorf("got != want")
	}
}*/




func TestSum(t *testing.T) {

    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}

        got := Sum(numbers)
        want := 15

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

	t.Run("sum all", func(t *testing.T) {
        numbers1 := []int{1, 2, 3}
		numbers2 := []int{1, 2, 3}
        got := SumAll(numbers1, numbers2)
        want := 12

        if got != want {
            t.Errorf("got %d want %d", got, want)
        }
    })
}