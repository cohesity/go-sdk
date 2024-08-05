# MongoDbOnPremSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MongoDBObjectTypes** | **[]string** | Specifies one or more MongoDB object types be searched. | 
**SearchString** | **NullableString** | Specifies the search string to search the MongoDB Objects | 
**SourceIds** | Pointer to **[]int64** | Specifies a list of source ids. Only files found in these sources will be returned. | [optional] 

## Methods

### NewMongoDbOnPremSearchParams

`func NewMongoDbOnPremSearchParams(mongoDBObjectTypes []string, searchString NullableString, ) *MongoDbOnPremSearchParams`

NewMongoDbOnPremSearchParams instantiates a new MongoDbOnPremSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDbOnPremSearchParamsWithDefaults

`func NewMongoDbOnPremSearchParamsWithDefaults() *MongoDbOnPremSearchParams`

NewMongoDbOnPremSearchParamsWithDefaults instantiates a new MongoDbOnPremSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMongoDBObjectTypes

`func (o *MongoDbOnPremSearchParams) GetMongoDBObjectTypes() []string`

GetMongoDBObjectTypes returns the MongoDBObjectTypes field if non-nil, zero value otherwise.

### GetMongoDBObjectTypesOk

`func (o *MongoDbOnPremSearchParams) GetMongoDBObjectTypesOk() (*[]string, bool)`

GetMongoDBObjectTypesOk returns a tuple with the MongoDBObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBObjectTypes

`func (o *MongoDbOnPremSearchParams) SetMongoDBObjectTypes(v []string)`

SetMongoDBObjectTypes sets MongoDBObjectTypes field to given value.


### GetSearchString

`func (o *MongoDbOnPremSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *MongoDbOnPremSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *MongoDbOnPremSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *MongoDbOnPremSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *MongoDbOnPremSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetSourceIds

`func (o *MongoDbOnPremSearchParams) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *MongoDbOnPremSearchParams) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *MongoDbOnPremSearchParams) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *MongoDbOnPremSearchParams) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *MongoDbOnPremSearchParams) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *MongoDbOnPremSearchParams) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


