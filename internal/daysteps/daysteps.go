package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат строки дневной активности")
	}

	// Без TrimSpace, чтобы строки с пробелами вызывали ошибку парсинга числа
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %w", err)
	}

	// Проверка на ноль или отрицательные шаги
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}

	// Проверка на нулевую или отрицательную продолжительность
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	// Перенос строки \n в конце для успешного прохождения тестов
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", 
		ds.Steps, distance, calories)
		
	return result, nil
}

func (ds DaySteps) ActionString() string {
	str, err := ds.ActionInfo()
	if err != nil {
		return fmt.Sprintf("Ошибка: %v", err)
	}
	return str
}