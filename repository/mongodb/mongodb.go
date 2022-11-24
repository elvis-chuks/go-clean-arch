package mongodb

import (
	"github.com/Kamva/mgm/v2"
	"github.com/elvis-chuks/go-clean-arch/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func New(l *zap.Logger) *domain.MongoRepository {
	var dbName, connectionString string

	if viper.Get("DB_NAME") != nil {
		dbName = viper.GetString("DB_NAME")
	}

	if viper.Get("DB_URL") != nil {
		connectionString = viper.GetString("DB_URL")
	}

	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(connectionString))

	if err != nil {
		l.Error(err.Error(), zap.Error(err))
	}

	return &domain.MongoRepository{
		UserRepo: NewMongoUserRepository(l),
	}
}
