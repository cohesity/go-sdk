# ProtectedObjectInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this accessing given protected object. | [optional] 
**ObjectBackupConfiguration** | Pointer to [**ProtectedObjectBackupConfig**](ProtectedObjectBackupConfig.md) |  | [optional] 
**ProtectionGroupConfigurations** | Pointer to [**[]ProtectedObjectGroupBackupConfig**](ProtectedObjectGroupBackupConfig.md) | Specifies the protection info associated with every object. There can be multiple instances of protection info since the same object can be protected in multiple protection groups. | [optional] 
**LastRun** | Pointer to [**ObjectProtectionRunSummary**](ObjectProtectionRunSummary.md) |  | [optional] 

## Methods

### NewProtectedObjectInfoAllOf

`func NewProtectedObjectInfoAllOf() *ProtectedObjectInfoAllOf`

NewProtectedObjectInfoAllOf instantiates a new ProtectedObjectInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectInfoAllOfWithDefaults

`func NewProtectedObjectInfoAllOfWithDefaults() *ProtectedObjectInfoAllOf`

NewProtectedObjectInfoAllOfWithDefaults instantiates a new ProtectedObjectInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *ProtectedObjectInfoAllOf) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ProtectedObjectInfoAllOf) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ProtectedObjectInfoAllOf) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ProtectedObjectInfoAllOf) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ProtectedObjectInfoAllOf) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ProtectedObjectInfoAllOf) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetObjectBackupConfiguration

`func (o *ProtectedObjectInfoAllOf) GetObjectBackupConfiguration() ProtectedObjectBackupConfig`

GetObjectBackupConfiguration returns the ObjectBackupConfiguration field if non-nil, zero value otherwise.

### GetObjectBackupConfigurationOk

`func (o *ProtectedObjectInfoAllOf) GetObjectBackupConfigurationOk() (*ProtectedObjectBackupConfig, bool)`

GetObjectBackupConfigurationOk returns a tuple with the ObjectBackupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectBackupConfiguration

`func (o *ProtectedObjectInfoAllOf) SetObjectBackupConfiguration(v ProtectedObjectBackupConfig)`

SetObjectBackupConfiguration sets ObjectBackupConfiguration field to given value.

### HasObjectBackupConfiguration

`func (o *ProtectedObjectInfoAllOf) HasObjectBackupConfiguration() bool`

HasObjectBackupConfiguration returns a boolean if a field has been set.

### GetProtectionGroupConfigurations

`func (o *ProtectedObjectInfoAllOf) GetProtectionGroupConfigurations() []ProtectedObjectGroupBackupConfig`

GetProtectionGroupConfigurations returns the ProtectionGroupConfigurations field if non-nil, zero value otherwise.

### GetProtectionGroupConfigurationsOk

`func (o *ProtectedObjectInfoAllOf) GetProtectionGroupConfigurationsOk() (*[]ProtectedObjectGroupBackupConfig, bool)`

GetProtectionGroupConfigurationsOk returns a tuple with the ProtectionGroupConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupConfigurations

`func (o *ProtectedObjectInfoAllOf) SetProtectionGroupConfigurations(v []ProtectedObjectGroupBackupConfig)`

SetProtectionGroupConfigurations sets ProtectionGroupConfigurations field to given value.

### HasProtectionGroupConfigurations

`func (o *ProtectedObjectInfoAllOf) HasProtectionGroupConfigurations() bool`

HasProtectionGroupConfigurations returns a boolean if a field has been set.

### SetProtectionGroupConfigurationsNil

`func (o *ProtectedObjectInfoAllOf) SetProtectionGroupConfigurationsNil(b bool)`

 SetProtectionGroupConfigurationsNil sets the value for ProtectionGroupConfigurations to be an explicit nil

### UnsetProtectionGroupConfigurations
`func (o *ProtectedObjectInfoAllOf) UnsetProtectionGroupConfigurations()`

UnsetProtectionGroupConfigurations ensures that no value is present for ProtectionGroupConfigurations, not even an explicit nil
### GetLastRun

`func (o *ProtectedObjectInfoAllOf) GetLastRun() ObjectProtectionRunSummary`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ProtectedObjectInfoAllOf) GetLastRunOk() (*ObjectProtectionRunSummary, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ProtectedObjectInfoAllOf) SetLastRun(v ObjectProtectionRunSummary)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ProtectedObjectInfoAllOf) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


