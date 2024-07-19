package api

import (
	"context"

	returns "gitlab.ozon.dev/yuweebix/homework-1/gen/returns/v1/proto"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AcceptReturn реализует метод /v1/returns/accept
func (api *API) AcceptReturn(ctx context.Context, req *returns.AcceptReturnRequest) (resp *returns.AcceptReturnResponse, err error) {
	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	err = api.service.AcceptReturn(ctx, &models.Order{
		ID:   req.GetOrderId(),
		User: &models.User{ID: req.GetUserId()},
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &returns.AcceptReturnResponse{}, nil
}

// ListReturn реализует метод /v1/returns/accept
func (api *API) ListReturns(ctx context.Context, req *returns.ListReturnsRequest) (resp *returns.ListReturnsResponse, err error) {
	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	list, err := api.service.ListReturns(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// переведём в вид респоса
	listResp := fromModelsOrderForListReturns(list)
	return &returns.ListReturnsResponse{Orders: listResp}, nil
}

// хелпер-функции для преобразования

func fromModelsOrderForListReturns(list []*models.Order) (listResp []*returns.ListReturnsResponse_Order) {
	listResp = make([]*returns.ListReturnsResponse_Order, 0, len(list))
	for _, m := range list {
		listResp = append(listResp, &returns.ListReturnsResponse_Order{
			OrderId:   m.ID,
			UserId:    m.User.ID,
			Expiry:    timestamppb.New(m.Expiry),
			ReturnBy:  timestamppb.New(m.ReturnBy),
			Status:    returns.Status(returns.Status_value[string(m.Status)]),
			Hash:      m.Hash,
			CreatedAt: timestamppb.New(m.CreatedAt),
			Cost:      utils.ConvertFromMicrocurrency(m.Cost),
			Weight:    m.Weight,
			Packaging: returns.PackagingType(returns.PackagingType_value[string(m.Packaging)]),
		})
	}

	return listResp
}
