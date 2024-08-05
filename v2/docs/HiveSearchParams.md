# HiveSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HiveObjectTypes** | **[]string** | Specifies one or more Hive object types be searched. | 
**SearchString** | **NullableString** | Specifies the search string to search the Hive Objects | 

## Methods

### NewHiveSearchParams

`func NewHiveSearchParams(hiveObjectTypes []string, searchString NullableString, ) *HiveSearchParams`

NewHiveSearchParams instantiates a new HiveSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveSearchParamsWithDefaults

`func NewHiveSearchParamsWithDefaults() *HiveSearchParams`

NewHiveSearchParamsWithDefaults instantiates a new HiveSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHiveObjectTypes

`func (o *HiveSearchParams) GetHiveObjectTypes() []string`

GetHiveObjectTypes returns the HiveObjectTypes field if non-nil, zero value otherwise.

### GetHiveObjectTypesOk

`func (o *HiveSearchParams) GetHiveObjectTypesOk() (*[]string, bool)`

GetHiveObjectTypesOk returns a tuple with the HiveObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveObjectTypes

`func (o *HiveSearchParams) SetHiveObjectTypes(v []string)`

SetHiveObjectTypes sets HiveObjectTypes field to given value.


### GetSearchString

`func (o *HiveSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *HiveSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *HiveSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *HiveSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *HiveSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


