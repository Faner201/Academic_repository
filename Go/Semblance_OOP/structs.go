package main

import (
	"time"
)

type Searcher interface {
	Contains(text string) bool
}

type Timestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTimestamp() Timestamp {
	return Timestamp{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Document struct {
	ID     int
	UserID int
	Name   string
	Value  string
	Timestamp
}

func NewDocument(userID int, name, value string) *Document {
	return &Document{
		Name:      name,
		UserID:    userID,
		Value:     value,
		Timestamp: NewTimestamp(),
	}
}

type Message struct {
	ID         int
	SenderID   int
	ReceiverID int
	Text       string
	Timestamp
}

func NewMessage(senderID, receiverID int, text string) *Message {
	return &Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Text:       text,
		Timestamp:  NewTimestamp(),
	}
}

func SearchText(text string, values []Searcher) []Searcher {
	result := make([]Searcher, 0)
	for _, v := range values {
		if v.Contains(text) {
			result = append(result, v)
		}
	}
	return result
}
