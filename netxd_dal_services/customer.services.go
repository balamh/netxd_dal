package netxd_dal_services

import (
	"context"
	"github.com/balamh/netxd_dal/netxd_dal_interfaces"
	"github.com/balamh/netxd_dal/netxd_dal_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerServices struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) netxd_dal_interfaces.INetxdCustomers {
	return &CustomerServices{collection, ctx}
}
func (c *CustomerServices) CreateCustomer(user *netxd_dal_models.NetxdCustomer) (*netxd_dal_models.CustomerResponse, error) {
	res, err := c.CustomerCollection.InsertOne(c.ctx, &user)
	if err != nil {
		return nil, err
	}
	var newUser *netxd_dal_models.CustomerResponse
	query := bson.M{"_id": res.InsertedID}

	err = c.CustomerCollection.FindOne(c.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
