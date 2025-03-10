/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RoutingRuleXo struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// Determines what should be done with requests when their path matches any of the matchers
	Mode string `json:"mode,omitempty"`
	// Regular expressions used to identify request paths that are allowed or blocked (depending on mode)
	Matchers []string `json:"matchers,omitempty"`
}
