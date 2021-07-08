# GcpCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientEmailAddress** | Pointer to **NullableString** | Specifies Client email address associated with the service account. | [optional] 
**ClientPrivateKey** | Pointer to **NullableString** | Specifies Client private associated with the service account. | [optional] 
**GcpType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kIAMUser&#39; if the environment is kGCP. Specifies the type of a GCP source entity. &#39;kIAMUser&#39; indicates a unique user within a GCP account. &#39;kProject&#39; represents compute resources and storage. &#39;kRegion&#39; indicates a geographical region in the global infrastructure. &#39;kAvailabilityZone&#39; indicates an availability zone within a region. &#39;kVirtualMachine&#39; indicates a Virtual Machine running in GCP environment. &#39;kVPC&#39; indicates a virtual private cloud (VPC) network within GCP. &#39;kSubnet&#39; indicates a subnet inside the VPC. &#39;kNetworkSecurityGroup&#39; represents a network security group. &#39;kInstanceType&#39; represents various machine types. &#39;kLabel&#39; represents a label present on the instances. &#39;kMetaData&#39; represents a custom metadata present on instances. &#39;kTag&#39; represents a network tag on instances. &#39;kVPCConnector&#39; represents a VPC connector used for serverless VPC access. | [optional] 
**ProjectId** | Pointer to **NullableString** | Specifies Id of the project associated with Google cloud account. | [optional] 
**VpcNetwork** | Pointer to **NullableString** | Specifies the VPC Network to deploy proxy VMs. | [optional] 
**VpcSubnetwork** | Pointer to **NullableString** | Specifies the subnetwork to deploy proxy VMs. | [optional] 

## Methods

### NewGcpCredentials

`func NewGcpCredentials() *GcpCredentials`

NewGcpCredentials instantiates a new GcpCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpCredentialsWithDefaults

`func NewGcpCredentialsWithDefaults() *GcpCredentials`

NewGcpCredentialsWithDefaults instantiates a new GcpCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientEmailAddress

`func (o *GcpCredentials) GetClientEmailAddress() string`

GetClientEmailAddress returns the ClientEmailAddress field if non-nil, zero value otherwise.

### GetClientEmailAddressOk

`func (o *GcpCredentials) GetClientEmailAddressOk() (*string, bool)`

GetClientEmailAddressOk returns a tuple with the ClientEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmailAddress

`func (o *GcpCredentials) SetClientEmailAddress(v string)`

SetClientEmailAddress sets ClientEmailAddress field to given value.

### HasClientEmailAddress

`func (o *GcpCredentials) HasClientEmailAddress() bool`

HasClientEmailAddress returns a boolean if a field has been set.

### SetClientEmailAddressNil

`func (o *GcpCredentials) SetClientEmailAddressNil(b bool)`

 SetClientEmailAddressNil sets the value for ClientEmailAddress to be an explicit nil

### UnsetClientEmailAddress
`func (o *GcpCredentials) UnsetClientEmailAddress()`

UnsetClientEmailAddress ensures that no value is present for ClientEmailAddress, not even an explicit nil
### GetClientPrivateKey

`func (o *GcpCredentials) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *GcpCredentials) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *GcpCredentials) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *GcpCredentials) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### SetClientPrivateKeyNil

`func (o *GcpCredentials) SetClientPrivateKeyNil(b bool)`

 SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil

### UnsetClientPrivateKey
`func (o *GcpCredentials) UnsetClientPrivateKey()`

UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil
### GetGcpType

`func (o *GcpCredentials) GetGcpType() string`

GetGcpType returns the GcpType field if non-nil, zero value otherwise.

### GetGcpTypeOk

`func (o *GcpCredentials) GetGcpTypeOk() (*string, bool)`

GetGcpTypeOk returns a tuple with the GcpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpType

`func (o *GcpCredentials) SetGcpType(v string)`

SetGcpType sets GcpType field to given value.

### HasGcpType

`func (o *GcpCredentials) HasGcpType() bool`

HasGcpType returns a boolean if a field has been set.

### SetGcpTypeNil

`func (o *GcpCredentials) SetGcpTypeNil(b bool)`

 SetGcpTypeNil sets the value for GcpType to be an explicit nil

### UnsetGcpType
`func (o *GcpCredentials) UnsetGcpType()`

UnsetGcpType ensures that no value is present for GcpType, not even an explicit nil
### GetProjectId

`func (o *GcpCredentials) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpCredentials) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpCredentials) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GcpCredentials) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *GcpCredentials) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *GcpCredentials) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetVpcNetwork

`func (o *GcpCredentials) GetVpcNetwork() string`

GetVpcNetwork returns the VpcNetwork field if non-nil, zero value otherwise.

### GetVpcNetworkOk

`func (o *GcpCredentials) GetVpcNetworkOk() (*string, bool)`

GetVpcNetworkOk returns a tuple with the VpcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNetwork

`func (o *GcpCredentials) SetVpcNetwork(v string)`

SetVpcNetwork sets VpcNetwork field to given value.

### HasVpcNetwork

`func (o *GcpCredentials) HasVpcNetwork() bool`

HasVpcNetwork returns a boolean if a field has been set.

### SetVpcNetworkNil

`func (o *GcpCredentials) SetVpcNetworkNil(b bool)`

 SetVpcNetworkNil sets the value for VpcNetwork to be an explicit nil

### UnsetVpcNetwork
`func (o *GcpCredentials) UnsetVpcNetwork()`

UnsetVpcNetwork ensures that no value is present for VpcNetwork, not even an explicit nil
### GetVpcSubnetwork

`func (o *GcpCredentials) GetVpcSubnetwork() string`

GetVpcSubnetwork returns the VpcSubnetwork field if non-nil, zero value otherwise.

### GetVpcSubnetworkOk

`func (o *GcpCredentials) GetVpcSubnetworkOk() (*string, bool)`

GetVpcSubnetworkOk returns a tuple with the VpcSubnetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcSubnetwork

`func (o *GcpCredentials) SetVpcSubnetwork(v string)`

SetVpcSubnetwork sets VpcSubnetwork field to given value.

### HasVpcSubnetwork

`func (o *GcpCredentials) HasVpcSubnetwork() bool`

HasVpcSubnetwork returns a boolean if a field has been set.

### SetVpcSubnetworkNil

`func (o *GcpCredentials) SetVpcSubnetworkNil(b bool)`

 SetVpcSubnetworkNil sets the value for VpcSubnetwork to be an explicit nil

### UnsetVpcSubnetwork
`func (o *GcpCredentials) UnsetVpcSubnetwork()`

UnsetVpcSubnetwork ensures that no value is present for VpcSubnetwork, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


