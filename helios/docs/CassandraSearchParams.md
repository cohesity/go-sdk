# CassandraSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | **NullableString** | Specifies the search string to search the Cassandra Objects | 
**CassandraObjectTypes** | **[]string** | Specifies one or more Cassandra object types to be searched. | 

## Methods

### NewCassandraSearchParams

`func NewCassandraSearchParams(searchString NullableString, cassandraObjectTypes []string, ) *CassandraSearchParams`

NewCassandraSearchParams instantiates a new CassandraSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraSearchParamsWithDefaults

`func NewCassandraSearchParamsWithDefaults() *CassandraSearchParams`

NewCassandraSearchParamsWithDefaults instantiates a new CassandraSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchString

`func (o *CassandraSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *CassandraSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *CassandraSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *CassandraSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *CassandraSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetCassandraObjectTypes

`func (o *CassandraSearchParams) GetCassandraObjectTypes() []string`

GetCassandraObjectTypes returns the CassandraObjectTypes field if non-nil, zero value otherwise.

### GetCassandraObjectTypesOk

`func (o *CassandraSearchParams) GetCassandraObjectTypesOk() (*[]string, bool)`

GetCassandraObjectTypesOk returns a tuple with the CassandraObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraObjectTypes

`func (o *CassandraSearchParams) SetCassandraObjectTypes(v []string)`

SetCassandraObjectTypes sets CassandraObjectTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


