/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type AssetXo struct {
	DownloadUrl string `json:"downloadUrl,omitempty"`
	Path string `json:"path,omitempty"`
	Id string `json:"id,omitempty"`
	Repository string `json:"repository,omitempty"`
	Format string `json:"format,omitempty"`
	Checksum map[string]interface{} `json:"checksum,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	LastModified time.Time `json:"lastModified,omitempty"`
	LastDownloaded time.Time `json:"lastDownloaded,omitempty"`
	Uploader string `json:"uploader,omitempty"`
	UploaderIp string `json:"uploaderIp,omitempty"`
	FileSize int64 `json:"fileSize,omitempty"`
	BlobCreated time.Time `json:"blobCreated,omitempty"`
}
