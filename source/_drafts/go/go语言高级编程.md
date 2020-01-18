[go语言高级编程](https://chai2010.cn/advanced-go-programming-book/)

### 函数、方法和接口
* init不是普通函数，可以定义有多个，所以也不能被其它函数调用
* 如果某个init函数内部用go关键字启动了新的goroutine的话，新的goroutine只有在进入main.main函数之后才可能被执行到。
* 每种类型对应的方法必须和类型的定义在同一个包中

### 面向并发的内存模型
* Goroutine采用的是半抢占式的协作调度，只有在当前Goroutine发生阻塞时才会导致调度；
* 并发编程的核心概念是同步通信