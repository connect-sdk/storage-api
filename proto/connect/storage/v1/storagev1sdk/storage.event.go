package storagev1sdk

import (
	"context"
	"encoding/json"

	connect "connectrpc.com/connect"
	pubsubv1 "github.com/connect-sdk/pubsub-api/proto/connect/pubsub/v1"

	storagev1 "github.com/connect-sdk/storage-api/proto/connect/storage/v1"
)

var _ pubsubv1.PubsubService = &StorageEventPubsubService{}

// StorageEventPubsubService represents a storage pubsub service.
type StorageEventPubsubService struct {
	// StorageObjectEventHandler is a storage object event handler.
	StorageObjectEventHandler storagev1.StorageObjectEventHandler
}

// PushPubsubMessage implements pubsubv1.PubsubService.
func (x *StorageEventPubsubService) PushPubsubMessage(ctx context.Context, r *pubsubv1.PushPubsubMessageRequest) (*pubsubv1.PushPubsubMessageResponse, error) {
	var event storagev1.StorageObjectEvent
	// prepare the event
	switch r.Message.Attributes["eventType"] {
	case storagev1.StorageObjectEventType_OBJECT_FINALIZE.String():
		event = &storagev1.StorageObjectFinalizedEvent{}
	case storagev1.StorageObjectEventType_OBJECT_ARCHIVE.String():
		event = &storagev1.StorageObjectArchivedEvent{}
	case storagev1.StorageObjectEventType_OBJECT_DELETE.String():
		event = &storagev1.StorageObjectDeletedEvent{}
	case storagev1.StorageObjectEventType_OBJECT_METADATA_UPDATE.String():
		event = &storagev1.StorageObjectMetadataUpdatedEvent{}
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, storagev1.ErrStorageEventTypeUnknown)
	}

	resource := &storagev1.StorageObject{}
	// unmarshal the object
	if err := json.Unmarshal(r.Message.Data, resource); err != nil {
		return nil, err
	}

	event.SetObject(resource)
	// handle the event
	if err := x.StorageObjectEventHandler.HandleStorageObjectEvent(ctx, event); err != nil {
		return nil, err
	}

	return &pubsubv1.PushPubsubMessageResponse{}, nil
}
