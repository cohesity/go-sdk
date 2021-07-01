# MongoDBRecoverJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered entities | [optional] 

## Methods

### NewMongoDBRecoverJobParams

`func NewMongoDBRecoverJobParams() *MongoDBRecoverJobParams`

NewMongoDBRecoverJobParams instantiates a new MongoDBRecoverJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBRecoverJobParamsWithDefaults

`func NewMongoDBRecoverJobParamsWithDefaults() *MongoDBRecoverJobParams`

NewMongoDBRecoverJobParamsWithDefaults instantiates a new MongoDBRecoverJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuffix

`func (o *MongoDBRecoverJobParams) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *MongoDBRecoverJobParams) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *MongoDBRecoverJobParams) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *MongoDBRecoverJobParams) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *MongoDBRecoverJobParams) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *MongoDBRecoverJobParams) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


