# \ReadOnlyApi

All URIs are relative to *https://localhost/service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForceRelease**](ReadOnlyApi.md#ForceRelease) | **Post** /v1/read-only/force-release | Forcibly release read-only and allow changes to embedded OrientDB
[**Freeze**](ReadOnlyApi.md#Freeze) | **Post** /v1/read-only/freeze | Prevent changes to embedded OrientDB
[**Get**](ReadOnlyApi.md#Get) | **Get** /v1/read-only | Get read-only state
[**Release**](ReadOnlyApi.md#Release) | **Post** /v1/read-only/release | Release read-only and allow changes to embedded OrientDB


# **ForceRelease**
> ForceRelease(ctx, )
Forcibly release read-only and allow changes to embedded OrientDB

Forcibly release read-only status, including if caused by system tasks. Warning: may result in data loss.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Freeze**
> Freeze(ctx, )
Prevent changes to embedded OrientDB

For low-level system maintenance purposes only; do not use if you want users to still be able to download components.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> ReadOnlyState Get(ctx, )
Get read-only state



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ReadOnlyState**](ReadOnlyState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Release**
> Release(ctx, )
Release read-only and allow changes to embedded OrientDB

Releases administrator-initiated read-only status. Will not release read-only status caused by system tasks.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

