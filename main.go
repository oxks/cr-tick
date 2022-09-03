package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	mdb "cr-tick/mdb"
	priceLogger "cr-tick/postgres/sqlc"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func wsWsAllMiniMarketsStatServeHandler(event binance.WsAllMiniMarketsStatEvent) {

	ctx := context.Background()
	db := mdb.DBCon

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	defer tx.Rollback()

	queries := priceLogger.New(db).WithTx(tx)

	var args priceLogger.AddPriceParams

	for _, v := range event {
		args.Event = v.Event
		args.Symbol = v.Symbol
		args.Time = v.Time
		args.LastPrice = v.LastPrice
		args.OpenPrice = v.OpenPrice
		args.HighPrice = v.HighPrice
		args.LowPrice = v.LowPrice
		args.BaseVolume = v.BaseVolume
		args.LowPrice = v.LowPrice
		args.QuoteVolume = v.QuoteVolume

		err = queries.AddPrice(ctx, args)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}

	//write time to healthcheck
	out := strconv.Itoa(int(time.Now().Unix()))
	os.WriteFile("healthcheck", []byte(out), 0644)

}

func main() {
	fmt.Println("Price logger is started.")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	mdb.DBCon = mdb.GetDB()

	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsAllMiniMarketsStatServe(wsWsAllMiniMarketsStatServeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
	fmt.Println("Price logger is stopped.")
}
