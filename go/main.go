package main

import (
	"fmt"
	"sync"
	"time"
)

type taskResult struct {
	id       int
	duration time.Duration
}

// função que executa uma tarefa hipotética, marcando o tempo de iníco e fim e recebendo um id
func task(id int, start time.Time) taskResult {
	taskStart := time.Now()
	fmt.Printf("Task %d started at [%06d]\n", id, int(time.Since(start).Seconds()))

	// sleep por 2 segundos para simular delay
	time.Sleep(2 * time.Second)

	fmt.Printf("Task %d terminated at [%06d]\n", id, int(time.Since(start).Seconds()))
	return taskResult{id: id, duration: time.Since(taskStart)}
}

// funcão que executa 3 tasks sequencialmente
func sequential(start time.Time) {
	for i := 1; i <= 3; i++ {
		result := task(i, start)
		fmt.Printf("Task %d took %s\n", result.id, result.duration.Round(time.Second))
	}
}

// função que executa 3 tasks com goroutines
func concurrent(start time.Time) {
	// cria um channel que tem 3 slots para resultados das tasks
	resultsChannel := make(chan taskResult, 3)
	// cria um waitgroup
	var wg sync.WaitGroup

	// roda 3 tasks
	for i := 1; i <= 3; i++ {
		// para cada task adiciona 1 na lista
		wg.Add(1)

		// função anonima que recebe um id
		go func(id int) {

			resultsChannel <- task(id, start) // roda a task com o id e passa o resultado dela para o channel
			defer wg.Done()            // espera todas as funções ao redor acabarem e marca como "done" a task

		}(i) // função anonima imediatamente é chamada
	}

	// espera até o waitgroup estar vazio com 0 items
	wg.Wait()
	close(resultsChannel)
	for result := range resultsChannel {
		fmt.Printf("Task %d took %s\n", result.id, result.duration.Round(time.Second))
	}

}

func main() {
	// cada task demora 2 segundos
	start := time.Now()

	fmt.Println("=== Sequential ===")
	fmt.Printf("Time: [%06d]\n", int(time.Since(start).Seconds()))
	sequential(start)
	fmt.Printf("Time: [%06d]\n", int(time.Since(start).Seconds()))

	start = time.Now()

	fmt.Println("\n=== Goroutines + Channels ===")
	fmt.Printf("Time: [%06d]\n", int(time.Since(start).Seconds()))
	concurrent(start)
	fmt.Printf("Time: [%06d]\n", int(time.Since(start).Seconds()))
}
