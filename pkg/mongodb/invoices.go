package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CollectionInvoices(ctx context.Context) (*mongo.Collection, error) {
	return collection(ctx, "invoices")
}

type InvoiceItem struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	PassportNumber string             `bson:"passportNumber" json:"passportNumber"`
	Name           string             `bson:"name" json:"name"`
	Country        string             `bson:"country" json:"country"`
	Service        string             `bson:"service" json:"service"`
	OrderDate      string             `bson:"orderDate" json:"orderDate"`
	ArrivalDate    string             `bson:"arrivalDate" json:"arrivalDate"`
	Port           string             `bson:"port" json:"port"`
	Cost           string             `bson:"cost" json:"cost"`

	Found   bool   `bson:"found" json:"found"`
	Comment string `bson:"comment" json:"comment"`
}

type Invoice struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Items []InvoiceItem      `bson:"items" json:"items"`
	Title string             `bson:"title" json:"title"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
