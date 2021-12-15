package application

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/inspectorvitya/x-technology-test/internal/model"
	"github.com/inspectorvitya/x-technology-test/internal/storage"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

type App struct {
	db storage.StorageStocks
}

func New(storage storage.StorageStocks) *App {
	return &App{storage}
}

func (app *App) Init(ctx context.Context) {
	go func(ctx context.Context) {
		parseNewStocks := app.parseStocks()
		emptyDb, err := app.db.CheckEmptyDb(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		if emptyDb {
			stocks, err := parseNewStocks()
			if err != nil {
				log.Fatalln(err)
			}
			err = app.db.CreateStockQuotes(ctx, stocks)
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second * 30)
		}
		for {
			stocks, err := parseNewStocks()
			if err != nil {
				log.Fatalln(err)
			}
			err = app.db.UpdateStockQuotes(ctx, stocks)
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second * 30)
		}
	}(ctx)
}

func (app *App) GetStockQuotes(ctx context.Context) ([]model.Stocks, error) {
	return app.db.GetStockQuotes(ctx)
}

func (app *App) parseStocks() func() ([]model.Stocks, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://api.blockchain.com/v3/exchange/tickers", nil)
	return func() ([]model.Stocks, error) {
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println(resp.Status)
			return nil, errors.New("bad status code")
		}
		ct := resp.Header.Get("Content-Type")
		if ct != "application/json" {
			return nil, errors.New("unexpected content-type")
		}
		var body []model.Stocks
		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			return nil, err
		}
		return body, err
	}
}
