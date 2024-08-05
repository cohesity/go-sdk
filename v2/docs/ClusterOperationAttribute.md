# ClusterOperationAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the attribute. Following attributres are possible. * &#x60;kUpgradePackageName&#x60; - Indicates the name of the package for   operation types is from enum &#x60;cluster_software_operation_type&#x60; that   represent upgrade related operations. * &#x60;kPatchPackageName&#x60; - Indicates the   name of the package for operation types is from enum   &#x60;cluster_software_operation_type&#x60; that represents patch related   operations. * &#x60;kPackageType&#x60; specifies whether operation is related to upgrade   or patch.  This will have values from enum &#x60;cluster_package_type&#x60;. * &#x60;kPackageSubType&#x60; specifies package sub type.  This will have values   from enum &#x60;cluster_package_sub_type&#x60;.  | 
**Value** | **string** |  | 

## Methods

### NewClusterOperationAttribute

`func NewClusterOperationAttribute(name string, value string, ) *ClusterOperationAttribute`

NewClusterOperationAttribute instantiates a new ClusterOperationAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOperationAttributeWithDefaults

`func NewClusterOperationAttributeWithDefaults() *ClusterOperationAttribute`

NewClusterOperationAttributeWithDefaults instantiates a new ClusterOperationAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterOperationAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterOperationAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterOperationAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ClusterOperationAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClusterOperationAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClusterOperationAttribute) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


