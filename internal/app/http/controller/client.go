package controller

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http/model/request"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http/model/response"
	"github.com/valyala/fasthttp"
)

func ClientTrasactionController(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	id := ctx.UserValue("id").(string)
	_, err := strconv.Atoi(id)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		errorResponse := response.NewErrorResponse("o id deve ser um inteiro")
		ctx.SetBody(response.GetBytes(errorResponse))
		return
	}

	var req request.ClientTransactionRequest
	err = json.Unmarshal(ctx.Request.Body(), &req)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		errorResponse := response.NewErrorResponse("Failed to parse request body")
		ctx.SetBody(response.GetBytes(errorResponse))
		return
	}

	if err = req.ValidateFields(); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		errorResponse := response.NewErrorResponse(err.Error())
		ctx.SetBody(response.GetBytes(errorResponse))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	b, _ := json.Marshal(response.ClientTransactionResponse{
		Limit:   100000,
		Balance: -9098,
	})

	ctx.SetBody(b)
}

func ClientExtractController(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	id := ctx.UserValue("id").(string)

	_, err := strconv.Atoi(id)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		errorResponse := response.NewErrorResponse("o id deve ser um inteiro")
		ctx.SetBody(response.GetBytes(errorResponse))
		return
	}

	res := response.ClientExtractResponse{
		Balance: &response.ExtractBalance{
			Date: time.Now(),
		},
		LastTransactions: []*response.ExtractTransaction{},
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	b, _ := json.Marshal(res)
	ctx.Response.SetBody(b)
}
