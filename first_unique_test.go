package firstunique

import "testing"

func TestExample1(t *testing.T) {
    fu := Constructor([]int{2, 3, 5})

    assertEqual(t, 2, fu.ShowFirstUnique())

    fu.Add(5)
    assertEqual(t, 2, fu.ShowFirstUnique())

    fu.Add(2)
    assertEqual(t, 3, fu.ShowFirstUnique())

    fu.Add(3)
    assertEqual(t, -1, fu.ShowFirstUnique())
}

func TestExample2(t *testing.T) {
    fu := Constructor([]int{7, 7, 7, 7, 7, 7})

    assertEqual(t, -1, fu.ShowFirstUnique())

    fu.Add(7)
    fu.Add(3)
    fu.Add(3)
    fu.Add(7)
    fu.Add(17)

    assertEqual(t, 17, fu.ShowFirstUnique())
}

func TestExample3(t *testing.T) {
    fu := Constructor([]int{809})

    assertEqual(t, 809, fu.ShowFirstUnique())

    fu.Add(809)
    assertEqual(t, -1, fu.ShowFirstUnique())
}

func TestNewValueBecomesFirstUnique(t *testing.T) {
    fu := Constructor([]int{4, 4, 8, 8})

    assertEqual(t, -1, fu.ShowFirstUnique())

    fu.Add(10)
    assertEqual(t, 10, fu.ShowFirstUnique())

    fu.Add(3)
    assertEqual(t, 10, fu.ShowFirstUnique())
}

func TestNewDuplicateDoesNotChangeCurrentFirst(t *testing.T) {
    fu := Constructor([]int{1, 2, 3})

    fu.Add(2)
    assertEqual(t, 1, fu.ShowFirstUnique())

    fu.Add(1)
    assertEqual(t, 3, fu.ShowFirstUnique())
}

func TestRepeatedShowDoesNotChangeAnswer(t *testing.T) {
    fu := Constructor([]int{5, 6, 7})

    for i := 0; i < 5; i++ {
        assertEqual(t, 5, fu.ShowFirstUnique())
    }

    fu.Add(5)
    assertEqual(t, 6, fu.ShowFirstUnique())
}

func TestAllValuesBecomeNonUnique(t *testing.T) {
    fu := Constructor([]int{1, 2, 3})

    fu.Add(1)
    fu.Add(2)
    fu.Add(3)

    assertEqual(t, -1, fu.ShowFirstUnique())

    fu.Add(4)
    assertEqual(t, 4, fu.ShowFirstUnique())
}

func assertEqual(t *testing.T, want, got int) {
    t.Helper()
    if got != want {
        t.Fatalf("want %d, got %d", want, got)
    }
}
