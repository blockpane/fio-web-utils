package stat

import (
	"github.com/blockpane/fio-web-utils/stats-api/models"
	ops "github.com/blockpane/fio-web-utils/stats-api/restapi/operations"
	"github.com/fioprotocol/fio-go/eos"
	"github.com/go-openapi/runtime/middleware"
	"math/big"
)

func fee(actionName bool, format string) middleware.Responder {
	if !state.ready() {
		return &ops.GetFeeServiceUnavailable{
			Payload: &models.Error{
				Code:    503,
				Message: "data is not up to date, please try again later",
			},
		}
	}
	payload := &ops.GetFeeOK{Payload: make([]*models.Price, 0)}
	for i := range state.Fees {
		ep := state.Fees[i].EndPoint
		if actionName {
			feeMapMux.RLock()
			ep = feeMap[ep]
			feeMapMux.RUnlock()
			if ep == "" {
				ep = state.Fees[i].EndPoint
			}
		}
		var price float64
		var value uint64
		switch format {
		case "suf":
			value = state.Fees[i].SufAmount
		case "float":
			price = float64(state.Fees[i].SufAmount) / 1_000_000_000.0
		case "usd":
			price = float64(state.Fees[i].SufAmount) / 1_000_000_000.0 * state.Price
		}
		payload.Payload = append(payload.Payload, &models.Price{
			EndPoint: &ep,
			Price:    price,
			Value:    value,
		})
	}
	return payload
}

func Fee(p ops.GetFeeParams) middleware.Responder {
	return fee(false, *p.Format)
}

func FeeByActionName(p ops.GetFeeByActionNameParams) middleware.Responder {
	return fee(true, *p.Format)
}

func FeeMultiplierProducer(params ops.GetFeeMultiplierProducerParams) middleware.Responder {
	if _, e := eos.StringToName(params.Producer); e != nil || len(params.Producer) != 12 {
		return &ops.GetFeeVotesMultiplierProducerBadRequest{
			Payload: &models.Error{
				Code:    400,
				Message: "invalid account",
			},
		}
	}
	if !state.ready() {
		return &ops.GetFeeVotesMultiplierProducerServiceUnavailable{
			Payload: &models.Error{
				Code:    503,
				Message: "data is not up to date, please try again later",
			},
		}
	}
	var vote float64
	for i := range state.FeeVoters {
		if string(state.FeeVoters[i].BlockProducerName) == params.Producer {
			vote = state.FeeVoters[i].FeeMultiplier
			ts := state.FeeVoters[i].LastVoteTimestamp
			if vote == 0 {
				return &ops.GetFeeVotesMultiplierProducerNotFound{
					Payload: &models.Error{
						Code:    404,
						Message: "multiplier not found",
					},
				}
			}
			return &ops.GetFeeVotesMultiplierProducerOK{Payload: &ops.GetFeeVotesMultiplierProducerOKBody{
				Multiplier: vote,
				Timestamp:  ts,
			}}
		}
	}
	return &ops.GetFeeVotesMultiplierProducerNotFound{
		Payload: &models.Error{
			Code:    404,
			Message: "producer not found",
		},
	}
}



func getProducerVotes(bp string, format string, applyMultiplier bool) middleware.Responder {
	if _, e := eos.StringToName(bp); e != nil || len(bp) != 12 {
		return &ops.GetFeeVotesFeevoteProducerBadRequest{
			Payload: &models.Error{
				Code:    400,
				Message: "invalid account",
			},
		}
	}
	if !state.ready() {
		return &ops.GetFeeVotesFeevoteProducerServiceUnavailable{
			Payload: &models.Error{
				Code:    503,
				Message: "data is not up to date, please try again later",
			},
		}
	}
	for i := range state.FeeVotes {
		if string(state.FeeVotes[i].BlockProducerName) == bp {
			var multiplier float64
			payload := &ops.GetFeeVotesProducerOK{Payload: make([]*models.Price, 0)}
			if applyMultiplier {
				for _, mult := range state.FeeVoters {
					if string(mult.BlockProducerName) == bp {
						multiplier = mult.FeeMultiplier
						break
					}
				}
				if multiplier == 0 {
					return &ops.GetFeeVotesFeevoteProducerNotFound{
						Payload: &models.Error{
							Code:    404,
							Message: "multiplier not found, unable to calculate effective fee",
						},
					}
				}
			}
			for _, v := range state.FeeVotes[i].FeeVotes {
				var price float64
				endpoint := v.EndPoint // dereference
				value := uint64(v.Value)
				if applyMultiplier {
					switch format {
					case "suf":
						// warning: overflows if we attempt to maintain precision, use big:
						value = big.NewInt(0).Div(big.NewInt(0).Mul(big.NewInt(v.Value), big.NewInt(int64(multiplier*1_000_000_000))), big.NewInt(1_000_000_000)).Uint64()
					case "float":
						price = (float64(v.Value) * multiplier) / 1_000_000_000.0
					case "usd":
						price = (float64(v.Value) * multiplier) / 1_000_000_000.0 * state.Price
					}
				}
				payload.Payload = append(payload.Payload, &models.Price{
					EndPoint: &endpoint,
					Price:    price,
					Value:    value,
				})
			}
			return payload
		}
	}
	return &ops.GetFeeVotesFeevoteProducerNotFound{
		Payload: &models.Error{
			Code:    404,
			Message: "producer not found",
		},
	}
}

func FeeVotesProducer(params ops.GetFeeVotesProducerParams) middleware.Responder {
	return getProducerVotes(params.Producer, *params.Format, *params.ApplyMultiplier)
}

func GetPrice(ops.GetPriceParams) middleware.Responder {
	if !state.ready() {
		return &ops.GetPriceServiceUnavailable{
			Payload: &models.Error{
				Code:    503,
				Message: "data is not up to date, please try again later",
			},
		}
	}
	return &ops.GetPriceOK{
		Payload: &ops.GetPriceOKBody{
			Price: state.Price,
		},
	}
}

func SuggestFee(params ops.GetSuggestFeevoteParams) middleware.Responder {
	return &ops.GetSuggestFeevoteOK{Payload: suggestFee()}
}

func SuggestMult(p ops.GetSuggestSetfeemultParams) middleware.Responder {
	if !state.ready() {
		return &ops.GetSuggestSetfeemultServiceUnavailable{
			Payload: &models.Error{
				Code:    503,
				Message: "data is not up to date, please try again later",
			},
		}
	}
	var fv float64
	for _, v := range suggestFee() {
		if *v.EndPoint == "register_fio_address" {
			fv = float64(*v.Value) / 1_000_000_000
		}
	}
	return &ops.GetSuggestSetfeemultOK{Payload: &ops.GetSuggestSetfeemultOKBody{SuggestedMultiplier: *p.Target / (state.Price * fv)}}
}
