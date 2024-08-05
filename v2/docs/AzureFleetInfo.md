# AzureFleetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppSecurityGroupName** | **NullableString** | Specifies the appplication security group of Azure dataplane cluster. | 
**AvailabilitySetName** | **NullableString** | Specifies the availabilty set of Azure dataplane cluster. | 
**ClientId** | **NullableString** | Specifies the client Id. | 
**ProximityPlacementGroupName** | **NullableString** | Specifies the proximity placement group of Azure dataplane cluster. | 
**Region** | **NullableString** | Specifies the region of Azure dataplane cluster. | 
**ResourceGroupName** | **NullableString** | Specifies the resource group of Azure dataplane cluster. | 
**SubnetName** | **NullableString** | Specifies the subnet of Azure dataplane cluster. | 
**VnetName** | **NullableString** | Specifies the virtual network of Azure dataplane cluster. | 

## Methods

### NewAzureFleetInfo

`func NewAzureFleetInfo(appSecurityGroupName NullableString, availabilitySetName NullableString, clientId NullableString, proximityPlacementGroupName NullableString, region NullableString, resourceGroupName NullableString, subnetName NullableString, vnetName NullableString, ) *AzureFleetInfo`

NewAzureFleetInfo instantiates a new AzureFleetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureFleetInfoWithDefaults

`func NewAzureFleetInfoWithDefaults() *AzureFleetInfo`

NewAzureFleetInfoWithDefaults instantiates a new AzureFleetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppSecurityGroupName

`func (o *AzureFleetInfo) GetAppSecurityGroupName() string`

GetAppSecurityGroupName returns the AppSecurityGroupName field if non-nil, zero value otherwise.

### GetAppSecurityGroupNameOk

`func (o *AzureFleetInfo) GetAppSecurityGroupNameOk() (*string, bool)`

GetAppSecurityGroupNameOk returns a tuple with the AppSecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSecurityGroupName

`func (o *AzureFleetInfo) SetAppSecurityGroupName(v string)`

SetAppSecurityGroupName sets AppSecurityGroupName field to given value.


### SetAppSecurityGroupNameNil

`func (o *AzureFleetInfo) SetAppSecurityGroupNameNil(b bool)`

 SetAppSecurityGroupNameNil sets the value for AppSecurityGroupName to be an explicit nil

### UnsetAppSecurityGroupName
`func (o *AzureFleetInfo) UnsetAppSecurityGroupName()`

UnsetAppSecurityGroupName ensures that no value is present for AppSecurityGroupName, not even an explicit nil
### GetAvailabilitySetName

`func (o *AzureFleetInfo) GetAvailabilitySetName() string`

GetAvailabilitySetName returns the AvailabilitySetName field if non-nil, zero value otherwise.

### GetAvailabilitySetNameOk

`func (o *AzureFleetInfo) GetAvailabilitySetNameOk() (*string, bool)`

GetAvailabilitySetNameOk returns a tuple with the AvailabilitySetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySetName

`func (o *AzureFleetInfo) SetAvailabilitySetName(v string)`

SetAvailabilitySetName sets AvailabilitySetName field to given value.


### SetAvailabilitySetNameNil

`func (o *AzureFleetInfo) SetAvailabilitySetNameNil(b bool)`

 SetAvailabilitySetNameNil sets the value for AvailabilitySetName to be an explicit nil

### UnsetAvailabilitySetName
`func (o *AzureFleetInfo) UnsetAvailabilitySetName()`

UnsetAvailabilitySetName ensures that no value is present for AvailabilitySetName, not even an explicit nil
### GetClientId

`func (o *AzureFleetInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureFleetInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureFleetInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### SetClientIdNil

`func (o *AzureFleetInfo) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *AzureFleetInfo) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetProximityPlacementGroupName

`func (o *AzureFleetInfo) GetProximityPlacementGroupName() string`

GetProximityPlacementGroupName returns the ProximityPlacementGroupName field if non-nil, zero value otherwise.

### GetProximityPlacementGroupNameOk

`func (o *AzureFleetInfo) GetProximityPlacementGroupNameOk() (*string, bool)`

GetProximityPlacementGroupNameOk returns a tuple with the ProximityPlacementGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityPlacementGroupName

`func (o *AzureFleetInfo) SetProximityPlacementGroupName(v string)`

SetProximityPlacementGroupName sets ProximityPlacementGroupName field to given value.


### SetProximityPlacementGroupNameNil

`func (o *AzureFleetInfo) SetProximityPlacementGroupNameNil(b bool)`

 SetProximityPlacementGroupNameNil sets the value for ProximityPlacementGroupName to be an explicit nil

### UnsetProximityPlacementGroupName
`func (o *AzureFleetInfo) UnsetProximityPlacementGroupName()`

UnsetProximityPlacementGroupName ensures that no value is present for ProximityPlacementGroupName, not even an explicit nil
### GetRegion

`func (o *AzureFleetInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureFleetInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureFleetInfo) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AzureFleetInfo) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AzureFleetInfo) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetResourceGroupName

`func (o *AzureFleetInfo) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureFleetInfo) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureFleetInfo) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### SetResourceGroupNameNil

`func (o *AzureFleetInfo) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *AzureFleetInfo) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetSubnetName

`func (o *AzureFleetInfo) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *AzureFleetInfo) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *AzureFleetInfo) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.


### SetSubnetNameNil

`func (o *AzureFleetInfo) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *AzureFleetInfo) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetVnetName

`func (o *AzureFleetInfo) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *AzureFleetInfo) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *AzureFleetInfo) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.


### SetVnetNameNil

`func (o *AzureFleetInfo) SetVnetNameNil(b bool)`

 SetVnetNameNil sets the value for VnetName to be an explicit nil

### UnsetVnetName
`func (o *AzureFleetInfo) UnsetVnetName()`

UnsetVnetName ensures that no value is present for VnetName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


