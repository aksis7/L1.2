package main

import (
	"fmt"
	"sync"
)

func main() {
	// Исходный массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Канал для передачи результатов
	results := make(chan int, len(numbers))

	// Группа ожидания для синхронизации горутин
	var wg sync.WaitGroup

	// Запускаем вычисления в отдельных горутинах
	for _, num := range numbers {
		//увеличить счетчик на 1
		wg.Add(1)
		go func(n int) {
			// горутин завершает свою задачу, она сообщает об этом WaitGroup(уменьшить счетчик на 1)
			defer wg.Done()
			results <- n * n // Отправляем квадрат числа в канал
		}(num)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Читаем и выводим результаты из канала
	for square := range results {
		fmt.Println(square)
	}
}
