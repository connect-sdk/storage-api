package storagev1fake

import (
	"bytes"

	"cloud.google.com/go/storage"
)

func NewFakeStorageBucketHandle() *FakeStorageBucketHandle {
	object := NewFakeStorageObjectHandle()
	// prepare the storage bucket mock
	bucket := &FakeStorageBucketHandle{}
	bucket.ObjectReturns(object)

	return bucket
}

func NewFakeStorageObjectHandle() *FakeStorageObjectHandle {
	// prepare the storage writer mock
	writer := &FakeStorageObjectWriter{}
	writer.ObjectAttrsReturns(&storage.ObjectAttrs{})
	// prepare the storage reader mock
	reader := &FakeStorageObjectReader{}
	reader.ReadStub = bytes.NewBuffer(nil).Read
	// prepare the storage object mock
	object := &FakeStorageObjectHandle{}
	object.NewWriterReturns(writer)
	object.NewReaderReturns(reader, nil)

	return object
}
