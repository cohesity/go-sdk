# ExchangeDagHostsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeDagProtectionPreference** | Pointer to [**ExchangeDAGProtectionPreference**](ExchangeDAGProtectionPreference.md) |  | [optional] 
**ExchangeHostInfoList** | Pointer to [**[]ExchangeHostInfo**](ExchangeHostInfo.md) | Specifies the list of exchange hosts that belong to Exchange DAG. | [optional] 
**Guid** | Pointer to **NullableString** | Specifies the Unique GUID for the DAG. | [optional] 
**IsStandaloneHost** | Pointer to **NullableBool** | Specifies if the endpoint provided in the request is a standlone exchange server or not. exchangeHostInfoList is not populated if it is a standalone exchange server. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the display name of the DAG. | [optional] 
**ProtectionSourceId** | Pointer to **NullableInt64** | Specifies the Protection Source Id of the Exchange DAG if it is already created. | [optional] 

## Methods

### NewExchangeDagHostsResponse

`func NewExchangeDagHostsResponse() *ExchangeDagHostsResponse`

NewExchangeDagHostsResponse instantiates a new ExchangeDagHostsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeDagHostsResponseWithDefaults

`func NewExchangeDagHostsResponseWithDefaults() *ExchangeDagHostsResponse`

NewExchangeDagHostsResponseWithDefaults instantiates a new ExchangeDagHostsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeDagProtectionPreference

`func (o *ExchangeDagHostsResponse) GetExchangeDagProtectionPreference() ExchangeDAGProtectionPreference`

GetExchangeDagProtectionPreference returns the ExchangeDagProtectionPreference field if non-nil, zero value otherwise.

### GetExchangeDagProtectionPreferenceOk

`func (o *ExchangeDagHostsResponse) GetExchangeDagProtectionPreferenceOk() (*ExchangeDAGProtectionPreference, bool)`

GetExchangeDagProtectionPreferenceOk returns a tuple with the ExchangeDagProtectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDagProtectionPreference

`func (o *ExchangeDagHostsResponse) SetExchangeDagProtectionPreference(v ExchangeDAGProtectionPreference)`

SetExchangeDagProtectionPreference sets ExchangeDagProtectionPreference field to given value.

### HasExchangeDagProtectionPreference

`func (o *ExchangeDagHostsResponse) HasExchangeDagProtectionPreference() bool`

HasExchangeDagProtectionPreference returns a boolean if a field has been set.

### GetExchangeHostInfoList

`func (o *ExchangeDagHostsResponse) GetExchangeHostInfoList() []ExchangeHostInfo`

GetExchangeHostInfoList returns the ExchangeHostInfoList field if non-nil, zero value otherwise.

### GetExchangeHostInfoListOk

`func (o *ExchangeDagHostsResponse) GetExchangeHostInfoListOk() (*[]ExchangeHostInfo, bool)`

GetExchangeHostInfoListOk returns a tuple with the ExchangeHostInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeHostInfoList

`func (o *ExchangeDagHostsResponse) SetExchangeHostInfoList(v []ExchangeHostInfo)`

SetExchangeHostInfoList sets ExchangeHostInfoList field to given value.

### HasExchangeHostInfoList

`func (o *ExchangeDagHostsResponse) HasExchangeHostInfoList() bool`

HasExchangeHostInfoList returns a boolean if a field has been set.

### SetExchangeHostInfoListNil

`func (o *ExchangeDagHostsResponse) SetExchangeHostInfoListNil(b bool)`

 SetExchangeHostInfoListNil sets the value for ExchangeHostInfoList to be an explicit nil

### UnsetExchangeHostInfoList
`func (o *ExchangeDagHostsResponse) UnsetExchangeHostInfoList()`

UnsetExchangeHostInfoList ensures that no value is present for ExchangeHostInfoList, not even an explicit nil
### GetGuid

`func (o *ExchangeDagHostsResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ExchangeDagHostsResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ExchangeDagHostsResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ExchangeDagHostsResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *ExchangeDagHostsResponse) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *ExchangeDagHostsResponse) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetIsStandaloneHost

`func (o *ExchangeDagHostsResponse) GetIsStandaloneHost() bool`

GetIsStandaloneHost returns the IsStandaloneHost field if non-nil, zero value otherwise.

### GetIsStandaloneHostOk

`func (o *ExchangeDagHostsResponse) GetIsStandaloneHostOk() (*bool, bool)`

GetIsStandaloneHostOk returns a tuple with the IsStandaloneHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandaloneHost

`func (o *ExchangeDagHostsResponse) SetIsStandaloneHost(v bool)`

SetIsStandaloneHost sets IsStandaloneHost field to given value.

### HasIsStandaloneHost

`func (o *ExchangeDagHostsResponse) HasIsStandaloneHost() bool`

HasIsStandaloneHost returns a boolean if a field has been set.

### SetIsStandaloneHostNil

`func (o *ExchangeDagHostsResponse) SetIsStandaloneHostNil(b bool)`

 SetIsStandaloneHostNil sets the value for IsStandaloneHost to be an explicit nil

### UnsetIsStandaloneHost
`func (o *ExchangeDagHostsResponse) UnsetIsStandaloneHost()`

UnsetIsStandaloneHost ensures that no value is present for IsStandaloneHost, not even an explicit nil
### GetName

`func (o *ExchangeDagHostsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeDagHostsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeDagHostsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeDagHostsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExchangeDagHostsResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExchangeDagHostsResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtectionSourceId

`func (o *ExchangeDagHostsResponse) GetProtectionSourceId() int64`

GetProtectionSourceId returns the ProtectionSourceId field if non-nil, zero value otherwise.

### GetProtectionSourceIdOk

`func (o *ExchangeDagHostsResponse) GetProtectionSourceIdOk() (*int64, bool)`

GetProtectionSourceIdOk returns a tuple with the ProtectionSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceId

`func (o *ExchangeDagHostsResponse) SetProtectionSourceId(v int64)`

SetProtectionSourceId sets ProtectionSourceId field to given value.

### HasProtectionSourceId

`func (o *ExchangeDagHostsResponse) HasProtectionSourceId() bool`

HasProtectionSourceId returns a boolean if a field has been set.

### SetProtectionSourceIdNil

`func (o *ExchangeDagHostsResponse) SetProtectionSourceIdNil(b bool)`

 SetProtectionSourceIdNil sets the value for ProtectionSourceId to be an explicit nil

### UnsetProtectionSourceId
`func (o *ExchangeDagHostsResponse) UnsetProtectionSourceId()`

UnsetProtectionSourceId ensures that no value is present for ProtectionSourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


