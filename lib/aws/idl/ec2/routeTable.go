// Copyright 2017 Pulumi, Inc. All rights reserved.

package ec2

import (
	"github.com/pulumi/coconut/pkg/resource/idl"
)

// RouteTable is a route table within your VPC.  After creating a route table, you can add routes and associate the
// table with a subnet.  For more information, see
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route-table.html.
type RouteTable struct {
	idl.NamedResource
	// The VPC where the route table will be created.
	VPC *VPC `coco:"vpc,replaces"`
}