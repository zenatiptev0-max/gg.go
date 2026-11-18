
package modifiers

import (
	"financialExchange/enum"
)

// ModifierForOrderBuy изменяет заявку на покупку
func ModifierForOrderBuy(sellerQuantity float64) func(enum.Order) enum.Order {
	return func(order enum.Order) enum.Order {

		// Добавляем количество проданного товара к выполненному количеству
		order.Filled += sellerQuantity

		// Проверяем, выполнена ли заявка полностью
		if order.Filled >= order.Quantity {
			// Если да, меняем статус на выполненный
			order.Status = enum.OrderStatusFilled
		} else {
			// Если нет, значит заявка выполнена только частично
			order.Status = enum.OrderStatusPartiallyFilled
		}

		// Возвращаем измененную заявку
		return order
	}
}
