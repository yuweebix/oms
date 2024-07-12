package kafka_test

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/kafka/pub"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/kafka/sub"
)

type SingleSuite struct {
	suite.Suite
	producer *pub.Producer
	consumer *sub.Consumer
}

// TestSingleSuite запускает все orders int-тесты
func TestSingleSuite(t *testing.T) {
	suite.Run(t, new(SingleSuite))
}

func (s *SingleSuite) SetupSuite() {
	brokers, err := getBrokers()

	if err != nil {
		s.FailNowf("Could not get brokers", err.Error())
	}

	s.producer, err = pub.NewProducer(brokers, topic)
	if err != nil {
		s.FailNowf("Error creating producer", err.Error())
	}

	s.consumer, err = sub.NewConsumer(brokers)
	if err != nil {
		s.FailNowf("Error creating consumer", err.Error())
	}

	ch = make(chan string, 100)
	if err := s.consumer.Start(topic, ch); err != nil {
		s.FailNowf("Error starting consumer", err.Error())
	}
}

func (s *SingleSuite) TearDownSuite() {
	if err := s.producer.Close(); err != nil {
		s.T().Log(err)
	}
	if err := s.consumer.Stop(); err != nil {
		s.T().Log(err)
	}
	if err := s.consumer.Close(); err != nil {
		s.T().Log(err)
	}
}

func (s *SingleSuite) TestString() {
	test := struct {
		expected string
		actual   string
	}{
		"test",
		"",
	}

	err := s.producer.Send(test.expected)
	test.actual = <-ch

	s.NoError(err)
	s.Contains(test.actual, test.expected)
}

func (s *SingleSuite) TestStruct() {
	type numStruct struct {
		Num int `json:"num"`
	}

	test := struct {
		expected numStruct
		actual   numStruct
	}{
		numStruct{1},
		numStruct{},
	}

	err := s.producer.Send(test.expected)
	data := <-ch

	json.Unmarshal([]byte(data), &test.actual)

	s.NoError(err)
	s.Equal(test.actual, test.expected)
}

func (s *SingleSuite) TestConcurrency() {
	type numStruct struct {
		Num int `json:"num"`
	}

	tests := struct {
		expected []numStruct
		actual   []numStruct
	}{
		[]numStruct{
			{1},
			{2},
			{3},
			{4},
			{5},
		},
		make([]numStruct, 5),
	}
	errChan := make(chan error, len(tests.expected))
	defer close(errChan)

	send := func(i int, ns numStruct) {
		err := s.producer.Send(ns)

		// т.к. тестим конкурентно, надо бы записывать ошибки в канал
		if err != nil {
			errChan <- err
			return
		}

		data := <-ch
		errChan <- json.Unmarshal([]byte(data), &tests.actual[i])
	}

	// отправляем сами сообщения
	for i, msg := range tests.expected {
		go send(i, msg)
	}

	// тут мы чекаем, что нет ошибок
	for range tests.expected {
		err := <-errChan
		s.NoError(err)
	}

	// скорее всего придут не по порядку
	sort.Slice(tests.actual, func(i, j int) bool {
		return tests.actual[i].Num < tests.actual[j].Num
	})

	for i := range tests.expected {
		s.Equal(tests.expected[i], tests.actual[i])
	}
}
