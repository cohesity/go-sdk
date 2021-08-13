# ObjectOneDriveParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerInfo** | Pointer to [**CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the OneDrive owner info. | [optional] 
**OneDriveParams** | Pointer to [**[]OneDriveParam**](OneDriveParam.md) | Specifies parameters to recover a OneDrive. | [optional] 

## Methods

### NewObjectOneDriveParam

`func NewObjectOneDriveParam() *ObjectOneDriveParam`

NewObjectOneDriveParam instantiates a new ObjectOneDriveParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectOneDriveParamWithDefaults

`func NewObjectOneDriveParamWithDefaults() *ObjectOneDriveParam`

NewObjectOneDriveParamWithDefaults instantiates a new ObjectOneDriveParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerInfo

`func (o *ObjectOneDriveParam) GetOwnerInfo() CommonRecoverObjectSnapshotParams`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *ObjectOneDriveParam) GetOwnerInfoOk() (*CommonRecoverObjectSnapshotParams, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *ObjectOneDriveParam) SetOwnerInfo(v CommonRecoverObjectSnapshotParams)`

SetOwnerInfo sets OwnerInfo field to given value.

### HasOwnerInfo

`func (o *ObjectOneDriveParam) HasOwnerInfo() bool`

HasOwnerInfo returns a boolean if a field has been set.

### GetOneDriveParams

`func (o *ObjectOneDriveParam) GetOneDriveParams() []OneDriveParam`

GetOneDriveParams returns the OneDriveParams field if non-nil, zero value otherwise.

### GetOneDriveParamsOk

`func (o *ObjectOneDriveParam) GetOneDriveParamsOk() (*[]OneDriveParam, bool)`

GetOneDriveParamsOk returns a tuple with the OneDriveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneDriveParams

`func (o *ObjectOneDriveParam) SetOneDriveParams(v []OneDriveParam)`

SetOneDriveParams sets OneDriveParams field to given value.

### HasOneDriveParams

`func (o *ObjectOneDriveParam) HasOneDriveParams() bool`

HasOneDriveParams returns a boolean if a field has been set.

### SetOneDriveParamsNil

`func (o *ObjectOneDriveParam) SetOneDriveParamsNil(b bool)`

 SetOneDriveParamsNil sets the value for OneDriveParams to be an explicit nil

### UnsetOneDriveParams
`func (o *ObjectOneDriveParam) UnsetOneDriveParams()`

UnsetOneDriveParams ensures that no value is present for OneDriveParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


