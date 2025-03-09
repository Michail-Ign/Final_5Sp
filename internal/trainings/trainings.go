package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {

	slice := strings.Split(datastring, ",")

	if len(slice) != 3 {
		return errors.New("Ошибка разделения строки")
	}

	step, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("[Parse int] %w", err)
	}
	t.Steps = step

	str_activ := slice[1]
	if str_activ != "Бег" && str_activ != "Ходьба" {
		return errors.New("Неизвестный тип тренировки!")
	}
	t.TrainingType = str_activ

	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("[Parse Duration] %w", err)
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {

	//1
	dist_km := spentenergy.Distance(t.Steps)

	//2
	if t.Duration == 0 {
		return "", errors.New("Продолжительность тренировки должна быть больше 0!")
	}

	//3
	speed_avg := spentenergy.МeanSpeed(t.Steps, t.Duration)

	//4
	var ccal float64

	switch t.TrainingType {
	case "Ходьба":
		ccal = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Бег":
		ccal = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	default:
		return "", errors.New("unknown training type")
	}

	//5
	str1 := fmt.Sprintf("Тип тренировки: %s.\n", t.TrainingType)
	str2 := fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	str3 := fmt.Sprintf("Дистанция: %.2f км.\n", dist_km)
	str4 := fmt.Sprintf("Скорость: %.2f км/ч\n", speed_avg)
	str5 := fmt.Sprintf("Сожгли калорий: %.2f", ccal)

	return str1 + str2 + str3 + str4 + str5, nil
}
