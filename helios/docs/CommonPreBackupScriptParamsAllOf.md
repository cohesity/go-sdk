# CommonPreBackupScriptParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails. | [optional] 

## Methods

### NewCommonPreBackupScriptParamsAllOf

`func NewCommonPreBackupScriptParamsAllOf() *CommonPreBackupScriptParamsAllOf`

NewCommonPreBackupScriptParamsAllOf instantiates a new CommonPreBackupScriptParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonPreBackupScriptParamsAllOfWithDefaults

`func NewCommonPreBackupScriptParamsAllOfWithDefaults() *CommonPreBackupScriptParamsAllOf`

NewCommonPreBackupScriptParamsAllOfWithDefaults instantiates a new CommonPreBackupScriptParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *CommonPreBackupScriptParamsAllOf) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *CommonPreBackupScriptParamsAllOf) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *CommonPreBackupScriptParamsAllOf) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *CommonPreBackupScriptParamsAllOf) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *CommonPreBackupScriptParamsAllOf) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *CommonPreBackupScriptParamsAllOf) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


