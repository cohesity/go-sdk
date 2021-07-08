# AWSFleetParams1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetSubnetType** | Pointer to **NullableInt32** | Fleet&#39;s subnet type. This field should always be set when specifying fleet params. | [optional] 
**FleetTagVec** | Pointer to [**[]AWSFleetParamsTag**](AWSFleetParamsTag.md) | Optional list of tags to be associated with the fleets. | [optional] 
**NetworkParams** | Pointer to [**AWSFleetParamsNetworkParams**](AWSFleetParamsNetworkParams.md) |  | [optional] 
**NetworkParamsMap** | Pointer to [**[]AWSFleetParamsNetworkParamsMapEntry**](AWSFleetParamsNetworkParamsMapEntry.md) | Map for a region to network params, as network params can be defined per region. Only set when kCustom fleet subnet type is being used. | [optional] 
**NetworkParamsVec** | Pointer to [**[]AWSFleetParamsNetworkParams**](AWSFleetParamsNetworkParams.md) | Network information for the fleet. This will be only set when fleet_subnet_type is kCustom. | [optional] 

## Methods

### NewAWSFleetParams1

`func NewAWSFleetParams1() *AWSFleetParams1`

NewAWSFleetParams1 instantiates a new AWSFleetParams1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSFleetParams1WithDefaults

`func NewAWSFleetParams1WithDefaults() *AWSFleetParams1`

NewAWSFleetParams1WithDefaults instantiates a new AWSFleetParams1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetSubnetType

`func (o *AWSFleetParams1) GetFleetSubnetType() int32`

GetFleetSubnetType returns the FleetSubnetType field if non-nil, zero value otherwise.

### GetFleetSubnetTypeOk

`func (o *AWSFleetParams1) GetFleetSubnetTypeOk() (*int32, bool)`

GetFleetSubnetTypeOk returns a tuple with the FleetSubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetSubnetType

`func (o *AWSFleetParams1) SetFleetSubnetType(v int32)`

SetFleetSubnetType sets FleetSubnetType field to given value.

### HasFleetSubnetType

`func (o *AWSFleetParams1) HasFleetSubnetType() bool`

HasFleetSubnetType returns a boolean if a field has been set.

### SetFleetSubnetTypeNil

`func (o *AWSFleetParams1) SetFleetSubnetTypeNil(b bool)`

 SetFleetSubnetTypeNil sets the value for FleetSubnetType to be an explicit nil

### UnsetFleetSubnetType
`func (o *AWSFleetParams1) UnsetFleetSubnetType()`

UnsetFleetSubnetType ensures that no value is present for FleetSubnetType, not even an explicit nil
### GetFleetTagVec

`func (o *AWSFleetParams1) GetFleetTagVec() []AWSFleetParamsTag`

GetFleetTagVec returns the FleetTagVec field if non-nil, zero value otherwise.

### GetFleetTagVecOk

`func (o *AWSFleetParams1) GetFleetTagVecOk() (*[]AWSFleetParamsTag, bool)`

GetFleetTagVecOk returns a tuple with the FleetTagVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetTagVec

`func (o *AWSFleetParams1) SetFleetTagVec(v []AWSFleetParamsTag)`

SetFleetTagVec sets FleetTagVec field to given value.

### HasFleetTagVec

`func (o *AWSFleetParams1) HasFleetTagVec() bool`

HasFleetTagVec returns a boolean if a field has been set.

### SetFleetTagVecNil

`func (o *AWSFleetParams1) SetFleetTagVecNil(b bool)`

 SetFleetTagVecNil sets the value for FleetTagVec to be an explicit nil

### UnsetFleetTagVec
`func (o *AWSFleetParams1) UnsetFleetTagVec()`

UnsetFleetTagVec ensures that no value is present for FleetTagVec, not even an explicit nil
### GetNetworkParams

`func (o *AWSFleetParams1) GetNetworkParams() AWSFleetParamsNetworkParams`

GetNetworkParams returns the NetworkParams field if non-nil, zero value otherwise.

### GetNetworkParamsOk

`func (o *AWSFleetParams1) GetNetworkParamsOk() (*AWSFleetParamsNetworkParams, bool)`

GetNetworkParamsOk returns a tuple with the NetworkParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkParams

`func (o *AWSFleetParams1) SetNetworkParams(v AWSFleetParamsNetworkParams)`

SetNetworkParams sets NetworkParams field to given value.

### HasNetworkParams

`func (o *AWSFleetParams1) HasNetworkParams() bool`

HasNetworkParams returns a boolean if a field has been set.

### GetNetworkParamsMap

`func (o *AWSFleetParams1) GetNetworkParamsMap() []AWSFleetParamsNetworkParamsMapEntry`

GetNetworkParamsMap returns the NetworkParamsMap field if non-nil, zero value otherwise.

### GetNetworkParamsMapOk

`func (o *AWSFleetParams1) GetNetworkParamsMapOk() (*[]AWSFleetParamsNetworkParamsMapEntry, bool)`

GetNetworkParamsMapOk returns a tuple with the NetworkParamsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkParamsMap

`func (o *AWSFleetParams1) SetNetworkParamsMap(v []AWSFleetParamsNetworkParamsMapEntry)`

SetNetworkParamsMap sets NetworkParamsMap field to given value.

### HasNetworkParamsMap

`func (o *AWSFleetParams1) HasNetworkParamsMap() bool`

HasNetworkParamsMap returns a boolean if a field has been set.

### SetNetworkParamsMapNil

`func (o *AWSFleetParams1) SetNetworkParamsMapNil(b bool)`

 SetNetworkParamsMapNil sets the value for NetworkParamsMap to be an explicit nil

### UnsetNetworkParamsMap
`func (o *AWSFleetParams1) UnsetNetworkParamsMap()`

UnsetNetworkParamsMap ensures that no value is present for NetworkParamsMap, not even an explicit nil
### GetNetworkParamsVec

`func (o *AWSFleetParams1) GetNetworkParamsVec() []AWSFleetParamsNetworkParams`

GetNetworkParamsVec returns the NetworkParamsVec field if non-nil, zero value otherwise.

### GetNetworkParamsVecOk

`func (o *AWSFleetParams1) GetNetworkParamsVecOk() (*[]AWSFleetParamsNetworkParams, bool)`

GetNetworkParamsVecOk returns a tuple with the NetworkParamsVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkParamsVec

`func (o *AWSFleetParams1) SetNetworkParamsVec(v []AWSFleetParamsNetworkParams)`

SetNetworkParamsVec sets NetworkParamsVec field to given value.

### HasNetworkParamsVec

`func (o *AWSFleetParams1) HasNetworkParamsVec() bool`

HasNetworkParamsVec returns a boolean if a field has been set.

### SetNetworkParamsVecNil

`func (o *AWSFleetParams1) SetNetworkParamsVecNil(b bool)`

 SetNetworkParamsVecNil sets the value for NetworkParamsVec to be an explicit nil

### UnsetNetworkParamsVec
`func (o *AWSFleetParams1) UnsetNetworkParamsVec()`

UnsetNetworkParamsVec ensures that no value is present for NetworkParamsVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


