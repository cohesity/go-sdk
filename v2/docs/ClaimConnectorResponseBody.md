# ClaimConnectorResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **int64** | Connection ID. | [optional] 
**ConnectorCertBundle** | [**CertificateBundle**](CertificateBundle.md) |  | 
**ConnectorGuid** | Pointer to **int64** | Unique id of the connector. | [optional] 
**TenantCaChain** | **[]string** | Tenant&#39;s CA chain. | 
**TenantId** | Pointer to **string** | Tenant id associated with the claimed connector. | [optional] 

## Methods

### NewClaimConnectorResponseBody

`func NewClaimConnectorResponseBody(connectorCertBundle CertificateBundle, tenantCaChain []string, ) *ClaimConnectorResponseBody`

NewClaimConnectorResponseBody instantiates a new ClaimConnectorResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimConnectorResponseBodyWithDefaults

`func NewClaimConnectorResponseBodyWithDefaults() *ClaimConnectorResponseBody`

NewClaimConnectorResponseBodyWithDefaults instantiates a new ClaimConnectorResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ClaimConnectorResponseBody) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ClaimConnectorResponseBody) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ClaimConnectorResponseBody) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ClaimConnectorResponseBody) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectorCertBundle

`func (o *ClaimConnectorResponseBody) GetConnectorCertBundle() CertificateBundle`

GetConnectorCertBundle returns the ConnectorCertBundle field if non-nil, zero value otherwise.

### GetConnectorCertBundleOk

`func (o *ClaimConnectorResponseBody) GetConnectorCertBundleOk() (*CertificateBundle, bool)`

GetConnectorCertBundleOk returns a tuple with the ConnectorCertBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCertBundle

`func (o *ClaimConnectorResponseBody) SetConnectorCertBundle(v CertificateBundle)`

SetConnectorCertBundle sets ConnectorCertBundle field to given value.


### GetConnectorGuid

`func (o *ClaimConnectorResponseBody) GetConnectorGuid() int64`

GetConnectorGuid returns the ConnectorGuid field if non-nil, zero value otherwise.

### GetConnectorGuidOk

`func (o *ClaimConnectorResponseBody) GetConnectorGuidOk() (*int64, bool)`

GetConnectorGuidOk returns a tuple with the ConnectorGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorGuid

`func (o *ClaimConnectorResponseBody) SetConnectorGuid(v int64)`

SetConnectorGuid sets ConnectorGuid field to given value.

### HasConnectorGuid

`func (o *ClaimConnectorResponseBody) HasConnectorGuid() bool`

HasConnectorGuid returns a boolean if a field has been set.

### GetTenantCaChain

`func (o *ClaimConnectorResponseBody) GetTenantCaChain() []string`

GetTenantCaChain returns the TenantCaChain field if non-nil, zero value otherwise.

### GetTenantCaChainOk

`func (o *ClaimConnectorResponseBody) GetTenantCaChainOk() (*[]string, bool)`

GetTenantCaChainOk returns a tuple with the TenantCaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCaChain

`func (o *ClaimConnectorResponseBody) SetTenantCaChain(v []string)`

SetTenantCaChain sets TenantCaChain field to given value.


### GetTenantId

`func (o *ClaimConnectorResponseBody) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ClaimConnectorResponseBody) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ClaimConnectorResponseBody) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ClaimConnectorResponseBody) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


