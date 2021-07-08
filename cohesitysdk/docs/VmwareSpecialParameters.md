# VmwareSpecialParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationParameters** | Pointer to [**ApplicationParameters**](ApplicationParameters.md) |  | [optional] 
**ExcludedDisks** | Pointer to [**[]DiskUnit**](DiskUnit.md) | Specifies the list of Disks to be excluded from backing up. These disks are excluded from all Protection Sources in the Protection Job. | [optional] 
**VmCredentials** | Pointer to [**NullableCredentials**](Credentials.md) | Specifies the administrator credentials to log in to the guest Windows system of a VM that hosts the Microsoft Exchange Server. If truncateExchangeLog is set to true and the specified source is a VM, administrator credentials to log in to the guest Windows system of the VM must be provided to truncate the logs. This field is only applicable to Sources in the kVMware environment. | [optional] 

## Methods

### NewVmwareSpecialParameters

`func NewVmwareSpecialParameters() *VmwareSpecialParameters`

NewVmwareSpecialParameters instantiates a new VmwareSpecialParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareSpecialParametersWithDefaults

`func NewVmwareSpecialParametersWithDefaults() *VmwareSpecialParameters`

NewVmwareSpecialParametersWithDefaults instantiates a new VmwareSpecialParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationParameters

`func (o *VmwareSpecialParameters) GetApplicationParameters() ApplicationParameters`

GetApplicationParameters returns the ApplicationParameters field if non-nil, zero value otherwise.

### GetApplicationParametersOk

`func (o *VmwareSpecialParameters) GetApplicationParametersOk() (*ApplicationParameters, bool)`

GetApplicationParametersOk returns a tuple with the ApplicationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationParameters

`func (o *VmwareSpecialParameters) SetApplicationParameters(v ApplicationParameters)`

SetApplicationParameters sets ApplicationParameters field to given value.

### HasApplicationParameters

`func (o *VmwareSpecialParameters) HasApplicationParameters() bool`

HasApplicationParameters returns a boolean if a field has been set.

### GetExcludedDisks

`func (o *VmwareSpecialParameters) GetExcludedDisks() []DiskUnit`

GetExcludedDisks returns the ExcludedDisks field if non-nil, zero value otherwise.

### GetExcludedDisksOk

`func (o *VmwareSpecialParameters) GetExcludedDisksOk() (*[]DiskUnit, bool)`

GetExcludedDisksOk returns a tuple with the ExcludedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDisks

`func (o *VmwareSpecialParameters) SetExcludedDisks(v []DiskUnit)`

SetExcludedDisks sets ExcludedDisks field to given value.

### HasExcludedDisks

`func (o *VmwareSpecialParameters) HasExcludedDisks() bool`

HasExcludedDisks returns a boolean if a field has been set.

### SetExcludedDisksNil

`func (o *VmwareSpecialParameters) SetExcludedDisksNil(b bool)`

 SetExcludedDisksNil sets the value for ExcludedDisks to be an explicit nil

### UnsetExcludedDisks
`func (o *VmwareSpecialParameters) UnsetExcludedDisks()`

UnsetExcludedDisks ensures that no value is present for ExcludedDisks, not even an explicit nil
### GetVmCredentials

`func (o *VmwareSpecialParameters) GetVmCredentials() Credentials`

GetVmCredentials returns the VmCredentials field if non-nil, zero value otherwise.

### GetVmCredentialsOk

`func (o *VmwareSpecialParameters) GetVmCredentialsOk() (*Credentials, bool)`

GetVmCredentialsOk returns a tuple with the VmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCredentials

`func (o *VmwareSpecialParameters) SetVmCredentials(v Credentials)`

SetVmCredentials sets VmCredentials field to given value.

### HasVmCredentials

`func (o *VmwareSpecialParameters) HasVmCredentials() bool`

HasVmCredentials returns a boolean if a field has been set.

### SetVmCredentialsNil

`func (o *VmwareSpecialParameters) SetVmCredentialsNil(b bool)`

 SetVmCredentialsNil sets the value for VmCredentials to be an explicit nil

### UnsetVmCredentials
`func (o *VmwareSpecialParameters) UnsetVmCredentials()`

UnsetVmCredentials ensures that no value is present for VmCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


