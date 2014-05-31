package indicators_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/thetruetrade/gotrade/indicators"
)

var _ = Describe("when calculating an average true range (atr) with DOHLCV source data", func() {
	var (
		period                      int = 14
		indicator                   *indicators.ATR
		indicatorWithLookbackInputs IndicatorWithLookbackSharedSpecInputs
	)

	BeforeEach(func() {
		indicator, _ = indicators.NewATR(period)
		indicatorWithLookbackInputs = IndicatorWithLookbackSharedSpecInputs{IndicatorUnderTest: indicator,
			SourceDataLength: len(sourceDOHLCVData),
			GetMaximum: func() float64 {
				return GetDataMax(indicator.Data)
			},
			GetMinimum: func() float64 {
				return GetDataMin(indicator.Data)
			}}
	})

	Context("and the indicator has not yet received any ticks", func() {
		ShouldBeAnInitialisedIndicatorWithLookback(&indicatorWithLookbackInputs)
	})

	Context("and the indicator has received less ticks than the lookback period", func() {

		BeforeEach(func() {
			for i := 0; i < indicator.GetLookbackPeriod(); i++ {
				indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
			}
		})

		ShouldBeAnIndicatorThatHasReceivedFewerTicksThanItsLookbackPeriod(&indicatorWithLookbackInputs)
	})

	Context("and the indicator has received ticks equal to the lookback period", func() {

		BeforeEach(func() {
			for i := 0; i <= indicator.GetLookbackPeriod(); i++ {
				indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
			}
		})

		ShouldBeAnIndicatorThatHasReceivedTicksEqualToItsLookbackPeriod(&indicatorWithLookbackInputs)
	})

	Context("and the indicator has received more ticks than the lookback period", func() {

		BeforeEach(func() {
			for i := range sourceDOHLCVData {
				indicator.ReceiveDOHLCVTick(sourceDOHLCVData[i], i+1)
			}
		})

		ShouldBeAnIndicatorThatHasReceivedMoreTicksThanItsLookbackPeriod(&indicatorWithLookbackInputs)
	})
})