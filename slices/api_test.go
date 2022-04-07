package slices

import (
	"strings"
	"testing"

	"github.com/jwhittle933/gonads/option"
)

func strContains(contains string) func(s string, _ int, _ []string) bool {
	return func(s string, _ int, _ []string) bool {
		return strings.Contains(s, contains)
	}
}

func strRemove(remove string) func(s string, _ int, _ []string) string {
	return func(s string, _ int, _ []string) string {
		return strings.ReplaceAll(s, remove, "")
	}
}

func lessThan(compare int) func(i int, _ int, _ []int) bool {
	return func(i, _ int, _ []int) bool {
		return i <= compare
	}
}

func multiplyBy(factor int) func(i int, _ int, _ []int) int {
	return func(i, _ int, _ []int) int {
		return i * factor
	}
}

func boolFnComposer[T comparable](expected bool) BoolFn[T] {
	return func(T, int, []T) bool {
		return expected
	}
}

func assertLen[T comparable](t *testing.T, want int, got []T) {
	if want != len(got) {
		t.Errorf("Expecting len() == %d, got %d", want, len(got))
	}
}

func assertSliceEquals[T comparable](t *testing.T, want, got []T) {
	if len(want) != len(got) {
		t.Errorf("Expecting len of %d, got %d", len(want), len(got))
	}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("Expecting %+v at index %d, got %+v", want[i], i, len(got))
		}
	}
}

func assertEqual(t *testing.T, want, got bool) {
	if want != got {
		t.Errorf("Expecting %t, got %t", want, got)
	}
}

func assertOption[T comparable](t *testing.T, want, got option.Option[T]) {
	if want.IsSome() && got.IsSome() {
		if wantVal, gotVal := want.Unwrap(), got.Unwrap(); wantVal != gotVal {
			t.Errorf("Want %+v, got %+v", wantVal, gotVal)
		}
	}

	if (want.IsNone() && got.IsSome()) || (want.IsSome() && got.IsNone()) {
		t.Error("Option types do not match")
	}
}

func TestFilter(t *testing.T) {
	t.Run(`Filters all strings without "cat"`, func(t *testing.T) {
		original := From("caterpillar", "cat", "battle", "incanto")
		assertLen(t, 2, Filter(original, strContains("cat")))
	})

	t.Run("Filters all ints over 100", func(t *testing.T) {
		original := From(10, 100, 150, 1000, 3, 4, 50)
		assertLen(t, 5, Filter(original, lessThan(100)))
	})
}

func TestMap(t *testing.T) {
	t.Run(`should strip "cat from each item"`, func(t *testing.T) {
		original := From("caterpillar", "cat", "battle", "incanto")
		assertSliceEquals(
			t,
			From("erpillar", "", "battle", "incanto"),
			Map(original, strRemove("cat")),
		)
	})

	t.Run("multiples each item by 10", func(t *testing.T) {
		original := From(10, 1, 100, 50)
		assertSliceEquals(
			t,
			From(100, 10, 1000, 500),
			Map(original, multiplyBy(10)),
		)
	})
}

func TestSome(t *testing.T) {
	t.Run(`Should return true if string contains "cat"`, func(t *testing.T) {
		original := From("caterpillar", "cat", "battle", "incanto")
		assertEqual(t, true, Some(original, strContains("cat")))
	})

	t.Run(`Should return false if no string contains "apples"`, func(t *testing.T) {
		original := From("caterpillar", "cat", "battle", "incanto")
		assertEqual(t, false, Some(original, strContains("apples")))
	})

	t.Run(`Should return true if int is <= 100`, func(t *testing.T) {
		original := From(10, 100, 150, 1000, 3, 4, 50)
		assertEqual(t, true, Some(original, lessThan(100)))
	})

	t.Run(`Should return false if no int is < 1`, func(t *testing.T) {
		original := From(10, 100, 150, 1000, 3, 4, 50)
		assertEqual(t, false, Some(original, lessThan(1)))
	})
}

func TestEvery(t *testing.T) {
	t.Run(`Should return true if every string contains "cat"`, func(t *testing.T) {
		original := From("caterpillar", "cat", "cattle", "concat")
		assertEqual(t, true, Every(original, strContains("cat")))
	})

	t.Run(`Should return false if some strings do not contain "apples"`, func(t *testing.T) {
		original := From("caterpillar", "cat", "apples", "incanto")
		assertEqual(t, false, Every(original, strContains("apples")))
	})

	t.Run(`Should return true if every int is <= 100`, func(t *testing.T) {
		original := From(10, 99, 50, 11, 3, 4, 22)
		assertEqual(t, true, Every(original, lessThan(100)))
	})

	t.Run(`Should return false if some ints are not is <= 100`, func(t *testing.T) {
		original := From(10, 100, 50, 1000, 3, 4, 50)
		assertEqual(t, false, Every(original, lessThan(100)))
	})
}

func TestFind(t *testing.T) {
	t.Run("Should find the item", func(t *testing.T) {
		item := Find(From("capable", "cattle"), strContains("cat"))

		if item.IsNone() {
			t.Error("Expected Option(Some), got None")
		}

		val := item.Unwrap()
		if val != "cattle" {
			t.Errorf("Expected 'cattle', got %s", val)
		}
	})

	t.Run("Should not find the item", func(t *testing.T) {
		item := Find(From(11, 1500, 22), lessThan(10))

		if item.IsSome() {
			t.Error("Expected Option(None), got Some")
		}
	})
}

func TestFindIndex(t *testing.T) {
	t.Run("Should find the item", func(t *testing.T) {
		item := FindIndex(From("capable", "cattle"), strContains("cat"))

		if item.IsNone() {
			t.Error("Expected Option(Some), got None")
		}

		val := item.Unwrap()
		if val != 1 {
			t.Errorf("Expected '1', got %d", val)
		}
	})

	t.Run("Should not find the item", func(t *testing.T) {
		item := FindIndex(From(11, 1500, 22), lessThan(10))

		if item.IsSome() {
			t.Error("Expected Option(None), got Some")
		}
	})
}

func TestLast(t *testing.T) {
	t.Run("Should return Some", func(t *testing.T) {
		assertOption(t, option.Some(0), Last(From(0)))
	})

	t.Run("Should return None", func(t *testing.T) {
		assertOption(t, option.None[int](), Last(From[int]()))
	})
}

func TestFirst(t *testing.T) {
	t.Run("Should return Some", func(t *testing.T) {
		assertOption(t, option.Some(0), First(From(0, 1, 2, 3)))
	})

	t.Run("Should return None", func(t *testing.T) {
		assertOption(t, option.None[int](), First(From[int]()))
	})
}

func TestAt(t *testing.T) {
	t.Run("Should return Some", func(t *testing.T) {
		assertOption(t, option.Some(1), At(From(0, 1, 2, 3), 1))
	})

	t.Run("Should return None", func(t *testing.T) {
		assertOption(t, option.None[int](), At(From(1, 2, 3), 10))
	})
}

func TestPrepend(t *testing.T) {
	t.Run("Should insert 1000 to the beginning of the list", func(t *testing.T) {
		original := From(0, 1, 2, 3, 4)
		actual := Prepend(original, 1000)
		if len(actual) != 6 {
			t.Errorf("Expecting len() == 6, got %d", len(actual))
		}

		if actual[0] != 1000 {
			t.Errorf("Expecting actual[0] == 1000, got %d", actual[0])
		}
	})

	t.Run("Should return a new list if nil is given", func(t *testing.T) {
		actual := Prepend(nil, 1000)
		if len(actual) != 1 {
			t.Errorf("Expecting len() == 6, got %d", len(actual))
		}

		if actual[0] != 1000 {
			t.Errorf("Expecting actual[0] == 1000, got %d", actual[0])
		}
	})
}

func TestEmpty(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		if !Empty([]int{}) {
			t.Error("List is empty!")
		}
	})

	t.Run("Should return false", func(t *testing.T) {
		if Empty([]int{1, 2, 3}) {
			t.Error("List is not emtpy!")
		}
	})
}

func TestSlice_ToMap(t *testing.T) {
	t.Run("Should create a map", func(t *testing.T) {
		actual := From("first", "second", "third", "fourth", "fifth").ToMap()

		t.Run("it should have len of 5", func(t *testing.T) {
			if len(actual) != 5 {
				t.Errorf("Want len() == 5, got %d", len(actual))
			}
		})
	})
}
