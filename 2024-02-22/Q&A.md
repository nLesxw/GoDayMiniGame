# Question:
**code.go的输出会是什么？**

- A.`F M D`
- B.`D F M`
- C.`F D Mi`

# Answer:
**正确答案:C**

**解析**

被调用函数里的 defer 语句在返回之前就会被执行，所以输出顺序是 F D M。