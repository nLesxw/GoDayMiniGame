# Question:
**Code.go的输出会是什么？**


# Answer:
**正确答案:输出5**
这个例子中，hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5.