// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Database An Oracle database on a bare metal or virtual machine DB system. For more information, see Bare Metal and Virtual Machine DB Systems (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type Database struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// A system-generated name for the database to ensure uniqueness within an Oracle Data Guard group (a primary database and its standby databases). The unique name cannot be changed.
	DbUniqueName *string `mandatory:"true" json:"dbUniqueName"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the database.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the database.
	LifecycleState DatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The character set for the database.
	CharacterSet *string `mandatory:"false" json:"characterSet"`

	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the database home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// The database workload type.
	DbWorkload *string `mandatory:"false" json:"dbWorkload"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"false" json:"ncharacterSet"`

	// The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	PdbName *string `mandatory:"false" json:"pdbName"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m Database) String() string {
	return common.PointerString(m)
}

// DatabaseLifecycleStateEnum Enum with underlying type: string
type DatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseLifecycleStateEnum
const (
	DatabaseLifecycleStateProvisioning     DatabaseLifecycleStateEnum = "PROVISIONING"
	DatabaseLifecycleStateAvailable        DatabaseLifecycleStateEnum = "AVAILABLE"
	DatabaseLifecycleStateUpdating         DatabaseLifecycleStateEnum = "UPDATING"
	DatabaseLifecycleStateBackupInProgress DatabaseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	DatabaseLifecycleStateTerminating      DatabaseLifecycleStateEnum = "TERMINATING"
	DatabaseLifecycleStateTerminated       DatabaseLifecycleStateEnum = "TERMINATED"
	DatabaseLifecycleStateRestoreFailed    DatabaseLifecycleStateEnum = "RESTORE_FAILED"
	DatabaseLifecycleStateFailed           DatabaseLifecycleStateEnum = "FAILED"
)

var mappingDatabaseLifecycleState = map[string]DatabaseLifecycleStateEnum{
	"PROVISIONING":       DatabaseLifecycleStateProvisioning,
	"AVAILABLE":          DatabaseLifecycleStateAvailable,
	"UPDATING":           DatabaseLifecycleStateUpdating,
	"BACKUP_IN_PROGRESS": DatabaseLifecycleStateBackupInProgress,
	"TERMINATING":        DatabaseLifecycleStateTerminating,
	"TERMINATED":         DatabaseLifecycleStateTerminated,
	"RESTORE_FAILED":     DatabaseLifecycleStateRestoreFailed,
	"FAILED":             DatabaseLifecycleStateFailed,
}

// GetDatabaseLifecycleStateEnumValues Enumerates the set of values for DatabaseLifecycleStateEnum
func GetDatabaseLifecycleStateEnumValues() []DatabaseLifecycleStateEnum {
	values := make([]DatabaseLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseLifecycleState {
		values = append(values, v)
	}
	return values
}
