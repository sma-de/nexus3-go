# AptProxyApiRepository

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique identifier for this repository | [optional] [default to null]
**Online** | **bool** | Whether this repository accepts incoming requests | [default to null]
**Storage** | [***StorageAttributes**](StorageAttributes.md) |  | [default to null]
**Cleanup** | [***CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] [default to null]
**Proxy** | [***ProxyAttributes**](ProxyAttributes.md) |  | [default to null]
**NegativeCache** | [***NegativeCacheAttributes**](NegativeCacheAttributes.md) |  | [default to null]
**HttpClient** | [***HttpClientAttributes**](HttpClientAttributes.md) |  | [default to null]
**RoutingRuleName** | **string** | The name of the routing rule assigned to this repository | [optional] [default to null]
**Replication** | [***ReplicationAttributes**](ReplicationAttributes.md) |  | [optional] [default to null]
**Apt** | [***AptProxyRepositoriesAttributes**](AptProxyRepositoriesAttributes.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


