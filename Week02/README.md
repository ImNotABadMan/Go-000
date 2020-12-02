# 学习笔记


### 1. Error 和Exeption 的区别
Error is Values : 所有的错误都需要处理


### 2. error的处理方式
error处理成堆栈


在初始化程序，什么是强依赖，什么是弱依赖
强依赖抛panic，还有配置，不会导致在不正确的上下文进行

阻塞和不阻塞 rpc

Error和Exception
exception需要被捕捉，然后可能回在外层被catch处理，
error由本函数控制，作为返回值

包级别的error，预定义，业务全局，静态

### 3. Error Types:
#### Sentinel Error
> 预定义的特定错误
> 最不灵活，需要调用者判断error类型，另一方面在API方面，增加了API的表面积
> 也形成了包之间的依赖
>

#### Error Types
> 实现了error接口的结构体，可断言类型，获取结构体包含的内容
> 比sentinel error好，但还是会造成依赖，因为需要断言类型才可以获取error信息
>
#### Opaque errors
> 对外暴露方法判断是否是error，比抛错误好，包里面想怎么处理就怎么处理，
不会造成外面需要依赖报的error类型。

### 4. Handling Error:

error接口
最后抛出来，不会一直if, 写法优雅。代码简洁

wrap errors
> 只处理异常一次
> 不能对返回值的内容做假设
>
> 方法处理异常，吞掉，不往上抛，返回值是默认值或者空。
保证100%完整性
之后不再报告错误，
>
> 错误一定要日志记录
最上层打日志

好用的error包：github.com/pkg/errors

正确使用wrap errors
> 业务代码，调用第三方，基础库，自己的库，wrap，把根因保存
基础库最上层写的不wrap，不然上层wrap，即使两次了
不打算处理的就是可以wrap
>
> 处理过的错误，不应该往上抛，记录日志都是处理
error.Cause获取根因可以与sentinel error作比较


首次报错需要wrap一下

GO1.13起error引入了Unwrap方法，
> Is方法判断类型是否相等，错误类型实现了Is接口，就与当前实现了Is接口的error作比较
>

#### 额外：
> 学习第一手资料，外国，官方
>
> 野生go挂了
这样写不会造成野生
Go(x func()) {
//xxxx
}
>
> 代码写法应该是一条直线，流水线，不是缩进很多，难以阅读


