
## Concepts
### Multitasking
- `Process`: unit to **allocate** computing resources of a program
- `Thread`: unit to **compute** some parts of a process
- `Coroutine`: time slice of thread to **compute** some parts of a thread

### Context Switch
> store the `state` of a **process** or a **thread** so that it can be `restored` and `resume` execution **later**
- `Process`: switch is controlled by **OS**
- `Thread`: switch is controlled by **OS**: **smallest** unit for context switching for `CPU`
- `Coroutine`: switch is controlled by **User Application**, ex. go-runtime

<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-05-21-01-06-51.png alt="no image found"></p>

## Go Routine
- an implementation of Coroutine
<p align="center"><img style="display: block; width: 600px; margin: 0 auto;" src=img/2020-05-21-01-07-20.png alt="no image found"></p>

- One machine corresponds to exact one KSE
- One Processors corresponds to one Machine at a time, but will switch to other Machine when time goes
- G and P layer both have a queue of free goroutine, which are reused for new task coming in
- **The default number of P is the CPU core number**, upper limit is 256, can be changed manually
```sh
GOMAXPROCS=10 go run main.go
```
- `GOMAXPROCS` represents the **Parallism**, **which is determined essentially by the number of CPU cores**.
- Number of M is bigger than the number of P since some M will be blocked  so new M will be generated
<!-- TODO: need deeper understaning of preemtive scheduling -->
- Go is preemptive scheduling
## Projects References
<!-- TODO: Add sample projects here -->
## References
- [Context Switch](https://osr507doc.xinuos.com/en/PERFORM/context_switching_cpu.html)
- [进程、线程、协程、例程、过程的区别是什么？](https://www.cnblogs.com/f-ck-need-u/p/10802716.html)
- [Achieving concurrency in Go](https://medium.com/rungo/achieving-concurrency-in-go-3f84cbf870ca)
- [Illustrated Tales of Go Runtime Scheduler](https://medium.com/@ankur_anand/illustrated-tales-of-go-runtime-scheduler-74809ef6d19b)
- [也谈goroutine调度器](https://tonybai.com/2017/06/23/an-intro-about-goroutine-scheduler/)