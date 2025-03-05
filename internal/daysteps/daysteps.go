package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
	"strconv"
	"strings"
	"time"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 2 {
		return 0, 0, errors.New("invalid data")
	}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return 0, 0, err
	}
	duration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	steps, time, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps == 0 {
		return ""
	}
	distance := (StepLength * float64(steps)) / 1000.0
	call := spentcalories.WalkingSpentCalories(steps, weight, height, time)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f", steps, distance, call)
}
