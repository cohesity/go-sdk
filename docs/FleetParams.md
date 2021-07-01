# FleetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetNetworkParams** | Pointer to [**FleetNetworkParams**](FleetNetworkParams.md) |  | [optional] 
**FleetSubnetType** | Pointer to **NullableString** | Specifies the subnet type of the fleet. Specifies the type of the fleet subnet. &#39;kCluster&#39; implies same subnet as of Cluster, valid only for Cloud Edition cluster. &#39;kSourceVM&#39; implies same subnet as of source vm. &#39;kCustom&#39; implies the custome subnet. | [optional] 
**FleetTags** | Pointer to [**[]FleetTag**](FleetTag.md) | Specifies the tag information for the fleet. | [optional] 

## Methods

### NewFleetParams

`func NewFleetParams() *FleetParams`

NewFleetParams instantiates a new FleetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetParamsWithDefaults

`func NewFleetParamsWithDefaults() *FleetParams`

NewFleetParamsWithDefaults instantiates a new FleetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetNetworkParams

`func (o *FleetParams) GetFleetNetworkParams() FleetNetworkParams`

GetFleetNetworkParams returns the FleetNetworkParams field if non-nil, zero value otherwise.

### GetFleetNetworkParamsOk

`func (o *FleetParams) GetFleetNetworkParamsOk() (*FleetNetworkParams, bool)`

GetFleetNetworkParamsOk returns a tuple with the FleetNetworkParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetNetworkParams

`func (o *FleetParams) SetFleetNetworkParams(v FleetNetworkParams)`

SetFleetNetworkParams sets FleetNetworkParams field to given value.

### HasFleetNetworkParams

`func (o *FleetParams) HasFleetNetworkParams() bool`

HasFleetNetworkParams returns a boolean if a field has been set.

### GetFleetSubnetType

`func (o *FleetParams) GetFleetSubnetType() string`

GetFleetSubnetType returns the FleetSubnetType field if non-nil, zero value otherwise.

### GetFleetSubnetTypeOk

`func (o *FleetParams) GetFleetSubnetTypeOk() (*string, bool)`

GetFleetSubnetTypeOk returns a tuple with the FleetSubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetSubnetType

`func (o *FleetParams) SetFleetSubnetType(v string)`

SetFleetSubnetType sets FleetSubnetType field to given value.

### HasFleetSubnetType

`func (o *FleetParams) HasFleetSubnetType() bool`

HasFleetSubnetType returns a boolean if a field has been set.

### SetFleetSubnetTypeNil

`func (o *FleetParams) SetFleetSubnetTypeNil(b bool)`

 SetFleetSubnetTypeNil sets the value for FleetSubnetType to be an explicit nil

### UnsetFleetSubnetType
`func (o *FleetParams) UnsetFleetSubnetType()`

UnsetFleetSubnetType ensures that no value is present for FleetSubnetType, not even an explicit nil
### GetFleetTags

`func (o *FleetParams) GetFleetTags() []FleetTag`

GetFleetTags returns the FleetTags field if non-nil, zero value otherwise.

### GetFleetTagsOk

`func (o *FleetParams) GetFleetTagsOk() (*[]FleetTag, bool)`

GetFleetTagsOk returns a tuple with the FleetTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetTags

`func (o *FleetParams) SetFleetTags(v []FleetTag)`

SetFleetTags sets FleetTags field to given value.

### HasFleetTags

`func (o *FleetParams) HasFleetTags() bool`

HasFleetTags returns a boolean if a field has been set.

### SetFleetTagsNil

`func (o *FleetParams) SetFleetTagsNil(b bool)`

 SetFleetTagsNil sets the value for FleetTags to be an explicit nil

### UnsetFleetTags
`func (o *FleetParams) UnsetFleetTags()`

UnsetFleetTags ensures that no value is present for FleetTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


