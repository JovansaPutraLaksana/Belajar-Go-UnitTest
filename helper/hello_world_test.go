package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldWithTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{name: "Benchmark with Jovansa", request: "Jovansa"},
		{name: "Benchmark with empty", request: ""},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkHelloWorldWithSubTest(b *testing.B) {
	b.Run("Benchmark with Jovansa", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jovansa")
		}
	})
	b.Run("Benchmark with empty", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Jovansa")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{name: "Jovansa", request: "Jovansa", expected: "Hello, Jovansa!"},
		{name: "Empty", request: "", expected: "Hello, !"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Jovansa")
	if result != "Hello, Jovansa!" {
		t.Error("Result should be 'Hello, Jovansa!'")
	}
}

func TestHelloWorldWithEmpty(t *testing.T) {
	result := HelloWorld("")
	if result != "Hello, !" {
		t.Fatal("Result should be 'Hello, !'")
	}
}
func TestHelloWorldWithSubTest(t *testing.T) {
	t.Run("Test with Jovansa", func(t *testing.T) {
		result := HelloWorld("Jovansa")
		if result != "Hello, Jovansa!" {
			t.Fail()
		}
	})
	t.Run("Test with empty", func(t *testing.T) {
		result := HelloWorld("")
		if result != "Hello, !" {
			t.FailNow()
		}
	})
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Jovansa")
	assert.Equal(t, "Hello, Jovansa!", result, "Result should be 'Hello, Jovansa!'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("")
	require.Equal(t, "Hello, !", result, "Result should be 'Hello, !'")
}

func TestHelloWorldSkip(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip testing in short mode")
	}
	result := HelloWorld("Jovansa")
	assert.Equal(t, "Hello, Jovansa!", result, "Result should be 'Hello, Jovansa!'")
}

func TestMain(m *testing.M) {
	println("Before Unit Test")
	m.Run()
	println("After Unit Test")
}
