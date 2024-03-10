package utils

import (
	"cmp"
	"time"
)

func List[T any](args ...T) []T {
	return append(make([]T, 0), args...)
}

func Map[T, U any](s []T, f func(T, int) U) []U {
	ns := make([]U, len(s))
	for i := range s {
		ns[i] = f(s[i], i)
	}
	return ns
}

func Accumulate[T, U any](s []T, f func(T, U) U, d U) U {
	acc := d
	for i := range s {
		acc = f(s[i], acc)
	}
	return acc
}

func Filter[T any](s []T, f func(T, int) bool) []T {
	ns := make([]T, 0)
	for i := range s {
		if f(s[i], i) {
			ns = append(ns, s[i])
		}
	}
	return ns
}

// func Join[T fmt.Stringer](slice []T, sep string) string {
// 	return strings.Join(Map(slice, func(item T, i int) string {
// 		return item.String()
// 	}), sep)
// }

func Max[T cmp.Ordered](s []T) int {
	if len(s) == 0 {
		return -1
	}

	max := 0
	for i := range s {
		if s[max] < s[i] {
			max = i
		}
	}
	return max
}

func WorkTime[T any](f func() T) (time.Time, time.Time, T) {
	startedAt := time.Now()
	res := f()
	finishedAt := time.Now()

	return startedAt, finishedAt, res
}
