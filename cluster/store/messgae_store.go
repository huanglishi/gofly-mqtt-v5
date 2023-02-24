package store

import (
	"context"

	"github.com/huanglishi/gofly-mqttv5/corev5/messagev5"
)

type MessageStore interface {
	BaseStore

	StoreWillMessage(ctx context.Context, clientId string, message *messagev5.PublishMessage) error
	ClearWillMessage(ctx context.Context, clientId string) error
	GetWillMessage(ctx context.Context, clientId string) (*messagev5.PublishMessage, error)

	StoreRetainMessage(ctx context.Context, topic string, message *messagev5.PublishMessage) error
	ClearRetainMessage(ctx context.Context, topic string) error
	GetRetainMessage(ctx context.Context, topic string) (*messagev5.PublishMessage, error)

	GetAllRetainMsg(ctx context.Context) ([]*messagev5.PublishMessage, error)
}
