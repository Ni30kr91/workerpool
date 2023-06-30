package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data interface{}
}

func worker(id int, tasks <-chan Task, results chan<- Task, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task.ID)
		time.Sleep(time.Second)

		fmt.Printf("Worker %d finished task %d\n", id, task.ID)
		results <- task
	}
	wg.Done()
}

func main() {
	numTasks := 10  // Set the number of tasks here
	numWorkers := 3 // Set the number of workers here

	tasks := make(chan Task, numTasks)
	results := make(chan Task, numTasks)
	wg := sync.WaitGroup{}

	// Start worker pool
	startTime := time.Now()
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Send tasks to the worker pool
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i, Data: fmt.Sprintf("Task Data %d", i)}
	}
	close(tasks)

	// Wait for all tasks to be completed
	wg.Wait()

	// Close the results channel after all tasks are processed
	close(results)
	timeTaken := time.Since(startTime)
	// Display the tasks that are running (results received from the workers)
	for result := range results {
		fmt.Printf("Received result from task %d\n", result.ID)
	}

	// Display the total time taken to finish the tasks
	fmt.Println("All tasks completed. Total time taken: ", timeTaken)

}
