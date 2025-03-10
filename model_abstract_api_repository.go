/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AbstractApiRepository struct {
	// A unique identifier for this repository
	Name string `json:"name,omitempty"`
	// Component format held in this repository
	Format string `json:"format,omitempty"`
	// Controls if deployments of and updates to artifacts are allowed
	Type_ string `json:"type,omitempty"`
	// URL to the repository
	Url string `json:"url,omitempty"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
}
