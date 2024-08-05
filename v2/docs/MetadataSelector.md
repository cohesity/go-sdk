# MetadataSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitySelector** | Pointer to [**EntitySelector**](EntitySelector.md) |  | [optional] 
**IncludeEnvironments** | Pointer to **[]string** | Contains the environments info that will be exported. If this is empty, all the environments info will be exported.. | [optional] 
**IncludeJobDescriptions** | Pointer to **bool** | If true, exports the job descriptions. This is automatically true if the include_restore_tasks is true. | [optional] 
**IncludePolicies** | Pointer to **bool** | If true, exports the accessible policies. | [optional] 
**IncludeRestoreTasks** | Pointer to **bool** | If true, exports the qualified restore tasks. | [optional] 
**RunSelector** | Pointer to [**ProtectionRunSelector**](ProtectionRunSelector.md) |  | [optional] 
**TenantId** | Pointer to **string** | Denotes the tenant whose metadata needs to be exported | [optional] 

## Methods

### NewMetadataSelector

`func NewMetadataSelector() *MetadataSelector`

NewMetadataSelector instantiates a new MetadataSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataSelectorWithDefaults

`func NewMetadataSelectorWithDefaults() *MetadataSelector`

NewMetadataSelectorWithDefaults instantiates a new MetadataSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitySelector

`func (o *MetadataSelector) GetEntitySelector() EntitySelector`

GetEntitySelector returns the EntitySelector field if non-nil, zero value otherwise.

### GetEntitySelectorOk

`func (o *MetadataSelector) GetEntitySelectorOk() (*EntitySelector, bool)`

GetEntitySelectorOk returns a tuple with the EntitySelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelector

`func (o *MetadataSelector) SetEntitySelector(v EntitySelector)`

SetEntitySelector sets EntitySelector field to given value.

### HasEntitySelector

`func (o *MetadataSelector) HasEntitySelector() bool`

HasEntitySelector returns a boolean if a field has been set.

### GetIncludeEnvironments

`func (o *MetadataSelector) GetIncludeEnvironments() []string`

GetIncludeEnvironments returns the IncludeEnvironments field if non-nil, zero value otherwise.

### GetIncludeEnvironmentsOk

`func (o *MetadataSelector) GetIncludeEnvironmentsOk() (*[]string, bool)`

GetIncludeEnvironmentsOk returns a tuple with the IncludeEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEnvironments

`func (o *MetadataSelector) SetIncludeEnvironments(v []string)`

SetIncludeEnvironments sets IncludeEnvironments field to given value.

### HasIncludeEnvironments

`func (o *MetadataSelector) HasIncludeEnvironments() bool`

HasIncludeEnvironments returns a boolean if a field has been set.

### GetIncludeJobDescriptions

`func (o *MetadataSelector) GetIncludeJobDescriptions() bool`

GetIncludeJobDescriptions returns the IncludeJobDescriptions field if non-nil, zero value otherwise.

### GetIncludeJobDescriptionsOk

`func (o *MetadataSelector) GetIncludeJobDescriptionsOk() (*bool, bool)`

GetIncludeJobDescriptionsOk returns a tuple with the IncludeJobDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJobDescriptions

`func (o *MetadataSelector) SetIncludeJobDescriptions(v bool)`

SetIncludeJobDescriptions sets IncludeJobDescriptions field to given value.

### HasIncludeJobDescriptions

`func (o *MetadataSelector) HasIncludeJobDescriptions() bool`

HasIncludeJobDescriptions returns a boolean if a field has been set.

### GetIncludePolicies

`func (o *MetadataSelector) GetIncludePolicies() bool`

GetIncludePolicies returns the IncludePolicies field if non-nil, zero value otherwise.

### GetIncludePoliciesOk

`func (o *MetadataSelector) GetIncludePoliciesOk() (*bool, bool)`

GetIncludePoliciesOk returns a tuple with the IncludePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePolicies

`func (o *MetadataSelector) SetIncludePolicies(v bool)`

SetIncludePolicies sets IncludePolicies field to given value.

### HasIncludePolicies

`func (o *MetadataSelector) HasIncludePolicies() bool`

HasIncludePolicies returns a boolean if a field has been set.

### GetIncludeRestoreTasks

`func (o *MetadataSelector) GetIncludeRestoreTasks() bool`

GetIncludeRestoreTasks returns the IncludeRestoreTasks field if non-nil, zero value otherwise.

### GetIncludeRestoreTasksOk

`func (o *MetadataSelector) GetIncludeRestoreTasksOk() (*bool, bool)`

GetIncludeRestoreTasksOk returns a tuple with the IncludeRestoreTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRestoreTasks

`func (o *MetadataSelector) SetIncludeRestoreTasks(v bool)`

SetIncludeRestoreTasks sets IncludeRestoreTasks field to given value.

### HasIncludeRestoreTasks

`func (o *MetadataSelector) HasIncludeRestoreTasks() bool`

HasIncludeRestoreTasks returns a boolean if a field has been set.

### GetRunSelector

`func (o *MetadataSelector) GetRunSelector() ProtectionRunSelector`

GetRunSelector returns the RunSelector field if non-nil, zero value otherwise.

### GetRunSelectorOk

`func (o *MetadataSelector) GetRunSelectorOk() (*ProtectionRunSelector, bool)`

GetRunSelectorOk returns a tuple with the RunSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunSelector

`func (o *MetadataSelector) SetRunSelector(v ProtectionRunSelector)`

SetRunSelector sets RunSelector field to given value.

### HasRunSelector

`func (o *MetadataSelector) HasRunSelector() bool`

HasRunSelector returns a boolean if a field has been set.

### GetTenantId

`func (o *MetadataSelector) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MetadataSelector) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MetadataSelector) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MetadataSelector) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


