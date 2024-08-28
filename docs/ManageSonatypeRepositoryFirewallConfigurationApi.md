# \ManageSonatypeRepositoryFirewallConfigurationApi

All URIs are relative to *https://localhost/service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableIq**](ManageSonatypeRepositoryFirewallConfigurationApi.md#DisableIq) | **Post** /v1/iq/disable | Disable Sonatype Repository Firewall
[**EnableIq**](ManageSonatypeRepositoryFirewallConfigurationApi.md#EnableIq) | **Post** /v1/iq/enable | Enable Sonatype Repository Firewall
[**GetConfiguration**](ManageSonatypeRepositoryFirewallConfigurationApi.md#GetConfiguration) | **Get** /v1/iq | Get Sonatype Repository Firewall configuration
[**UpdateConfiguration**](ManageSonatypeRepositoryFirewallConfigurationApi.md#UpdateConfiguration) | **Put** /v1/iq | Update Sonatype Repository Firewall configuration
[**VerifyConnection**](ManageSonatypeRepositoryFirewallConfigurationApi.md#VerifyConnection) | **Post** /v1/iq/verify-connection | Verify Sonatype Repository Firewall connection


# **DisableIq**
> DisableIq(ctx, )
Disable Sonatype Repository Firewall



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

# **EnableIq**
> EnableIq(ctx, )
Enable Sonatype Repository Firewall



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

# **GetConfiguration**
> GetConfiguration(ctx, )
Get Sonatype Repository Firewall configuration



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

# **UpdateConfiguration**
> UpdateConfiguration(ctx, optional)
Update Sonatype Repository Firewall configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManageSonatypeRepositoryFirewallConfigurationApiUpdateConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManageSonatypeRepositoryFirewallConfigurationApiUpdateConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of IqConnectionXo**](IqConnectionXo.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyConnection**
> VerifyConnection(ctx, )
Verify Sonatype Repository Firewall connection



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

