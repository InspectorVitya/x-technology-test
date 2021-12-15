package httpserver

import (
	"encoding/json"
	"github.com/inspectorvitya/x-technology-test/internal/model"
	"net/http"
)

func (s *Server) GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := s.App.GetStockQuotes(r.Context())
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp := make(model.Response, len(stocks))
	for _, val := range stocks {
		resp[val.Symbol] = model.ResponseStocks{
			Price:     val.Price,
			Volume:    val.Volume,
			LastTrade: val.LastTradePrice,
		}

	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}
