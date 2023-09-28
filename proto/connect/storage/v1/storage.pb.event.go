package storagev1

import (
	"context"
	"errors"
)

// The error occurs when the event is unknown.
var ErrStorageEventTypeUnknown = errors.New("unknown storage event type")

//counterfeiter:generate -o storagev1fake . StorageObjectEvent

// StorageObjectEvent is the interface for events that contain a StorageObject.
type StorageObjectEvent interface {
	GetObject() *StorageObject
	SetObject(*StorageObject)
}

var _ StorageObjectEvent = &StorageObjectDeletedEvent{}

// SetStorageObject sets the StorageObject.
func (x *StorageObjectDeletedEvent) SetObject(obj *StorageObject) {
	x.Object = obj
}

var _ StorageObjectEvent = &StorageObjectArchivedEvent{}

// SetStorageObject sets the StorageObject.
func (x *StorageObjectArchivedEvent) SetObject(obj *StorageObject) {
	x.Object = obj
}

var _ StorageObjectEvent = &StorageObjectFinalizedEvent{}

// SetStorageObject sets the StorageObject.
func (x *StorageObjectFinalizedEvent) SetObject(obj *StorageObject) {
	x.Object = obj
}

var _ StorageObjectEvent = &StorageObjectMetadataUpdatedEvent{}

// SetStorageObject sets the StorageObject.
func (x *StorageObjectMetadataUpdatedEvent) SetObject(obj *StorageObject) {
	x.Object = obj
}

//counterfeiter:generate -o storagev1fake . StorageObjectEventHandler

// StorageEventHandler handles connect.storage.v1.StorageObjectEvent event.
type StorageObjectEventHandler interface {
	// HandleStorageObjectEvent handles the connect.storage.v1.StorageObjectEvent event.
	HandleStorageObjectEvent(context.Context, StorageObjectEvent) error
}
