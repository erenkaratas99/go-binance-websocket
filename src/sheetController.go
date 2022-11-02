package main

import (
	"strconv"
	"time"
)

type LoopController struct {
	isFirst bool
	j       int
	ss      *SpreadsheetService
}

func NewLoopController(ss *SpreadsheetService) *LoopController {
	return &LoopController{
		isFirst: true,
		j:       1,
		ss:      ss,
	}
}

func (lc *LoopController) LoopObserver() {
	if lc.j >= 1 {
		lc.isFirst = false
	}
	if lc.j < 11 {
		lc.j += 1
	}
	if lc.j >= 11 || lc.j < 1 {
		lc.j = 1
	}
}

func (lc *LoopController) GoogleSheetController(msg *BinanceResponseBody) {
	lc.LoopObserver()
	tm := time.Unix(int64(msg.EventTime/1000), 0)
	jStr := strconv.Itoa(lc.j)
	sheetRange := "A" + jStr
	obj := &SpreadsheetPushRequest{
		SpreadsheetId: "1R3xn0eJ_qoKIDPM49iMoop8v92ecoxJWltqmRA0xtsI",
		Range:         sheetRange,
		Values:        []interface{}{msg.Price, tm, msg.EventType},
	}
	if lc.isFirst == false {
		lc.ss.WriteToSpreadsheet(obj)
	}
}
