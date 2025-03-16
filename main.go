package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

const (
	buttonWidth  = 100  // Ширина кнопки
	buttonHeight = 50   // Высота кнопки
	windowWidth  = 1728 // Ширина окна
	windowHeight = 1117 // Высота окна
)

var (
	resultsText = "" // Переменная для хранения текста результатов
)

func main() {
	// Создаем приложение и главное окно
	myApp := app.New()
	myWindow := myApp.NewWindow("Эксперимент по закону Фиттса")

	// Создаем кнопку
	button := widget.NewButton("Click Me", func() {})
	button.Resize(fyne.NewSize(buttonWidth, buttonHeight))
	button.Move(fyne.NewPos(50, (windowHeight-buttonHeight)/2)) // Кнопка ближе к левому краю

	// Создаем текстовую область для вывода результатов
	results := widget.NewLabel("Результаты появятся здесь\n")
	results.Wrapping = fyne.TextWrapWord // Включение переноса текста

	// Создаем контейнер с прокруткой для результатов
	resultsScroll := container.NewScroll(results)
	resultsScroll.SetMinSize(fyne.NewSize(windowWidth-20, 200)) // Устанавливаем высоту 200 пикселей

	// Создаем выпадающий список для выбора уровня
	levelSelector := widget.NewSelect([]string{"Уровень 1", "Уровень 2", "Уровень 3"}, func(selected string) {
		results.SetText("") // Очищаем результаты при смене уровня
		resultsText = ""    // Очищаем текстовую переменную
		button.Show()       // Показываем кнопку
		switch selected {
		case "Уровень 1":
			go startLevel1(button, results, myWindow) // Запуск уровня 1
		case "Уровень 2":
			go startLevel2(button, results, myWindow) // Запуск уровня 2
		case "Уровень 3":
			go startLevel3(button, results, myWindow) // Запуск уровня 3
		}
	})
	levelSelector.SetSelected("Уровень 1") // Устанавливаем уровень по умолчанию

	// Создаем контейнер для размещения элементов
	content := container.NewVBox(
		levelSelector,                      // Выбор уровня
		container.NewWithoutLayout(button), // Кнопка
		resultsScroll,                      // Прокручиваемая область для результатов
	)

	// Устанавливаем контейнер в окно
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(windowWidth, windowHeight)) // Размер окна

	// Показываем окно и запускаем приложение
	myWindow.ShowAndRun()
}

// Функция для расчета времени по формуле Фиттса
func calculateFittsTime(distance, width float64) float64 {
	return 50 + 150*math.Log2(distance/width+1)
}

// Уровень 1: Курсор появляется на фиксированном расстоянии от кнопки
func startLevel1(button *widget.Button, results *widget.Label, window fyne.Window) {
	for i := 0; i < 10; i++ {
		// Случайная задержка перед появлением курсора
		delay := time.Duration(rand.Intn(3000)) * time.Millisecond
		time.Sleep(delay)

		// Фиксированное расстояние от кнопки
		cursorX := 400                               // Фиксированная позиция по X
		cursorY := (windowHeight - buttonHeight) / 2 // Та же высота, что и у кнопки

		// Перемещаем курсор в начальную позицию
		robotgo.Move(cursorX, cursorY)

		// Фиксируем время появления курсора
		startTime := time.Now()

		// Ждем клика по кнопке
		clicked := false
		button.OnTapped = func() {
			if !clicked {
				clicked = true

				// Фиксируем время клика
				reactionTime := time.Since(startTime)

				// Вычисляем время по формуле Фиттса
				distance := float64(cursorX - 50) // Расстояние от начальной позиции до кнопки
				width := float64(buttonWidth)
				fittsTime := calculateFittsTime(distance, width)

				// Обновляем результаты
				resultsText += fmt.Sprintf("Нажатие №%d\nВремя реакции: %v\nВремя по Фиттсу: %v мс\n\n", i+1, reactionTime, fittsTime)
				results.SetText(resultsText)
			}
		}

		// Ждем, пока пользователь нажмет кнопку
		for !clicked {
			time.Sleep(100 * time.Millisecond)
		}
	}

	// После 10 попыток скрываем кнопку и выводим сообщение
	button.Hide()
	resultsText += "Выберите уровень или начните заново.\n"
	results.SetText(resultsText)
}

// Уровень 2: Курсор появляется на случайном расстоянии от кнопки, но на одной линии
func startLevel2(button *widget.Button, results *widget.Label, window fyne.Window) {
	for i := 0; i < 10; i++ {
		// Случайная задержка перед появлением курсора
		delay := time.Duration(rand.Intn(3000)) * time.Millisecond
		time.Sleep(delay)

		// Случайное расстояние от курсора до кнопки
		distance := rand.Intn(600) + 100             // От 100 до 700 пикселей
		cursorX := 50 + distance                     // Начальная позиция курсора
		cursorY := (windowHeight - buttonHeight) / 2 // Та же высота, что и у кнопки

		// Перемещаем курсор в начальную позицию
		robotgo.Move(cursorX, cursorY)

		// Фиксируем время появления курсора
		startTime := time.Now()

		// Ждем клика по кнопке
		clicked := false
		button.OnTapped = func() {
			if !clicked {
				clicked = true

				// Фиксируем время клика
				reactionTime := time.Since(startTime)

				// Вычисляем время по формуле Фиттса
				width := float64(buttonWidth)
				fittsTime := calculateFittsTime(float64(distance), width)

				// Обновляем результаты
				resultsText += fmt.Sprintf("Нажатие №%d\nВремя реакции: %v\nВремя по Фиттсу: %v мс\nРасстояние: %v px\n\n", i+1, reactionTime, fittsTime, distance)
				results.SetText(resultsText)
			}
		}

		// Ждем, пока пользователь нажмет кнопку
		for !clicked {
			time.Sleep(100 * time.Millisecond)
		}
	}

	// После 10 попыток скрываем кнопку и выводим сообщение
	button.Hide()
	resultsText += "Выберите уровень или начните заново.\n"
	results.SetText(resultsText)
}

// Уровень 3: Курсор появляется в произвольном месте окна
func startLevel3(button *widget.Button, results *widget.Label, window fyne.Window) {
	for i := 0; i < 10; i++ {
		// Случайная задержка перед появлением курсора
		delay := time.Duration(rand.Intn(3000)) * time.Millisecond
		time.Sleep(delay)

		// Случайное положение курсора в пределах окна
		cursorX := rand.Intn(windowWidth - buttonWidth)
		cursorY := rand.Intn(windowHeight - buttonHeight)

		// Перемещаем курсор в начальную позицию
		robotgo.Move(cursorX, cursorY)

		// Вычисляем расстояние до кнопки
		buttonX := 50 // Кнопка находится на фиксированной позиции
		buttonY := (windowHeight - buttonHeight) / 2
		distance := math.Sqrt(math.Pow(float64(cursorX-buttonX), 2) + math.Pow(float64(cursorY-buttonY), 2))

		// Фиксируем время появления курсора
		startTime := time.Now()

		// Ждем клика по кнопке
		clicked := false
		button.OnTapped = func() {
			if !clicked {
				clicked = true

				// Фиксируем время клика
				reactionTime := time.Since(startTime)

				// Вычисляем время по формуле Фиттса
				width := float64(buttonWidth)
				fittsTime := calculateFittsTime(distance, width)

				// Обновляем результаты
				resultsText += fmt.Sprintf("Нажатие №%d\nВремя реакции: %v\nВремя по Фиттсу: %v мс\nРасстояние: %v px\n\n", i+1, reactionTime, fittsTime, distance)
				results.SetText(resultsText)
			}
		}

		// Ждем, пока пользователь нажмет кнопку
		for !clicked {
			time.Sleep(100 * time.Millisecond)
		}
	}

	// После 10 попыток скрываем кнопку и выводим сообщение
	button.Hide()
	resultsText += "Выберите уровень или начните заново.\n"
	results.SetText(resultsText)
}
