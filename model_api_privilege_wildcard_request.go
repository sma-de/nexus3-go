/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiPrivilegeWildcardRequest struct {
	// The name of the privilege.  This value cannot be changed.
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// A colon separated list of parts that create a permission string.
	Pattern string `json:"pattern,omitempty"`
}
