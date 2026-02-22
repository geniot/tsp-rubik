package main

func orPanic(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func IfInt(cond bool, vTrue int, vFalse int) int {
	if cond {
		return vTrue
	}
	return vFalse
}

func orPanicRes[T any](res T, err interface{}) T {
	orPanic(err)
	return res
}
