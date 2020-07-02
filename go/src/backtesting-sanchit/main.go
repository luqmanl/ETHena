package main

import (
	"fmt"
  luno "github.com/luno/luno-go"
	"github.com/luno/luno-go/decimal"
)

func testSMA(bot *smaBot) {
	var i int64 = 0
	for i < bot.numOfDecisions {
		bot.tradeSMA()
		i++
	}
}


func testRSI(bot *rsiBot) {
	var i int64 = 0
	for i < bot.numOfDecisions {
		bot.tradeRSI()
		i++
	}
}


var client *luno.Client
var reqPointer *luno.GetTickerRequest

func main() {
	startingFunds := decimal.NewFromInt64(int64(100))

	parseXlsx()

	maxProfit := decimal.NewFromInt64(0)
	maxTradingPeriod := int64(0)
	maxOverSold := int64(0)
	maxOverBought := int64(0)

	for x := 0; x < 50; x+=10 {
		for y := 100; y > 50; y-=10 {
			i:=4
			for i < 60{
				overSold := int64(x)
				overBought := int64(y)
				tradingPeriod := int64(i)
				var numOfDecisions int64 = 50000/tradingPeriod
				//	var offset int64 = 40

				pf := portfolio{startingFunds, decimal.NewFromInt64(int64(0)), tradingPeriod/*tradingPeriod*/, tradingPeriod/*tradingPeriod*/, 0}
				//botSMA := smaBot{&pf, decimal.NewFromInt64(offset), numOfDecisions}
				bot := rsiBot{&pf, numOfDecisions, overSold, overBought}
				testRSI(&bot)
				currBid := getBid(bot.pf.currRow)
				portfolioValue := bot.pf.funds.Add(currBid.Mul(bot.pf.stock))
				profit := portfolioValue.Sub(startingFunds)

				if profit.Cmp(maxProfit) == 1 {
					maxProfit = profit
					maxOverSold = bot.overSold
					maxOverBought = bot.overBought
					maxTradingPeriod = bot.pf.tradingPeriod
					fmt.Println("maximum profit made: £", maxProfit)
					fmt.Println("at trading periods:  ", maxTradingPeriod)
					fmt.Println("upper bound: 			  ", maxOverBought)
					fmt.Println("upper Sold:  				", maxOverSold)
					fmt.Println(".")
				}
				i +=4
			}
		}
	}





	days := ((50000 / 60)/ 24)
  fmt.Println("Days: ",days)
//	fmt.Println("Profit/Loss:     £", profit)
  fmt.Println("maximum profit made: £", maxProfit)
	fmt.Println(" at trading periods:  ", maxTradingPeriod)
	fmt.Println(" 			 upper bound:  ", maxOverBought)
	fmt.Println(" 			 upper Sold:  ", maxOverSold)
//	fmt.Println(bot.pf.tradesMade," trades made")

}
