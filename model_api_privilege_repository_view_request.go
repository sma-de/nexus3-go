/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiPrivilegeRepositoryViewRequest struct {
	// The name of the privilege.  This value cannot be changed.
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as 'run' for script privileges.
	Actions []string `json:"actions,omitempty"`
	// The repository format (i.e 'nuget', 'npm') this privilege will grant access to (or * for all).
	Format string `json:"format,omitempty"`
	// The name of the repository this privilege will grant access to (or * for all).
	Repository string `json:"repository,omitempty"`
}
