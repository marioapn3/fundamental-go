package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// unit test failed
		panic("Result is not Hello Eko")
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be Hello Eko")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be Hello Eko")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on MacOS")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be Hello Eko")
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "Result must be Hello Eko")
	})
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result, "Result must be Hello Budi")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
	b.Run("Kurniawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurniawan")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Mario",
			request:  "Mario",
			expected: "Hello Mario",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
