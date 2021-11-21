package util

func Jint(i interface{}) int {
	return int(i.(float64))
}

func Jint64(i interface{}) int64 {
	return int64(i.(float64))
}

func Jnint(i interface{}) *int {
	if i == nil {
		return nil
	}

	ni := new(int)
	*ni = int(i.(float64))
	return ni
}

func Jaf64(i interface{}) (a []float64) {
	ai := i.([]interface{})

	a = make([]float64, len(ai))
	for n, e := range ai {
		a[n] = e.(float64)
	}
	return
}
