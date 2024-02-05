# Question:
**对 add() 函数调用正确的是？**

`
func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
`

- A.`add(1, 2)`
- B.`add(1, 3, 7)`
- C.`add([]int{1, 2})`
- D.`add([]int{1, 3, 7}...)`


# Answer:
**正确答案:ABD**

可变函数