# CassandraOnPremSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | **NullableString** | Specifies the search string to search the Cassandra Objects | 
**CassandraObjectTypes** | **[]string** | Specifies one or more Cassandra object types to be searched. | 
**SourceIds** | Pointer to **[]int64** | Specifies a list of source ids. Only files found in these sources will be returned. | [optional] 

## Methods

### NewCassandraOnPremSearchParams

`func NewCassandraOnPremSearchParams(searchString NullableString, cassandraObjectTypes []string, ) *CassandraOnPremSearchParams`

NewCassandraOnPremSearchParams instantiates a new CassandraOnPremSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraOnPremSearchParamsWithDefaults

`func NewCassandraOnPremSearchParamsWithDefaults() *CassandraOnPremSearchParams`

NewCassandraOnPremSearchParamsWithDefaults instantiates a new CassandraOnPremSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchString

`func (o *CassandraOnPremSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *CassandraOnPremSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *CassandraOnPremSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *CassandraOnPremSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *CassandraOnPremSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetCassandraObjectTypes

`func (o *CassandraOnPremSearchParams) GetCassandraObjectTypes() []string`

GetCassandraObjectTypes returns the CassandraObjectTypes field if non-nil, zero value otherwise.

### GetCassandraObjectTypesOk

`func (o *CassandraOnPremSearchParams) GetCassandraObjectTypesOk() (*[]string, bool)`

GetCassandraObjectTypesOk returns a tuple with the CassandraObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraObjectTypes

`func (o *CassandraOnPremSearchParams) SetCassandraObjectTypes(v []string)`

SetCassandraObjectTypes sets CassandraObjectTypes field to given value.


### GetSourceIds

`func (o *CassandraOnPremSearchParams) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *CassandraOnPremSearchParams) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *CassandraOnPremSearchParams) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *CassandraOnPremSearchParams) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *CassandraOnPremSearchParams) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *CassandraOnPremSearchParams) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


