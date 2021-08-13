# KubernetesFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to **[]int64** | Array of objects that are to be included. | [optional] 
**LabelIds** | Pointer to **[][]int64** | Array of Arrays of Label Ids that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude. Optionally specify a list of items to include in filter by listing Source Ids of Labels in this two dimensional array. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. For example a Namespace is selected to be protected but you want to exclude all the &#39;Employees:Former&#39; items in the &#39;Location:East&#39; and &#39;Location:West&#39; but keep all the items for &#39;Employees:Former&#39; in the South which are also stored in this Namespace, by specifying the following source id array: [ [1000, 2221], [1000, 3031] ], where 1000 is the &#39;Employee:Former&#39; Label id, 2221 is the &#39;Location:East&#39; Label id and 3031 is the &#39;West&#39; Label id. The first inner array [1000, 2221] produces a list of items that are both labeled with &#39;Employees:Former&#39; and &#39;Location:East&#39; (an intersection). The second inner array [1000, 3031] produces a list of items that are both labeled with &#39;Employees:Former&#39; and &#39;Location:West&#39; (an intersection). The outer array combines the items from the two inner arrays. The list of resulting items, when combined with isExclude&#x3D;true, are then excluded from being protected this Job. | [optional] 

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
### GetLabelIds

`func (o *KubernetesFilterParams) GetLabelIds() [][]int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *KubernetesFilterParams) GetLabelIdsOk() (*[][]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *KubernetesFilterParams) SetLabelIds(v [][]int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *KubernetesFilterParams) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### SetLabelIdsNil

`func (o *KubernetesFilterParams) SetLabelIdsNil(b bool)`

 SetLabelIdsNil sets the value for LabelIds to be an explicit nil

### UnsetLabelIds
`func (o *KubernetesFilterParams) UnsetLabelIds()`

UnsetLabelIds ensures that no value is present for LabelIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


