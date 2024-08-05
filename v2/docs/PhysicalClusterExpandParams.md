# PhysicalClusterExpandParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisRackConfigs** | Pointer to [**[]ChassisRackConfigParams**](ChassisRackConfigParams.md) | Chassis serial to rack id mapping configuration. | [optional] 
**NodeConfigs** | [**[]PhysicalNodeConfigParams**](PhysicalNodeConfigParams.md) | Configuration of the nodes. | 
**Vips** | Pointer to **[]string** | Virtual IPs to add to the cluster. | [optional] 

## Methods

### NewPhysicalClusterExpandParams

`func NewPhysicalClusterExpandParams(nodeConfigs []PhysicalNodeConfigParams, ) *PhysicalClusterExpandParams`

NewPhysicalClusterExpandParams instantiates a new PhysicalClusterExpandParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalClusterExpandParamsWithDefaults

`func NewPhysicalClusterExpandParamsWithDefaults() *PhysicalClusterExpandParams`

NewPhysicalClusterExpandParamsWithDefaults instantiates a new PhysicalClusterExpandParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisRackConfigs

`func (o *PhysicalClusterExpandParams) GetChassisRackConfigs() []ChassisRackConfigParams`

GetChassisRackConfigs returns the ChassisRackConfigs field if non-nil, zero value otherwise.

### GetChassisRackConfigsOk

`func (o *PhysicalClusterExpandParams) GetChassisRackConfigsOk() (*[]ChassisRackConfigParams, bool)`

GetChassisRackConfigsOk returns a tuple with the ChassisRackConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisRackConfigs

`func (o *PhysicalClusterExpandParams) SetChassisRackConfigs(v []ChassisRackConfigParams)`

SetChassisRackConfigs sets ChassisRackConfigs field to given value.

### HasChassisRackConfigs

`func (o *PhysicalClusterExpandParams) HasChassisRackConfigs() bool`

HasChassisRackConfigs returns a boolean if a field has been set.

### GetNodeConfigs

`func (o *PhysicalClusterExpandParams) GetNodeConfigs() []PhysicalNodeConfigParams`

GetNodeConfigs returns the NodeConfigs field if non-nil, zero value otherwise.

### GetNodeConfigsOk

`func (o *PhysicalClusterExpandParams) GetNodeConfigsOk() (*[]PhysicalNodeConfigParams, bool)`

GetNodeConfigsOk returns a tuple with the NodeConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConfigs

`func (o *PhysicalClusterExpandParams) SetNodeConfigs(v []PhysicalNodeConfigParams)`

SetNodeConfigs sets NodeConfigs field to given value.


### GetVips

`func (o *PhysicalClusterExpandParams) GetVips() []string`

GetVips returns the Vips field if non-nil, zero value otherwise.

### GetVipsOk

`func (o *PhysicalClusterExpandParams) GetVipsOk() (*[]string, bool)`

GetVipsOk returns a tuple with the Vips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVips

`func (o *PhysicalClusterExpandParams) SetVips(v []string)`

SetVips sets Vips field to given value.

### HasVips

`func (o *PhysicalClusterExpandParams) HasVips() bool`

HasVips returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


