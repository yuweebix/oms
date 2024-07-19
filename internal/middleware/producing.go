package middleware

import (
	"context"
	"fmt"
	"os"
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/kafka/pub"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type producer struct {
	*pub.Producer
}

type ProducerKeyType string

const Producer ProducerKeyType = "producer"

// SetProducerContext добавляет продьюсера в контекст
func SetProducerContext(producer *pub.Producer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = context.WithValue(ctx, Producer, producer)
		return handler(ctx, req)
	}
}

// Producing представляет собой логирование вызванных методов и ошибок в брокер сообщений (кафку)
func Producing(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	p := ctx.Value(Producer).(*pub.Producer)
	producer := &producer{p}

	raw, err := protojson.Marshal((req).(proto.Message))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	msg := &models.Message{
		CreatedAt:  time.Now().UTC(),
		MethodName: info.FullMethod,
		RawRequest: string(raw),
	}

	resp, err = handler(ctx, req)
	if err != nil {
		producer.sendWithError(msg, err)
		return nil, err
	}

	producer.send(msg)
	return
}

// send функция-обертка, что залоггирует сообщение в брокер сообшений (кафку)
func (p *producer) send(msg *models.Message) {
	err := p.Send(msg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// sendWithError функция-обертка, что залоггирует сообщение с ошибкой
func (p *producer) sendWithError(msg *models.Message, err error) {
	msgWithErr := models.MessageWithError{
		Message: msg,
		Error:   err.Error(),
	}

	err = p.Send(msgWithErr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
