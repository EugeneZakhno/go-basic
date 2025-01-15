package main

import "fmt"

type EmptyItemInfoError struct{}

// Кастомные ошибки
func (e *EmptyItemInfoError) Error() string {
	return "Вы ввели пустую строку."
}

type NoPriceInItemInfoError struct {
	itemInfo string
}

func (e *NoPriceInItemInfoError) Error() string {
	return fmt.Sprintf("В строке [%s] не указана стоимость.", e.itemInfo)
}
