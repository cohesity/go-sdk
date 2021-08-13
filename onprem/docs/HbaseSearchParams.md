# HbaseSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | **NullableString** | Specifies the search string to search the Hbase Objects | 
**HbaseObjectTypes** | **[]string** | Specifies one or more Hbase object types be searched. | 

## Methods

### NewHbaseSearchParams

`func NewHbaseSearchParams(searchString NullableString, hbaseObjectTypes []string, ) *HbaseSearchParams`

NewHbaseSearchParams instantiates a new HbaseSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHbaseSearchParamsWithDefaults

`func NewHbaseSearchParamsWithDefaults() *HbaseSearchParams`

NewHbaseSearchParamsWithDefaults instantiates a new HbaseSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchString

`func (o *HbaseSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *HbaseSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *HbaseSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *HbaseSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *HbaseSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetHbaseObjectTypes

`func (o *HbaseSearchParams) GetHbaseObjectTypes() []string`

GetHbaseObjectTypes returns the HbaseObjectTypes field if non-nil, zero value otherwise.

### GetHbaseObjectTypesOk

`func (o *HbaseSearchParams) GetHbaseObjectTypesOk() (*[]string, bool)`

GetHbaseObjectTypesOk returns a tuple with the HbaseObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseObjectTypes

`func (o *HbaseSearchParams) SetHbaseObjectTypes(v []string)`

SetHbaseObjectTypes sets HbaseObjectTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


