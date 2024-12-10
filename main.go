package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем списокc c языками программирования
	programming_languages := []string{
		"Go", "Python", "Java", "C", "JavaScript", "Rust", "Ruby", "PHP", "Swift", "Kotlin",
	}

	// Создаем мапу для хранения данных о языках
	languageMap := make(map[string]int)

	// Мьютекс для синхронизации доступа к map
	// Используется для того, чтобы гарантировать, что только одна горутина одновременно может изменять map
	// Это делается для того чтобы не было таких случаев что одновременно две горутины начали изменять одну ячейку памяти
	var mute_language sync.Mutex

	// Группа ожидания для горутин
	var languageGroup sync.WaitGroup

	// Запускаем горутины для записи данных в map
	for _, lang := range programming_languages {
		languageGroup.Add(1)
		go func(language string) {
			defer languageGroup.Done()

			// Блокируем мьютекс для безопасного доступа к map
			mute_language.Lock()
			defer mute_language.Unlock()

			// Записываем в map
			languageMap[language] = len(language)
			fmt.Printf("Язык: %s, длина: %d\n", language, len(language))
		}(lang)
	}

	// Ждем завершения всех горутин
	languageGroup.Wait()

	// Выводим содержимое map
	fmt.Println("Результат записи в map:")
	for lang, length := range languageMap {
		fmt.Printf("%s: %d\n", lang, length)
	}
}
