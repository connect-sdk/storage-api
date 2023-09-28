package storagev1sdk

import (
	"context"

	storage "cloud.google.com/go/storage"
	iterator "google.golang.org/api/iterator"
	option "google.golang.org/api/option"

	storagev1 "github.com/connect-sdk/storage-api/proto/connect/storage/v1"
)

var _ storagev1.StorageServiceClient = &StorageServiceClient{}

// StorageServiceClient represents a storage service client.
type StorageServiceClient struct {
	client *storage.Client
}

// NewStorageServiceClient creates a new StorageServiceClient.
func NewStorageServiceClient(ctx context.Context, options ...option.ClientOption) (*StorageServiceClient, error) {
	client, err := storage.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &StorageServiceClient{client: client}, nil
}

// Bucket implements storagev1.StorageServiceClient.
func (x *StorageServiceClient) Bucket(name string) storagev1.StorageBucketHandle {
	handle := x.client.Bucket(name)
	// done!
	return &StorageBucketHandle{handle: handle}
}

// Buckets implements storagev1.StorageServiceClient.
func (x *StorageServiceClient) Buckets(ctx context.Context, projectID string) storagev1.StorageBucketIterator {
	iterator := x.client.Buckets(ctx, projectID)
	// done
	return &StorageBucketIterator{iterator: iterator}
}

// Close implements storagev1.StorageServiceClient.
func (x *StorageServiceClient) Close() error {
	return x.client.Close()
}

var _ storagev1.StorageBucketHandle = &StorageBucketHandle{}

type StorageBucketHandle struct {
	handle *storage.BucketHandle
}

// ACL implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) ACL() storagev1.StorageACLHandle {
	handle := x.handle.ACL()
	// done!
	return &StorageACLHandle{handle: handle}
}

// AddNotification implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) AddNotification(ctx context.Context, notification *storage.Notification) (*storage.Notification, error) {
	result, err := x.handle.AddNotification(ctx, (*storage.Notification)(notification))
	if err != nil {
		return nil, err
	}

	return (*storage.Notification)(result), nil
}

// Attrs implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) Attrs(ctx context.Context) (*storage.BucketAttrs, error) {
	result, err := x.handle.Attrs(ctx)
	if err != nil {
		return nil, err
	}

	return (*storage.BucketAttrs)(result), nil
}

// Create implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) Create(ctx context.Context, projectID string, attr *storage.BucketAttrs) error {
	return x.handle.Create(ctx, projectID, (*storage.BucketAttrs)(attr))
}

// DefaultObjectACL implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) DefaultObjectACL() storagev1.StorageACLHandle {
	handle := x.handle.DefaultObjectACL()
	// done!
	return &StorageACLHandle{handle: handle}
}

// Delete implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) Delete(ctx context.Context) error {
	return x.handle.Delete(ctx)
}

// DeleteNotification implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) DeleteNotification(ctx context.Context, id string) error {
	return x.handle.DeleteNotification(ctx, id)
}

// If implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) If(cond storage.BucketConditions) storagev1.StorageBucketHandle {
	handle := x.handle.If(cond)
	// done!
	return &StorageBucketHandle{handle: handle}
}

// LockRetentionPolicy implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) LockRetentionPolicy(ctx context.Context) error {
	return x.handle.LockRetentionPolicy(ctx)
}

// Notifications implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) Notifications(ctx context.Context) (map[string]*storage.Notification, error) {
	result, err := x.handle.Notifications(ctx)
	if err != nil {
		return nil, err
	}

	dictionary := make(map[string]*storage.Notification)
	// translate the dictionary
	for k, v := range result {
		dictionary[k] = (*storage.Notification)(v)
	}
	// done!
	return dictionary, nil
}

// Object implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) Object(name string) storagev1.StorageObjectHandle {
	handle := x.handle.Object(name)
	// done!
	return &StorageObjectHandle{handle: handle}
}

// Objects implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) Objects(ctx context.Context, query *storage.Query) storagev1.StorageObjectIterator {
	iterator := x.handle.Objects(ctx, (*storage.Query)(query))
	// done
	return &StorageObjectIterator{iterator: iterator}
}

// Update implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) Update(ctx context.Context, attr storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error) {
	result, err := x.handle.Update(ctx, storage.BucketAttrsToUpdate(attr))
	if err != nil {
		return nil, err
	}
	// done!
	return (*storage.BucketAttrs)(result), nil
}

// UserProject implements storagev1.StorageBucketHandle.
func (x *StorageBucketHandle) UserProject(projectID string) storagev1.StorageBucketHandle {
	handle := x.handle.UserProject(projectID)
	// done!
	return &StorageBucketHandle{handle: handle}
}

var _ storagev1.StorageBucketIterator = &StorageBucketIterator{}

type StorageBucketIterator struct {
	iterator *storage.BucketIterator
}

// Next implements storagev1.StorageBucketIterator.
func (x *StorageBucketIterator) Next() (*storage.BucketAttrs, error) {
	result, err := x.iterator.Next()
	if err != nil {
		return nil, err
	}

	return (*storage.BucketAttrs)(result), nil
}

// PageInfo implements storagev1.StorageBucketIterator.
func (x *StorageBucketIterator) PageInfo() storagev1.StoragePageInfo {
	info := x.iterator.PageInfo()
	return &StoragePageInfo{info: info}
}

var _ storagev1.StorageACLHandle = &StorageACLHandle{}

type StorageACLHandle struct {
	handle *storage.ACLHandle
}

// Delete implements storagev1.StorageACLHandle.
func (x *StorageACLHandle) Delete(ctx context.Context, entity storage.ACLEntity) error {
	return x.handle.Delete(ctx, storage.ACLEntity(entity))
}

// List implements storagev1.StorageACLHandle.
func (x *StorageACLHandle) List(ctx context.Context) ([]storage.ACLRule, error) {
	items, err := x.handle.List(ctx)
	if err != nil {
		return nil, err
	}

	collection := make([]storage.ACLRule, len(items))

	for index, item := range items {
		collection[index] = storage.ACLRule(item)
	}

	return collection, nil
}

// Set implements storagev1.StorageACLHandle.
func (x *StorageACLHandle) Set(ctx context.Context, entity storage.ACLEntity, role storage.ACLRole) error {
	return x.handle.Set(ctx, storage.ACLEntity(entity), storage.ACLRole(role))
}

var _ storagev1.StorageObjectHandle = &StorageObjectHandle{}

type StorageObjectHandle struct {
	handle *storage.ObjectHandle
}

// ACL implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) ACL() storagev1.StorageACLHandle {
	result := x.handle.ACL()
	// done!
	return &StorageACLHandle{handle: result}
}

// Attrs implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) Attrs(ctx context.Context) (*storage.ObjectAttrs, error) {
	result, err := x.handle.Attrs(ctx)
	if err != nil {
		return nil, err
	}

	return (*storage.ObjectAttrs)(result), nil
}

// ComposerFrom implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) ComposerFrom(handles ...storagev1.StorageObjectHandle) storagev1.StorageObjectComposer {
	args := make([]*storage.ObjectHandle, len(handles))

	for index, item := range handles {
		args[index] = item.(*StorageObjectHandle).handle
	}

	composer := x.handle.ComposerFrom(args...)
	return &StorageComposer{composer: composer}
}

// CopierFrom implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) CopierFrom(handle storagev1.StorageObjectHandle) storagev1.StorageObjectCopier {
	copier := x.handle.CopierFrom(handle.(*StorageObjectHandle).handle)
	return &StorageCopier{copier: copier}
}

// Delete implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) Delete(ctx context.Context) error {
	return x.handle.Delete(ctx)
}

// Generation implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) Generation(gen int64) storagev1.StorageObjectHandle {
	result := x.handle.Generation(gen)
	return &StorageObjectHandle{handle: result}
}

// If implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) If(cond storage.Conditions) storagev1.StorageObjectHandle {
	result := x.handle.If(cond)
	return &StorageObjectHandle{handle: result}
}

// Key implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) Key(data []byte) storagev1.StorageObjectHandle {
	result := x.handle.Key(data)
	return &StorageObjectHandle{handle: result}
}

// NewRangeReader implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) NewRangeReader(ctx context.Context, start int64, end int64) (storagev1.StorageObjectReader, error) {
	return x.handle.NewRangeReader(ctx, start, end)
}

// NewReader implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) NewReader(ctx context.Context) (storagev1.StorageObjectReader, error) {
	return x.handle.NewReader(ctx)
}

// NewWriter implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) NewWriter(ctx context.Context) storagev1.StorageObjectWriter {
	writer := x.handle.NewWriter(ctx)
	return &StorageWriter{writer: writer}
}

// ReadCompressed implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) ReadCompressed(compressed bool) storagev1.StorageObjectHandle {
	result := x.handle.ReadCompressed(compressed)
	return &StorageObjectHandle{handle: result}
}

// Update implements storagev1.StorageObjectHandle.
func (x *StorageObjectHandle) Update(ctx context.Context, uattrs storage.ObjectAttrsToUpdate) (*storage.ObjectAttrs, error) {
	result, err := x.handle.Update(ctx, storage.ObjectAttrsToUpdate(uattrs))
	if err != nil {
		return nil, err
	}

	return (*storage.ObjectAttrs)(result), nil
}

var _ storagev1.StorageObjectIterator = &StorageObjectIterator{}

type StorageObjectIterator struct {
	iterator *storage.ObjectIterator
}

// Next implements storagev1.StorageObjectIterator.
func (x *StorageObjectIterator) Next() (*storage.ObjectAttrs, error) {
	result, err := x.iterator.Next()
	if err != nil {
		return nil, err
	}

	return (*storage.ObjectAttrs)(result), nil
}

// PageInfo implements storagev1.StorageObjectIterator.
func (x *StorageObjectIterator) PageInfo() storagev1.StoragePageInfo {
	info := x.iterator.PageInfo()
	return &StoragePageInfo{info: info}
}

var _ storagev1.StorageObjectWriter = &StorageWriter{}

type StorageWriter struct {
	writer *storage.Writer
}

// Attrs implements storagev1.StorageWriter.
func (x *StorageWriter) Attrs() *storage.ObjectAttrs {
	result := x.writer.Attrs()
	return (*storage.ObjectAttrs)(result)
}

// Close implements storagev1.StorageWriter.
func (x *StorageWriter) Close() error {
	return x.writer.Close()
}

// CloseWithError implements storagev1.StorageWriter.
func (x *StorageWriter) CloseWithError(err error) error {
	return x.writer.CloseWithError(err)
}

// ObjectAttrs implements storagev1.StorageWriter.
func (x *StorageWriter) ObjectAttrs() *storage.ObjectAttrs {
	return (*storage.ObjectAttrs)(&x.writer.ObjectAttrs)
}

// SetCRC32C implements storagev1.StorageWriter.
func (x *StorageWriter) SetCRC32C(v uint32) {
	x.writer.CRC32C = v
}

// SetChunkSize implements storagev1.StorageWriter.
func (x *StorageWriter) SetChunkSize(v int) {
	x.writer.ChunkSize = v
}

// SetProgressFunc implements storagev1.StorageWriter.
func (x *StorageWriter) SetProgressFunc(fn func(int64)) {
	x.writer.ProgressFunc = fn
}

// SetContentType implements storagev1.StorageWriter.
func (x *StorageWriter) SetContentType(value string) {
	x.writer.ContentType = value
}

// Write implements storagev1.StorageWriter.
func (x *StorageWriter) Write(p []byte) (n int, err error) {
	return x.writer.Write(p)
}

var _ storagev1.StorageObjectCopier = &StorageCopier{}

type StorageCopier struct {
	copier *storage.Copier
}

// ObjectAttrs implements storagev1.StorageCopier.
func (x *StorageCopier) ObjectAttrs() *storage.ObjectAttrs {
	result := x.copier.ObjectAttrs
	// done!
	return (*storage.ObjectAttrs)(&result)
}

// Run implements storagev1.StorageCopier.
func (x *StorageCopier) Run(ctx context.Context) (*storage.ObjectAttrs, error) {
	result, err := x.copier.Run(ctx)
	if err != nil {
		return nil, err
	}

	return (*storage.ObjectAttrs)(result), nil
}

// SetDestinationKMSKeyName implements storagev1.StorageCopier.
func (x *StorageCopier) SetDestinationKMSKeyName(name string) {
	x.copier.DestinationKMSKeyName = name
}

// SetProgressFunc implements storagev1.StorageCopier.
func (x *StorageCopier) SetProgressFunc(fn func(uint64, uint64)) {
	x.copier.ProgressFunc = fn
}

// SetRewriteToken implements storagev1.StorageCopier.
func (x *StorageCopier) SetRewriteToken(token string) {
	x.copier.RewriteToken = token
}

var _ storagev1.StorageObjectComposer = &StorageComposer{}

type StorageComposer struct {
	composer *storage.Composer
}

// ObjectAttrs implements storagev1.StorageComposer.
func (x *StorageComposer) ObjectAttrs() *storage.ObjectAttrs {
	result := x.composer.ObjectAttrs
	// done!
	return (*storage.ObjectAttrs)(&result)
}

// Run implements storagev1.StorageComposer.
func (x *StorageComposer) Run(ctx context.Context) (*storage.ObjectAttrs, error) {
	result, err := x.composer.Run(ctx)
	if err != nil {
		return nil, err
	}

	return (*storage.ObjectAttrs)(result), nil
}

var _ storagev1.StoragePageInfo = &StoragePageInfo{}

// StoragePageInfo represents a storage page info.
type StoragePageInfo struct {
	info *iterator.PageInfo
}

// MaxSize implements storagev1.StoragePageInfo.
func (x *StoragePageInfo) MaxSize() int {
	return x.info.MaxSize
}

// Remaining implements storagev1.StoragePageInfo.
func (x *StoragePageInfo) Remaining() int {
	return x.info.Remaining()
}

// Token implements storagev1.StoragePageInfo.
func (x *StoragePageInfo) Token() string {
	return x.info.Token
}
