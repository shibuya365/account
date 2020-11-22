package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Calc(key string, value int) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key] += value
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	for j := range jobs {
		// 1行を表示
		defer wg.Done()
		// fmt.Println("worker", id, "job", j)

		// 1行を単語ごとに分割
		slice := strings.Fields(j)
		value, _ := strconv.Atoi(slice[2])
		fmt.Printf("worker: %v date:%s %s %v %s \n", id, slice[0], slice[1], value, slice[3])
		c.Calc(slice[1], value)
		c.Calc(slice[3], value*-1)
	}
}

var c = SafeCounter{v: make(map[string]int)}

func main() {
	balance := [3]string{"現金", "預金", "資本金"}
	const numJobs = 5
	jobs := make(chan string, numJobs)
	// results := make(chan string, numJobs)
	var wg sync.WaitGroup

	// ファイルの読み込み
	f, err := os.Open("./data.txt")
	if err != nil {
		return
	}

	// CPUコア数分のgoroutine起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, jobs, &wg)
	}

	// 仕事登録
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wg.Add(1)
		jobs <- scanner.Text()
	}

	// ジョブ終了
	close(jobs)

	// 処理待ち
	wg.Wait()

	// 貸借科目
	fmt.Println("# 貸借科目")
	for _, key := range balance {
		fmt.Printf("%s %v \n", key, c.Value(key))
		delete(c.v, key)
	}

	// 損益科目
	fmt.Println("# 損益科目")
	for key, _ := range c.v {
		fmt.Printf("%s %v \n", key, c.Value(key))
	}
}
