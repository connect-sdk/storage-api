package storagev1

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

//go:generate counterfeiter -generate

//counterfeiter:generate -o storagev1fake . StoragePageInfo

// StoragePageInfo contains information about an iterator's paging state.
type StoragePageInfo interface {
	Token() string
	MaxSize() int
	Remaining() int
}

//counterfeiter:generate -o storagev1fake . StorageServiceClient

type StorageServiceClient interface {
	Bucket(name string) StorageBucketHandle
	Buckets(ctx context.Context, projectID string) StorageBucketIterator
	Close() error
}

//counterfeiter:generate -o storagev1fake . StorageBucketIterator

type StorageBucketIterator interface {
	Next() (*storage.BucketAttrs, error)
	PageInfo() StoragePageInfo
}

//counterfeiter:generate -o storagev1fake . StorageBucketHandle

type StorageBucketHandle interface {
	Create(context.Context, string, *storage.BucketAttrs) error
	Delete(context.Context) error
	DefaultObjectACL() StorageACLHandle
	Object(string) StorageObjectHandle
	Attrs(context.Context) (*storage.BucketAttrs, error)
	Update(context.Context, storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error)
	If(storage.BucketConditions) StorageBucketHandle
	Objects(context.Context, *storage.Query) StorageObjectIterator
	ACL() StorageACLHandle
	UserProject(projectID string) StorageBucketHandle
	Notifications(context.Context) (map[string]*storage.Notification, error)
	AddNotification(context.Context, *storage.Notification) (*storage.Notification, error)
	DeleteNotification(context.Context, string) error
	LockRetentionPolicy(context.Context) error
}

//counterfeiter:generate -o storagev1fake . StorageObjectIterator

type StorageObjectIterator interface {
	Next() (*storage.ObjectAttrs, error)
	PageInfo() StoragePageInfo
}

//counterfeiter:generate -o storagev1fake . StorageObjectHandle

type StorageObjectHandle interface {
	ACL() StorageACLHandle
	Generation(int64) StorageObjectHandle
	If(storage.Conditions) StorageObjectHandle
	Key([]byte) StorageObjectHandle
	ReadCompressed(bool) StorageObjectHandle
	Attrs(context.Context) (*storage.ObjectAttrs, error)
	Update(context.Context, storage.ObjectAttrsToUpdate) (*storage.ObjectAttrs, error)
	NewReader(context.Context) (StorageObjectReader, error)
	NewRangeReader(context.Context, int64, int64) (StorageObjectReader, error)
	NewWriter(context.Context) StorageObjectWriter
	Delete(context.Context) error
	CopierFrom(StorageObjectHandle) StorageObjectCopier
	ComposerFrom(...StorageObjectHandle) StorageObjectComposer
}

//counterfeiter:generate -o storagev1fake . StorageACLHandle

type StorageACLHandle interface {
	Delete(context.Context, storage.ACLEntity) error
	Set(context.Context, storage.ACLEntity, storage.ACLRole) error
	List(context.Context) ([]storage.ACLRule, error)
}

//counterfeiter:generate -o storagev1fake . StorageObjectReader

type StorageObjectReader interface {
	io.ReadCloser
	Size() int64
	Remain() int64
	ContentType() string
	ContentEncoding() string
	CacheControl() string
}

//counterfeiter:generate -o storagev1fake . StorageObjectWriter

type StorageObjectWriter interface {
	io.WriteCloser
	ObjectAttrs() *storage.ObjectAttrs
	SetChunkSize(int)
	SetProgressFunc(func(int64))
	SetContentType(string)
	SetCRC32C(uint32)
	CloseWithError(err error) error
	Attrs() *storage.ObjectAttrs
}

//counterfeiter:generate -o storagev1fake . StorageObjectCopier

type StorageObjectCopier interface {
	ObjectAttrs() *storage.ObjectAttrs
	SetRewriteToken(string)
	SetProgressFunc(func(uint64, uint64))
	SetDestinationKMSKeyName(string)
	Run(context.Context) (*storage.ObjectAttrs, error)
}

//counterfeiter:generate -o storagev1fake . StorageObjectComposer

type StorageObjectComposer interface {
	ObjectAttrs() *storage.ObjectAttrs
	Run(context.Context) (*storage.ObjectAttrs, error)
}

var _ StorageServiceClient = &NopStorageServiceClient{}

// NopStorageServiceClient is a no-op implementation of StorageServiceClient.
type NopStorageServiceClient struct{}

// Bucket implements StorageServiceClient.
func (x *NopStorageServiceClient) Bucket(name string) StorageBucketHandle {
	return &NopStorageBucketHandle{}
}

// Buckets implements StorageServiceClient.
func (x *NopStorageServiceClient) Buckets(ctx context.Context, projectID string) StorageBucketIterator {
	return &NopStorageBucketIterator{}
}

// Close implements StorageServiceClient.
func (x *NopStorageServiceClient) Close() error {
	return nil
}

var _ StorageBucketHandle = &NopStorageBucketHandle{}

// NopStorageBucketHandle is a no-op implementation of StorageBucketHandle.
type NopStorageBucketHandle struct{}

// ACL implements StorageBucketHandle.
func (*NopStorageBucketHandle) ACL() StorageACLHandle {
	return &NopStorageACLHandle{}
}

// AddNotification implements StorageBucketHandle.
func (*NopStorageBucketHandle) AddNotification(ctx context.Context, x *storage.Notification) (*storage.Notification, error) {
	return x, nil
}

// Attrs implements StorageBucketHandle.
func (*NopStorageBucketHandle) Attrs(context.Context) (*storage.BucketAttrs, error) {
	return &storage.BucketAttrs{}, nil
}

// Create implements StorageBucketHandle.
func (*NopStorageBucketHandle) Create(context.Context, string, *storage.BucketAttrs) error {
	return nil
}

// DefaultObjectACL implements StorageBucketHandle.
func (*NopStorageBucketHandle) DefaultObjectACL() StorageACLHandle {
	return &NopStorageACLHandle{}
}

// Delete implements StorageBucketHandle.
func (*NopStorageBucketHandle) Delete(context.Context) error {
	return nil
}

// DeleteNotification implements StorageBucketHandle.
func (*NopStorageBucketHandle) DeleteNotification(context.Context, string) error {
	return nil
}

// If implements StorageBucketHandle.
func (*NopStorageBucketHandle) If(storage.BucketConditions) StorageBucketHandle {
	return nil
}

// LockRetentionPolicy implements StorageBucketHandle.
func (*NopStorageBucketHandle) LockRetentionPolicy(context.Context) error {
	return nil
}

// Notifications implements StorageBucketHandle.
func (*NopStorageBucketHandle) Notifications(context.Context) (map[string]*storage.Notification, error) {
	return make(map[string]*storage.Notification), nil
}

// Object implements StorageBucketHandle.
func (*NopStorageBucketHandle) Object(string) StorageObjectHandle {
	return &NopStorageObjectHandle{}
}

// Objects implements StorageBucketHandle.
func (*NopStorageBucketHandle) Objects(context.Context, *storage.Query) StorageObjectIterator {
	return &NopStorageObjectIterator{}
}

// Update implements StorageBucketHandle.
func (*NopStorageBucketHandle) Update(context.Context, storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error) {
	return &storage.BucketAttrs{}, nil
}

// UserProject implements StorageBucketHandle.
func (*NopStorageBucketHandle) UserProject(projectID string) StorageBucketHandle {
	return &NopStorageBucketHandle{}
}

var _ StorageBucketIterator = &NopStorageBucketIterator{}

// NopStorageBucketIterator is a no-op implementation of StorageBucketIterator.
type NopStorageBucketIterator struct{}

// Next implements StorageBucketIterator.
func (*NopStorageBucketIterator) Next() (*storage.BucketAttrs, error) {
	return nil, iterator.Done
}

// PageInfo implements StorageBucketIterator.
func (x *NopStorageBucketIterator) PageInfo() StoragePageInfo {
	return &NopStoragePageInfo{}
}

var _ StorageACLHandle = &NopStorageACLHandle{}

// NopStorageACLHandle is a no-op implementation of StorageACLHandle.
type NopStorageACLHandle struct{}

// Delete implements StorageACLHandle.
func (*NopStorageACLHandle) Delete(context.Context, storage.ACLEntity) error {
	return nil
}

// List implements StorageACLHandle.
func (*NopStorageACLHandle) List(context.Context) ([]storage.ACLRule, error) {
	return []storage.ACLRule{}, nil
}

// Set implements StorageACLHandle.
func (*NopStorageACLHandle) Set(context.Context, storage.ACLEntity, storage.ACLRole) error {
	return nil
}

var _ StorageObjectHandle = &NopStorageObjectHandle{}

// NopStorageObjectHandle is a no-op implementation of StorageObjectHandle.
type NopStorageObjectHandle struct{}

// ACL implements StorageObjectHandle.
func (*NopStorageObjectHandle) ACL() StorageACLHandle {
	return &NopStorageACLHandle{}
}

// Attrs implements StorageObjectHandle.
func (*NopStorageObjectHandle) Attrs(context.Context) (*storage.ObjectAttrs, error) {
	return &storage.ObjectAttrs{}, nil
}

// ComposerFrom implements StorageObjectHandle.
func (*NopStorageObjectHandle) ComposerFrom(...StorageObjectHandle) StorageObjectComposer {
	return &NopStorageObjectComposer{}
}

// CopierFrom implements StorageObjectHandle.
func (*NopStorageObjectHandle) CopierFrom(StorageObjectHandle) StorageObjectCopier {
	return &NopStorageObjectCopier{}
}

// Delete implements StorageObjectHandle.
func (*NopStorageObjectHandle) Delete(context.Context) error {
	return nil
}

// Generation implements StorageObjectHandle.
func (*NopStorageObjectHandle) Generation(int64) StorageObjectHandle {
	return &NopStorageObjectHandle{}
}

// If implements StorageObjectHandle.
func (*NopStorageObjectHandle) If(storage.Conditions) StorageObjectHandle {
	return &NopStorageObjectHandle{}
}

// Key implements StorageObjectHandle.
func (*NopStorageObjectHandle) Key([]byte) StorageObjectHandle {
	return &NopStorageObjectHandle{}
}

// NewRangeReader implements StorageObjectHandle.
func (*NopStorageObjectHandle) NewRangeReader(context.Context, int64, int64) (StorageObjectReader, error) {
	panic("unimplemented")
}

// NewReader implements StorageObjectHandle.
func (*NopStorageObjectHandle) NewReader(context.Context) (StorageObjectReader, error) {
	panic("unimplemented")
}

// NewWriter implements StorageObjectHandle.
func (*NopStorageObjectHandle) NewWriter(context.Context) StorageObjectWriter {
	panic("unimplemented")
}

// ReadCompressed implements StorageObjectHandle.
func (*NopStorageObjectHandle) ReadCompressed(bool) StorageObjectHandle {
	return &NopStorageObjectHandle{}
}

// Update implements StorageObjectHandle.
func (*NopStorageObjectHandle) Update(context.Context, storage.ObjectAttrsToUpdate) (*storage.ObjectAttrs, error) {
	return &storage.ObjectAttrs{}, nil
}

var _ StorageObjectIterator = &NopStorageObjectIterator{}

// NopStorageObjectIterator is a no-op implementation of StorageObjectIterator.
type NopStorageObjectIterator struct{}

// Next implements StorageObjectIterator.
func (*NopStorageObjectIterator) Next() (*storage.ObjectAttrs, error) {
	return nil, iterator.Done
}

// PageInfo implements StorageObjectIterator.
func (*NopStorageObjectIterator) PageInfo() StoragePageInfo {
	return &NopStoragePageInfo{}
}

var _ StorageObjectCopier = &NopStorageObjectCopier{}

// NopStorageObjectCopier is a no-op implementation of StorageObjectCopier.
type NopStorageObjectCopier struct{}

// ObjectAttrs implements StorageObjectCopier.
func (*NopStorageObjectCopier) ObjectAttrs() *storage.ObjectAttrs {
	return &storage.ObjectAttrs{}
}

// Run implements StorageObjectCopier.
func (*NopStorageObjectCopier) Run(context.Context) (*storage.ObjectAttrs, error) {
	return &storage.ObjectAttrs{}, nil
}

// SetDestinationKMSKeyName implements StorageObjectCopier.
func (*NopStorageObjectCopier) SetDestinationKMSKeyName(string) {

}

// SetProgressFunc implements StorageObjectCopier.
func (*NopStorageObjectCopier) SetProgressFunc(func(uint64, uint64)) {

}

// SetRewriteToken implements StorageObjectCopier.
func (*NopStorageObjectCopier) SetRewriteToken(string) {
	panic("unimplemented")
}

var _ StorageObjectComposer = &NopStorageObjectComposer{}

// NopStorageObjectComposer is a no-op implementation of StorageObjectComposer.
type NopStorageObjectComposer struct{}

// ObjectAttrs implements StorageObjectComposer.
func (*NopStorageObjectComposer) ObjectAttrs() *storage.ObjectAttrs {
	return &storage.ObjectAttrs{}
}

// Run implements StorageObjectComposer.
func (*NopStorageObjectComposer) Run(context.Context) (*storage.ObjectAttrs, error) {
	return &storage.ObjectAttrs{}, nil
}

var _ StoragePageInfo = &NopStoragePageInfo{}

// NopStoragePageInfo is a no-op implementation of StoragePageInfo.
type NopStoragePageInfo struct{}

// MaxSize implements StoragePageInfo.
func (*NopStoragePageInfo) MaxSize() int {
	return 100
}

// Remaining implements StoragePageInfo.
func (*NopStoragePageInfo) Remaining() int {
	return 0
}

// Token implements StoragePageInfo.
func (*NopStoragePageInfo) Token() string {
	return ""
}
