/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HttpClientConnectionAuthenticationAttributes struct {
	// Authentication type
	Type_ string `json:"type,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	NtlmHost string `json:"ntlmHost,omitempty"`
	NtlmDomain string `json:"ntlmDomain,omitempty"`
}
