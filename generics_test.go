package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Wrapped struct {
	Str Opt[string]
}

func TestOptStr(t *testing.T) {
	var o OptStr
	assert.Equal(t, "bar", o.Or("bar"), "Or")
	assert.False(t, o.IsSet(), "expected not set")
	assert.False(t, o.Set, "expected not set")
	assert.Equal(t, "", o.Value, "expected empty string")

	o.SetTo("foo")
	assert.True(t, o.Set, "expected set")
	assert.Equal(t, "foo", o.Value, "expected value")
	assert.True(t, o.IsSet(), "expected set")
	if v, ok := o.Get(); !ok || v != "foo" {
		t.Fatalf("expected foo, got %v", v)
	}
	assert.Equal(t, "foo", o.Or("bar"), "Or")

	o.Reset()
	assert.False(t, o.IsSet(), "expected not set")
	assert.False(t, o.Set, "expected not set")
	assert.Equal(t, "", o.Value, "expected empty string")
}

func TestWrapped(t *testing.T) {
	t.Run("Set", func(t *testing.T) {
		var w Wrapped

		var v OptStr
		v.SetTo("foo")

		w.Str = v

		assert.True(t, w.Str.Set, "expected set")
	})
	t.Run("Positive", func(t *testing.T) {
		var w Wrapped
		assert.False(t, w.Str.IsSet(), "expected not set")
		assert.False(t, w.Str.Set, "expected not set")
		assert.Equal(t, "", w.Str.Value, "expected empty string")

		w.Str.SetTo("foo")
		assert.True(t, w.Str.Set, "expected set")
		assert.Equal(t, "foo", w.Str.Value, "expected value")
		assert.True(t, w.Str.IsSet(), "expected set")
		if v, ok := w.Str.Get(); !ok || v != "foo" {
			t.Fatalf("expected foo, got %v", v)
		}
		assert.Equal(t, "foo", w.Str.Or("bar"), "Or")

		w.Str.Reset()
		assert.False(t, w.Str.IsSet(), "expected not set")
		assert.False(t, w.Str.Set, "expected not set")
		assert.Equal(t, "", w.Str.Value, "expected empty string")
	})
}
