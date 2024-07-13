package api

import (
	"context"

	orders "gitlab.ozon.dev/yuweebix/homework-1/gen/orders/v1/proto"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"gitlab.ozon.dev/yuweebix/homework-1/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AcceptOrder реализует метод /v1/orders/accept
func (api *API) AcceptOrder(ctx context.Context, req *orders.AcceptRequest) (resp *orders.AcceptResponse, err error) {
	// составляем сообщения, что пойдет в брокер
	msg, err := getMessage(ctx, req.ProtoReflect())
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	err = api.domain.AcceptOrder(ctx, &models.Order{
		ID:        req.GetOrderId(),
		User:      &models.User{ID: req.GetUserId()},
		Expiry:    req.GetExpiry().AsTime(),
		Cost:      utils.ConvertToMicrocurrency(req.GetCost()),
		Weight:    req.GetWeight(),
		Packaging: models.PackagingType(req.GetPackaging()),
	})
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// отправляем сообщение в брокер
	err = api.producer.Send(msg)
	if err != nil {
		return nil, status.Error((codes.Internal), err.Error())
	}

	return &orders.AcceptResponse{}, nil
}

// DeliverOrders реализует метод /v1/orders/deliver
func (api *API) DeliverOrders(ctx context.Context, req *orders.DeliverRequest) (resp *orders.DeliverResponse, err error) {
	// составляем сообщения, что пойдет в брокер
	msg, err := getMessage(ctx, req.ProtoReflect())
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	err = api.domain.DeliverOrders(ctx, req.GetOrderIds())
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// отправляем сообщение в брокер
	err = api.producer.Send(msg)
	if err != nil {
		return nil, status.Error((codes.Internal), err.Error())
	}

	return &orders.DeliverResponse{}, nil
}

// ListOrders реализует метод /v1/orders/list
func (api *API) ListOrders(ctx context.Context, req *orders.ListRequest) (resp *orders.ListResponse, err error) {
	// составляем сообщения, что пойдет в брокер
	msg, err := getMessage(ctx, req.ProtoReflect())
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	list, err := api.domain.ListOrders(
		ctx,
		req.GetUserId(),
		req.GetLimit(),
		req.GetOffset(),
		req.GetIsStored(),
	)
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// переведём в вид респоса
	listResp := make([]*orders.ListResponse_Order, 0, len(list))
	for _, m := range list {
		listResp = append(listResp, &orders.ListResponse_Order{
			OrderId:   m.ID,
			UserId:    m.User.ID,
			Expiry:    timestamppb.New(m.Expiry),
			ReturnBy:  timestamppb.New(m.ReturnBy),
			Status:    string(m.Status),
			Hash:      m.Hash,
			CreatedAt: timestamppb.New(m.CreatedAt),
			Cost:      utils.ConvertFromMicrocurrency(m.Cost),
			Weight:    m.Weight,
			Packaging: string(m.Packaging),
		})
	}

	// отправляем сообщение в брокер
	err = api.producer.Send(msg)
	if err != nil {
		return nil, status.Error((codes.Internal), err.Error())
	}

	return &orders.ListResponse{
		Orders: listResp,
	}, nil
}

// ReturnOrder реализует метод /v1/orders/return
func (api *API) ReturnOrder(ctx context.Context, req *orders.ReturnRequest) (resp *orders.ReturnResponse, err error) {
	// составляем сообщения, что пойдет в брокер
	msg, err := getMessage(ctx, req.ProtoReflect())
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// валидация заданного в прото контракте
	err = req.ValidateAll()
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// проход в БЛ
	err = api.domain.ReturnOrder(ctx, &models.Order{ID: req.GetOrderId()})
	if err != nil {
		if err := api.producer.Send(models.MessageWithError{Message: msg, Error: err.Error()}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// отправляем сообщение в брокер
	err = api.producer.Send(msg)
	if err != nil {
		return nil, status.Error((codes.Internal), err.Error())
	}

	return &orders.ReturnResponse{}, nil
}
