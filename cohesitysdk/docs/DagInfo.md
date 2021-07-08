# DagInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagApplicationServerInfoList** | Pointer to [**[]DagApplicationServerInfo**](DagApplicationServerInfo.md) | Specifies the status of all the Exchange Application Servers that are part of this DAG. | [optional] 
**ExchangeDagProtectionPreference** | Pointer to [**ExchangeDAGProtectionPreference**](ExchangeDAGProtectionPreference.md) |  | [optional] 
**Guid** | Pointer to **NullableString** | Specifies Unique GUID for the DAG. | [optional] 
**Name** | Pointer to **NullableString** | Specifies display name of the DAG. | [optional] 

## Methods

### NewDagInfo

`func NewDagInfo() *DagInfo`

NewDagInfo instantiates a new DagInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDagInfoWithDefaults

`func NewDagInfoWithDefaults() *DagInfo`

NewDagInfoWithDefaults instantiates a new DagInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagApplicationServerInfoList

`func (o *DagInfo) GetDagApplicationServerInfoList() []DagApplicationServerInfo`

GetDagApplicationServerInfoList returns the DagApplicationServerInfoList field if non-nil, zero value otherwise.

### GetDagApplicationServerInfoListOk

`func (o *DagInfo) GetDagApplicationServerInfoListOk() (*[]DagApplicationServerInfo, bool)`

GetDagApplicationServerInfoListOk returns a tuple with the DagApplicationServerInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagApplicationServerInfoList

`func (o *DagInfo) SetDagApplicationServerInfoList(v []DagApplicationServerInfo)`

SetDagApplicationServerInfoList sets DagApplicationServerInfoList field to given value.

### HasDagApplicationServerInfoList

`func (o *DagInfo) HasDagApplicationServerInfoList() bool`

HasDagApplicationServerInfoList returns a boolean if a field has been set.

### SetDagApplicationServerInfoListNil

`func (o *DagInfo) SetDagApplicationServerInfoListNil(b bool)`

 SetDagApplicationServerInfoListNil sets the value for DagApplicationServerInfoList to be an explicit nil

### UnsetDagApplicationServerInfoList
`func (o *DagInfo) UnsetDagApplicationServerInfoList()`

UnsetDagApplicationServerInfoList ensures that no value is present for DagApplicationServerInfoList, not even an explicit nil
### GetExchangeDagProtectionPreference

`func (o *DagInfo) GetExchangeDagProtectionPreference() ExchangeDAGProtectionPreference`

GetExchangeDagProtectionPreference returns the ExchangeDagProtectionPreference field if non-nil, zero value otherwise.

### GetExchangeDagProtectionPreferenceOk

`func (o *DagInfo) GetExchangeDagProtectionPreferenceOk() (*ExchangeDAGProtectionPreference, bool)`

GetExchangeDagProtectionPreferenceOk returns a tuple with the ExchangeDagProtectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDagProtectionPreference

`func (o *DagInfo) SetExchangeDagProtectionPreference(v ExchangeDAGProtectionPreference)`

SetExchangeDagProtectionPreference sets ExchangeDagProtectionPreference field to given value.

### HasExchangeDagProtectionPreference

`func (o *DagInfo) HasExchangeDagProtectionPreference() bool`

HasExchangeDagProtectionPreference returns a boolean if a field has been set.

### GetGuid

`func (o *DagInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *DagInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *DagInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *DagInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *DagInfo) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *DagInfo) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetName

`func (o *DagInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DagInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DagInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DagInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DagInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DagInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


