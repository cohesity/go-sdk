# ObjectMsTeamParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverObject** | [**CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the Microsoft 365 Team recover object info. | 
**MsTeamParam** | [**MsTeamParam**](MsTeamParam.md) | Specifies parameters to recover a Microsoft 365 Team. | 

## Methods

### NewObjectMsTeamParam

`func NewObjectMsTeamParam(recoverObject CommonRecoverObjectSnapshotParams, msTeamParam MsTeamParam, ) *ObjectMsTeamParam`

NewObjectMsTeamParam instantiates a new ObjectMsTeamParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMsTeamParamWithDefaults

`func NewObjectMsTeamParamWithDefaults() *ObjectMsTeamParam`

NewObjectMsTeamParamWithDefaults instantiates a new ObjectMsTeamParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverObject

`func (o *ObjectMsTeamParam) GetRecoverObject() CommonRecoverObjectSnapshotParams`

GetRecoverObject returns the RecoverObject field if non-nil, zero value otherwise.

### GetRecoverObjectOk

`func (o *ObjectMsTeamParam) GetRecoverObjectOk() (*CommonRecoverObjectSnapshotParams, bool)`

GetRecoverObjectOk returns a tuple with the RecoverObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverObject

`func (o *ObjectMsTeamParam) SetRecoverObject(v CommonRecoverObjectSnapshotParams)`

SetRecoverObject sets RecoverObject field to given value.


### GetMsTeamParam

`func (o *ObjectMsTeamParam) GetMsTeamParam() MsTeamParam`

GetMsTeamParam returns the MsTeamParam field if non-nil, zero value otherwise.

### GetMsTeamParamOk

`func (o *ObjectMsTeamParam) GetMsTeamParamOk() (*MsTeamParam, bool)`

GetMsTeamParamOk returns a tuple with the MsTeamParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsTeamParam

`func (o *ObjectMsTeamParam) SetMsTeamParam(v MsTeamParam)`

SetMsTeamParam sets MsTeamParam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


