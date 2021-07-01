# CloudDeployTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsParams** | Pointer to [**AwsParams**](AwsParams.md) |  | [optional] 
**AzureParams** | Pointer to [**AzureParams**](AzureParams.md) |  | [optional] 
**GcpParams** | Pointer to [**GcpParams**](GcpParams.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Entity corresponding to the cloud deploy target.  Specifies the id field inside the EntityProto. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the inner object&#39;s name or a human-readable string made off the salient attributes. This is only plumbed when Entity objects are exposed to Iris BE or to Yoda. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the CloudDeploy target. &#39;kAzure&#39; indicates that Azure as a cloud deploy target type. &#39;kAWS&#39; indicates that AWS as a cloud deploy target type. &#39;kGCP&#39; indicates that GCP as a cloud deploy target type. | [optional] 

## Methods

### NewCloudDeployTargetDetails

`func NewCloudDeployTargetDetails() *CloudDeployTargetDetails`

NewCloudDeployTargetDetails instantiates a new CloudDeployTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDeployTargetDetailsWithDefaults

`func NewCloudDeployTargetDetailsWithDefaults() *CloudDeployTargetDetails`

NewCloudDeployTargetDetailsWithDefaults instantiates a new CloudDeployTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsParams

`func (o *CloudDeployTargetDetails) GetAwsParams() AwsParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *CloudDeployTargetDetails) GetAwsParamsOk() (*AwsParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *CloudDeployTargetDetails) SetAwsParams(v AwsParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *CloudDeployTargetDetails) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *CloudDeployTargetDetails) GetAzureParams() AzureParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *CloudDeployTargetDetails) GetAzureParamsOk() (*AzureParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *CloudDeployTargetDetails) SetAzureParams(v AzureParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *CloudDeployTargetDetails) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *CloudDeployTargetDetails) GetGcpParams() GcpParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *CloudDeployTargetDetails) GetGcpParamsOk() (*GcpParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *CloudDeployTargetDetails) SetGcpParams(v GcpParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *CloudDeployTargetDetails) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetId

`func (o *CloudDeployTargetDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudDeployTargetDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudDeployTargetDetails) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CloudDeployTargetDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CloudDeployTargetDetails) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CloudDeployTargetDetails) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CloudDeployTargetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudDeployTargetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudDeployTargetDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudDeployTargetDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CloudDeployTargetDetails) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CloudDeployTargetDetails) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *CloudDeployTargetDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudDeployTargetDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudDeployTargetDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudDeployTargetDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CloudDeployTargetDetails) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CloudDeployTargetDetails) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


