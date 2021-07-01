# AwsProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **NullableString** | Specifies Access key of the AWS account. | [optional] 
**AmazonResourceName** | Pointer to **NullableString** | Specifies Amazon Resource Name (owner ID) of the IAM user, act as an unique identifier of as AWS entity. | [optional] 
**AuthMethod** | Pointer to **NullableString** | Specifies the authentication method to be used for API calls. Specifies the authentication method to be used for API calls. &#39;kUseIAMUser&#39; indicates a user based authentication. &#39;kUseIAMRole&#39; indicates a role based authentication, used only for AWS CE. | [optional] 
**AwsFleetParams** | Pointer to [**AwsFleetParams**](AwsFleetParams.md) |  | [optional] 
**AwsType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kIAMUser&#39; if the environment is kAWS. Specifies the type of an AWS source entity. &#39;kIAMUser&#39; indicates a unique user within an AWS account. &#39;kRegion&#39; indicates a geographical region in the global infrastructure. &#39;kAvailabilityZone&#39; indicates an availability zone within a region. &#39;kEC2Instance&#39; indicates a Virtual Machine running in AWS environment. &#39;kVPC&#39; indicates a virtual private cloud (VPC) network within AWS. &#39;kSubnet&#39; indicates a subnet inside the VPC. &#39;kNetworkSecurityGroup&#39; represents a network security group. &#39;kInstanceType&#39; represents various machine types. &#39;kKeyPair&#39; represents a pair of public and private key used to login into a Virtual Machine. &#39;kTag&#39; represents a tag attached to EC2 instance. &#39;kRDSOptionGroup&#39; represents a RDS option group for configuring database features. &#39;kRDSParameterGroup&#39; represents a RDS parameter group. &#39;kRDSInstance&#39; represents a RDS DB instance. &#39;kRDSSubnet&#39; represents a RDS subnet. &#39;kRDSTag&#39; represents a tag attached to RDS instance. | [optional] 
**ClusterNetworkInfo** | Pointer to [**FleetNetworkParams**](FleetNetworkParams.md) |  | [optional] 
**DbEngineId** | Pointer to **NullableString** | Specifies DB engine version info of the entity. This is populated only for RDSInstance, RDSOptionGroup and RDSParameterGroup entity types. | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the OS type of the Protection Source of type &#39;kVirtualMachine&#39; such as &#39;kWindows&#39; or &#39;kLinux&#39;. overrideDescription: true &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**IamRoleArn** | Pointer to **NullableString** | Specifies the IAM role which will be used to access the security credentials required for API calls. | [optional] 
**IpAddresses** | Pointer to **NullableString** | Specifies the IP address of the entity of type &#39;kVirtualMachine&#39;. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Object set by the Cloud Provider. If the provider did not set a name for the object, this field is not set. | [optional] 
**OwnerId** | Pointer to **NullableString** | Specifies the owner id of the resource in AWS environment. With type, name and ownerId gives a globally unique identity to the AWS entity. | [optional] 
**PhysicalSourceId** | Pointer to **NullableInt64** | Specifies the Protection Source id of the registered Physical Host. If the cloud entity is protected using a Physical Agent, it must be registered as a physical host. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region Id of the entity if the entity is an EC2 instance. | [optional] 
**ResourceId** | Pointer to **NullableString** | Specifies the unique Id of the resource given by the cloud provider. | [optional] 
**RestoreTaskId** | Pointer to **NullableInt64** | Specifies the id of the \&quot;convert and deploy\&quot; restore task that created the entity in the cloud.  It is required to support the DR-to-cloud usecase where we replicate an on-prem entity to a cluster running in cloud, bring it up using \&quot;convert and deploy\&quot; mechanism, protect it using a cloud job that uses physical adapter, and convert it back to the on-prem format before replication.  Before replicating, we need to update the backup task state of the backed up entity using the on-prem entity and on-prem entity&#39;s parent. The id is used to lookup the restore entity that contains details about the on-prem entity.  It is set at the time of refreshing the cloud entity hierarchy if all the following conditions are met: Name of the current entity matches with name of any cloud entity deployed using the \&quot;convert and deploy\&quot; restore task. Restore entity associated with the above matched cloud entity has &#39;failed_over&#39; flag set to true in its cloud extension. | [optional] 
**SecretAccessKey** | Pointer to **NullableString** | Specifies Secret Access key of the AWS account. | [optional] 
**SubscriptionType** | Pointer to **NullableString** | Specifies the subscription type of AWS such as &#39;kAWSCommercial&#39; or &#39;kAWSGovCloud&#39;. Specifies the subscription type of an AWS source entity. &#39;kAWSCommercial&#39; indicates a standard AWS subscription. &#39;kAWSGovCloud&#39; indicates a govt AWS subscription. | [optional] 
**TagAttributes** | Pointer to [**[]TagAttribute**](TagAttribute.md) | Specifies the list of AWS tag attributes. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of an AWS Protection Source Object such as &#39;kStorageContainer&#39;, &#39;kVirtualMachine&#39;, &#39;kVirtualNetwork&#39;, etc. Specifies the type of an AWS source entity. &#39;kIAMUser&#39; indicates a unique user within an AWS account. &#39;kRegion&#39; indicates a geographical region in the global infrastructure. &#39;kAvailabilityZone&#39; indicates an availability zone within a region. &#39;kEC2Instance&#39; indicates a Virtual Machine running in AWS environment. &#39;kVPC&#39; indicates a virtual private cloud (VPC) network within AWS. &#39;kSubnet&#39; indicates a subnet inside the VPC. &#39;kNetworkSecurityGroup&#39; represents a network security group. &#39;kInstanceType&#39; represents various machine types. &#39;kKeyPair&#39; represents a pair of public and private key used to login into a Virtual Machine. &#39;kTag&#39; represents a tag attached to EC2 instance. &#39;kRDSOptionGroup&#39; represents a RDS option group for configuring database features. &#39;kRDSParameterGroup&#39; represents a RDS parameter group. &#39;kRDSInstance&#39; represents a RDS DB instance. &#39;kRDSSubnet&#39; represents a RDS subnet. &#39;kRDSTag&#39; represents a tag attached to RDS instance. | [optional] 
**UserAccountId** | Pointer to **NullableString** | Specifies the account id derived from the ARN of the user. | [optional] 
**UserResourceName** | Pointer to **NullableString** | Specifies the Amazon Resource Name (ARN) of the user. | [optional] 

## Methods

### NewAwsProtectionSource

`func NewAwsProtectionSource() *AwsProtectionSource`

NewAwsProtectionSource instantiates a new AwsProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsProtectionSourceWithDefaults

`func NewAwsProtectionSourceWithDefaults() *AwsProtectionSource`

NewAwsProtectionSourceWithDefaults instantiates a new AwsProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AwsProtectionSource) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AwsProtectionSource) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AwsProtectionSource) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AwsProtectionSource) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### SetAccessKeyNil

`func (o *AwsProtectionSource) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *AwsProtectionSource) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetAmazonResourceName

`func (o *AwsProtectionSource) GetAmazonResourceName() string`

GetAmazonResourceName returns the AmazonResourceName field if non-nil, zero value otherwise.

### GetAmazonResourceNameOk

`func (o *AwsProtectionSource) GetAmazonResourceNameOk() (*string, bool)`

GetAmazonResourceNameOk returns a tuple with the AmazonResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonResourceName

`func (o *AwsProtectionSource) SetAmazonResourceName(v string)`

SetAmazonResourceName sets AmazonResourceName field to given value.

### HasAmazonResourceName

`func (o *AwsProtectionSource) HasAmazonResourceName() bool`

HasAmazonResourceName returns a boolean if a field has been set.

### SetAmazonResourceNameNil

`func (o *AwsProtectionSource) SetAmazonResourceNameNil(b bool)`

 SetAmazonResourceNameNil sets the value for AmazonResourceName to be an explicit nil

### UnsetAmazonResourceName
`func (o *AwsProtectionSource) UnsetAmazonResourceName()`

UnsetAmazonResourceName ensures that no value is present for AmazonResourceName, not even an explicit nil
### GetAuthMethod

`func (o *AwsProtectionSource) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AwsProtectionSource) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AwsProtectionSource) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *AwsProtectionSource) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### SetAuthMethodNil

`func (o *AwsProtectionSource) SetAuthMethodNil(b bool)`

 SetAuthMethodNil sets the value for AuthMethod to be an explicit nil

### UnsetAuthMethod
`func (o *AwsProtectionSource) UnsetAuthMethod()`

UnsetAuthMethod ensures that no value is present for AuthMethod, not even an explicit nil
### GetAwsFleetParams

`func (o *AwsProtectionSource) GetAwsFleetParams() AwsFleetParams`

GetAwsFleetParams returns the AwsFleetParams field if non-nil, zero value otherwise.

### GetAwsFleetParamsOk

`func (o *AwsProtectionSource) GetAwsFleetParamsOk() (*AwsFleetParams, bool)`

GetAwsFleetParamsOk returns a tuple with the AwsFleetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsFleetParams

`func (o *AwsProtectionSource) SetAwsFleetParams(v AwsFleetParams)`

SetAwsFleetParams sets AwsFleetParams field to given value.

### HasAwsFleetParams

`func (o *AwsProtectionSource) HasAwsFleetParams() bool`

HasAwsFleetParams returns a boolean if a field has been set.

### GetAwsType

`func (o *AwsProtectionSource) GetAwsType() string`

GetAwsType returns the AwsType field if non-nil, zero value otherwise.

### GetAwsTypeOk

`func (o *AwsProtectionSource) GetAwsTypeOk() (*string, bool)`

GetAwsTypeOk returns a tuple with the AwsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsType

`func (o *AwsProtectionSource) SetAwsType(v string)`

SetAwsType sets AwsType field to given value.

### HasAwsType

`func (o *AwsProtectionSource) HasAwsType() bool`

HasAwsType returns a boolean if a field has been set.

### SetAwsTypeNil

`func (o *AwsProtectionSource) SetAwsTypeNil(b bool)`

 SetAwsTypeNil sets the value for AwsType to be an explicit nil

### UnsetAwsType
`func (o *AwsProtectionSource) UnsetAwsType()`

UnsetAwsType ensures that no value is present for AwsType, not even an explicit nil
### GetClusterNetworkInfo

`func (o *AwsProtectionSource) GetClusterNetworkInfo() FleetNetworkParams`

GetClusterNetworkInfo returns the ClusterNetworkInfo field if non-nil, zero value otherwise.

### GetClusterNetworkInfoOk

`func (o *AwsProtectionSource) GetClusterNetworkInfoOk() (*FleetNetworkParams, bool)`

GetClusterNetworkInfoOk returns a tuple with the ClusterNetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNetworkInfo

`func (o *AwsProtectionSource) SetClusterNetworkInfo(v FleetNetworkParams)`

SetClusterNetworkInfo sets ClusterNetworkInfo field to given value.

### HasClusterNetworkInfo

`func (o *AwsProtectionSource) HasClusterNetworkInfo() bool`

HasClusterNetworkInfo returns a boolean if a field has been set.

### GetDbEngineId

`func (o *AwsProtectionSource) GetDbEngineId() string`

GetDbEngineId returns the DbEngineId field if non-nil, zero value otherwise.

### GetDbEngineIdOk

`func (o *AwsProtectionSource) GetDbEngineIdOk() (*string, bool)`

GetDbEngineIdOk returns a tuple with the DbEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbEngineId

`func (o *AwsProtectionSource) SetDbEngineId(v string)`

SetDbEngineId sets DbEngineId field to given value.

### HasDbEngineId

`func (o *AwsProtectionSource) HasDbEngineId() bool`

HasDbEngineId returns a boolean if a field has been set.

### SetDbEngineIdNil

`func (o *AwsProtectionSource) SetDbEngineIdNil(b bool)`

 SetDbEngineIdNil sets the value for DbEngineId to be an explicit nil

### UnsetDbEngineId
`func (o *AwsProtectionSource) UnsetDbEngineId()`

UnsetDbEngineId ensures that no value is present for DbEngineId, not even an explicit nil
### GetHostType

`func (o *AwsProtectionSource) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *AwsProtectionSource) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *AwsProtectionSource) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *AwsProtectionSource) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *AwsProtectionSource) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *AwsProtectionSource) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetIamRoleArn

`func (o *AwsProtectionSource) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *AwsProtectionSource) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *AwsProtectionSource) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.

### HasIamRoleArn

`func (o *AwsProtectionSource) HasIamRoleArn() bool`

HasIamRoleArn returns a boolean if a field has been set.

### SetIamRoleArnNil

`func (o *AwsProtectionSource) SetIamRoleArnNil(b bool)`

 SetIamRoleArnNil sets the value for IamRoleArn to be an explicit nil

### UnsetIamRoleArn
`func (o *AwsProtectionSource) UnsetIamRoleArn()`

UnsetIamRoleArn ensures that no value is present for IamRoleArn, not even an explicit nil
### GetIpAddresses

`func (o *AwsProtectionSource) GetIpAddresses() string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *AwsProtectionSource) GetIpAddressesOk() (*string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *AwsProtectionSource) SetIpAddresses(v string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *AwsProtectionSource) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *AwsProtectionSource) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *AwsProtectionSource) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetName

`func (o *AwsProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *AwsProtectionSource) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AwsProtectionSource) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AwsProtectionSource) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AwsProtectionSource) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *AwsProtectionSource) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *AwsProtectionSource) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetPhysicalSourceId

`func (o *AwsProtectionSource) GetPhysicalSourceId() int64`

GetPhysicalSourceId returns the PhysicalSourceId field if non-nil, zero value otherwise.

### GetPhysicalSourceIdOk

`func (o *AwsProtectionSource) GetPhysicalSourceIdOk() (*int64, bool)`

GetPhysicalSourceIdOk returns a tuple with the PhysicalSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalSourceId

`func (o *AwsProtectionSource) SetPhysicalSourceId(v int64)`

SetPhysicalSourceId sets PhysicalSourceId field to given value.

### HasPhysicalSourceId

`func (o *AwsProtectionSource) HasPhysicalSourceId() bool`

HasPhysicalSourceId returns a boolean if a field has been set.

### SetPhysicalSourceIdNil

`func (o *AwsProtectionSource) SetPhysicalSourceIdNil(b bool)`

 SetPhysicalSourceIdNil sets the value for PhysicalSourceId to be an explicit nil

### UnsetPhysicalSourceId
`func (o *AwsProtectionSource) UnsetPhysicalSourceId()`

UnsetPhysicalSourceId ensures that no value is present for PhysicalSourceId, not even an explicit nil
### GetRegionId

`func (o *AwsProtectionSource) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AwsProtectionSource) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AwsProtectionSource) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *AwsProtectionSource) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *AwsProtectionSource) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *AwsProtectionSource) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetResourceId

`func (o *AwsProtectionSource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AwsProtectionSource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AwsProtectionSource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AwsProtectionSource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AwsProtectionSource) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AwsProtectionSource) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetRestoreTaskId

`func (o *AwsProtectionSource) GetRestoreTaskId() int64`

GetRestoreTaskId returns the RestoreTaskId field if non-nil, zero value otherwise.

### GetRestoreTaskIdOk

`func (o *AwsProtectionSource) GetRestoreTaskIdOk() (*int64, bool)`

GetRestoreTaskIdOk returns a tuple with the RestoreTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTaskId

`func (o *AwsProtectionSource) SetRestoreTaskId(v int64)`

SetRestoreTaskId sets RestoreTaskId field to given value.

### HasRestoreTaskId

`func (o *AwsProtectionSource) HasRestoreTaskId() bool`

HasRestoreTaskId returns a boolean if a field has been set.

### SetRestoreTaskIdNil

`func (o *AwsProtectionSource) SetRestoreTaskIdNil(b bool)`

 SetRestoreTaskIdNil sets the value for RestoreTaskId to be an explicit nil

### UnsetRestoreTaskId
`func (o *AwsProtectionSource) UnsetRestoreTaskId()`

UnsetRestoreTaskId ensures that no value is present for RestoreTaskId, not even an explicit nil
### GetSecretAccessKey

`func (o *AwsProtectionSource) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AwsProtectionSource) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AwsProtectionSource) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AwsProtectionSource) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### SetSecretAccessKeyNil

`func (o *AwsProtectionSource) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *AwsProtectionSource) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetSubscriptionType

`func (o *AwsProtectionSource) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *AwsProtectionSource) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *AwsProtectionSource) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *AwsProtectionSource) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### SetSubscriptionTypeNil

`func (o *AwsProtectionSource) SetSubscriptionTypeNil(b bool)`

 SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil

### UnsetSubscriptionType
`func (o *AwsProtectionSource) UnsetSubscriptionType()`

UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
### GetTagAttributes

`func (o *AwsProtectionSource) GetTagAttributes() []TagAttribute`

GetTagAttributes returns the TagAttributes field if non-nil, zero value otherwise.

### GetTagAttributesOk

`func (o *AwsProtectionSource) GetTagAttributesOk() (*[]TagAttribute, bool)`

GetTagAttributesOk returns a tuple with the TagAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagAttributes

`func (o *AwsProtectionSource) SetTagAttributes(v []TagAttribute)`

SetTagAttributes sets TagAttributes field to given value.

### HasTagAttributes

`func (o *AwsProtectionSource) HasTagAttributes() bool`

HasTagAttributes returns a boolean if a field has been set.

### SetTagAttributesNil

`func (o *AwsProtectionSource) SetTagAttributesNil(b bool)`

 SetTagAttributesNil sets the value for TagAttributes to be an explicit nil

### UnsetTagAttributes
`func (o *AwsProtectionSource) UnsetTagAttributes()`

UnsetTagAttributes ensures that no value is present for TagAttributes, not even an explicit nil
### GetType

`func (o *AwsProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AwsProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AwsProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AwsProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUserAccountId

`func (o *AwsProtectionSource) GetUserAccountId() string`

GetUserAccountId returns the UserAccountId field if non-nil, zero value otherwise.

### GetUserAccountIdOk

`func (o *AwsProtectionSource) GetUserAccountIdOk() (*string, bool)`

GetUserAccountIdOk returns a tuple with the UserAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountId

`func (o *AwsProtectionSource) SetUserAccountId(v string)`

SetUserAccountId sets UserAccountId field to given value.

### HasUserAccountId

`func (o *AwsProtectionSource) HasUserAccountId() bool`

HasUserAccountId returns a boolean if a field has been set.

### SetUserAccountIdNil

`func (o *AwsProtectionSource) SetUserAccountIdNil(b bool)`

 SetUserAccountIdNil sets the value for UserAccountId to be an explicit nil

### UnsetUserAccountId
`func (o *AwsProtectionSource) UnsetUserAccountId()`

UnsetUserAccountId ensures that no value is present for UserAccountId, not even an explicit nil
### GetUserResourceName

`func (o *AwsProtectionSource) GetUserResourceName() string`

GetUserResourceName returns the UserResourceName field if non-nil, zero value otherwise.

### GetUserResourceNameOk

`func (o *AwsProtectionSource) GetUserResourceNameOk() (*string, bool)`

GetUserResourceNameOk returns a tuple with the UserResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResourceName

`func (o *AwsProtectionSource) SetUserResourceName(v string)`

SetUserResourceName sets UserResourceName field to given value.

### HasUserResourceName

`func (o *AwsProtectionSource) HasUserResourceName() bool`

HasUserResourceName returns a boolean if a field has been set.

### SetUserResourceNameNil

`func (o *AwsProtectionSource) SetUserResourceNameNil(b bool)`

 SetUserResourceNameNil sets the value for UserResourceName to be an explicit nil

### UnsetUserResourceName
`func (o *AwsProtectionSource) UnsetUserResourceName()`

UnsetUserResourceName ensures that no value is present for UserResourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


