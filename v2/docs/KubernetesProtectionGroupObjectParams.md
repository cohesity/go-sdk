# KubernetesProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludePvcs** | Pointer to [**[]KubernetesPvcInfo**](KubernetesPvcInfo.md) | Specifies a list of pvcs to exclude from being protected. This is only applicable to kubernetes. | [optional] 
**Id** | **int64** | Specifies the id of the object. | 
**IncludePvcs** | Pointer to [**[]KubernetesPvcInfo**](KubernetesPvcInfo.md) | Specifies a list of Pvcs to include in the protection. This is only applicable to kubernetes. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 
**QuiesceGroups** | Pointer to [**[]QuiesceGroup**](QuiesceGroup.md) | Specifies the quiescing rules are which specified by the user for doing backup. | [optional] 

## Methods

### NewKubernetesProtectionGroupObjectParams

`func NewKubernetesProtectionGroupObjectParams(id int64, ) *KubernetesProtectionGroupObjectParams`

NewKubernetesProtectionGroupObjectParams instantiates a new KubernetesProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesProtectionGroupObjectParamsWithDefaults

`func NewKubernetesProtectionGroupObjectParamsWithDefaults() *KubernetesProtectionGroupObjectParams`

NewKubernetesProtectionGroupObjectParamsWithDefaults instantiates a new KubernetesProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludePvcs

`func (o *KubernetesProtectionGroupObjectParams) GetExcludePvcs() []KubernetesPvcInfo`

GetExcludePvcs returns the ExcludePvcs field if non-nil, zero value otherwise.

### GetExcludePvcsOk

`func (o *KubernetesProtectionGroupObjectParams) GetExcludePvcsOk() (*[]KubernetesPvcInfo, bool)`

GetExcludePvcsOk returns a tuple with the ExcludePvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePvcs

`func (o *KubernetesProtectionGroupObjectParams) SetExcludePvcs(v []KubernetesPvcInfo)`

SetExcludePvcs sets ExcludePvcs field to given value.

### HasExcludePvcs

`func (o *KubernetesProtectionGroupObjectParams) HasExcludePvcs() bool`

HasExcludePvcs returns a boolean if a field has been set.

### SetExcludePvcsNil

`func (o *KubernetesProtectionGroupObjectParams) SetExcludePvcsNil(b bool)`

 SetExcludePvcsNil sets the value for ExcludePvcs to be an explicit nil

### UnsetExcludePvcs
`func (o *KubernetesProtectionGroupObjectParams) UnsetExcludePvcs()`

UnsetExcludePvcs ensures that no value is present for ExcludePvcs, not even an explicit nil
### GetId

`func (o *KubernetesProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### GetIncludePvcs

`func (o *KubernetesProtectionGroupObjectParams) GetIncludePvcs() []KubernetesPvcInfo`

GetIncludePvcs returns the IncludePvcs field if non-nil, zero value otherwise.

### GetIncludePvcsOk

`func (o *KubernetesProtectionGroupObjectParams) GetIncludePvcsOk() (*[]KubernetesPvcInfo, bool)`

GetIncludePvcsOk returns a tuple with the IncludePvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePvcs

`func (o *KubernetesProtectionGroupObjectParams) SetIncludePvcs(v []KubernetesPvcInfo)`

SetIncludePvcs sets IncludePvcs field to given value.

### HasIncludePvcs

`func (o *KubernetesProtectionGroupObjectParams) HasIncludePvcs() bool`

HasIncludePvcs returns a boolean if a field has been set.

### SetIncludePvcsNil

`func (o *KubernetesProtectionGroupObjectParams) SetIncludePvcsNil(b bool)`

 SetIncludePvcsNil sets the value for IncludePvcs to be an explicit nil

### UnsetIncludePvcs
`func (o *KubernetesProtectionGroupObjectParams) UnsetIncludePvcs()`

UnsetIncludePvcs ensures that no value is present for IncludePvcs, not even an explicit nil
### GetName

`func (o *KubernetesProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KubernetesProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KubernetesProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuiesceGroups

`func (o *KubernetesProtectionGroupObjectParams) GetQuiesceGroups() []QuiesceGroup`

GetQuiesceGroups returns the QuiesceGroups field if non-nil, zero value otherwise.

### GetQuiesceGroupsOk

`func (o *KubernetesProtectionGroupObjectParams) GetQuiesceGroupsOk() (*[]QuiesceGroup, bool)`

GetQuiesceGroupsOk returns a tuple with the QuiesceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesceGroups

`func (o *KubernetesProtectionGroupObjectParams) SetQuiesceGroups(v []QuiesceGroup)`

SetQuiesceGroups sets QuiesceGroups field to given value.

### HasQuiesceGroups

`func (o *KubernetesProtectionGroupObjectParams) HasQuiesceGroups() bool`

HasQuiesceGroups returns a boolean if a field has been set.

### SetQuiesceGroupsNil

`func (o *KubernetesProtectionGroupObjectParams) SetQuiesceGroupsNil(b bool)`

 SetQuiesceGroupsNil sets the value for QuiesceGroups to be an explicit nil

### UnsetQuiesceGroups
`func (o *KubernetesProtectionGroupObjectParams) UnsetQuiesceGroups()`

UnsetQuiesceGroups ensures that no value is present for QuiesceGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


