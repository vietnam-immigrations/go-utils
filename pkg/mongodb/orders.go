package mongodb

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CollectionOrders(ctx context.Context, log *logrus.Entry, stage string) (*mongo.Collection, error) {
	return collection(ctx, log, stage, "orders")
}

type Billing struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Email     string `bson:"email" json:"email"`
	Email2    string `bson:"email2" json:"email2"`
	Phone     string `bson:"phone" json:"phone"`
}

type Trip struct {
	ArrivalDate      string `bson:"arrivalDate" json:"arrivalDate"`
	Checkpoint       string `bson:"checkpoint" json:"checkpoint"`
	ProcessingTime   string `bson:"processingTime" json:"processingTime"`
	FastTrack        string `bson:"fastTrack" json:"fastTrack"`
	CarPickup        bool   `bson:"carPickup" json:"carPickup"`
	Flight           string `bson:"flight" json:"flight"`
	CarPickupAddress string `bson:"carPickupAddress" json:"carPickupAddress"`
}

type Applicant struct {
	FirstName      string `bson:"firstName" json:"firstName"`
	LastName       string `bson:"lastName" json:"lastName"`
	DateOfBirth    string `bson:"dateOfBirth" json:"dateOfBirth"`
	Sex            string `bson:"sex" json:"sex"`
	Nationality    string `bson:"nationality" json:"nationality"`
	PassportNumber string `bson:"passportNumber" json:"passportNumber"`
	PassportExpiry string `bson:"passportExpiry" json:"passportExpiry"`
}

type Order struct {
	ID                 primitive.ObjectID `bson:"_id" json:"id"`
	OrderID            int                `bson:"orderId" json:"orderId"`
	Total              string             `bson:"total" json:"total"`
	OrderKey           string             `bson:"orderKey" json:"orderKey"`
	Billing            Billing            `bson:"billing" json:"billing"`
	PaymentMethodTitle string             `bson:"paymentMethodTitle" json:"paymentMethodTitle"`
	Number             string             `bson:"number" json:"number"`
	Trip               Trip               `bson:"trip" json:"trip"`
	Applicants         []Applicant        `bson:"applicants" json:"applicants"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
