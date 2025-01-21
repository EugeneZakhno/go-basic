package main

import "fmt"

type EmptyItemInfoError struct{}

// Кастомные ошибки
func (e *EmptyItemInfoError) Error() string {
	return "Попытка преобразовать пустую строку к числу."
}

type ConvertToInt8Error struct {
	inputString string
}

func (e *ConvertToInt8Error) Error() string {
	return fmt.Sprintf("Невозможно преобразовать [%s] к int8.", e.inputString)
}
