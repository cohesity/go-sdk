# DagApplicationServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **NullableString** | Specifies the fully qualified domain name of the Exchange Server. | [optional] 
**Guid** | Pointer to **NullableString** | Specifies the Guid of the Exchange Application Server. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the entity id of the Exchange Application server. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the display name of the Exchange Application Server. | [optional] 
**OwnerId** | Pointer to **NullableInt64** | Specifies the entity id of the owner entity of the Exchange Application Server. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the registration of the Exchange Application Server. Specifies the status of registration of Exchange Application Server. &#39;kUnknown&#39; indicates the status is not known. &#39;kHealthy&#39; indicates the status is healty and is registered as Exchange Server. &#39;kUnHealthy&#39; indicates the exchange application is registered on the physical server but it is unreachable now. &#39;kUnregistered&#39; indicates the server is not registered as physical source. &#39;kUnreachable&#39; indicates the server is not reachable from the cohesity cluster or the cohesity protection server is not installed on the exchange server. &#39;kDetached&#39; indicates the server is removed from the ExchangeDAG. | [optional] 
**TotalSizeBytes** | Pointer to **NullableInt64** | Specifies the total size of all Exchange database copies in all the Exchange Application Servers that are part of the DAG. | [optional] 

## Methods

### NewDagApplicationServerInfo

`func NewDagApplicationServerInfo() *DagApplicationServerInfo`

NewDagApplicationServerInfo instantiates a new DagApplicationServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDagApplicationServerInfoWithDefaults

`func NewDagApplicationServerInfoWithDefaults() *DagApplicationServerInfo`

NewDagApplicationServerInfoWithDefaults instantiates a new DagApplicationServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *DagApplicationServerInfo) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DagApplicationServerInfo) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DagApplicationServerInfo) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *DagApplicationServerInfo) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *DagApplicationServerInfo) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *DagApplicationServerInfo) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetGuid

`func (o *DagApplicationServerInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *DagApplicationServerInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *DagApplicationServerInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *DagApplicationServerInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *DagApplicationServerInfo) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *DagApplicationServerInfo) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetId

`func (o *DagApplicationServerInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DagApplicationServerInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DagApplicationServerInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DagApplicationServerInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DagApplicationServerInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DagApplicationServerInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DagApplicationServerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DagApplicationServerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DagApplicationServerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DagApplicationServerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DagApplicationServerInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DagApplicationServerInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *DagApplicationServerInfo) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DagApplicationServerInfo) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DagApplicationServerInfo) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DagApplicationServerInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *DagApplicationServerInfo) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *DagApplicationServerInfo) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetStatus

`func (o *DagApplicationServerInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DagApplicationServerInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DagApplicationServerInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DagApplicationServerInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DagApplicationServerInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DagApplicationServerInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTotalSizeBytes

`func (o *DagApplicationServerInfo) GetTotalSizeBytes() int64`

GetTotalSizeBytes returns the TotalSizeBytes field if non-nil, zero value otherwise.

### GetTotalSizeBytesOk

`func (o *DagApplicationServerInfo) GetTotalSizeBytesOk() (*int64, bool)`

GetTotalSizeBytesOk returns a tuple with the TotalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeBytes

`func (o *DagApplicationServerInfo) SetTotalSizeBytes(v int64)`

SetTotalSizeBytes sets TotalSizeBytes field to given value.

### HasTotalSizeBytes

`func (o *DagApplicationServerInfo) HasTotalSizeBytes() bool`

HasTotalSizeBytes returns a boolean if a field has been set.

### SetTotalSizeBytesNil

`func (o *DagApplicationServerInfo) SetTotalSizeBytesNil(b bool)`

 SetTotalSizeBytesNil sets the value for TotalSizeBytes to be an explicit nil

### UnsetTotalSizeBytes
`func (o *DagApplicationServerInfo) UnsetTotalSizeBytes()`

UnsetTotalSizeBytes ensures that no value is present for TotalSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


