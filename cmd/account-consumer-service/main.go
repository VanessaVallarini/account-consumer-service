package main

import (
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/entities"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/repository"
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/gocql/gocql"
	"github.com/joomcode/errorx"
)

func main() {

	ctx := context.Background()
	config := config.NewConfig()
	var err error
	/*
		kafkaClient, err := kafka.NewKafkaClient(config.Kafka)
		if err != nil {
			zap.S().Fatal(err)
		}

		Producer(10, ctx, config.Kafka, kafkaClient)

		err = consumer.NewConsumer(ctx, config.Kafka, kafkaClient)
		if err != nil {
			zap.S().Fatal(err)
		}
	*/

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

	registry := repository.NewRegistryRepository(session)
	var errorx *errorx.Error

	//createAddress
	fmt.Println("CRIANDO ENDEREÇO...")
	aCreate := entities.Address{
		Alias:       "SP",
		City:        "São Paulo",
		District:    "Sé",
		PublicPlace: "Praça da Sé",
		ZipCode:     "01001-000",
	}
	errorx = registry.AddressRepository().Create(ctx, aCreate)
	if errorx != nil {
		fmt.Println(errorx)
	}
	fmt.Println("ENDEREÇO CRIADO!")

	fmt.Println()

	//getAddressAll
	fmt.Println("PEGANDO TODOS OS ENDEREÇOS...")
	aList, errorx := registry.AddressRepository().List(ctx)
	if errorx != nil {
		fmt.Println(errorx)
	}
	for _, a := range aList {
		fmt.Printf("Id:%v. Alias:%s, City:%s, District:%s, PublicPlace:%s, ZipCode:%s \n", a.Id, a.Alias, a.City, a.District, a.PublicPlace, a.ZipCode)
	}
	fmt.Println("PEGAMOS TODOS OS ENDEREÇOS!")

	fmt.Println()

	//getAddressById
	fmt.Println("PEGANDO O ENDEREÇO POR ID...")
	reqAById := entities.AddressRequestById{
		Id: gocql.UUID.String(aList[0].Id),
	}
	retAById, errorx := registry.AddressRepository().GetById(ctx, reqAById)
	if errorx != nil {
		fmt.Println(errorx)
	}
	fmt.Printf("Id:%v. Alias:%s, City:%s, District:%s, PublicPlace:%s, ZipCode:%s \n", retAById.Id, retAById.Alias, retAById.City, retAById.District, retAById.PublicPlace, retAById.ZipCode)
	fmt.Println("PEGAMOS O ENDEREÇO POR ID!")

	fmt.Println()
	fmt.Println()

	//createPhone
	fmt.Println("CRIANDO PHONE...")
	pCreate := entities.Phone{
		CountryCode: "55",
		AreaCode:    "11",
		Number:      "964127229",
	}
	errorx = registry.PhoneRepository().Create(ctx, pCreate)
	if errorx != nil {
		fmt.Println(errorx)
	}
	fmt.Println("PHONE CRIADO!")

	fmt.Println()

	//getPhoneAll
	fmt.Println("PEGANDO TODOS OS PHONES...")
	pList, errorx := registry.PhoneRepository().List(ctx)
	if errorx != nil {
		fmt.Println(errorx)
	}
	for _, p := range pList {
		fmt.Printf("Id:%v. CountryCode:%s, AreaCode:%s, Number:%s \n", p.Id, p.CountryCode, p.AreaCode, p.Number)
	}
	fmt.Println("PEGAMOS TODOS OS PHONES!")

	fmt.Println()

	//getPhoneById
	fmt.Println("PEGANDO O PHONE POR ID...")
	reqPById := entities.PhoneRequestById{
		Id: gocql.UUID.String(pList[0].Id),
	}
	retPById, errorx := registry.PhoneRepository().GetById(ctx, reqPById)
	if errorx != nil {
		fmt.Println(errorx)
	}
	fmt.Printf("Id:%v. CountryCode:%s, AreaCode:%s, Number:%s \n", retPById.Id, retPById.CountryCode, retPById.AreaCode, retPById.Number)
	fmt.Println("PEGAMOS O PHONE POR ID!")

	fmt.Println()
	fmt.Println()

	//createUser
	fmt.Println("CRIANDO USER...")
	uCreate := entities.User{
		AddressId: retAById.Id.String(),
		PhoneId:   retPById.Id.String(),
		Name:      "Van",
		Email:     "van@email.com",
	}
	errorx = registry.UserRepository().Create(ctx, uCreate)
	if errorx != nil {
		fmt.Println(errorx)
	}
	fmt.Println("USER CRIADO!")

	fmt.Println()

	//getUserAll
	fmt.Println("PEGANDO TODOS OS USERS...")
	uList, errorx := registry.UserRepository().List(ctx)
	if errorx != nil {
		fmt.Println(errorx)
	}
	for _, u := range uList {
		fmt.Printf("AddressId:%v. PhoneId:%s, Name:%s, Email:%s \n", u.AddressId, u.PhoneId, u.Name, u.Email)
	}
	fmt.Println("PEGAMOS TODOS OS USERS!")

	fmt.Println()

	//getUserById
	fmt.Println("PEGANDO O USER POR ID...")
	reqUById := entities.UserRequestById{
		Id: gocql.UUID.String(uList[0].Id),
	}
	retUById, errorx := registry.UserRepository().GetById(ctx, reqUById)
	if errorx != nil {
		fmt.Println(errorx)
	}
	fmt.Printf("AddressId:%v. PhoneId:%s, Name:%s, Email:%s \n", retUById.AddressId, retUById.PhoneId, retUById.Name, retUById.Email)
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
