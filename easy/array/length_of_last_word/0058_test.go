package length_of_last_word

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	t.Run("Test-way-1", func(t *testing.T) {
		data := "Hello World"
		got := lengthOfLastWord(data)
		want := 5
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

}

func Test_lengthOfLastWordSecond(t *testing.T) {
	t.Run("Test-way-2", func(t *testing.T) {
		data := "Hello World"
		got := lengthOfLastWordSecond(data)
		want := 5
		if got != want {
			t.Error("GOT", got, "WNT:", want)
		}
	})

}
