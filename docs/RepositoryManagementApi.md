# \RepositoryManagementApi

All URIs are relative to *https://localhost/service/rest/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRepository**](RepositoryManagementApi.md#CreateRepository) | **Post** /v1/repositories/maven/group | Create Maven group repository
[**CreateRepository1**](RepositoryManagementApi.md#CreateRepository1) | **Post** /v1/repositories/maven/hosted | Create Maven hosted repository
[**CreateRepository10**](RepositoryManagementApi.md#CreateRepository10) | **Post** /v1/repositories/npm/proxy | Create npm proxy repository
[**CreateRepository11**](RepositoryManagementApi.md#CreateRepository11) | **Post** /v1/repositories/nuget/group | Create NuGet group repository
[**CreateRepository12**](RepositoryManagementApi.md#CreateRepository12) | **Post** /v1/repositories/nuget/hosted | Create NuGet hosted repository
[**CreateRepository13**](RepositoryManagementApi.md#CreateRepository13) | **Post** /v1/repositories/nuget/proxy | Create NuGet proxy repository
[**CreateRepository14**](RepositoryManagementApi.md#CreateRepository14) | **Post** /v1/repositories/rubygems/group | Create RubyGems group repository
[**CreateRepository15**](RepositoryManagementApi.md#CreateRepository15) | **Post** /v1/repositories/rubygems/hosted | Create RubyGems hosted repository
[**CreateRepository16**](RepositoryManagementApi.md#CreateRepository16) | **Post** /v1/repositories/rubygems/proxy | Create RubyGems proxy repository
[**CreateRepository17**](RepositoryManagementApi.md#CreateRepository17) | **Post** /v1/repositories/docker/group | Create Docker group repository
[**CreateRepository18**](RepositoryManagementApi.md#CreateRepository18) | **Post** /v1/repositories/docker/hosted | Create Docker hosted repository
[**CreateRepository19**](RepositoryManagementApi.md#CreateRepository19) | **Post** /v1/repositories/docker/proxy | Create Docker proxy repository
[**CreateRepository2**](RepositoryManagementApi.md#CreateRepository2) | **Post** /v1/repositories/maven/proxy | Create Maven proxy repository
[**CreateRepository20**](RepositoryManagementApi.md#CreateRepository20) | **Post** /v1/repositories/yum/group | Create Yum group repository
[**CreateRepository21**](RepositoryManagementApi.md#CreateRepository21) | **Post** /v1/repositories/yum/hosted | Create Yum hosted repository
[**CreateRepository22**](RepositoryManagementApi.md#CreateRepository22) | **Post** /v1/repositories/yum/proxy | Create Yum proxy repository
[**CreateRepository23**](RepositoryManagementApi.md#CreateRepository23) | **Post** /v1/repositories/helm/hosted | Create Helm hosted repository
[**CreateRepository24**](RepositoryManagementApi.md#CreateRepository24) | **Post** /v1/repositories/helm/proxy | Create Helm proxy repository
[**CreateRepository25**](RepositoryManagementApi.md#CreateRepository25) | **Post** /v1/repositories/gitlfs/hosted | Create Git LFS hosted repository
[**CreateRepository26**](RepositoryManagementApi.md#CreateRepository26) | **Post** /v1/repositories/pypi/group | Create PyPI group repository
[**CreateRepository27**](RepositoryManagementApi.md#CreateRepository27) | **Post** /v1/repositories/pypi/hosted | Create PyPI hosted repository
[**CreateRepository28**](RepositoryManagementApi.md#CreateRepository28) | **Post** /v1/repositories/pypi/proxy | Create PyPI proxy repository
[**CreateRepository29**](RepositoryManagementApi.md#CreateRepository29) | **Post** /v1/repositories/conda/proxy | Create conda proxy repository
[**CreateRepository3**](RepositoryManagementApi.md#CreateRepository3) | **Post** /v1/repositories/apt/hosted | Create APT hosted repository
[**CreateRepository30**](RepositoryManagementApi.md#CreateRepository30) | **Post** /v1/repositories/conan/proxy | Create Conan proxy repository
[**CreateRepository31**](RepositoryManagementApi.md#CreateRepository31) | **Post** /v1/repositories/r/group | Create R group repository
[**CreateRepository32**](RepositoryManagementApi.md#CreateRepository32) | **Post** /v1/repositories/r/hosted | Create R hosted repository
[**CreateRepository33**](RepositoryManagementApi.md#CreateRepository33) | **Post** /v1/repositories/r/proxy | Create R proxy repository
[**CreateRepository34**](RepositoryManagementApi.md#CreateRepository34) | **Post** /v1/repositories/cocoapods/proxy | Create Cocoapods proxy repository
[**CreateRepository35**](RepositoryManagementApi.md#CreateRepository35) | **Post** /v1/repositories/go/group | Create a Go group repository
[**CreateRepository36**](RepositoryManagementApi.md#CreateRepository36) | **Post** /v1/repositories/go/proxy | Create a Go proxy repository
[**CreateRepository37**](RepositoryManagementApi.md#CreateRepository37) | **Post** /v1/repositories/p2/proxy | Create p2 proxy repository
[**CreateRepository38**](RepositoryManagementApi.md#CreateRepository38) | **Post** /v1/repositories/bower/group | Create Bower group repository
[**CreateRepository39**](RepositoryManagementApi.md#CreateRepository39) | **Post** /v1/repositories/bower/hosted | Create Bower hosted repository
[**CreateRepository4**](RepositoryManagementApi.md#CreateRepository4) | **Post** /v1/repositories/apt/proxy | Create APT proxy repository
[**CreateRepository40**](RepositoryManagementApi.md#CreateRepository40) | **Post** /v1/repositories/bower/proxy | Create Bower proxy repository
[**CreateRepository5**](RepositoryManagementApi.md#CreateRepository5) | **Post** /v1/repositories/raw/group | Create raw group repository
[**CreateRepository6**](RepositoryManagementApi.md#CreateRepository6) | **Post** /v1/repositories/raw/hosted | Create raw hosted repository
[**CreateRepository7**](RepositoryManagementApi.md#CreateRepository7) | **Post** /v1/repositories/raw/proxy | Create raw proxy repository
[**CreateRepository8**](RepositoryManagementApi.md#CreateRepository8) | **Post** /v1/repositories/npm/group | Create npm group repository
[**CreateRepository9**](RepositoryManagementApi.md#CreateRepository9) | **Post** /v1/repositories/npm/hosted | Create npm hosted repository
[**DeleteRepository**](RepositoryManagementApi.md#DeleteRepository) | **Delete** /v1/repositories/{repositoryName} | Delete repository of any format
[**DisableRepositoryHealthCheck**](RepositoryManagementApi.md#DisableRepositoryHealthCheck) | **Delete** /v1/repositories/{repositoryName}/health-check | Disable repository health check. Proxy repositories only.
[**EnableRepositoryHealthCheck**](RepositoryManagementApi.md#EnableRepositoryHealthCheck) | **Post** /v1/repositories/{repositoryName}/health-check | Enable repository health check. Proxy repositories only.
[**GetRepositories**](RepositoryManagementApi.md#GetRepositories) | **Get** /v1/repositorySettings | List repositories
[**GetRepositories1**](RepositoryManagementApi.md#GetRepositories1) | **Get** /v1/repositories | List repositories
[**GetRepository**](RepositoryManagementApi.md#GetRepository) | **Get** /v1/repositories/{repositoryName} | Get repository details
[**GetRepository1**](RepositoryManagementApi.md#GetRepository1) | **Get** /v1/repositories/maven/group/{repositoryName} | Get repository
[**GetRepository10**](RepositoryManagementApi.md#GetRepository10) | **Get** /v1/repositories/npm/hosted/{repositoryName} | Get repository
[**GetRepository11**](RepositoryManagementApi.md#GetRepository11) | **Get** /v1/repositories/npm/proxy/{repositoryName} | Get repository
[**GetRepository12**](RepositoryManagementApi.md#GetRepository12) | **Get** /v1/repositories/nuget/group/{repositoryName} | Get repository
[**GetRepository13**](RepositoryManagementApi.md#GetRepository13) | **Get** /v1/repositories/nuget/hosted/{repositoryName} | Get repository
[**GetRepository14**](RepositoryManagementApi.md#GetRepository14) | **Get** /v1/repositories/nuget/proxy/{repositoryName} | Get repository
[**GetRepository15**](RepositoryManagementApi.md#GetRepository15) | **Get** /v1/repositories/rubygems/group/{repositoryName} | Get repository
[**GetRepository16**](RepositoryManagementApi.md#GetRepository16) | **Get** /v1/repositories/rubygems/hosted/{repositoryName} | Get repository
[**GetRepository17**](RepositoryManagementApi.md#GetRepository17) | **Get** /v1/repositories/rubygems/proxy/{repositoryName} | Get repository
[**GetRepository18**](RepositoryManagementApi.md#GetRepository18) | **Get** /v1/repositories/docker/group/{repositoryName} | Get repository
[**GetRepository19**](RepositoryManagementApi.md#GetRepository19) | **Get** /v1/repositories/docker/hosted/{repositoryName} | Get repository
[**GetRepository2**](RepositoryManagementApi.md#GetRepository2) | **Get** /v1/repositories/maven/hosted/{repositoryName} | Get repository
[**GetRepository20**](RepositoryManagementApi.md#GetRepository20) | **Get** /v1/repositories/docker/proxy/{repositoryName} | Get repository
[**GetRepository21**](RepositoryManagementApi.md#GetRepository21) | **Get** /v1/repositories/yum/group/{repositoryName} | Get repository
[**GetRepository22**](RepositoryManagementApi.md#GetRepository22) | **Get** /v1/repositories/yum/hosted/{repositoryName} | Get repository
[**GetRepository23**](RepositoryManagementApi.md#GetRepository23) | **Get** /v1/repositories/yum/proxy/{repositoryName} | Get repository
[**GetRepository24**](RepositoryManagementApi.md#GetRepository24) | **Get** /v1/repositories/helm/hosted/{repositoryName} | Get repository
[**GetRepository25**](RepositoryManagementApi.md#GetRepository25) | **Get** /v1/repositories/helm/proxy/{repositoryName} | Get repository
[**GetRepository26**](RepositoryManagementApi.md#GetRepository26) | **Get** /v1/repositories/gitlfs/hosted/{repositoryName} | Get repository
[**GetRepository27**](RepositoryManagementApi.md#GetRepository27) | **Get** /v1/repositories/pypi/group/{repositoryName} | Get repository
[**GetRepository28**](RepositoryManagementApi.md#GetRepository28) | **Get** /v1/repositories/pypi/hosted/{repositoryName} | Get repository
[**GetRepository29**](RepositoryManagementApi.md#GetRepository29) | **Get** /v1/repositories/pypi/proxy/{repositoryName} | Get repository
[**GetRepository3**](RepositoryManagementApi.md#GetRepository3) | **Get** /v1/repositories/maven/proxy/{repositoryName} | Get repository
[**GetRepository30**](RepositoryManagementApi.md#GetRepository30) | **Get** /v1/repositories/conda/proxy/{repositoryName} | Get repository
[**GetRepository31**](RepositoryManagementApi.md#GetRepository31) | **Get** /v1/repositories/conan/proxy/{repositoryName} | Get repository
[**GetRepository32**](RepositoryManagementApi.md#GetRepository32) | **Get** /v1/repositories/r/group/{repositoryName} | Get repository
[**GetRepository33**](RepositoryManagementApi.md#GetRepository33) | **Get** /v1/repositories/r/hosted/{repositoryName} | Get repository
[**GetRepository34**](RepositoryManagementApi.md#GetRepository34) | **Get** /v1/repositories/r/proxy/{repositoryName} | Get repository
[**GetRepository35**](RepositoryManagementApi.md#GetRepository35) | **Get** /v1/repositories/cocoapods/proxy/{repositoryName} | Get repository
[**GetRepository36**](RepositoryManagementApi.md#GetRepository36) | **Get** /v1/repositories/go/group/{repositoryName} | Get repository
[**GetRepository37**](RepositoryManagementApi.md#GetRepository37) | **Get** /v1/repositories/go/proxy/{repositoryName} | Get repository
[**GetRepository38**](RepositoryManagementApi.md#GetRepository38) | **Get** /v1/repositories/p2/proxy/{repositoryName} | Get repository
[**GetRepository39**](RepositoryManagementApi.md#GetRepository39) | **Get** /v1/repositories/bower/group/{repositoryName} | Get repository
[**GetRepository4**](RepositoryManagementApi.md#GetRepository4) | **Get** /v1/repositories/apt/hosted/{repositoryName} | Get repository
[**GetRepository40**](RepositoryManagementApi.md#GetRepository40) | **Get** /v1/repositories/bower/hosted/{repositoryName} | Get repository
[**GetRepository41**](RepositoryManagementApi.md#GetRepository41) | **Get** /v1/repositories/bower/proxy/{repositoryName} | Get repository
[**GetRepository5**](RepositoryManagementApi.md#GetRepository5) | **Get** /v1/repositories/apt/proxy/{repositoryName} | Get repository
[**GetRepository6**](RepositoryManagementApi.md#GetRepository6) | **Get** /v1/repositories/raw/group/{repositoryName} | Get repository
[**GetRepository7**](RepositoryManagementApi.md#GetRepository7) | **Get** /v1/repositories/raw/hosted/{repositoryName} | Get repository
[**GetRepository8**](RepositoryManagementApi.md#GetRepository8) | **Get** /v1/repositories/raw/proxy/{repositoryName} | Get repository
[**GetRepository9**](RepositoryManagementApi.md#GetRepository9) | **Get** /v1/repositories/npm/group/{repositoryName} | Get repository
[**InvalidateCache**](RepositoryManagementApi.md#InvalidateCache) | **Post** /v1/repositories/{repositoryName}/invalidate-cache | Invalidate repository cache. Proxy or group repositories only.
[**RebuildIndex**](RepositoryManagementApi.md#RebuildIndex) | **Post** /v1/repositories/{repositoryName}/rebuild-index | Schedule a &#39;Repair - Rebuild repository search&#39; Task. Hosted or proxy repositories only.
[**UpdateRepository**](RepositoryManagementApi.md#UpdateRepository) | **Put** /v1/repositories/maven/group/{repositoryName} | Update Maven group repository
[**UpdateRepository1**](RepositoryManagementApi.md#UpdateRepository1) | **Put** /v1/repositories/maven/hosted/{repositoryName} | Update Maven hosted repository
[**UpdateRepository10**](RepositoryManagementApi.md#UpdateRepository10) | **Put** /v1/repositories/npm/proxy/{repositoryName} | Update npm proxy repository
[**UpdateRepository11**](RepositoryManagementApi.md#UpdateRepository11) | **Put** /v1/repositories/nuget/group/{repositoryName} | Update NuGet group repository
[**UpdateRepository12**](RepositoryManagementApi.md#UpdateRepository12) | **Put** /v1/repositories/nuget/hosted/{repositoryName} | Update NuGet hosted repository
[**UpdateRepository13**](RepositoryManagementApi.md#UpdateRepository13) | **Put** /v1/repositories/nuget/proxy/{repositoryName} | Update NuGet proxy repository
[**UpdateRepository14**](RepositoryManagementApi.md#UpdateRepository14) | **Put** /v1/repositories/rubygems/group/{repositoryName} | Update RubyGems group repository
[**UpdateRepository15**](RepositoryManagementApi.md#UpdateRepository15) | **Put** /v1/repositories/rubygems/hosted/{repositoryName} | Update RubyGems hosted repository
[**UpdateRepository16**](RepositoryManagementApi.md#UpdateRepository16) | **Put** /v1/repositories/rubygems/proxy/{repositoryName} | Update RubyGems proxy repository
[**UpdateRepository17**](RepositoryManagementApi.md#UpdateRepository17) | **Put** /v1/repositories/docker/group/{repositoryName} | Update Docker group repository
[**UpdateRepository18**](RepositoryManagementApi.md#UpdateRepository18) | **Put** /v1/repositories/docker/hosted/{repositoryName} | Update Docker hosted repository
[**UpdateRepository19**](RepositoryManagementApi.md#UpdateRepository19) | **Put** /v1/repositories/docker/proxy/{repositoryName} | Update Docker proxy repository
[**UpdateRepository2**](RepositoryManagementApi.md#UpdateRepository2) | **Put** /v1/repositories/maven/proxy/{repositoryName} | Update Maven proxy repository
[**UpdateRepository20**](RepositoryManagementApi.md#UpdateRepository20) | **Put** /v1/repositories/yum/group/{repositoryName} | Update Yum group repository
[**UpdateRepository21**](RepositoryManagementApi.md#UpdateRepository21) | **Put** /v1/repositories/yum/hosted/{repositoryName} | Update Yum hosted repository
[**UpdateRepository22**](RepositoryManagementApi.md#UpdateRepository22) | **Put** /v1/repositories/yum/proxy/{repositoryName} | Update Yum proxy repository
[**UpdateRepository23**](RepositoryManagementApi.md#UpdateRepository23) | **Put** /v1/repositories/helm/hosted/{repositoryName} | Update Helm hosted repository
[**UpdateRepository24**](RepositoryManagementApi.md#UpdateRepository24) | **Put** /v1/repositories/helm/proxy/{repositoryName} | Update Helm proxy repository
[**UpdateRepository25**](RepositoryManagementApi.md#UpdateRepository25) | **Put** /v1/repositories/gitlfs/hosted/{repositoryName} | Update Git LFS hosted repository
[**UpdateRepository26**](RepositoryManagementApi.md#UpdateRepository26) | **Put** /v1/repositories/pypi/group/{repositoryName} | Update PyPI group repository
[**UpdateRepository27**](RepositoryManagementApi.md#UpdateRepository27) | **Put** /v1/repositories/pypi/hosted/{repositoryName} | Update PyPI hosted repository
[**UpdateRepository28**](RepositoryManagementApi.md#UpdateRepository28) | **Put** /v1/repositories/pypi/proxy/{repositoryName} | Update PyPI proxy repository
[**UpdateRepository29**](RepositoryManagementApi.md#UpdateRepository29) | **Put** /v1/repositories/conda/proxy/{repositoryName} | Update conda proxy repository
[**UpdateRepository3**](RepositoryManagementApi.md#UpdateRepository3) | **Put** /v1/repositories/apt/hosted/{repositoryName} | Update APT hosted repository
[**UpdateRepository30**](RepositoryManagementApi.md#UpdateRepository30) | **Put** /v1/repositories/conan/proxy/{repositoryName} | Update Conan proxy repository
[**UpdateRepository31**](RepositoryManagementApi.md#UpdateRepository31) | **Put** /v1/repositories/r/group/{repositoryName} | Update R group repository
[**UpdateRepository32**](RepositoryManagementApi.md#UpdateRepository32) | **Put** /v1/repositories/r/hosted/{repositoryName} | Update R hosted repository
[**UpdateRepository33**](RepositoryManagementApi.md#UpdateRepository33) | **Put** /v1/repositories/r/proxy/{repositoryName} | Update R proxy repository
[**UpdateRepository34**](RepositoryManagementApi.md#UpdateRepository34) | **Put** /v1/repositories/cocoapods/proxy/{repositoryName} | Update Cocoapods proxy repository
[**UpdateRepository35**](RepositoryManagementApi.md#UpdateRepository35) | **Put** /v1/repositories/go/group/{repositoryName} | Update a Go group repository
[**UpdateRepository36**](RepositoryManagementApi.md#UpdateRepository36) | **Put** /v1/repositories/go/proxy/{repositoryName} | Update a Go proxy repository
[**UpdateRepository37**](RepositoryManagementApi.md#UpdateRepository37) | **Put** /v1/repositories/p2/proxy/{repositoryName} | Update p2 proxy repository
[**UpdateRepository38**](RepositoryManagementApi.md#UpdateRepository38) | **Put** /v1/repositories/bower/group/{repositoryName} | Update Bower group repository
[**UpdateRepository39**](RepositoryManagementApi.md#UpdateRepository39) | **Put** /v1/repositories/bower/hosted/{repositoryName} | Update Bower hosted repository
[**UpdateRepository4**](RepositoryManagementApi.md#UpdateRepository4) | **Put** /v1/repositories/apt/proxy/{repositoryName} | Update APT proxy repository
[**UpdateRepository40**](RepositoryManagementApi.md#UpdateRepository40) | **Put** /v1/repositories/bower/proxy/{repositoryName} | Update Bower proxy repository
[**UpdateRepository5**](RepositoryManagementApi.md#UpdateRepository5) | **Put** /v1/repositories/raw/group/{repositoryName} | Update raw group repository
[**UpdateRepository6**](RepositoryManagementApi.md#UpdateRepository6) | **Put** /v1/repositories/raw/hosted/{repositoryName} | Update raw hosted repository
[**UpdateRepository7**](RepositoryManagementApi.md#UpdateRepository7) | **Put** /v1/repositories/raw/proxy/{repositoryName} | Update raw proxy repository
[**UpdateRepository8**](RepositoryManagementApi.md#UpdateRepository8) | **Put** /v1/repositories/npm/group/{repositoryName} | Update npm group repository
[**UpdateRepository9**](RepositoryManagementApi.md#UpdateRepository9) | **Put** /v1/repositories/npm/hosted/{repositoryName} | Update npm hosted repository


# **CreateRepository**
> CreateRepository(ctx, optional)
Create Maven group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepositoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MavenGroupRepositoryApiRequest**](MavenGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository1**
> CreateRepository1(ctx, optional)
Create Maven hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MavenHostedRepositoryApiRequest**](MavenHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository10**
> CreateRepository10(ctx, optional)
Create npm proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository10Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository10Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NpmProxyRepositoryApiRequest**](NpmProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository11**
> CreateRepository11(ctx, optional)
Create NuGet group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository11Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NugetGroupRepositoryApiRequest**](NugetGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository12**
> CreateRepository12(ctx, optional)
Create NuGet hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository12Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository12Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NugetHostedRepositoryApiRequest**](NugetHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository13**
> CreateRepository13(ctx, optional)
Create NuGet proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository13Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository13Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NugetProxyRepositoryApiRequest**](NugetProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository14**
> CreateRepository14(ctx, optional)
Create RubyGems group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository14Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository14Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RubyGemsGroupRepositoryApiRequest**](RubyGemsGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository15**
> CreateRepository15(ctx, optional)
Create RubyGems hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository15Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository15Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RubyGemsHostedRepositoryApiRequest**](RubyGemsHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository16**
> CreateRepository16(ctx, optional)
Create RubyGems proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository16Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository16Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RubyGemsProxyRepositoryApiRequest**](RubyGemsProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository17**
> CreateRepository17(ctx, optional)
Create Docker group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository17Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository17Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DockerGroupRepositoryApiRequest**](DockerGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository18**
> CreateRepository18(ctx, optional)
Create Docker hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository18Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository18Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DockerHostedRepositoryApiRequest**](DockerHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository19**
> CreateRepository19(ctx, optional)
Create Docker proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository19Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository19Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DockerProxyRepositoryApiRequest**](DockerProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository2**
> CreateRepository2(ctx, optional)
Create Maven proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MavenProxyRepositoryApiRequest**](MavenProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository20**
> CreateRepository20(ctx, optional)
Create Yum group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository20Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository20Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of YumGroupRepositoryApiRequest**](YumGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository21**
> CreateRepository21(ctx, optional)
Create Yum hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository21Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository21Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of YumHostedRepositoryApiRequest**](YumHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository22**
> CreateRepository22(ctx, optional)
Create Yum proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository22Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository22Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of YumProxyRepositoryApiRequest**](YumProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository23**
> CreateRepository23(ctx, optional)
Create Helm hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository23Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository23Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HelmHostedRepositoryApiRequest**](HelmHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository24**
> CreateRepository24(ctx, optional)
Create Helm proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository24Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository24Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HelmProxyRepositoryApiRequest**](HelmProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository25**
> CreateRepository25(ctx, optional)
Create Git LFS hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository25Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository25Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GitLfsHostedRepositoryApiRequest**](GitLfsHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository26**
> CreateRepository26(ctx, optional)
Create PyPI group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository26Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository26Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PypiGroupRepositoryApiRequest**](PypiGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository27**
> CreateRepository27(ctx, optional)
Create PyPI hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository27Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository27Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PypiHostedRepositoryApiRequest**](PypiHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository28**
> CreateRepository28(ctx, optional)
Create PyPI proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository28Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository28Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PypiProxyRepositoryApiRequest**](PypiProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository29**
> CreateRepository29(ctx, optional)
Create conda proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository29Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository29Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CondaProxyRepositoryApiRequest**](CondaProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository3**
> CreateRepository3(ctx, optional)
Create APT hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AptHostedRepositoryApiRequest**](AptHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository30**
> CreateRepository30(ctx, optional)
Create Conan proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository30Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository30Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ConanProxyRepositoryApiRequest**](ConanProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository31**
> CreateRepository31(ctx, optional)
Create R group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository31Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository31Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RGroupRepositoryApiRequest**](RGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository32**
> CreateRepository32(ctx, optional)
Create R hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository32Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository32Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RHostedRepositoryApiRequest**](RHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository33**
> CreateRepository33(ctx, optional)
Create R proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository33Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository33Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RProxyRepositoryApiRequest**](RProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository34**
> CreateRepository34(ctx, optional)
Create Cocoapods proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository34Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository34Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CocoapodsProxyRepositoryApiRequest**](CocoapodsProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository35**
> CreateRepository35(ctx, optional)
Create a Go group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository35Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository35Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GolangGroupRepositoryApiRequest**](GolangGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository36**
> CreateRepository36(ctx, optional)
Create a Go proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository36Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository36Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GolangProxyRepositoryApiRequest**](GolangProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository37**
> CreateRepository37(ctx, optional)
Create p2 proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository37Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository37Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of P2ProxyRepositoryApiRequest**](P2ProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository38**
> CreateRepository38(ctx, optional)
Create Bower group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository38Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository38Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BowerGroupRepositoryApiRequest**](BowerGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository39**
> CreateRepository39(ctx, optional)
Create Bower hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository39Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository39Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BowerHostedRepositoryApiRequest**](BowerHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository4**
> CreateRepository4(ctx, optional)
Create APT proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AptProxyRepositoryApiRequest**](AptProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository40**
> CreateRepository40(ctx, optional)
Create Bower proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository40Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository40Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BowerProxyRepositoryApiRequest**](BowerProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository5**
> CreateRepository5(ctx, optional)
Create raw group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository5Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RawGroupRepositoryApiRequest**](RawGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository6**
> CreateRepository6(ctx, optional)
Create raw hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository6Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RawHostedRepositoryApiRequest**](RawHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository7**
> CreateRepository7(ctx, optional)
Create raw proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository7Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository7Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RawProxyRepositoryApiRequest**](RawProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository8**
> CreateRepository8(ctx, optional)
Create npm group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository8Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository8Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NpmGroupRepositoryApiRequest**](NpmGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepository9**
> CreateRepository9(ctx, optional)
Create npm hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryManagementApiCreateRepository9Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiCreateRepository9Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NpmHostedRepositoryApiRequest**](NpmHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepository**
> DeleteRepository(ctx, repositoryName)
Delete repository of any format



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableRepositoryHealthCheck**
> DisableRepositoryHealthCheck(ctx, repositoryName)
Disable repository health check. Proxy repositories only.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to disable Repository Health Check for | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableRepositoryHealthCheck**
> EnableRepositoryHealthCheck(ctx, repositoryName)
Enable repository health check. Proxy repositories only.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to enable Repository Health Check for | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositories**
> []AbstractApiRepository GetRepositories(ctx, )
List repositories



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AbstractApiRepository**](AbstractApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositories1**
> []RepositoryXo GetRepositories1(ctx, )
List repositories



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]RepositoryXo**](RepositoryXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository**
> RepositoryXo GetRepository(ctx, repositoryName)
Get repository details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to get | 

### Return type

[**RepositoryXo**](RepositoryXO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository1**
> SimpleApiGroupRepository GetRepository1(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository10**
> SimpleApiHostedRepository GetRepository10(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository11**
> NpmProxyApiRepository GetRepository11(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**NpmProxyApiRepository**](NpmProxyApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository12**
> SimpleApiGroupRepository GetRepository12(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository13**
> SimpleApiHostedRepository GetRepository13(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository14**
> NugetProxyApiRepository GetRepository14(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**NugetProxyApiRepository**](NugetProxyApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository15**
> SimpleApiGroupRepository GetRepository15(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository16**
> SimpleApiHostedRepository GetRepository16(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository17**
> SimpleApiProxyRepository GetRepository17(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository18**
> DockerGroupApiRepository GetRepository18(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**DockerGroupApiRepository**](DockerGroupApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository19**
> DockerHostedApiRepository GetRepository19(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**DockerHostedApiRepository**](DockerHostedApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository2**
> MavenHostedApiRepository GetRepository2(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**MavenHostedApiRepository**](MavenHostedApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository20**
> DockerProxyApiRepository GetRepository20(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**DockerProxyApiRepository**](DockerProxyApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository21**
> SimpleApiGroupRepository GetRepository21(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository22**
> YumHostedApiRepository GetRepository22(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**YumHostedApiRepository**](YumHostedApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository23**
> SimpleApiProxyRepository GetRepository23(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository24**
> SimpleApiHostedRepository GetRepository24(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository25**
> SimpleApiProxyRepository GetRepository25(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository26**
> SimpleApiHostedRepository GetRepository26(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository27**
> SimpleApiGroupRepository GetRepository27(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository28**
> SimpleApiHostedRepository GetRepository28(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository29**
> SimpleApiProxyRepository GetRepository29(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository3**
> MavenProxyApiRepository GetRepository3(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**MavenProxyApiRepository**](MavenProxyApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository30**
> SimpleApiProxyRepository GetRepository30(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository31**
> SimpleApiProxyRepository GetRepository31(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository32**
> SimpleApiGroupRepository GetRepository32(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository33**
> SimpleApiHostedRepository GetRepository33(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository34**
> SimpleApiProxyRepository GetRepository34(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository35**
> SimpleApiProxyRepository GetRepository35(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository36**
> SimpleApiGroupRepository GetRepository36(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository37**
> SimpleApiProxyRepository GetRepository37(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository38**
> SimpleApiProxyRepository GetRepository38(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository39**
> SimpleApiGroupRepository GetRepository39(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository4**
> AptHostedApiRepository GetRepository4(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**AptHostedApiRepository**](AptHostedApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository40**
> SimpleApiHostedRepository GetRepository40(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository41**
> BowerProxyApiRepository GetRepository41(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**BowerProxyApiRepository**](BowerProxyApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository5**
> AptProxyApiRepository GetRepository5(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**AptProxyApiRepository**](AptProxyApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository6**
> SimpleApiGroupRepository GetRepository6(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository7**
> SimpleApiHostedRepository GetRepository7(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository8**
> SimpleApiProxyRepository GetRepository8(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository9**
> SimpleApiGroupDeployRepository GetRepository9(ctx, repositoryName)
Get repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**|  | 

### Return type

[**SimpleApiGroupDeployRepository**](SimpleApiGroupDeployRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvalidateCache**
> InvalidateCache(ctx, repositoryName)
Invalidate repository cache. Proxy or group repositories only.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to invalidate cache | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebuildIndex**
> RebuildIndex(ctx, repositoryName)
Schedule a 'Repair - Rebuild repository search' Task. Hosted or proxy repositories only.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to rebuild index | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository**
> UpdateRepository(ctx, repositoryName, optional)
Update Maven group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepositoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MavenGroupRepositoryApiRequest**](MavenGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository1**
> UpdateRepository1(ctx, repositoryName, optional)
Update Maven hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MavenHostedRepositoryApiRequest**](MavenHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository10**
> UpdateRepository10(ctx, repositoryName, optional)
Update npm proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository10Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository10Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NpmProxyRepositoryApiRequest**](NpmProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository11**
> UpdateRepository11(ctx, repositoryName, optional)
Update NuGet group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository11Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NugetGroupRepositoryApiRequest**](NugetGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository12**
> UpdateRepository12(ctx, repositoryName, optional)
Update NuGet hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository12Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository12Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NugetHostedRepositoryApiRequest**](NugetHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository13**
> UpdateRepository13(ctx, repositoryName, optional)
Update NuGet proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository13Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository13Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NugetProxyRepositoryApiRequest**](NugetProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository14**
> UpdateRepository14(ctx, repositoryName, optional)
Update RubyGems group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository14Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository14Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RubyGemsGroupRepositoryApiRequest**](RubyGemsGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository15**
> UpdateRepository15(ctx, repositoryName, optional)
Update RubyGems hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository15Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository15Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RubyGemsHostedRepositoryApiRequest**](RubyGemsHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository16**
> UpdateRepository16(ctx, repositoryName, optional)
Update RubyGems proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository16Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository16Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RubyGemsProxyRepositoryApiRequest**](RubyGemsProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository17**
> UpdateRepository17(ctx, repositoryName, optional)
Update Docker group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository17Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository17Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DockerGroupRepositoryApiRequest**](DockerGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository18**
> UpdateRepository18(ctx, repositoryName, optional)
Update Docker hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository18Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository18Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DockerHostedRepositoryApiRequest**](DockerHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository19**
> UpdateRepository19(ctx, repositoryName, optional)
Update Docker proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository19Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository19Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DockerProxyRepositoryApiRequest**](DockerProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository2**
> UpdateRepository2(ctx, repositoryName, optional)
Update Maven proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MavenProxyRepositoryApiRequest**](MavenProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository20**
> UpdateRepository20(ctx, repositoryName, optional)
Update Yum group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository20Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository20Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of YumGroupRepositoryApiRequest**](YumGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository21**
> UpdateRepository21(ctx, repositoryName, optional)
Update Yum hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository21Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository21Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of YumHostedRepositoryApiRequest**](YumHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository22**
> UpdateRepository22(ctx, repositoryName, optional)
Update Yum proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository22Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository22Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of YumProxyRepositoryApiRequest**](YumProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository23**
> UpdateRepository23(ctx, repositoryName, optional)
Update Helm hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository23Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository23Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HelmHostedRepositoryApiRequest**](HelmHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository24**
> UpdateRepository24(ctx, repositoryName, optional)
Update Helm proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository24Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository24Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HelmProxyRepositoryApiRequest**](HelmProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository25**
> UpdateRepository25(ctx, repositoryName, optional)
Update Git LFS hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository25Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository25Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GitLfsHostedRepositoryApiRequest**](GitLfsHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository26**
> UpdateRepository26(ctx, repositoryName, optional)
Update PyPI group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository26Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository26Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PypiGroupRepositoryApiRequest**](PypiGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository27**
> UpdateRepository27(ctx, repositoryName, optional)
Update PyPI hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository27Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository27Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PypiHostedRepositoryApiRequest**](PypiHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository28**
> UpdateRepository28(ctx, repositoryName, optional)
Update PyPI proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository28Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository28Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PypiProxyRepositoryApiRequest**](PypiProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository29**
> UpdateRepository29(ctx, repositoryName, optional)
Update conda proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository29Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository29Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CondaProxyRepositoryApiRequest**](CondaProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository3**
> UpdateRepository3(ctx, repositoryName, optional)
Update APT hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AptHostedRepositoryApiRequest**](AptHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository30**
> UpdateRepository30(ctx, repositoryName, optional)
Update Conan proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository30Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository30Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ConanProxyRepositoryApiRequest**](ConanProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository31**
> UpdateRepository31(ctx, repositoryName, optional)
Update R group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository31Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository31Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RGroupRepositoryApiRequest**](RGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository32**
> UpdateRepository32(ctx, repositoryName, optional)
Update R hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository32Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository32Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RHostedRepositoryApiRequest**](RHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository33**
> UpdateRepository33(ctx, repositoryName, optional)
Update R proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository33Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository33Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RProxyRepositoryApiRequest**](RProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository34**
> UpdateRepository34(ctx, repositoryName, optional)
Update Cocoapods proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository34Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository34Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CocoapodsProxyRepositoryApiRequest**](CocoapodsProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository35**
> UpdateRepository35(ctx, repositoryName, optional)
Update a Go group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository35Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository35Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GolangGroupRepositoryApiRequest**](GolangGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository36**
> UpdateRepository36(ctx, repositoryName, optional)
Update a Go proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository36Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository36Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GolangProxyRepositoryApiRequest**](GolangProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository37**
> UpdateRepository37(ctx, repositoryName, optional)
Update p2 proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository37Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository37Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of P2ProxyRepositoryApiRequest**](P2ProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository38**
> UpdateRepository38(ctx, repositoryName, optional)
Update Bower group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository38Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository38Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BowerGroupRepositoryApiRequest**](BowerGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository39**
> UpdateRepository39(ctx, repositoryName, optional)
Update Bower hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository39Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository39Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BowerHostedRepositoryApiRequest**](BowerHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository4**
> UpdateRepository4(ctx, repositoryName, optional)
Update APT proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AptProxyRepositoryApiRequest**](AptProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository40**
> UpdateRepository40(ctx, repositoryName, optional)
Update Bower proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository40Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository40Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BowerProxyRepositoryApiRequest**](BowerProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository5**
> UpdateRepository5(ctx, repositoryName, optional)
Update raw group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository5Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RawGroupRepositoryApiRequest**](RawGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository6**
> UpdateRepository6(ctx, repositoryName, optional)
Update raw hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository6Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RawHostedRepositoryApiRequest**](RawHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository7**
> UpdateRepository7(ctx, repositoryName, optional)
Update raw proxy repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository7Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository7Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RawProxyRepositoryApiRequest**](RawProxyRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository8**
> UpdateRepository8(ctx, repositoryName, optional)
Update npm group repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository8Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository8Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NpmGroupRepositoryApiRequest**](NpmGroupRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepository9**
> UpdateRepository9(ctx, repositoryName, optional)
Update npm hosted repository



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryName** | **string**| Name of the repository to update | 
 **optional** | ***RepositoryManagementApiUpdateRepository9Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryManagementApiUpdateRepository9Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NpmHostedRepositoryApiRequest**](NpmHostedRepositoryApiRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

