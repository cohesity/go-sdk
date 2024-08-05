# ClusterStateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **int64** | Specifies the id of the cluster. | [optional] 
**ClusterIncarnationId** | Pointer to **int64** | Specifies the incarnation id of the cluster. | [optional] 
**Name** | Pointer to **string** | Specifies the name of the cluster. | [optional] 
**Operations** | Pointer to **[]string** | Specifies the operations running on the cluster. * &#x60;None&#x60; indicates that there are no operations currently running on the cluster. * &#x60;Destroy&#x60; indicates that the cluster is currently being destroyed. * &#x60;Clean&#x60; indicates that the cluster is being cleaned. * &#x60;NodeRemoval&#x60; indicates that a node is being removed from the cluster. * &#x60;DiskRemoval&#x60; indicates that a disk is being removed from the cluster. * &#x60;DiskAddition&#x60; indicates that a disk is being added tos the cluster. * &#x60;Upgrade&#x60; indicates to upgrade the software on the cluster. * &#x60;ApplyPatch&#x60; indicates to apply the patch. * &#x60;RevertPatch&#x60; indicates to revert the patch. * &#x60;BaseOSUpgrade&#x60; indicates that the BaseOSUpgrade operation on the cluster is set. * &#x60;ServiceRestart&#x60; indicates that the services on the Cluster are currently being restarted. * &#x60;SystemServiceRestart&#39; indicates that system services on the Cluster are currently being restarted.  | [optional] 
**SoftwareVersion** | Pointer to **string** | Specifies the software version of the cluster. | [optional] 
**SystemApps** | Pointer to [**[]SystemAppStatusParams**](SystemAppStatusParams.md) | Specifies the details of each system app state on the cluster. | [optional] 

## Methods

### NewClusterStateParams

`func NewClusterStateParams() *ClusterStateParams`

NewClusterStateParams instantiates a new ClusterStateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStateParamsWithDefaults

`func NewClusterStateParamsWithDefaults() *ClusterStateParams`

NewClusterStateParamsWithDefaults instantiates a new ClusterStateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterStateParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterStateParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterStateParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterStateParams) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterIncarnationId

`func (o *ClusterStateParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ClusterStateParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ClusterStateParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ClusterStateParams) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### GetName

`func (o *ClusterStateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterStateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterStateParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterStateParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperations

`func (o *ClusterStateParams) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ClusterStateParams) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ClusterStateParams) SetOperations(v []string)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ClusterStateParams) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *ClusterStateParams) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *ClusterStateParams) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *ClusterStateParams) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *ClusterStateParams) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetSystemApps

`func (o *ClusterStateParams) GetSystemApps() []SystemAppStatusParams`

GetSystemApps returns the SystemApps field if non-nil, zero value otherwise.

### GetSystemAppsOk

`func (o *ClusterStateParams) GetSystemAppsOk() (*[]SystemAppStatusParams, bool)`

GetSystemAppsOk returns a tuple with the SystemApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemApps

`func (o *ClusterStateParams) SetSystemApps(v []SystemAppStatusParams)`

SetSystemApps sets SystemApps field to given value.

### HasSystemApps

`func (o *ClusterStateParams) HasSystemApps() bool`

HasSystemApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


