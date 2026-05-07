package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
