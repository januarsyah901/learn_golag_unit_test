package helper

import (
	"fmt"
	"runtime"
	// "structs"
	"testing"

	assert "github.com/stretchr/testify/assert"   // ini pake fail()
	require "github.com/stretchr/testify/require" // ini pake failNow()
)

func BenchmarkSub(b *testing.B) {
	// TODO: Initialize
	b.Run("janu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("janu")
		}
	})
	b.Run("akbar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("akbar")
		}
	})
}
func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("januarsyah")
	}
}

func BenchmarkTableHello(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "janu",
			request: "janu",
		},
		{
			name:    "akbar",
			request: "akbar",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.request)
			}
		})

	}
}
 

func BenchmarkSayGoodBye(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("akbar")
	}
}
func TestTableHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "janu",
			request:  "janu",
			expected: "Hello janu",
		},
		{
			name:     "budi",
			request:  "budi",
			expected: "Hello budi",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("DB Connected")

	m.Run()

	// after
	fmt.Println("CRUD Success")
}
func TestSubTest(t *testing.T) {
	t.Run("TestSub", func(t *testing.T) {
		result := SayHello("SubTest")
		assert.Equal(t, "Hello SubTest", result, "result must be 'Hello SubTest'")
		fmt.Println("TestSub Pass")
	})

	t.Run("TestSub2", func(t *testing.T) {
		result := SayGoodBye("SubTest2")
		assert.Equal(t, "Goodbye SubTest2", result, "result must be 'Goodbye SubTest2'")
		fmt.Println("TestSub2 Pass")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Tidak berjalan pada sistem operasi ini")
	}
}

/*
FailNow dan Fatal digunakan untuk menghentikan pengujian saat ini dan melaporkan kesalahan.
Perbedaan antara keduanya adalah bahwa FailNow akan melanjutkan pengujian lainnya setelahnya,
sedangkan Fatal akan menghentikan semua pengujian yang sedang berlangsung.
*/

// func TestSayHello(t *testing.T) {
// 	result := SayHello("januarsyah")

// 	if result != "Hello januarsyah" {
// 		// error
// 		t.Error("result must be 'januarsyah'")
// 	}
// 	fmt.Println("TestSayHello Pass")
// }

func TestSayHelloAssert(t *testing.T) {
	result := SayHello("januarsyah")
	assert.Equal(t, "Hello januarsyah", result, "result must be 'Hello januarsyah'")
	fmt.Println("TestSayHelloAssert Pass")
}

func TestSayGoodBye(t *testing.T) {
	result := SayGoodBye("surya")

	if result != "Goodbye surya" {
		// error
		t.Fatal("must be 'surya'")
	}
	fmt.Println("TestSayGoodBye Pass")
}
func TestSayGoodByeRequire(t *testing.T) {
	result := SayGoodBye("surya")
	require.Equal(t, "Goodbye surya", result, "result must be 'Goodbye surya'")
	fmt.Println("TestSayGoodByeRequire Pass")
}

// create Table test
