# \ComponentsApi

All URIs are relative to *https://localhost/service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComponent**](ComponentsApi.md#DeleteComponent) | **Delete** /v1/components/{id} | Delete a single component
[**GetComponentById**](ComponentsApi.md#GetComponentById) | **Get** /v1/components/{id} | Get a single component
[**GetComponents**](ComponentsApi.md#GetComponents) | **Get** /v1/components | List components
[**UploadComponent**](ComponentsApi.md#UploadComponent) | **Post** /v1/components | Upload a single component


# **DeleteComponent**
> DeleteComponent(ctx, id)
Delete a single component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the component to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentById**
> ComponentXo GetComponentById(ctx, id)
Get a single component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the component to retrieve | 

### Return type

[**ComponentXo**](ComponentXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponents**
> PageComponentXo GetComponents(ctx, repository, optional)
List components



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| Repository from which you would like to retrieve components | 
 **optional** | ***ComponentsApiGetComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentsApiGetComponentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| A token returned by a prior request. If present, the next page of results are returned | 

### Return type

[**PageComponentXo**](PageComponentXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadComponent**
> UploadComponent(ctx, repository, optional)
Upload a single component



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| Name of the repository to which you would like to upload the component | 
 **optional** | ***ComponentsApiUploadComponentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentsApiUploadComponentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rAsset** | **optional.Interface of *os.File**| r Asset  | 
 **rAssetPathId** | **optional.String**| r Asset  Package Path | 
 **pypiAsset** | **optional.Interface of *os.File**| pypi Asset  | 
 **helmAsset** | **optional.Interface of *os.File**| helm Asset  | 
 **yumDirectory** | **optional.String**| yum Directory | 
 **yumAsset** | **optional.Interface of *os.File**| yum Asset  | 
 **yumAssetFilename** | **optional.String**| yum Asset  Filename | 
 **dockerAsset** | **optional.Interface of *os.File**| docker Asset  | 
 **rubygemsAsset** | **optional.Interface of *os.File**| rubygems Asset  | 
 **nugetAsset** | **optional.Interface of *os.File**| nuget Asset  | 
 **npmAsset** | **optional.Interface of *os.File**| npm Asset  | 
 **rawDirectory** | **optional.String**| raw Directory | 
 **rawAsset1** | **optional.Interface of *os.File**| raw Asset 1 | 
 **rawAsset1Filename** | **optional.String**| raw Asset 1 Filename | 
 **rawAsset2** | **optional.Interface of *os.File**| raw Asset 2 | 
 **rawAsset2Filename** | **optional.String**| raw Asset 2 Filename | 
 **rawAsset3** | **optional.Interface of *os.File**| raw Asset 3 | 
 **rawAsset3Filename** | **optional.String**| raw Asset 3 Filename | 
 **aptAsset** | **optional.Interface of *os.File**| apt Asset  | 
 **maven2GroupId** | **optional.String**| maven2 Group ID | 
 **maven2ArtifactId** | **optional.String**| maven2 Artifact ID | 
 **maven2Version** | **optional.String**| maven2 Version | 
 **maven2GeneratePom** | **optional.Bool**| maven2 Generate a POM file with these coordinates | 
 **maven2Packaging** | **optional.String**| maven2 Packaging | 
 **maven2Asset1** | **optional.Interface of *os.File**| maven2 Asset 1 | 
 **maven2Asset1Classifier** | **optional.String**| maven2 Asset 1 Classifier | 
 **maven2Asset1Extension** | **optional.String**| maven2 Asset 1 Extension | 
 **maven2Asset2** | **optional.Interface of *os.File**| maven2 Asset 2 | 
 **maven2Asset2Classifier** | **optional.String**| maven2 Asset 2 Classifier | 
 **maven2Asset2Extension** | **optional.String**| maven2 Asset 2 Extension | 
 **maven2Asset3** | **optional.Interface of *os.File**| maven2 Asset 3 | 
 **maven2Asset3Classifier** | **optional.String**| maven2 Asset 3 Classifier | 
 **maven2Asset3Extension** | **optional.String**| maven2 Asset 3 Extension | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

