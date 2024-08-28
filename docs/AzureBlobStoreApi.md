# \AzureBlobStoreApi

All URIs are relative to *https://localhost/service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyConnection1**](AzureBlobStoreApi.md#VerifyConnection1) | **Post** /v1/azureblobstore/test-connection | Verify connection using supplied Azure Blob Store settings


# **VerifyConnection1**
> VerifyConnection1(ctx, optional)
Verify connection using supplied Azure Blob Store settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AzureBlobStoreApiVerifyConnection1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AzureBlobStoreApiVerifyConnection1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AzureConnectionXo**](AzureConnectionXo.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

