package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECTED  = "Successfully connected to database: %s"
	caFilePath = ".certs/rds-combined-ca-bundle.pem"
)

type MongoDB struct {
	Database *mongo.Database
	Client   *mongo.Client
	Logger   *log.Logger
}

type MongoConfig struct {
	MongoUri      string
	MongoTls      *tls.Config
	MongoDatabase string
}

func connect(config *MongoConfig, logger *logrus.Logger) (*mongo.Database, *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client

	connectOnce.Do(func() {
		db, session = connectToMongo(config, logger)
	})

	return db, session
}

func connectToMongo(config *MongoConfig, logger *logrus.Logger) (a *mongo.Database, b *mongo.Client) {

	var err error
	var session *mongo.Client

	if os.Getenv("GO_ENV") == "AWS" {
		session, err = mongo.NewClient(options.Client().ApplyURI(config.MongoUri).SetTLSConfig(config.MongoTls))
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		session, err = mongo.NewClient(options.Client().ApplyURI(config.MongoUri))
		if err != nil {
			logger.Fatal(err)
		}
	}

	session.Connect(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}

	var DB = session.Database(config.MongoDatabase)
	logger.Infof(CONNECTED, config.MongoDatabase)

	return DB, session
}

func setMongoConfig() *MongoConfig {
	var err error
	config := &MongoConfig{}

	if os.Getenv("GO_ENV") == "AWS" {
		config.MongoUri = fmt.Sprintf(os.Getenv("AWS_MONGO_HOST"), os.Getenv("AWS_MONGO_USER"), os.Getenv("AWS_MONGO_PASSWORD"))
		config.MongoTls, err = getCustomTLSConfig(caFilePath)
		if err != nil {
			log.Fatalf("DAMNIT: %s\n\n", err)
		}
	} else {
		config.MongoUri = os.Getenv("LOCAL_MONGO_HOST")
	}

	config.MongoDatabase = os.Getenv("MONGO_DATABASE")
	return config
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("failed parsing pem file")
	}

	return tlsConfig, nil
}

func GetConnection() (*MongoDB, error) {
	var err error = nil
	var log = logrus.New()

	config := setMongoConfig()

	db, session := connect(config, log)

	if db == nil || session == nil {
		log.Fatalf("Failed to connect to database: %v", config.MongoDatabase)
		err = fmt.Errorf("Failed to connect to database: %v", config.MongoDatabase)
	}

	return &MongoDB{
		Database: db,
		Client:   session,
		Logger:   log,
	}, err
}
