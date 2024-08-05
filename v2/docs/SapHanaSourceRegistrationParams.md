# SapHanaSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | **[]string** | Specifies the IPs/hostnames for the nodes forming the SAP HANA source cluster. | 
**ScriptDir** | **string** | Specifies the absolute path of scripts used to interact with the SAP HANA source. | 
**SourceName** | **string** | Specifies user friendly name for the source. | 

## Methods

### NewSapHanaSourceRegistrationParams

`func NewSapHanaSourceRegistrationParams(hosts []string, scriptDir string, sourceName string, ) *SapHanaSourceRegistrationParams`

NewSapHanaSourceRegistrationParams instantiates a new SapHanaSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSapHanaSourceRegistrationParamsWithDefaults

`func NewSapHanaSourceRegistrationParamsWithDefaults() *SapHanaSourceRegistrationParams`

NewSapHanaSourceRegistrationParamsWithDefaults instantiates a new SapHanaSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *SapHanaSourceRegistrationParams) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *SapHanaSourceRegistrationParams) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *SapHanaSourceRegistrationParams) SetHosts(v []string)`

SetHosts sets Hosts field to given value.


### GetScriptDir

`func (o *SapHanaSourceRegistrationParams) GetScriptDir() string`

GetScriptDir returns the ScriptDir field if non-nil, zero value otherwise.

### GetScriptDirOk

`func (o *SapHanaSourceRegistrationParams) GetScriptDirOk() (*string, bool)`

GetScriptDirOk returns a tuple with the ScriptDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptDir

`func (o *SapHanaSourceRegistrationParams) SetScriptDir(v string)`

SetScriptDir sets ScriptDir field to given value.


### GetSourceName

`func (o *SapHanaSourceRegistrationParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *SapHanaSourceRegistrationParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *SapHanaSourceRegistrationParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


