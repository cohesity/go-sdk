# RecoverSalesforceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other salesforce objects if one of Object failed to recover. Default value is false. | [optional] 
**Objects** | [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. | 
**RecoverSfdcObjectParams** | Pointer to [**RecoverSfdcObjectParams**](RecoverSfdcObjectParams.md) |  | [optional] 
**RecoverTo** | Pointer to **NullableInt64** | Specifies the id of registered source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location. | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverSalesforceParams

`func NewRecoverSalesforceParams(objects []CommonRecoverObjectSnapshotParams, recoveryAction string, ) *RecoverSalesforceParams`

NewRecoverSalesforceParams instantiates a new RecoverSalesforceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverSalesforceParamsWithDefaults

`func NewRecoverSalesforceParamsWithDefaults() *RecoverSalesforceParams`

NewRecoverSalesforceParamsWithDefaults instantiates a new RecoverSalesforceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverSalesforceParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverSalesforceParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverSalesforceParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverSalesforceParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverSalesforceParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverSalesforceParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetObjects

`func (o *RecoverSalesforceParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverSalesforceParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverSalesforceParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverSalesforceParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverSalesforceParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverSfdcObjectParams

`func (o *RecoverSalesforceParams) GetRecoverSfdcObjectParams() RecoverSfdcObjectParams`

GetRecoverSfdcObjectParams returns the RecoverSfdcObjectParams field if non-nil, zero value otherwise.

### GetRecoverSfdcObjectParamsOk

`func (o *RecoverSalesforceParams) GetRecoverSfdcObjectParamsOk() (*RecoverSfdcObjectParams, bool)`

GetRecoverSfdcObjectParamsOk returns a tuple with the RecoverSfdcObjectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverSfdcObjectParams

`func (o *RecoverSalesforceParams) SetRecoverSfdcObjectParams(v RecoverSfdcObjectParams)`

SetRecoverSfdcObjectParams sets RecoverSfdcObjectParams field to given value.

### HasRecoverSfdcObjectParams

`func (o *RecoverSalesforceParams) HasRecoverSfdcObjectParams() bool`

HasRecoverSfdcObjectParams returns a boolean if a field has been set.

### GetRecoverTo

`func (o *RecoverSalesforceParams) GetRecoverTo() int64`

GetRecoverTo returns the RecoverTo field if non-nil, zero value otherwise.

### GetRecoverToOk

`func (o *RecoverSalesforceParams) GetRecoverToOk() (*int64, bool)`

GetRecoverToOk returns a tuple with the RecoverTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTo

`func (o *RecoverSalesforceParams) SetRecoverTo(v int64)`

SetRecoverTo sets RecoverTo field to given value.

### HasRecoverTo

`func (o *RecoverSalesforceParams) HasRecoverTo() bool`

HasRecoverTo returns a boolean if a field has been set.

### SetRecoverToNil

`func (o *RecoverSalesforceParams) SetRecoverToNil(b bool)`

 SetRecoverToNil sets the value for RecoverTo to be an explicit nil

### UnsetRecoverTo
`func (o *RecoverSalesforceParams) UnsetRecoverTo()`

UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverSalesforceParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverSalesforceParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverSalesforceParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


