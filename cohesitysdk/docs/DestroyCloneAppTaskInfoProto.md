# DestroyCloneAppTaskInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEnv** | Pointer to **NullableInt32** | The application environment. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Finished** | Pointer to **NullableBool** | This will be set to true if the task is complete on the slave. | [optional] 
**TargetEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TargetEntityCredentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 

## Methods

### NewDestroyCloneAppTaskInfoProto

`func NewDestroyCloneAppTaskInfoProto() *DestroyCloneAppTaskInfoProto`

NewDestroyCloneAppTaskInfoProto instantiates a new DestroyCloneAppTaskInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestroyCloneAppTaskInfoProtoWithDefaults

`func NewDestroyCloneAppTaskInfoProtoWithDefaults() *DestroyCloneAppTaskInfoProto`

NewDestroyCloneAppTaskInfoProtoWithDefaults instantiates a new DestroyCloneAppTaskInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEnv

`func (o *DestroyCloneAppTaskInfoProto) GetAppEnv() int32`

GetAppEnv returns the AppEnv field if non-nil, zero value otherwise.

### GetAppEnvOk

`func (o *DestroyCloneAppTaskInfoProto) GetAppEnvOk() (*int32, bool)`

GetAppEnvOk returns a tuple with the AppEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEnv

`func (o *DestroyCloneAppTaskInfoProto) SetAppEnv(v int32)`

SetAppEnv sets AppEnv field to given value.

### HasAppEnv

`func (o *DestroyCloneAppTaskInfoProto) HasAppEnv() bool`

HasAppEnv returns a boolean if a field has been set.

### SetAppEnvNil

`func (o *DestroyCloneAppTaskInfoProto) SetAppEnvNil(b bool)`

 SetAppEnvNil sets the value for AppEnv to be an explicit nil

### UnsetAppEnv
`func (o *DestroyCloneAppTaskInfoProto) UnsetAppEnv()`

UnsetAppEnv ensures that no value is present for AppEnv, not even an explicit nil
### GetError

`func (o *DestroyCloneAppTaskInfoProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DestroyCloneAppTaskInfoProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DestroyCloneAppTaskInfoProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *DestroyCloneAppTaskInfoProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinished

`func (o *DestroyCloneAppTaskInfoProto) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *DestroyCloneAppTaskInfoProto) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *DestroyCloneAppTaskInfoProto) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *DestroyCloneAppTaskInfoProto) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *DestroyCloneAppTaskInfoProto) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *DestroyCloneAppTaskInfoProto) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetTargetEntity

`func (o *DestroyCloneAppTaskInfoProto) GetTargetEntity() EntityProto`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *DestroyCloneAppTaskInfoProto) GetTargetEntityOk() (*EntityProto, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *DestroyCloneAppTaskInfoProto) SetTargetEntity(v EntityProto)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *DestroyCloneAppTaskInfoProto) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.

### GetTargetEntityCredentials

`func (o *DestroyCloneAppTaskInfoProto) GetTargetEntityCredentials() Credentials`

GetTargetEntityCredentials returns the TargetEntityCredentials field if non-nil, zero value otherwise.

### GetTargetEntityCredentialsOk

`func (o *DestroyCloneAppTaskInfoProto) GetTargetEntityCredentialsOk() (*Credentials, bool)`

GetTargetEntityCredentialsOk returns a tuple with the TargetEntityCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntityCredentials

`func (o *DestroyCloneAppTaskInfoProto) SetTargetEntityCredentials(v Credentials)`

SetTargetEntityCredentials sets TargetEntityCredentials field to given value.

### HasTargetEntityCredentials

`func (o *DestroyCloneAppTaskInfoProto) HasTargetEntityCredentials() bool`

HasTargetEntityCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


