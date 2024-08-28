/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ReadLdapServerXo struct {
	// LDAP server name
	Name string `json:"name"`
	// LDAP server connection Protocol to use
	Protocol string `json:"protocol"`
	// Whether to use certificates stored in Nexus Repository Manager's truststore
	UseTrustStore bool `json:"useTrustStore,omitempty"`
	// LDAP server connection hostname
	Host string `json:"host"`
	// LDAP server connection port to use
	Port int32 `json:"port"`
	// LDAP location to be added to the connection URL
	SearchBase string `json:"searchBase"`
	// Authentication scheme used for connecting to LDAP server
	AuthScheme string `json:"authScheme"`
	// The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	AuthRealm string `json:"authRealm,omitempty"`
	// This must be a fully qualified username if simple authentication is used. Required if authScheme other than none.
	AuthUsername string `json:"authUsername,omitempty"`
	// How long to wait before timeout
	ConnectionTimeoutSeconds int32 `json:"connectionTimeoutSeconds"`
	// How long to wait before retrying
	ConnectionRetryDelaySeconds int32 `json:"connectionRetryDelaySeconds"`
	// How many retry attempts
	MaxIncidentsCount int32 `json:"maxIncidentsCount"`
	// The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN.
	UserBaseDn string `json:"userBaseDn,omitempty"`
	// Are users located in structures below the user base DN?
	UserSubtree bool `json:"userSubtree,omitempty"`
	// LDAP class for user objects
	UserObjectClass string `json:"userObjectClass,omitempty"`
	// LDAP search filter to limit user search
	UserLdapFilter string `json:"userLdapFilter,omitempty"`
	// This is used to find a user given its user ID
	UserIdAttribute string `json:"userIdAttribute,omitempty"`
	// This is used to find a real name given the user ID
	UserRealNameAttribute string `json:"userRealNameAttribute,omitempty"`
	// This is used to find an email address given the user ID
	UserEmailAddressAttribute string `json:"userEmailAddressAttribute,omitempty"`
	// If this field is blank the user will be authenticated against a bind with the LDAP server
	UserPasswordAttribute string `json:"userPasswordAttribute,omitempty"`
	// Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	LdapGroupsAsRoles bool `json:"ldapGroupsAsRoles,omitempty"`
	// Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true.
	GroupType string `json:"groupType"`
	// The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN.
	GroupBaseDn string `json:"groupBaseDn,omitempty"`
	// Are groups located in structures below the group base DN
	GroupSubtree bool `json:"groupSubtree,omitempty"`
	// LDAP class for group objects. Required if groupType is static
	GroupObjectClass string `json:"groupObjectClass,omitempty"`
	// This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	GroupIdAttribute string `json:"groupIdAttribute,omitempty"`
	// LDAP attribute containing the usernames for the group. Required if groupType is static
	GroupMemberAttribute string `json:"groupMemberAttribute,omitempty"`
	// The format of user ID stored in the group member attribute. Required if groupType is static
	GroupMemberFormat string `json:"groupMemberFormat,omitempty"`
	// Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	UserMemberOfAttribute string `json:"userMemberOfAttribute,omitempty"`
	// LDAP server ID
	Id string `json:"id,omitempty"`
	// Order number in which the server is being used when looking for a user
	Order int32 `json:"order,omitempty"`
}
