# ApiPrivilegeApplicationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the privilege.  This value cannot be changed. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Actions** | **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | [optional] [default to null]
**Domain** | **string** | The domain (i.e. &#39;blobstores&#39;, &#39;capabilities&#39; or even &#39;*&#39; for all) that this privilege is granting access to.  Note that creating new privileges with a domain is only necessary when using plugins that define their own domain(s). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


