/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.69.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiCertificate struct {
	ExpiresOn int64 `json:"expiresOn,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Id string `json:"id,omitempty"`
	IssuedOn int64 `json:"issuedOn,omitempty"`
	IssuerCommonName string `json:"issuerCommonName,omitempty"`
	IssuerOrganization string `json:"issuerOrganization,omitempty"`
	IssuerOrganizationalUnit string `json:"issuerOrganizationalUnit,omitempty"`
	Pem string `json:"pem,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	SubjectCommonName string `json:"subjectCommonName,omitempty"`
	SubjectOrganization string `json:"subjectOrganization,omitempty"`
	SubjectOrganizationalUnit string `json:"subjectOrganizationalUnit,omitempty"`
}
