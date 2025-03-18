## 1. Lịch sử
- Golang lần đầu được ra đời vào năm 2007 để có thể cải thiện hiệu năng các chương trình đa nhân, mạng máy tính.
- Golang được xác định là ngôn ngữ yêu cầu khai báo kiểu dữ liệu trước (Compile Time).
- Dễ hiểu, dễ dùng.
- Hiệu năng cao trong các tác vụ về xử lý đa nhân.


---
## 2. So sánh
- Nhanh, hiệu quả với các tiến trình đa nhân, lập trình mạng.
- Phù hợp để lập trình multithreading vì có Goroutines và Channel hỗ trợ tốt.
- Được kế thừa tính dễ dùng của các dynamic languague như Python và Javascript.


---
## 3. Function
- Variadic function: Hàm có thể truyền vào tùy ý các đối số bằng cách sử dụng syntax: `...`


```go
package main

import "fmt"

func printEven(nums ...int) {
	for _, n := range nums {
		if (n & 1) == 0 {
			fmt.Println(n)
		}
	}
}

func main() {
	printEven(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    // 2 4 6 8 10
}
```
- Multiple return: Hàm trả về nhiều dữ liệu với các kiểu dữ liệu tùy ý


```go
package main

import "fmt"

func cntNum(nums ...int) (int, int) {
	cnt := 0
	for _, n := range nums {
		if (n & 1) == 0 {
			cnt += 1
		}
	}
	return cnt, len(nums) - cnt
}

func main() {
	even, odd := cntNum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(even, odd)
	// 5 5
}

```
- Error / Exception handling

Golang không có cơ chế exception như các ngôn ngữ khác và sử dụng giá trị trả về để xử lý lỗi. Hàm thường trả về giá trị thứ hai là lỗi, kiểu `error`

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	res, err := divide(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
    // Result: 2

	res, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
    // Error: division by zero
}
```
---
## 4. For-loop
Vòng lặp qua các phần tử của mảng, slice, map hoặc string

```go
package main

import "fmt"

func main() {
	// Slice
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	// Map
	dict := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	for key, val := range dict {
		fmt.Println(key, val)
	}

	// String
	str := "helloworld"
	for idx, char := range str {
		fmt.Println(idx, string(char))
	}
}
```

---
## 5. Package, Import

Có 2 loại package:

- Executable package là package chứa hàm main, để thực thi hàm main, là package được khai báo là main.

- Reusable package là package được dùng để viết các library và được tái sử dụng nhờ câu lệnh `import`, là các package không được khai báo là main.

Import statement: Câu lệnh `import` được dùng để import các package cần được sử dụng trong chương trình như một số standard package như: `fmt`, `io`, `time`, ... và các thư viện người dùng tự định nghĩa nhờ sử dụng reusable package.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Square root of 16 is: ", math.Sqrt(16))
	// Square root of 16 is:  4
}
```
---
## 6. Go module
Go module là hệ thống quản lý các thư viện và dependencies trong các dự án Go, cho phép việc quản lý phiên bản và các gói thư viện dễ dàng hơn.

- Tạo một module mới

```sh
go mod init example.com/mymodule
```

- Thêm một dependency

Để thêm một dependency vào module, chỉ cần import nó vào mã nguồn và chạy lệnh `go build` hoặc `go mod tidy`. Điều này sẽ tự động thêm dependency vào file `go.mod`.
```go
package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("Generated UUID:", id)
}
```
Sau đó, chạy lệnh:

```sh
go mod tidy
```
- Cập nhật và quản lý dependencies

```sh
go get github.com/google/uuid@latest
```
---
## 7. Naming Convetion
> Tên package

- Tên package nên ngắn gọn, thường là một chữ đơn, viết bằng chữ thường.
- Tên package nên mô tả chính xác chức năng hoặc mục đích của nó.

```go
package math
package http
```

> Tên biến và hàm

- Tên biến và hàm nên sử dụng camelCase.
- Tên nên ngắn gọn nhưng có ý nghĩa, đủ để mô tả chức năng hoặc dữ liệu mà chúng nắm giữ.

```go
var userName string
func calculateSum(a int, b int) int {
    return a + b
}
```

> Tên Struct và Interface

- Tên struct và interface nên để dạng PascalCase.
- Tên struct thường là danh từ, trong khi interface thường là tính từ hoặc danh từ kết thúc bằng "er".

```go
type User struct {
    Name  string
    Email string
}

type Reader interface {
    Read(p []byte) (n int, err error)
}
```

> Tên Hằng số

- Tên hằng số nên sử dụng kiểu viết hoa với dấu gạch dưới giữa các từ (SNAKE_CASE)

```go
const MAX_BUFFER_SIZE = 1024
const PI = 3.14
```

> Exported vs Unexported Names

- Tên bắt đầu bằng chữ hoa sẽ được export và các package khác bên ngoài có thể truy cập được.
- Tên bắt đầu bằng chữ thường không được export và chỉ có thể truy cập trong package hiện tại.

```go
// Exported
func PrintMessage(msg string) {
    fmt.Println(msg)
}

// Unexported
func calculateSum(a int, b int) int {
    return a + b
}
```
---

## 8. Data structure

> Phân biệt giữa new và make

- **new**: Từ khóa `new` không được dùng để khởi tạo kiểu dữ liệu nào đó mà được dùng để khởi tạo một pointer trỏ đến vùng nhớ lưu zero-value của kiểu dữ liệu đó.

```go
var p *int
p = new(int)
// p sẽ trỏ đến một vùng nhớ chứa zero-value của int (0)
```

- **make**: từ khóa `make` chỉ được dùng để khởi tạo các data structures như map, slice, channel khi chúng cần được khởi tạo để có thể sử dụng được. Hàm `make` trả về kiểu dữ liệu giống với kiểu dữ liệu được truyền vào và được khởi tạo size và capacity tương ứng với các tham số truyền vào.

```go
s := make([]int, 5, 6)
m := make(map[string]int)
c := make(chan int)
```

- **Slice**: là references đến array, mỗi khi tạo ra một slice, sẽ tạo ra 2 phần chính là slice chưa pointer, size và capacity của array và actual array. Slice hỗ trợ thêm phần tử, cấp phát động thêm size, và lấy subarray nhờ toán tử: `:`.

```go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:3]
fmt.Println(s)
// [2 3]
```

- Trong trường hợp slice ban đầu không đủ sức chứa khi thêm phần tử, hàm append sẽ thực hiện cấp phát lại vùng nhớ có kích thước:
	+ Nếu kích thước cũ (cap) < 1024: cấp phát gấp đôi (x2) vùng nhớ cũ.
	+ Nếu kích thước cũ >= 1024: cấp pháp 1.25x vùng nhớ cũ.

```go
package main
import "fmt"
func main() {
	s := make([]int, 0, 1)

	for i := 1; i <= 1300; i++ {
		s = append(s, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
	}
}
```

- **Map**: Giống với slice, mao cũng là kiểu dữ liệu references đến data structures. Được định nghĩa bởi cặp (key, val).

```go
m := make(map[string]int)
m["one"] = 1
fmt.Println(m["one"]) // 1
```

- Nếu truy cập tới một phần tử không có trong map thì nó sẽ trả về giá trị zero. Để chắc chắn thì mỗi khi truy cập một phần tử có trong map, kiểm tra giá trị trả về thứ hai.

```go
val, ok := m["two"]
if !ok {
	fmt.Println("Key not found")
} else {
	fmt.Println(val)
}
```

- **Array**: Là một tập hợp các phần tử có cùng kiểu dữ liệu, được lưu trữ liên tiếp trong bộ nhớ. Kích thước của array là cố định và không thể thay đổi.

```go
var a [5]int
a[0] = 1
a[1] = 2
fmt.Println(a) // [1 2 0 0 0]

b := [3]int{1, 2, 3}
fmt.Println(b) // [1 2 3]
```

- **Struct**: Là một tập hợp cá trường (fields), có thể chứa các kiểu dữ liệu khác nhau. Struct được sử dụng để nhóm các trường dữ liệu liên quan lại với nhau.

```go
type Person struct {
	Name string
	Age int
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p) // {Alice 25}
}
```

> Channel

- **Channel**: Là cấu trúc dữ liệu đặc biệt được dùng để giao tiếp giữa các goroutine. Channel cho phép một goroutine gửi dữ liệu và một goroutine khác nhận dữ liệu.

```go
func main() {
	ch := make(chan int)
	go func() {
		c <- 42
	}()
	val <- c
	fmt.Println(val)
}
```

> Linked List

- **Linked List**: Là một cấu trúc dữ liệu bao gồm các node, mỗi node chứa một giá trị và một con trỏ trỏ đến node kế tiếp.

```go
type Node struct {
	Value int
	Next *Node
}

func main() {
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}

	for n := head; n != nil; n = n.Next {
		fmt.Println(n.Value)
	}
}
```

> Stack

- **Stack**: Là một cấu trúc dữ liệu LIFO. Các phần tử được thêm vào và lấy ra từ đỉnh của stack.

```go
type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1
	}
	value := s.elements[len(s.elements) - 1]
	s.elements = s.elements[ : len(s.elements) - 1]
	return value
}
```

> Queue

- **Queue**: Là một cấu trúc dữ liệu FIFO. Các phần tử được thêm vào cuối queue và lấy ra từ đầu queue.

```go
type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		return -1
	}
	value := q.elements[0]
	q.elements = q.elements[1 : ]
	return value
}
```

---
## 9.Multithreading

Multithreading trong Golang được hỗ trợ bởi Goroutines và Channels. Goroutines là các hàm hoặc phương thức chạy đồng thời với các hàm khác, trong khi Channels cung cấp cơ chế giao tiếp và đồng bộ hóa giữa các Goroutine.


- Goroutines là các luồng nhẹ được quản lý bởi runtime của Go. Để tạo một goroutines, sử dụng từ khóa `go` trước một hàm hoặc mọt phương thức.

```go
package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers()
	time.Sleep(600 * time.Millisecond)
}
```

- Channels

Channels là cơ chế dùng để giao tiếp giữa các Goroutine. Các giá trị có thể gửi và nhận thông qua Channels.

```go
package main

import "fmt"

func sum(a, b int, c chan int) {
	c <- a + b
}

func main() {
	c := make(chan int)
	go sum(1, 2, c)
	res := <-c
	fmt.Println(res)
}
```

- Buffered Channels

Buffered Channels cho phép định nghĩa số lượng giá trị có thể lưu trữ trong Channel mà không cần đợi giá trị được nhận.

```go
package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
```

- Select

Câu lệnh `select` cho phép một Goroutine chờ nhiều hoạt động trên nhiều Channels.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
			case msg1 := <-c1:
				fmt.Println("Received", msg1)
			case msg2 := <-c2:
				fmt.Println("Received", msg2)
		}
	}
}
```

- WaitGroup

WaitGroup là công cụ đồng bộ hóa cho phép đợi cho đến khi một tập hợp các Goroutine hoàn thành công việc của chúng.

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go.worker(1, &wg)
	}
	wg.Wait()
}
```

- Mutex

Mutex là cơ chế giúp đảm bảo tại một thời điểm, chỉ có một Goroutine truy cập được vào các phần tử dữ liệu.

```go
package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
}
```

- Worker Pool

Worker Pool là một mô hình trong đó một số lượng cố định các Goroutine (workders) xử lý các công việc từ một hàng đợi.
```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		results <- j * 2
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Result:", result)
	}
}
```

- Fan-in Fan-out

Fan-in Fan-out là một mô hình trong đó nhiều Goroutine (producers) gửi dữ liệu vào một Channel (fan-in) và nhiều Goroutine khác (consumers) xử lý dữ liệu từ Channel đó (fan-out).

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// fan-in
func producer(id int, data chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(id) * time.Second)
		data <- i
		fmt.Printf("Producer %d produced %d\n", id, i)
	}
}

// fan-out
func consumer(id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range data {
		fmt.Printf("Consumer %d consumed %d\n", id, d)
	}
}

func main() {
	data := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		go producer(i, data)
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go consumer(i, data, &wg)
	}
	time.Sleep(10 * time.Second)
	close(data)
	wg.Wait()
}

```

---

## 10. Common Package

- fmt

Package `fmt` cung cấp các hàm định dạng và in dữ liệu

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	name := "Alice"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

- time

Package `time` cung cấp các hàm và kiểu dữ liệu để làm việc với thời gian và ngày tháng.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time: " + now)

	later := now.Add(time.Hour * 2)
	fmt.Println("Two Hours Later", later)
}
```

- net/http 

Package `net/http` cung cấp các hàm để xây dựng các ứng dụng web và HTTP client

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting at 8080")
	http.ListenAndServe(":8080", nil)
}
```
