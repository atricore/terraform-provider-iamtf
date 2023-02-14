package jossoappi

func IPtrString(s interface{}) *string {
	if _, ok := s.(string); ok {
		return PtrString(s.(string))
	}
	return nil
}

func IPtrBool(s interface{}) *bool {
	if _, ok := s.(bool); ok {
		return PtrBool(s.(bool))
	}
	return nil
}

func AsBool(i interface{}, d bool) bool {
	if i == nil {
		return d
	}

	switch v := i.(type) {
	case bool:
		return v
	case *bool:
		return *v

	}

	return d
}

func AsInt32Def(i interface{}, d int32, zeroAsNil bool) int32 {
	if i == nil {
		return d
	}

	var result int32

	switch v := i.(type) {
	case int32:
		result = v
	case int:
		result = int32(v)
	case int64:
		result = int32(v)
	case float32:
		result = int32(v)
	case float64:
		result = int32(v)
	}

	if zeroAsNil && result == 0 {
		return d
	}

	return result

}

func AsInt32(i interface{}, d int32) int32 {
	if i == nil {
		return d
	}

	switch v := i.(type) {
	case int32:
		return v
	case int:
		return int32(v)
	case int64:
		return int32(v)
	case float32:
		return int32(v)
	case float64:
		return int32(v)
	}

	return d

}

func AsInt64(i interface{}, d int64) int64 {
	if i == nil {
		return d
	}

	switch v := i.(type) {
	case int32:
		return int64(v)
	case int:
		return int64(v)
	case int64:
		return v
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	}

	return d

}

// Returns the received value 'casted' as a string. If the value is nil, returs the default.
func AsStringDef(value interface{},
	defautlValue string,
	emptyAsNil bool) string {
	if value == nil {
		return defautlValue
	}

	s := value.(string)
	if emptyAsNil && s == "" {
		return defautlValue
	}

	return s
}

func AsString(i interface{}, d string) string {
	return AsStringDef(i, d, false)
}

func AsStringArr(i interface{}) []string {
	if i == nil {
		var vec []string
		return vec
	}

	l := i.([]interface{})
	if len(l) == 0 {
		var v []string
		return v
	}

	return i.([]string)
}
