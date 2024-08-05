# McmRigelClaimRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimToken** | **string** | Claim token used for authentication. | 
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster incarnation id. | [optional] 
**RigelGuid** | **int64** | Unique id for rigel instance. | 
**RigelIp** | Pointer to **NullableString** | Specifies the Rigel IP. | [optional] 
**RigelName** | Pointer to **NullableString** | Specifies the Rigel name. | [optional] 
**RigelType** | Pointer to **NullableString** | Specifies the Rigel type that is being claimed. | [optional] 
**SoftwareVersion** | Pointer to **NullableString** | Specifies the Rigel Software version. | [optional] 

## Methods

### NewMcmRigelClaimRequestParams

`func NewMcmRigelClaimRequestParams(claimToken string, rigelGuid int64, ) *McmRigelClaimRequestParams`

NewMcmRigelClaimRequestParams instantiates a new McmRigelClaimRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmRigelClaimRequestParamsWithDefaults

`func NewMcmRigelClaimRequestParamsWithDefaults() *McmRigelClaimRequestParams`

NewMcmRigelClaimRequestParamsWithDefaults instantiates a new McmRigelClaimRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimToken

`func (o *McmRigelClaimRequestParams) GetClaimToken() string`

GetClaimToken returns the ClaimToken field if non-nil, zero value otherwise.

### GetClaimTokenOk

`func (o *McmRigelClaimRequestParams) GetClaimTokenOk() (*string, bool)`

GetClaimTokenOk returns a tuple with the ClaimToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimToken

`func (o *McmRigelClaimRequestParams) SetClaimToken(v string)`

SetClaimToken sets ClaimToken field to given value.


### GetClusterId

`func (o *McmRigelClaimRequestParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmRigelClaimRequestParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmRigelClaimRequestParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmRigelClaimRequestParams) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmRigelClaimRequestParams) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmRigelClaimRequestParams) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *McmRigelClaimRequestParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmRigelClaimRequestParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmRigelClaimRequestParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *McmRigelClaimRequestParams) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *McmRigelClaimRequestParams) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *McmRigelClaimRequestParams) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetRigelGuid

`func (o *McmRigelClaimRequestParams) GetRigelGuid() int64`

GetRigelGuid returns the RigelGuid field if non-nil, zero value otherwise.

### GetRigelGuidOk

`func (o *McmRigelClaimRequestParams) GetRigelGuidOk() (*int64, bool)`

GetRigelGuidOk returns a tuple with the RigelGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelGuid

`func (o *McmRigelClaimRequestParams) SetRigelGuid(v int64)`

SetRigelGuid sets RigelGuid field to given value.


### GetRigelIp

`func (o *McmRigelClaimRequestParams) GetRigelIp() string`

GetRigelIp returns the RigelIp field if non-nil, zero value otherwise.

### GetRigelIpOk

`func (o *McmRigelClaimRequestParams) GetRigelIpOk() (*string, bool)`

GetRigelIpOk returns a tuple with the RigelIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelIp

`func (o *McmRigelClaimRequestParams) SetRigelIp(v string)`

SetRigelIp sets RigelIp field to given value.

### HasRigelIp

`func (o *McmRigelClaimRequestParams) HasRigelIp() bool`

HasRigelIp returns a boolean if a field has been set.

### SetRigelIpNil

`func (o *McmRigelClaimRequestParams) SetRigelIpNil(b bool)`

 SetRigelIpNil sets the value for RigelIp to be an explicit nil

### UnsetRigelIp
`func (o *McmRigelClaimRequestParams) UnsetRigelIp()`

UnsetRigelIp ensures that no value is present for RigelIp, not even an explicit nil
### GetRigelName

`func (o *McmRigelClaimRequestParams) GetRigelName() string`

GetRigelName returns the RigelName field if non-nil, zero value otherwise.

### GetRigelNameOk

`func (o *McmRigelClaimRequestParams) GetRigelNameOk() (*string, bool)`

GetRigelNameOk returns a tuple with the RigelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelName

`func (o *McmRigelClaimRequestParams) SetRigelName(v string)`

SetRigelName sets RigelName field to given value.

### HasRigelName

`func (o *McmRigelClaimRequestParams) HasRigelName() bool`

HasRigelName returns a boolean if a field has been set.

### SetRigelNameNil

`func (o *McmRigelClaimRequestParams) SetRigelNameNil(b bool)`

 SetRigelNameNil sets the value for RigelName to be an explicit nil

### UnsetRigelName
`func (o *McmRigelClaimRequestParams) UnsetRigelName()`

UnsetRigelName ensures that no value is present for RigelName, not even an explicit nil
### GetRigelType

`func (o *McmRigelClaimRequestParams) GetRigelType() string`

GetRigelType returns the RigelType field if non-nil, zero value otherwise.

### GetRigelTypeOk

`func (o *McmRigelClaimRequestParams) GetRigelTypeOk() (*string, bool)`

GetRigelTypeOk returns a tuple with the RigelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelType

`func (o *McmRigelClaimRequestParams) SetRigelType(v string)`

SetRigelType sets RigelType field to given value.

### HasRigelType

`func (o *McmRigelClaimRequestParams) HasRigelType() bool`

HasRigelType returns a boolean if a field has been set.

### SetRigelTypeNil

`func (o *McmRigelClaimRequestParams) SetRigelTypeNil(b bool)`

 SetRigelTypeNil sets the value for RigelType to be an explicit nil

### UnsetRigelType
`func (o *McmRigelClaimRequestParams) UnsetRigelType()`

UnsetRigelType ensures that no value is present for RigelType, not even an explicit nil
### GetSoftwareVersion

`func (o *McmRigelClaimRequestParams) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *McmRigelClaimRequestParams) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *McmRigelClaimRequestParams) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *McmRigelClaimRequestParams) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *McmRigelClaimRequestParams) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *McmRigelClaimRequestParams) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


