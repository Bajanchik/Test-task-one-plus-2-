package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Coin struct {
	ID        string  `json:"id"`
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Price     float64 `json:"current_price"`
	MarketCap int64   `json:"market_cap"`
}

func main() {
	// Получаем курсы всех криптовалют
	cryptoCurrencies, err := getCryptoCurrencies()
	if err != nil {
		log.Fatal("Ошибка получения криптовалют:", err)
	}

	// Печатаем курсы всех криптовалют
	fmt.Println("Курсы всех криптовалют:")
	for _, crypto := range cryptoCurrencies {
		fmt.Printf("%s (%s): $%.2f\n", crypto.Name, crypto.Symbol, crypto.Price)
	}

	// Получаем курс для определенной криптовалюты (например, BTC)
	fmt.Println("Enter crypto:")
	var symbol string
	fmt.Scan(&symbol)
	rate, err := getCurrencyRate(symbol, cryptoCurrencies)
	if err != nil {
		log.Fatal("Ошибка получения курса для криптовалюты", symbol, ":", err)
	}

	// Печатаем курс для определенной криптовалюты
	fmt.Printf("Курс для %s (%s): $%.2f\n", rate.Name, rate.Symbol, rate.Price)
}

func getCryptoCurrencies() ([]Coin, error) {
	// Отправляем GET запрос к API
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Декодируем JSON ответ в структуру данных
	var coins []Coin
	err = json.NewDecoder(resp.Body).Decode(&coins)
	if err != nil {
		return nil, err
	}

	return coins, nil
}

func getCurrencyRate(symbol string, coins []Coin) (*Coin, error) {
	// Ищем криптовалюту по символу
	for _, coin := range coins {
		if coin.Symbol == symbol {
			return &coin, nil
		}
	}

	return nil, fmt.Errorf("криптовалюта с символом %s не найдена", symbol)
}
