package suffixarray_test

import (
	"slices"
	"os"
	"bytes"
	"testing"
	"github.com/BobuSumisu/aho-corasick"
	simplesuffixarray "github.com/MarcinCiura/AT-lab/5/suffixarray"
	//"github.com/MarcinCiura/AT-lab/5/suffixtree"
	"index/suffixarray"
)

const GENES = "geny.txt"
const MDNA = "mdna.txt"


func TestLookupAll(t *testing.T) {
	sa := simplesuffixarray.New([]byte("ananas"))
	tests := []struct {
		in   string
		want []int
	}{
		{"ana", []int{0, 2}},
		{"nan", []int{1}},
		{"as", []int{4}},
		{"na", []int{1, 3}},
		{"nana", []int{1}},
		{"ananas", []int{0}},
		{"ananasx", []int{}},
	}
	for _, tt := range tests {
		got := sa.LookupAll([]byte(tt.in))
		if !slices.Equal(got, tt.want) {
			t.Errorf("LookupAll(%q) = %v; want %v", tt.in, got, tt.want)
		}
	}
}

func BenchmarkAhoCorasickBuild(b *testing.B) {
	genes, err := os.ReadFile(GENES)
	if err != nil {
		b.Fatal(err)
	}
	genesArray := bytes.Split(genes, []byte("\n"))
	var genesString []string
    for _, b := range genesArray {
        genesString = append(genesString, string(b))
    }
	builder := ahocorasick.NewTrieBuilder()
	builder.AddStrings(genesString)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie := builder.Build()
		_ = trie
	}
}

func BenchmarkAhoCorasickSearch(b *testing.B) {
	genes, err := os.ReadFile(GENES)
	if err != nil {
		b.Fatal(err)
	}
	mdna, err := os.ReadFile(MDNA)
	if err != nil {
		b.Fatal(err)
	}
	genesArray := bytes.Split(genes, []byte("\n"))
	var genesString []string
    for _, b := range genesArray {
        genesString = append(genesString, string(b))
    }
	builder := ahocorasick.NewTrieBuilder()
	builder.AddStrings(genesString)
	trie := builder.Build()
	mdna_str := string(mdna)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.MatchString(mdna_str)
	}
}

func BenchmarkSimpleSuffixArrayBuild(b *testing.B) {
	mdna, err := os.ReadFile(MDNA)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sa := simplesuffixarray.New(mdna)
		_ = sa
	}
}

func BenchmarkSimpleSuffixArraySearch(b *testing.B) {
	genes, err := os.ReadFile(GENES)
	if err != nil {
		b.Fatal(err)
	}
	mdna, err := os.ReadFile(MDNA)
	if err != nil {
		b.Fatal(err)
	}
	genesArray := bytes.Split(genes, []byte("\n"))
	sa := simplesuffixarray.New(mdna)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, gene := range genesArray {
			sa.LookupAll(gene)
		}
	}
}

func BenchmarkLibrarySuffixArrayBuild(b *testing.B) {
	mdna, err := os.ReadFile(MDNA)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sa := suffixarray.New(mdna)
		_ = sa
	}
}

func BenchmarkLibrarySuffixArraySearch(b *testing.B) {
	genes, err := os.ReadFile(GENES)
	if err != nil {
		b.Fatal(err)
	}
	mdna, err := os.ReadFile(MDNA)
	if err != nil {
		b.Fatal(err)
	}
	sa := suffixarray.New(mdna)
	genesArray := bytes.Split(genes, []byte("\n"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, gene := range genesArray {
			sa.Lookup(gene, 1)
		}
	}
}

