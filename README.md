**goroutine【协程】**

    多个goroutine可以并行执行，相互不受影响
    net/http 库内置并发
    代码实现： go xxx()


**channel【通道】**
    
    数据结构， 可以让goroutine之间进行安全的数据通信（避免内存共享问题）
    保证： 同一时刻只会有一个 goroutine 修改数据    
    
    goroutine之间通过 channel进行数据传输，这个传输过程是 同步的，不存在异步场景下的问题； 
   
        需要强调的是，通道并不提供跨 goroutine 的数据访问保护机制。如果通过通道传输数据的 一份副本，那么每个 goroutine 都持有一份副本，各自对自己的副本做修改是安全的。
    当传输的 是指向数据的指针时，如果读和写是由不同的 goroutine 完成的，每个 goroutine 依旧需要 "额外的同步动作"。


无继承
组合设计模式（composition）
使用接口作为代码复用的基础模块
