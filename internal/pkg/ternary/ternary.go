package ternary

// Ternary[T any]
// @Description: 实现三元表达式
// @param condition 表达式条件
// @param a 结果a
// @param b 结果b
// @return T
func Ternary[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}
