# DockerHostedStorageAttributes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobStoreName** | **string** | Blob store used to store repository contents | [default to null]
**StrictContentTypeValidation** | **bool** | Whether to validate uploaded content&#39;s MIME type appropriate for the repository format | [default to null]
**WritePolicy** | **string** | Controls if deployments of and updates to assets are allowed | [default to null]
**LatestPolicy** | **bool** | Whether to allow redeploying the &#39;latest&#39; tag but defer to the Deployment Policy for all other tags | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


