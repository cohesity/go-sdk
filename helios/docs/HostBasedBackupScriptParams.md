# HostBasedBackupScriptParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | [**ScriptHost**](ScriptHost.md) |  | 
**PreScript** | Pointer to [**CommonPreBackupScriptParams**](CommonPreBackupScriptParams.md) |  | [optional] 
**PostScript** | Pointer to [**CommonPostBackupScriptParams**](CommonPostBackupScriptParams.md) |  | [optional] 

## Methods

### NewHostBasedBackupScriptParams

`func NewHostBasedBackupScriptParams(host ScriptHost, ) *HostBasedBackupScriptParams`

NewHostBasedBackupScriptParams instantiates a new HostBasedBackupScriptParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostBasedBackupScriptParamsWithDefaults

`func NewHostBasedBackupScriptParamsWithDefaults() *HostBasedBackupScriptParams`

NewHostBasedBackupScriptParamsWithDefaults instantiates a new HostBasedBackupScriptParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HostBasedBackupScriptParams) GetHost() ScriptHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostBasedBackupScriptParams) GetHostOk() (*ScriptHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostBasedBackupScriptParams) SetHost(v ScriptHost)`

SetHost sets Host field to given value.


### GetPreScript

`func (o *HostBasedBackupScriptParams) GetPreScript() CommonPreBackupScriptParams`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *HostBasedBackupScriptParams) GetPreScriptOk() (*CommonPreBackupScriptParams, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *HostBasedBackupScriptParams) SetPreScript(v CommonPreBackupScriptParams)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *HostBasedBackupScriptParams) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *HostBasedBackupScriptParams) GetPostScript() CommonPostBackupScriptParams`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *HostBasedBackupScriptParams) GetPostScriptOk() (*CommonPostBackupScriptParams, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *HostBasedBackupScriptParams) SetPostScript(v CommonPostBackupScriptParams)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *HostBasedBackupScriptParams) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


