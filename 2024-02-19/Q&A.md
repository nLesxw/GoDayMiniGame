# Question:
**code.go的输出会是什么？**

# Answer:
**正确答案:1 5 1**

- f1闭包修改了r的值,所以结果是1
- f2闭包修改了t的值,对r没有任何影响,所以结果是5
- f3闭包修改了r的值,但是func(r int)参数中的值,非返回值r的值,所以结果是1