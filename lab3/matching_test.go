package matching

import (
	"os"
	"slices"
	"testing"
)

var shortWord []byte = []byte("Julia")
var longWord  []byte = []byte("Ministerstwo")
var filename string = "orwell-rok-1984.txt"

func TestPreprocess(t *testing.T) {
	data := []string{
		"aaaaaaa",
		"pies",
		"dźwiedź",
		"owocowo",
		"indianin",
		"nienapełnienie",
	}
	for _, in := range data {
		got := Preprocess([]byte(in))
		want := SimplePreprocess([]byte(in))
		if !slices.Equal(got, want) {
			t.Errorf(`Preprocess(%#v) == %#v want %#v`,
				in, got, want)
		}
	}
}

func indices(pat, text []byte) []int {
	r := []int{}
	for i := 0; i+len(pat) <= len(text); i++ {
		if slices.Equal(text[i:i+len(pat)], pat) {
			r = append(r, i)
		}
	}
	return r
}


func TestNaive(t *testing.T) {
	pat := []byte("abc")
	text := []byte("abcabcabc")
	got := []int{}
	Naive(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf(`Naive(%#v, %#v) == %#v want %#v`, pat, text, got, want)
	}
}

func TestBackwardNaive(t *testing.T) {
	pat := []byte("abc")
	text := []byte("abcabcabc")
	got := []int{}
	BackwardNaive(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf(`BackwardNaive(%#v, %#v) == %#v want %#v`, pat, text, got, want)
	}
}

func TestBoyerMoore(t *testing.T) {
	pat := []byte("abc")
	text := []byte("abcabcabc")
	got := []int{}
	BoyerMoore(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf(`BoyerMoore(%#v, %#v) == %#v want %#v`, pat, text, got, want)
	}
}

func TestKMP(t *testing.T) {
	pat := []byte("abc")
	text := []byte("abcabcabc")
	got := []int{}
	KMP(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf(`KMP(%#v, %#v) == %#v want %#v`, pat, text, got, want)
	}
}

func TestKarpRabin(t *testing.T) {
	pat := []byte("abc")
	text := []byte("abcabcabc")
	got := []int{}
	KarpRabin(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf(`RabinKarp(%#v, %#v) == %#v want %#v`, pat, text, got, want)
	}
}

func TestShiftOr(t *testing.T) {
	pat := []byte("abc")
	text := []byte("abcabcabc")
	got := []int{}
	ShiftOr(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		t.Errorf(`ShiftOr(%#v, %#v) == %#v want %#v`, pat, text, got, want)
	}
}

func BenchmarkShortNaive(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := shortWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Naive(pat, text, func(int) {})
	}
}

func BenchmarkShortBackwardNaive(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := shortWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BackwardNaive(pat, text, func(int) {})
	}
}

func BenchmarkShortBoyerMoore(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := shortWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoyerMoore(pat, text, func(int) {})
	}
}

func BenchmarkShortKMP(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := shortWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMP(pat, text, func(int) {})
	}
}

func BenchmarkShortKarpRabin(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := shortWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KarpRabin(pat, text, func(int) {})
	}
}

func BenchmarkShortShiftOr(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := shortWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftOr(pat, text, func(int) {})
	}
}

func BenchmarkShortBoyerMooreFast(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := shortWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoyerMooreFast(pat, text, func(int) {})
	}
}

func BenchmarkLongNaive(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := longWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Naive(pat, text, func(int) {})
	}
}

func BenchmarkLongBackwardNaive(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := longWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BackwardNaive(pat, text, func(int) {})
	}
}

func BenchmarkLongBoyerMoore(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := longWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoyerMoore(pat, text, func(int) {})
	}
}

func BenchmarkLongKMP(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := longWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMP(pat, text, func(int) {})
	}
}

func BenchmarkLongKarpRabin(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := longWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KarpRabin(pat, text, func(int) {})
	}
}

func BenchmarkLongShiftOr(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := longWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftOr(pat, text, func(int) {})
	}
}

func BenchmarkLongBoyerMooreFast(b *testing.B) {
	text, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	pat := longWord
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoyerMooreFast(pat, text, func(int) {})
	}
}
