package remove_element

import "testing"

func Test_removeElement(t *testing.T) {
	t.Run("Test remove element", func(t *testing.T) {
		data := []int{1, 2, 4, 5, 6, 6, 7, 6, 7, 5, 12}
		got := removeElement(data, 5)
		want := 9
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
