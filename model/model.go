package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	SORT_PARAM_EVENTTYPE = "eventType"
	SORT_PARAM_EVENTNAME = "eventName"
	SORT_PARAM_EVENTDESC = "eventDesc"
	SORT_PARAM_DATETIME = "dateTime"
	SORT_PARAM_DISTANCE = "distance"

	SORT_ORDER_ASC = "asc"
	SORT_ORDER_DESC = "desc"
)

type RequestGeo struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latiude"`
	SortParam string `json:"sort"`
	SortOrder string `json:"order"`
	Limit int `json:"limit"`
	Radius int `json:"radius"`
}

type ResponseGeo struct {
	Events []EventLocation `json:"events,omitempty"`
	Summary ResponseSummary `json:"summary"`
}

type ResponseSummary struct {
	Count int `json:"count"`
	Duration string `json:"duration"`
	Date time.Time `json:"date"`
	Params RequestGeo `json:"parameters"`
}

type EventLocation struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Location GeoJson  `bson:"location" json:"location"`
	EventType int  `bson:"eventType" json:"eventType,omitempty"`
	EventName string  `bson:"eventName" json:"eventName,omitempty"`
	EventDesc string  `bson:"eventDesc" json:"eventDesc,omitempty"`
	DateTime  time.Time `bson:"dateTime" json:"dateTime"`
	Distance  float64 `bson:"distance" json:"distance"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}