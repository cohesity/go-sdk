# AwsS3ObjectLevelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the object being protected. This can be a leaf level or non leaf level object. | 
**ObjectPrefixExclusions** | Pointer to **[]string** | Specifies the list of prefix paths excluded. Objects containing any of these prefixes in their path will be excluded. | [optional] 
**ObjectPrefixInclusions** | Pointer to **[]string** | Specifies the list of prefix paths included. Objects containing any of these prefixes in their path will be included. Among inclusion and exclusion, inclusion will take precedence. | [optional] 

## Methods

### NewAwsS3ObjectLevelParams

`func NewAwsS3ObjectLevelParams(id NullableInt64, ) *AwsS3ObjectLevelParams`

NewAwsS3ObjectLevelParams instantiates a new AwsS3ObjectLevelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsS3ObjectLevelParamsWithDefaults

`func NewAwsS3ObjectLevelParamsWithDefaults() *AwsS3ObjectLevelParams`

NewAwsS3ObjectLevelParamsWithDefaults instantiates a new AwsS3ObjectLevelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwsS3ObjectLevelParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsS3ObjectLevelParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsS3ObjectLevelParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AwsS3ObjectLevelParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AwsS3ObjectLevelParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectPrefixExclusions

`func (o *AwsS3ObjectLevelParams) GetObjectPrefixExclusions() []string`

GetObjectPrefixExclusions returns the ObjectPrefixExclusions field if non-nil, zero value otherwise.

### GetObjectPrefixExclusionsOk

`func (o *AwsS3ObjectLevelParams) GetObjectPrefixExclusionsOk() (*[]string, bool)`

GetObjectPrefixExclusionsOk returns a tuple with the ObjectPrefixExclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPrefixExclusions

`func (o *AwsS3ObjectLevelParams) SetObjectPrefixExclusions(v []string)`

SetObjectPrefixExclusions sets ObjectPrefixExclusions field to given value.

### HasObjectPrefixExclusions

`func (o *AwsS3ObjectLevelParams) HasObjectPrefixExclusions() bool`

HasObjectPrefixExclusions returns a boolean if a field has been set.

### GetObjectPrefixInclusions

`func (o *AwsS3ObjectLevelParams) GetObjectPrefixInclusions() []string`

GetObjectPrefixInclusions returns the ObjectPrefixInclusions field if non-nil, zero value otherwise.

### GetObjectPrefixInclusionsOk

`func (o *AwsS3ObjectLevelParams) GetObjectPrefixInclusionsOk() (*[]string, bool)`

GetObjectPrefixInclusionsOk returns a tuple with the ObjectPrefixInclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPrefixInclusions

`func (o *AwsS3ObjectLevelParams) SetObjectPrefixInclusions(v []string)`

SetObjectPrefixInclusions sets ObjectPrefixInclusions field to given value.

### HasObjectPrefixInclusions

`func (o *AwsS3ObjectLevelParams) HasObjectPrefixInclusions() bool`

HasObjectPrefixInclusions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


