package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Ping Redis to check the connection.
    pong, err := rdb.Ping(ctx).Result()
    fmt.Println(pong, err)

    // Set a key-value pair.
    err = rdb.Set(ctx, "key1", "value1", 0).Err()
    if err != nil {
        panic(err)
    }

    // Get the value of the key.
    val, err := rdb.Get(ctx, "key1").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key1", val)
}





package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Push values to a list.
    err := rdb.LPush(ctx, "list1", "value1", "value2", "value3").Err()
    if err != nil {
        panic(err)
    }

    // Retrieve list elements.
    vals, err := rdb.LRange(ctx, "list1", 0, -1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("list1:", vals)
}





package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Set multiple fields in a hash.
    err := rdb.HSet(ctx, "hash1", map[string]interface{}{
        "field1": "value1",
        "field2": "value2",
    }).Err()
    if err != nil {
        panic(err)
    }

    // Retrieve hash values.
    vals, err := rdb.HGetAll(ctx, "hash1").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("hash1:", vals)
}





package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Add members to a set.
    err := rdb.SAdd(ctx, "set1", "member1", "member2", "member3").Err()
    if err != nil {
        panic(err)
    }

    // Retrieve set members.
    members, err := rdb.SMembers(ctx, "set1").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("set1 members:", members)
}





package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Add members with scores to a sorted set.
    err := rdb.ZAdd(ctx, "sortedset1", &redis.Z{
        Score:  1.0,
        Member: "member1",
    }, &redis.Z{
        Score:  2.0,
        Member: "member2",
    }).Err()
    if err != nil {
        panic(err)
    }

    // Retrieve sorted set members by rank.
    vals, err := rdb.ZRange(ctx, "sortedset1", 0, -1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("sortedset1 members:", vals)
}





package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Subscribe to a channel.
    pubsub := rdb.Subscribe(ctx, "channel1")
    defer pubsub.Close()

    // Wait for confirmation that subscription is created before publishing anything.
    _, err := pubsub.Receive(ctx)
    if err != nil {
        panic(err)
    }

    // Publish a message to the channel.
    err = rdb.Publish(ctx, "channel1", "hello world").Err()
    if err != nil {
        panic(err)
    }

    // Read message from the channel.
    msg, err := pubsub.ReceiveMessage(ctx)
    if err != nil {
        panic(err)
    }
    fmt.Println("Message received:", msg.Payload)
}





package main

import (
    "context"
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Set a key that expires in 10 seconds.
    err := rdb.Set(ctx, "key2", "value2", 10*time.Second).Err()
    if err != nil {
        panic(err)
    }

    // Retrieve the value of the key immediately.
    val, err := rdb.Get(ctx, "key2").Result()
    if err != nil {
        if err == redis.Nil {
            fmt.Println("key2 does not exist")
        } else {
            panic(err)
        }
    } else {
        fmt.Println("key2", val)
    }

    // Wait for key to expire.
    time.Sleep(11 * time.Second)

    // Try to retrieve the expired key.
    val, err = rdb.Get(ctx, "key2").Result()
    if err != nil {
        if err == redis.Nil {
            fmt.Println("key2 does not exist after expiry")
        } else {
            panic(err)
        }
    } else {
        fmt.Println("key2", val)
    }
}





package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Create a pipeline.
    pipe := rdb.Pipeline()

    // Execute multiple commands in a pipeline.
    pipe.Set(ctx, "key3", "value3", 0)
    pipe.Get(ctx, "key3")

    // Execute the pipeline and retrieve results.
    _, err := pipe.Exec(ctx)
    if err != nil {
        panic(err)
    }

    // Retrieve the value of the key from Redis.
    val, err := rdb.Get(ctx, "key3").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key3", val)
}





package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Begin a transaction.
    tx := rdb.TxPipeline()

    // Queue commands inside the transaction.
    tx.Set(ctx, "key4", "value4", 0)
    tx.Get(ctx, "key4")

    // Execute the transaction.
    _, err := tx.Exec(ctx)
    if err != nil {
        panic(err)
    }

    // Retrieve the value of the key from Redis.
    val, err := rdb.Get(ctx, "key4").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key4", val)
}





package main

import (
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Define a Lua script.
    luaScript := `
        local val = redis.call('GET', KEYS[1])
        return val
    `

    // Load the Lua script.
    script := redis.NewScript(luaScript)

    // Execute the Lua script.
    val, err := script.Run(ctx, rdb, []string{"key1"}).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("Lua script result:", val)
}





package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    // Connect to Redis Cluster.
    rdb := redis.NewClusterClient(&redis.ClusterOptions{
        Addrs: []string{"localhost:7000", "localhost:7001"}, // Redis cluster nodes
    })

    // Set a key-value pair in Redis Cluster.
    err := rdb.Set(ctx, "key5", "value5", 0).Err()
    if err != nil {
        panic(err)
    }

    // Get the value of the key from Redis Cluster.
    val, err := rdb.Get(ctx, "key5").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key5", val)
}





package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    // Connect to Redis.
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Use default DB
    })

    // Set keys for demo purpose.
    for i := 0; i < 10; i++ {
        key := fmt.Sprintf("key%d", i)
        err := rdb.Set(ctx, key, fmt.Sprintf("value%d", i), 0).Err()
        if err != nil {
            panic(err)
        }
    }

    // Use SCAN to iterate over keys.
    cursor := uint64(0)
    for {
        keys, nextCursor, err := rdb.Scan(ctx, cursor, "key*", 10).Result()
        if err != nil {
            panic(err)
        }
        for _, key := range keys {
            val, err := rdb.Get(ctx, key).Result()
            if err != nil {
                panic(err)
            }
            fmt.Printf("%s: %s\n", key, val)
        }
        cursor = nextCursor
        if cursor == 0 {
            break
        }
    }
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set bits in a bitmap.
	err := rdb.SetBit(ctx, "bitmap1", 0, 1).Err()
	if err != nil {
		panic(err)
	}

	err = rdb.SetBit(ctx, "bitmap1", 2, 1).Err()
	if err != nil {
		panic(err)
	}

	// Get bits from the bitmap.
	bit, err := rdb.GetBit(ctx, "bitmap1", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Bit at index 0:", bit)

	bit, err = rdb.GetBit(ctx, "bitmap1", 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Bit at index 1:", bit)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add elements to a HyperLogLog.
	err := rdb.PFAdd(ctx, "hll1", "elem1", "elem2", "elem3").Err()
	if err != nil {
		panic(err)
	}

	// Count unique elements in the HyperLogLog.
	count, err := rdb.PFCount(ctx, "hll1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HyperLogLog count:", count)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add locations to a Geo set.
	geo := []*redis.GeoLocation{
		{Name: "location1", Longitude: 13.361389, Latitude: 38.115556},
		{Name: "location2", Longitude: 15.087269, Latitude: 37.502669},
	}
	err := rdb.GeoAdd(ctx, "locations", geo...).Err()
	if err != nil {
		panic(err)
	}

	// Get Geo coordinates.
	coords, err := rdb.GeoPos(ctx, "locations", "location1", "location2").Result()
	if err != nil {
		panic(err)
	}
	for _, coord := range coords {
		fmt.Printf("Location: %s, Latitude: %f, Longitude: %f\n", coord.Name, coord.Latitude, coord.Longitude)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add entries to a Redis Stream.
	_, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "mystream",
		Values: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}).Result()
	if err != nil {
		panic(err)
	}

	// Read entries from the Redis Stream.
	streams, err := rdb.XRange(ctx, "mystream", "-", "+").Result()
	if err != nil {
		panic(err)
	}
	for _, msg := range streams {
		fmt.Printf("Message ID: %s, Fields: %v\n", msg.ID, msg.Values)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Define a Lua script.
	luaScript := `
        local val = redis.call('GET', KEYS[1])
        return val
    `

	// Load the Lua script.
	script := redis.NewScript(luaScript)

	// Load the script and get its SHA1 hash.
	sha1, err := script.Load(ctx, rdb).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Lua script SHA1:", sha1)

	// Execute the Lua script using EvalSha.
	val, err := rdb.EvalSha(ctx, sha1, []string{"key1"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("EvalSha result:", val)
}





package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Create Redis client options with custom configurations.
	opt := &redis.Options{
		Addr:         "localhost:6379", // Redis server address
		Password:     "",               // no password set
		DB:           0,                // use default DB
		MaxRetries:   3,                // retry up to 3 times
		DialTimeout:  5 * time.Second,  // connect timeout
		ReadTimeout:  3 * time.Second,  // read timeout
		WriteTimeout: 3 * time.Second,  // write timeout
		PoolSize:     10,               // connection pool size
		PoolTimeout:  4 * time.Second,  // connection pool timeout
	}

	// Connect to Redis with custom options.
	rdb := redis.NewClient(opt)

	// Ping Redis to check the connection.
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping:", pong)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Subscribe to a channel with a context.
	pubsub := rdb.Subscribe(ctx, "channel1")
	defer pubsub.Close()

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	// Publish a message to the channel.
	err = rdb.Publish(ctx, "channel1", "hello world").Err()
	if err != nil {
		panic(err)
	}

	// Read message from the channel.
	msg, err := pubsub.ReceiveMessage(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message received:", msg.Payload)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Create a pipeline for atomic counter.
	pipe := rdb.Pipeline()

	// Increment the counter multiple times.
	for i := 0; i < 5; i++ {
		pipe.Incr(ctx, "counter")
	}

	// Execute the pipeline and retrieve results.
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	// Retrieve the value of the counter from Redis.
	val, err := rdb.Get(ctx, "counter").Int()
	if err != nil {
		panic(err)
	}
	fmt.Println("Counter value:", val)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Define a Lua script for atomic operations.
	luaScript := `
        local current = redis.call('GET', KEYS[1])
        local amount = tonumber(ARGV[1])
        local new = current and tonumber(current) + amount or amount
        redis.call('SET', KEYS[1], new)
        return new
    `

	// Load the Lua script.
	script := redis.NewScript(luaScript)

	// Execute the Lua script atomically.
	val, err := script.Run(ctx, rdb, []string{"counter"}, 5).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Counter value after Lua script:", val)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Begin a transaction.
	tx := rdb.TxPipeline()

	// Queue commands inside the transaction.
	tx.Set(ctx, "key1", "value1", 0)
	tx.Set(ctx, "key2", "value2", 0)

	// Execute the transaction.
	_, err := tx.Exec(ctx)
	if err != nil {
		panic(err)
	}

	// Retrieve values of keys from Redis.
	val1, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", val1)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key2", val2)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Create a context with cancellation.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a pipeline with context.
	pipe := rdb.PipelineWithContext(ctx)

	// Set a key-value pair.
	pipe.Set(ctx, "key3", "value3", 0)

	// Get the value of the key.
	pipe.Get(ctx, "key3")

	// Execute the pipeline and retrieve results.
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	// Access results from individual commands.
	setResult := cmds[0].(*redis.StatusCmd)
	fmt.Println("SET result:", setResult)

	getResult := cmds[1].(*redis.StringCmd)
	val, err := getResult.Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("GET result:", val)
}





package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a connection pool.
	pool := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // no password set
		DB:       0,                // use default DB
		PoolSize: 10,               // connection pool size
	})

	// Close the connection pool after program execution.
	defer pool.Close()

	// Connect to Redis using the connection pool.
	ctx := context.Background()
	err := pool.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	// Example operations with connection pooling.
	err = pool.Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := pool.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", val)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Key to watch for transactions.
	key := "watch_key"

	// Begin a transaction with Watch.
	tx := rdb.Watch(context.Background(), func(tx *redis.Tx) error {
		// Get the current value of the key.
		currVal, err := tx.Get(context.Background(), key).Result()
		if err != nil && err != redis.Nil {
			return err
		}

		// Perform transaction operations.
		pipe := tx.TxPipeline()
		pipe.Set(context.Background(), key, "new_value", 0)

		// Execute the transaction.
		_, err = pipe.Exec(context.Background())
		if err != nil {
			return err
		}

		return nil
	}, key)

	// Check for transaction errors.
	if tx.Err() != nil {
		panic(tx.Err())
	}

	// Retrieve the new value of the key from Redis.
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("New value of key:", val)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add members with scores to a sorted set.
	err := rdb.ZAdd(context.Background(), "sortedset1", &redis.Z{
		Score:  1.0,
		Member: "member1",
	}, &redis.Z{
		Score:  2.0,
		Member: "member2",
	}).Err()
	if err != nil {
		panic(err)
	}

	// Retrieve sorted set members with scores.
	vals, err := rdb.ZRangeWithScores(context.Background(), "sortedset1", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sorted set members with scores:")
	for _, z := range vals {
		fmt.Printf("Member: %s, Score: %f\n", z.Member, z.Score)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set keys for demo purpose.
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		err := rdb.Set(context.Background(), key, fmt.Sprintf("value%d", i), 0).Err()
		if err != nil {
			panic(err)
		}
	}

	// Use SCAN to iterate over keys.
	var cursor uint64
	var keys []string
	for {
		var err error
		keys, cursor, err = rdb.Scan(context.Background(), cursor, "key*", 10).Result()
		if err != nil {
			panic(err)
		}
		for _, key := range keys {
			val, err := rdb.Get(context.Background(), key).Result()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s: %s\n", key, val)
		}
		if cursor == 0 {
			break
		}
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set multiple fields in a hash.
	err := rdb.HSet(context.Background(), "hash1", map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}).Err()
	if err != nil {
		panic(err)
	}

	// Retrieve all fields and values of the hash.
	vals, err := rdb.HGetAll(context.Background(), "hash1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All fields and values of hash1:")
	for field, value := range vals {
		fmt.Printf("%s: %s\n", field, value)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
			Addr: "localhost:6379", // Redis server address
			DB:   0,                // Use default DB
	})

	// Add members to a set.
	err := rdb.SAdd(context.Background(), "set1", "member1", "member2", "member3").Err()
	if err != nil {
			panic(err)
	}

	// Retrieve all members of the set.
	vals, err := rdb.SMembers(context.Background(), "set1").Result()
	if err != nil {
			panic(err)
	}
	fmt.Println("All members of set1:")
	for _, member := range vals {
			fmt.Println(member)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add members with scores to a sorted set.
	err := rdb.ZAdd(context.Background(), "sortedset2", &redis.Z{
		Score:  1.0,
		Member: "member1",
	}, &redis.Z{
		Score:  2.0,
		Member: "member2",
	}).Err()
	if err != nil {
		panic(err)
	}

	// Retrieve members within a specific score range from the sorted set.
	vals, err := rdb.ZRangeByScore(context.Background(), "sortedset2", &redis.ZRangeBy{
		Min: "1",
		Max: "2",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members within score range in sortedset2:")
	for _, member := range vals {
		fmt.Println(member)
	}
}





package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Create a user object.
	user := User{
		ID:   "1",
		Name: "John Doe",
		Age:  30,
	}

	// Marshal user object to JSON and set it in Redis.
	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	err = rdb.Set(context.Background(), "user1", userJSON, 0).Err()
	if err != nil {
		panic(err)
	}

	// Get JSON data from Redis and unmarshal it into a user object.
	val, err := rdb.Get(context.Background(), "user1").Result()
	if err != nil {
		panic(err)
	}
	var retrievedUser User
	err = json.Unmarshal([]byte(val), &retrievedUser)
	if err != nil {
		panic(err)
	}
	fmt.Println("Retrieved user from Redis:")
	fmt.Printf("ID: %s, Name: %s, Age: %d\n", retrievedUser.ID, retrievedUser.Name, retrievedUser.Age)
}





package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Push items into a list.
	err := rdb.LPush(context.Background(), "list1", "item1", "item2", "item3").Err()
	if err != nil {
		panic(err)
	}

	// Pop items from the list with BRPOP, which blocks until an item is available.
	for {
		result, err := rdb.BRPop(context.Background(), 0*time.Second, "list1").Result()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Popped item: %v\n", result)
	}
}





package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set a key that expires in 5 seconds using SETEX.
	err := rdb.SetEX(context.Background(), "key2", "value2", 5*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// Retrieve the value of the key immediately.
	val, err := rdb.Get(context.Background(), "key2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key2", val)

	// Wait for key to expire.
	time.Sleep(6 * time.Second)

	// Try to retrieve the expired key.
	val, err = rdb.Get(context.Background(), "key2").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key2 does not exist after expiry")
		} else {
			panic(err)
		}
	} else {
		fmt.Println("key2", val)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Increment a counter using INCR.
	_, err := rdb.Incr(context.Background(), "counter").Result()
	if err != nil {
		panic(err)
	}

	// Retrieve the value of the counter.
	val, err := rdb.Get(context.Background(), "counter").Int()
	if err != nil {
		panic(err)
	}
	fmt.Println("Counter value:", val)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Ping Redis to check connectivity.
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping:", pong)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set a key in Redis.
	err := rdb.Set(context.Background(), "key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	// Check if key exists using EXISTS.
	exists, err := rdb.Exists(context.Background(), "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Key exists:", exists)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set multiple fields in a hash using HMSET.
	err := rdb.HMSet(context.Background(), "hash2", map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}).Err()
	if err != nil {
		panic(err)
	}

	// Retrieve all fields and values of the hash using HGETALL.
	vals, err := rdb.HGetAll(context.Background(), "hash2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All fields and values of hash2:")
	for field, value := range vals {
		fmt.Printf("%s: %s\n", field, value)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set a field in a hash only if it does not exist using HSETNX.
	set, err := rdb.HSetNX(context.Background(), "hash3", "field1", "value1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HSETNX result:", set)

	// Attempt to set the same field again.
	set, err = rdb.HSetNX(context.Background(), "hash3", "field1", "new_value").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HSETNX result:", set)

	// Retrieve the value of the field.
	val, err := rdb.HGet(context.Background(), "hash3", "field1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Value of field1 in hash3:", val)
}





package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set a key with an expiration using SETEX.
	err := rdb.SetEX(context.Background(), "key3", "value3", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// Get the TTL (time-to-live) of the key.
	ttl, err := rdb.TTL(context.Background(), "key3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("TTL of key3:", ttl)

	// Get the PTTL (time-to-live in milliseconds) of the key.
	pttl, err := rdb.PTTL(context.Background(), "key3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("PTTL of key3:", pttl)

	// Sleep for 11 seconds to let the key expire.
	time.Sleep(11 * time.Second)

	// Check if the key still exists.
	exists, err := rdb.Exists(context.Background(), "key3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Key3 exists:", exists)
}





package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Create a subscriber.
	pubsub := rdb.Subscribe(context.Background(), "channel1")

	// Go routine to handle messages received.
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				panic(err)
			}
			fmt.Printf("Message received on channel %s: %s\n", msg.Channel, msg.Payload)
		}
	}()

	// Publish messages to the channel.
	for i := 0; i < 5; i++ {
		err := rdb.Publish(context.Background(), "channel1", fmt.Sprintf("message%d", i+1)).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second) // Add delay between publishes for demonstration
	}

	// Unsubscribe from the channel.
	err := pubsub.Unsubscribe(context.Background(), "channel1")
	if err != nil {
		panic(err)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Push items into a list.
	err := rdb.LPush(context.Background(), "list2", "item1", "item2", "item3").Err()
	if err != nil {
		panic(err)
	}

	// Retrieve the length of the list.
	length, err := rdb.LLen(context.Background(), "list2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Length of list2:", length)

	// Retrieve an item from the list by index.
	item, err := rdb.LIndex(context.Background(), "list2", 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Item at index 1 in list2:", item)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add members with initial scores to a sorted set.
	err := rdb.ZAdd(context.Background(), "sortedset3", &redis.Z{
		Score:  1.0,
		Member: "member1",
	}, &redis.Z{
		Score:  2.0,
		Member: "member2",
	}).Err()
	if err != nil {
		panic(err)
	}

	// Increment the score of a member in the sorted set using ZINCRBY.
	newScore, err := rdb.ZIncrBy(context.Background(), "sortedset3", 2.5, "member1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("New score of member1 in sortedset3:", newScore)

	// Retrieve members with scores from the sorted set.
	vals, err := rdb.ZRangeWithScores(context.Background(), "sortedset3", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members with scores in sortedset3:")
	for _, z := range vals {
		fmt.Printf("Member: %s, Score: %f\n", z.Member, z.Score)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Lua script to increment a counter.
	luaScript := `
        local current = redis.call('GET', KEYS[1])
        local amount = tonumber(ARGV[1])
        local new = current and tonumber(current) + amount or amount
        redis.call('SET', KEYS[1], new)
        return new
    `

	// Load the Lua script.
	script := redis.NewScript(luaScript)

	// Get the SHA1 digest of the Lua script.
	sha1, err := script.Load(context.Background(), rdb).Result()
	if err != nil {
		panic(err)
	}

	// Execute the Lua script using EVALSHA.
	val, err := rdb.EvalSha(context.Background(), sha1, []string{"counter"}, 5).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Counter value after Lua script with EVALSHA:", val)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set a key in Redis.
	err := rdb.Set(context.Background(), "oldkey", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	// Rename the key.
	err = rdb.Rename(context.Background(), "oldkey", "newkey").Err()
	if err != nil {
		panic(err)
	}

	// Retrieve the value of the new key.
	val, err := rdb.Get(context.Background(), "newkey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Value of newkey:", val)

	// Attempt to retrieve the old key.
	val, err = rdb.Get(context.Background(), "oldkey").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("oldkey does not exist")
		} else {
			panic(err)
		}
	} else {
		fmt.Println("Value of oldkey:", val)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Set multiple fields in a hash.
	err := rdb.HSet(context.Background(), "hash4", map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}).Err()
	if err != nil {
		panic(err)
	}

	// Delete a field from the hash using HDEL.
	deletedCount, err := rdb.HDel(context.Background(), "hash4", "field1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted fields count:", deletedCount)

	// Check if the deleted field exists.
	exists, err := rdb.HExists(context.Background(), "hash4", "field1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Field1 exists in hash4:", exists)
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add members to a sorted set.
	err := rdb.ZAdd(context.Background(), "sortedset4", &redis.Z{
		Score:  1.0,
		Member: "member1",
	}, &redis.Z{
		Score:  2.0,
		Member: "member2",
	}).Err()
	if err != nil {
		panic(err)
	}

	// Remove a member from the sorted set using ZREM.
	removedCount, err := rdb.ZRem(context.Background(), "sortedset4", "member1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Removed members count:", removedCount)

	// Retrieve all members of the sorted set.
	vals, err := rdb.ZRange(context.Background(), "sortedset4", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members of sortedset4 after removal:")
	for _, member := range vals {
		fmt.Println(member)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add members to two different sets.
	err := rdb.SAdd(context.Background(), "set2", "member1", "member2", "member3").Err()
	if err != nil {
		panic(err)
	}
	err = rdb.SAdd(context.Background(), "set3", "member2", "member3", "member4").Err()
	if err != nil {
		panic(err)
	}

	// Calculate the difference between two sets using SDIFF.
	diff, err := rdb.SDiff(context.Background(), "set2", "set3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Difference between set2 and set3:")
	for _, member := range diff {
		fmt.Println(member)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add members to two different sets.
	err := rdb.SAdd(context.Background(), "set4", "member1", "member2", "member3").Err()
	if err != nil {
		panic(err)
	}
	err = rdb.SAdd(context.Background(), "set5", "member2", "member3", "member4").Err()
	if err != nil {
		panic(err)
	}

	// Calculate the intersection of two sets and store the result using SINTERSTORE.
	intersectCount, err := rdb.SInterStore(context.Background(), "set_intersection", "set4", "set5").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Intersection members count:", intersectCount)

	// Retrieve all members of the intersection set.
	vals, err := rdb.SMembers(context.Background(), "set_intersection").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members of set_intersection:")
	for _, member := range vals {
		fmt.Println(member)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add members to two different sets.
	err := rdb.SAdd(context.Background(), "set6", "member1", "member2", "member3").Err()
	if err != nil {
		panic(err)
	}
	err = rdb.SAdd(context.Background(), "set7", "member2", "member3", "member4").Err()
	if err != nil {
		panic(err)
	}

	// Calculate the union of two sets and store the result using SUNIONSTORE.
	unionCount, err := rdb.SUnionStore(context.Background(), "set_union", "set6", "set7").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Union members count:", unionCount)

	// Retrieve all members of the union set.
	vals, err := rdb.SMembers(context.Background(), "set_union").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members of set_union:")
	for _, member := range vals {
		fmt.Println(member)
	}
}





package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Use default DB
	})

	// Add members with scores to a sorted set.
	err := rdb.ZAdd(context.Background(), "sortedset5", &redis.Z{
		Score:  1.0,
		Member: "member1",
	}, &redis.Z{
		Score:  2.0,
		Member: "member2",
	}, &redis.Z{
		Score:  3.0,
		Member: "member3",
	}).Err()
	if err != nil {
		panic(err)
	}

	// Retrieve members in reverse order from the sorted set using ZREVRANGE.
	vals, err := rdb.ZRevRange(context.Background(), "sortedset5", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members of sortedset5 in reverse order:")
	for _, member := range vals {
		fmt.Println(member)
	}
}

