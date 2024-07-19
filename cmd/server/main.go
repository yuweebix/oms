package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	orders "gitlab.ozon.dev/yuweebix/homework-1/gen/orders/v1/proto"
	returns "gitlab.ozon.dev/yuweebix/homework-1/gen/returns/v1/proto"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/api"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cache"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/domain"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/kafka/pub"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/kafka/sub/group"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/middleware"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	topic   = "api"
	groupID = "apiID"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:32269", "gRPC server endpoint")
)

func main() {
	// читаем данные из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatalln("Error reading DATABASE_URL from .env file")
	}

	brokersStr := os.Getenv("BROKERS")
	if brokersStr == "" {
		log.Fatalln("Error reading BROKERS from .env file")
	}
	brokers := strings.Split(brokersStr, ",")

	redisAddr := os.Getenv("REDIS_ADDR")
	if brokersStr == "" {
		log.Fatalln("Error reading REDIS_ADDR from .env file")
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		log.Fatalln("Error reading GRPC_PORT from .env file")
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatalln("Error reading HTTP_PORT from .env file")
	}

	outputMode := os.Getenv("OUTPUT_MODE")
	if outputMode == "" {
		log.Fatalln("Error reading OUTPUT_MODE from .env file")
	}

	// вг - для горутин + контекст
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	// бд
	repository, err := repository.NewRepository(ctx, connString)
	if err != nil {
		log.Fatalln(err)
	}
	defer repository.Close()

	// кэш
	cache, err := cache.NewCache(redisAddr, "", 0)
	if err != nil {
		log.Fatalln(err)
	}

	// сервис
	domain := domain.NewDomain(repository, cache)

	// опции для интерцептора
	opts := make([]grpc.ServerOption, 0)

	// логгинг в кафку
	notificationChan := make(chan string, 100)
	producer := &pub.Producer{}
	cg := &group.Group{}
	if outputMode == "kafka" {
		// kafka продьюсер
		producer, err = pub.NewProducer(brokers, topic)
		if err != nil {
			log.Fatalln(err)
		}

		// kafka группа консьюмеров
		cg, err = group.NewConsumerGroup(brokers, []string{topic}, groupID, notificationChan)

		// начинаем работу
		if err != nil {
			log.Fatalln(err)
		}
		if err := cg.Start(ctx, []string{topic}); err != nil {
			log.Fatalln(err)
		}

		// горутина для обработки уведомлений
		wg.Add(1)
		go func() {
			defer wg.Done()
			for notification := range notificationChan {
				fmt.Println(notification)
			}
		}()

		// сначала засеттим, а потом добавим миддлик
		opts = append(opts, grpc.ChainUnaryInterceptor(middleware.SetProducerContext(producer), middleware.Producing))
	}

	// логгинг в консоль
	if outputMode == "console" {
		opts = append(opts, grpc.ChainUnaryInterceptor(middleware.Logging))
	}

	// в api имплеменитораны методы и orders контракта, и returns контракта, поэтому можно использовать её одну
	// всё идёт на один сервер
	api := api.NewAPI(domain)
	grpcServer := grpc.NewServer(opts...)
	orders.RegisterOrdersServer(grpcServer, api)
	returns.RegisterReturnsServer(grpcServer, api)

	// для постмана полезно
	reflection.Register(grpcServer)

	// запуск grpc сервера
	go func() {
		// слушаем
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("grpc server listening on", grpcPort)

		// сёрвим
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()

	// http gateway
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

		// orders
		err := orders.RegisterOrdersHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
		if err != nil {
			log.Fatalln(err)
		}

		// returns
		err = returns.RegisterReturnsHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
		if err != nil {
			log.Fatalln(err)
		}

		// слушаем и сёрвим
		log.Println("http gateway listening on", httpPort)
		if err := http.ListenAndServe(httpPort, mux); err != nil {
			log.Fatalln(err)
		}
	}()

	// не выходим, пока не придёт сигнал
	for {
		select {
		case <-ctx.Done():
			grpcServer.GracefulStop()
			if outputMode == "kafka" {
				if err := producer.Close(); err != nil {
					log.Println(err)
				}
				if err := cg.Stop(); err != nil {
					log.Println(err)
				}
			}
			close(notificationChan)

			wg.Wait()
			return
		case <-sigs:
			cancel()
		}
	}
}
