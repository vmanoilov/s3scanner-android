package scanner

// BucketInfo represents information about a bucket
type BucketInfo struct {
	Name     string
	Region   string
	Exists   bool
	IsPublic bool
}

// Bucket represents an S3 bucket
type Bucket struct {
	Name string
}

// Provider interface defines methods that cloud providers must implement
type Provider interface {
	CheckBucket(bucketName string) (*BucketInfo, error)
}

// ScanResult represents the result of a bucket scan
type ScanResult struct {
	BucketName string
	Region     string
	Exists     bool
	IsPublic   bool
	Error      error
}

// NewBucket creates a new Bucket instance
func NewBucket(name string) *Bucket {
	return &Bucket{Name: name}
}

// IsValidS3BucketName checks if a bucket name is valid according to S3 naming rules
func IsValidS3BucketName(name string) bool {
	// Basic validation - can be expanded based on full S3 bucket naming rules
	if len(name) < 3 || len(name) > 63 {
		return false
	}
	return true
}

// Scan performs a scan of the bucket using the provided provider
func (b *Bucket) Scan(p Provider) (*ScanResult, error) {
	info, err := p.CheckBucket(b.Name)
	if err != nil {
		return &ScanResult{
			BucketName: b.Name,
			Error:      err,
		}, err
	}

	return &ScanResult{
		BucketName: b.Name,
		Region:     info.Region,
		Exists:     info.Exists,
		IsPublic:   info.IsPublic,
		Error:      nil,
	}, nil
}

// GetSupportedProviders returns a list of supported cloud providers
func GetSupportedProviders() []string {
	return []string{
		"aws",
		"gcp",
		"azure",
		"digitalocean",
		"custom",
	}
}
