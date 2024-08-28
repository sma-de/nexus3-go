/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type S3BlobStoreApiAdvancedBucketConnection struct {
	// A custom endpoint URL for third party object stores using the S3 API.
	Endpoint string `json:"endpoint,omitempty"`
	// An API signature version which may be required for third party object stores using the S3 API.
	SignerType string `json:"signerType,omitempty"`
	// Setting this flag will result in path-style access being used for all requests.
	ForcePathStyle bool `json:"forcePathStyle,omitempty"`
	// Setting this value will override the default connection pool size of Nexus of the s3 client for this blobstore.
	MaxConnectionPoolSize int32 `json:"maxConnectionPoolSize,omitempty"`
}
