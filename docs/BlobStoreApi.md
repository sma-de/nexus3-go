# \BlobStoreApi

All URIs are relative to *https://localhost/service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlobStore**](BlobStoreApi.md#CreateBlobStore) | **Post** /v1/blobstores/s3 | Create an S3 blob store
[**CreateBlobStore1**](BlobStoreApi.md#CreateBlobStore1) | **Post** /v1/blobstores/azure | Create an Azure blob store
[**CreateFileBlobStore**](BlobStoreApi.md#CreateFileBlobStore) | **Post** /v1/blobstores/file | Create a file blob store
[**DeleteBlobStore**](BlobStoreApi.md#DeleteBlobStore) | **Delete** /v1/blobstores/{name} | Delete a blob store by name
[**GetBlobStore**](BlobStoreApi.md#GetBlobStore) | **Get** /v1/blobstores/s3/{name} | Get a S3 blob store configuration by name
[**GetBlobStore1**](BlobStoreApi.md#GetBlobStore1) | **Get** /v1/blobstores/azure/{name} | Get an Azure blob store configuration by name
[**GetFileBlobStoreConfiguration**](BlobStoreApi.md#GetFileBlobStoreConfiguration) | **Get** /v1/blobstores/file/{name} | Get a file blob store configuration by name
[**ListBlobStores**](BlobStoreApi.md#ListBlobStores) | **Get** /v1/blobstores | List the blob stores
[**QuotaStatus**](BlobStoreApi.md#QuotaStatus) | **Get** /v1/blobstores/{name}/quota-status | Get quota status for a given blob store
[**UpdateBlobStore**](BlobStoreApi.md#UpdateBlobStore) | **Put** /v1/blobstores/s3/{name} | Update an S3 blob store configuration by name
[**UpdateBlobStore1**](BlobStoreApi.md#UpdateBlobStore1) | **Put** /v1/blobstores/azure/{name} | Update an Azure blob store configuration by name
[**UpdateFileBlobStore**](BlobStoreApi.md#UpdateFileBlobStore) | **Put** /v1/blobstores/file/{name} | Update a file blob store configuration by name


# **CreateBlobStore**
> CreateBlobStore(ctx, optional)
Create an S3 blob store



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlobStoreApiCreateBlobStoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStoreApiCreateBlobStoreOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of S3BlobStoreApiModel**](S3BlobStoreApiModel.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBlobStore1**
> CreateBlobStore1(ctx, optional)
Create an Azure blob store



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlobStoreApiCreateBlobStore1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStoreApiCreateBlobStore1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AzureBlobStoreApiModel**](AzureBlobStoreApiModel.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFileBlobStore**
> CreateFileBlobStore(ctx, optional)
Create a file blob store



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BlobStoreApiCreateFileBlobStoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStoreApiCreateFileBlobStoreOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FileBlobStoreApiCreateRequest**](FileBlobStoreApiCreateRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBlobStore**
> DeleteBlobStore(ctx, name)
Delete a blob store by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the blob store to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlobStore**
> S3BlobStoreApiModel GetBlobStore(ctx, name)
Get a S3 blob store configuration by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the blob store configuration to fetch | 

### Return type

[**S3BlobStoreApiModel**](S3BlobStoreApiModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlobStore1**
> AzureBlobStoreApiModel GetBlobStore1(ctx, name)
Get an Azure blob store configuration by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the blob store configuration to fetch | 

### Return type

[**AzureBlobStoreApiModel**](AzureBlobStoreApiModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileBlobStoreConfiguration**
> FileBlobStoreApiModel GetFileBlobStoreConfiguration(ctx, name)
Get a file blob store configuration by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file blob store to read | 

### Return type

[**FileBlobStoreApiModel**](FileBlobStoreApiModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBlobStores**
> []GenericBlobStoreApiResponse ListBlobStores(ctx, )
List the blob stores



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GenericBlobStoreApiResponse**](GenericBlobStoreApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotaStatus**
> BlobStoreQuotaResultXo QuotaStatus(ctx, name)
Get quota status for a given blob store



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**BlobStoreQuotaResultXo**](BlobStoreQuotaResultXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBlobStore**
> UpdateBlobStore(ctx, name, optional)
Update an S3 blob store configuration by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the blob store to update | 
 **optional** | ***BlobStoreApiUpdateBlobStoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStoreApiUpdateBlobStoreOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of S3BlobStoreApiModel**](S3BlobStoreApiModel.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBlobStore1**
> UpdateBlobStore1(ctx, name, optional)
Update an Azure blob store configuration by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of the blob store to update | 
 **optional** | ***BlobStoreApiUpdateBlobStore1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStoreApiUpdateBlobStore1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AzureBlobStoreApiModel**](AzureBlobStoreApiModel.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFileBlobStore**
> UpdateFileBlobStore(ctx, name, optional)
Update a file blob store configuration by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file blob store to update | 
 **optional** | ***BlobStoreApiUpdateFileBlobStoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlobStoreApiUpdateFileBlobStoreOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of FileBlobStoreApiUpdateRequest**](FileBlobStoreApiUpdateRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

