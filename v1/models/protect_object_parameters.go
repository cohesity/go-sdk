// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProtectObjectParameters Protect Object Parameters
//
// Specifies the parameters to protect an object.
//
// swagger:model ProtectObjectParameters
type ProtectObjectParameters struct {

	// Specifies the environment type of the Protection Source object.
	// Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
	// NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter.
	// 'kVMware' indicates the VMware Protection Source environment.
	// 'kHyperV' indicates the HyperV Protection Source environment.
	// 'kSQL' indicates the SQL Protection Source environment.
	// 'kView' indicates the View Protection Source environment.
	// 'kPuppeteer' indicates the Cohesity's Remote Adapter.
	// 'kPhysical' indicates the physical Protection Source environment.
	// 'kPure' indicates the Pure Storage Protection Source environment.
	// 'kNimble' indicates the Nimble Storage Protection Source environment.
	// 'kAzure' indicates the Microsoft's Azure Protection Source environment.
	// 'kNetapp' indicates the Netapp Protection Source environment.
	// 'kAgent' indicates the Agent Protection Source environment.
	// 'kGenericNas' indicates the Generic Network Attached Storage Protection
	// Source environment.
	// 'kAcropolis' indicates the Acropolis Protection Source environment.
	// 'kPhysicalFiles' indicates the Physical Files Protection Source environment.
	// 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment.
	// 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment.
	// 'kGPFS' indicates IBM's GPFS Protection Source environment.
	// 'kKVM' indicates the KVM Protection Source environment.
	// 'kAWS' indicates the AWS Protection Source environment.
	// 'kExchange' indicates the Exchange Protection Source environment.
	// 'kHyperVVSS' indicates the HyperV VSS Protection Source
	// environment.
	// 'kOracle' indicates the Oracle Protection Source environment.
	// 'kGCP' indicates the Google Cloud Platform Protection Source environment.
	// 'kFlashBlade' indicates the Flash Blade Protection Source environment.
	// 'kAWSNative' indicates the AWS Native Protection Source environment.
	// 'kO365' indicates the Office 365 Protection Source environment.
	// 'kO365Outlook' indicates Office 365 outlook Protection Source environment.
	// 'kHyperFlex' indicates the Hyper Flex Protection Source environment.
	// 'kGCPNative' indicates the GCP Native Protection Source environment.
	// 'kAzureNative' indicates the Azure Native Protection Source environment.
	// 'kKubernetes' indicates a Kubernetes Protection Source environment.
	// 'kElastifile' indicates Elastifile Protection Source environment.
	// 'kAD' indicates Active Directory Protection Source environment.
	// 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment.
	// 'kCassandra' indicates Cassandra Protection Source environment.
	// 'kMongoDB' indicates MongoDB Protection Source environment.
	// 'kCouchbase' indicates Couchbase Protection Source environment.
	// 'kHdfs' indicates Hdfs Protection Source environment.
	// 'kHive' indicates Hive Protection Source environment.
	// 'kHBase' indicates HBase Protection Source environment.
	// 'kUDA' indicates Universal Data Adapter Protection Source environment.
	// 'kO365Teams' indicates the Office365 Teams Protection Source environment.
	// 'kO365Group' indicates the Office365 Groups Protection Source environment.
	// 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment.
	// 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment.
	// 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment.
	// 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment.
	// Enum: ["kVMware","kHyperV","kSQL","kView","kPuppeteer","kPhysical","kPure","kNimble","kIbmFlashSystem","kAzure","kNetapp","kAgent","kGenericNas","kAcropolis","kPhysicalFiles","kIsilon","kGPFS","kKVM","kAWS","kExchange","kHyperVVSS","kOracle","kGCP","kFlashBlade","kAWSNative","kO365","kO365Outlook","kHyperFlex","kGCPNative","kAzureNative","kKubernetes","kElastifile","kAD","kRDSSnapshotManager","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kUDA","kO365Teams","kO365Group","kO365Exchange","kO365OneDrive","kO365Sharepoint","kO365PublicFolders"]
	ProtectionSourceEnvironment *string `json:"protectionSourceEnvironment,omitempty"`

	// Specifies the ids of the Protection Sources to protect.
	// Required: true
	ProtectionSourceIds []int64 `json:"protectionSourceIds"`

	// Specifies the Rpo policy id.
	// Required: true
	RpoPolicyID *string `json:"rpoPolicyId"`
}

// Validate validates this protect object parameters
func (m *ProtectObjectParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtectionSourceEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionSourceIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRpoPolicyID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var protectObjectParametersTypeProtectionSourceEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVMware","kHyperV","kSQL","kView","kPuppeteer","kPhysical","kPure","kNimble","kIbmFlashSystem","kAzure","kNetapp","kAgent","kGenericNas","kAcropolis","kPhysicalFiles","kIsilon","kGPFS","kKVM","kAWS","kExchange","kHyperVVSS","kOracle","kGCP","kFlashBlade","kAWSNative","kO365","kO365Outlook","kHyperFlex","kGCPNative","kAzureNative","kKubernetes","kElastifile","kAD","kRDSSnapshotManager","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kUDA","kO365Teams","kO365Group","kO365Exchange","kO365OneDrive","kO365Sharepoint","kO365PublicFolders"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectObjectParametersTypeProtectionSourceEnvironmentPropEnum = append(protectObjectParametersTypeProtectionSourceEnvironmentPropEnum, v)
	}
}

const (

	// ProtectObjectParametersProtectionSourceEnvironmentKVMware captures enum value "kVMware"
	ProtectObjectParametersProtectionSourceEnvironmentKVMware string = "kVMware"

	// ProtectObjectParametersProtectionSourceEnvironmentKHyperV captures enum value "kHyperV"
	ProtectObjectParametersProtectionSourceEnvironmentKHyperV string = "kHyperV"

	// ProtectObjectParametersProtectionSourceEnvironmentKSQL captures enum value "kSQL"
	ProtectObjectParametersProtectionSourceEnvironmentKSQL string = "kSQL"

	// ProtectObjectParametersProtectionSourceEnvironmentKView captures enum value "kView"
	ProtectObjectParametersProtectionSourceEnvironmentKView string = "kView"

	// ProtectObjectParametersProtectionSourceEnvironmentKPuppeteer captures enum value "kPuppeteer"
	ProtectObjectParametersProtectionSourceEnvironmentKPuppeteer string = "kPuppeteer"

	// ProtectObjectParametersProtectionSourceEnvironmentKPhysical captures enum value "kPhysical"
	ProtectObjectParametersProtectionSourceEnvironmentKPhysical string = "kPhysical"

	// ProtectObjectParametersProtectionSourceEnvironmentKPure captures enum value "kPure"
	ProtectObjectParametersProtectionSourceEnvironmentKPure string = "kPure"

	// ProtectObjectParametersProtectionSourceEnvironmentKNimble captures enum value "kNimble"
	ProtectObjectParametersProtectionSourceEnvironmentKNimble string = "kNimble"

	// ProtectObjectParametersProtectionSourceEnvironmentKIbmFlashSystem captures enum value "kIbmFlashSystem"
	ProtectObjectParametersProtectionSourceEnvironmentKIbmFlashSystem string = "kIbmFlashSystem"

	// ProtectObjectParametersProtectionSourceEnvironmentKAzure captures enum value "kAzure"
	ProtectObjectParametersProtectionSourceEnvironmentKAzure string = "kAzure"

	// ProtectObjectParametersProtectionSourceEnvironmentKNetapp captures enum value "kNetapp"
	ProtectObjectParametersProtectionSourceEnvironmentKNetapp string = "kNetapp"

	// ProtectObjectParametersProtectionSourceEnvironmentKAgent captures enum value "kAgent"
	ProtectObjectParametersProtectionSourceEnvironmentKAgent string = "kAgent"

	// ProtectObjectParametersProtectionSourceEnvironmentKGenericNas captures enum value "kGenericNas"
	ProtectObjectParametersProtectionSourceEnvironmentKGenericNas string = "kGenericNas"

	// ProtectObjectParametersProtectionSourceEnvironmentKAcropolis captures enum value "kAcropolis"
	ProtectObjectParametersProtectionSourceEnvironmentKAcropolis string = "kAcropolis"

	// ProtectObjectParametersProtectionSourceEnvironmentKPhysicalFiles captures enum value "kPhysicalFiles"
	ProtectObjectParametersProtectionSourceEnvironmentKPhysicalFiles string = "kPhysicalFiles"

	// ProtectObjectParametersProtectionSourceEnvironmentKIsilon captures enum value "kIsilon"
	ProtectObjectParametersProtectionSourceEnvironmentKIsilon string = "kIsilon"

	// ProtectObjectParametersProtectionSourceEnvironmentKGPFS captures enum value "kGPFS"
	ProtectObjectParametersProtectionSourceEnvironmentKGPFS string = "kGPFS"

	// ProtectObjectParametersProtectionSourceEnvironmentKKVM captures enum value "kKVM"
	ProtectObjectParametersProtectionSourceEnvironmentKKVM string = "kKVM"

	// ProtectObjectParametersProtectionSourceEnvironmentKAWS captures enum value "kAWS"
	ProtectObjectParametersProtectionSourceEnvironmentKAWS string = "kAWS"

	// ProtectObjectParametersProtectionSourceEnvironmentKExchange captures enum value "kExchange"
	ProtectObjectParametersProtectionSourceEnvironmentKExchange string = "kExchange"

	// ProtectObjectParametersProtectionSourceEnvironmentKHyperVVSS captures enum value "kHyperVVSS"
	ProtectObjectParametersProtectionSourceEnvironmentKHyperVVSS string = "kHyperVVSS"

	// ProtectObjectParametersProtectionSourceEnvironmentKOracle captures enum value "kOracle"
	ProtectObjectParametersProtectionSourceEnvironmentKOracle string = "kOracle"

	// ProtectObjectParametersProtectionSourceEnvironmentKGCP captures enum value "kGCP"
	ProtectObjectParametersProtectionSourceEnvironmentKGCP string = "kGCP"

	// ProtectObjectParametersProtectionSourceEnvironmentKFlashBlade captures enum value "kFlashBlade"
	ProtectObjectParametersProtectionSourceEnvironmentKFlashBlade string = "kFlashBlade"

	// ProtectObjectParametersProtectionSourceEnvironmentKAWSNative captures enum value "kAWSNative"
	ProtectObjectParametersProtectionSourceEnvironmentKAWSNative string = "kAWSNative"

	// ProtectObjectParametersProtectionSourceEnvironmentKO365 captures enum value "kO365"
	ProtectObjectParametersProtectionSourceEnvironmentKO365 string = "kO365"

	// ProtectObjectParametersProtectionSourceEnvironmentKO365Outlook captures enum value "kO365Outlook"
	ProtectObjectParametersProtectionSourceEnvironmentKO365Outlook string = "kO365Outlook"

	// ProtectObjectParametersProtectionSourceEnvironmentKHyperFlex captures enum value "kHyperFlex"
	ProtectObjectParametersProtectionSourceEnvironmentKHyperFlex string = "kHyperFlex"

	// ProtectObjectParametersProtectionSourceEnvironmentKGCPNative captures enum value "kGCPNative"
	ProtectObjectParametersProtectionSourceEnvironmentKGCPNative string = "kGCPNative"

	// ProtectObjectParametersProtectionSourceEnvironmentKAzureNative captures enum value "kAzureNative"
	ProtectObjectParametersProtectionSourceEnvironmentKAzureNative string = "kAzureNative"

	// ProtectObjectParametersProtectionSourceEnvironmentKKubernetes captures enum value "kKubernetes"
	ProtectObjectParametersProtectionSourceEnvironmentKKubernetes string = "kKubernetes"

	// ProtectObjectParametersProtectionSourceEnvironmentKElastifile captures enum value "kElastifile"
	ProtectObjectParametersProtectionSourceEnvironmentKElastifile string = "kElastifile"

	// ProtectObjectParametersProtectionSourceEnvironmentKAD captures enum value "kAD"
	ProtectObjectParametersProtectionSourceEnvironmentKAD string = "kAD"

	// ProtectObjectParametersProtectionSourceEnvironmentKRDSSnapshotManager captures enum value "kRDSSnapshotManager"
	ProtectObjectParametersProtectionSourceEnvironmentKRDSSnapshotManager string = "kRDSSnapshotManager"

	// ProtectObjectParametersProtectionSourceEnvironmentKCassandra captures enum value "kCassandra"
	ProtectObjectParametersProtectionSourceEnvironmentKCassandra string = "kCassandra"

	// ProtectObjectParametersProtectionSourceEnvironmentKMongoDB captures enum value "kMongoDB"
	ProtectObjectParametersProtectionSourceEnvironmentKMongoDB string = "kMongoDB"

	// ProtectObjectParametersProtectionSourceEnvironmentKCouchbase captures enum value "kCouchbase"
	ProtectObjectParametersProtectionSourceEnvironmentKCouchbase string = "kCouchbase"

	// ProtectObjectParametersProtectionSourceEnvironmentKHdfs captures enum value "kHdfs"
	ProtectObjectParametersProtectionSourceEnvironmentKHdfs string = "kHdfs"

	// ProtectObjectParametersProtectionSourceEnvironmentKHive captures enum value "kHive"
	ProtectObjectParametersProtectionSourceEnvironmentKHive string = "kHive"

	// ProtectObjectParametersProtectionSourceEnvironmentKHBase captures enum value "kHBase"
	ProtectObjectParametersProtectionSourceEnvironmentKHBase string = "kHBase"

	// ProtectObjectParametersProtectionSourceEnvironmentKUDA captures enum value "kUDA"
	ProtectObjectParametersProtectionSourceEnvironmentKUDA string = "kUDA"

	// ProtectObjectParametersProtectionSourceEnvironmentKO365Teams captures enum value "kO365Teams"
	ProtectObjectParametersProtectionSourceEnvironmentKO365Teams string = "kO365Teams"

	// ProtectObjectParametersProtectionSourceEnvironmentKO365Group captures enum value "kO365Group"
	ProtectObjectParametersProtectionSourceEnvironmentKO365Group string = "kO365Group"

	// ProtectObjectParametersProtectionSourceEnvironmentKO365Exchange captures enum value "kO365Exchange"
	ProtectObjectParametersProtectionSourceEnvironmentKO365Exchange string = "kO365Exchange"

	// ProtectObjectParametersProtectionSourceEnvironmentKO365OneDrive captures enum value "kO365OneDrive"
	ProtectObjectParametersProtectionSourceEnvironmentKO365OneDrive string = "kO365OneDrive"

	// ProtectObjectParametersProtectionSourceEnvironmentKO365Sharepoint captures enum value "kO365Sharepoint"
	ProtectObjectParametersProtectionSourceEnvironmentKO365Sharepoint string = "kO365Sharepoint"

	// ProtectObjectParametersProtectionSourceEnvironmentKO365PublicFolders captures enum value "kO365PublicFolders"
	ProtectObjectParametersProtectionSourceEnvironmentKO365PublicFolders string = "kO365PublicFolders"
)

// prop value enum
func (m *ProtectObjectParameters) validateProtectionSourceEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectObjectParametersTypeProtectionSourceEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectObjectParameters) validateProtectionSourceEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionSourceEnvironment) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtectionSourceEnvironmentEnum("protectionSourceEnvironment", "body", *m.ProtectionSourceEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *ProtectObjectParameters) validateProtectionSourceIds(formats strfmt.Registry) error {

	if err := validate.Required("protectionSourceIds", "body", m.ProtectionSourceIds); err != nil {
		return err
	}

	return nil
}

func (m *ProtectObjectParameters) validateRpoPolicyID(formats strfmt.Registry) error {

	if err := validate.Required("rpoPolicyId", "body", m.RpoPolicyID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this protect object parameters based on context it is used
func (m *ProtectObjectParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectObjectParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectObjectParameters) UnmarshalBinary(b []byte) error {
	var res ProtectObjectParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
