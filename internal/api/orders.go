package api

import (
	"context"

	orders "gitlab.ozon.dev/yuweebix/homework-1/gen/orders/v1/proto"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/metrics"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AcceptOrder реализует метод /v1/orders/accept
func (api *API) AcceptOrder(ctx context.Context, req *orders.AcceptOrderRequest) (resp *orders.AcceptOrderResponse, err error) {
	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	err = api.service.AcceptOrder(ctx, toModelsOrderForAcceptOrder(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &orders.AcceptOrderResponse{}, nil
}

// DeliverOrders реализует метод /v1/orders/deliver
func (api *API) DeliverOrders(ctx context.Context, req *orders.DeliverOrdersRequest) (resp *orders.DeliverOrdersResponse, err error) {
	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	err = api.service.DeliverOrders(ctx, req.GetOrderIds())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	metrics.IncDeliveriesTotal()
	return &orders.DeliverOrdersResponse{}, nil
}

// ListOrders реализует метод /v1/orders/list
func (api *API) ListOrders(ctx context.Context, req *orders.ListOrdersRequest) (resp *orders.ListOrdersResponse, err error) {
	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	list, err := api.service.ListOrders(ctx, req.GetUserId(), req.GetLimit(), req.GetOffset(), req.GetIsStored())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// переведём в вид респонса
	listResp := fromModelsOrderForListOrders(list)
	return &orders.ListOrdersResponse{
		Orders: listResp,
	}, nil
}

// ReturnOrder реализует метод /v1/orders/return
func (api *API) ReturnOrder(ctx context.Context, req *orders.ReturnOrderRequest) (resp *orders.ReturnOrderResponse, err error) {
	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	err = api.service.ReturnOrder(ctx, &models.Order{ID: req.GetOrderId()})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &orders.ReturnOrderResponse{}, nil
}

// хелпер-функции для преобразования

func toModelsOrderForAcceptOrder(req *orders.AcceptOrderRequest) *models.Order {
	return &models.Order{
		ID:        req.GetOrderId(),
		User:      &models.User{ID: req.GetUserId()},
		Expiry:    req.GetExpiry().AsTime(),
		Cost:      utils.ConvertToMicrocurrency(req.GetCost()),
		Weight:    req.GetWeight(),
		Packaging: models.PackagingType(req.GetPackaging().String()),
	}
}

func fromModelsOrderForListOrders(list []*models.Order) (listResp []*orders.ListOrdersResponse_Order) {
	listResp = make([]*orders.ListOrdersResponse_Order, 0, len(list))
	for _, m := range list {
		listResp = append(listResp, &orders.ListOrdersResponse_Order{
			OrderId:   m.ID,
			UserId:    m.User.ID,
			Expiry:    timestamppb.New(m.Expiry),
			ReturnBy:  timestamppb.New(m.ReturnBy),
			Status:    orders.Status(orders.Status_value[string(m.Status)]),
			Hash:      m.Hash,
			CreatedAt: timestamppb.New(m.CreatedAt),
			Cost:      utils.ConvertFromMicrocurrency(m.Cost),
			Weight:    m.Weight,
			Packaging: orders.PackagingType(orders.PackagingType_value[string(m.Packaging)]),
		})
	}

	return listResp
}
