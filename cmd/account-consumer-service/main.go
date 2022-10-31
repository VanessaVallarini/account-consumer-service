package main

import (
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/entities"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/kafka/consumer"
	"account-consumer-service/internal/pkg/repository"
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func main() {

	ctx := context.Background()

	config := config.NewConfig()

	var err error

	kafkaClient, err := kafka.NewKafkaClient(config.Kafka)
	if err != nil {
		zap.S().Fatal(err)
	}

	Producer(10, ctx, config.Kafka, kafkaClient)

	//scylla
	cluster := gocql.NewCluster(config.DatabaseConnStr.DatabaseHost)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: config.DatabaseConnStr.DatabaseUser,
		Password: config.DatabaseConnStr.DatabasePassword,
	}
	cluster.Keyspace = config.DatabaseConnStr.DatabaseKeyspace
	cluster.ConnectTimeout = cluster.ConnectTimeout * 5
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
	}

	//kafka

	err = consumer.NewConsumer(ctx, config.Kafka, kafkaClient)
	if err != nil {
		zap.S().Fatal(err)
	}

	registry := repository.NewRegistryRepository(session)

	//createAddress
	fmt.Println("CRIANDO ENDEREÇO...")
	a := entities.Address{
		Alias:       "SP",
		City:        "São Paulo",
		District:    "Sé",
		PublicPlace: "Praça da Sé",
		ZipCode:     "01001-000",
	}
	errCreateA := registry.AddressRepository().Create(ctx, a)
	if errCreateA != nil {
		fmt.Println(errCreateA)
	}
	fmt.Println("ENDEREÇO CRIADO!")

	//getAddressAll
	fmt.Println("PEGANDO TODOS OS ENDEREÇOS...")
	aList, errListA := registry.AddressRepository().List(ctx)
	if errListA != nil {
		fmt.Println(errListA)
	}
	for _, a := range aList {
		fmt.Println(a.Id)
		fmt.Println(a.Alias)
		fmt.Println(a.City)
		fmt.Println(a.District)
		fmt.Println(a.PublicPlace)
		fmt.Println(a.ZipCode)
		fmt.Printf("Id:%v. Alias:%s, City:%s, District:%s, PublicPlace:%s, ZipCode:%s", a.Id, a.Alias, a.City, a.District, a.PublicPlace, a.ZipCode)
	}
	fmt.Println("PEGAMOS TODOS OS ENDEREÇOS!")

	//getAddressById
	fmt.Println("PEGANDO O ENDEREÇO POR ID...")
	reqA := entities.Address{
		Id: aList[0].Id,
	}
	retA, errGetAById := registry.AddressRepository().GetById(ctx, reqA)
	if errGetAById != nil {
		fmt.Println(errGetAById)
	}
	fmt.Println(retA.Id)
	fmt.Println("PEGAMOS O ENDEREÇO POR ID!")

	//createPhone
	fmt.Println("CRIANDO PHONE...")
	p := entities.Phone{
		CountryCode: "55",
		AreaCode:    "11",
		Number:      "964127229",
	}
	errCreateP := registry.PhoneRepository().Create(ctx, p)
	if errCreateP != nil {
		fmt.Println(errCreateP)
	}
	fmt.Println("PHONE CRIADO!")

	//getPhoneAll
	fmt.Println("PEGANDO TODOS OS PHONES...")
	pList, errListP := registry.PhoneRepository().List(ctx)
	if errListP != nil {
		fmt.Println(errListP)
	}
	for _, p := range pList {
		fmt.Println(p.Id)
		fmt.Println(p.CountryCode)
		fmt.Println(p.AreaCode)
		fmt.Println(p.Number)
		fmt.Printf("Id:%v. CountryCode:%s, AreaCode:%s, Number:%s", p.Id, p.CountryCode, p.AreaCode, p.Number)
	}
	fmt.Println("PEGAMOS TODOS OS PHONES!")

	//getPhoneById
	fmt.Println("PEGANDO O PHONE POR ID...")
	reqP := entities.Phone{
		Id: pList[0].Id,
	}
	retP, errGetPById := registry.PhoneRepository().GetById(ctx, reqP)
	if errGetPById != nil {
		fmt.Println(errGetPById)
	}
	fmt.Println(retP.Id)
	fmt.Println("PEGAMOS O PHONE POR ID!")

	//createUser
	fmt.Println("CRIANDO USER...")
	u := entities.User{
		AddressId: retA.Id.String(),
		PhoneId:   retP.Id.String(),
		Name:      "Van",
		Email:     "van@email.com",
	}
	errCreateU := registry.UserRepository().Create(ctx, u)
	if errCreateU != nil {
		fmt.Println(errCreateU)
	}
	fmt.Println("USER CRIADO!")

	//getUserAll
	fmt.Println("PEGANDO TODOS OS USERS...")
	uList, errListU := registry.UserRepository().List(ctx)
	if errListU != nil {
		fmt.Println(errListU)
	}
	for _, u := range uList {
		fmt.Println(u.AddressId)
		fmt.Println(u.PhoneId)
		fmt.Println(u.Name)
		fmt.Println(u.Email)
		fmt.Printf("AddressId:%v. PhoneId:%s, Name:%s, Email:%s", u.AddressId, u.PhoneId, u.Name, u.Email)
	}
	fmt.Println("PEGAMOS TODOS OS USERS!")

	//getUserById
	reqUser := entities.User{
		Id: uList[0].Id,
	}
	retUser, errGetUById := registry.UserRepository().GetById(ctx, reqUser)
	if errGetUById != nil {
		fmt.Println(errGetUById)
	}
	fmt.Println(retUser.Id)
	fmt.Println("PEGAMOS O USER POR ID!")

}

func Producer(limit int, ctx context.Context, kafkaConfig *config.KafkaConfig, kafkaClient *kafka.KafkaClient) {

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.StringEncoder("testing 123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}
