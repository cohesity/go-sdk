# HiveOnPremSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HiveObjectTypes** | **[]string** | Specifies one or more Hive object types be searched. | 
**SearchString** | **NullableString** | Specifies the search string to search the Hive Objects | 
**SourceIds** | Pointer to **[]int64** | Specifies a list of source ids. Only files found in these sources will be returned. | [optional] 

## Methods

### NewHiveOnPremSearchParams

`func NewHiveOnPremSearchParams(hiveObjectTypes []string, searchString NullableString, ) *HiveOnPremSearchParams`

NewHiveOnPremSearchParams instantiates a new HiveOnPremSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveOnPremSearchParamsWithDefaults

`func NewHiveOnPremSearchParamsWithDefaults() *HiveOnPremSearchParams`

NewHiveOnPremSearchParamsWithDefaults instantiates a new HiveOnPremSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHiveObjectTypes

`func (o *HiveOnPremSearchParams) GetHiveObjectTypes() []string`

GetHiveObjectTypes returns the HiveObjectTypes field if non-nil, zero value otherwise.

### GetHiveObjectTypesOk

`func (o *HiveOnPremSearchParams) GetHiveObjectTypesOk() (*[]string, bool)`

GetHiveObjectTypesOk returns a tuple with the HiveObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveObjectTypes

`func (o *HiveOnPremSearchParams) SetHiveObjectTypes(v []string)`

SetHiveObjectTypes sets HiveObjectTypes field to given value.


### GetSearchString

`func (o *HiveOnPremSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *HiveOnPremSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *HiveOnPremSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *HiveOnPremSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *HiveOnPremSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetSourceIds

`func (o *HiveOnPremSearchParams) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *HiveOnPremSearchParams) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *HiveOnPremSearchParams) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *HiveOnPremSearchParams) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *HiveOnPremSearchParams) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *HiveOnPremSearchParams) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


