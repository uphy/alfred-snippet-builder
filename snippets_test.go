package snippet

import "testing"

func TestEscapeFilename(t *testing.T) {
	expect(t, escapeFilename("ab"), "ab")
	expect(t, escapeFilename("a\nb"), "a b")
	expect(t, escapeFilename(`a"#%&*:<>?/{|}b`), "a             b")
}

func expect(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
