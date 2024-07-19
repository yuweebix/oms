package middleware

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Logging представляет собой логирование вызванных методов и ошибок в консоль
func Logging(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	raw, err := protojson.Marshal((req).(proto.Message))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	resp, err = handler(ctx, req)
	log.Println(info.FullMethod, string(raw))
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}

	return
}
