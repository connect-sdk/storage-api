package storagev1

// Key returns the key representation.
func (x ParsedProjectName) Key() string {
	return x.ProjectID
}

// String returns the string representation.
func (x ParsedProjectName) String() string {
	return x.Name()
}

// Key returns the key representation.
func (x ParsedBucketName) Key() string {
	return x.BucketID
}

// Parent returns the parent representation.
func (x ParsedBucketName) Parent() ParsedProjectName {
	return ParsedProjectName{ProjectID: x.ProjectID}
}

// String returns the string representation.
func (x ParsedBucketName) String() string {
	return x.Name()
}

// Key returns the key representation.
func (x ParsedObjectName) Key() string {
	return x.ObjectID
}

// Parent returns the parent representation.
func (x ParsedObjectName) Parent() ParsedBucketName {
	return ParsedBucketName{
		ProjectID: x.ProjectID,
		BucketID:  x.BucketID,
	}
}

// String returns the string representation.
func (x ParsedObjectName) String() string {
	return x.Name()
}
