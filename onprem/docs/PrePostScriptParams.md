# PrePostScriptParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreScript** | Pointer to [**CommonPreBackupScriptParams**](CommonPreBackupScriptParams.md) |  | [optional] 
**PostScript** | Pointer to [**CommonPostBackupScriptParams**](CommonPostBackupScriptParams.md) |  | [optional] 

## Methods

### NewPrePostScriptParams

`func NewPrePostScriptParams() *PrePostScriptParams`

NewPrePostScriptParams instantiates a new PrePostScriptParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrePostScriptParamsWithDefaults

`func NewPrePostScriptParamsWithDefaults() *PrePostScriptParams`

NewPrePostScriptParamsWithDefaults instantiates a new PrePostScriptParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreScript

`func (o *PrePostScriptParams) GetPreScript() CommonPreBackupScriptParams`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *PrePostScriptParams) GetPreScriptOk() (*CommonPreBackupScriptParams, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *PrePostScriptParams) SetPreScript(v CommonPreBackupScriptParams)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *PrePostScriptParams) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *PrePostScriptParams) GetPostScript() CommonPostBackupScriptParams`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *PrePostScriptParams) GetPostScriptOk() (*CommonPostBackupScriptParams, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *PrePostScriptParams) SetPostScript(v CommonPostBackupScriptParams)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *PrePostScriptParams) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


