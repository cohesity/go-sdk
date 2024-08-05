# ClusterExpandParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudClusterParams** | Pointer to [**CloudClusterExpandParams**](CloudClusterExpandParams.md) |  | [optional] 
**PhysicalClusterParams** | Pointer to [**PhysicalClusterExpandParams**](PhysicalClusterExpandParams.md) |  | [optional] 
**Type** | **string** | Type of the cluster. &#39;Cloud&#39; indicates cloud edition cluster. &#39;Physical&#39; indicates physical edition cluster. &#39;Virtual&#39; indicates virtual edition cluster. | 

## Methods

### NewClusterExpandParams

`func NewClusterExpandParams(type_ string, ) *ClusterExpandParams`

NewClusterExpandParams instantiates a new ClusterExpandParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterExpandParamsWithDefaults

`func NewClusterExpandParamsWithDefaults() *ClusterExpandParams`

NewClusterExpandParamsWithDefaults instantiates a new ClusterExpandParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudClusterParams

`func (o *ClusterExpandParams) GetCloudClusterParams() CloudClusterExpandParams`

GetCloudClusterParams returns the CloudClusterParams field if non-nil, zero value otherwise.

### GetCloudClusterParamsOk

`func (o *ClusterExpandParams) GetCloudClusterParamsOk() (*CloudClusterExpandParams, bool)`

GetCloudClusterParamsOk returns a tuple with the CloudClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudClusterParams

`func (o *ClusterExpandParams) SetCloudClusterParams(v CloudClusterExpandParams)`

SetCloudClusterParams sets CloudClusterParams field to given value.

### HasCloudClusterParams

`func (o *ClusterExpandParams) HasCloudClusterParams() bool`

HasCloudClusterParams returns a boolean if a field has been set.

### GetPhysicalClusterParams

`func (o *ClusterExpandParams) GetPhysicalClusterParams() PhysicalClusterExpandParams`

GetPhysicalClusterParams returns the PhysicalClusterParams field if non-nil, zero value otherwise.

### GetPhysicalClusterParamsOk

`func (o *ClusterExpandParams) GetPhysicalClusterParamsOk() (*PhysicalClusterExpandParams, bool)`

GetPhysicalClusterParamsOk returns a tuple with the PhysicalClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalClusterParams

`func (o *ClusterExpandParams) SetPhysicalClusterParams(v PhysicalClusterExpandParams)`

SetPhysicalClusterParams sets PhysicalClusterParams field to given value.

### HasPhysicalClusterParams

`func (o *ClusterExpandParams) HasPhysicalClusterParams() bool`

HasPhysicalClusterParams returns a boolean if a field has been set.

### GetType

`func (o *ClusterExpandParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterExpandParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterExpandParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


