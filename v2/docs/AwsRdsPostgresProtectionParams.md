# AwsRdsPostgresProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]AwsObjectLevelParams**](AwsObjectLevelParams.md) | Specifies the objects to be protected. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the source of the objects. | [optional] [readonly] 

## Methods

### NewAwsRdsPostgresProtectionParams

`func NewAwsRdsPostgresProtectionParams() *AwsRdsPostgresProtectionParams`

NewAwsRdsPostgresProtectionParams instantiates a new AwsRdsPostgresProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRdsPostgresProtectionParamsWithDefaults

`func NewAwsRdsPostgresProtectionParamsWithDefaults() *AwsRdsPostgresProtectionParams`

NewAwsRdsPostgresProtectionParamsWithDefaults instantiates a new AwsRdsPostgresProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *AwsRdsPostgresProtectionParams) GetObjects() []AwsObjectLevelParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AwsRdsPostgresProtectionParams) GetObjectsOk() (*[]AwsObjectLevelParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AwsRdsPostgresProtectionParams) SetObjects(v []AwsObjectLevelParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AwsRdsPostgresProtectionParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetSourceId

`func (o *AwsRdsPostgresProtectionParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AwsRdsPostgresProtectionParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AwsRdsPostgresProtectionParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AwsRdsPostgresProtectionParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *AwsRdsPostgresProtectionParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AwsRdsPostgresProtectionParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


