# AwsFleetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetSubnetType** | Pointer to **NullableString** | Specifies the subnet type of the fleet. Specifies the type of the fleet subnet. &#39;kCluster&#39; implies same subnet as of Cluster, valid only for Cloud Edition cluster. &#39;kSourceVM&#39; implies same subnet as of source vm. &#39;kCustom&#39; implies the custome subnet. | [optional] 
**FleetTags** | Pointer to [**[]FleetTag**](FleetTag.md) | Specifies the tag information for the fleet. | [optional] 
**NetworkParamsList** | Pointer to [**[]FleetNetworkParams**](FleetNetworkParams.md) | Specifies the list of network params for the fleet. | [optional] 

## Methods

### NewAwsFleetParams

`func NewAwsFleetParams() *AwsFleetParams`

NewAwsFleetParams instantiates a new AwsFleetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsFleetParamsWithDefaults

`func NewAwsFleetParamsWithDefaults() *AwsFleetParams`

NewAwsFleetParamsWithDefaults instantiates a new AwsFleetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetSubnetType

`func (o *AwsFleetParams) GetFleetSubnetType() string`

GetFleetSubnetType returns the FleetSubnetType field if non-nil, zero value otherwise.

### GetFleetSubnetTypeOk

`func (o *AwsFleetParams) GetFleetSubnetTypeOk() (*string, bool)`

GetFleetSubnetTypeOk returns a tuple with the FleetSubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetSubnetType

`func (o *AwsFleetParams) SetFleetSubnetType(v string)`

SetFleetSubnetType sets FleetSubnetType field to given value.

### HasFleetSubnetType

`func (o *AwsFleetParams) HasFleetSubnetType() bool`

HasFleetSubnetType returns a boolean if a field has been set.

### SetFleetSubnetTypeNil

`func (o *AwsFleetParams) SetFleetSubnetTypeNil(b bool)`

 SetFleetSubnetTypeNil sets the value for FleetSubnetType to be an explicit nil

### UnsetFleetSubnetType
`func (o *AwsFleetParams) UnsetFleetSubnetType()`

UnsetFleetSubnetType ensures that no value is present for FleetSubnetType, not even an explicit nil
### GetFleetTags

`func (o *AwsFleetParams) GetFleetTags() []FleetTag`

GetFleetTags returns the FleetTags field if non-nil, zero value otherwise.

### GetFleetTagsOk

`func (o *AwsFleetParams) GetFleetTagsOk() (*[]FleetTag, bool)`

GetFleetTagsOk returns a tuple with the FleetTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetTags

`func (o *AwsFleetParams) SetFleetTags(v []FleetTag)`

SetFleetTags sets FleetTags field to given value.

### HasFleetTags

`func (o *AwsFleetParams) HasFleetTags() bool`

HasFleetTags returns a boolean if a field has been set.

### SetFleetTagsNil

`func (o *AwsFleetParams) SetFleetTagsNil(b bool)`

 SetFleetTagsNil sets the value for FleetTags to be an explicit nil

### UnsetFleetTags
`func (o *AwsFleetParams) UnsetFleetTags()`

UnsetFleetTags ensures that no value is present for FleetTags, not even an explicit nil
### GetNetworkParamsList

`func (o *AwsFleetParams) GetNetworkParamsList() []FleetNetworkParams`

GetNetworkParamsList returns the NetworkParamsList field if non-nil, zero value otherwise.

### GetNetworkParamsListOk

`func (o *AwsFleetParams) GetNetworkParamsListOk() (*[]FleetNetworkParams, bool)`

GetNetworkParamsListOk returns a tuple with the NetworkParamsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkParamsList

`func (o *AwsFleetParams) SetNetworkParamsList(v []FleetNetworkParams)`

SetNetworkParamsList sets NetworkParamsList field to given value.

### HasNetworkParamsList

`func (o *AwsFleetParams) HasNetworkParamsList() bool`

HasNetworkParamsList returns a boolean if a field has been set.

### SetNetworkParamsListNil

`func (o *AwsFleetParams) SetNetworkParamsListNil(b bool)`

 SetNetworkParamsListNil sets the value for NetworkParamsList to be an explicit nil

### UnsetNetworkParamsList
`func (o *AwsFleetParams) UnsetNetworkParamsList()`

UnsetNetworkParamsList ensures that no value is present for NetworkParamsList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


