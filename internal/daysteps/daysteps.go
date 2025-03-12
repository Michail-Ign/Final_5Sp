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
		return errors.New("line splitting errors")
	}

	ds.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("[Parse int] %w", err)
	}

	ds.Duration, err = time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("[Parse Duration] %w", err)
	}

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {

	//1
	if ds.Duration <= 0 {
		return "", errors.New("line splitting errors")
	}

	//2
	distanse_km := spentenergy.Distance(ds.Steps)

	//3
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)

	//4
	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, distanse_km, calories)
	
	return str, nil
}
