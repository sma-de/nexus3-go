/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PageTaskXo struct {
	Items []TaskXo `json:"items,omitempty"`
	ContinuationToken string `json:"continuationToken,omitempty"`
}
