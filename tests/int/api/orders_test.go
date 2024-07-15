package api_test

import (
	"context"
	"time"

	orders "gitlab.ozon.dev/yuweebix/homework-1/gen/orders/v1/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// validation tests
func (s *APISuite) TestAcceptOrder_Validation() {
	ctx := context.Background()

	tests := []struct {
		name string
		req  *orders.AcceptOrderRequest
	}{
		{
			"zero order id",
			&orders.AcceptOrderRequest{
				OrderId:   0,
				UserId:    1,
				Expiry:    timestamppb.New(time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC)),
				Cost:      1,
				Weight:    1,
				Packaging: orders.PackagingType_bag,
			},
		},
		{
			"zero user id",
			&orders.AcceptOrderRequest{
				OrderId:   1,
				UserId:    0,
				Expiry:    timestamppb.New(time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC)),
				Cost:      1,
				Weight:    1,
				Packaging: orders.PackagingType_bag,
			},
		},
		{
			"zero cost",
			&orders.AcceptOrderRequest{
				OrderId:   1,
				UserId:    1,
				Expiry:    timestamppb.New(time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC)),
				Cost:      0,
				Weight:    1,
				Packaging: orders.PackagingType_bag,
			},
		},
		{
			"zero weight",
			&orders.AcceptOrderRequest{
				OrderId:   1,
				UserId:    1,
				Expiry:    timestamppb.New(time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC)),
				Cost:      1,
				Weight:    0,
				Packaging: orders.PackagingType_bag,
			},
		},
		{
			"wrong packaging type",
			&orders.AcceptOrderRequest{
				OrderId:   1,
				UserId:    1,
				Expiry:    timestamppb.New(time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC)),
				Cost:      1,
				Weight:    1,
				Packaging: orders.PackagingType_unknown_packaging,
			},
		},
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := orders.NewOrdersClient(conn)

	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, err := client.AcceptOrder(ctx, tt.req)
			s.ErrorContains(err, status.Error(codes.InvalidArgument, "").Error())
		})
	}
}

func (s *APISuite) TestDeliverOrders_Validation() {
	ctx := context.Background()

	tests := []struct {
		name string
		req  *orders.DeliverOrdersRequest
	}{
		{
			"empty list",
			&orders.DeliverOrdersRequest{
				OrderIds: []uint64{},
			},
		},
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := orders.NewOrdersClient(conn)

	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, err := client.DeliverOrders(ctx, tt.req)
			s.ErrorContains(err, status.Error(codes.InvalidArgument, "").Error())
		})
	}
}

func (s *APISuite) TestListOrders_Validation() {
	ctx := context.Background()

	tests := []struct {
		name string
		req  *orders.ListOrdersRequest
	}{
		{
			"zero user id",
			&orders.ListOrdersRequest{
				UserId: 0,
			},
		},
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := orders.NewOrdersClient(conn)

	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, err := client.ListOrders(ctx, tt.req)
			s.ErrorContains(err, status.Error(codes.InvalidArgument, "").Error())
		})
	}
}

func (s *APISuite) TestReturnOrder_Validation() {
	ctx := context.Background()

	tests := []struct {
		name string
		req  *orders.ReturnOrderRequest
	}{
		{
			"zero order id",
			&orders.ReturnOrderRequest{
				OrderId: 0,
			},
		},
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := orders.NewOrdersClient(conn)

	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, err := client.ReturnOrder(ctx, tt.req)
			s.ErrorContains(err, status.Error(codes.InvalidArgument, "").Error())
		})
	}
}

// successful tests

func (s *APISuite) TestAcceptOrder_Success() {
	ctx := context.Background()
	req := &orders.AcceptOrderRequest{
		OrderId:   1,
		UserId:    1,
		Expiry:    timestamppb.New(time.Date(10000, 0, 0, 0, 0, 0, 0, time.UTC)),
		Cost:      1,
		Weight:    1,
		Packaging: orders.PackagingType_bag,
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := orders.NewOrdersClient(conn)
	_, err := client.AcceptOrder(ctx, req)

	s.NoError(err)
}

func (s *APISuite) TestDeliverOrders_Success() {
	ctx := context.Background()
	req := &orders.DeliverOrdersRequest{
		OrderIds: []uint64{1, 2},
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := orders.NewOrdersClient(conn)
	_, err := client.DeliverOrders(ctx, req)

	s.NoError(err)
}

func (s *APISuite) TestListOrders_Success() {
	ctx := context.Background()
	req := &orders.ListOrdersRequest{
		UserId: 1,
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := orders.NewOrdersClient(conn)
	_, err := client.ListOrders(ctx, req)

	s.NoError(err)
}

func (s *APISuite) TestReturnOrder_Success() {
	ctx := context.Background()
	req := &orders.ReturnOrderRequest{
		OrderId: 1,
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := orders.NewOrdersClient(conn)
	_, err := client.ReturnOrder(ctx, req)

	s.NoError(err)
}
