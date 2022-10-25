package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) WebSocketHandler(c echo.Context) error {
	conn, _, err := websocket.DefaultDialer.Dial("wss://stream.binancefuture.com/ws/btcusdt@markPrice", nil)
	defer conn.Close()
	if err != nil {
		fmt.Errorf("DialErr :", err)
		return err
	}
	reqBody := &BinanceRequestBody{
		Method: "SUBSCRIBE",
		Params: [1]string{"btcusdt@aggTrade"},
		ID:     1,
	}
	err = conn.WriteJSON(*reqBody)
	if err != nil {
		fmt.Errorf("WriteJSONErr : ", err)
		return err
	}
	for {
		_, rawMsg, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		//fmt.Println("raw : ", rawMsg)
		wsMsg := &BinanceResponseBody{}
		err = json.Unmarshal(rawMsg, &wsMsg)
		if err != nil {
			fmt.Println("CastErr : ", err)
			break
		}
		fmt.Println(*wsMsg)
	}
	return nil
}
