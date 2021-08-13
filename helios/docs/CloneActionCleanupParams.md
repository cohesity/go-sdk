# CloneActionCleanupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CleanupType** | **NullableString** | Specifies the type of cleanup to be done. | 
**PowerOffVmParams** | Pointer to [**CloneActionCleanupPowerOffVmParams**](CloneActionCleanupPowerOffVmParams.md) |  | [optional] 
**CloudResourcesCleanupParams** | Pointer to [**CloneActionCleanupCloudResourcesCleanupParams**](CloneActionCleanupCloudResourcesCleanupParams.md) |  | [optional] 

## Methods

### NewCloneActionCleanupParams

`func NewCloneActionCleanupParams(cleanupType NullableString, ) *CloneActionCleanupParams`

NewCloneActionCleanupParams instantiates a new CloneActionCleanupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneActionCleanupParamsWithDefaults

`func NewCloneActionCleanupParamsWithDefaults() *CloneActionCleanupParams`

NewCloneActionCleanupParamsWithDefaults instantiates a new CloneActionCleanupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanupType

`func (o *CloneActionCleanupParams) GetCleanupType() string`

GetCleanupType returns the CleanupType field if non-nil, zero value otherwise.

### GetCleanupTypeOk

`func (o *CloneActionCleanupParams) GetCleanupTypeOk() (*string, bool)`

GetCleanupTypeOk returns a tuple with the CleanupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupType

`func (o *CloneActionCleanupParams) SetCleanupType(v string)`

SetCleanupType sets CleanupType field to given value.


### SetCleanupTypeNil

`func (o *CloneActionCleanupParams) SetCleanupTypeNil(b bool)`

 SetCleanupTypeNil sets the value for CleanupType to be an explicit nil

### UnsetCleanupType
`func (o *CloneActionCleanupParams) UnsetCleanupType()`

UnsetCleanupType ensures that no value is present for CleanupType, not even an explicit nil
### GetPowerOffVmParams

`func (o *CloneActionCleanupParams) GetPowerOffVmParams() CloneActionCleanupPowerOffVmParams`

GetPowerOffVmParams returns the PowerOffVmParams field if non-nil, zero value otherwise.

### GetPowerOffVmParamsOk

`func (o *CloneActionCleanupParams) GetPowerOffVmParamsOk() (*CloneActionCleanupPowerOffVmParams, bool)`

GetPowerOffVmParamsOk returns a tuple with the PowerOffVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffVmParams

`func (o *CloneActionCleanupParams) SetPowerOffVmParams(v CloneActionCleanupPowerOffVmParams)`

SetPowerOffVmParams sets PowerOffVmParams field to given value.

### HasPowerOffVmParams

`func (o *CloneActionCleanupParams) HasPowerOffVmParams() bool`

HasPowerOffVmParams returns a boolean if a field has been set.

### GetCloudResourcesCleanupParams

`func (o *CloneActionCleanupParams) GetCloudResourcesCleanupParams() CloneActionCleanupCloudResourcesCleanupParams`

GetCloudResourcesCleanupParams returns the CloudResourcesCleanupParams field if non-nil, zero value otherwise.

### GetCloudResourcesCleanupParamsOk

`func (o *CloneActionCleanupParams) GetCloudResourcesCleanupParamsOk() (*CloneActionCleanupCloudResourcesCleanupParams, bool)`

GetCloudResourcesCleanupParamsOk returns a tuple with the CloudResourcesCleanupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudResourcesCleanupParams

`func (o *CloneActionCleanupParams) SetCloudResourcesCleanupParams(v CloneActionCleanupCloudResourcesCleanupParams)`

SetCloudResourcesCleanupParams sets CloudResourcesCleanupParams field to given value.

### HasCloudResourcesCleanupParams

`func (o *CloneActionCleanupParams) HasCloudResourcesCleanupParams() bool`

HasCloudResourcesCleanupParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


