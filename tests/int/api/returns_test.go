package api_test

import (
	"context"

	returns "gitlab.ozon.dev/yuweebix/homework-1/gen/returns/v1/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// validation tests
func (s *APISuite) TestAcceptReturn_Validation() {
	ctx := context.Background()

	tests := []struct {
		name string
		req  *returns.AcceptReturnRequest
	}{
		{
			"zero order id",
			&returns.AcceptReturnRequest{
				OrderId: 0,
				UserId:  1,
			},
		},
		{
			"zero user id",
			&returns.AcceptReturnRequest{
				OrderId: 1,
				UserId:  0,
			},
		},
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := returns.NewReturnsClient(conn)

	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, err := client.AcceptReturn(ctx, tt.req)
			s.ErrorContains(err, status.Error(codes.InvalidArgument, "").Error())
		})
	}
}

// у ListReturns ничего не провалидировать, т.к. unint64 всегда больше или равен нулю

// successful tests

func (s *APISuite) TestAcceptReturn_Success() {
	ctx := context.Background()
	req := &returns.AcceptReturnRequest{
		OrderId: 1,
		UserId:  1,
	}

	conn := s.GetClientConn()
	defer conn.Close()

	client := returns.NewReturnsClient(conn)
	_, err := client.AcceptReturn(ctx, req)

	s.NoError(err)
}

func (s *APISuite) TestListReturns_Success() {
	ctx := context.Background()
	req := &returns.ListReturnsRequest{}

	conn := s.GetClientConn()
	defer conn.Close()

	client := returns.NewReturnsClient(conn)
	_, err := client.ListReturns(ctx, req)

	s.NoError(err)
}
