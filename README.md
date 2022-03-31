<style>
table {
margin: auto;
}
</style>

# Go语言并发之道学习笔记

<font color="#f29e8e">**写在前面的话：**</font> 近期一直在学Go语言，并发是必不可少的，看《Go语言并发之道》做个记录，方便回顾，里面很多内容也不完善，希望有所收获


## 1. 基本概念

- 并发和并行的区别
    - 本书的说法: 并发属于代码，并行属于一个运行中的程序。
    - 《Go语言区块链应用开发》书中的说法: 并行强调的是同一时刻上同时做事的能力，并发强调的是交替做不同事情的能力。并发是不同的代码块交替执行，并行是不同的代码块同时执行，并行一定要是多核CPU才行
- sync 和 channel 的区别
    - sync 对性能要求高，保护某个结构的内部状态，不关心操作的结果
    - channel 需要转让数据的所有权，协调多个逻辑判断
    - 尽量使用 channel


## 2. goroutine 是否改变闭包内变量的值

**先来看一段代码，输出是什么？**

```go
package main
import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup
func main() {
    salutation := "hello"
    wg.Add(1)
    go func() {
        defer wg.Done()
        salutation = "welcome"
    }()
    wg.Wait()
    fmt.Println(salutation)
}
```

- 输出结果: 
    > welcome
- 结论: 
    - `goroutine在他们所创建的相同地址空间内执行`


**我们更改一下main部分**

```go
func main() {
    for _, salutation := range []string{"hello", "greetings", "good day"} {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println(salutation)
        }()
    }
    wg.Wait()
}
```

- 输出结果: 
    > good day
        good day
        good day

- 结论: 
    > goroutine 在运行一个闭包，在闭包使用变量 salutation 时，字符串的迭代已经结束。goroutine 开始之前循环又很高的概率会退出，意味着变量 salutation 的值不在范围内。所以 salutation 会被转移到堆中。

**若要循环打出，正确方式如下:**

- 声明参数，将变量显示的映射到闭包中
- 将当前迭代的变量传递给闭包，创建了一个字符串结构的副本，从而确保当 goroutine 运行时，我们可以引用适当的字符串

```go
for _, salutation := range []string{"hello", "greetings", "good day"} {
    wg.Add(1)
    go func(s string) {
        defer wg.Done()
        fmt.Println(s)
    }(salutation)
}
wg.Wait()
```

## 3. 关键字

### 3.1 sync

#### 3.1.1 WaitGroup

- `Add: ` 表示一个 goroutine 开始
- `Done: ` 表明退出
- `Wait: ` 阻塞 main goroutine，直到所有 goroutine 表明已经退出

`WaitGroup` 可以视为一个并发安全的计数器: 调用通过传入的证书执行 `add` 方法增加计数器的增量，并调用 `Done` 方法进行递减。 `Wait` 阻塞，直到计数器为 `0`.

#### 3.1.2 Mutex 互斥锁

互斥，保护程序中临界区的一种方式。临界区是程序中需要独占访问资源的区域。为了使用一个资源， `channel` 通过通信共享内存，而 `Mutex` 通过开发人员的约定同步访问共享内存。

#### 3.1.3 cond 

一个 `goroutine` 的集合点，等待或发布一个 `event` 。管理调度队列空间。

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    c := sync.NewCond(&sync.Mutex{})
    queue := make([]interface{}, 0, 10)

    romoveFromQueue := func(delay time.Duration) {
        time.Sleep(delay)
        c.L.Lock()
        defer c.L.Unlock()
        queue = queue[1:]
        fmt.Println("Remove fomr queue")
        c.Signal()
    }

    for i := 0; i < 10; i++ {
        c.L.Lock()
        for len(queue) == 2 {
            c.Wait()
        }
        fmt.Println("Adding to a queue")
        queue = append(queue, struct{}{})
        go romoveFromQueue(1 * time.Second)
        c.L.Unlock()
    }
}
```

> 相比 `channel` ， `Cond` 类型的性能要高很多。与 `sync` 包中所含的大多数其他东西一样， `Cond` 的使用最好被限制在一个紧凑的范围中，或者是通过封装它的类型来暴露在更大的范围内。

#### 3.1.4 once

**下例会出现死锁**
```go
var onceA, onceB sync.Once
var initB func()

initA := func() { onceB.Do(initB)}
initB := func() { onceA.Do(initA)}
onceA.Do(initA)
```

`sync.Once` 就是为了防止多重初始化，唯一能保证的是函数只被调用一次。


#### 3.1.5 Pool(还未细细研究)

并发进程需要请求一个对象，但是在实例化之后很快地处理掉它们时，或者在这些对象的构造可能会对内存产生负面影响，这时最好使用 `Pool` 设计模式

### 3.2 channel

#### 3.2.1 基本操作

- 创建
    ```Go
    var dataStream chan interface{}
    dataStream = make(chan interface{}) // 无缓冲
    dataStream = make(chan interface{}, 4) // 有缓冲
    ```
- 声明单向 channel
    - 只能读取
        ```go
        var dataStream <-chan interface{}
        dataStream = make(<-chan interface{})
        ```
    - 只能发送
        ```go
        var dataStream chan<- interface{}
        dataStream = make(chan<- interface{})
        ```
- 关闭 channel
    - 关闭 channel 用来打开所有的 goroutine
        ```go
        close(channel)
        ```


#### 3.2.2 使用channel

<div class="center">

|  | nil | 打开且非空 | 打开但空 | 关闭的 | 只写 |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| Read | 阻塞 | 输出值 | 阻塞 | <默认值>，false | 编译错误 |

</div>
<div class="center">

|  | nil | 打开且填满 | 打开但不满 | 关闭的 | 只读 |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| Write | 阻塞 | 阻塞 | 写入值 | panic | 编译错误 |
</div>

<div class="center">

|  | nil | 打开且非空 | 打开但空 | 关闭的 | 只读 |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| close | panic | 关闭Channel；读取成功，直到通道耗尽，然后读取产生值的默认值 | 关闭Channel；读到生产者的默认值  | panic | 编译错误 |
</div>


- 拥有 channel 的 goroutine 应该具备
    - 实例化 channel
    - 执行写操作，或将所有权传递给另一个 goroutine
    - 关闭 channel
    - 执行在次列表的前三件事，并通过一个只读 channel 将它们暴露出来

```go
package main

import (
    "fmt"
)

func main() {
    // 返回的是一个只读channel，resultStream被隐式的转换为只读消费者
    chanOwner := func() <-chan int {
        // 实例化一个缓冲channel
        resultStream := make(chan int, 5)
        // 启用匿名的 goroutine
        go func() {
            // 确保执行完成后通道关闭
            defer close(resultStream)
            for i := 0; i < 5; i++ {
                resultStream <- i
            }
        }()
        return resultStream
    }

    resultStream := chanOwner()
    for result := range resultStream {
        fmt.Printf("Received: %d\n", result)
    }
    fmt.Println("Done Receiving!")
}
```

- 输出结果:
    > Received: 0
    Received: 1
    Received: 2
    Received: 3
    Received: 4
    Done Receiving!

#### 3.2.3 select

<font color="red">暂时跳过</font>


## 4. Go语言的并发模式

### 4.1 约束（不是很理解）

`"限制"` 是一个用来确保某信息在并发的过程中仅能被其中之一进程进行访问的简单且强大的技术

- 特定约束
- 词法约束

### 4.2 防止 goroutine 泄露

#### 4.2.1 goroutine 泄露 —— 读泄露
```go
package main

import (
    "fmt"
)

func main() {
    doWork := func(strings <-chan string) <-chan interface{} {
        completed := make(chan interface{})
        go func() {
            defer fmt.Println("doWork exited")
            defer close(completed)
            for s := range strings {
                fmt.Println(s)
            }
        }()
        return completed
    }

    doWork(nil)
    fmt.Println("done.")
}
```

- 输出结果
    > done.


`main goroutine` 将一个空的 `channel` 传递给了函数，因此内部的 `strings` 永远获取不到任何`string`，并且包含 `doWork` 函数的 `goroutine` 会一直在程序的生命周期内保持在内存中。

**如何关闭一个通道，即结束一个goroutine**

将父子 `goroutine` 进行成功整合的一种方法就是在父子 `goroutine` 之间建立一个 "信号通道"，让父 `goroutine` 可以向子 `goroutine` 发出取消信号。按照惯例，这个信号通常是一个名为 `done` 的只读 `channel`。父 `goroutine` 将该 `channel` 传递给子 `goroutine`，然后在想要取消子 `goroutine` 时关闭该 `channel`。

```go 
package main

import (
    "fmt"
    "time"
)

func main() {
    doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
        terminated := make(chan interface{})
        go func() {
            defer fmt.Println("doWork exited.")
            defer close(terminated)
            for {
                select {
                case s := <-strings:
                    fmt.Println(s)
                case <-done:
                    return
                }
            }
        }()
        return terminated
    }

    done := make(chan interface{})
    terminated := doWork(done, nil)

    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("canceling doWork goroutine...")
        close(done)
    }()

    <-terminated
    fmt.Println("Done.")
}
```

- 输出结果
    > canceling doWork goroutine...
doWork exited.
Done.


对比上述两段代码，可以发现第一段是没有进入到子 `goroutine` 中的，第二段通过传递额外 `channel` 来控制，没有造成死锁的原因是，加入两个 `goroutine` 之前，创建了第三个 `goroutine` 来在 `doWork` 执行 `1s` 后取消 `doWork` 中的 `goroutine`。这样就消除了 `goroutine` 泄露。

#### 4.2.2 goroutine 泄露 —— 写泄露

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    newRandStream := func() <-chan int {
        randStream := make(chan int)
        go func() {
            defer fmt.Println("newRandStream closure exited")
            defer close(randStream)
            for {
                randStream <- rand.Int()
            }
        }()
        return randStream
    }

    randStream := newRandStream()
    fmt.Println("3 random ints:")
    for i := 1; i <= 3; i++ {
        fmt.Printf("%d: %d\n", i, <-randStream)
    }
}
```

- 输出结果: 
    > 3 random ints:
1: 5577006791947779410
2: 8674665223082153551
3: 6129484611666145821

**如何关闭一个通道，即结束一个goroutine**

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    newRandStream := func(done <-chan interface{}) <-chan int {
        randStream := make(chan int)
        go func() {
            defer fmt.Println("newRandStream closure exited")
            defer close(randStream)
            for {
                select {
                case randStream <- rand.Int():
                case <-done:
                    return
                }
            }
        }()
        return randStream
    }
    done := make(chan interface{})
    randStream := newRandStream(done)
    fmt.Println("3 random ints:")
    for i := 1; i <= 3; i++ {
        fmt.Printf("%d: %d\n", i, <-randStream)
    }
    close(done)

    time.Sleep(1 * time.Second)
}
```

- 输出结果
    > 3 random ints:
1: 5577006791947779410
2: 8674665223082153551
3: 6129484611666145821
newRandStream closure exited

#### 4.3 or-channel（我跳过了）

#### 4.4 错误处理

先看一段代码

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
        responses := make(chan *http.Response)
        go func() {
            defer close(responses)
            for _, url := range urls {
                resp, err := http.Get(url)
                if err != nil {
                    fmt.Println(err)
                    continue
                }
                select {
                case <-done:
                    return
                case responses <- resp:
                }
            }
        }()
        return responses
    }

    done := make(chan interface{})
    defer close(done)
    urls := []string{"https://www.baidu.com", "https://badhost"}

    for res := range checkStatus(done, urls...) {
        fmt.Printf("Response: %v\n", res.Status)
    }
}
```

- 输出结果
    > Response: 200 OK
    Get "https://badhost": dial tcp: lookup badhost: no such host


```go
func main() {

    type Result struct {
        Error    error
        Response *http.Response
    }

    checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
        results := make(chan Result)
        go func() {
            defer close(results)
            for _, url := range urls {
                var result Result
                resp, err := http.Get(url)
                result = Result{Error: err, Response: resp}
                select {
                case <-done:
                    return
                case results <- result:
                }
            }
        }()
        return results
    }

    done := make(chan interface{})
    defer close(done)
    urls := []string{"https://www.baidu.com", "https://badhost"}

    for result := range checkStatus(done, urls...) {
        if result.Error != nil {
            fmt.Printf("error: %v", result.Error)
            continue
        }
        fmt.Printf("Response: %v\n", result.Response.Status)
    }
}
```

- 输出结果 (原书是谷歌肯定访问不到，所以我换成百度了)
    > Response: 200 OK
    error: Get "https://badhost": dial tcp: lookup badhost: no such host

是不是又去看上面的结果了？区别就多了个<font color="red"> error: </font>，其实代码想表达的是这样处理能够将错误处理从生产者 `goroutine` 中分离出来。

### 4.5 pipeline

#### 4.5.1 概念

不贴书中的例子了，直接说一下我的理解，就是函数的返回结果不需要给定特定参数，直接可以作为参数传递给其他函数，比如 `f1(f2(f3()))`，和python一样的机制。

#### 4.5.2 channel 构建 pipeline


**一个便利的生成器** （这部分代码没跑成功，求大佬告知原因，返回结果是地址）

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
        valueStream := make(chan interface{})
        go func() {
            defer close(valueStream)
            for {
                for _, v := range values {
                    select {
                    case <-done:
                        return
                    case valueStream <- v:
                    }
                }
            }
        }()
        return valueStream
    }

    take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
        takeStream := make(chan interface{})
        go func() {
            defer close(takeStream)
            for i := 0; i < num; i++ {
                select {
                case <-done:
                    return
                case takeStream <- valueStream:
                }
            }
        }()
        return takeStream
    }
    
    done := make(chan interface{})
    defer close(done)

    for num := range take(done, repeat(done, 1), 10) {
        fmt.Printf("%v", num)
    }
}
```

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
        takeStream := make(chan interface{})
        go func() {
            defer close(takeStream)
            for i := 0; i < num; i++ {
                select {
                case <-done:
                    return
                case takeStream <- valueStream:
                }
            }
        }()
        return takeStream
    }
    
    repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
        valueStream := make(chan interface{})
        go func() {
            defer close(valueStream)
            for {
                select {
                case <-done:
                    return
                case valueStream <- fn():
                }
            }
        }()
        return valueStream
    }
    done := make(chan interface{})
    defer close(done)

    randnum := func() interface{} {
        return rand.Int()
    }
    for num := range take(done, repeatFn(done, randnum), 10) {
        fmt.Println(num)
    }
}
```

后面还有一段关于pipeline代码的消耗测试，这里只说结论了。类型绑定的stage速度是空接口类型stage的两倍，但在整个过程中也仅仅是稍微快了一点。一般来说，pipeline上的限制因素是生成器，或者是计算密集型的一个stage。如果生成器不像repeat 和 repeatFn 生成器那样从内存中创建流，则可能会受I/O限制。为了缓解这种情况（stage消耗很大），引入扇出扇入技术

### 4.6 扇出，扇入（fan-out,fan-in）

pipeline的属性，可使用独立的，可重新排序的stage的组合来操作数据流。可以重复使用pipeline的各个stage。在多个goroutine上重用我们的pipeline的单个stage以试图并行化来自上游stage的pull。

- 扇出模式
    - 不依赖于之前stage计算的值
    - 运行需要很长时间

- 扇入模式
    - 将多个数据流复用或合并成一个流

```go
fanIn := func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
    var wg sync.WaitGroup
    multiplexedStream := make(chan interface{})

    multiplex := func(c <-chan interface{}) {
        defer wg.Done()
        for i:= range c {
            select {
            case <-done:
                return 
            case multiplexedStream <-i:
            }
        }
    }

    // 从所有的channel里取值
    wg.Add(len(channels))
    for _, c := range channels {
        go multiplex(c)
    }

    // 等待所有的读操作结束
    go func() {
        wg.Wait()
        close(multiplexedStream)
    }

    return multiplexedStream
}
```

### 4.7 or-done-channel

处理来自系统各个分散部分的channel

### 4.8 tee-channel 

分割一个来自channel的值，以便将它们发送到代码的两个独立区域中

### 4.9 桥接channel

从一系列的channel中消费产生的值

### 4.10 context 

作用: 
1. 提供一个可以取消调用图中分支的API
2. 提供用于通过呼叫传输请求范围数据的数据包
#### 4.10.1 常用方法

可能有三种情况
1. goroutine 的父 goroutine 想取消它
2. 一个 goroutine 可能想取消子 goroutine
3. goroutine 中的任何阻塞操作都必须是可抢占的，以便可以被取消

**先看一段代码**

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func printGreeting(ctx context.Context) error {
    greeting, err := genGreeting(ctx)
    if err != nil {
        return err
    }
    fmt.Printf("%s world!\n", greeting)
    return nil
}

func genGreeting(ctx context.Context) (string, error) {
    // 使用WithTimeout包装，将在1s后自动取消返回的context，从而取消它传递该context的任何子函数
    ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
    // ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
    defer cancel()
    switch locale, err := locale(ctx); {
    case err != nil:
        return "", err
    case locale == "EN/US":
        return "hello", nil
    }
    return "", fmt.Errorf("unsupported locale")
}

func printFarewell(ctx context.Context) error {
    farewell, err := genFarewell(ctx)
    if err != nil {
        return err
    }
    fmt.Printf("%s world!\n", farewell)
    return nil
}

func genFarewell(ctx context.Context) (string, error) {
    switch locale, err := locale(ctx); {
    case err != nil:
        return "", err
    case locale == "EN/US":
        return "goodbye", nil
    }
    return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
    select {
    case <-ctx.Done():
        // 这一行返回为什么Context被取消的原因。该错误会已知弹出到main，这会导致取消
        return "", ctx.Err()
    case <-time.After(1 * time.Minute):
    // case <-time.After(1 * time.Second):
    }
    return "EN/US", nil
}

func main() {
    var wg sync.WaitGroup
    // 创建context对象
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    wg.Add(1)
    go func() {
        defer wg.Done()
        if err := printGreeting(ctx); err != nil {
            fmt.Printf("cannot print greeting: %v\n", err)
            // 如果从打印问候语返回错误，main将取消这个context
            cancel()
        }
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        if err := printFarewell(ctx); err != nil {
            fmt.Printf("cannot print farewell: %v\n", err)
        }
    }()
    wg.Wait()
}
```

- 输出结果
    > cannot print greeting: context deadline exceeded
        cannot print farewell: context canceled

将我注释掉的代码更换一下

```go
func genGreeting(ctx context.Context) (string, error) {
    // 使用WithTimeout包装，将在1s后自动取消返回的context，从而取消它传递该context的任何子函数
    // ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
    ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
    defer cancel()
    switch locale, err := locale(ctx); {
    case err != nil:
        return "", err
    case locale == "EN/US":
        return "hello", nil
    }
    return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
    select {
    case <-ctx.Done():
        // 这一行返回为什么Context被取消的原因。该错误会已知弹出到main，这会导致取消
        return "", ctx.Err()
    // case <-time.After(1 * time.Minute):
    case <-time.After(1 * time.Second):
    }
    return "EN/US", nil
}
```

- 输出结果
    > goodbye world!
    hello world!

发现就可以正常输出了。原因就在于local的运行时间已超时而导致cancel，所以 `genGreeting` 检测到了ctx的超时也因此无法正常运行

接下来再对函数进行更改，基于第一部分的代码将locale函数更改为

```go
func locale(ctx context.Context) (string, error) {
    if deadline, ok := ctx.Deadline(); ok {
        if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
            return "", context.DeadlineExceeded
        }
    }
    select {
    case <-ctx.Done():
        // 这一行返回为什么Context被取消的原因。该错误会已知弹出到main，这会导致取消
        return "", ctx.Err()
    case <-time.After(1 * time.Minute):
        // case <-time.After(1 * time.Second):
    }
    return "EN/US", nil
}
```

`ctx.Deadline()` 是为了检查context是否提供了超时时间。如果超过截止时间，那么返回context包中的特定错误。**允许函数立即失败**，这是我在使用python的时候一直想要的一个功能。

#### 4.10.2 Cancel

`context` 的另一个功能: **用于存储和检索请求范围数据的Context的数据包**

```go
package main

import (
    "context"
    "fmt"
)

func ProcessRequest(userID, authToken string) {
    ctx := context.WithValue(context.Background(), "userID", userID)
    ctx = context.WithValue(ctx, "authToken", authToken)
    HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
    fmt.Printf("handling response for %v (%v)", ctx.Value("userID"), ctx.Value("authToken"))
}

func main() {
    ProcessRequest("hunt", "123")
}
```

其中需要注意的是需要使用类型安全的函数来从context获取值

```go
type ctxKey int

const (
    ctxUserID ctxKey = iota
    ctxAuthToken
)

func UserID(c context.Context) string {
    return c.Value(ctxUserID).(string)
}

func AuthToken(c context.Context) string {
    return c.Value(ctxAuthToken).(string)
}

func HandleResponse(ctx context.Context) {
    // fmt.Printf("handling response for %v (%v)", ctx.Value("userID"), ctx.Value("authToken"))
    fmt.Printf("handling response for %v (auth: %v)", UserID(ctx), AuthToken(ctx))
}
```

## 5. Go语言的并发模式

### 5.1 异常传递



