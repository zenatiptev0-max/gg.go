package main

import (
	"financialExchange/store"
	"financialExchange/enum"
	"financialExchange/modifiers"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	options := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	logger := slog.New(handler)

	// currensies := store.NewStorage[string](5)
	// currensies.Add("BTC-RUB")
	// currensies.Add("ETH-RUB")
	// currensies.Add("USD-RUB")
	// currensies.Add("USD-RUB")

	// lenghtOfCurrensies := currensies.Lenght()
	// logger.Info("работа с валютами:", 
	//  slog.Int("количество уникальных комбинаций", lenghtOfCurrensies))

	// currensies.Has("ETH-RUB")
	// logger.Info("Проверка наличия элемента в хранилище:", 
	//  slog.Bool("наличие ETH-RUB", currensies.Has("ETH-RUB")))

	// currensies.Remove("ETH-RUB")
	// logger.Info("ETH-RUB удалена из хранилища")
	
	// logger.Info("работа с валютами:", 
	//  slog.Int("количество уникальных комбинаций", currensies.Lenght()))


	// traderStore := store.NewStore[uuid.UUID, float64](10)
	// traderID := uuid.New()

	// traderStore.Put(traderID, 50000.0004)
	// logger.Info("Добавлен новый трейдер в хранилище:",
	//  slog.String("UUID", traderID.String()),
	// slog.Float64("начальный баланс", 50000.0004))
	
	// if balance, ok := traderStore.Get(traderID); ok {
	// 	logger.Info("поиск в хранилище:",
	// 	slog.String("найден трейдер с UUID", traderID.String()),
	// 	slog.Float64("текущий баланс", balance))
	// }else{
	// 	logger.Info("элемент не найден в хранилище:",
	//  	slog.String("Проверьте UUID", traderID.String()))
	// }
	
	orderBook := store.NewStore[uuid.UUID, enum.Order](100)
	uuidOrder1 := uuid.New()
	uuidOrder2 := uuid.New()
	
	orderBuy := 	enum.Order{
								ID: uuidOrder1,
								TraderID: uuid.New(),
								Side: enum.SideBuy,
								Type: enum.OrderTypeLimit,
								Symbo: "USD",
								Price: 1000.0,
								Quantity: 2.00,
								Filled: 0.00,
								Status: enum.OrderStatusNew,
								CreatedAt: time.Now().UTC(),
								UpdatedAt: time.Now().UTC(),
								}

	orderSell := 	enum.Order{
								ID: uuidOrder2,
								TraderID: uuid.New(),
								Side: enum.SideSell,
								Type: enum.OrderTypeLimit,
								Symbo: "USD",
								Price: 1000.0,
								Quantity: 2.00,
								Filled: 0.00,
								Status: enum.OrderStatusNew,
								CreatedAt: time.Now().UTC(),
								UpdatedAt: time.Now().UTC(),
								}

    orderBook.Put(uuidOrder1, orderBuy)
	orderBook.Put(uuidOrder2, orderSell)							
	
	logger.Info("сделака зарегестрирована",
			slog.String("Type", orderBuy.Type.String()),
			slog.Float64("Value", orderBuy.Price),
			slog.Float64("Quantity", orderBuy.Quantity),
			slog.String("Symbo", orderBuy.Symbo),
			slog.String("TraderID", orderBuy.TraderID.String()),)


	// if orderBook.Delete(uuidOrder1) {
    //     logger.Info("заказ успешно удалён",
    //         slog.String("OrderID", uuidOrder1.String()),
    //     )
    // } else {
    //     logger.Info("Не удалось удалить заказ, так как он не был найден",
    //         slog.String("OrderID", uuidOrder1.String()),
    //     )
    // }


	// orderBook := store.NewIndexedStore[uuid.UUID, *enum.Order](100)
	
	// uuidOrder := uuid.New()

	// orderBuy := enum.Order{
	// 	ID:        uuidOrder,
	// 	TraderID:  uuid.New(),
	// 	Side:      enum.SideBuy,
	// 	Type:      enum.OrderTypeLimit,
	// 	Symbo:     "USD",
	// 	Price:     1000.0,
	// 	Quantity:  8.00,
	// 	Filled:    0.00,
	// 	Status:    enum.OrderStatusNew,
	// 	CreatedAt: time.Now().UTC(),
	// 	UpdatedAt: time.Now().UTC(),
	// }

	// orderBook.Put(uuidOrder, &orderBuy)

	// logger.Info("Заказ на покупку зарегистрирован в системе",
	// 	slog.String("OrderID", uuidOrder.String()),
	// 	slog.String("Type", orderBuy.Type.String()),
	// 	slog.Float64("TotalQuantity", orderBuy.Quantity),
	// )

	// sellOrders := []enum.Order{
	// 	{
	// 		ID:        uuid.New(),
	// 		TraderID:  uuid.New(),
	// 		Side:      enum.SideSell,
	// 		Type:      enum.OrderTypeLimit,
	// 		Symbo:     "USD",
	// 		Price:     1000.0,
	// 		Quantity:  2.00,
	// 		Filled:    0.00,
	// 		Status:    enum.OrderStatusNew,
	// 		CreatedAt: time.Now().UTC(),
	// 		UpdatedAt: time.Now().UTC(),
	// 	},
	// 	{
	// 		ID:        uuid.New(),
	// 		TraderID:  uuid.New(),
	// 		Side:      enum.SideSell,
	// 		Type:      enum.OrderTypeLimit,
	// 		Symbo:     "USD",
	// 		Price:     1000.0,
	// 		Quantity:  2.00,
	// 		Filled:    0.00,
	// 		Status:    enum.OrderStatusNew,
	// 		CreatedAt: time.Now().UTC(),
	// 		UpdatedAt: time.Now().UTC(),
	// 	},
	// 	{
	// 		ID:        uuid.New(),
	// 		TraderID:  uuid.New(),
	// 		Side:      enum.SideSell,
	// 		Type:      enum.OrderTypeLimit,
	// 		Symbo:     "USD",
	// 		Price:     1000.0,
	// 		Quantity:  2.00,
	// 		Filled:    0.00,
	// 		Status:    enum.OrderStatusNew,
	// 		CreatedAt: time.Now().UTC(),
	// 		UpdatedAt: time.Now().UTC(),
	// 	},
	// 	{
	// 		ID:        uuid.New(),
	// 		TraderID:  uuid.New(),
	// 		Side:      enum.SideSell,
	// 		Type:      enum.OrderTypeLimit,
	// 		Symbo:     "USD",
	// 		Price:     1000.0,
	// 		Quantity:  2.00,
	// 		Filled:    0.00,
	// 		Status:    enum.OrderStatusNew,
	// 		CreatedAt: time.Now().UTC(),
	// 		UpdatedAt: time.Now().UTC(),
	// 	},
	// }

	// for _, sellOrder := range sellOrders{
	// 	currentBuy, ok := orderBook.Get(uuidOrder)
	// 	if !ok{
	// 		logger.Info("Заказ на покупку не найден")
	// 		break
	// 	}

	// 	if !currentBuy.IsActive(){
	// 		logger.Info("заказ уже либо выполнен, либо отменён")
	// 	}

	// 	remainingQty := currentBuy.Remaining()

	// 	tradeQty := sellOrder.Quantity
	// 	if tradeQty > remainingQty {
	// 		tradeQty = remainingQty
	// 	}

	// 	trade := enum.Trade{
	// 		ID:          uuid.New(),
	// 		Symbol:      currentBuy.Symbo,
	// 		BuyOrderID:  currentBuy.ID,
	// 		SellOrderID: sellOrder.ID,
	// 		BuyerID:     currentBuy.TraderID,
	// 		SellerID:    sellOrder.TraderID,
	// 		Price:       currentBuy.Price,
	// 		Quantity:    tradeQty,
	// 		ExecutedAt:  time.Now().UTC(),
	// 	}

	// 	currentBuy.Filled += tradeQty
	// 	currentBuy.UpdatedAt = time.Now().UTC()

	// 	if currentBuy.Remaining() == 0 {
	// 		currentBuy.Status = enum.OrderStatusFilled
	// 	} else {
	// 		currentBuy.Status = enum.OrderStatusPartiallyFilled
	// 	}

	// 	orderBook.Put(uuidOrder, currentBuy)

	// 	logger.Info("Выполнена промежуточная сделка",
	// 		slog.String("TradeID", trade.ID.String()),
	// 		slog.Float64("TradeQuantity", trade.Quantity),
	// 		slog.Float64("FilledTotal", currentBuy.Filled),
	// 		slog.Float64("RemainingQuantity", currentBuy.Remaining()),
	// 		slog.String("CurrentStatus", currentBuy.Status.String()),
	// 	)
	// }

	// if finalBuy, ok := orderBook.Get(uuidOrder); ok {
	// 	logger.Info("Итоговое состояние заказа на покупку",
	// 		slog.String("OrderID", finalBuy.ID.String()),
	// 		slog.Float64("TotalQuantity", finalBuy.Quantity),
	// 		slog.Float64("Filled", finalBuy.Filled),
	// 		slog.Float64("Remaining", finalBuy.Remaining()),
	// 		slog.String("FinalStatus", finalBuy.Status.String()),
	// 	)
	// }

	updatedOrder, ok := orderBook.Update(uuidOrder1, modifiers.ModifierForOrderBuy(orderSell.Quantity))

	if ok{
		logger.Info("Статус сделки на покупку изменён",
				slog.String("Type", updatedOrder.Type.String()),
				slog.String("Status", updatedOrder.Status.String()),
				slog.Float64("Value", updatedOrder.Price),
				slog.Float64("Quantity", updatedOrder.Quantity),
				slog.String("Symbo", updatedOrder.Symbo),
				slog.String("TraderID", updatedOrder.TraderID.String()),
				slog.String("Date ", updatedOrder.UpdatedAt.String()))
	}

	fmt.Scanln()
}