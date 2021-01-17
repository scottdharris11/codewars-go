package main

func IsTriangle(a, b, c int) bool {
	return !(a < 1 || b < 1 || c < 1 || a+b <= c || b+c <= a || a+c <= b)
}
