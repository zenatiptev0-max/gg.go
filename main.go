
package main

import (
	"financialExchange/enum"
	"financialExchange/modifiers"
	"financialExchange/store"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {

	// Настраиваем вывод сообщений в формате JSON
	options := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewJSONHandler(os.Stdout, &options)
	logger := slog.New(handler)

	// Создаём хранилище для заявок
	orderBook := store.NewStore[uuid.UUID, enum.Order](100)

	// Создаём ID для заявок
	buyOrderID := uuid.New()

	// Создаём заявку на покупку
	orderBuy := enum.Order{
		ID:        buyOrderID,
		TraderID:  uuid.New(),
		Side:      enum.SideBuy,
		Type:      enum.OrderTypeLimit,
		Symbo:     "USD",
		Price:     1000.0,
		Quantity:  5.0,
		Filled:    0.0,
		Status:    enum.OrderStatusNew,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// Добавляем заявку в хранилище
	orderBook.Put(buyOrderID, orderBuy)

	logger.Info("Создана заявка на покупку",
		slog.String("OrderID", buyOrderID.String()),
		slog.Float64("Quantity", orderBuy.Quantity),
		slog.Float64("Price", orderBuy.Price),
		slog.String("Symbol", orderBuy.Symbo),
	)

	// Список заявок от продавцов
	sellOrders := []enum.Order{
		{
			ID:        uuid.New(),
			TraderID:  uuid.New(),
			Side:      enum.SideSell,
			Type:      enum.OrderTypeLimit,
			Symbo:     "USD",
			Price:     1000.0,
			Quantity:  2.0,
			Filled:    0.0,
			Status:    enum.OrderStatusNew,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		{
			ID:        uuid.New(),
			TraderID:  uuid.New(),
			Side:      enum.SideSell,
			Type:      enum.OrderTypeLimit,
			Symbo:     "USD",
			Price:     1000.0,
			Quantity:  1.5,
			Filled:    0.0,
			Status:    enum.OrderStatusNew,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		{
			ID:        uuid.New(),
			TraderID:  uuid.New(),
			Side:      enum.SideSell,
			Type:      enum.OrderTypeLimit,
			Symbo:     "USD",
			Price:     1000.0,
			Quantity:  1.5,
			Filled:    0.0,
			Status:    enum.OrderStatusNew,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}

	// Перебираем все заявки на продажу
	for i, sellOrder := range sellOrders {

		// Получаем актуальную заявку на покупку
		currentOrder, ok := orderBook.Get(buyOrderID)

		if !ok {
			logger.Info("Заявка на покупку не найдена")
			break
		}

		// Считаем, сколько ещё осталось купить
		remaining := currentOrder.Quantity - currentOrder.Filled

		// Если заявка уже выполнена, дальше не продолжаем
		if remaining <= 0 {
			logger.Info("Заявка полностью выполнена")
			break
		}

		// Определяем количество для текущей сделки
		tradeQuantity := sellOrder.Quantity

		// Не позволяем купить больше, чем указано в заявке
		if tradeQuantity > remaining {
			tradeQuantity = remaining
		}

		// Обновляем заявку после продажи
		updatedOrder, ok := orderBook.Update(
			buyOrderID,
			modifiers.ModifierForOrderBuy(tradeQuantity),
		)

		if !ok {
			logger.Info("Не удалось обновить заявку")
			break
		}

		// Выводим информацию о текущей сделке
		logger.Info("Заявка на покупку обновлена",
			slog.Int("DealNumber", i+1),
			slog.Float64("SoldQuantity", tradeQuantity),
			slog.Float64("Filled", updatedOrder.Filled),
			slog.Float64("Remaining", updatedOrder.Quantity-updatedOrder.Filled),
			slog.String("Status", updatedOrder.Status.String()),
			slog.String("TraderID", updatedOrder.TraderID.String()),
		)
	}

	// Получаем итоговое состояние заявки
	finalOrder, ok := orderBook.Get(buyOrderID)

	if ok {
		logger.Info("Итоговое состояние заявки",
			slog.Float64("TotalQuantity", finalOrder.Quantity),
			slog.Float64("Filled", finalOrder.Filled),
			slog.Float64("Remaining", finalOrder.Quantity-finalOrder.Filled),
			slog.String("FinalStatus", finalOrder.Status.String()),
		)
	}

	fmt.Scanln()
}
