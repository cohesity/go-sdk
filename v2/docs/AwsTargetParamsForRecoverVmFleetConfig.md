# AwsTargetParamsForRecoverVmFleetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetNetworkParams** | Pointer to [**NullableFleetConfigFleetNetworkParams**](FleetConfigFleetNetworkParams.md) |  | [optional] 
**FleetSubnetType** | Pointer to **NullableString** | Specifies the subnet type of the fleet. | [optional] 
**FleetTags** | Pointer to [**[]FleetTags**](FleetTags.md) | Specifies the network security groups within above VPC. | [optional] 

## Methods

### NewAwsTargetParamsForRecoverVmFleetConfig

`func NewAwsTargetParamsForRecoverVmFleetConfig() *AwsTargetParamsForRecoverVmFleetConfig`

NewAwsTargetParamsForRecoverVmFleetConfig instantiates a new AwsTargetParamsForRecoverVmFleetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsTargetParamsForRecoverVmFleetConfigWithDefaults

`func NewAwsTargetParamsForRecoverVmFleetConfigWithDefaults() *AwsTargetParamsForRecoverVmFleetConfig`

NewAwsTargetParamsForRecoverVmFleetConfigWithDefaults instantiates a new AwsTargetParamsForRecoverVmFleetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetNetworkParams

`func (o *AwsTargetParamsForRecoverVmFleetConfig) GetFleetNetworkParams() FleetConfigFleetNetworkParams`

GetFleetNetworkParams returns the FleetNetworkParams field if non-nil, zero value otherwise.

### GetFleetNetworkParamsOk

`func (o *AwsTargetParamsForRecoverVmFleetConfig) GetFleetNetworkParamsOk() (*FleetConfigFleetNetworkParams, bool)`

GetFleetNetworkParamsOk returns a tuple with the FleetNetworkParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetNetworkParams

`func (o *AwsTargetParamsForRecoverVmFleetConfig) SetFleetNetworkParams(v FleetConfigFleetNetworkParams)`

SetFleetNetworkParams sets FleetNetworkParams field to given value.

### HasFleetNetworkParams

`func (o *AwsTargetParamsForRecoverVmFleetConfig) HasFleetNetworkParams() bool`

HasFleetNetworkParams returns a boolean if a field has been set.

### SetFleetNetworkParamsNil

`func (o *AwsTargetParamsForRecoverVmFleetConfig) SetFleetNetworkParamsNil(b bool)`

 SetFleetNetworkParamsNil sets the value for FleetNetworkParams to be an explicit nil

### UnsetFleetNetworkParams
`func (o *AwsTargetParamsForRecoverVmFleetConfig) UnsetFleetNetworkParams()`

UnsetFleetNetworkParams ensures that no value is present for FleetNetworkParams, not even an explicit nil
### GetFleetSubnetType

`func (o *AwsTargetParamsForRecoverVmFleetConfig) GetFleetSubnetType() string`

GetFleetSubnetType returns the FleetSubnetType field if non-nil, zero value otherwise.

### GetFleetSubnetTypeOk

`func (o *AwsTargetParamsForRecoverVmFleetConfig) GetFleetSubnetTypeOk() (*string, bool)`

GetFleetSubnetTypeOk returns a tuple with the FleetSubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetSubnetType

`func (o *AwsTargetParamsForRecoverVmFleetConfig) SetFleetSubnetType(v string)`

SetFleetSubnetType sets FleetSubnetType field to given value.

### HasFleetSubnetType

`func (o *AwsTargetParamsForRecoverVmFleetConfig) HasFleetSubnetType() bool`

HasFleetSubnetType returns a boolean if a field has been set.

### SetFleetSubnetTypeNil

`func (o *AwsTargetParamsForRecoverVmFleetConfig) SetFleetSubnetTypeNil(b bool)`

 SetFleetSubnetTypeNil sets the value for FleetSubnetType to be an explicit nil

### UnsetFleetSubnetType
`func (o *AwsTargetParamsForRecoverVmFleetConfig) UnsetFleetSubnetType()`

UnsetFleetSubnetType ensures that no value is present for FleetSubnetType, not even an explicit nil
### GetFleetTags

`func (o *AwsTargetParamsForRecoverVmFleetConfig) GetFleetTags() []FleetTags`

GetFleetTags returns the FleetTags field if non-nil, zero value otherwise.

### GetFleetTagsOk

`func (o *AwsTargetParamsForRecoverVmFleetConfig) GetFleetTagsOk() (*[]FleetTags, bool)`

GetFleetTagsOk returns a tuple with the FleetTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetTags

`func (o *AwsTargetParamsForRecoverVmFleetConfig) SetFleetTags(v []FleetTags)`

SetFleetTags sets FleetTags field to given value.

### HasFleetTags

`func (o *AwsTargetParamsForRecoverVmFleetConfig) HasFleetTags() bool`

HasFleetTags returns a boolean if a field has been set.

### SetFleetTagsNil

`func (o *AwsTargetParamsForRecoverVmFleetConfig) SetFleetTagsNil(b bool)`

 SetFleetTagsNil sets the value for FleetTags to be an explicit nil

### UnsetFleetTags
`func (o *AwsTargetParamsForRecoverVmFleetConfig) UnsetFleetTags()`

UnsetFleetTags ensures that no value is present for FleetTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


