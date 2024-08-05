# MongodbSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MongoDBObjectTypes** | **[]string** | Specifies one or more MongoDB object types be searched. | 
**SearchString** | **NullableString** | Specifies the search string to search the MongoDB Objects | 

## Methods

### NewMongodbSearchParams

`func NewMongodbSearchParams(mongoDBObjectTypes []string, searchString NullableString, ) *MongodbSearchParams`

NewMongodbSearchParams instantiates a new MongodbSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongodbSearchParamsWithDefaults

`func NewMongodbSearchParamsWithDefaults() *MongodbSearchParams`

NewMongodbSearchParamsWithDefaults instantiates a new MongodbSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMongoDBObjectTypes

`func (o *MongodbSearchParams) GetMongoDBObjectTypes() []string`

GetMongoDBObjectTypes returns the MongoDBObjectTypes field if non-nil, zero value otherwise.

### GetMongoDBObjectTypesOk

`func (o *MongodbSearchParams) GetMongoDBObjectTypesOk() (*[]string, bool)`

GetMongoDBObjectTypesOk returns a tuple with the MongoDBObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBObjectTypes

`func (o *MongodbSearchParams) SetMongoDBObjectTypes(v []string)`

SetMongoDBObjectTypes sets MongoDBObjectTypes field to given value.


### GetSearchString

`func (o *MongodbSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *MongodbSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *MongodbSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *MongodbSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *MongodbSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


