package scanner

import (
	"fmt"
	"net/http"
)

// SimpleMobileProvider implements a basic Provider for mobile apps
type SimpleMobileProvider struct {
	Name string
}

// NewSimpleMobileProvider creates a new provider instance
func NewSimpleMobileProvider(name string) *SimpleMobileProvider {
	return &SimpleMobileProvider{Name: name}
}

// CheckBucket implements the Provider interface
func (p *SimpleMobileProvider) CheckBucket(bucketName string) (*BucketInfo, error) {
	info := &BucketInfo{
		Name: bucketName,
	}

	// Try to access the bucket with the appropriate URL format
	var url string
	switch p.Name {
	case "aws":
		url = fmt.Sprintf("https://%s.s3.amazonaws.com", bucketName)
	case "gcp":
		url = fmt.Sprintf("https://storage.googleapis.com/%s", bucketName)
	case "azure":
		url = fmt.Sprintf("https://%s.blob.core.windows.net", bucketName)
	case "digitalocean":
		url = fmt.Sprintf("https://%s.nyc3.digitaloceanspaces.com", bucketName)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", p.Name)
	}

	resp, err := http.Head(url)
	if err != nil {
		// Network error, but bucket might still exist
		info.Exists = false
		info.IsPublic = false
		return info, nil
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK, http.StatusForbidden:
		info.Exists = true
		info.IsPublic = resp.StatusCode == http.StatusOK
	case http.StatusNotFound:
		info.Exists = false
		info.IsPublic = false
	default:
		info.Exists = true
		info.IsPublic = false
	}

	// Try to get region from headers
	if region := resp.Header.Get("x-amz-bucket-region"); region != "" {
		info.Region = region
	} else {
		info.Region = "unknown"
	}

	return info, nil
}

// MobileScanner provides the mobile interface for S3 scanning
type MobileScanner struct {
	providers map[string]*SimpleMobileProvider
}

// NewMobileScanner creates a new MobileScanner instance
func NewMobileScanner() *MobileScanner {
	return &MobileScanner{
		providers: make(map[string]*SimpleMobileProvider),
	}
}

// ScanBucket performs a scan of the specified bucket
func (s *MobileScanner) ScanBucket(bucketName string, providerName string) *ScanResult {
	result := &ScanResult{
		BucketName: bucketName,
	}

	// Validate bucket name
	if !IsValidS3BucketName(bucketName) {
		result.Error = fmt.Errorf("invalid bucket name: %s", bucketName)
		return result
	}

	// Get or create provider
	provider, exists := s.providers[providerName]
	if !exists {
		provider = NewSimpleMobileProvider(providerName)
		s.providers[providerName] = provider
	}

	// Create bucket and scan
	bucket := NewBucket(bucketName)
	scanResult, err := bucket.Scan(provider)
	if err != nil {
		result.Error = err
		return result
	}

	result.Region = scanResult.Region
	result.Exists = scanResult.Exists
	result.IsPublic = scanResult.IsPublic

	return result
}

// GetVersion returns the version of the scanner
func (s *MobileScanner) GetVersion() string {
	return "1.0.0"
}