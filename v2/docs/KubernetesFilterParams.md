# KubernetesFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelCombinationMethod** | Pointer to **NullableString** | Whether to include all the labels or any of them while performing inclusion/exclusion of objects. | [optional] 
**LabelVector** | Pointer to [**[]KubernetesLabel**](KubernetesLabel.md) | Array of Object to represent Label that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude.It will be a two-dimensional array, where each inner array will consist of a key and value representing labels. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections or the union of these labels, as decided by operation parameter. | [optional] 
**Objects** | Pointer to **[]int64** | Array of objects that are to be included. | [optional] 

## Methods

### NewKubernetesFilterParams

`func NewKubernetesFilterParams() *KubernetesFilterParams`

NewKubernetesFilterParams instantiates a new KubernetesFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesFilterParamsWithDefaults

`func NewKubernetesFilterParamsWithDefaults() *KubernetesFilterParams`

NewKubernetesFilterParamsWithDefaults instantiates a new KubernetesFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelCombinationMethod

`func (o *KubernetesFilterParams) GetLabelCombinationMethod() string`

GetLabelCombinationMethod returns the LabelCombinationMethod field if non-nil, zero value otherwise.

### GetLabelCombinationMethodOk

`func (o *KubernetesFilterParams) GetLabelCombinationMethodOk() (*string, bool)`

GetLabelCombinationMethodOk returns a tuple with the LabelCombinationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCombinationMethod

`func (o *KubernetesFilterParams) SetLabelCombinationMethod(v string)`

SetLabelCombinationMethod sets LabelCombinationMethod field to given value.

### HasLabelCombinationMethod

`func (o *KubernetesFilterParams) HasLabelCombinationMethod() bool`

HasLabelCombinationMethod returns a boolean if a field has been set.

### SetLabelCombinationMethodNil

`func (o *KubernetesFilterParams) SetLabelCombinationMethodNil(b bool)`

 SetLabelCombinationMethodNil sets the value for LabelCombinationMethod to be an explicit nil

### UnsetLabelCombinationMethod
`func (o *KubernetesFilterParams) UnsetLabelCombinationMethod()`

UnsetLabelCombinationMethod ensures that no value is present for LabelCombinationMethod, not even an explicit nil
### GetLabelVector

`func (o *KubernetesFilterParams) GetLabelVector() []KubernetesLabel`

GetLabelVector returns the LabelVector field if non-nil, zero value otherwise.

### GetLabelVectorOk

`func (o *KubernetesFilterParams) GetLabelVectorOk() (*[]KubernetesLabel, bool)`

GetLabelVectorOk returns a tuple with the LabelVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelVector

`func (o *KubernetesFilterParams) SetLabelVector(v []KubernetesLabel)`

SetLabelVector sets LabelVector field to given value.

### HasLabelVector

`func (o *KubernetesFilterParams) HasLabelVector() bool`

HasLabelVector returns a boolean if a field has been set.

### SetLabelVectorNil

`func (o *KubernetesFilterParams) SetLabelVectorNil(b bool)`

 SetLabelVectorNil sets the value for LabelVector to be an explicit nil

### UnsetLabelVector
`func (o *KubernetesFilterParams) UnsetLabelVector()`

UnsetLabelVector ensures that no value is present for LabelVector, not even an explicit nil
### GetObjects

`func (o *KubernetesFilterParams) GetObjects() []int64`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *KubernetesFilterParams) GetObjectsOk() (*[]int64, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *KubernetesFilterParams) SetObjects(v []int64)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *KubernetesFilterParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *KubernetesFilterParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *KubernetesFilterParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


