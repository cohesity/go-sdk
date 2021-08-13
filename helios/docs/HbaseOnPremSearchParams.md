# HbaseOnPremSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | **NullableString** | Specifies the search string to search the Hbase Objects | 
**HbaseObjectTypes** | **[]string** | Specifies one or more Hbase object types be searched. | 
**SourceIds** | Pointer to **[]int64** | Specifies a list of source ids. Only files found in these sources will be returned. | [optional] 

## Methods

### NewHbaseOnPremSearchParams

`func NewHbaseOnPremSearchParams(searchString NullableString, hbaseObjectTypes []string, ) *HbaseOnPremSearchParams`

NewHbaseOnPremSearchParams instantiates a new HbaseOnPremSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHbaseOnPremSearchParamsWithDefaults

`func NewHbaseOnPremSearchParamsWithDefaults() *HbaseOnPremSearchParams`

NewHbaseOnPremSearchParamsWithDefaults instantiates a new HbaseOnPremSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchString

`func (o *HbaseOnPremSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *HbaseOnPremSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *HbaseOnPremSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *HbaseOnPremSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *HbaseOnPremSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetHbaseObjectTypes

`func (o *HbaseOnPremSearchParams) GetHbaseObjectTypes() []string`

GetHbaseObjectTypes returns the HbaseObjectTypes field if non-nil, zero value otherwise.

### GetHbaseObjectTypesOk

`func (o *HbaseOnPremSearchParams) GetHbaseObjectTypesOk() (*[]string, bool)`

GetHbaseObjectTypesOk returns a tuple with the HbaseObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseObjectTypes

`func (o *HbaseOnPremSearchParams) SetHbaseObjectTypes(v []string)`

SetHbaseObjectTypes sets HbaseObjectTypes field to given value.


### GetSourceIds

`func (o *HbaseOnPremSearchParams) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *HbaseOnPremSearchParams) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *HbaseOnPremSearchParams) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *HbaseOnPremSearchParams) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *HbaseOnPremSearchParams) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *HbaseOnPremSearchParams) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


