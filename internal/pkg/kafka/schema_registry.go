package kafka

import (
	"account-consumer-service/internal/pkg/utils"
	"encoding/binary"
	"errors"
	"strings"

	"github.com/hamba/avro"
	"github.com/riferrei/srclient"
)

type SchemaRegistry struct {
	*srclient.SchemaRegistryClient
}

func NewSchemaRegistry(host, user, password string) (*SchemaRegistry, error) {
	src := srclient.CreateSchemaRegistryClient(host)
	if user != "" || password != "" {
		src.SetCredentials(user, password)
	}

	return &SchemaRegistry{src}, nil
}

// ValidateSchema checks for the existence and compatibility of a schema.
// if the subject does not exist it will be created, if it is incompatible it will return an error.
func (sr *SchemaRegistry) ValidateSchema(rawSchema, subject string, schemaType string) error {
	schema, err := sr.GetLatestSchema(subject)

	if err != nil && !strings.Contains(err.Error(), "404") {
		utils.Logger.Error("kafka schema registry: %v", err)
		return err
	}

	if schema == nil {
		_, err := sr.CreateSchema(subject, rawSchema, srclient.SchemaType(schemaType))
		if err != nil {
			return err
		}
		return nil
	}

	isCompatible, err := sr.IsSchemaCompatible(subject, rawSchema, "latest", srclient.SchemaType(schemaType))
	if err != nil || !isCompatible {
		return err
	}

	if !isCompatible {
		return errors.New("schema registry invalid schema is not compatible")
	}

	return nil
}

func (sr *SchemaRegistry) GetSchema(subject string) (*srclient.Schema, error) {
	schema, err := sr.GetLatestSchema(subject)
	if err != nil {
		return nil, err
	}

	if schema == nil {
		return nil, errors.New("schema registry unexpected behavior retrieving schema, got 'nil' from registry")
	}

	return schema, nil
}

func (sr *SchemaRegistry) Decode(data []byte, value interface{}, subject string) error {
	schema, err := sr.GetSchema(subject)
	if err != nil {
		return err
	}

	schemaDecoder, err := avro.Parse(schema.Schema())
	if err != nil {
		return err
	}

	err = avro.Unmarshal(schemaDecoder, data[5:], value)
	if err != nil {
		return err
	}

	return nil
}

func (sr *SchemaRegistry) Encode(value interface{}, subject string) ([]byte, error) {
	schema, err := sr.GetSchema(subject)
	if err != nil {
		return nil, err
	}

	schemaEncoder, err := avro.Parse(schema.Schema())
	if err != nil {
		return []byte{}, err
	}

	avroNative, err := avro.Marshal(schemaEncoder, value)
	if err != nil {
		return []byte{}, err
	}

	var recordValue []byte
	schemaIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(schemaIDBytes, uint32(schema.ID()))
	recordValue = append(recordValue, byte(0))
	recordValue = append(recordValue, schemaIDBytes...)
	recordValue = append(recordValue, avroNative...)

	return recordValue, nil
}
