package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {

	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return errors.New("Ошибка разделения строки")
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("[Parse int] %w", err)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("[Parse Duration] %w", err)
	}
	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {

	//1
	if ds.Duration <= 0 {
		return "", errors.New("Ошибка разделения строки")
	}

	//2
	distanse_km := (StepLength * float64(ds.Steps)) / 1000

	//3
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)

	//4
	str1 := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	str2 := fmt.Sprintf("Дистанция составила %.2f км.\n", distanse_km)
	str3 := fmt.Sprintf("Вы сожгли %.2f ккал.", calories)

	return str1 + str2 + str3, nil
}
