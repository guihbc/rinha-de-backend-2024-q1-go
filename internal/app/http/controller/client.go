package controller

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http/model/request"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http/model/response"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/usecase"
	"github.com/valyala/fasthttp"
)

func ClientTrasactionController(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	id := ctx.UserValue("id").(string)
	_, err := strconv.Atoi(id)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		errorResponse := response.NewErrorResponse("o id deve ser um inteiro")
		log.Println(errorResponse.Message)
		ctx.SetBody(response.GetBytes(errorResponse))
		return
	}

	var req request.ClientTransactionRequest
	err = json.Unmarshal(ctx.Request.Body(), &req)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		errorResponse := response.NewErrorResponse("Failed to parse request body")
		log.Println(errorResponse.Message)
		ctx.SetBody(response.GetBytes(errorResponse))
		return
	}

	if err = req.ValidateFields(); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		errorResponse := response.NewErrorResponse(err.Error())
		log.Println(errorResponse.Message)
		ctx.SetBody(response.GetBytes(errorResponse))
		return
	}

	res, err := usecase.TransactionUseCase(id, req)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		errorResponse := response.NewErrorResponse(err.Error())
		log.Println(errorResponse.Message)
		ctx.SetBody(response.GetBytes(errorResponse))

		if err.Error() == "client not found" {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
		}

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(response.GetBytes(res))
}

func ClientExtractController(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	id := ctx.UserValue("id").(string)

	_, err := strconv.Atoi(id)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		errorResponse := response.NewErrorResponse("o id deve ser um inteiro")
		log.Println(errorResponse.Message)
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
