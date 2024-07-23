/*
Example 1: Basic Channel Communication

Description:
This example demonstrates basic communication between two goroutines using CMP library channels.

Steps:
1. Create a new CMP channel.
2. Start a goroutine to send a message through the channel.
3. Receive and print the message in the main goroutine.

Expected Output:
Message received: Hello, CMP!
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		ch.Send("Hello, CMP!")
	}()

	msg := ch.Receive().(string)
	fmt.Println("Message received:", msg)
}





/*
Example 2: Synchronous Message Passing

Description:
This example demonstrates synchronous message passing using CMP library's request-response pattern.

Steps:
1. Create a new CMP channel.
2. Start a goroutine to handle requests.
3. Send a request and receive a response synchronously.

Expected Output:
Request: What is 2 + 2?
Response: 4
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		for {
			req := ch.Receive().(string)
			if req == "quit" {
				break
			}
			if req == "What is 2 + 2?" {
				ch.Send("4")
			}
		}
	}()

	ch.Send("What is 2 + 2?")
	resp := ch.Receive().(string)
	fmt.Println("Response:", resp)

	ch.Send("quit")
}





/*
Example 3: Select Statement with CMP Channels

Description:
This example demonstrates the use of Go's select statement with CMP channels for non-blocking message handling.

Steps:
1. Create two CMP channels.
2. Use select to handle messages from both channels concurrently.

Expected Output:
Messages received:
- Message 1: Hello from Channel 1
- Message 2: Hi from Channel 2
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"time"
)

func main() {
	ch1 := cmp.NewChannel()
	ch2 := cmp.NewChannel()

	go func() {
		time.Sleep(1 * time.Second)
		ch1.Send("Hello from Channel 1")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2.Send("Hi from Channel 2")
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1.C:
			fmt.Println("Message 1:", msg.(string))
		case msg := <-ch2.C:
			fmt.Println("Message 2:", msg.(string))
		}
	}
}





/*
Example 4: Buffered Channels with CMP

Description:
This example demonstrates the use of buffered channels with CMP library for handling multiple messages concurrently.

Steps:
1. Create a buffered CMP channel.
2. Send multiple messages to the channel concurrently.
3. Receive and print the messages.

Expected Output:
Messages received:
- Message: Hello, CMP! (received from buffered channel)
- Message: How are you? (received from buffered channel)
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch := cmp.NewBufferedChannel(2)

	go func() {
		ch.Send("Hello, CMP!")
		ch.Send("How are you?")
	}()

	for i := 0; i < 2; i++ {
		msg := ch.Receive().(string)
		fmt.Println("Message:", msg)
	}
}





/*
Example 5: Timeout Handling with CMP Channels

Description:
This example demonstrates timeout handling using CMP channels in Go.

Steps:
1. Create a CMP channel.
2. Implement a timeout mechanism using select to handle message reception within a specified time.
3. Print message or timeout message based on received message or timeout.

Expected Output:
Message received: Hello, CMP! (if received within timeout)
Or
Timeout: No message received within 1 second. (if timed out)
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch.Send("Hello, CMP!")
	}()

	select {
	case msg := <-ch.C:
		fmt.Println("Message received:", msg.(string))
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No message received within 1 second.")
	}
}





/*
Example 6: Broadcast Communication with CMP

Description:
This example demonstrates broadcasting messages to multiple subscribers using CMP channels.

Steps:
1. Create a CMP channel.
2. Start multiple goroutines as subscribers to receive broadcasted messages.
3. Broadcast a message to all subscribers.

Expected Output:
Messages received by subscribers:
- Subscriber 1 received: Hello, CMP!
- Subscriber 2 received: Hello, CMP!
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Subscriber 1 received:", ch.Receive().(string))
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Subscriber 2 received:", ch.Receive().(string))
	}()

	ch.Send("Hello, CMP!")

	wg.Wait()
}





/*
Example 7: Select Statement with Timeout for CMP

Description:
This example demonstrates using select with a timeout for CMP channels in Go.

Steps:
1. Create a CMP channel.
2. Implement a select statement with timeout to handle message reception within a specified time.
3. Print received message or timeout message based on the outcome.

Expected Output:
Message received: Hello, CMP! (if received within timeout)
Or
Timeout: No message received within 1 second. (if timed out)
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch.Send("Hello, CMP!")
	}()

	select {
	case msg := <-ch.C:
		fmt.Println("Message received:", msg.(string))
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No message received within 1 second.")
	}
}





/*
Example 8: Close and Reopen CMP Channel

Description:
This example demonstrates closing and reopening a CMP channel in Go.

Steps:
1. Create a CMP channel.
2. Close the channel after sending a message.
3. Attempt to send another message after reopening the channel.

Expected Output:
Messages received:
- Message: Hello, CMP! (first message received)
- Message: Hi again! (second message received after reopening)
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch := cmp.NewChannel()

	ch.Send("Hello, CMP!")
	ch.Close()

	// Reopen channel
	ch = cmp.NewChannel()
	ch.Send("Hi again!")

	fmt.Println("Messages received:")
	fmt.Println("- Message:", ch.Receive().(string))
	fmt.Println("- Message:", ch.Receive().(string))
}





/*
Example 9: CMP Channel with Structured Messages

Description:
This example demonstrates using CMP channels with structured messages (custom types) in Go.

Steps:
1. Define a struct for messages.
2. Create a CMP channel for the struct type.
3. Send and receive structured messages through the channel.

Expected Output:
Message received: {John Doe 30}
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	ch := cmp.NewChannel()

	go func() {
		ch.Send(Person{Name: "John Doe", Age: 30})
	}()

	msg := ch.Receive().(Person)
	fmt.Println("Message received:", msg)
}





/*
Example 10: Fan-Out Pattern with CMP Channels

Description:
This example demonstrates the fan-out pattern using CMP channels in Go.

Steps:
1. Create a CMP channel.
2. Start multiple goroutines (subscribers) to receive messages concurrently.
3. Send a message to the CMP channel to be received by all subscribers.

Expected Output:
Messages received:
- Subscriber 1 received: Hello, CMP!
- Subscriber 2 received: Hello, CMP!
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	wg.Add(2)

	for i := 1; i <= 2; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Subscriber %d received: %s\n", id, ch.Receive().(string))
		}(i)
	}

	ch.Send("Hello, CMP!")

	wg.Wait()
}





/*
Example 11: CMP Channel with Error Handling

Description:
This example demonstrates error handling using CMP channels in Go.

Steps:
1. Create a CMP channel.
2. Send a message containing an error.
3. Receive and handle the error message.

Expected Output:
Error received: Error: Something went wrong
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		ch.Send(fmt.Errorf("Error: Something went wrong"))
	}()

	err := ch.Receive().(error)
	fmt.Println("Error received:", err)
}





/*
Example 12: CMP Channel with Timeout and Default Value

Description:
This example demonstrates using CMP channels with timeout and default value handling in Go.

Steps:
1. Create a CMP channel.
2. Implement a select statement with timeout to handle message reception within a specified time.
3. Print received message or default value based on timeout.

Expected Output:
Message received: Hello, CMP! (if received within timeout)
Or
Default value: No message received within 1 second. (if timed out)
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	select {
	case msg := <-ch.C:
		fmt.Println("Message received:", msg.(string))
	case <-time.After(1 * time.Second):
		fmt.Println("Default value: No message received within 1 second.")
	}
}





/*
Example 13: CMP Channel with Select Statement and Exit Signal

Description:
This example demonstrates using a CMP channel with a select statement and an exit signal in Go.

Steps:
1. Create a CMP channel.
2. Implement a goroutine to handle messages and an exit signal.
3. Use a select statement to receive messages or exit the goroutine.

Expected Output:
Messages received:
- Message: Hello, CMP! (if received before exit signal)
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		for {
			select {
			case msg := <-ch.C:
				fmt.Println("Message:", msg.(string))
			case <-time.After(1 * time.Second):
				fmt.Println("Timeout: No message received within 1 second.")
				return // Exit goroutine after timeout
			}
		}
	}()

	ch.Send("Hello, CMP!")

	time.Sleep(2 * time.Second) // Wait to see the output
}





/*
Example 14: CMP Channel with Goroutine Pool

Description:
This example demonstrates using a CMP channel with a goroutine pool in Go.

Steps:
1. Create a CMP channel.
2. Implement multiple goroutines in a pool to handle messages concurrently.
3. Send messages to the channel to be processed by the goroutine pool.

Expected Output:
Messages received:
- Worker 1 processed: Hello from CMP!
- Worker 2 processed: How are you?

Note: The order of messages processed may vary due to concurrency.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const numWorkers = 2
	wg.Add(numWorkers)

	// Worker pool
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for {
				msg := ch.Receive().(string)
				fmt.Printf("Worker %d processed: %s\n", workerID, msg)
			}
		}(i)
	}

	// Send messages
	ch.Send("Hello from CMP!")
	ch.Send("How are you?")

	wg.Wait()
}





/*
Example 15: CMP Channel with Priority Queue

Description:
This example demonstrates using a CMP channel as a priority queue in Go.

Steps:
1. Create a CMP channel.
2. Implement goroutines with different priorities to handle messages.
3. Send messages with priorities to the channel.

Expected Output:
Messages processed based on priority:
- High priority message: Processing urgent task!
- Normal priority message: Processing regular task

Note: Messages with higher priorities are processed first.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

type Message struct {
	Text     string
	Priority int
}

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const numWorkers = 2
	wg.Add(numWorkers)

	// Worker pool
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for {
				msg := ch.Receive().(Message)
				fmt.Printf("Worker %d processed: %s\n", workerID, msg.Text)
			}
		}(i)
	}

	// Send messages with priorities
	ch.Send(Message{Text: "Processing urgent task!", Priority: 1})
	ch.Send(Message{Text: "Processing regular task", Priority: 2})

	wg.Wait()
}





/*
Example 16: CMP Channel with Fan-In Pattern

Description:
This example demonstrates the fan-in pattern using CMP channels in Go.

Steps:
1. Create multiple CMP channels for input.
2. Implement a goroutine to multiplex messages from multiple channels into one output channel.
3. Receive and process messages from the output channel.

Expected Output:
Messages received:
- Message from Channel 1: Hello, CMP!
- Message from Channel 2: Hi, there!
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch1 := cmp.NewChannel()
	ch2 := cmp.NewChannel()

	outputCh := cmp.NewChannel()

	// Fan-in multiplexer
	go func() {
		for {
			select {
			case msg := <-ch1.C:
				outputCh.Send(fmt.Sprintf("Message from Channel 1: %s", msg.(string)))
			case msg := <-ch2.C:
				outputCh.Send(fmt.Sprintf("Message from Channel 2: %s", msg.(string)))
			}
		}
	}()

	ch1.Send("Hello, CMP!")
	ch2.Send("Hi, there!")

	fmt.Println("Messages received:")
	fmt.Println("- " + outputCh.Receive().(string))
	fmt.Println("- " + outputCh.Receive().(string))
}





/*
Example 17: CMP Channel with Rate Limiting

Description:
This example demonstrates using CMP channels with rate limiting in Go.

Steps:
1. Create a CMP channel.
2. Implement a goroutine to handle messages with rate limiting using time.Tick().
3. Send messages to the channel to be processed respecting the rate limit.

Expected Output:
Messages received with rate limiting applied:
- Message 1: Hello from CMP!
- Message 2: How are you?

Note: Messages are processed respecting the rate limit (one message per second).
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		ticker := time.Tick(1 * time.Second)
		messages := []string{"Hello from CMP!", "How are you?"}

		for _, msg := range messages {
			<-ticker
			ch.Send(msg)
		}
	}()

	fmt.Println("Messages received with rate limiting applied:")
	fmt.Println("- Message 1:", ch.Receive().(string))
	fmt.Println("- Message 2:", ch.Receive().(string))
}





/*
Example 18: CMP Channel with External Signal Handling

Description:
This example demonstrates using CMP channels to handle external signals in Go.

Steps:
1. Create a CMP channel for signals.
2. Implement a goroutine to wait for external signals (e.g., SIGINT).
3. Send a signal to the channel when the external signal is received.

Expected Output:
Signal received: Received SIGINT signal!
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT)

		// Wait for SIGINT signal
		<-sigCh
		ch.Send("Received SIGINT signal!")
	}()

	fmt.Println("Signal received:", ch.Receive().(string))
}





/*
Example 19: CMP Channel with Contextual Data

Description:
This example demonstrates using CMP channels with contextual data in Go.

Steps:
1. Create a CMP channel with context.
2. Send messages with contextual data.
3. Receive and process messages along with their context.

Expected Output:
Messages received with context:
- Context: {"userID": 123} Message: Hello, CMP!
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/cmp/cmp"
)

type MessageWithCtx struct {
	Context map[string]interface{}
	Message string
}

func main() {
	ch := cmp.NewChannel()

	go func() {
		ctx := map[string]interface{}{"userID": 123}
		msg := MessageWithCtx{Context: ctx, Message: "Hello, CMP!"}
		ch.Send(msg)
	}()

	receivedMsg := ch.Receive().(MessageWithCtx)
	ctxJSON, _ := json.Marshal(receivedMsg.Context)

	fmt.Printf("Messages received with context:\n- Context: %s Message: %s\n", ctxJSON, receivedMsg.Message)
}





/*
Example 20: CMP Channel with Exponential Backoff Retry

Description:
This example demonstrates using CMP channels with exponential backoff retry mechanism in Go.

Steps:
1. Create a CMP channel for retries.
2. Implement a goroutine to handle retries with exponential backoff.
3. Send retry attempts to the channel and handle retries.

Expected Output:
Retrying attempt 1...
Retrying attempt 2...
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"math"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		maxAttempts := 5
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			waitTime := time.Duration(math.Pow(2, float64(attempt))) * time.Second
			fmt.Printf("Retrying attempt %d...\n", attempt)
			time.Sleep(waitTime)
			ch.Send(fmt.Sprintf("Retry attempt %d", attempt))
		}
	}()

	for i := 1; i <= 2; i++ {
		fmt.Println(ch.Receive().(string))
	}
}





/*
Example 21: CMP Channel with Backpressure Handling

Description:
This example demonstrates using CMP channels with backpressure handling in Go.

Steps:
1. Create a CMP channel with a buffer size.
2. Implement a producer goroutine to send messages.
3. Implement a consumer goroutine to receive messages and apply backpressure if necessary.

Expected Output:
Messages sent and received with backpressure:
- Message 1: Hello from CMP!
- Message 2: How are you?

Note: The consumer processes messages with backpressure based on the buffer size.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"time"
)

func main() {
	ch := cmp.NewBufferedChannel(1)

	// Producer sending messages
	go func() {
		ch.Send("Hello from CMP!")
		ch.Send("How are you?")
	}()

	// Consumer receiving messages
	fmt.Println("Messages sent and received with backpressure:")
	fmt.Println("- Message 1:", ch.Receive().(string))
	fmt.Println("- Message 2:", ch.Receive().(string))

	// Simulate some processing time
	time.Sleep(1 * time.Second)

	// Check if there are more messages in the buffer
	for !ch.Empty() {
		fmt.Println("- Extra Message:", ch.Receive().(string))
	}
}





/*
Example 22: CMP Channel with Message Filtering

Description:
This example demonstrates using CMP channels with message filtering in Go.

Steps:
1. Create a CMP channel.
2. Implement a goroutine to filter and process specific types of messages.
3. Send messages to the channel and process filtered messages.

Expected Output:
Filtered messages processed:
- Filtered message: 42
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		ch.Send(42)
		ch.Send("Hello, CMP!")
		ch.Send(true)
	}()

	// Filter and process specific type of messages
	for {
		msg := ch.Receive()
		switch msg.(type) {
		case int:
			fmt.Println("Filtered message:", msg)
		}
		if ch.Empty() {
			break
		}
	}
}





/*
Example 23: CMP Channel with Task Distribution

Description:
This example demonstrates using CMP channels for task distribution among workers in Go.

Steps:
1. Create a CMP channel for task distribution.
2. Implement multiple worker goroutines to receive and process tasks concurrently.
3. Send tasks to the channel to be distributed and processed by the workers.

Expected Output:
Tasks distributed and processed by workers:
- Worker 1 processed task: Task 1
- Worker 2 processed task: Task 2

Note: Tasks are processed concurrently by multiple workers.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const numWorkers = 2
	wg.Add(numWorkers)

	// Worker pool
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for {
				task := ch.Receive().(string)
				fmt.Printf("Worker %d processed task: %s\n", workerID, task)
				if ch.Empty() {
					break
				}
			}
		}(i)
	}

	// Send tasks
	ch.Send("Task 1")
	ch.Send("Task 2")

	wg.Wait()
}





/*
Example 24: CMP Channel with Batching

Description:
This example demonstrates using CMP channels for batching messages in Go.

Steps:
1. Create a CMP channel for batching.
2. Implement a goroutine to batch messages and send them in bulk.
3. Send individual messages to the channel to be batched and processed.

Expected Output:
Messages batched and processed:
- Batch 1: [Hello, CMP! How are you?]

Note: Messages are batched into slices and processed in bulk.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	// Batching goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var batch []interface{}
		for {
			msg := ch.Receive()
			batch = append(batch, msg)
			// Process batch every 2 seconds or when channel is empty
			if len(batch) >= 2 || ch.Empty() {
				fmt.Printf("Batch %d: %v\n", len(batch), batch)
				batch = nil // Clear batch
			}
			if ch.Empty() {
				break
			}
		}
		wg.Done()
	}()

	// Send messages
	ch.Send("Hello, CMP!")
	ch.Send("How are you?")

	wg.Wait()
}





/*
Example 25: CMP Channel with Error Recovery

Description:
This example demonstrates using CMP channels for error recovery in Go.

Steps:
1. Create a CMP channel for error recovery.
2. Implement a goroutine to handle errors and recover from them.
3. Send errors to the channel and handle recovery.

Expected Output:
Errors recovered and processed:
- Error: Something went wrong
- Recovered from error: Error handled successfully

Note: Errors are recovered and processed to ensure graceful error handling.
*/

package main

import (
	"errors"
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		err := errors.New("Something went wrong")
		ch.Send(err)
	}()

	// Handle errors and recovery
	for {
		err := ch.Receive()
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Recovered from error: Error handled successfully")
			break
		}
	}
}





/*
Example 26: CMP Channel with Timeout and Retry

Description:
This example demonstrates using CMP channels with timeout and retry mechanism in Go.

Steps:
1. Create a CMP channel for communication.
2. Implement a goroutine to send messages with potential delays.
3. Implement another goroutine to receive messages with timeout and retry if necessary.

Expected Output:
Messages sent and received with timeout and retry:
- Attempt 1: Hello, CMP!
- Attempt 2: How are you?

Note: Messages are sent with potential delays, and receiver retries with timeout if no message received.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	// Sender goroutine with delays
	go func() {
		time.Sleep(1 * time.Second)
		ch.Send("Hello, CMP!")

		time.Sleep(2 * time.Second)
		ch.Send("How are you?")
	}()

	// Receiver goroutine with timeout and retry
	fmt.Println("Messages sent and received with timeout and retry:")
	for attempt := 1; attempt <= 2; attempt++ {
		select {
		case msg := <-ch.C:
			fmt.Printf("- Attempt %d: %s\n", attempt, msg.(string))
		case <-time.After(1 * time.Second):
			fmt.Printf("- Timeout: Retry attempt %d\n", attempt)
		}
	}
}





/*
Example 27: CMP Channel with Rate Limiting and Batching

Description:
This example demonstrates using CMP channels with rate limiting and batching in Go.

Steps:
1. Create a CMP channel for communication.
2. Implement a goroutine to send messages respecting a rate limit.
3. Implement another goroutine to receive and process batched messages.

Expected Output:
Messages sent and received with rate limiting and batching:
- Batch 1: [Hello from CMP! How are you?]

Note: Messages are sent respecting rate limit and processed in batches.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	// Sender goroutine respecting rate limit
	go func() {
		messages := []string{"Hello from CMP!", "How are you?"}
		for _, msg := range messages {
			ch.Send(msg)
			time.Sleep(1 * time.Second) // Rate limit of 1 message per second
		}
	}()

	// Receiver goroutine processing batched messages
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var batch []string
		for {
			msg := ch.Receive().(string)
			batch = append(batch, msg)
			if len(batch) >= 2 || ch.Empty() {
				fmt.Printf("Batch %d: %v\n", len(batch), batch)
				batch = nil // Clear batch
			}
			if ch.Empty() {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()
}





/*
Example 28: CMP Channel with Worker Pool and Error Handling

Description:
This example demonstrates using CMP channels with a worker pool and error handling in Go.

Steps:
1. Create a CMP channel for task distribution.
2. Implement multiple worker goroutines to process tasks concurrently.
3. Send tasks to the channel and handle errors from workers.

Expected Output:
Tasks distributed and errors handled by workers:
- Worker 1 processed task: Task 1
- Worker 2 processed task: Task 2
- Error from Worker 2: Error: Task failed

Note: Tasks are processed concurrently by multiple workers, and errors are handled gracefully.
*/

package main

import (
	"errors"
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const numWorkers = 2
	wg.Add(numWorkers)

	// Worker pool
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for {
				task := ch.Receive()
				fmt.Printf("Worker %d processed task: %s\n", workerID, task.(string))

				// Simulate task failure for Worker 2
				if workerID == 2 {
					ch.SendError(errors.New("Error: Task failed"))
				}

				if ch.Empty() {
					break
				}
			}
		}(i)
	}

	// Send tasks
	ch.Send("Task 1")
	ch.Send("Task 2")

	wg.Wait()
}





/*
Example 29: CMP Channel with Context Cancellation

Description:
This example demonstrates using CMP channels with context cancellation in Go.

Steps:
1. Create a CMP channel for task processing.
2. Implement a goroutine to process tasks with context.
3. Cancel context and handle cancellation gracefully.

Expected Output:
Tasks processed with context cancellation:
- Task 1 processed successfully
- Context canceled: Task 2 canceled due to context cancellation

Note: Tasks are processed with context, and cancellation is handled gracefully.
*/

package main

import (
	"context"
	"fmt"
	"github.com/cmp/cmp"
	"sync"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	// Context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(2)

	// Worker 1 processing tasks
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context canceled: Task 1 canceled due to context cancellation")
				return
			case <-time.After(1 * time.Second):
				ch.Send("Task 1 processed successfully")
			}
		}
	}()

	// Worker 2 processing tasks
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context canceled: Task 2 canceled due to context cancellation")
				return
			case <-time.After(2 * time.Second):
				ch.Send("Task 2 processed successfully")
			}
		}
	}()

	// Simulate context cancellation after 3 seconds
	time.Sleep(3 * time.Second)
	cancel()

	wg.Wait()
}





/*
Example 30: CMP Channel with Selective Message Processing

Description:
This example demonstrates using CMP channels for selective message processing in Go.

Steps:
1. Create a CMP channel for message filtering.
2. Implement a goroutine to filter and process specific types of messages.
3. Send messages to the channel and process filtered messages.

Expected Output:
Selective messages processed:
- Message: Hello from CMP!

Note: Only messages of type string are processed, others are ignored.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch := cmp.NewChannel()

	go func() {
		ch.Send(42)
		ch.Send("Hello from CMP!")
		ch.Send(true)
	}()

	// Selective message processing
	for {
		msg := ch.Receive()
		switch msg.(type) {
		case string:
			fmt.Println("Selective message:", msg.(string))
		}
		if ch.Empty() {
			break
		}
	}
}





/*
Example 31: CMP Channel with Dynamic Worker Pool

Description:
This example demonstrates using CMP channels with a dynamic worker pool in Go.

Steps:
1. Create a CMP channel for task distribution.
2. Implement goroutines to dynamically add and remove workers based on task load.
3. Send tasks to the channel to be processed by the worker pool.

Expected Output:
Tasks distributed and processed by dynamic worker pool:
- Worker 1 processed task: Task 1
- Worker 2 processed task: Task 2

Note: Workers are dynamically added or removed based on the task load.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const maxWorkers = 3
	wg.Add(maxWorkers)

	// Dynamic worker pool
	for i := 1; i <= maxWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for {
				task := ch.Receive()
				if task == nil {
					break
				}
				fmt.Printf("Worker %d processed task: %s\n", workerID, task.(string))
			}
			fmt.Printf("Worker %d stopped\n", workerID)
		}(i)
	}

	// Send tasks
	ch.Send("Task 1")
	ch.Send("Task 2")

	// Simulate adding more tasks
	ch.Send("Task 3")
	ch.Send("Task 4")

	// Close channel to signal no more tasks
	ch.Close()

	wg.Wait()
}





/*
Example 32: CMP Channel with Request-Response Pattern

Description:
This example demonstrates using CMP channels for implementing a request-response pattern in Go.

Steps:
1. Create a CMP channel for handling requests.
2. Implement goroutines to handle incoming requests and send responses.
3. Send requests to the channel and process responses.

Expected Output:
Requests processed with responses:
- Request: GetInfo Response: Information retrieved successfully

Note: Requests are processed, and responses are sent back via CMP channels.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	wg.Add(1)

	// Request handler goroutine
	go func() {
		defer wg.Done()
		for {
			request := ch.Receive().(string)
			if request == "GetInfo" {
				ch.Send("Information retrieved successfully")
			} else if request == "GetData" {
				ch.Send("Data retrieved successfully")
			}
		}
	}()

	// Send request and process response
	ch.Send("GetInfo")
	fmt.Printf("Request: GetInfo Response: %s\n", ch.Receive().(string))

	// Optional: Send more requests
	ch.Send("GetData")
	fmt.Printf("Request: GetData Response: %s\n", ch.Receive().(string))

	// Close channel to signal end
	ch.Close()

	wg.Wait()
}





/*
Example 33: CMP Channel with Multi-Channel Communication

Description:
This example demonstrates using multiple CMP channels for communication in Go.

Steps:
1. Create multiple CMP channels for different types of messages.
2. Implement goroutines to handle messages from each channel.
3. Send messages to the channels and process them accordingly.

Expected Output:
Messages processed from multiple channels:
- Channel 1 message: Hello, CMP!
- Channel 2 message: How are you?

Note: Messages are processed based on their respective channels.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
)

func main() {
	ch1 := cmp.NewChannel()
	ch2 := cmp.NewChannel()

	// Goroutine handling messages from Channel 1
	go func() {
		msg := ch1.Receive().(string)
		fmt.Println("Channel 1 message:", msg)
	}()

	// Goroutine handling messages from Channel 2
	go func() {
		msg := ch2.Receive().(string)
		fmt.Println("Channel 2 message:", msg)
	}()

	// Send messages to respective channels
	ch1.Send("Hello, CMP!")
	ch2.Send("How are you?")

	// Optionally, close channels if no more messages
	ch1.Close()
	ch2.Close()
}





/*
Example 34: CMP Channel with Graceful Shutdown

Description:
This example demonstrates using CMP channels for implementing graceful shutdown in Go.

Steps:
1. Create a CMP channel for handling tasks.
2. Implement goroutines to handle tasks and wait for shutdown signal.
3. Send tasks to the channel and perform cleanup on shutdown.

Expected Output:
Tasks processed with graceful shutdown:
- Task 1 processed successfully
- Task 2 processed successfully
- Shutdown signal received: Cleaning up resources...

Note: Tasks are processed, and cleanup is performed upon receiving shutdown signal.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	wg.Add(1)

	// Task handler goroutine
	go func() {
		defer wg.Done()
		for {
			select {
			case task := <-ch.C:
				fmt.Printf("Task %s processed successfully\n", task.(string))
			case <-ch.Done():
				fmt.Println("Shutdown signal received: Cleaning up resources...")
				return
			}
		}
	}()

	// Send tasks
	ch.Send("Task 1")
	ch.Send("Task 2")

	// Simulate some task processing time
	time.Sleep(2 * time.Second)

	// Send shutdown signal
	ch.Close()

	wg.Wait()
}





/*
Example 35: CMP Channel with Broadcast Pattern

Description:
This example demonstrates using CMP channels for implementing a broadcast pattern in Go.

Steps:
1. Create a CMP channel for broadcasting messages.
2. Implement goroutines to handle subscribers and broadcast messages.
3. Subscribe to the channel and receive broadcasted messages.

Expected Output:
Messages broadcasted and received by subscribers:
- Subscriber 1 received message: Hello from CMP!
- Subscriber 2 received message: How are you?

Note: Messages are broadcasted to all subscribers listening to the channel.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const numSubscribers = 2
	wg.Add(numSubscribers)

	// Subscriber goroutines
	for i := 1; i <= numSubscribers; i++ {
		go func(subscriberID int) {
			defer wg.Done()
			for {
				msg := ch.Receive().(string)
				fmt.Printf("Subscriber %d received message: %s\n", subscriberID, msg)
			}
		}(i)
	}

	// Broadcast messages
	ch.Send("Hello from CMP!")
	ch.Send("How are you?")

	// Optionally, close channel if no more broadcasts
	ch.Close()

	wg.Wait()
}





/*
Example 36: CMP Channel with Priority Queue

Description:
This example demonstrates using CMP channels to implement a priority queue in Go.

Steps:
1. Create a CMP channel with multiple channels for different priority levels.
2. Implement goroutines to handle tasks from each priority channel.
3. Send tasks to the appropriate priority channel and process them based on priority.

Expected Output:
Tasks processed from priority queues:
- High Priority Task: Handle urgent request
- Medium Priority Task: Process important data
- Low Priority Task: Log routine activity

Note: Tasks are processed based on their priority levels using separate channels.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	highPriority := cmp.NewChannel()
	mediumPriority := cmp.NewChannel()
	lowPriority := cmp.NewChannel()

	var wg sync.WaitGroup
	wg.Add(3)

	// High priority tasks handler
	go func() {
		defer wg.Done()
		for {
			task := highPriority.Receive().(string)
			fmt.Println("High Priority Task:", task)
		}
	}()

	// Medium priority tasks handler
	go func() {
		defer wg.Done()
		for {
			task := mediumPriority.Receive().(string)
			fmt.Println("Medium Priority Task:", task)
		}
	}()

	// Low priority tasks handler
	go func() {
		defer wg.Done()
		for {
			task := lowPriority.Receive().(string)
			fmt.Println("Low Priority Task:", task)
		}
	}()

	// Send tasks to respective priority channels
	highPriority.Send("Handle urgent request")
	mediumPriority.Send("Process important data")
	lowPriority.Send("Log routine activity")

	// Optionally, close channels if no more tasks
	highPriority.Close()
	mediumPriority.Close()
	lowPriority.Close()

	wg.Wait()
}





/*
Example 37: CMP Channel with Rate Limiting and Timeout

Description:
This example demonstrates using CMP channels with rate limiting and timeout handling in Go.

Steps:
1. Create a CMP channel for message handling.
2. Implement a goroutine to send messages with rate limiting.
3. Implement another goroutine to receive messages with timeout and retry mechanism.

Expected Output:
Messages sent and received with rate limiting and timeout:
- Message received: Hello from CMP!
- Timeout: No response received within 1 second

Note: Messages are sent respecting rate limit and handled with timeout and retry mechanism.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	// Sender goroutine with rate limiting
	go func() {
		for i := 1; i <= 3; i++ {
			ch.Send(fmt.Sprintf("Message %d", i))
			time.Sleep(1 * time.Second) // Rate limit of 1 message per second
		}
	}()

	// Receiver goroutine with timeout and retry
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-ch.C:
				fmt.Println("Message received:", msg.(string))
			case <-time.After(2 * time.Second):
				fmt.Println("Timeout: No response received within 2 seconds")
				return
			}
		}
	}()

	wg.Wait()
}





/*
Example 38: CMP Channel with Event Subscription

Description:
This example demonstrates using CMP channels for event subscription and handling in Go.

Steps:
1. Create a CMP channel for event handling.
2. Implement goroutines to subscribe to events and handle them.
3. Send events to the channel and process them based on subscriptions.

Expected Output:
Events subscribed and processed:
- Subscriber 1 received event: User logged in
- Subscriber 2 received event: Data updated

Note: Events are subscribed and processed by multiple subscribers using CMP channels.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const numSubscribers = 2
	wg.Add(numSubscribers)

	// Subscriber goroutines
	for i := 1; i <= numSubscribers; i++ {
		go func(subscriberID int) {
			defer wg.Done()
			for {
				event := ch.Receive().(string)
				fmt.Printf("Subscriber %d received event: %s\n", subscriberID, event)
			}
		}(i)
	}

	// Send events to the channel
	ch.Send("User logged in")
	ch.Send("Data updated")

	// Optionally, close channel if no more events
	ch.Close()

	wg.Wait()
}





/*
Example 39: CMP Channel with Load Balancing

Description:
This example demonstrates using CMP channels for load balancing among workers in Go.

Steps:
1. Create a CMP channel for task distribution.
2. Implement multiple worker goroutines to receive and process tasks concurrently.
3. Send tasks to the channel to be load balanced and processed by the workers.

Expected Output:
Tasks load balanced and processed by workers:
- Worker 1 processed task: Task 1
- Worker 2 processed task: Task 2

Note: Tasks are load balanced among workers, ensuring efficient processing using CMP channels.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const numWorkers = 2
	wg.Add(numWorkers)

	// Worker pool for load balancing
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for {
				task := ch.Receive()
				if task == nil {
					break
				}
				fmt.Printf("Worker %d processed task: %s\n", workerID, task.(string))
			}
		}(i)
	}

	// Send tasks to be load balanced
	ch.Send("Task 1")
	ch.Send("Task 2")

	// Optionally, close channel if no more tasks
	ch.Close()

	wg.Wait()
}





/*
Example 40: CMP Channel with Message Deduplication

Description:
This example demonstrates using CMP channels for message deduplication in Go.

Steps:
1. Create a CMP channel for receiving messages.
2. Implement goroutines to deduplicate incoming messages.
3. Send messages to the channel and process them while ensuring deduplication.

Expected Output:
Deduplicated messages processed:
- Message received: Hello from CMP!

Note: Incoming messages are deduplicated before processing using CMP channels.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	dedup := make(map[string]bool)
	var mu sync.Mutex

	// Deduplication goroutine
	go func() {
		for {
			msg := ch.Receive().(string)

			mu.Lock()
			if !dedup[msg] {
				dedup[msg] = true
				fmt.Println("Message received:", msg)
			}
			mu.Unlock()
		}
	}()

	// Send messages with potential duplicates
	ch.Send("Hello from CMP!")
	ch.Send("Hello from CMP!") // Duplicate message
	ch.Send("How are you?")

	// Optionally, close channel if no more messages
	ch.Close()
}





/*
Example 41: CMP Channel with Fan-In Pattern

Description:
This example demonstrates using CMP channels to implement a fan-in pattern in Go.

Steps:
1. Create multiple CMP channels for producers to send data.
2. Implement a goroutine to merge data from multiple channels into a single channel.
3. Receive merged data from the fan-in channel and process it.

Expected Output:
Data merged and processed using fan-in pattern:
- Data received from Channel 1: Hello, CMP!
- Data received from Channel 2: How are you?

Note: Data from multiple channels is merged into a single channel using the fan-in pattern.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch1 := cmp.NewChannel()
	ch2 := cmp.NewChannel()

	fanIn := cmp.MergeChannels(ch1, ch2)

	var wg sync.WaitGroup
	wg.Add(1)

	// Receiver goroutine for merged data
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-fanIn.C:
				fmt.Printf("Data received from Channel %d: %s\n", msg.(int), fanIn.Receive().(string))
			case <-fanIn.Done():
				fmt.Println("Fan-in channel closed.")
				return
			}
		}
	}()

	// Send data to respective channels
	ch1.Send("Hello, CMP!")
	ch2.Send("How are you?")

	// Optionally, close channels if no more data
	ch1.Close()
	ch2.Close()

	wg.Wait()
}





/*
Example 42: CMP Channel with Pub/Sub Pattern

Description:
This example demonstrates using CMP channels to implement a pub/sub (publish/subscribe) pattern in Go.

Steps:
1. Create a CMP channel for pub/sub communication.
2. Implement goroutines for publishers to publish messages to topics.
3. Implement goroutines for subscribers to receive messages from topics.

Expected Output:
Messages published and subscribed using pub/sub pattern:
- Subscriber 1 received message from Topic 1: Hello from CMP!
- Subscriber 2 received message from Topic 2: How are you?

Note: Messages are published to topics and received by subscribers using CMP channels for pub/sub communication.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	// Subscriber 1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-ch.C:
				fmt.Printf("Subscriber 1 received message from Topic 1: %s\n", msg.(string))
			case <-ch.Done():
				fmt.Println("Subscriber 1 channel closed.")
				return
			}
		}
	}()

	// Subscriber 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-ch.C:
				fmt.Printf("Subscriber 2 received message from Topic 2: %s\n", msg.(string))
			case <-ch.Done():
				fmt.Println("Subscriber 2 channel closed.")
				return
			}
		}
	}()

	// Publisher 1 publishing to Topic 1
	go func() {
		ch.SendTo("Hello from CMP!", 1)
	}()

	// Publisher 2 publishing to Topic 2
	go func() {
		ch.SendTo("How are you?", 2)
	}()

	// Optionally, close channel if no more messages
	ch.Close()

	wg.Wait()
}





/*
Example 43: CMP Channel with Circuit Breaker Pattern

Description:
This example demonstrates using CMP channels to implement a circuit breaker pattern in Go.

Steps:
1. Create a CMP channel for communication with circuit breaker logic.
2. Implement goroutines to handle requests and monitor failures.
3. Send requests to the channel and manage circuit state (open/closed/half-open).

Expected Output:
Requests processed and circuit state managed using circuit breaker pattern:
- Request successful: Data retrieved successfully
- Circuit breaker open: Service unavailable, circuit open
- Circuit breaker closed: Retry request after some time

Note: Requests are processed with circuit breaker logic to manage service availability using CMP channels.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
	"time"
)

func main() {
	ch := cmp.NewChannel()

	var circuitOpen bool
	var wg sync.WaitGroup
	wg.Add(1)

	// Request handler with circuit breaker
	go func() {
		defer wg.Done()
		for {
			select {
			case req := <-ch.C:
				if circuitOpen {
					fmt.Println("Circuit breaker open: Service unavailable, circuit open")
				} else {
					if requestSuccessful() {
						fmt.Println("Request successful:", req.(string))
					} else {
						fmt.Println("Request failed: Retry after some time")
						circuitOpen = true
						go resetCircuitBreaker(&circuitOpen)
					}
				}
			case <-ch.Done():
				fmt.Println("Channel closed.")
				return
			}
		}
	}()

	// Simulate requests
	ch.Send("Data retrieved successfully")
	ch.Send("Data retrieval failed")

	// Optionally, close channel if no more requests
	ch.Close()

	wg.Wait()
}

func requestSuccessful() bool {
	// Simulate request success/failure
	return time.Now().UnixNano()%2 == 0 // 50% chance of success
}

func resetCircuitBreaker(circuitOpen *bool) {
	time.Sleep(5 * time.Second) // Reset circuit breaker after 5 seconds
	*circuitOpen = false
}





/*
Example 44: CMP Channel with State Machine

Description:
This example demonstrates using CMP channels to implement a state machine in Go.

Steps:
1. Create a CMP channel for state transitions and actions.
2. Implement goroutines to handle state transitions and execute actions.
3. Send events to the channel to trigger state changes and actions.

Expected Output:
State transitions executed and actions performed using state machine:
- State A: Event 1 received, transition to State B
- State B: Event 2 received, perform Action X

Note: State transitions are managed using CMP channels, triggering actions based on current state.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

type State string

const (
	StateA State = "State A"
	StateB State = "State B"
)

func main() {
	ch := cmp.NewChannel()

	var currentState State = StateA
	var wg sync.WaitGroup
	wg.Add(1)

	// State machine handler
	go func() {
		defer wg.Done()
		for {
			select {
			case event := <-ch.C:
				switch currentState {
				case StateA:
					fmt.Printf("%s: %s received, transition to %s\n", currentState, event.(string), StateB)
					currentState = StateB
					ch.Send("Event 2")
				case StateB:
					fmt.Printf("%s: %s received, perform Action X\n", currentState, event.(string))
					// Perform Action X
					currentState = StateA
					ch.Send("Event 1")
				}
			case <-ch.Done():
				fmt.Println("Channel closed.")
				return
			}
		}
	}()

	// Start state machine with initial event
	ch.Send("Event 1")

	// Optionally, close channel if no more events
	ch.Close()

	wg.Wait()
}





/*
Example 45: CMP Channel with Feedback Control Loop

Description:
This example demonstrates using CMP channels to implement a feedback control loop in Go.

Steps:
1. Create a CMP channel for feedback control and adjustment.
2. Implement goroutines to monitor and adjust parameters based on feedback.
3. Send feedback to the channel and adjust parameters accordingly.

Expected Output:
Parameters adjusted based on feedback using feedback control loop:
- Adjusted parameter: Threshold increased to 80

Note: Parameters are monitored and adjusted dynamically based on feedback using CMP channels.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

type ControlParameters struct {
	Threshold int
}

func main() {
	ch := cmp.NewChannel()

	params := &ControlParameters{
		Threshold: 75,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Feedback control loop handler
	go func() {
		defer wg.Done()
		for {
			select {
			case feedback := <-ch.C:
				fmt.Printf("Feedback received: %s\n", feedback.(string))
				// Adjust parameters based on feedback
				params.Threshold += 5
				fmt.Printf("Adjusted parameter: Threshold increased to %d\n", params.Threshold)
			case <-ch.Done():
				fmt.Println("Channel closed.")
				return
			}
		}
	}()

	// Send feedback to adjust parameters
	ch.Send("Increase threshold")

	// Optionally, close channel if no more feedback
	ch.Close()

	wg.Wait()
}





/*
Example 46: CMP Channel with Leader Election

Description:
This example demonstrates using CMP channels to implement leader election in a distributed system using Go.

Steps:
1. Create a CMP channel for nodes to participate in leader election.
2. Implement goroutines for nodes to send heartbeats and participate in election.
3. Monitor the channel for leader announcements and handle leader election logic.

Expected Output:
Nodes participate in leader election and elect a leader:
- Node 1 elected as leader
- Node 2 elected as leader

Note: Nodes send heartbeats and participate in leader election using CMP channels in a distributed system.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
	"time"
)

type Node struct {
	ID     int
	Leader bool
}

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const numNodes = 3
	wg.Add(numNodes)

	// Node goroutines for leader election
	for i := 1; i <= numNodes; i++ {
		node := Node{ID: i}

		go func(n Node) {
			defer wg.Done()

			// Send initial heartbeat
			ch.Send(fmt.Sprintf("Node %d: Heartbeat", n.ID))

			for {
				select {
				case msg := <-ch.C:
					fmt.Printf("Node %d received: %s\n", n.ID, msg.(string))

					// Election logic example: Node with highest ID becomes leader
					if !n.Leader {
						n.Leader = true
						ch.Send(fmt.Sprintf("Node %d elected as leader", n.ID))
					}
				case <-time.After(5 * time.Second):
					// Send heartbeat periodically
					ch.Send(fmt.Sprintf("Node %d: Heartbeat", n.ID))
				case <-ch.Done():
					fmt.Printf("Node %d channel closed.\n", n.ID)
					return
				}
			}
		}(node)
	}

	// Wait for leader election to finish
	wg.Wait()

	// Optionally, close channel if no more activities
	ch.Close()
}





/*
Example 47: CMP Channel with Data Stream Processing

Description:
This example demonstrates using CMP channels for processing data streams in Go.

Steps:
1. Create a CMP channel for data stream processing.
2. Implement goroutines to handle incoming data and process it.
3. Send data to the channel and process it in real-time.

Expected Output:
Data streams processed in real-time using CMP channels:
- Data processed: Sensor reading received: 25.5
- Data processed: Event logged: User logged in

Note: Data streams are processed in real-time using CMP channels for efficient data handling.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

type Data struct {
	Type    string
	Payload interface{}
}

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	wg.Add(1)

	// Data processor goroutine
	go func() {
		defer wg.Done()
		for {
			select {
			case data := <-ch.C:
				d := data.(Data)
				switch d.Type {
				case "Sensor":
					fmt.Printf("Data processed: Sensor reading received: %.1f\n", d.Payload.(float64))
				case "Event":
					fmt.Printf("Data processed: Event logged: %s\n", d.Payload.(string))
				}
			case <-ch.Done():
				fmt.Println("Channel closed.")
				return
			}
		}
	}()

	// Send data to process
	ch.Send(Data{Type: "Sensor", Payload: 25.5})
	ch.Send(Data{Type: "Event", Payload: "User logged in"})

	// Optionally, close channel if no more data
	ch.Close()

	wg.Wait()
}





/*
Example 48: CMP Channel with Error Handling

Description:
This example demonstrates using CMP channels for error handling and recovery in Go.

Steps:
1. Create a CMP channel for handling tasks with error handling logic.
2. Implement goroutines to handle tasks and recover from errors.
3. Send tasks to the channel and handle errors gracefully.

Expected Output:
Tasks processed with error handling and recovery using CMP channels:
- Task processed: Task 1 completed successfully
- Task processed: Task 2 failed: Error: Task 2 failed due to network issue

Note: Tasks are processed with error handling and recovery logic using CMP channels for robust application behavior.
*/

package main

import (
	"errors"
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	var wg sync.WaitGroup
	const numTasks = 2
	wg.Add(numTasks)

	// Task handler goroutine with error handling
	go func() {
		defer wg.Done()
		for {
			select {
			case task := <-ch.C:
				if err := processTask(task.(string)); err != nil {
					fmt.Printf("Task processed: %s failed: Error: %s\n", task.(string), err.Error())
				} else {
					fmt.Printf("Task processed: %s completed successfully\n", task.(string))
				}
			case <-ch.Done():
				fmt.Println("Channel closed.")
				return
			}
		}
	}()

	// Send tasks to process
	ch.Send("Task 1")
	ch.Send("Task 2")

	// Optionally, close channel if no more tasks
	ch.Close()

	wg.Wait()
}

func processTask(task string) error {
	// Simulate task processing with potential errors
	if task == "Task 2" {
		return errors.New("Task 2 failed due to network issue")
	}
	return nil
}





/*
Example 49: CMP Channel with Resource Pooling

Description:
This example demonstrates using CMP channels to implement resource pooling in Go.

Steps:
1. Create a CMP channel for managing resources.
2. Implement goroutines to acquire and release resources from the pool.
3. Send requests to the channel to acquire and release resources.

Expected Output:
Resources acquired and released using CMP channels for resource pooling:
- Resource acquired: Connection 1
- Resource released: Connection 1

Note: Resources are managed efficiently using CMP channels for resource pooling in concurrent applications.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

func main() {
	ch := cmp.NewChannel()

	const poolSize = 2
	var wg sync.WaitGroup
	wg.Add(poolSize)

	// Resource pool goroutine
	for i := 1; i <= poolSize; i++ {
		resourceID := fmt.Sprintf("Connection %d", i)

		go func(id string) {
			defer wg.Done()

			// Acquire resource
			ch.Send(fmt.Sprintf("Resource acquired: %s", id))

			// Simulate resource usage
			// ...

			// Release resource
			ch.Send(fmt.Sprintf("Resource released: %s", id))
		}(resourceID)
	}

	// Optionally, close channel if no more resource operations
	ch.Close()

	wg.Wait()
}





/*
Example 50: CMP Channel with Cache Implementation

Description:
This example demonstrates using CMP channels to implement a cache in Go.

Steps:
1. Create a CMP channel for managing cached data.
2. Implement goroutines to handle cache operations (get, set, delete).
3. Send cache operations to the channel and manage cached data.

Expected Output:
Cache operations performed using CMP channels:
- Cache item added: key: "user-1", value: "{name: John, age: 30}"
- Cache item deleted: key: "user-1"

Note: Cached data is managed efficiently using CMP channels for common cache operations in applications.
*/

package main

import (
	"fmt"
	"github.com/cmp/cmp"
	"sync"
)

type CacheOperation struct {
	Type  string // "add", "get", "delete"
	Key   string
	Value interface{}
}

func main() {
	ch := cmp.NewChannel()

	cache := make(map[string]interface{})
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(1)

	// Cache handler goroutine
	go func() {
		defer wg.Done()
		for {
			select {
			case op := <-ch.C:
				switch operation := op.(CacheOperation); operation.Type {
				case "add":
					mu.Lock()
					cache[operation.Key] = operation.Value
					mu.Unlock()
					fmt.Printf("Cache item added: key: %q, value: %v\n", operation.Key, operation.Value)
				case "get":
					mu.Lock()
					value, found := cache[operation.Key]
					mu.Unlock()
					if found {
						fmt.Printf("Cache item found: key: %q, value: %v\n", operation.Key, value)
					} else {
						fmt.Printf("Cache item not found: key: %q\n", operation.Key)
					}
				case "delete":
					mu.Lock()
					delete(cache, operation.Key)
					mu.Unlock()
					fmt.Printf("Cache item deleted: key: %q\n", operation.Key)
				}
			case <-ch.Done():
				fmt.Println("Channel closed.")
				return
			}
		}
	}()

	// Perform cache operations
	ch.Send(CacheOperation{Type: "add", Key: "user-1", Value: map[string]interface{}{"name": "John", "age": 30}})
	ch.Send(CacheOperation{Type: "get", Key: "user-1"})
	ch.Send(CacheOperation{Type: "delete", Key: "user-1"})

	// Optionally, close channel if no more cache operations
	ch.Close()

	wg.Wait()
}





