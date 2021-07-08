# SearchProductionAdObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinguishedNames** | Pointer to **[]string** | Specifies the list of the distinguished names of the AD objects. | [optional] 
**ObjectGuids** | Pointer to **[]string** | Specifies the list of the guids of the AD objects. | [optional] 
**ProtectionSourceId** | Pointer to **NullableInt64** | ProtectionSourceId is the Id of the Domain Controller host on which we want to search for AD objects. | [optional] 
**SamAccountNames** | Pointer to **[]string** | Specifies the list of the sam account names of the AD objects. | [optional] 

## Methods

### NewSearchProductionAdObjectsRequest

`func NewSearchProductionAdObjectsRequest() *SearchProductionAdObjectsRequest`

NewSearchProductionAdObjectsRequest instantiates a new SearchProductionAdObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchProductionAdObjectsRequestWithDefaults

`func NewSearchProductionAdObjectsRequestWithDefaults() *SearchProductionAdObjectsRequest`

NewSearchProductionAdObjectsRequestWithDefaults instantiates a new SearchProductionAdObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinguishedNames

`func (o *SearchProductionAdObjectsRequest) GetDistinguishedNames() []string`

GetDistinguishedNames returns the DistinguishedNames field if non-nil, zero value otherwise.

### GetDistinguishedNamesOk

`func (o *SearchProductionAdObjectsRequest) GetDistinguishedNamesOk() (*[]string, bool)`

GetDistinguishedNamesOk returns a tuple with the DistinguishedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedNames

`func (o *SearchProductionAdObjectsRequest) SetDistinguishedNames(v []string)`

SetDistinguishedNames sets DistinguishedNames field to given value.

### HasDistinguishedNames

`func (o *SearchProductionAdObjectsRequest) HasDistinguishedNames() bool`

HasDistinguishedNames returns a boolean if a field has been set.

### SetDistinguishedNamesNil

`func (o *SearchProductionAdObjectsRequest) SetDistinguishedNamesNil(b bool)`

 SetDistinguishedNamesNil sets the value for DistinguishedNames to be an explicit nil

### UnsetDistinguishedNames
`func (o *SearchProductionAdObjectsRequest) UnsetDistinguishedNames()`

UnsetDistinguishedNames ensures that no value is present for DistinguishedNames, not even an explicit nil
### GetObjectGuids

`func (o *SearchProductionAdObjectsRequest) GetObjectGuids() []string`

GetObjectGuids returns the ObjectGuids field if non-nil, zero value otherwise.

### GetObjectGuidsOk

`func (o *SearchProductionAdObjectsRequest) GetObjectGuidsOk() (*[]string, bool)`

GetObjectGuidsOk returns a tuple with the ObjectGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGuids

`func (o *SearchProductionAdObjectsRequest) SetObjectGuids(v []string)`

SetObjectGuids sets ObjectGuids field to given value.

### HasObjectGuids

`func (o *SearchProductionAdObjectsRequest) HasObjectGuids() bool`

HasObjectGuids returns a boolean if a field has been set.

### SetObjectGuidsNil

`func (o *SearchProductionAdObjectsRequest) SetObjectGuidsNil(b bool)`

 SetObjectGuidsNil sets the value for ObjectGuids to be an explicit nil

### UnsetObjectGuids
`func (o *SearchProductionAdObjectsRequest) UnsetObjectGuids()`

UnsetObjectGuids ensures that no value is present for ObjectGuids, not even an explicit nil
### GetProtectionSourceId

`func (o *SearchProductionAdObjectsRequest) GetProtectionSourceId() int64`

GetProtectionSourceId returns the ProtectionSourceId field if non-nil, zero value otherwise.

### GetProtectionSourceIdOk

`func (o *SearchProductionAdObjectsRequest) GetProtectionSourceIdOk() (*int64, bool)`

GetProtectionSourceIdOk returns a tuple with the ProtectionSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceId

`func (o *SearchProductionAdObjectsRequest) SetProtectionSourceId(v int64)`

SetProtectionSourceId sets ProtectionSourceId field to given value.

### HasProtectionSourceId

`func (o *SearchProductionAdObjectsRequest) HasProtectionSourceId() bool`

HasProtectionSourceId returns a boolean if a field has been set.

### SetProtectionSourceIdNil

`func (o *SearchProductionAdObjectsRequest) SetProtectionSourceIdNil(b bool)`

 SetProtectionSourceIdNil sets the value for ProtectionSourceId to be an explicit nil

### UnsetProtectionSourceId
`func (o *SearchProductionAdObjectsRequest) UnsetProtectionSourceId()`

UnsetProtectionSourceId ensures that no value is present for ProtectionSourceId, not even an explicit nil
### GetSamAccountNames

`func (o *SearchProductionAdObjectsRequest) GetSamAccountNames() []string`

GetSamAccountNames returns the SamAccountNames field if non-nil, zero value otherwise.

### GetSamAccountNamesOk

`func (o *SearchProductionAdObjectsRequest) GetSamAccountNamesOk() (*[]string, bool)`

GetSamAccountNamesOk returns a tuple with the SamAccountNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountNames

`func (o *SearchProductionAdObjectsRequest) SetSamAccountNames(v []string)`

SetSamAccountNames sets SamAccountNames field to given value.

### HasSamAccountNames

`func (o *SearchProductionAdObjectsRequest) HasSamAccountNames() bool`

HasSamAccountNames returns a boolean if a field has been set.

### SetSamAccountNamesNil

`func (o *SearchProductionAdObjectsRequest) SetSamAccountNamesNil(b bool)`

 SetSamAccountNamesNil sets the value for SamAccountNames to be an explicit nil

### UnsetSamAccountNames
`func (o *SearchProductionAdObjectsRequest) UnsetSamAccountNames()`

UnsetSamAccountNames ensures that no value is present for SamAccountNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


