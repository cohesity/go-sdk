# AwsRdsSnapshotManagerObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]AwsObjectLevelParams**](AwsObjectLevelParams.md) | Specifies the objects to be protected. | [optional] 

## Methods

### NewAwsRdsSnapshotManagerObjectProtectionParams

`func NewAwsRdsSnapshotManagerObjectProtectionParams() *AwsRdsSnapshotManagerObjectProtectionParams`

NewAwsRdsSnapshotManagerObjectProtectionParams instantiates a new AwsRdsSnapshotManagerObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRdsSnapshotManagerObjectProtectionParamsWithDefaults

`func NewAwsRdsSnapshotManagerObjectProtectionParamsWithDefaults() *AwsRdsSnapshotManagerObjectProtectionParams`

NewAwsRdsSnapshotManagerObjectProtectionParamsWithDefaults instantiates a new AwsRdsSnapshotManagerObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *AwsRdsSnapshotManagerObjectProtectionParams) GetObjects() []AwsObjectLevelParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AwsRdsSnapshotManagerObjectProtectionParams) GetObjectsOk() (*[]AwsObjectLevelParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AwsRdsSnapshotManagerObjectProtectionParams) SetObjects(v []AwsObjectLevelParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AwsRdsSnapshotManagerObjectProtectionParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


