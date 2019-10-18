package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// 将依赖关系定义为一个接口，我们可以在 main 使用真实的 Sleeper，同时在
// 测试中使用 spy sleeper
type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurablesSleeper struct {
	duration time.Duration
}

func (c *ConfigurablesSleeper) Sleep() {
	time.Sleep(c.duration)
}

func main() {
	sleeper := &ConfigurablesSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
