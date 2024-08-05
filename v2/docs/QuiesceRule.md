# QuiesceRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodSelectorLabels** | Pointer to [**[]KubernetesLabel**](KubernetesLabel.md) | Specifies the labels to select a pod. | [optional] 
**PostSnapshotHooks** | [**[]KubernetesHook**](KubernetesHook.md) | Specifies the hooks to be applied after taking snapshot. | 
**PreSnapshotHooks** | [**[]KubernetesHook**](KubernetesHook.md) | Specifies the hooks to be applied before taking snapshot. | 

## Methods

### NewQuiesceRule

`func NewQuiesceRule(postSnapshotHooks []KubernetesHook, preSnapshotHooks []KubernetesHook, ) *QuiesceRule`

NewQuiesceRule instantiates a new QuiesceRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuiesceRuleWithDefaults

`func NewQuiesceRuleWithDefaults() *QuiesceRule`

NewQuiesceRuleWithDefaults instantiates a new QuiesceRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodSelectorLabels

`func (o *QuiesceRule) GetPodSelectorLabels() []KubernetesLabel`

GetPodSelectorLabels returns the PodSelectorLabels field if non-nil, zero value otherwise.

### GetPodSelectorLabelsOk

`func (o *QuiesceRule) GetPodSelectorLabelsOk() (*[]KubernetesLabel, bool)`

GetPodSelectorLabelsOk returns a tuple with the PodSelectorLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSelectorLabels

`func (o *QuiesceRule) SetPodSelectorLabels(v []KubernetesLabel)`

SetPodSelectorLabels sets PodSelectorLabels field to given value.

### HasPodSelectorLabels

`func (o *QuiesceRule) HasPodSelectorLabels() bool`

HasPodSelectorLabels returns a boolean if a field has been set.

### SetPodSelectorLabelsNil

`func (o *QuiesceRule) SetPodSelectorLabelsNil(b bool)`

 SetPodSelectorLabelsNil sets the value for PodSelectorLabels to be an explicit nil

### UnsetPodSelectorLabels
`func (o *QuiesceRule) UnsetPodSelectorLabels()`

UnsetPodSelectorLabels ensures that no value is present for PodSelectorLabels, not even an explicit nil
### GetPostSnapshotHooks

`func (o *QuiesceRule) GetPostSnapshotHooks() []KubernetesHook`

GetPostSnapshotHooks returns the PostSnapshotHooks field if non-nil, zero value otherwise.

### GetPostSnapshotHooksOk

`func (o *QuiesceRule) GetPostSnapshotHooksOk() (*[]KubernetesHook, bool)`

GetPostSnapshotHooksOk returns a tuple with the PostSnapshotHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSnapshotHooks

`func (o *QuiesceRule) SetPostSnapshotHooks(v []KubernetesHook)`

SetPostSnapshotHooks sets PostSnapshotHooks field to given value.


### SetPostSnapshotHooksNil

`func (o *QuiesceRule) SetPostSnapshotHooksNil(b bool)`

 SetPostSnapshotHooksNil sets the value for PostSnapshotHooks to be an explicit nil

### UnsetPostSnapshotHooks
`func (o *QuiesceRule) UnsetPostSnapshotHooks()`

UnsetPostSnapshotHooks ensures that no value is present for PostSnapshotHooks, not even an explicit nil
### GetPreSnapshotHooks

`func (o *QuiesceRule) GetPreSnapshotHooks() []KubernetesHook`

GetPreSnapshotHooks returns the PreSnapshotHooks field if non-nil, zero value otherwise.

### GetPreSnapshotHooksOk

`func (o *QuiesceRule) GetPreSnapshotHooksOk() (*[]KubernetesHook, bool)`

GetPreSnapshotHooksOk returns a tuple with the PreSnapshotHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSnapshotHooks

`func (o *QuiesceRule) SetPreSnapshotHooks(v []KubernetesHook)`

SetPreSnapshotHooks sets PreSnapshotHooks field to given value.


### SetPreSnapshotHooksNil

`func (o *QuiesceRule) SetPreSnapshotHooksNil(b bool)`

 SetPreSnapshotHooksNil sets the value for PreSnapshotHooks to be an explicit nil

### UnsetPreSnapshotHooks
`func (o *QuiesceRule) UnsetPreSnapshotHooks()`

UnsetPreSnapshotHooks ensures that no value is present for PreSnapshotHooks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


