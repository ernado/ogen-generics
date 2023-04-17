package generics

type OptStr = Opt[string]

type Opt[T any] struct {
	Value T
	Set   bool
}

// IsSet returns true if OptMessage was set.
func (o Opt[T]) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *Opt[T]) Reset() {
	var v T
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *Opt[T]) SetTo(v T) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o Opt[T]) Get() (v T, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o Opt[T]) Or(d T) T {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type OptNil[T any] struct {
	Value T
	Set   bool
	Null  bool
}

// IsSet returns true if OptNil[T] was set.
func (o OptNil[T]) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNil[T]) Reset() {
	var v T
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNil[T]) SetTo(v T) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o OptNil[T]) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *OptNil[T]) SetToNull() {
	o.Set = true
	o.Null = true
	var v T
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNil[T]) Get() (v T, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if it does not.
func (o OptNil[T]) Or(d T) T {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}
