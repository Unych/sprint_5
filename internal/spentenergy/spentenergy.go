package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkParams(steps, weight, height, duration); err != nil {
		return 0, err
	}

	speed := MeanSpeed(steps, height, duration)
	calories := weight * speed * duration.Minutes() / minInH

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkParams(steps, weight, height, duration); err != nil {
		return 0, err
	}

	speed := MeanSpeed(steps, height, duration)

	return weight * speed * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}

func checkParams(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errors.New("steps must be greater than 0")
	}

	if weight <= 0 {
		return errors.New("weight must be greater than 0")
	}

	if height <= 0 {
		return errors.New("height must be greater than 0")
	}

	if duration <= 0 {
		return errors.New("duration must be greater than 0")
	}

	return nil
}
