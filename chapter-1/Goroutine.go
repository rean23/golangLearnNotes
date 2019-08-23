package main

/**
Goroutine以及channel配合使用
*/
import (
	"fmt"
	"time"
)

func main() {
	//通道channel

	/**
	创建channel
	ci := make(chan int)
	cs := make(chan string)
	cf := make(chan interface{})
	ch <- v    // 发送 v 到 channel ch.
	v := <-ch  // 从 ch 中接收数据，并赋值给v
	*/

	/**
	channel默认情况下读或者写都是阻塞的除非另一端已经准备好，这样就使得 Goroutines 同步变的更加的简单，
	而不需要显式的 lock。
	所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
	其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲 channel 是在多个 goroutine 之间同步很棒的工具。
	*/

	//这里并没有打印,因为主进程没有等goroutine 里的执行就结束了
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go func(i int) {
			c <- i
		}(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	/**
	这里说一个死锁的问题 在单一继承中操作操作channel(比如main里面操作)但是没有其他goroutine来接收消息，这样所有线程或进程都在等待释放,这样称之为死锁。
	*/

	/**
	缓冲渠道
	创建语法:ch := make(chan type, value)
	当 value = 0 时，channel 是无缓冲阻塞读写的，
	当 value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。
	*/

	/*ch := make(chan int, 3)//修改成1，则造成死锁 因为写入一个值就会阻塞,而没有其他地方来读取数据来解除阻塞
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	*/
	ch1 := make(chan int, 10)

	go func(n int, c chan int) {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}(cap(ch1), ch1)

	/*
		for range 能够不断读取channel里的数据,直到其被显式的关闭。
		关闭之后的channel无法再发送数据.
		如果没有关闭channel,则无法从channel获取数据的时候会报:goroutine 1 [chan receive]:错误
	*/
	for item := range ch1 {
		fmt.Println(item)
	}

	//判断channel是否已经关闭
	v, ok := <-ch1
	fmt.Println(v, ok) //ok返回false表示channel已经关闭

	/**
	记住应该在生产者的地方关闭 channel，而不是消费的地方去关闭它，这样容易引起 panic
	另外记住一点的就是 channel 不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束 range 循环之类的
	*/

	/**
	select
	类似于switch,但是每个case都是一个通信操作,要么发送要么
	*/
	ch2 := make(chan int)
	ch3 := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int) {
			ch2 <- i
			ch3 <- i
		}(i)
	}

	timeOut := make(chan bool)
	for {
		select {
		//如果能从ch2读出数据 则执行此case
		case c := <-ch2:
			fmt.Println("Recvice value", c)
			fmt.Println("from channel")
			goto Foo //这里无法使用break跳出循环,只能跳出当前case
		//如果能从ch3读出数据 则执行此case
		case s := <-ch3:
			fmt.Println("Recvice value", s)
			fmt.Println("from channel1")
			goto Foo
		//超时处理
		case <-time.After(5 * time.Second):
			timeOut <- true
			break
		//如果case都不满足,则执行default
		default:
			fmt.Println("nothing Recvice")
		}
	}
	// 使用goto跳出for-select循环
	Foo:
}
