package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

// Person struct with two fields: Name and Age
type Person struct {
	Name string
	Age  int
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a Redis connection pool
	pool := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Your Redis server address
		DB:       0,                // Use the default DB
		PoolSize: 10,               // Set the size of the connection pool
	})

	// Check if the connection to Redis was successful
	pong, err := pool.Ping().Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// Prompt the user for their choice
	fmt.Println("Choose an option:")
	fmt.Println("Press 1 to Create 100k objects")
	fmt.Println("Press 2 to Create 200k objects")
	fmt.Println("Press 3 to Create 300k objects")
	fmt.Println("Press 4 to Create 400k objects")
	fmt.Println("Press 5 to Create 500k objects")

	var choice int
	fmt.Scan(&choice)

	// Record start time
	startTime := time.Now()

	// Create and store objects based on the user's choice
	queueName := "person_queue"
	switch choice {
	case 1:
		createAndStoreObjectsInQueue(pool, queueName, 100000)
	case 2:
		createAndStoreObjectsInQueue(pool, queueName, 200000)
	case 3:
		createAndStoreObjectsInQueue(pool, queueName, 300000)
	case 4:
		createAndStoreObjectsInQueue(pool, queueName, 400000)
	case 5:
		createAndStoreObjectsInQueue(pool, queueName, 500000)
	default:
		fmt.Println("Invalid choice. Exiting...")
		return
	}

	// Record end time
	endTime := time.Now()

	// Calculate and print the elapsed time
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Time taken to enqueue objects: %v\n", elapsedTime)
}

func createAndStoreObjectsInQueue(pool *redis.Client, queueName string, numObjects int) {
	// Record start time for insertion
	insertStartTime := time.Now()

	for i := 1; i <= numObjects; i++ {
		person := Person{
			Name: generateRandomName(),
			Age:  generateRandomAge(),
		}

		// Serialize Person object to JSON
		personJSON, err := json.Marshal(person)
		if err != nil {
			log.Println("Error marshaling JSON:", err)
			return
		}

		// Enqueue the serialized JSON into a Redis list (queue)
		_, err = pool.LPush(queueName, string(personJSON)).Result()
		if err != nil {
			log.Println("Failed to enqueue object in Redis:", err)
			return
		}
	}

	// Record end time for insertion
	insertEndTime := time.Now()

	// Calculate and print the elapsed time for insertion
	insertElapsedTime := insertEndTime.Sub(insertStartTime)
	fmt.Printf("Time taken to insert %d objects: %v\n", numObjects, insertElapsedTime)

	fmt.Printf("Enqueued %d objects in Redis\n", numObjects)
}

func generateRandomName() string {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Henry", "Ivy", "Jack"}
	return names[rand.Intn(len(names))]
}

func generateRandomAge() int {
	return rand.Intn(31) + 20
}
