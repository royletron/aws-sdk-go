// Package ec2 provides a client for Amazon Elastic Compute Cloud.
package ec2

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// EC2 is a client for Amazon Elastic Compute Cloud.
type EC2 struct {
	client *aws.EC2Client
}

// New returns a new EC2 client.
func New(creds aws.Credentials, region string, client *http.Client) *EC2 {
	if client == nil {
		client = http.DefaultClient
	}

	service := "ec2"
	endpoint, service, region := endpoints.Lookup("ec2", region)

	return &EC2{
		client: &aws.EC2Client{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2014-10-01",
		},
	}
}

// AcceptVpcPeeringConnection accept a VPC peering connection request. To
// accept a request, the VPC peering connection must be in the
// pending-acceptance state, and you must be the owner of the peer Use the
// DescribeVpcPeeringConnections request to view your outstanding VPC
// peering connection requests.
func (c *EC2) AcceptVpcPeeringConnection(req AcceptVpcPeeringConnectionRequest) (resp *AcceptVpcPeeringConnectionResult, err error) {
	resp = &AcceptVpcPeeringConnectionResult{}
	err = c.client.Do("AcceptVpcPeeringConnection", "POST", "/", req, resp)
	return
}

// AllocateAddress acquires an Elastic IP address. An Elastic IP address is
// for use either in the EC2-Classic platform or in a For more information,
// see Elastic IP Addresses in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) AllocateAddress(req AllocateAddressRequest) (resp *AllocateAddressResult, err error) {
	resp = &AllocateAddressResult{}
	err = c.client.Do("AllocateAddress", "POST", "/", req, resp)
	return
}

// AssignPrivateIPAddresses assigns one or more secondary private IP
// addresses to the specified network interface. You can specify one or
// more specific secondary IP addresses, or you can specify the number of
// secondary IP addresses to be automatically assigned within the subnet's
// block range. The number of secondary IP addresses that you can assign to
// an instance varies by instance type. For information about instance
// types, see Instance Types in the Amazon Elastic Compute Cloud User Guide
// . For more information about Elastic IP addresses, see Elastic IP
// Addresses in the Amazon Elastic Compute Cloud User Guide
// AssignPrivateIpAddresses is available only in EC2-VPC.
func (c *EC2) AssignPrivateIPAddresses(req AssignPrivateIPAddressesRequest) (err error) {
	// NRE
	err = c.client.Do("AssignPrivateIpAddresses", "POST", "/", req, nil)
	return
}

// AssociateAddress associates an Elastic IP address with an instance or a
// network interface. An Elastic IP address is for use in either the
// EC2-Classic platform or in a For more information, see Elastic IP
// Addresses in the Amazon Elastic Compute Cloud User Guide [EC2-Classic,
// VPC in an EC2-VPC-only account] If the Elastic IP address is already
// associated with a different instance, it is disassociated from that
// instance and associated with the specified instance. in an EC2-Classic
// account] If you don't specify a private IP address, the Elastic IP
// address is associated with the primary IP address. If the Elastic IP
// address is already associated with a different instance or a network
// interface, you get an error unless you allow reassociation. This is an
// idempotent operation. If you perform the operation more than once,
// Amazon EC2 doesn't return an error.
func (c *EC2) AssociateAddress(req AssociateAddressRequest) (resp *AssociateAddressResult, err error) {
	resp = &AssociateAddressResult{}
	err = c.client.Do("AssociateAddress", "POST", "/", req, resp)
	return
}

// AssociateDhcpOptions associates a set of options (that you've previously
// created) with the specified or associates no options with the After you
// associate the options with the any existing instances and all new
// instances that you launch in that VPC use the options. You don't need to
// restart or relaunch the instances. They automatically pick up the
// changes within a few hours, depending on how frequently the instance
// renews its lease. You can explicitly renew the lease using the operating
// system on the instance. For more information, see Options Sets in the
// Amazon Virtual Private Cloud User Guide
func (c *EC2) AssociateDhcpOptions(req AssociateDhcpOptionsRequest) (err error) {
	// NRE
	err = c.client.Do("AssociateDhcpOptions", "POST", "/", req, nil)
	return
}

// AssociateRouteTable associates a subnet with a route table. The subnet
// and route table must be in the same This association causes traffic
// originating from the subnet to be routed according to the routes in the
// route table. The action returns an association ID, which you need in
// order to disassociate the route table from the subnet later. A route
// table can be associated with multiple subnets. For more information
// about route tables, see Route Tables in the Amazon Virtual Private Cloud
// User Guide
func (c *EC2) AssociateRouteTable(req AssociateRouteTableRequest) (resp *AssociateRouteTableResult, err error) {
	resp = &AssociateRouteTableResult{}
	err = c.client.Do("AssociateRouteTable", "POST", "/", req, resp)
	return
}

// AttachInternetGateway attaches an Internet gateway to a enabling
// connectivity between the Internet and the For more information about
// your VPC and Internet gateway, see the Amazon Virtual Private Cloud User
// Guide
func (c *EC2) AttachInternetGateway(req AttachInternetGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("AttachInternetGateway", "POST", "/", req, nil)
	return
}

// AttachNetworkInterface is undocumented.
func (c *EC2) AttachNetworkInterface(req AttachNetworkInterfaceRequest) (resp *AttachNetworkInterfaceResult, err error) {
	resp = &AttachNetworkInterfaceResult{}
	err = c.client.Do("AttachNetworkInterface", "POST", "/", req, resp)
	return
}

// AttachVolume attaches an Amazon EBS volume to a running or stopped
// instance and exposes it to the instance with the specified device name.
// Encrypted Amazon EBS volumes may only be attached to instances that
// support Amazon EBS encryption. For more information, see Amazon EBS
// Encryption in the Amazon Elastic Compute Cloud User Guide For a list of
// supported device names, see Attaching an Amazon EBS Volume to an
// Instance . Any device names that aren't reserved for instance store
// volumes can be used for Amazon EBS volumes. For more information, see
// Amazon EC2 Instance Store in the Amazon Elastic Compute Cloud User Guide
// If a volume has an AWS Marketplace product code: The volume can only be
// attached as the root device of a stopped instance. You must be
// subscribed to the AWS Marketplace code that is on the volume. The
// configuration (instance type, operating system) of the instance must
// support that specific AWS Marketplace code. For example, you cannot take
// a volume from a Windows instance and attach it to a Linux instance. AWS
// Marketplace product codes are copied from the volume to the instance.
// For an overview of the AWS Marketplace, see
// https://aws.amazon.com/marketplace/help/200900000 . For more information
// about how to use the AWS Marketplace, see AWS Marketplace For more
// information about Amazon EBS volumes, see Attaching Amazon EBS Volumes
// in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) AttachVolume(req AttachVolumeRequest) (resp *VolumeAttachment, err error) {
	resp = &VolumeAttachment{}
	err = c.client.Do("AttachVolume", "POST", "/", req, resp)
	return
}

// AttachVpnGateway attaches a virtual private gateway to a For more
// information, see Adding a Hardware Virtual Private Gateway to Your in
// the Amazon Virtual Private Cloud User Guide
func (c *EC2) AttachVpnGateway(req AttachVpnGatewayRequest) (resp *AttachVpnGatewayResult, err error) {
	resp = &AttachVpnGatewayResult{}
	err = c.client.Do("AttachVpnGateway", "POST", "/", req, resp)
	return
}

// AuthorizeSecurityGroupEgress adds one or more egress rules to a security
// group for use with a Specifically, this action permits instances to send
// traffic to one or more destination IP address ranges, or to one or more
// destination security groups for the same You can have up to 50 rules per
// security group (covering both ingress and egress rules). A security
// group is for use with instances either in the EC2-Classic platform or in
// a specific This action doesn't apply to security groups for use in
// EC2-Classic. For more information, see Security Groups for Your in the
// Amazon Virtual Private Cloud User Guide Each rule consists of the
// protocol (for example, plus either a range or a source group. For the
// TCP and UDP protocols, you must also specify the destination port or
// port range. For the protocol, you must also specify the type and code.
// You can use -1 for the type or code to mean all types or all codes. Rule
// changes are propagated to affected instances as quickly as possible.
// However, a small delay might occur.
func (c *EC2) AuthorizeSecurityGroupEgress(req AuthorizeSecurityGroupEgressRequest) (err error) {
	// NRE
	err = c.client.Do("AuthorizeSecurityGroupEgress", "POST", "/", req, nil)
	return
}

// AuthorizeSecurityGroupIngress adds one or more ingress rules to a
// security group. EC2-Classic: You can have up to 100 rules per group.
// EC2-VPC: You can have up to 50 rules per group (covering both ingress
// and egress rules). Rule changes are propagated to instances within the
// security group as quickly as possible. However, a small delay might
// occur. [EC2-Classic] This action gives one or more IP address ranges
// permission to access a security group in your account, or gives one or
// more security groups (called the source groups ) permission to access a
// security group for your account. A source group can be for your own AWS
// account, or another. [EC2-VPC] This action gives one or more IP address
// ranges permission to access a security group in your or gives one or
// more other security groups (called the source groups ) permission to
// access a security group for your The security groups must all be for the
// same
func (c *EC2) AuthorizeSecurityGroupIngress(req AuthorizeSecurityGroupIngressRequest) (err error) {
	// NRE
	err = c.client.Do("AuthorizeSecurityGroupIngress", "POST", "/", req, nil)
	return
}

// BundleInstance bundles an Amazon instance store-backed Windows instance.
// During bundling, only the root device volume is bundled. Data on other
// instance store volumes is not preserved. This procedure is not
// applicable for Linux/Unix instances or Windows instances that are backed
// by Amazon For more information, see Creating an Instance Store-Backed
// Windows
func (c *EC2) BundleInstance(req BundleInstanceRequest) (resp *BundleInstanceResult, err error) {
	resp = &BundleInstanceResult{}
	err = c.client.Do("BundleInstance", "POST", "/", req, resp)
	return
}

// CancelBundleTask cancels a bundling operation for an instance
// store-backed Windows instance.
func (c *EC2) CancelBundleTask(req CancelBundleTaskRequest) (resp *CancelBundleTaskResult, err error) {
	resp = &CancelBundleTaskResult{}
	err = c.client.Do("CancelBundleTask", "POST", "/", req, resp)
	return
}

// CancelConversionTask cancels an active conversion task. The task can be
// the import of an instance or volume. The action removes all artifacts of
// the conversion, including a partially uploaded volume or instance. If
// the conversion is complete or is in the process of transferring the
// final disk image, the command fails and returns an exception. For more
// information, see Using the Command Line Tools to Import Your Virtual
// Machine to Amazon EC2 in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CancelConversionTask(req CancelConversionRequest) (err error) {
	// NRE
	err = c.client.Do("CancelConversionTask", "POST", "/", req, nil)
	return
}

// CancelExportTask cancels an active export task. The request removes all
// artifacts of the export, including any partially-created Amazon S3
// objects. If the export task is complete or is in the process of
// transferring the final disk image, the command fails and returns an
// error.
func (c *EC2) CancelExportTask(req CancelExportTaskRequest) (err error) {
	// NRE
	err = c.client.Do("CancelExportTask", "POST", "/", req, nil)
	return
}

// CancelReservedInstancesListing cancels the specified Reserved Instance
// listing in the Reserved Instance Marketplace. For more information, see
// Reserved Instance Marketplace in the Amazon Elastic Compute Cloud User
// Guide
func (c *EC2) CancelReservedInstancesListing(req CancelReservedInstancesListingRequest) (resp *CancelReservedInstancesListingResult, err error) {
	resp = &CancelReservedInstancesListingResult{}
	err = c.client.Do("CancelReservedInstancesListing", "POST", "/", req, resp)
	return
}

// CancelSpotInstanceRequests cancels one or more Spot Instance requests.
// Spot Instances are instances that Amazon EC2 starts on your behalf when
// the maximum price that you specify exceeds the current Spot Price.
// Amazon EC2 periodically sets the Spot Price based on available Spot
// Instance capacity and current Spot Instance requests. For more
// information about Spot Instances, see Spot Instances in the Amazon
// Elastic Compute Cloud User Guide Canceling a Spot Instance request does
// not terminate running Spot Instances associated with the request.
func (c *EC2) CancelSpotInstanceRequests(req CancelSpotInstanceRequestsRequest) (resp *CancelSpotInstanceRequestsResult, err error) {
	resp = &CancelSpotInstanceRequestsResult{}
	err = c.client.Do("CancelSpotInstanceRequests", "POST", "/", req, resp)
	return
}

// ConfirmProductInstance determines whether a product code is associated
// with an instance. This action can only be used by the owner of the
// product code. It is useful when a product code owner needs to verify
// whether another user's instance is eligible for support.
func (c *EC2) ConfirmProductInstance(req ConfirmProductInstanceRequest) (resp *ConfirmProductInstanceResult, err error) {
	resp = &ConfirmProductInstanceResult{}
	err = c.client.Do("ConfirmProductInstance", "POST", "/", req, resp)
	return
}

// CopyImage initiates the copy of an AMI from the specified source region
// to the region in which the request was made. You specify the destination
// region by using its endpoint when making the request. AMIs that use
// encrypted Amazon EBS snapshots cannot be copied with this method. For
// more information, see Copying AMIs in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) CopyImage(req CopyImageRequest) (resp *CopyImageResult, err error) {
	resp = &CopyImageResult{}
	err = c.client.Do("CopyImage", "POST", "/", req, resp)
	return
}

// CopySnapshot copies a point-in-time snapshot of an Amazon EBS volume and
// stores it in Amazon S3. You can copy the snapshot within the same region
// or from one region to another. You can use the snapshot to create Amazon
// EBS volumes or Amazon Machine Images (AMIs). The snapshot is copied to
// the regional endpoint that you send the request to. Copies of encrypted
// Amazon EBS snapshots remain encrypted. Copies of unencrypted snapshots
// remain unencrypted. For more information, see Copying an Amazon EBS
// Snapshot in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CopySnapshot(req CopySnapshotRequest) (resp *CopySnapshotResult, err error) {
	resp = &CopySnapshotResult{}
	err = c.client.Do("CopySnapshot", "POST", "/", req, resp)
	return
}

// CreateCustomerGateway provides information to AWS about your VPN
// customer gateway device. The customer gateway is the appliance at your
// end of the VPN connection. (The device on the AWS side of the VPN
// connection is the virtual private gateway.) You must provide the
// Internet-routable IP address of the customer gateway's external
// interface. The IP address must be static and can't be behind a device
// performing network address translation For devices that use Border
// Gateway Protocol you can also provide the device's BGP Autonomous System
// Number You can use an existing ASN assigned to your network. If you
// don't have an ASN already, you can use a private ASN (in the 64512 -
// 65534 range). Amazon EC2 supports all 2-byte ASN numbers in the range of
// 1 - 65534, with the exception of 7224, which is reserved in the
// us-east-1 region, and 9059, which is reserved in the eu-west-1 region.
// For more information about VPN customer gateways, see Adding a Hardware
// Virtual Private Gateway to Your in the Amazon Virtual Private Cloud User
// Guide
func (c *EC2) CreateCustomerGateway(req CreateCustomerGatewayRequest) (resp *CreateCustomerGatewayResult, err error) {
	resp = &CreateCustomerGatewayResult{}
	err = c.client.Do("CreateCustomerGateway", "POST", "/", req, resp)
	return
}

// CreateDhcpOptions creates a set of options for your After creating the
// set, you must associate it with the causing all existing and new
// instances that you launch in the VPC to use this set of options. The
// following are the individual options you can specify. For more
// information about the options, see RFC 2132 domain-name-servers - The IP
// addresses of up to four domain name servers, or AmazonProvidedDNS . The
// default option set specifies AmazonProvidedDNS . If specifying more than
// one domain name server, specify the IP addresses in a single parameter,
// separated by commas. domain-name - If you're using AmazonProvidedDNS in
// us-east-1 , specify ec2.internal . If you're using AmazonProvidedDNS in
// another region, specify region.compute.internal (for example,
// ap-northeast-1.compute.internal ). Otherwise, specify a domain name (for
// example, MyCompany.com ). If specifying more than one domain name,
// separate them with spaces. ntp-servers - The IP addresses of up to four
// Network Time Protocol servers. netbios-name-servers - The IP addresses
// of up to four NetBIOS name servers. netbios-node-type - The NetBIOS node
// type (1, 2, 4, or 8). We recommend that you specify 2 (broadcast and
// multicast are not currently supported). For more information about these
// node types, see RFC 2132 . Your VPC automatically starts out with a set
// of options that includes only a DNS server that we provide
// (AmazonProvidedDNS). If you create a set of options, and if your VPC has
// an Internet gateway, make sure to set the domain-name-servers option
// either to AmazonProvidedDNS or to a domain name server of your choice.
// For more information about options, see Options Sets in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) CreateDhcpOptions(req CreateDhcpOptionsRequest) (resp *CreateDhcpOptionsResult, err error) {
	resp = &CreateDhcpOptionsResult{}
	err = c.client.Do("CreateDhcpOptions", "POST", "/", req, resp)
	return
}

// CreateImage creates an Amazon EBS-backed AMI from an Amazon EBS-backed
// instance that is either running or stopped. If you customized your
// instance with instance store volumes or EBS volumes in addition to the
// root device volume, the new AMI contains block device mapping
// information for those volumes. When you launch an instance from this new
// the instance automatically launches with those additional volumes. For
// more information, see Creating Amazon EBS-Backed Linux AMIs in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateImage(req CreateImageRequest) (resp *CreateImageResult, err error) {
	resp = &CreateImageResult{}
	err = c.client.Do("CreateImage", "POST", "/", req, resp)
	return
}

// CreateInstanceExportTask exports a running or stopped instance to an
// Amazon S3 bucket. For information about the supported operating systems,
// image formats, and known limitations for the types of instances you can
// export, see Exporting EC2 Instances in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) CreateInstanceExportTask(req CreateInstanceExportTaskRequest) (resp *CreateInstanceExportTaskResult, err error) {
	resp = &CreateInstanceExportTaskResult{}
	err = c.client.Do("CreateInstanceExportTask", "POST", "/", req, resp)
	return
}

// CreateInternetGateway creates an Internet gateway for use with a After
// creating the Internet gateway, you attach it to a VPC using
// AttachInternetGateway For more information about your VPC and Internet
// gateway, see the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateInternetGateway(req CreateInternetGatewayRequest) (resp *CreateInternetGatewayResult, err error) {
	resp = &CreateInternetGatewayResult{}
	err = c.client.Do("CreateInternetGateway", "POST", "/", req, resp)
	return
}

// CreateKeyPair creates a 2048-bit RSA key pair with the specified name.
// Amazon EC2 stores the public key and displays the private key for you to
// save to a file. The private key is returned as an unencrypted PEM
// encoded PKCS#8 private key. If a key with the specified name already
// exists, Amazon EC2 returns an error. You can have up to five thousand
// key pairs per region. The key pair returned to you is available only in
// the region in which you create it. To create a key pair that is
// available in all regions, use ImportKeyPair For more information about
// key pairs, see Key Pairs in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateKeyPair(req CreateKeyPairRequest) (resp *KeyPair, err error) {
	resp = &KeyPair{}
	err = c.client.Do("CreateKeyPair", "POST", "/", req, resp)
	return
}

// CreateNetworkAcl creates a network ACL in a Network ACLs provide an
// optional layer of security (in addition to security groups) for the
// instances in your For more information about network ACLs, see Network
// ACLs in the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateNetworkAcl(req CreateNetworkAclRequest) (resp *CreateNetworkAclResult, err error) {
	resp = &CreateNetworkAclResult{}
	err = c.client.Do("CreateNetworkAcl", "POST", "/", req, resp)
	return
}

// CreateNetworkAclEntry creates an entry (a rule) in a network ACL with
// the specified rule number. Each network ACL has a set of numbered
// ingress rules and a separate set of numbered egress rules. When
// determining whether a packet should be allowed in or out of a subnet
// associated with the we process the entries in the ACL according to the
// rule numbers, in ascending order. Each network ACL has a set of ingress
// rules and a separate set of egress rules. We recommend that you leave
// room between the rule numbers (for example, 100, 110, 120, and not
// number them one right after the other (for example, 101, 102, 103, This
// makes it easier to add a rule between existing ones without having to
// renumber the rules. After you add an entry, you can't modify it; you
// must either replace it, or create an entry and delete the old one. For
// more information about network ACLs, see Network ACLs in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) CreateNetworkAclEntry(req CreateNetworkAclEntryRequest) (err error) {
	// NRE
	err = c.client.Do("CreateNetworkAclEntry", "POST", "/", req, nil)
	return
}

// CreateNetworkInterface creates a network interface in the specified
// subnet. For more information about network interfaces, see Elastic
// Network Interfaces in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateNetworkInterface(req CreateNetworkInterfaceRequest) (resp *CreateNetworkInterfaceResult, err error) {
	resp = &CreateNetworkInterfaceResult{}
	err = c.client.Do("CreateNetworkInterface", "POST", "/", req, resp)
	return
}

// CreatePlacementGroup creates a placement group that you launch cluster
// instances into. You must give the group a name that's unique within the
// scope of your account. For more information about placement groups and
// cluster instances, see Cluster Instances in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) CreatePlacementGroup(req CreatePlacementGroupRequest) (err error) {
	// NRE
	err = c.client.Do("CreatePlacementGroup", "POST", "/", req, nil)
	return
}

// CreateReservedInstancesListing creates a listing for Amazon EC2 Reserved
// Instances to be sold in the Reserved Instance Marketplace. You can
// submit one Reserved Instance listing at a time. To get a list of your
// Reserved Instances, you can use the DescribeReservedInstances operation.
// The Reserved Instance Marketplace matches sellers who want to resell
// Reserved Instance capacity that they no longer need with buyers who want
// to purchase additional capacity. Reserved Instances bought and sold
// through the Reserved Instance Marketplace work like any other Reserved
// Instances. To sell your Reserved Instances, you must first register as a
// Seller in the Reserved Instance Marketplace. After completing the
// registration process, you can create a Reserved Instance Marketplace
// listing of some or all of your Reserved Instances, and specify the
// upfront price to receive for them. Your Reserved Instance listings then
// become available for purchase. To view the details of your Reserved
// Instance listing, you can use the DescribeReservedInstancesListings
// operation. For more information, see Reserved Instance Marketplace in
// the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateReservedInstancesListing(req CreateReservedInstancesListingRequest) (resp *CreateReservedInstancesListingResult, err error) {
	resp = &CreateReservedInstancesListingResult{}
	err = c.client.Do("CreateReservedInstancesListing", "POST", "/", req, resp)
	return
}

// CreateRoute creates a route in a route table within a You must specify
// one of the following targets: Internet gateway or virtual private
// gateway, NAT instance, VPC peering connection, or network interface.
// When determining how to route traffic, we use the route with the most
// specific match. For example, let's say the traffic is destined for
// 192.0.2.3 , and the route table includes the following two routes:
// 192.0.2.0/24 (goes to some target 192.0.2.0/28 (goes to some target Both
// routes apply to the traffic destined for 192.0.2.3 . However, the second
// route in the list covers a smaller number of IP addresses and is
// therefore more specific, so we use that route to determine where to
// target the traffic. For more information about route tables, see Route
// Tables in the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateRoute(req CreateRouteRequest) (err error) {
	// NRE
	err = c.client.Do("CreateRoute", "POST", "/", req, nil)
	return
}

// CreateRouteTable creates a route table for the specified After you
// create a route table, you can add routes and associate the table with a
// subnet. For more information about route tables, see Route Tables in the
// Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateRouteTable(req CreateRouteTableRequest) (resp *CreateRouteTableResult, err error) {
	resp = &CreateRouteTableResult{}
	err = c.client.Do("CreateRouteTable", "POST", "/", req, resp)
	return
}

// CreateSecurityGroup creates a security group. A security group is for
// use with instances either in the EC2-Classic platform or in a specific
// For more information, see Amazon EC2 Security Groups in the Amazon
// Elastic Compute Cloud User Guide and Security Groups for Your in the
// Amazon Virtual Private Cloud User Guide EC2-Classic: You can have up to
// 500 security groups. EC2-VPC: You can create up to 100 security groups
// per When you create a security group, you specify a friendly name of
// your choice. You can have a security group for use in EC2-Classic with
// the same name as a security group for use in a However, you can't have
// two security groups for use in EC2-Classic with the same name or two
// security groups for use in a VPC with the same name. You have a default
// security group for use in EC2-Classic and a default security group for
// use in your If you don't specify a security group when you launch an
// instance, the instance is launched into the appropriate default security
// group. A default security group includes a default rule that grants
// instances unrestricted network access to each other. You can add or
// remove rules from your security groups using
// AuthorizeSecurityGroupIngress , AuthorizeSecurityGroupEgress ,
// RevokeSecurityGroupIngress , and RevokeSecurityGroupEgress
func (c *EC2) CreateSecurityGroup(req CreateSecurityGroupRequest) (resp *CreateSecurityGroupResult, err error) {
	resp = &CreateSecurityGroupResult{}
	err = c.client.Do("CreateSecurityGroup", "POST", "/", req, resp)
	return
}

// CreateSnapshot creates a snapshot of an Amazon EBS volume and stores it
// in Amazon S3. You can use snapshots for backups, to make copies of
// Amazon EBS volumes, and to save data before shutting down an instance.
// When a snapshot is created, any AWS Marketplace product codes that are
// associated with the source volume are propagated to the snapshot. You
// can take a snapshot of an attached volume that is in use. However,
// snapshots only capture data that has been written to your Amazon EBS
// volume at the time the snapshot command is issued; this may exclude any
// data that has been cached by any applications or the operating system.
// If you can pause any file systems on the volume long enough to take a
// snapshot, your snapshot should be complete. However, if you cannot pause
// all file writes to the volume, you should unmount the volume from within
// the instance, issue the snapshot command, and then remount the volume to
// ensure a consistent and complete snapshot. You may remount and use your
// volume while the snapshot status is pending To create a snapshot for
// Amazon EBS volumes that serve as root devices, you should stop the
// instance before taking the snapshot. Snapshots that are taken from
// encrypted volumes are automatically encrypted. Volumes that are created
// from encrypted snapshots are also automatically encrypted. Your
// encrypted volumes and any associated snapshots always remain protected.
// For more information, see Amazon Elastic Block Store and Amazon EBS
// Encryption in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateSnapshot(req CreateSnapshotRequest) (resp *Snapshot, err error) {
	resp = &Snapshot{}
	err = c.client.Do("CreateSnapshot", "POST", "/", req, resp)
	return
}

// CreateSpotDatafeedSubscription creates a datafeed for Spot Instances,
// enabling you to view Spot Instance usage logs. You can create one data
// feed per AWS account. For more information, see Spot Instances in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateSpotDatafeedSubscription(req CreateSpotDatafeedSubscriptionRequest) (resp *CreateSpotDatafeedSubscriptionResult, err error) {
	resp = &CreateSpotDatafeedSubscriptionResult{}
	err = c.client.Do("CreateSpotDatafeedSubscription", "POST", "/", req, resp)
	return
}

// CreateSubnet creates a subnet in an existing When you create each
// subnet, you provide the VPC ID and the block you want for the subnet.
// After you create a subnet, you can't change its block. The subnet's
// block can be the same as the VPC's block (assuming you want only a
// single subnet in the or a subset of the VPC's block. If you create more
// than one subnet in a the subnets' blocks must not overlap. The smallest
// subnet (and you can create uses a /28 netmask (16 IP addresses), and the
// largest uses a /16 netmask (65,536 IP addresses). AWS reserves both the
// first four and the last IP address in each subnet's block. They're not
// available for use. If you add more than one subnet to a they're set up
// in a star topology with a logical router in the middle. If you launch an
// instance in a VPC using an Amazon EBS-backed the IP address doesn't
// change if you stop and restart the instance (unlike a similar instance
// launched outside a which gets a new IP address when restarted). It's
// therefore possible to have a subnet with no running instances (they're
// all stopped), but no remaining IP addresses available. For more
// information about subnets, see Your VPC and Subnets in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) CreateSubnet(req CreateSubnetRequest) (resp *CreateSubnetResult, err error) {
	resp = &CreateSubnetResult{}
	err = c.client.Do("CreateSubnet", "POST", "/", req, resp)
	return
}

// CreateTags adds or overwrites one or more tags for the specified EC2
// resource or resources. Each resource can have a maximum of 10 tags. Each
// tag consists of a key and optional value. Tag keys must be unique per
// resource. For more information about tags, see Tagging Your Resources in
// the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateTags(req CreateTagsRequest) (err error) {
	// NRE
	err = c.client.Do("CreateTags", "POST", "/", req, nil)
	return
}

// CreateVolume creates an Amazon EBS volume that can be attached to an
// instance in the same Availability Zone. The volume is created in the
// specified region. You can create a new empty volume or restore a volume
// from an Amazon EBS snapshot. Any AWS Marketplace product codes from the
// snapshot are propagated to the volume. You can create encrypted volumes
// with the Encrypted parameter. Encrypted volumes may only be attached to
// instances that support Amazon EBS encryption. Volumes that are created
// from encrypted snapshots are also automatically encrypted. For more
// information, see Amazon EBS Encryption in the Amazon Elastic Compute
// Cloud User Guide For more information, see Creating or Restoring an
// Amazon EBS Volume in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) CreateVolume(req CreateVolumeRequest) (resp *Volume, err error) {
	resp = &Volume{}
	err = c.client.Do("CreateVolume", "POST", "/", req, resp)
	return
}

// CreateVpc creates a VPC with the specified block. The smallest VPC you
// can create uses a /28 netmask (16 IP addresses), and the largest uses a
// /16 netmask (65,536 IP addresses). To help you decide how big to make
// your see Your VPC and Subnets in the Amazon Virtual Private Cloud User
// Guide By default, each instance you launch in the VPC has the default
// options, which includes only a default DNS server that we provide
// (AmazonProvidedDNS). For more information about options, see Options
// Sets in the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateVpc(req CreateVpcRequest) (resp *CreateVpcResult, err error) {
	resp = &CreateVpcResult{}
	err = c.client.Do("CreateVpc", "POST", "/", req, resp)
	return
}

// CreateVpcPeeringConnection requests a VPC peering connection between two
// VPCs: a requester VPC that you own and a peer VPC with which to create
// the connection. The peer VPC can belong to another AWS account. The
// requester VPC and peer VPC cannot have overlapping blocks. The owner of
// the peer VPC must accept the peering request to activate the peering
// connection. The VPC peering connection request expires after 7 days,
// after which it cannot be accepted or rejected. A
// CreateVpcPeeringConnection request between VPCs with overlapping blocks
// results in the VPC peering connection having a status of failed
func (c *EC2) CreateVpcPeeringConnection(req CreateVpcPeeringConnectionRequest) (resp *CreateVpcPeeringConnectionResult, err error) {
	resp = &CreateVpcPeeringConnectionResult{}
	err = c.client.Do("CreateVpcPeeringConnection", "POST", "/", req, resp)
	return
}

// CreateVpnConnection creates a VPN connection between an existing virtual
// private gateway and a VPN customer gateway. The only supported
// connection type is ipsec.1 The response includes information that you
// need to give to your network administrator to configure your customer
// gateway. We strongly recommend that you use when calling this operation
// because the response contains sensitive cryptographic information for
// configuring your customer gateway. If you decide to shut down your VPN
// connection for any reason and later create a new VPN connection, you
// must reconfigure your customer gateway with the new information returned
// from this call. For more information about VPN connections, see Adding a
// Hardware Virtual Private Gateway to Your in the Amazon Virtual Private
// Cloud User Guide
func (c *EC2) CreateVpnConnection(req CreateVpnConnectionRequest) (resp *CreateVpnConnectionResult, err error) {
	resp = &CreateVpnConnectionResult{}
	err = c.client.Do("CreateVpnConnection", "POST", "/", req, resp)
	return
}

// CreateVpnConnectionRoute creates a static route associated with a VPN
// connection between an existing virtual private gateway and a VPN
// customer gateway. The static route allows traffic to be routed from the
// virtual private gateway to the VPN customer gateway. For more
// information about VPN connections, see Adding a Hardware Virtual Private
// Gateway to Your in the Amazon Virtual Private Cloud User Guide
func (c *EC2) CreateVpnConnectionRoute(req CreateVpnConnectionRouteRequest) (err error) {
	// NRE
	err = c.client.Do("CreateVpnConnectionRoute", "POST", "/", req, nil)
	return
}

// CreateVpnGateway creates a virtual private gateway. A virtual private
// gateway is the endpoint on the VPC side of your VPN connection. You can
// create a virtual private gateway before creating the VPC itself. For
// more information about virtual private gateways, see Adding a Hardware
// Virtual Private Gateway to Your in the Amazon Virtual Private Cloud User
// Guide
func (c *EC2) CreateVpnGateway(req CreateVpnGatewayRequest) (resp *CreateVpnGatewayResult, err error) {
	resp = &CreateVpnGatewayResult{}
	err = c.client.Do("CreateVpnGateway", "POST", "/", req, resp)
	return
}

// DeleteCustomerGateway deletes the specified customer gateway. You must
// delete the VPN connection before you can delete the customer gateway.
func (c *EC2) DeleteCustomerGateway(req DeleteCustomerGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteCustomerGateway", "POST", "/", req, nil)
	return
}

// DeleteDhcpOptions deletes the specified set of options. You must
// disassociate the set of options before you can delete it. You can
// disassociate the set of options by associating either a new set of
// options or the default set of options with the
func (c *EC2) DeleteDhcpOptions(req DeleteDhcpOptionsRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteDhcpOptions", "POST", "/", req, nil)
	return
}

// DeleteInternetGateway deletes the specified Internet gateway. You must
// detach the Internet gateway from the VPC before you can delete it.
func (c *EC2) DeleteInternetGateway(req DeleteInternetGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteInternetGateway", "POST", "/", req, nil)
	return
}

// DeleteKeyPair deletes the specified key pair, by removing the public key
// from Amazon EC2.
func (c *EC2) DeleteKeyPair(req DeleteKeyPairRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteKeyPair", "POST", "/", req, nil)
	return
}

// DeleteNetworkAcl deletes the specified network You can't delete the ACL
// if it's associated with any subnets. You can't delete the default
// network
func (c *EC2) DeleteNetworkAcl(req DeleteNetworkAclRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteNetworkAcl", "POST", "/", req, nil)
	return
}

// DeleteNetworkAclEntry deletes the specified ingress or egress entry
// (rule) from the specified network
func (c *EC2) DeleteNetworkAclEntry(req DeleteNetworkAclEntryRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteNetworkAclEntry", "POST", "/", req, nil)
	return
}

// DeleteNetworkInterface deletes the specified network interface. You must
// detach the network interface before you can delete it.
func (c *EC2) DeleteNetworkInterface(req DeleteNetworkInterfaceRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteNetworkInterface", "POST", "/", req, nil)
	return
}

// DeletePlacementGroup deletes the specified placement group. You must
// terminate all instances in the placement group before you can delete the
// placement group. For more information about placement groups and cluster
// instances, see Cluster Instances in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) DeletePlacementGroup(req DeletePlacementGroupRequest) (err error) {
	// NRE
	err = c.client.Do("DeletePlacementGroup", "POST", "/", req, nil)
	return
}

// DeleteRoute deletes the specified route from the specified route table.
func (c *EC2) DeleteRoute(req DeleteRouteRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteRoute", "POST", "/", req, nil)
	return
}

// DeleteRouteTable deletes the specified route table. You must
// disassociate the route table from any subnets before you can delete it.
// You can't delete the main route table.
func (c *EC2) DeleteRouteTable(req DeleteRouteTableRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteRouteTable", "POST", "/", req, nil)
	return
}

// DeleteSecurityGroup deletes a security group. If you attempt to delete a
// security group that is associated with an instance, or is referenced by
// another security group, the operation fails with InvalidGroup.InUse in
// EC2-Classic or DependencyViolation in EC2-VPC.
func (c *EC2) DeleteSecurityGroup(req DeleteSecurityGroupRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSecurityGroup", "POST", "/", req, nil)
	return
}

// DeleteSnapshot deletes the specified snapshot. When you make periodic
// snapshots of a volume, the snapshots are incremental, and only the
// blocks on the device that have changed since your last snapshot are
// saved in the new snapshot. When you delete a snapshot, only the data not
// needed for any other snapshot is removed. So regardless of which prior
// snapshots have been deleted, all active snapshots will have access to
// all the information needed to restore the volume. You cannot delete a
// snapshot of the root device of an Amazon EBS volume used by a registered
// You must first de-register the AMI before you can delete the snapshot.
// For more information, see Deleting an Amazon EBS Snapshot in the Amazon
// Elastic Compute Cloud User Guide
func (c *EC2) DeleteSnapshot(req DeleteSnapshotRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSnapshot", "POST", "/", req, nil)
	return
}

// DeleteSpotDatafeedSubscription deletes the datafeed for Spot Instances.
// For more information, see Spot Instances in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) DeleteSpotDatafeedSubscription(req DeleteSpotDatafeedSubscriptionRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSpotDatafeedSubscription", "POST", "/", req, nil)
	return
}

// DeleteSubnet deletes the specified subnet. You must terminate all
// running instances in the subnet before you can delete the subnet.
func (c *EC2) DeleteSubnet(req DeleteSubnetRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteSubnet", "POST", "/", req, nil)
	return
}

// DeleteTags deletes the specified set of tags from the specified set of
// resources. This call is designed to follow a DescribeTags request. For
// more information about tags, see Tagging Your Resources in the Amazon
// Elastic Compute Cloud User Guide
func (c *EC2) DeleteTags(req DeleteTagsRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteTags", "POST", "/", req, nil)
	return
}

// DeleteVolume deletes the specified Amazon EBS volume. The volume must be
// in the available state (not attached to an instance). The volume may
// remain in the deleting state for several minutes. For more information,
// see Deleting an Amazon EBS Volume in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) DeleteVolume(req DeleteVolumeRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVolume", "POST", "/", req, nil)
	return
}

// DeleteVpc deletes the specified You must detach or delete all gateways
// and resources that are associated with the VPC before you can delete it.
// For example, you must terminate all instances running in the delete all
// security groups associated with the VPC (except the default one), delete
// all route tables associated with the VPC (except the default one), and
// so on.
func (c *EC2) DeleteVpc(req DeleteVpcRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVpc", "POST", "/", req, nil)
	return
}

// DeleteVpcPeeringConnection deletes a VPC peering connection. Either the
// owner of the requester VPC or the owner of the peer VPC can delete the
// VPC peering connection if it's in the active state. The owner of the
// requester VPC can delete a VPC peering connection in the
// pending-acceptance state.
func (c *EC2) DeleteVpcPeeringConnection(req DeleteVpcPeeringConnectionRequest) (resp *DeleteVpcPeeringConnectionResult, err error) {
	resp = &DeleteVpcPeeringConnectionResult{}
	err = c.client.Do("DeleteVpcPeeringConnection", "POST", "/", req, resp)
	return
}

// DeleteVpnConnection deletes the specified VPN connection. If you're
// deleting the VPC and its associated components, we recommend that you
// detach the virtual private gateway from the VPC and delete the VPC
// before deleting the VPN connection. If you believe that the tunnel
// credentials for your VPN connection have been compromised, you can
// delete the VPN connection and create a new one that has new keys,
// without needing to delete the VPC or virtual private gateway. If you
// create a new VPN connection, you must reconfigure the customer gateway
// using the new configuration information returned with the new VPN
// connection
func (c *EC2) DeleteVpnConnection(req DeleteVpnConnectionRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVpnConnection", "POST", "/", req, nil)
	return
}

// DeleteVpnConnectionRoute deletes the specified static route associated
// with a VPN connection between an existing virtual private gateway and a
// VPN customer gateway. The static route allows traffic to be routed from
// the virtual private gateway to the VPN customer gateway.
func (c *EC2) DeleteVpnConnectionRoute(req DeleteVpnConnectionRouteRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVpnConnectionRoute", "POST", "/", req, nil)
	return
}

// DeleteVpnGateway deletes the specified virtual private gateway. We
// recommend that before you delete a virtual private gateway, you detach
// it from the VPC and delete the VPN connection. Note that you don't need
// to delete the virtual private gateway if you plan to delete and recreate
// the VPN connection between your VPC and your network.
func (c *EC2) DeleteVpnGateway(req DeleteVpnGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteVpnGateway", "POST", "/", req, nil)
	return
}

// DeregisterImage deregisters the specified After you deregister an it
// can't be used to launch new instances. This command does not delete the
func (c *EC2) DeregisterImage(req DeregisterImageRequest) (err error) {
	// NRE
	err = c.client.Do("DeregisterImage", "POST", "/", req, nil)
	return
}

// DescribeAccountAttributes describes the specified attribute of your AWS
// account.
func (c *EC2) DescribeAccountAttributes(req DescribeAccountAttributesRequest) (resp *DescribeAccountAttributesResult, err error) {
	resp = &DescribeAccountAttributesResult{}
	err = c.client.Do("DescribeAccountAttributes", "POST", "/", req, resp)
	return
}

// DescribeAddresses describes one or more of your Elastic IP addresses. An
// Elastic IP address is for use in either the EC2-Classic platform or in a
// For more information, see Elastic IP Addresses in the Amazon Elastic
// Compute Cloud User Guide
func (c *EC2) DescribeAddresses(req DescribeAddressesRequest) (resp *DescribeAddressesResult, err error) {
	resp = &DescribeAddressesResult{}
	err = c.client.Do("DescribeAddresses", "POST", "/", req, resp)
	return
}

// DescribeAvailabilityZones describes one or more of the Availability
// Zones that are available to you. The results include zones only for the
// region you're currently using. If there is an event impacting an
// Availability Zone, you can use this request to view the state and any
// provided message for that Availability Zone. For more information, see
// Regions and Availability Zones in the Amazon Elastic Compute Cloud User
// Guide
func (c *EC2) DescribeAvailabilityZones(req DescribeAvailabilityZonesRequest) (resp *DescribeAvailabilityZonesResult, err error) {
	resp = &DescribeAvailabilityZonesResult{}
	err = c.client.Do("DescribeAvailabilityZones", "POST", "/", req, resp)
	return
}

// DescribeBundleTasks describes one or more of your bundling tasks.
// Completed bundle tasks are listed for only a limited time. If your
// bundle task is no longer in the list, you can still register an AMI from
// it. Just use RegisterImage with the Amazon S3 bucket name and image
// manifest name you provided to the bundle task.
func (c *EC2) DescribeBundleTasks(req DescribeBundleTasksRequest) (resp *DescribeBundleTasksResult, err error) {
	resp = &DescribeBundleTasksResult{}
	err = c.client.Do("DescribeBundleTasks", "POST", "/", req, resp)
	return
}

// DescribeConversionTasks describes one or more of your conversion tasks.
// For more information, see Using the Command Line Tools to Import Your
// Virtual Machine to Amazon EC2 in the Amazon Elastic Compute Cloud User
// Guide
func (c *EC2) DescribeConversionTasks(req DescribeConversionTasksRequest) (resp *DescribeConversionTasksResult, err error) {
	resp = &DescribeConversionTasksResult{}
	err = c.client.Do("DescribeConversionTasks", "POST", "/", req, resp)
	return
}

// DescribeCustomerGateways describes one or more of your VPN customer
// gateways. For more information about VPN customer gateways, see Adding a
// Hardware Virtual Private Gateway to Your in the Amazon Virtual Private
// Cloud User Guide
func (c *EC2) DescribeCustomerGateways(req DescribeCustomerGatewaysRequest) (resp *DescribeCustomerGatewaysResult, err error) {
	resp = &DescribeCustomerGatewaysResult{}
	err = c.client.Do("DescribeCustomerGateways", "POST", "/", req, resp)
	return
}

// DescribeDhcpOptions describes one or more of your options sets. For more
// information about options sets, see Options Sets in the Amazon Virtual
// Private Cloud User Guide
func (c *EC2) DescribeDhcpOptions(req DescribeDhcpOptionsRequest) (resp *DescribeDhcpOptionsResult, err error) {
	resp = &DescribeDhcpOptionsResult{}
	err = c.client.Do("DescribeDhcpOptions", "POST", "/", req, resp)
	return
}

// DescribeExportTasks is undocumented.
func (c *EC2) DescribeExportTasks(req DescribeExportTasksRequest) (resp *DescribeExportTasksResult, err error) {
	resp = &DescribeExportTasksResult{}
	err = c.client.Do("DescribeExportTasks", "POST", "/", req, resp)
	return
}

// DescribeImageAttribute describes the specified attribute of the
// specified You can specify only one attribute at a time.
func (c *EC2) DescribeImageAttribute(req DescribeImageAttributeRequest) (resp *ImageAttribute, err error) {
	resp = &ImageAttribute{}
	err = c.client.Do("DescribeImageAttribute", "POST", "/", req, resp)
	return
}

// DescribeImages describes one or more of the images (AMIs, AKIs, and
// ARIs) available to you. Images available to you include public images,
// private images that you own, and private images owned by other AWS
// accounts but for which you have explicit launch permissions.
// Deregistered images are included in the returned results for an
// unspecified interval after deregistration.
func (c *EC2) DescribeImages(req DescribeImagesRequest) (resp *DescribeImagesResult, err error) {
	resp = &DescribeImagesResult{}
	err = c.client.Do("DescribeImages", "POST", "/", req, resp)
	return
}

// DescribeInstanceAttribute describes the specified attribute of the
// specified instance. You can specify only one attribute at a time. Valid
// attribute values are: instanceType | kernel | ramdisk | userData |
// disableApiTermination | instanceInitiatedShutdownBehavior |
// rootDeviceName | blockDeviceMapping | productCodes | sourceDestCheck |
// groupSet | ebsOptimized | sriovNetSupport
func (c *EC2) DescribeInstanceAttribute(req DescribeInstanceAttributeRequest) (resp *InstanceAttribute, err error) {
	resp = &InstanceAttribute{}
	err = c.client.Do("DescribeInstanceAttribute", "POST", "/", req, resp)
	return
}

// DescribeInstanceStatus describes the status of one or more instances,
// including any scheduled events. Instance status has two main components:
// System Status reports impaired functionality that stems from issues
// related to the systems that support an instance, such as such as
// hardware failures and network connectivity problems. This call reports
// such problems as impaired reachability. Instance Status reports impaired
// functionality that arises from problems internal to the instance. This
// call reports such problems as impaired reachability. Instance status
// provides information about four types of scheduled events for an
// instance that may require your attention: Scheduled Reboot: When Amazon
// EC2 determines that an instance must be rebooted, the instances status
// returns one of two event codes: system-reboot or instance-reboot .
// System reboot commonly occurs if certain maintenance or upgrade
// operations require a reboot of the underlying host that supports an
// instance. Instance reboot commonly occurs if the instance must be
// rebooted, rather than the underlying host. Rebooting events include a
// scheduled start and end time. System Maintenance: When Amazon EC2
// determines that an instance requires maintenance that requires power or
// network impact, the instance status is the event code system-maintenance
// . System maintenance is either power maintenance or network maintenance.
// For power maintenance, your instance will be unavailable for a brief
// period of time and then rebooted. For network maintenance, your instance
// will experience a brief loss of network connectivity. System maintenance
// events include a scheduled start and end time. You will also be notified
// by email if one of your instances is set for system maintenance. The
// email message indicates when your instance is scheduled for maintenance.
// Scheduled Retirement: When Amazon EC2 determines that an instance must
// be shut down, the instance status is the event code instance-retirement
// . Retirement commonly occurs when the underlying host is degraded and
// must be replaced. Retirement events include a scheduled start and end
// time. You will also be notified by email if one of your instances is set
// to retiring. The email message indicates when your instance will be
// permanently retired. Scheduled Stop: When Amazon EC2 determines that an
// instance must be shut down, the instances status returns an event code
// called instance-stop . Stop events include a scheduled start and end
// time. You will also be notified by email if one of your instances is set
// to stop. The email message indicates when your instance will be stopped.
// When your instance is retired, it will either be terminated (if its root
// device type is the instance-store) or stopped (if its root device type
// is an EBS volume). Instances stopped due to retirement will not be
// restarted, but you can do so manually. You can also avoid retirement of
// EBS-backed instances by manually restarting your instance when its event
// code is instance-retirement . This ensures that your instance is started
// on a different underlying host. For more information about failed status
// checks, see Troubleshooting Instances with Failed Status Checks in the
// Amazon Elastic Compute Cloud User Guide . For more information about
// working with scheduled events, see Working with an Instance That Has a
// Scheduled Event in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeInstanceStatus(req DescribeInstanceStatusRequest) (resp *DescribeInstanceStatusResult, err error) {
	resp = &DescribeInstanceStatusResult{}
	err = c.client.Do("DescribeInstanceStatus", "POST", "/", req, resp)
	return
}

// DescribeInstances describes one or more of your instances. If you
// specify one or more instance IDs, Amazon EC2 returns information for
// those instances. If you do not specify instance IDs, Amazon EC2 returns
// information for all relevant instances. If you specify an instance ID
// that is not valid, an error is returned. If you specify an instance that
// you do not own, it is not included in the returned results. Recently
// terminated instances might appear in the returned results. This interval
// is usually less than one hour.
func (c *EC2) DescribeInstances(req DescribeInstancesRequest) (resp *DescribeInstancesResult, err error) {
	resp = &DescribeInstancesResult{}
	err = c.client.Do("DescribeInstances", "POST", "/", req, resp)
	return
}

// DescribeInternetGateways is undocumented.
func (c *EC2) DescribeInternetGateways(req DescribeInternetGatewaysRequest) (resp *DescribeInternetGatewaysResult, err error) {
	resp = &DescribeInternetGatewaysResult{}
	err = c.client.Do("DescribeInternetGateways", "POST", "/", req, resp)
	return
}

// DescribeKeyPairs describes one or more of your key pairs. For more
// information about key pairs, see Key Pairs in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) DescribeKeyPairs(req DescribeKeyPairsRequest) (resp *DescribeKeyPairsResult, err error) {
	resp = &DescribeKeyPairsResult{}
	err = c.client.Do("DescribeKeyPairs", "POST", "/", req, resp)
	return
}

// DescribeNetworkAcls describes one or more of your network ACLs. For more
// information about network ACLs, see Network ACLs in the Amazon Virtual
// Private Cloud User Guide
func (c *EC2) DescribeNetworkAcls(req DescribeNetworkAclsRequest) (resp *DescribeNetworkAclsResult, err error) {
	resp = &DescribeNetworkAclsResult{}
	err = c.client.Do("DescribeNetworkAcls", "POST", "/", req, resp)
	return
}

// DescribeNetworkInterfaceAttribute describes a network interface
// attribute. You can specify only one attribute at a time.
func (c *EC2) DescribeNetworkInterfaceAttribute(req DescribeNetworkInterfaceAttributeRequest) (resp *DescribeNetworkInterfaceAttributeResult, err error) {
	resp = &DescribeNetworkInterfaceAttributeResult{}
	err = c.client.Do("DescribeNetworkInterfaceAttribute", "POST", "/", req, resp)
	return
}

// DescribeNetworkInterfaces is undocumented.
func (c *EC2) DescribeNetworkInterfaces(req DescribeNetworkInterfacesRequest) (resp *DescribeNetworkInterfacesResult, err error) {
	resp = &DescribeNetworkInterfacesResult{}
	err = c.client.Do("DescribeNetworkInterfaces", "POST", "/", req, resp)
	return
}

// DescribePlacementGroups describes one or more of your placement groups.
// For more information about placement groups and cluster instances, see
// Cluster Instances in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribePlacementGroups(req DescribePlacementGroupsRequest) (resp *DescribePlacementGroupsResult, err error) {
	resp = &DescribePlacementGroupsResult{}
	err = c.client.Do("DescribePlacementGroups", "POST", "/", req, resp)
	return
}

// DescribeRegions describes one or more regions that are currently
// available to you. For a list of the regions supported by Amazon EC2, see
// Regions and Endpoints
func (c *EC2) DescribeRegions(req DescribeRegionsRequest) (resp *DescribeRegionsResult, err error) {
	resp = &DescribeRegionsResult{}
	err = c.client.Do("DescribeRegions", "POST", "/", req, resp)
	return
}

// DescribeReservedInstances describes one or more of the Reserved
// Instances that you purchased. For more information about Reserved
// Instances, see Reserved Instances in the Amazon Elastic Compute Cloud
// User Guide
func (c *EC2) DescribeReservedInstances(req DescribeReservedInstancesRequest) (resp *DescribeReservedInstancesResult, err error) {
	resp = &DescribeReservedInstancesResult{}
	err = c.client.Do("DescribeReservedInstances", "POST", "/", req, resp)
	return
}

// DescribeReservedInstancesListings describes your account's Reserved
// Instance listings in the Reserved Instance Marketplace. The Reserved
// Instance Marketplace matches sellers who want to resell Reserved
// Instance capacity that they no longer need with buyers who want to
// purchase additional capacity. Reserved Instances bought and sold through
// the Reserved Instance Marketplace work like any other Reserved
// Instances. As a seller, you choose to list some or all of your Reserved
// Instances, and you specify the upfront price to receive for them. Your
// Reserved Instances are then listed in the Reserved Instance Marketplace
// and are available for purchase. As a buyer, you specify the
// configuration of the Reserved Instance to purchase, and the Marketplace
// matches what you're searching for with what's available. The Marketplace
// first sells the lowest priced Reserved Instances to you, and continues
// to sell available Reserved Instance listings to you until your demand is
// met. You are charged based on the total price of all of the listings
// that you purchase. For more information, see Reserved Instance
// Marketplace in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeReservedInstancesListings(req DescribeReservedInstancesListingsRequest) (resp *DescribeReservedInstancesListingsResult, err error) {
	resp = &DescribeReservedInstancesListingsResult{}
	err = c.client.Do("DescribeReservedInstancesListings", "POST", "/", req, resp)
	return
}

// DescribeReservedInstancesModifications describes the modifications made
// to your Reserved Instances. If no parameter is specified, information
// about all your Reserved Instances modification requests is returned. If
// a modification ID is specified, only information about the specific
// modification is returned. For more information, see Modifying Reserved
// Instances in the Amazon Elastic Compute Cloud User Guide.
func (c *EC2) DescribeReservedInstancesModifications(req DescribeReservedInstancesModificationsRequest) (resp *DescribeReservedInstancesModificationsResult, err error) {
	resp = &DescribeReservedInstancesModificationsResult{}
	err = c.client.Do("DescribeReservedInstancesModifications", "POST", "/", req, resp)
	return
}

// DescribeReservedInstancesOfferings describes Reserved Instance offerings
// that are available for purchase. With Reserved Instances, you purchase
// the right to launch instances for a period of time. During that time
// period, you do not receive insufficient capacity errors, and you pay a
// lower usage rate than the rate charged for On-Demand instances for the
// actual time used. For more information, see Reserved Instance
// Marketplace in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeReservedInstancesOfferings(req DescribeReservedInstancesOfferingsRequest) (resp *DescribeReservedInstancesOfferingsResult, err error) {
	resp = &DescribeReservedInstancesOfferingsResult{}
	err = c.client.Do("DescribeReservedInstancesOfferings", "POST", "/", req, resp)
	return
}

// DescribeRouteTables describes one or more of your route tables. For more
// information about route tables, see Route Tables in the Amazon Virtual
// Private Cloud User Guide
func (c *EC2) DescribeRouteTables(req DescribeRouteTablesRequest) (resp *DescribeRouteTablesResult, err error) {
	resp = &DescribeRouteTablesResult{}
	err = c.client.Do("DescribeRouteTables", "POST", "/", req, resp)
	return
}

// DescribeSecurityGroups describes one or more of your security groups. A
// security group is for use with instances either in the EC2-Classic
// platform or in a specific For more information, see Amazon EC2 Security
// Groups in the Amazon Elastic Compute Cloud User Guide and Security
// Groups for Your in the Amazon Virtual Private Cloud User Guide
func (c *EC2) DescribeSecurityGroups(req DescribeSecurityGroupsRequest) (resp *DescribeSecurityGroupsResult, err error) {
	resp = &DescribeSecurityGroupsResult{}
	err = c.client.Do("DescribeSecurityGroups", "POST", "/", req, resp)
	return
}

// DescribeSnapshotAttribute describes the specified attribute of the
// specified snapshot. You can specify only one attribute at a time. For
// more information about Amazon EBS snapshots, see Amazon EBS Snapshots in
// the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeSnapshotAttribute(req DescribeSnapshotAttributeRequest) (resp *DescribeSnapshotAttributeResult, err error) {
	resp = &DescribeSnapshotAttributeResult{}
	err = c.client.Do("DescribeSnapshotAttribute", "POST", "/", req, resp)
	return
}

// DescribeSnapshots describes one or more of the Amazon EBS snapshots
// available to you. Available snapshots include public snapshots available
// for any AWS account to launch, private snapshots that you own, and
// private snapshots owned by another AWS account but for which you've been
// given explicit create volume permissions. The create volume permissions
// fall into the following categories: public : The owner of the snapshot
// granted create volume permissions for the snapshot to the all group. All
// AWS accounts have create volume permissions for these snapshots.
// explicit : The owner of the snapshot granted create volume permissions
// to a specific AWS account. implicit : An AWS account has implicit create
// volume permissions for all snapshots it owns. The list of snapshots
// returned can be modified by specifying snapshot IDs, snapshot owners, or
// AWS accounts with create volume permissions. If no options are
// specified, Amazon EC2 returns all snapshots for which you have create
// volume permissions. If you specify one or more snapshot IDs, only
// snapshots that have the specified IDs are returned. If you specify an
// invalid snapshot ID, an error is returned. If you specify a snapshot ID
// for which you do not have access, it is not included in the returned
// results. If you specify one or more snapshot owners, only snapshots from
// the specified owners and for which you have access are returned. The
// results can include the AWS account IDs of the specified owners, amazon
// for snapshots owned by Amazon, or self for snapshots that you own. If
// you specify a list of restorable users, only snapshots with create
// snapshot permissions for those users are returned. You can specify AWS
// account IDs (if you own the snapshots), self for snapshots for which you
// own or have explicit permissions, or all for public snapshots. For more
// information about Amazon EBS snapshots, see Amazon EBS Snapshots in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeSnapshots(req DescribeSnapshotsRequest) (resp *DescribeSnapshotsResult, err error) {
	resp = &DescribeSnapshotsResult{}
	err = c.client.Do("DescribeSnapshots", "POST", "/", req, resp)
	return
}

// DescribeSpotDatafeedSubscription describes the datafeed for Spot
// Instances. For more information, see Spot Instances in the Amazon
// Elastic Compute Cloud User Guide
func (c *EC2) DescribeSpotDatafeedSubscription(req DescribeSpotDatafeedSubscriptionRequest) (resp *DescribeSpotDatafeedSubscriptionResult, err error) {
	resp = &DescribeSpotDatafeedSubscriptionResult{}
	err = c.client.Do("DescribeSpotDatafeedSubscription", "POST", "/", req, resp)
	return
}

// DescribeSpotInstanceRequests describes the Spot Instance requests that
// belong to your account. Spot Instances are instances that Amazon EC2
// starts on your behalf when the maximum price that you specify exceeds
// the current Spot Price. Amazon EC2 periodically sets the Spot Price
// based on available Spot Instance capacity and current Spot Instance
// requests. For more information about Spot Instances, see Spot Instances
// in the Amazon Elastic Compute Cloud User Guide You can use
// DescribeSpotInstanceRequests to find a running Spot Instance by
// examining the response. If the status of the Spot Instance is fulfilled
// , the instance ID appears in the response and contains the identifier of
// the instance. Alternatively, you can use DescribeInstances with a filter
// to look for instances where the instance lifecycle is spot
func (c *EC2) DescribeSpotInstanceRequests(req DescribeSpotInstanceRequestsRequest) (resp *DescribeSpotInstanceRequestsResult, err error) {
	resp = &DescribeSpotInstanceRequestsResult{}
	err = c.client.Do("DescribeSpotInstanceRequests", "POST", "/", req, resp)
	return
}

// DescribeSpotPriceHistory describes the Spot Price history. Spot
// Instances are instances that Amazon EC2 starts on your behalf when the
// maximum price that you specify exceeds the current Spot Price. Amazon
// EC2 periodically sets the Spot Price based on available Spot Instance
// capacity and current Spot Instance requests. For more information about
// Spot Instances, see Spot Instances in the Amazon Elastic Compute Cloud
// User Guide When you specify an Availability Zone, this operation
// describes the price history for the specified Availability Zone with the
// most recent set of prices listed first. If you don't specify an
// Availability Zone, you get the prices across all Availability Zones,
// starting with the most recent set. However, if you're using an API
// version earlier than 2011-05-15, you get the lowest price across the
// region for the specified time period. The prices returned are listed in
// chronological order, from the oldest to the most recent. When you
// specify the start and end time options, this operation returns two
// pieces of data: the prices of the instance types within the time range
// that you specified and the time when the price changed. The price is
// valid within the time period that you specified; the response merely
// indicates the last time that the price changed.
func (c *EC2) DescribeSpotPriceHistory(req DescribeSpotPriceHistoryRequest) (resp *DescribeSpotPriceHistoryResult, err error) {
	resp = &DescribeSpotPriceHistoryResult{}
	err = c.client.Do("DescribeSpotPriceHistory", "POST", "/", req, resp)
	return
}

// DescribeSubnets describes one or more of your subnets. For more
// information about subnets, see Your VPC and Subnets in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) DescribeSubnets(req DescribeSubnetsRequest) (resp *DescribeSubnetsResult, err error) {
	resp = &DescribeSubnetsResult{}
	err = c.client.Do("DescribeSubnets", "POST", "/", req, resp)
	return
}

// DescribeTags describes one or more of the tags for your EC2 resources.
// For more information about tags, see Tagging Your Resources in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeTags(req DescribeTagsRequest) (resp *DescribeTagsResult, err error) {
	resp = &DescribeTagsResult{}
	err = c.client.Do("DescribeTags", "POST", "/", req, resp)
	return
}

// DescribeVolumeAttribute describes the specified attribute of the
// specified volume. You can specify only one attribute at a time. For more
// information about Amazon EBS volumes, see Amazon EBS Volumes in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeVolumeAttribute(req DescribeVolumeAttributeRequest) (resp *DescribeVolumeAttributeResult, err error) {
	resp = &DescribeVolumeAttributeResult{}
	err = c.client.Do("DescribeVolumeAttribute", "POST", "/", req, resp)
	return
}

// DescribeVolumeStatus describes the status of the specified volumes.
// Volume status provides the result of the checks performed on your
// volumes to determine events that can impair the performance of your
// volumes. The performance of a volume can be affected if an issue occurs
// on the volume's underlying host. If the volume's underlying host
// experiences a power outage or system issue, after the system is
// restored, there could be data inconsistencies on the volume. Volume
// events notify you if this occurs. Volume actions notify you if any
// action needs to be taken in response to the event. The
// DescribeVolumeStatus operation provides the following information about
// the specified volumes: Status : Reflects the current status of the
// volume. The possible values are ok , impaired , warning , or
// insufficient-data . If all checks pass, the overall status of the volume
// is ok . If the check fails, the overall status is impaired . If the
// status is insufficient-data , then the checks may still be taking place
// on your volume at the time. We recommend that you retry the request. For
// more information on volume status, see Monitoring the Status of Your
// Volumes Events : Reflect the cause of a volume status and may require
// you to take action. For example, if your volume returns an impaired
// status, then the volume event might be potential-data-inconsistency .
// This means that your volume has been affected by an issue with the
// underlying host, has all I/O operations disabled, and may have
// inconsistent data. Actions : Reflect the actions you may have to take in
// response to an event. For example, if the status of the volume is
// impaired and the volume event shows potential-data-inconsistency , then
// the action shows enable-volume-io . This means that you may want to
// enable the I/O operations for the volume by calling the EnableVolumeIO
// action and then check the volume for data consistency. Volume status is
// based on the volume status checks, and does not reflect the volume
// state. Therefore, volume status does not indicate volumes in the error
// state (for example, when a volume is incapable of accepting
func (c *EC2) DescribeVolumeStatus(req DescribeVolumeStatusRequest) (resp *DescribeVolumeStatusResult, err error) {
	resp = &DescribeVolumeStatusResult{}
	err = c.client.Do("DescribeVolumeStatus", "POST", "/", req, resp)
	return
}

// DescribeVolumes describes the specified Amazon EBS volumes. If you are
// describing a long list of volumes, you can paginate the output to make
// the list more manageable. The MaxResults parameter sets the maximum
// number of results returned in a single page. If the list of results
// exceeds your MaxResults value, then that number of results is returned
// along with a NextToken value that can be passed to a subsequent
// DescribeVolumes request to retrieve the remaining results. For more
// information about Amazon EBS volumes, see Amazon EBS Volumes in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) DescribeVolumes(req DescribeVolumesRequest) (resp *DescribeVolumesResult, err error) {
	resp = &DescribeVolumesResult{}
	err = c.client.Do("DescribeVolumes", "POST", "/", req, resp)
	return
}

// DescribeVpcAttribute describes the specified attribute of the specified
// You can specify only one attribute at a time.
func (c *EC2) DescribeVpcAttribute(req DescribeVpcAttributeRequest) (resp *DescribeVpcAttributeResult, err error) {
	resp = &DescribeVpcAttributeResult{}
	err = c.client.Do("DescribeVpcAttribute", "POST", "/", req, resp)
	return
}

// DescribeVpcPeeringConnections describes one or more of your VPC peering
// connections.
func (c *EC2) DescribeVpcPeeringConnections(req DescribeVpcPeeringConnectionsRequest) (resp *DescribeVpcPeeringConnectionsResult, err error) {
	resp = &DescribeVpcPeeringConnectionsResult{}
	err = c.client.Do("DescribeVpcPeeringConnections", "POST", "/", req, resp)
	return
}

// DescribeVpcs is undocumented.
func (c *EC2) DescribeVpcs(req DescribeVpcsRequest) (resp *DescribeVpcsResult, err error) {
	resp = &DescribeVpcsResult{}
	err = c.client.Do("DescribeVpcs", "POST", "/", req, resp)
	return
}

// DescribeVpnConnections describes one or more of your VPN connections.
// For more information about VPN connections, see Adding a Hardware
// Virtual Private Gateway to Your in the Amazon Virtual Private Cloud User
// Guide
func (c *EC2) DescribeVpnConnections(req DescribeVpnConnectionsRequest) (resp *DescribeVpnConnectionsResult, err error) {
	resp = &DescribeVpnConnectionsResult{}
	err = c.client.Do("DescribeVpnConnections", "POST", "/", req, resp)
	return
}

// DescribeVpnGateways describes one or more of your virtual private
// gateways. For more information about virtual private gateways, see
// Adding an IPsec Hardware VPN to Your in the Amazon Virtual Private Cloud
// User Guide
func (c *EC2) DescribeVpnGateways(req DescribeVpnGatewaysRequest) (resp *DescribeVpnGatewaysResult, err error) {
	resp = &DescribeVpnGatewaysResult{}
	err = c.client.Do("DescribeVpnGateways", "POST", "/", req, resp)
	return
}

// DetachInternetGateway detaches an Internet gateway from a disabling
// connectivity between the Internet and the The VPC must not contain any
// running instances with Elastic IP addresses.
func (c *EC2) DetachInternetGateway(req DetachInternetGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DetachInternetGateway", "POST", "/", req, nil)
	return
}

// DetachNetworkInterface is undocumented.
func (c *EC2) DetachNetworkInterface(req DetachNetworkInterfaceRequest) (err error) {
	// NRE
	err = c.client.Do("DetachNetworkInterface", "POST", "/", req, nil)
	return
}

// DetachVolume detaches an Amazon EBS volume from an instance. Make sure
// to unmount any file systems on the device within your operating system
// before detaching the volume. Failure to do so results in the volume
// being stuck in a busy state while detaching. If an Amazon EBS volume is
// the root device of an instance, it can't be detached while the instance
// is running. To detach the root volume, stop the instance first. If the
// root volume is detached from an instance with an AWS Marketplace product
// code, then the AWS Marketplace product codes from that volume are no
// longer associated with the instance. For more information, see Detaching
// an Amazon EBS Volume in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) DetachVolume(req DetachVolumeRequest) (resp *VolumeAttachment, err error) {
	resp = &VolumeAttachment{}
	err = c.client.Do("DetachVolume", "POST", "/", req, resp)
	return
}

// DetachVpnGateway detaches a virtual private gateway from a You do this
// if you're planning to turn off the VPC and not use it anymore. You can
// confirm a virtual private gateway has been completely detached from a
// VPC by describing the virtual private gateway (any attachments to the
// virtual private gateway are also described). You must wait for the
// attachment's state to switch to detached before you can delete the VPC
// or attach a different VPC to the virtual private gateway.
func (c *EC2) DetachVpnGateway(req DetachVpnGatewayRequest) (err error) {
	// NRE
	err = c.client.Do("DetachVpnGateway", "POST", "/", req, nil)
	return
}

// DisableVgwRoutePropagation disables a virtual private gateway from
// propagating routes to a specified route table of a
func (c *EC2) DisableVgwRoutePropagation(req DisableVgwRoutePropagationRequest) (err error) {
	// NRE
	err = c.client.Do("DisableVgwRoutePropagation", "POST", "/", req, nil)
	return
}

// DisassociateAddress disassociates an Elastic IP address from the
// instance or network interface it's associated with. An Elastic IP
// address is for use in either the EC2-Classic platform or in a For more
// information, see Elastic IP Addresses in the Amazon Elastic Compute
// Cloud User Guide This is an idempotent operation. If you perform the
// operation more than once, Amazon EC2 doesn't return an error.
func (c *EC2) DisassociateAddress(req DisassociateAddressRequest) (err error) {
	// NRE
	err = c.client.Do("DisassociateAddress", "POST", "/", req, nil)
	return
}

// DisassociateRouteTable disassociates a subnet from a route table. After
// you perform this action, the subnet no longer uses the routes in the
// route table. Instead, it uses the routes in the VPC's main route table.
// For more information about route tables, see Route Tables in the Amazon
// Virtual Private Cloud User Guide
func (c *EC2) DisassociateRouteTable(req DisassociateRouteTableRequest) (err error) {
	// NRE
	err = c.client.Do("DisassociateRouteTable", "POST", "/", req, nil)
	return
}

// EnableVgwRoutePropagation enables a virtual private gateway to propagate
// routes to the specified route table of a
func (c *EC2) EnableVgwRoutePropagation(req EnableVgwRoutePropagationRequest) (err error) {
	// NRE
	err = c.client.Do("EnableVgwRoutePropagation", "POST", "/", req, nil)
	return
}

// EnableVolumeIO enables I/O operations for a volume that had I/O
// operations disabled because the data on the volume was potentially
// inconsistent.
func (c *EC2) EnableVolumeIO(req EnableVolumeIORequest) (err error) {
	// NRE
	err = c.client.Do("EnableVolumeIO", "POST", "/", req, nil)
	return
}

// GetConsoleOutput gets the console output for the specified instance.
// Instances do not have a physical monitor through which you can view
// their console output. They also lack physical controls that allow you to
// power up, reboot, or shut them down. To allow these actions, we provide
// them through the Amazon EC2 API and command line interface. Instance
// console output is buffered and posted shortly after instance boot,
// reboot, and termination. Amazon EC2 preserves the most recent 64 KB
// output which is available for at least one hour after the most recent
// post. For Linux/Unix instances, the instance console output displays the
// exact console output that would normally be displayed on a physical
// monitor attached to a machine. This output is buffered because the
// instance produces it and then posts it to a store where the instance's
// owner can retrieve it. For Windows instances, the instance console
// output displays the last three system event log errors.
func (c *EC2) GetConsoleOutput(req GetConsoleOutputRequest) (resp *GetConsoleOutputResult, err error) {
	resp = &GetConsoleOutputResult{}
	err = c.client.Do("GetConsoleOutput", "POST", "/", req, resp)
	return
}

// GetPasswordData retrieves the encrypted administrator password for an
// instance running Windows. The Windows password is generated at boot if
// the EC2Config service plugin, Ec2SetPassword , is enabled. This usually
// only happens the first time an AMI is launched, and then Ec2SetPassword
// is automatically disabled. The password is not generated for rebundled
// AMIs unless Ec2SetPassword is enabled before bundling. The password is
// encrypted using the key pair that you specified when you launched the
// instance. You must provide the corresponding key pair file. Password
// generation and encryption takes a few moments. We recommend that you
// wait up to 15 minutes after launching an instance before trying to
// retrieve the generated password.
func (c *EC2) GetPasswordData(req GetPasswordDataRequest) (resp *GetPasswordDataResult, err error) {
	resp = &GetPasswordDataResult{}
	err = c.client.Do("GetPasswordData", "POST", "/", req, resp)
	return
}

// ImportInstance creates an import instance task using metadata from the
// specified disk image. After importing the image, you then upload it
// using the command in the EC2 command line tools. For more information,
// see Using the Command Line Tools to Import Your Virtual Machine to
// Amazon EC2
func (c *EC2) ImportInstance(req ImportInstanceRequest) (resp *ImportInstanceResult, err error) {
	resp = &ImportInstanceResult{}
	err = c.client.Do("ImportInstance", "POST", "/", req, resp)
	return
}

// ImportKeyPair imports the public key from an RSA key pair that you
// created with a third-party tool. Compare this with CreateKeyPair , in
// which AWS creates the key pair and gives the keys to you keeps a copy of
// the public key). With ImportKeyPair, you create the key pair and give
// AWS just the public key. The private key is never transferred between
// you and For more information about key pairs, see Key Pairs in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) ImportKeyPair(req ImportKeyPairRequest) (resp *ImportKeyPairResult, err error) {
	resp = &ImportKeyPairResult{}
	err = c.client.Do("ImportKeyPair", "POST", "/", req, resp)
	return
}

// ImportVolume creates an import volume task using metadata from the
// specified disk image. After importing the image, you then upload it
// using the command in the Amazon EC2 command-line interface tools. For
// more information, see Using the Command Line Tools to Import Your
// Virtual Machine to Amazon EC2 in the Amazon Elastic Compute Cloud User
// Guide
func (c *EC2) ImportVolume(req ImportVolumeRequest) (resp *ImportVolumeResult, err error) {
	resp = &ImportVolumeResult{}
	err = c.client.Do("ImportVolume", "POST", "/", req, resp)
	return
}

// ModifyImageAttribute modifies the specified attribute of the specified
// You can specify only one attribute at a time. AWS Marketplace product
// codes cannot be modified. Images with an AWS Marketplace product code
// cannot be made public.
func (c *EC2) ModifyImageAttribute(req ModifyImageAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyImageAttribute", "POST", "/", req, nil)
	return
}

// ModifyInstanceAttribute modifies the specified attribute of the
// specified instance. You can specify only one attribute at a time. To
// modify some attributes, the instance must be stopped. For more
// information, see Modifying Attributes of a Stopped Instance in the
// Amazon Elastic Compute Cloud User Guide
func (c *EC2) ModifyInstanceAttribute(req ModifyInstanceAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyInstanceAttribute", "POST", "/", req, nil)
	return
}

// ModifyNetworkInterfaceAttribute modifies the specified network interface
// attribute. You can specify only one attribute at a time.
func (c *EC2) ModifyNetworkInterfaceAttribute(req ModifyNetworkInterfaceAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyNetworkInterfaceAttribute", "POST", "/", req, nil)
	return
}

// ModifyReservedInstances modifies the Availability Zone, instance count,
// instance type, or network platform (EC2-Classic or EC2-VPC) of your
// Reserved Instances. The Reserved Instances to be modified must be
// identical, except for Availability Zone, network platform, and instance
// type. For more information, see Modifying Reserved Instances in the
// Amazon Elastic Compute Cloud User Guide.
func (c *EC2) ModifyReservedInstances(req ModifyReservedInstancesRequest) (resp *ModifyReservedInstancesResult, err error) {
	resp = &ModifyReservedInstancesResult{}
	err = c.client.Do("ModifyReservedInstances", "POST", "/", req, resp)
	return
}

// ModifySnapshotAttribute adds or removes permission settings for the
// specified snapshot. You may add or remove specified AWS account IDs from
// a snapshot's list of create volume permissions, but you cannot do both
// in a single API call. If you need to both add and remove account IDs for
// a snapshot, you must use multiple API calls. For more information on
// modifying snapshot permissions, see Sharing Snapshots in the Amazon
// Elastic Compute Cloud User Guide Snapshots with AWS Marketplace product
// codes cannot be made public.
func (c *EC2) ModifySnapshotAttribute(req ModifySnapshotAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifySnapshotAttribute", "POST", "/", req, nil)
	return
}

// ModifySubnetAttribute is undocumented.
func (c *EC2) ModifySubnetAttribute(req ModifySubnetAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifySubnetAttribute", "POST", "/", req, nil)
	return
}

// ModifyVolumeAttribute modifies a volume attribute. By default, all I/O
// operations for the volume are suspended when the data on the volume is
// determined to be potentially inconsistent, to prevent undetectable,
// latent data corruption. The I/O access to the volume can be resumed by
// first enabling I/O access and then checking the data consistency on your
// volume. You can change the default behavior to resume I/O operations. We
// recommend that you change this only for boot volumes or for volumes that
// are stateless or disposable.
func (c *EC2) ModifyVolumeAttribute(req ModifyVolumeAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyVolumeAttribute", "POST", "/", req, nil)
	return
}

// ModifyVpcAttribute is undocumented.
func (c *EC2) ModifyVpcAttribute(req ModifyVpcAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ModifyVpcAttribute", "POST", "/", req, nil)
	return
}

// MonitorInstances enables monitoring for a running instance. For more
// information about monitoring instances, see Monitoring Your Instances
// and Volumes in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) MonitorInstances(req MonitorInstancesRequest) (resp *MonitorInstancesResult, err error) {
	resp = &MonitorInstancesResult{}
	err = c.client.Do("MonitorInstances", "POST", "/", req, resp)
	return
}

// PurchaseReservedInstancesOffering purchases a Reserved Instance for use
// with your account. With Amazon EC2 Reserved Instances, you obtain a
// capacity reservation for a certain instance configuration over a
// specified period of time. You pay a lower usage rate than with On-Demand
// instances for the time that you actually use the capacity reservation.
// Use DescribeReservedInstancesOfferings to get a list of Reserved
// Instance offerings that match your specifications. After you've
// purchased a Reserved Instance, you can check for your new Reserved
// Instance with DescribeReservedInstances For more information, see
// Reserved Instances and Reserved Instance Marketplace in the Amazon
// Elastic Compute Cloud User Guide
func (c *EC2) PurchaseReservedInstancesOffering(req PurchaseReservedInstancesOfferingRequest) (resp *PurchaseReservedInstancesOfferingResult, err error) {
	resp = &PurchaseReservedInstancesOfferingResult{}
	err = c.client.Do("PurchaseReservedInstancesOffering", "POST", "/", req, resp)
	return
}

// RebootInstances requests a reboot of one or more instances. This
// operation is asynchronous; it only queues a request to reboot the
// specified instances. The operation succeeds if the instances are valid
// and belong to you. Requests to reboot terminated instances are ignored.
// If a Linux/Unix instance does not cleanly shut down within four minutes,
// Amazon EC2 performs a hard reboot. For more information about
// troubleshooting, see Getting Console Output and Rebooting Instances in
// the Amazon Elastic Compute Cloud User Guide
func (c *EC2) RebootInstances(req RebootInstancesRequest) (err error) {
	// NRE
	err = c.client.Do("RebootInstances", "POST", "/", req, nil)
	return
}

// RegisterImage registers an When you're creating an this is the final
// step you must complete before you can launch an instance from the For
// more information about creating AMIs, see Creating Your Own AMIs in the
// Amazon Elastic Compute Cloud User Guide For Amazon EBS-backed instances,
// CreateImage creates and registers the AMI in a single request, so you
// don't have to register the AMI yourself. You can also use RegisterImage
// to create an Amazon EBS-backed AMI from a snapshot of a root device
// volume. For more information, see Launching an Instance from a Snapshot
// in the Amazon Elastic Compute Cloud User Guide If needed, you can
// deregister an AMI at any time. Any modifications you make to an AMI
// backed by an instance store volume invalidates its registration. If you
// make changes to an image, deregister the previous image and register the
// new image. You can't register an image where a secondary (non-root)
// snapshot has AWS Marketplace product codes.
func (c *EC2) RegisterImage(req RegisterImageRequest) (resp *RegisterImageResult, err error) {
	resp = &RegisterImageResult{}
	err = c.client.Do("RegisterImage", "POST", "/", req, resp)
	return
}

// RejectVpcPeeringConnection rejects a VPC peering connection request. The
// VPC peering connection must be in the pending-acceptance state. Use the
// DescribeVpcPeeringConnections request to view your outstanding VPC
// peering connection requests. To delete an active VPC peering connection,
// or to delete a VPC peering connection request that you initiated, use
// DeleteVpcPeeringConnection
func (c *EC2) RejectVpcPeeringConnection(req RejectVpcPeeringConnectionRequest) (resp *RejectVpcPeeringConnectionResult, err error) {
	resp = &RejectVpcPeeringConnectionResult{}
	err = c.client.Do("RejectVpcPeeringConnection", "POST", "/", req, resp)
	return
}

// ReleaseAddress releases the specified Elastic IP address. After
// releasing an Elastic IP address, it is released to the IP address pool
// and might be unavailable to you. Be sure to update your DNS records and
// any servers or devices that communicate with the address. If you attempt
// to release an Elastic IP address that you already released, you'll get
// an AuthFailure error if the address is already allocated to another AWS
// account. [EC2-Classic, default Releasing an Elastic IP address
// automatically disassociates it from any instance that it's associated
// with. To disassociate an Elastic IP address without releasing it, use
// DisassociateAddress [Nondefault You must use DisassociateAddress to
// disassociate the Elastic IP address before you try to release it.
// Otherwise, Amazon EC2 returns an error InvalidIPAddress.InUse
func (c *EC2) ReleaseAddress(req ReleaseAddressRequest) (err error) {
	// NRE
	err = c.client.Do("ReleaseAddress", "POST", "/", req, nil)
	return
}

// ReplaceNetworkAclAssociation changes which network ACL a subnet is
// associated with. By default when you create a subnet, it's automatically
// associated with the default network For more information about network
// ACLs, see Network ACLs in the Amazon Virtual Private Cloud User Guide
func (c *EC2) ReplaceNetworkAclAssociation(req ReplaceNetworkAclAssociationRequest) (resp *ReplaceNetworkAclAssociationResult, err error) {
	resp = &ReplaceNetworkAclAssociationResult{}
	err = c.client.Do("ReplaceNetworkAclAssociation", "POST", "/", req, resp)
	return
}

// ReplaceNetworkAclEntry replaces an entry (rule) in a network For more
// information about network ACLs, see Network ACLs in the Amazon Virtual
// Private Cloud User Guide
func (c *EC2) ReplaceNetworkAclEntry(req ReplaceNetworkAclEntryRequest) (err error) {
	// NRE
	err = c.client.Do("ReplaceNetworkAclEntry", "POST", "/", req, nil)
	return
}

// ReplaceRoute replaces an existing route within a route table in a You
// must provide only one of the following: Internet gateway or virtual
// private gateway, NAT instance, VPC peering connection, or network
// interface. For more information about route tables, see Route Tables in
// the Amazon Virtual Private Cloud User Guide
func (c *EC2) ReplaceRoute(req ReplaceRouteRequest) (err error) {
	// NRE
	err = c.client.Do("ReplaceRoute", "POST", "/", req, nil)
	return
}

// ReplaceRouteTableAssociation changes the route table associated with a
// given subnet in a After the operation completes, the subnet uses the
// routes in the new route table it's associated with. For more information
// about route tables, see Route Tables in the Amazon Virtual Private Cloud
// User Guide You can also use ReplaceRouteTableAssociation to change which
// table is the main route table in the You just specify the main route
// table's association ID and the route table to be the new main route
// table.
func (c *EC2) ReplaceRouteTableAssociation(req ReplaceRouteTableAssociationRequest) (resp *ReplaceRouteTableAssociationResult, err error) {
	resp = &ReplaceRouteTableAssociationResult{}
	err = c.client.Do("ReplaceRouteTableAssociation", "POST", "/", req, resp)
	return
}

// ReportInstanceStatus submits feedback about the status of an instance.
// The instance must be in the running state. If your experience with the
// instance differs from the instance status returned by
// DescribeInstanceStatus , use ReportInstanceStatus to report your
// experience with the instance. Amazon EC2 collects this information to
// improve the accuracy of status checks. Use of this action does not
// change the value returned by DescribeInstanceStatus
func (c *EC2) ReportInstanceStatus(req ReportInstanceStatusRequest) (err error) {
	// NRE
	err = c.client.Do("ReportInstanceStatus", "POST", "/", req, nil)
	return
}

// RequestSpotInstances creates a Spot Instance request. Spot Instances are
// instances that Amazon EC2 starts on your behalf when the maximum price
// that you specify exceeds the current Spot Price. Amazon EC2 periodically
// sets the Spot Price based on available Spot Instance capacity and
// current Spot Instance requests. For more information about Spot
// Instances, see Spot Instances in the Amazon Elastic Compute Cloud User
// Guide Users must be subscribed to the required product to run an
// instance with AWS Marketplace product codes.
func (c *EC2) RequestSpotInstances(req RequestSpotInstancesRequest) (resp *RequestSpotInstancesResult, err error) {
	resp = &RequestSpotInstancesResult{}
	err = c.client.Do("RequestSpotInstances", "POST", "/", req, resp)
	return
}

// ResetImageAttribute resets an attribute of an AMI to its default value.
func (c *EC2) ResetImageAttribute(req ResetImageAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ResetImageAttribute", "POST", "/", req, nil)
	return
}

// ResetInstanceAttribute resets an attribute of an instance to its default
// value. To reset the kernel or ramdisk , the instance must be in a
// stopped state. To reset the SourceDestCheck , the instance can be either
// running or stopped. The SourceDestCheck attribute controls whether
// source/destination checking is enabled. The default value is true ,
// which means checking is enabled. This value must be false for a NAT
// instance to perform For more information, see NAT Instances in the
// Amazon Virtual Private Cloud User Guide
func (c *EC2) ResetInstanceAttribute(req ResetInstanceAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ResetInstanceAttribute", "POST", "/", req, nil)
	return
}

// ResetNetworkInterfaceAttribute resets a network interface attribute. You
// can specify only one attribute at a time.
func (c *EC2) ResetNetworkInterfaceAttribute(req ResetNetworkInterfaceAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ResetNetworkInterfaceAttribute", "POST", "/", req, nil)
	return
}

// ResetSnapshotAttribute resets permission settings for the specified
// snapshot. For more information on modifying snapshot permissions, see
// Sharing Snapshots in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) ResetSnapshotAttribute(req ResetSnapshotAttributeRequest) (err error) {
	// NRE
	err = c.client.Do("ResetSnapshotAttribute", "POST", "/", req, nil)
	return
}

// RevokeSecurityGroupEgress removes one or more egress rules from a
// security group for EC2-VPC. The values that you specify in the revoke
// request (for example, ports) must match the existing rule's values for
// the rule to be revoked. Each rule consists of the protocol and the range
// or source security group. For the TCP and UDP protocols, you must also
// specify the destination port or range of ports. For the protocol, you
// must also specify the type and code. Rule changes are propagated to
// instances within the security group as quickly as possible. However, a
// small delay might occur.
func (c *EC2) RevokeSecurityGroupEgress(req RevokeSecurityGroupEgressRequest) (err error) {
	// NRE
	err = c.client.Do("RevokeSecurityGroupEgress", "POST", "/", req, nil)
	return
}

// RevokeSecurityGroupIngress removes one or more ingress rules from a
// security group. The values that you specify in the revoke request (for
// example, ports) must match the existing rule's values for the rule to be
// removed. Each rule consists of the protocol and the range or source
// security group. For the TCP and UDP protocols, you must also specify the
// destination port or range of ports. For the protocol, you must also
// specify the type and code. Rule changes are propagated to instances
// within the security group as quickly as possible. However, a small delay
// might occur.
func (c *EC2) RevokeSecurityGroupIngress(req RevokeSecurityGroupIngressRequest) (err error) {
	// NRE
	err = c.client.Do("RevokeSecurityGroupIngress", "POST", "/", req, nil)
	return
}

// RunInstances launches the specified number of instances using an AMI for
// which you have permissions. When you launch an instance, it enters the
// pending state. After the instance is ready for you, it enters the
// running state. To check the state of your instance, call
// DescribeInstances If you don't specify a security group when launching
// an instance, Amazon EC2 uses the default security group. For more
// information, see Security Groups in the Amazon Elastic Compute Cloud
// User Guide Linux instances have access to the public key of the key pair
// at boot. You can use this key to provide secure access to the instance.
// Amazon EC2 public images use this feature to provide secure access
// without passwords. For more information, see Key Pairs in the Amazon
// Elastic Compute Cloud User Guide You can provide optional user data when
// launching an instance. For more information, see Instance Metadata in
// the Amazon Elastic Compute Cloud User Guide If any of the AMIs have a
// product code attached for which the user has not subscribed,
// RunInstances fails. T2 instance types can only be launched into a If you
// do not have a default or if you do not specify a subnet ID in the
// request, RunInstances fails. For more information about troubleshooting,
// see What To Do If An Instance Immediately Terminates , and
// Troubleshooting Connecting to Your Instance in the Amazon Elastic
// Compute Cloud User Guide
func (c *EC2) RunInstances(req RunInstancesRequest) (resp *Reservation, err error) {
	resp = &Reservation{}
	err = c.client.Do("RunInstances", "POST", "/", req, resp)
	return
}

// StartInstances starts an Amazon EBS-backed AMI that you've previously
// stopped. Instances that use Amazon EBS volumes as their root devices can
// be quickly stopped and started. When an instance is stopped, the compute
// resources are released and you are not billed for hourly instance usage.
// However, your root partition Amazon EBS volume remains, continues to
// persist your data, and you are charged for Amazon EBS volume usage. You
// can restart your instance at any time. Each time you transition an
// instance from stopped to started, Amazon EC2 charges a full instance
// hour, even if transitions happen multiple times within a single hour.
// Before stopping an instance, make sure it is in a state from which it
// can be restarted. Stopping an instance does not preserve data stored in
// Performing this operation on an instance that uses an instance store as
// its root device returns an error. For more information, see Stopping
// Instances in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) StartInstances(req StartInstancesRequest) (resp *StartInstancesResult, err error) {
	resp = &StartInstancesResult{}
	err = c.client.Do("StartInstances", "POST", "/", req, resp)
	return
}

// StopInstances stops an Amazon EBS-backed instance. Each time you
// transition an instance from stopped to started, Amazon EC2 charges a
// full instance hour, even if transitions happen multiple times within a
// single hour. You can't start or stop Spot Instances. Instances that use
// Amazon EBS volumes as their root devices can be quickly stopped and
// started. When an instance is stopped, the compute resources are released
// and you are not billed for hourly instance usage. However, your root
// partition Amazon EBS volume remains, continues to persist your data, and
// you are charged for Amazon EBS volume usage. You can restart your
// instance at any time. Before stopping an instance, make sure it is in a
// state from which it can be restarted. Stopping an instance does not
// preserve data stored in Performing this operation on an instance that
// uses an instance store as its root device returns an error. You can
// stop, start, and terminate EBS-backed instances. You can only terminate
// instance store-backed instances. What happens to an instance differs if
// you stop it or terminate it. For example, when you stop an instance, the
// root device and any other devices attached to the instance persist. When
// you terminate an instance, the root device and any other devices
// attached during the instance launch are automatically deleted. For more
// information about the differences between stopping and terminating
// instances, see Instance Lifecycle in the Amazon Elastic Compute Cloud
// User Guide For more information about troubleshooting, see
// Troubleshooting Stopping Your Instance in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) StopInstances(req StopInstancesRequest) (resp *StopInstancesResult, err error) {
	resp = &StopInstancesResult{}
	err = c.client.Do("StopInstances", "POST", "/", req, resp)
	return
}

// TerminateInstances shuts down one or more instances. This operation is
// idempotent; if you terminate an instance more than once, each call
// succeeds. Terminated instances remain visible after termination (for
// approximately one hour). By default, Amazon EC2 deletes all Amazon EBS
// volumes that were attached when the instance launched. Volumes attached
// after instance launch continue running. You can stop, start, and
// terminate EBS-backed instances. You can only terminate instance
// store-backed instances. What happens to an instance differs if you stop
// it or terminate it. For example, when you stop an instance, the root
// device and any other devices attached to the instance persist. When you
// terminate an instance, the root device and any other devices attached
// during the instance launch are automatically deleted. For more
// information about the differences between stopping and terminating
// instances, see Instance Lifecycle in the Amazon Elastic Compute Cloud
// User Guide For more information about troubleshooting, see
// Troubleshooting Terminating Your Instance in the Amazon Elastic Compute
// Cloud User Guide
func (c *EC2) TerminateInstances(req TerminateInstancesRequest) (resp *TerminateInstancesResult, err error) {
	resp = &TerminateInstancesResult{}
	err = c.client.Do("TerminateInstances", "POST", "/", req, resp)
	return
}

// UnassignPrivateIPAddresses unassigns one or more secondary private IP
// addresses from a network interface.
func (c *EC2) UnassignPrivateIPAddresses(req UnassignPrivateIPAddressesRequest) (err error) {
	// NRE
	err = c.client.Do("UnassignPrivateIpAddresses", "POST", "/", req, nil)
	return
}

// UnmonitorInstances disables monitoring for a running instance. For more
// information about monitoring instances, see Monitoring Your Instances
// and Volumes in the Amazon Elastic Compute Cloud User Guide
func (c *EC2) UnmonitorInstances(req UnmonitorInstancesRequest) (resp *UnmonitorInstancesResult, err error) {
	resp = &UnmonitorInstancesResult{}
	err = c.client.Do("UnmonitorInstances", "POST", "/", req, resp)
	return
}

// AcceptVpcPeeringConnectionRequest is undocumented.
type AcceptVpcPeeringConnectionRequest struct {
	DryRun                 bool   `ec2:"dryRun" xml:"dryRun"`
	VpcPeeringConnectionID string `ec2:"vpcPeeringConnectionId" xml:"vpcPeeringConnectionId"`
}

// AcceptVpcPeeringConnectionResult is undocumented.
type AcceptVpcPeeringConnectionResult struct {
	VpcPeeringConnection VpcPeeringConnection `ec2:"vpcPeeringConnection" xml:"vpcPeeringConnection"`
}

// AccountAttribute is undocumented.
type AccountAttribute struct {
	AttributeName   string                  `ec2:"attributeName" xml:"attributeName"`
	AttributeValues []AccountAttributeValue `ec2:"attributeValueSet" xml:"attributeValueSet>item"`
}

// AccountAttributeValue is undocumented.
type AccountAttributeValue struct {
	AttributeValue string `ec2:"attributeValue" xml:"attributeValue"`
}

// Address is undocumented.
type Address struct {
	AllocationID            string `ec2:"allocationId" xml:"allocationId"`
	AssociationID           string `ec2:"associationId" xml:"associationId"`
	Domain                  string `ec2:"domain" xml:"domain"`
	InstanceID              string `ec2:"instanceId" xml:"instanceId"`
	NetworkInterfaceID      string `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	NetworkInterfaceOwnerID string `ec2:"networkInterfaceOwnerId" xml:"networkInterfaceOwnerId"`
	PrivateIPAddress        string `ec2:"privateIpAddress" xml:"privateIpAddress"`
	PublicIP                string `ec2:"publicIp" xml:"publicIp"`
}

// AllocateAddressRequest is undocumented.
type AllocateAddressRequest struct {
	Domain string `ec2:"" xml:"Domain"`
	DryRun bool   `ec2:"dryRun" xml:"dryRun"`
}

// AllocateAddressResult is undocumented.
type AllocateAddressResult struct {
	AllocationID string `ec2:"allocationId" xml:"allocationId"`
	Domain       string `ec2:"domain" xml:"domain"`
	PublicIP     string `ec2:"publicIp" xml:"publicIp"`
}

// AssignPrivateIPAddressesRequest is undocumented.
type AssignPrivateIPAddressesRequest struct {
	AllowReassignment              bool     `ec2:"allowReassignment" xml:"allowReassignment"`
	NetworkInterfaceID             string   `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	PrivateIPAddresses             []string `ec2:"privateIpAddress" xml:"privateIpAddress>PrivateIpAddress"`
	SecondaryPrivateIPAddressCount int      `ec2:"secondaryPrivateIpAddressCount" xml:"secondaryPrivateIpAddressCount"`
}

// AssociateAddressRequest is undocumented.
type AssociateAddressRequest struct {
	AllocationID       string `ec2:"" xml:"AllocationId"`
	AllowReassociation bool   `ec2:"allowReassociation" xml:"allowReassociation"`
	DryRun             bool   `ec2:"dryRun" xml:"dryRun"`
	InstanceID         string `ec2:"" xml:"InstanceId"`
	NetworkInterfaceID string `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	PrivateIPAddress   string `ec2:"privateIpAddress" xml:"privateIpAddress"`
	PublicIP           string `ec2:"" xml:"PublicIp"`
}

// AssociateAddressResult is undocumented.
type AssociateAddressResult struct {
	AssociationID string `ec2:"associationId" xml:"associationId"`
}

// AssociateDhcpOptionsRequest is undocumented.
type AssociateDhcpOptionsRequest struct {
	DhcpOptionsID string `ec2:"" xml:"DhcpOptionsId"`
	DryRun        bool   `ec2:"dryRun" xml:"dryRun"`
	VpcID         string `ec2:"" xml:"VpcId"`
}

// AssociateRouteTableRequest is undocumented.
type AssociateRouteTableRequest struct {
	DryRun       bool   `ec2:"dryRun" xml:"dryRun"`
	RouteTableID string `ec2:"routeTableId" xml:"routeTableId"`
	SubnetID     string `ec2:"subnetId" xml:"subnetId"`
}

// AssociateRouteTableResult is undocumented.
type AssociateRouteTableResult struct {
	AssociationID string `ec2:"associationId" xml:"associationId"`
}

// AttachInternetGatewayRequest is undocumented.
type AttachInternetGatewayRequest struct {
	DryRun            bool   `ec2:"dryRun" xml:"dryRun"`
	InternetGatewayID string `ec2:"internetGatewayId" xml:"internetGatewayId"`
	VpcID             string `ec2:"vpcId" xml:"vpcId"`
}

// AttachNetworkInterfaceRequest is undocumented.
type AttachNetworkInterfaceRequest struct {
	DeviceIndex        int    `ec2:"deviceIndex" xml:"deviceIndex"`
	DryRun             bool   `ec2:"dryRun" xml:"dryRun"`
	InstanceID         string `ec2:"instanceId" xml:"instanceId"`
	NetworkInterfaceID string `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
}

// AttachNetworkInterfaceResult is undocumented.
type AttachNetworkInterfaceResult struct {
	AttachmentID string `ec2:"attachmentId" xml:"attachmentId"`
}

// AttachVolumeRequest is undocumented.
type AttachVolumeRequest struct {
	Device     string `ec2:"" xml:"Device"`
	DryRun     bool   `ec2:"dryRun" xml:"dryRun"`
	InstanceID string `ec2:"" xml:"InstanceId"`
	VolumeID   string `ec2:"" xml:"VolumeId"`
}

// AttachVpnGatewayRequest is undocumented.
type AttachVpnGatewayRequest struct {
	DryRun       bool   `ec2:"dryRun" xml:"dryRun"`
	VpcID        string `ec2:"" xml:"VpcId"`
	VpnGatewayID string `ec2:"" xml:"VpnGatewayId"`
}

// AttachVpnGatewayResult is undocumented.
type AttachVpnGatewayResult struct {
	VpcAttachment VpcAttachment `ec2:"attachment" xml:"attachment"`
}

// AttributeBooleanValue is undocumented.
type AttributeBooleanValue struct {
	Value bool `ec2:"value" xml:"value"`
}

// AttributeValue is undocumented.
type AttributeValue struct {
	Value string `ec2:"value" xml:"value"`
}

// AuthorizeSecurityGroupEgressRequest is undocumented.
type AuthorizeSecurityGroupEgressRequest struct {
	CidrIP                     string         `ec2:"cidrIp" xml:"cidrIp"`
	DryRun                     bool           `ec2:"dryRun" xml:"dryRun"`
	FromPort                   int            `ec2:"fromPort" xml:"fromPort"`
	GroupID                    string         `ec2:"groupId" xml:"groupId"`
	IPPermissions              []IPPermission `ec2:"ipPermissions" xml:"ipPermissions>item"`
	IPProtocol                 string         `ec2:"ipProtocol" xml:"ipProtocol"`
	SourceSecurityGroupName    string         `ec2:"sourceSecurityGroupName" xml:"sourceSecurityGroupName"`
	SourceSecurityGroupOwnerID string         `ec2:"sourceSecurityGroupOwnerId" xml:"sourceSecurityGroupOwnerId"`
	ToPort                     int            `ec2:"toPort" xml:"toPort"`
}

// AuthorizeSecurityGroupIngressRequest is undocumented.
type AuthorizeSecurityGroupIngressRequest struct {
	CidrIP                     string         `ec2:"" xml:"CidrIp"`
	DryRun                     bool           `ec2:"dryRun" xml:"dryRun"`
	FromPort                   int            `ec2:"" xml:"FromPort"`
	GroupID                    string         `ec2:"" xml:"GroupId"`
	GroupName                  string         `ec2:"" xml:"GroupName"`
	IPPermissions              []IPPermission `ec2:"" xml:"IpPermissions>item"`
	IPProtocol                 string         `ec2:"" xml:"IpProtocol"`
	SourceSecurityGroupName    string         `ec2:"" xml:"SourceSecurityGroupName"`
	SourceSecurityGroupOwnerID string         `ec2:"" xml:"SourceSecurityGroupOwnerId"`
	ToPort                     int            `ec2:"" xml:"ToPort"`
}

// AvailabilityZone is undocumented.
type AvailabilityZone struct {
	Messages   []AvailabilityZoneMessage `ec2:"messageSet" xml:"messageSet>item"`
	RegionName string                    `ec2:"regionName" xml:"regionName"`
	State      string                    `ec2:"zoneState" xml:"zoneState"`
	ZoneName   string                    `ec2:"zoneName" xml:"zoneName"`
}

// AvailabilityZoneMessage is undocumented.
type AvailabilityZoneMessage struct {
	Message string `ec2:"message" xml:"message"`
}

// BlobAttributeValue is undocumented.
type BlobAttributeValue struct {
	Value []byte `ec2:"value" xml:"value"`
}

// BlockDeviceMapping is undocumented.
type BlockDeviceMapping struct {
	DeviceName  string         `ec2:"deviceName" xml:"deviceName"`
	Ebs         EbsBlockDevice `ec2:"ebs" xml:"ebs"`
	NoDevice    string         `ec2:"noDevice" xml:"noDevice"`
	VirtualName string         `ec2:"virtualName" xml:"virtualName"`
}

// BundleInstanceRequest is undocumented.
type BundleInstanceRequest struct {
	DryRun     bool    `ec2:"dryRun" xml:"dryRun"`
	InstanceID string  `ec2:"" xml:"InstanceId"`
	Storage    Storage `ec2:"" xml:"Storage"`
}

// BundleInstanceResult is undocumented.
type BundleInstanceResult struct {
	BundleTask BundleTask `ec2:"bundleInstanceTask" xml:"bundleInstanceTask"`
}

// BundleTask is undocumented.
type BundleTask struct {
	BundleID        string          `ec2:"bundleId" xml:"bundleId"`
	BundleTaskError BundleTaskError `ec2:"error" xml:"error"`
	InstanceID      string          `ec2:"instanceId" xml:"instanceId"`
	Progress        string          `ec2:"progress" xml:"progress"`
	StartTime       time.Time       `ec2:"startTime" xml:"startTime"`
	State           string          `ec2:"state" xml:"state"`
	Storage         Storage         `ec2:"storage" xml:"storage"`
	UpdateTime      time.Time       `ec2:"updateTime" xml:"updateTime"`
}

// BundleTaskError is undocumented.
type BundleTaskError struct {
	Code    string `ec2:"code" xml:"code"`
	Message string `ec2:"message" xml:"message"`
}

// CancelBundleTaskRequest is undocumented.
type CancelBundleTaskRequest struct {
	BundleID string `ec2:"" xml:"BundleId"`
	DryRun   bool   `ec2:"dryRun" xml:"dryRun"`
}

// CancelBundleTaskResult is undocumented.
type CancelBundleTaskResult struct {
	BundleTask BundleTask `ec2:"bundleInstanceTask" xml:"bundleInstanceTask"`
}

// CancelConversionRequest is undocumented.
type CancelConversionRequest struct {
	ConversionTaskID string `ec2:"conversionTaskId" xml:"conversionTaskId"`
	DryRun           bool   `ec2:"dryRun" xml:"dryRun"`
	ReasonMessage    string `ec2:"reasonMessage" xml:"reasonMessage"`
}

// CancelExportTaskRequest is undocumented.
type CancelExportTaskRequest struct {
	ExportTaskID string `ec2:"exportTaskId" xml:"exportTaskId"`
}

// CancelReservedInstancesListingRequest is undocumented.
type CancelReservedInstancesListingRequest struct {
	ReservedInstancesListingID string `ec2:"reservedInstancesListingId" xml:"reservedInstancesListingId"`
}

// CancelReservedInstancesListingResult is undocumented.
type CancelReservedInstancesListingResult struct {
	ReservedInstancesListings []ReservedInstancesListing `ec2:"reservedInstancesListingsSet" xml:"reservedInstancesListingsSet>item"`
}

// CancelSpotInstanceRequestsRequest is undocumented.
type CancelSpotInstanceRequestsRequest struct {
	DryRun                 bool     `ec2:"dryRun" xml:"dryRun"`
	SpotInstanceRequestIds []string `ec2:"SpotInstanceRequestId" xml:"SpotInstanceRequestId>SpotInstanceRequestId"`
}

// CancelSpotInstanceRequestsResult is undocumented.
type CancelSpotInstanceRequestsResult struct {
	CancelledSpotInstanceRequests []CancelledSpotInstanceRequest `ec2:"spotInstanceRequestSet" xml:"spotInstanceRequestSet>item"`
}

// CancelledSpotInstanceRequest is undocumented.
type CancelledSpotInstanceRequest struct {
	SpotInstanceRequestID string `ec2:"spotInstanceRequestId" xml:"spotInstanceRequestId"`
	State                 string `ec2:"state" xml:"state"`
}

// ConfirmProductInstanceRequest is undocumented.
type ConfirmProductInstanceRequest struct {
	DryRun      bool   `ec2:"dryRun" xml:"dryRun"`
	InstanceID  string `ec2:"" xml:"InstanceId"`
	ProductCode string `ec2:"" xml:"ProductCode"`
}

// ConfirmProductInstanceResult is undocumented.
type ConfirmProductInstanceResult struct {
	OwnerID string `ec2:"ownerId" xml:"ownerId"`
}

// ConversionTask is undocumented.
type ConversionTask struct {
	ConversionTaskID string                    `ec2:"conversionTaskId" xml:"conversionTaskId"`
	ExpirationTime   string                    `ec2:"expirationTime" xml:"expirationTime"`
	ImportInstance   ImportInstanceTaskDetails `ec2:"importInstance" xml:"importInstance"`
	ImportVolume     ImportVolumeTaskDetails   `ec2:"importVolume" xml:"importVolume"`
	State            string                    `ec2:"state" xml:"state"`
	StatusMessage    string                    `ec2:"statusMessage" xml:"statusMessage"`
	Tags             []Tag                     `ec2:"tagSet" xml:"tagSet>item"`
}

// CopyImageRequest is undocumented.
type CopyImageRequest struct {
	ClientToken   string `ec2:"" xml:"ClientToken"`
	Description   string `ec2:"" xml:"Description"`
	DryRun        bool   `ec2:"dryRun" xml:"dryRun"`
	Name          string `ec2:"" xml:"Name"`
	SourceImageID string `ec2:"" xml:"SourceImageId"`
	SourceRegion  string `ec2:"" xml:"SourceRegion"`
}

// CopyImageResult is undocumented.
type CopyImageResult struct {
	ImageID string `ec2:"imageId" xml:"imageId"`
}

// CopySnapshotRequest is undocumented.
type CopySnapshotRequest struct {
	Description       string `ec2:"" xml:"Description"`
	DestinationRegion string `ec2:"destinationRegion" xml:"destinationRegion"`
	DryRun            bool   `ec2:"dryRun" xml:"dryRun"`
	PresignedURL      string `ec2:"presignedUrl" xml:"presignedUrl"`
	SourceRegion      string `ec2:"" xml:"SourceRegion"`
	SourceSnapshotID  string `ec2:"" xml:"SourceSnapshotId"`
}

// CopySnapshotResult is undocumented.
type CopySnapshotResult struct {
	SnapshotID string `ec2:"snapshotId" xml:"snapshotId"`
}

// CreateCustomerGatewayRequest is undocumented.
type CreateCustomerGatewayRequest struct {
	BgpAsn   int    `ec2:"" xml:"BgpAsn"`
	DryRun   bool   `ec2:"dryRun" xml:"dryRun"`
	PublicIP string `ec2:"IpAddress" xml:"IpAddress"`
	Type     string `ec2:"" xml:"Type"`
}

// CreateCustomerGatewayResult is undocumented.
type CreateCustomerGatewayResult struct {
	CustomerGateway CustomerGateway `ec2:"customerGateway" xml:"customerGateway"`
}

// CreateDhcpOptionsRequest is undocumented.
type CreateDhcpOptionsRequest struct {
	DhcpConfigurations []NewDhcpConfiguration `ec2:"dhcpConfiguration" xml:"dhcpConfiguration>item"`
	DryRun             bool                   `ec2:"dryRun" xml:"dryRun"`
}

// CreateDhcpOptionsResult is undocumented.
type CreateDhcpOptionsResult struct {
	DhcpOptions DhcpOptions `ec2:"dhcpOptions" xml:"dhcpOptions"`
}

// CreateImageRequest is undocumented.
type CreateImageRequest struct {
	BlockDeviceMappings []BlockDeviceMapping `ec2:"blockDeviceMapping" xml:"blockDeviceMapping>BlockDeviceMapping"`
	Description         string               `ec2:"description" xml:"description"`
	DryRun              bool                 `ec2:"dryRun" xml:"dryRun"`
	InstanceID          string               `ec2:"instanceId" xml:"instanceId"`
	Name                string               `ec2:"name" xml:"name"`
	NoReboot            bool                 `ec2:"noReboot" xml:"noReboot"`
}

// CreateImageResult is undocumented.
type CreateImageResult struct {
	ImageID string `ec2:"imageId" xml:"imageId"`
}

// CreateInstanceExportTaskRequest is undocumented.
type CreateInstanceExportTaskRequest struct {
	Description       string                      `ec2:"description" xml:"description"`
	ExportToS3Task    ExportToS3TaskSpecification `ec2:"exportToS3" xml:"exportToS3"`
	InstanceID        string                      `ec2:"instanceId" xml:"instanceId"`
	TargetEnvironment string                      `ec2:"targetEnvironment" xml:"targetEnvironment"`
}

// CreateInstanceExportTaskResult is undocumented.
type CreateInstanceExportTaskResult struct {
	ExportTask ExportTask `ec2:"exportTask" xml:"exportTask"`
}

// CreateInternetGatewayRequest is undocumented.
type CreateInternetGatewayRequest struct {
	DryRun bool `ec2:"dryRun" xml:"dryRun"`
}

// CreateInternetGatewayResult is undocumented.
type CreateInternetGatewayResult struct {
	InternetGateway InternetGateway `ec2:"internetGateway" xml:"internetGateway"`
}

// CreateKeyPairRequest is undocumented.
type CreateKeyPairRequest struct {
	DryRun  bool   `ec2:"dryRun" xml:"dryRun"`
	KeyName string `ec2:"" xml:"KeyName"`
}

// CreateNetworkAclEntryRequest is undocumented.
type CreateNetworkAclEntryRequest struct {
	CidrBlock    string       `ec2:"cidrBlock" xml:"cidrBlock"`
	DryRun       bool         `ec2:"dryRun" xml:"dryRun"`
	Egress       bool         `ec2:"egress" xml:"egress"`
	IcmpTypeCode IcmpTypeCode `ec2:"Icmp" xml:"Icmp"`
	NetworkAclID string       `ec2:"networkAclId" xml:"networkAclId"`
	PortRange    PortRange    `ec2:"portRange" xml:"portRange"`
	Protocol     string       `ec2:"protocol" xml:"protocol"`
	RuleAction   string       `ec2:"ruleAction" xml:"ruleAction"`
	RuleNumber   int          `ec2:"ruleNumber" xml:"ruleNumber"`
}

// CreateNetworkAclRequest is undocumented.
type CreateNetworkAclRequest struct {
	DryRun bool   `ec2:"dryRun" xml:"dryRun"`
	VpcID  string `ec2:"vpcId" xml:"vpcId"`
}

// CreateNetworkAclResult is undocumented.
type CreateNetworkAclResult struct {
	NetworkAcl NetworkAcl `ec2:"networkAcl" xml:"networkAcl"`
}

// CreateNetworkInterfaceRequest is undocumented.
type CreateNetworkInterfaceRequest struct {
	Description                    string                          `ec2:"description" xml:"description"`
	DryRun                         bool                            `ec2:"dryRun" xml:"dryRun"`
	Groups                         []string                        `ec2:"SecurityGroupId" xml:"SecurityGroupId>SecurityGroupId"`
	PrivateIPAddress               string                          `ec2:"privateIpAddress" xml:"privateIpAddress"`
	PrivateIPAddresses             []PrivateIPAddressSpecification `ec2:"privateIpAddresses" xml:"privateIpAddresses>item"`
	SecondaryPrivateIPAddressCount int                             `ec2:"secondaryPrivateIpAddressCount" xml:"secondaryPrivateIpAddressCount"`
	SubnetID                       string                          `ec2:"subnetId" xml:"subnetId"`
}

// CreateNetworkInterfaceResult is undocumented.
type CreateNetworkInterfaceResult struct {
	NetworkInterface NetworkInterface `ec2:"networkInterface" xml:"networkInterface"`
}

// CreatePlacementGroupRequest is undocumented.
type CreatePlacementGroupRequest struct {
	DryRun    bool   `ec2:"dryRun" xml:"dryRun"`
	GroupName string `ec2:"groupName" xml:"groupName"`
	Strategy  string `ec2:"strategy" xml:"strategy"`
}

// CreateReservedInstancesListingRequest is undocumented.
type CreateReservedInstancesListingRequest struct {
	ClientToken         string                       `ec2:"clientToken" xml:"clientToken"`
	InstanceCount       int                          `ec2:"instanceCount" xml:"instanceCount"`
	PriceSchedules      []PriceScheduleSpecification `ec2:"priceSchedules" xml:"priceSchedules>item"`
	ReservedInstancesID string                       `ec2:"reservedInstancesId" xml:"reservedInstancesId"`
}

// CreateReservedInstancesListingResult is undocumented.
type CreateReservedInstancesListingResult struct {
	ReservedInstancesListings []ReservedInstancesListing `ec2:"reservedInstancesListingsSet" xml:"reservedInstancesListingsSet>item"`
}

// CreateRouteRequest is undocumented.
type CreateRouteRequest struct {
	DestinationCidrBlock   string `ec2:"destinationCidrBlock" xml:"destinationCidrBlock"`
	DryRun                 bool   `ec2:"dryRun" xml:"dryRun"`
	GatewayID              string `ec2:"gatewayId" xml:"gatewayId"`
	InstanceID             string `ec2:"instanceId" xml:"instanceId"`
	NetworkInterfaceID     string `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	RouteTableID           string `ec2:"routeTableId" xml:"routeTableId"`
	VpcPeeringConnectionID string `ec2:"vpcPeeringConnectionId" xml:"vpcPeeringConnectionId"`
}

// CreateRouteTableRequest is undocumented.
type CreateRouteTableRequest struct {
	DryRun bool   `ec2:"dryRun" xml:"dryRun"`
	VpcID  string `ec2:"vpcId" xml:"vpcId"`
}

// CreateRouteTableResult is undocumented.
type CreateRouteTableResult struct {
	RouteTable RouteTable `ec2:"routeTable" xml:"routeTable"`
}

// CreateSecurityGroupRequest is undocumented.
type CreateSecurityGroupRequest struct {
	Description string `ec2:"GroupDescription" xml:"GroupDescription"`
	DryRun      bool   `ec2:"dryRun" xml:"dryRun"`
	GroupName   string `ec2:"" xml:"GroupName"`
	VpcID       string `ec2:"" xml:"VpcId"`
}

// CreateSecurityGroupResult is undocumented.
type CreateSecurityGroupResult struct {
	GroupID string `ec2:"groupId" xml:"groupId"`
}

// CreateSnapshotRequest is undocumented.
type CreateSnapshotRequest struct {
	Description string `ec2:"" xml:"Description"`
	DryRun      bool   `ec2:"dryRun" xml:"dryRun"`
	VolumeID    string `ec2:"" xml:"VolumeId"`
}

// CreateSpotDatafeedSubscriptionRequest is undocumented.
type CreateSpotDatafeedSubscriptionRequest struct {
	Bucket string `ec2:"bucket" xml:"bucket"`
	DryRun bool   `ec2:"dryRun" xml:"dryRun"`
	Prefix string `ec2:"prefix" xml:"prefix"`
}

// CreateSpotDatafeedSubscriptionResult is undocumented.
type CreateSpotDatafeedSubscriptionResult struct {
	SpotDatafeedSubscription SpotDatafeedSubscription `ec2:"spotDatafeedSubscription" xml:"spotDatafeedSubscription"`
}

// CreateSubnetRequest is undocumented.
type CreateSubnetRequest struct {
	AvailabilityZone string `ec2:"" xml:"AvailabilityZone"`
	CidrBlock        string `ec2:"" xml:"CidrBlock"`
	DryRun           bool   `ec2:"dryRun" xml:"dryRun"`
	VpcID            string `ec2:"" xml:"VpcId"`
}

// CreateSubnetResult is undocumented.
type CreateSubnetResult struct {
	Subnet Subnet `ec2:"subnet" xml:"subnet"`
}

// CreateTagsRequest is undocumented.
type CreateTagsRequest struct {
	DryRun    bool     `ec2:"dryRun" xml:"dryRun"`
	Resources []string `ec2:"ResourceId" xml:"ResourceId>member"`
	Tags      []Tag    `ec2:"Tag" xml:"Tag>item"`
}

// CreateVolumePermission is undocumented.
type CreateVolumePermission struct {
	Group  string `ec2:"group" xml:"group"`
	UserID string `ec2:"userId" xml:"userId"`
}

// CreateVolumePermissionModifications is undocumented.
type CreateVolumePermissionModifications struct {
	Add    []CreateVolumePermission `ec2:"" xml:"Add>item"`
	Remove []CreateVolumePermission `ec2:"" xml:"Remove>item"`
}

// CreateVolumeRequest is undocumented.
type CreateVolumeRequest struct {
	AvailabilityZone string `ec2:"" xml:"AvailabilityZone"`
	DryRun           bool   `ec2:"dryRun" xml:"dryRun"`
	Encrypted        bool   `ec2:"encrypted" xml:"encrypted"`
	Iops             int    `ec2:"" xml:"Iops"`
	KmsKeyID         string `ec2:"" xml:"KmsKeyId"`
	Size             int    `ec2:"" xml:"Size"`
	SnapshotID       string `ec2:"" xml:"SnapshotId"`
	VolumeType       string `ec2:"" xml:"VolumeType"`
}

// CreateVpcPeeringConnectionRequest is undocumented.
type CreateVpcPeeringConnectionRequest struct {
	DryRun      bool   `ec2:"dryRun" xml:"dryRun"`
	PeerOwnerID string `ec2:"peerOwnerId" xml:"peerOwnerId"`
	PeerVpcID   string `ec2:"peerVpcId" xml:"peerVpcId"`
	VpcID       string `ec2:"vpcId" xml:"vpcId"`
}

// CreateVpcPeeringConnectionResult is undocumented.
type CreateVpcPeeringConnectionResult struct {
	VpcPeeringConnection VpcPeeringConnection `ec2:"vpcPeeringConnection" xml:"vpcPeeringConnection"`
}

// CreateVpcRequest is undocumented.
type CreateVpcRequest struct {
	CidrBlock       string `ec2:"" xml:"CidrBlock"`
	DryRun          bool   `ec2:"dryRun" xml:"dryRun"`
	InstanceTenancy string `ec2:"instanceTenancy" xml:"instanceTenancy"`
}

// CreateVpcResult is undocumented.
type CreateVpcResult struct {
	Vpc Vpc `ec2:"vpc" xml:"vpc"`
}

// CreateVpnConnectionRequest is undocumented.
type CreateVpnConnectionRequest struct {
	CustomerGatewayID string                            `ec2:"" xml:"CustomerGatewayId"`
	DryRun            bool                              `ec2:"dryRun" xml:"dryRun"`
	Options           VpnConnectionOptionsSpecification `ec2:"options" xml:"options"`
	Type              string                            `ec2:"" xml:"Type"`
	VpnGatewayID      string                            `ec2:"" xml:"VpnGatewayId"`
}

// CreateVpnConnectionResult is undocumented.
type CreateVpnConnectionResult struct {
	VpnConnection VpnConnection `ec2:"vpnConnection" xml:"vpnConnection"`
}

// CreateVpnConnectionRouteRequest is undocumented.
type CreateVpnConnectionRouteRequest struct {
	DestinationCidrBlock string `ec2:"" xml:"DestinationCidrBlock"`
	VpnConnectionID      string `ec2:"" xml:"VpnConnectionId"`
}

// CreateVpnGatewayRequest is undocumented.
type CreateVpnGatewayRequest struct {
	AvailabilityZone string `ec2:"" xml:"AvailabilityZone"`
	DryRun           bool   `ec2:"dryRun" xml:"dryRun"`
	Type             string `ec2:"" xml:"Type"`
}

// CreateVpnGatewayResult is undocumented.
type CreateVpnGatewayResult struct {
	VpnGateway VpnGateway `ec2:"vpnGateway" xml:"vpnGateway"`
}

// CustomerGateway is undocumented.
type CustomerGateway struct {
	BgpAsn            string `ec2:"bgpAsn" xml:"bgpAsn"`
	CustomerGatewayID string `ec2:"customerGatewayId" xml:"customerGatewayId"`
	IPAddress         string `ec2:"ipAddress" xml:"ipAddress"`
	State             string `ec2:"state" xml:"state"`
	Tags              []Tag  `ec2:"tagSet" xml:"tagSet>item"`
	Type              string `ec2:"type" xml:"type"`
}

// DeleteCustomerGatewayRequest is undocumented.
type DeleteCustomerGatewayRequest struct {
	CustomerGatewayID string `ec2:"" xml:"CustomerGatewayId"`
	DryRun            bool   `ec2:"dryRun" xml:"dryRun"`
}

// DeleteDhcpOptionsRequest is undocumented.
type DeleteDhcpOptionsRequest struct {
	DhcpOptionsID string `ec2:"" xml:"DhcpOptionsId"`
	DryRun        bool   `ec2:"dryRun" xml:"dryRun"`
}

// DeleteInternetGatewayRequest is undocumented.
type DeleteInternetGatewayRequest struct {
	DryRun            bool   `ec2:"dryRun" xml:"dryRun"`
	InternetGatewayID string `ec2:"internetGatewayId" xml:"internetGatewayId"`
}

// DeleteKeyPairRequest is undocumented.
type DeleteKeyPairRequest struct {
	DryRun  bool   `ec2:"dryRun" xml:"dryRun"`
	KeyName string `ec2:"" xml:"KeyName"`
}

// DeleteNetworkAclEntryRequest is undocumented.
type DeleteNetworkAclEntryRequest struct {
	DryRun       bool   `ec2:"dryRun" xml:"dryRun"`
	Egress       bool   `ec2:"egress" xml:"egress"`
	NetworkAclID string `ec2:"networkAclId" xml:"networkAclId"`
	RuleNumber   int    `ec2:"ruleNumber" xml:"ruleNumber"`
}

// DeleteNetworkAclRequest is undocumented.
type DeleteNetworkAclRequest struct {
	DryRun       bool   `ec2:"dryRun" xml:"dryRun"`
	NetworkAclID string `ec2:"networkAclId" xml:"networkAclId"`
}

// DeleteNetworkInterfaceRequest is undocumented.
type DeleteNetworkInterfaceRequest struct {
	DryRun             bool   `ec2:"dryRun" xml:"dryRun"`
	NetworkInterfaceID string `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
}

// DeletePlacementGroupRequest is undocumented.
type DeletePlacementGroupRequest struct {
	DryRun    bool   `ec2:"dryRun" xml:"dryRun"`
	GroupName string `ec2:"groupName" xml:"groupName"`
}

// DeleteRouteRequest is undocumented.
type DeleteRouteRequest struct {
	DestinationCidrBlock string `ec2:"destinationCidrBlock" xml:"destinationCidrBlock"`
	DryRun               bool   `ec2:"dryRun" xml:"dryRun"`
	RouteTableID         string `ec2:"routeTableId" xml:"routeTableId"`
}

// DeleteRouteTableRequest is undocumented.
type DeleteRouteTableRequest struct {
	DryRun       bool   `ec2:"dryRun" xml:"dryRun"`
	RouteTableID string `ec2:"routeTableId" xml:"routeTableId"`
}

// DeleteSecurityGroupRequest is undocumented.
type DeleteSecurityGroupRequest struct {
	DryRun    bool   `ec2:"dryRun" xml:"dryRun"`
	GroupID   string `ec2:"" xml:"GroupId"`
	GroupName string `ec2:"" xml:"GroupName"`
}

// DeleteSnapshotRequest is undocumented.
type DeleteSnapshotRequest struct {
	DryRun     bool   `ec2:"dryRun" xml:"dryRun"`
	SnapshotID string `ec2:"" xml:"SnapshotId"`
}

// DeleteSpotDatafeedSubscriptionRequest is undocumented.
type DeleteSpotDatafeedSubscriptionRequest struct {
	DryRun bool `ec2:"dryRun" xml:"dryRun"`
}

// DeleteSubnetRequest is undocumented.
type DeleteSubnetRequest struct {
	DryRun   bool   `ec2:"dryRun" xml:"dryRun"`
	SubnetID string `ec2:"" xml:"SubnetId"`
}

// DeleteTagsRequest is undocumented.
type DeleteTagsRequest struct {
	DryRun    bool     `ec2:"dryRun" xml:"dryRun"`
	Resources []string `ec2:"resourceId" xml:"resourceId>member"`
	Tags      []Tag    `ec2:"tag" xml:"tag>item"`
}

// DeleteVolumeRequest is undocumented.
type DeleteVolumeRequest struct {
	DryRun   bool   `ec2:"dryRun" xml:"dryRun"`
	VolumeID string `ec2:"" xml:"VolumeId"`
}

// DeleteVpcPeeringConnectionRequest is undocumented.
type DeleteVpcPeeringConnectionRequest struct {
	DryRun                 bool   `ec2:"dryRun" xml:"dryRun"`
	VpcPeeringConnectionID string `ec2:"vpcPeeringConnectionId" xml:"vpcPeeringConnectionId"`
}

// DeleteVpcPeeringConnectionResult is undocumented.
type DeleteVpcPeeringConnectionResult struct {
	Return bool `ec2:"return" xml:"return"`
}

// DeleteVpcRequest is undocumented.
type DeleteVpcRequest struct {
	DryRun bool   `ec2:"dryRun" xml:"dryRun"`
	VpcID  string `ec2:"" xml:"VpcId"`
}

// DeleteVpnConnectionRequest is undocumented.
type DeleteVpnConnectionRequest struct {
	DryRun          bool   `ec2:"dryRun" xml:"dryRun"`
	VpnConnectionID string `ec2:"" xml:"VpnConnectionId"`
}

// DeleteVpnConnectionRouteRequest is undocumented.
type DeleteVpnConnectionRouteRequest struct {
	DestinationCidrBlock string `ec2:"" xml:"DestinationCidrBlock"`
	VpnConnectionID      string `ec2:"" xml:"VpnConnectionId"`
}

// DeleteVpnGatewayRequest is undocumented.
type DeleteVpnGatewayRequest struct {
	DryRun       bool   `ec2:"dryRun" xml:"dryRun"`
	VpnGatewayID string `ec2:"" xml:"VpnGatewayId"`
}

// DeregisterImageRequest is undocumented.
type DeregisterImageRequest struct {
	DryRun  bool   `ec2:"dryRun" xml:"dryRun"`
	ImageID string `ec2:"" xml:"ImageId"`
}

// DescribeAccountAttributesRequest is undocumented.
type DescribeAccountAttributesRequest struct {
	AttributeNames []string `ec2:"attributeName" xml:"attributeName>attributeName"`
	DryRun         bool     `ec2:"dryRun" xml:"dryRun"`
}

// DescribeAccountAttributesResult is undocumented.
type DescribeAccountAttributesResult struct {
	AccountAttributes []AccountAttribute `ec2:"accountAttributeSet" xml:"accountAttributeSet>item"`
}

// DescribeAddressesRequest is undocumented.
type DescribeAddressesRequest struct {
	AllocationIds []string `ec2:"AllocationId" xml:"AllocationId>AllocationId"`
	DryRun        bool     `ec2:"dryRun" xml:"dryRun"`
	Filters       []Filter `ec2:"Filter" xml:"Filter>Filter"`
	PublicIPs     []string `ec2:"PublicIp" xml:"PublicIp>PublicIp"`
}

// DescribeAddressesResult is undocumented.
type DescribeAddressesResult struct {
	Addresses []Address `ec2:"addressesSet" xml:"addressesSet>item"`
}

// DescribeAvailabilityZonesRequest is undocumented.
type DescribeAvailabilityZonesRequest struct {
	DryRun    bool     `ec2:"dryRun" xml:"dryRun"`
	Filters   []Filter `ec2:"Filter" xml:"Filter>Filter"`
	ZoneNames []string `ec2:"ZoneName" xml:"ZoneName>ZoneName"`
}

// DescribeAvailabilityZonesResult is undocumented.
type DescribeAvailabilityZonesResult struct {
	AvailabilityZones []AvailabilityZone `ec2:"availabilityZoneInfo" xml:"availabilityZoneInfo>item"`
}

// DescribeBundleTasksRequest is undocumented.
type DescribeBundleTasksRequest struct {
	BundleIds []string `ec2:"BundleId" xml:"BundleId>BundleId"`
	DryRun    bool     `ec2:"dryRun" xml:"dryRun"`
	Filters   []Filter `ec2:"Filter" xml:"Filter>Filter"`
}

// DescribeBundleTasksResult is undocumented.
type DescribeBundleTasksResult struct {
	BundleTasks []BundleTask `ec2:"bundleInstanceTasksSet" xml:"bundleInstanceTasksSet>item"`
}

// DescribeConversionTasksRequest is undocumented.
type DescribeConversionTasksRequest struct {
	ConversionTaskIds []string `ec2:"conversionTaskId" xml:"conversionTaskId>item"`
	DryRun            bool     `ec2:"dryRun" xml:"dryRun"`
	Filters           []Filter `ec2:"filter" xml:"filter>Filter"`
}

// DescribeConversionTasksResult is undocumented.
type DescribeConversionTasksResult struct {
	ConversionTasks []ConversionTask `ec2:"conversionTasks" xml:"conversionTasks>item"`
}

// DescribeCustomerGatewaysRequest is undocumented.
type DescribeCustomerGatewaysRequest struct {
	CustomerGatewayIds []string `ec2:"CustomerGatewayId" xml:"CustomerGatewayId>CustomerGatewayId"`
	DryRun             bool     `ec2:"dryRun" xml:"dryRun"`
	Filters            []Filter `ec2:"Filter" xml:"Filter>Filter"`
}

// DescribeCustomerGatewaysResult is undocumented.
type DescribeCustomerGatewaysResult struct {
	CustomerGateways []CustomerGateway `ec2:"customerGatewaySet" xml:"customerGatewaySet>item"`
}

// DescribeDhcpOptionsRequest is undocumented.
type DescribeDhcpOptionsRequest struct {
	DhcpOptionsIds []string `ec2:"DhcpOptionsId" xml:"DhcpOptionsId>DhcpOptionsId"`
	DryRun         bool     `ec2:"dryRun" xml:"dryRun"`
	Filters        []Filter `ec2:"Filter" xml:"Filter>Filter"`
}

// DescribeDhcpOptionsResult is undocumented.
type DescribeDhcpOptionsResult struct {
	DhcpOptions []DhcpOptions `ec2:"dhcpOptionsSet" xml:"dhcpOptionsSet>item"`
}

// DescribeExportTasksRequest is undocumented.
type DescribeExportTasksRequest struct {
	ExportTaskIds []string `ec2:"exportTaskId" xml:"exportTaskId>ExportTaskId"`
}

// DescribeExportTasksResult is undocumented.
type DescribeExportTasksResult struct {
	ExportTasks []ExportTask `ec2:"exportTaskSet" xml:"exportTaskSet>item"`
}

// DescribeImageAttributeRequest is undocumented.
type DescribeImageAttributeRequest struct {
	Attribute string `ec2:"" xml:"Attribute"`
	DryRun    bool   `ec2:"dryRun" xml:"dryRun"`
	ImageID   string `ec2:"" xml:"ImageId"`
}

// DescribeImagesRequest is undocumented.
type DescribeImagesRequest struct {
	DryRun          bool     `ec2:"dryRun" xml:"dryRun"`
	ExecutableUsers []string `ec2:"ExecutableBy" xml:"ExecutableBy>ExecutableBy"`
	Filters         []Filter `ec2:"Filter" xml:"Filter>Filter"`
	ImageIds        []string `ec2:"ImageId" xml:"ImageId>ImageId"`
	Owners          []string `ec2:"Owner" xml:"Owner>Owner"`
}

// DescribeImagesResult is undocumented.
type DescribeImagesResult struct {
	Images []Image `ec2:"imagesSet" xml:"imagesSet>item"`
}

// DescribeInstanceAttributeRequest is undocumented.
type DescribeInstanceAttributeRequest struct {
	Attribute  string `ec2:"attribute" xml:"attribute"`
	DryRun     bool   `ec2:"dryRun" xml:"dryRun"`
	InstanceID string `ec2:"instanceId" xml:"instanceId"`
}

// DescribeInstanceStatusRequest is undocumented.
type DescribeInstanceStatusRequest struct {
	DryRun              bool     `ec2:"dryRun" xml:"dryRun"`
	Filters             []Filter `ec2:"Filter" xml:"Filter>Filter"`
	IncludeAllInstances bool     `ec2:"includeAllInstances" xml:"includeAllInstances"`
	InstanceIds         []string `ec2:"InstanceId" xml:"InstanceId>InstanceId"`
	MaxResults          int      `ec2:"" xml:"MaxResults"`
	NextToken           string   `ec2:"" xml:"NextToken"`
}

// DescribeInstanceStatusResult is undocumented.
type DescribeInstanceStatusResult struct {
	InstanceStatuses []InstanceStatus `ec2:"instanceStatusSet" xml:"instanceStatusSet>item"`
	NextToken        string           `ec2:"nextToken" xml:"nextToken"`
}

// DescribeInstancesRequest is undocumented.
type DescribeInstancesRequest struct {
	DryRun      bool     `ec2:"dryRun" xml:"dryRun"`
	Filters     []Filter `ec2:"Filter" xml:"Filter>Filter"`
	InstanceIds []string `ec2:"InstanceId" xml:"InstanceId>InstanceId"`
	MaxResults  int      `ec2:"maxResults" xml:"maxResults"`
	NextToken   string   `ec2:"nextToken" xml:"nextToken"`
}

// DescribeInstancesResult is undocumented.
type DescribeInstancesResult struct {
	NextToken    string        `ec2:"nextToken" xml:"nextToken"`
	Reservations []Reservation `ec2:"reservationSet" xml:"reservationSet>item"`
}

// DescribeInternetGatewaysRequest is undocumented.
type DescribeInternetGatewaysRequest struct {
	DryRun             bool     `ec2:"dryRun" xml:"dryRun"`
	Filters            []Filter `ec2:"Filter" xml:"Filter>Filter"`
	InternetGatewayIds []string `ec2:"internetGatewayId" xml:"internetGatewayId>item"`
}

// DescribeInternetGatewaysResult is undocumented.
type DescribeInternetGatewaysResult struct {
	InternetGateways []InternetGateway `ec2:"internetGatewaySet" xml:"internetGatewaySet>item"`
}

// DescribeKeyPairsRequest is undocumented.
type DescribeKeyPairsRequest struct {
	DryRun   bool     `ec2:"dryRun" xml:"dryRun"`
	Filters  []Filter `ec2:"Filter" xml:"Filter>Filter"`
	KeyNames []string `ec2:"KeyName" xml:"KeyName>KeyName"`
}

// DescribeKeyPairsResult is undocumented.
type DescribeKeyPairsResult struct {
	KeyPairs []KeyPairInfo `ec2:"keySet" xml:"keySet>item"`
}

// DescribeNetworkAclsRequest is undocumented.
type DescribeNetworkAclsRequest struct {
	DryRun        bool     `ec2:"dryRun" xml:"dryRun"`
	Filters       []Filter `ec2:"Filter" xml:"Filter>Filter"`
	NetworkAclIds []string `ec2:"NetworkAclId" xml:"NetworkAclId>item"`
}

// DescribeNetworkAclsResult is undocumented.
type DescribeNetworkAclsResult struct {
	NetworkAcls []NetworkAcl `ec2:"networkAclSet" xml:"networkAclSet>item"`
}

// DescribeNetworkInterfaceAttributeRequest is undocumented.
type DescribeNetworkInterfaceAttributeRequest struct {
	Attribute          string `ec2:"attribute" xml:"attribute"`
	DryRun             bool   `ec2:"dryRun" xml:"dryRun"`
	NetworkInterfaceID string `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
}

// DescribeNetworkInterfaceAttributeResult is undocumented.
type DescribeNetworkInterfaceAttributeResult struct {
	Attachment         NetworkInterfaceAttachment `ec2:"attachment" xml:"attachment"`
	Description        AttributeValue             `ec2:"description" xml:"description"`
	Groups             []GroupIdentifier          `ec2:"groupSet" xml:"groupSet>item"`
	NetworkInterfaceID string                     `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	SourceDestCheck    AttributeBooleanValue      `ec2:"sourceDestCheck" xml:"sourceDestCheck"`
}

// DescribeNetworkInterfacesRequest is undocumented.
type DescribeNetworkInterfacesRequest struct {
	DryRun              bool     `ec2:"dryRun" xml:"dryRun"`
	Filters             []Filter `ec2:"filter" xml:"filter>Filter"`
	NetworkInterfaceIds []string `ec2:"NetworkInterfaceId" xml:"NetworkInterfaceId>item"`
}

// DescribeNetworkInterfacesResult is undocumented.
type DescribeNetworkInterfacesResult struct {
	NetworkInterfaces []NetworkInterface `ec2:"networkInterfaceSet" xml:"networkInterfaceSet>item"`
}

// DescribePlacementGroupsRequest is undocumented.
type DescribePlacementGroupsRequest struct {
	DryRun     bool     `ec2:"dryRun" xml:"dryRun"`
	Filters    []Filter `ec2:"Filter" xml:"Filter>Filter"`
	GroupNames []string `ec2:"groupName" xml:"groupName>member"`
}

// DescribePlacementGroupsResult is undocumented.
type DescribePlacementGroupsResult struct {
	PlacementGroups []PlacementGroup `ec2:"placementGroupSet" xml:"placementGroupSet>item"`
}

// DescribeRegionsRequest is undocumented.
type DescribeRegionsRequest struct {
	DryRun      bool     `ec2:"dryRun" xml:"dryRun"`
	Filters     []Filter `ec2:"Filter" xml:"Filter>Filter"`
	RegionNames []string `ec2:"RegionName" xml:"RegionName>RegionName"`
}

// DescribeRegionsResult is undocumented.
type DescribeRegionsResult struct {
	Regions []Region `ec2:"regionInfo" xml:"regionInfo>item"`
}

// DescribeReservedInstancesListingsRequest is undocumented.
type DescribeReservedInstancesListingsRequest struct {
	Filters                    []Filter `ec2:"filters" xml:"filters>Filter"`
	ReservedInstancesID        string   `ec2:"reservedInstancesId" xml:"reservedInstancesId"`
	ReservedInstancesListingID string   `ec2:"reservedInstancesListingId" xml:"reservedInstancesListingId"`
}

// DescribeReservedInstancesListingsResult is undocumented.
type DescribeReservedInstancesListingsResult struct {
	ReservedInstancesListings []ReservedInstancesListing `ec2:"reservedInstancesListingsSet" xml:"reservedInstancesListingsSet>item"`
}

// DescribeReservedInstancesModificationsRequest is undocumented.
type DescribeReservedInstancesModificationsRequest struct {
	Filters                          []Filter `ec2:"Filter" xml:"Filter>Filter"`
	NextToken                        string   `ec2:"nextToken" xml:"nextToken"`
	ReservedInstancesModificationIds []string `ec2:"ReservedInstancesModificationId" xml:"ReservedInstancesModificationId>ReservedInstancesModificationId"`
}

// DescribeReservedInstancesModificationsResult is undocumented.
type DescribeReservedInstancesModificationsResult struct {
	NextToken                      string                          `ec2:"nextToken" xml:"nextToken"`
	ReservedInstancesModifications []ReservedInstancesModification `ec2:"reservedInstancesModificationsSet" xml:"reservedInstancesModificationsSet>item"`
}

// DescribeReservedInstancesOfferingsRequest is undocumented.
type DescribeReservedInstancesOfferingsRequest struct {
	AvailabilityZone             string   `ec2:"" xml:"AvailabilityZone"`
	DryRun                       bool     `ec2:"dryRun" xml:"dryRun"`
	Filters                      []Filter `ec2:"Filter" xml:"Filter>Filter"`
	IncludeMarketplace           bool     `ec2:"" xml:"IncludeMarketplace"`
	InstanceTenancy              string   `ec2:"instanceTenancy" xml:"instanceTenancy"`
	InstanceType                 string   `ec2:"" xml:"InstanceType"`
	MaxDuration                  int64    `ec2:"" xml:"MaxDuration"`
	MaxInstanceCount             int      `ec2:"" xml:"MaxInstanceCount"`
	MaxResults                   int      `ec2:"maxResults" xml:"maxResults"`
	MinDuration                  int64    `ec2:"" xml:"MinDuration"`
	NextToken                    string   `ec2:"nextToken" xml:"nextToken"`
	OfferingType                 string   `ec2:"offeringType" xml:"offeringType"`
	ProductDescription           string   `ec2:"" xml:"ProductDescription"`
	ReservedInstancesOfferingIds []string `ec2:"ReservedInstancesOfferingId" xml:"ReservedInstancesOfferingId>member"`
}

// DescribeReservedInstancesOfferingsResult is undocumented.
type DescribeReservedInstancesOfferingsResult struct {
	NextToken                  string                      `ec2:"nextToken" xml:"nextToken"`
	ReservedInstancesOfferings []ReservedInstancesOffering `ec2:"reservedInstancesOfferingsSet" xml:"reservedInstancesOfferingsSet>item"`
}

// DescribeReservedInstancesRequest is undocumented.
type DescribeReservedInstancesRequest struct {
	DryRun               bool     `ec2:"dryRun" xml:"dryRun"`
	Filters              []Filter `ec2:"Filter" xml:"Filter>Filter"`
	OfferingType         string   `ec2:"offeringType" xml:"offeringType"`
	ReservedInstancesIds []string `ec2:"ReservedInstancesId" xml:"ReservedInstancesId>ReservedInstancesId"`
}

// DescribeReservedInstancesResult is undocumented.
type DescribeReservedInstancesResult struct {
	ReservedInstances []ReservedInstances `ec2:"reservedInstancesSet" xml:"reservedInstancesSet>item"`
}

// DescribeRouteTablesRequest is undocumented.
type DescribeRouteTablesRequest struct {
	DryRun        bool     `ec2:"dryRun" xml:"dryRun"`
	Filters       []Filter `ec2:"Filter" xml:"Filter>Filter"`
	RouteTableIds []string `ec2:"RouteTableId" xml:"RouteTableId>item"`
}

// DescribeRouteTablesResult is undocumented.
type DescribeRouteTablesResult struct {
	RouteTables []RouteTable `ec2:"routeTableSet" xml:"routeTableSet>item"`
}

// DescribeSecurityGroupsRequest is undocumented.
type DescribeSecurityGroupsRequest struct {
	DryRun     bool     `ec2:"dryRun" xml:"dryRun"`
	Filters    []Filter `ec2:"Filter" xml:"Filter>Filter"`
	GroupIds   []string `ec2:"GroupId" xml:"GroupId>groupId"`
	GroupNames []string `ec2:"GroupName" xml:"GroupName>GroupName"`
}

// DescribeSecurityGroupsResult is undocumented.
type DescribeSecurityGroupsResult struct {
	SecurityGroups []SecurityGroup `ec2:"securityGroupInfo" xml:"securityGroupInfo>item"`
}

// DescribeSnapshotAttributeRequest is undocumented.
type DescribeSnapshotAttributeRequest struct {
	Attribute  string `ec2:"" xml:"Attribute"`
	DryRun     bool   `ec2:"dryRun" xml:"dryRun"`
	SnapshotID string `ec2:"" xml:"SnapshotId"`
}

// DescribeSnapshotAttributeResult is undocumented.
type DescribeSnapshotAttributeResult struct {
	CreateVolumePermissions []CreateVolumePermission `ec2:"createVolumePermission" xml:"createVolumePermission>item"`
	ProductCodes            []ProductCode            `ec2:"productCodes" xml:"productCodes>item"`
	SnapshotID              string                   `ec2:"snapshotId" xml:"snapshotId"`
}

// DescribeSnapshotsRequest is undocumented.
type DescribeSnapshotsRequest struct {
	DryRun              bool     `ec2:"dryRun" xml:"dryRun"`
	Filters             []Filter `ec2:"Filter" xml:"Filter>Filter"`
	OwnerIds            []string `ec2:"Owner" xml:"Owner>Owner"`
	RestorableByUserIds []string `ec2:"RestorableBy" xml:"RestorableBy>member"`
	SnapshotIds         []string `ec2:"SnapshotId" xml:"SnapshotId>SnapshotId"`
}

// DescribeSnapshotsResult is undocumented.
type DescribeSnapshotsResult struct {
	Snapshots []Snapshot `ec2:"snapshotSet" xml:"snapshotSet>item"`
}

// DescribeSpotDatafeedSubscriptionRequest is undocumented.
type DescribeSpotDatafeedSubscriptionRequest struct {
	DryRun bool `ec2:"dryRun" xml:"dryRun"`
}

// DescribeSpotDatafeedSubscriptionResult is undocumented.
type DescribeSpotDatafeedSubscriptionResult struct {
	SpotDatafeedSubscription SpotDatafeedSubscription `ec2:"spotDatafeedSubscription" xml:"spotDatafeedSubscription"`
}

// DescribeSpotInstanceRequestsRequest is undocumented.
type DescribeSpotInstanceRequestsRequest struct {
	DryRun                 bool     `ec2:"dryRun" xml:"dryRun"`
	Filters                []Filter `ec2:"Filter" xml:"Filter>Filter"`
	SpotInstanceRequestIds []string `ec2:"SpotInstanceRequestId" xml:"SpotInstanceRequestId>SpotInstanceRequestId"`
}

// DescribeSpotInstanceRequestsResult is undocumented.
type DescribeSpotInstanceRequestsResult struct {
	SpotInstanceRequests []SpotInstanceRequest `ec2:"spotInstanceRequestSet" xml:"spotInstanceRequestSet>item"`
}

// DescribeSpotPriceHistoryRequest is undocumented.
type DescribeSpotPriceHistoryRequest struct {
	AvailabilityZone    string    `ec2:"availabilityZone" xml:"availabilityZone"`
	DryRun              bool      `ec2:"dryRun" xml:"dryRun"`
	EndTime             time.Time `ec2:"endTime" xml:"endTime"`
	Filters             []Filter  `ec2:"Filter" xml:"Filter>Filter"`
	InstanceTypes       []string  `ec2:"InstanceType" xml:"InstanceType>member"`
	MaxResults          int       `ec2:"maxResults" xml:"maxResults"`
	NextToken           string    `ec2:"nextToken" xml:"nextToken"`
	ProductDescriptions []string  `ec2:"ProductDescription" xml:"ProductDescription>member"`
	StartTime           time.Time `ec2:"startTime" xml:"startTime"`
}

// DescribeSpotPriceHistoryResult is undocumented.
type DescribeSpotPriceHistoryResult struct {
	NextToken        string      `ec2:"nextToken" xml:"nextToken"`
	SpotPriceHistory []SpotPrice `ec2:"spotPriceHistorySet" xml:"spotPriceHistorySet>item"`
}

// DescribeSubnetsRequest is undocumented.
type DescribeSubnetsRequest struct {
	DryRun    bool     `ec2:"dryRun" xml:"dryRun"`
	Filters   []Filter `ec2:"Filter" xml:"Filter>Filter"`
	SubnetIds []string `ec2:"SubnetId" xml:"SubnetId>SubnetId"`
}

// DescribeSubnetsResult is undocumented.
type DescribeSubnetsResult struct {
	Subnets []Subnet `ec2:"subnetSet" xml:"subnetSet>item"`
}

// DescribeTagsRequest is undocumented.
type DescribeTagsRequest struct {
	DryRun     bool     `ec2:"dryRun" xml:"dryRun"`
	Filters    []Filter `ec2:"Filter" xml:"Filter>Filter"`
	MaxResults int      `ec2:"maxResults" xml:"maxResults"`
	NextToken  string   `ec2:"nextToken" xml:"nextToken"`
}

// DescribeTagsResult is undocumented.
type DescribeTagsResult struct {
	NextToken string           `ec2:"nextToken" xml:"nextToken"`
	Tags      []TagDescription `ec2:"tagSet" xml:"tagSet>item"`
}

// DescribeVolumeAttributeRequest is undocumented.
type DescribeVolumeAttributeRequest struct {
	Attribute string `ec2:"" xml:"Attribute"`
	DryRun    bool   `ec2:"dryRun" xml:"dryRun"`
	VolumeID  string `ec2:"" xml:"VolumeId"`
}

// DescribeVolumeAttributeResult is undocumented.
type DescribeVolumeAttributeResult struct {
	AutoEnableIO AttributeBooleanValue `ec2:"autoEnableIO" xml:"autoEnableIO"`
	ProductCodes []ProductCode         `ec2:"productCodes" xml:"productCodes>item"`
	VolumeID     string                `ec2:"volumeId" xml:"volumeId"`
}

// DescribeVolumeStatusRequest is undocumented.
type DescribeVolumeStatusRequest struct {
	DryRun     bool     `ec2:"dryRun" xml:"dryRun"`
	Filters    []Filter `ec2:"Filter" xml:"Filter>Filter"`
	MaxResults int      `ec2:"" xml:"MaxResults"`
	NextToken  string   `ec2:"" xml:"NextToken"`
	VolumeIds  []string `ec2:"VolumeId" xml:"VolumeId>VolumeId"`
}

// DescribeVolumeStatusResult is undocumented.
type DescribeVolumeStatusResult struct {
	NextToken      string             `ec2:"nextToken" xml:"nextToken"`
	VolumeStatuses []VolumeStatusItem `ec2:"volumeStatusSet" xml:"volumeStatusSet>item"`
}

// DescribeVolumesRequest is undocumented.
type DescribeVolumesRequest struct {
	DryRun     bool     `ec2:"dryRun" xml:"dryRun"`
	Filters    []Filter `ec2:"Filter" xml:"Filter>Filter"`
	MaxResults int      `ec2:"maxResults" xml:"maxResults"`
	NextToken  string   `ec2:"nextToken" xml:"nextToken"`
	VolumeIds  []string `ec2:"VolumeId" xml:"VolumeId>VolumeId"`
}

// DescribeVolumesResult is undocumented.
type DescribeVolumesResult struct {
	NextToken string   `ec2:"nextToken" xml:"nextToken"`
	Volumes   []Volume `ec2:"volumeSet" xml:"volumeSet>item"`
}

// DescribeVpcAttributeRequest is undocumented.
type DescribeVpcAttributeRequest struct {
	Attribute string `ec2:"" xml:"Attribute"`
	DryRun    bool   `ec2:"dryRun" xml:"dryRun"`
	VpcID     string `ec2:"" xml:"VpcId"`
}

// DescribeVpcAttributeResult is undocumented.
type DescribeVpcAttributeResult struct {
	EnableDNSHostnames AttributeBooleanValue `ec2:"enableDnsHostnames" xml:"enableDnsHostnames"`
	EnableDNSSupport   AttributeBooleanValue `ec2:"enableDnsSupport" xml:"enableDnsSupport"`
	VpcID              string                `ec2:"vpcId" xml:"vpcId"`
}

// DescribeVpcPeeringConnectionsRequest is undocumented.
type DescribeVpcPeeringConnectionsRequest struct {
	DryRun                  bool     `ec2:"dryRun" xml:"dryRun"`
	Filters                 []Filter `ec2:"Filter" xml:"Filter>Filter"`
	VpcPeeringConnectionIds []string `ec2:"VpcPeeringConnectionId" xml:"VpcPeeringConnectionId>item"`
}

// DescribeVpcPeeringConnectionsResult is undocumented.
type DescribeVpcPeeringConnectionsResult struct {
	VpcPeeringConnections []VpcPeeringConnection `ec2:"vpcPeeringConnectionSet" xml:"vpcPeeringConnectionSet>item"`
}

// DescribeVpcsRequest is undocumented.
type DescribeVpcsRequest struct {
	DryRun  bool     `ec2:"dryRun" xml:"dryRun"`
	Filters []Filter `ec2:"Filter" xml:"Filter>Filter"`
	VpcIds  []string `ec2:"VpcId" xml:"VpcId>VpcId"`
}

// DescribeVpcsResult is undocumented.
type DescribeVpcsResult struct {
	Vpcs []Vpc `ec2:"vpcSet" xml:"vpcSet>item"`
}

// DescribeVpnConnectionsRequest is undocumented.
type DescribeVpnConnectionsRequest struct {
	DryRun           bool     `ec2:"dryRun" xml:"dryRun"`
	Filters          []Filter `ec2:"Filter" xml:"Filter>Filter"`
	VpnConnectionIds []string `ec2:"VpnConnectionId" xml:"VpnConnectionId>VpnConnectionId"`
}

// DescribeVpnConnectionsResult is undocumented.
type DescribeVpnConnectionsResult struct {
	VpnConnections []VpnConnection `ec2:"vpnConnectionSet" xml:"vpnConnectionSet>item"`
}

// DescribeVpnGatewaysRequest is undocumented.
type DescribeVpnGatewaysRequest struct {
	DryRun        bool     `ec2:"dryRun" xml:"dryRun"`
	Filters       []Filter `ec2:"Filter" xml:"Filter>Filter"`
	VpnGatewayIds []string `ec2:"VpnGatewayId" xml:"VpnGatewayId>VpnGatewayId"`
}

// DescribeVpnGatewaysResult is undocumented.
type DescribeVpnGatewaysResult struct {
	VpnGateways []VpnGateway `ec2:"vpnGatewaySet" xml:"vpnGatewaySet>item"`
}

// DetachInternetGatewayRequest is undocumented.
type DetachInternetGatewayRequest struct {
	DryRun            bool   `ec2:"dryRun" xml:"dryRun"`
	InternetGatewayID string `ec2:"internetGatewayId" xml:"internetGatewayId"`
	VpcID             string `ec2:"vpcId" xml:"vpcId"`
}

// DetachNetworkInterfaceRequest is undocumented.
type DetachNetworkInterfaceRequest struct {
	AttachmentID string `ec2:"attachmentId" xml:"attachmentId"`
	DryRun       bool   `ec2:"dryRun" xml:"dryRun"`
	Force        bool   `ec2:"force" xml:"force"`
}

// DetachVolumeRequest is undocumented.
type DetachVolumeRequest struct {
	Device     string `ec2:"" xml:"Device"`
	DryRun     bool   `ec2:"dryRun" xml:"dryRun"`
	Force      bool   `ec2:"" xml:"Force"`
	InstanceID string `ec2:"" xml:"InstanceId"`
	VolumeID   string `ec2:"" xml:"VolumeId"`
}

// DetachVpnGatewayRequest is undocumented.
type DetachVpnGatewayRequest struct {
	DryRun       bool   `ec2:"dryRun" xml:"dryRun"`
	VpcID        string `ec2:"" xml:"VpcId"`
	VpnGatewayID string `ec2:"" xml:"VpnGatewayId"`
}

// DhcpConfiguration is undocumented.
type DhcpConfiguration struct {
	Key    string           `ec2:"key" xml:"key"`
	Values []AttributeValue `ec2:"valueSet" xml:"valueSet>item"`
}

// DhcpOptions is undocumented.
type DhcpOptions struct {
	DhcpConfigurations []DhcpConfiguration `ec2:"dhcpConfigurationSet" xml:"dhcpConfigurationSet>item"`
	DhcpOptionsID      string              `ec2:"dhcpOptionsId" xml:"dhcpOptionsId"`
	Tags               []Tag               `ec2:"tagSet" xml:"tagSet>item"`
}

// DisableVgwRoutePropagationRequest is undocumented.
type DisableVgwRoutePropagationRequest struct {
	GatewayID    string `ec2:"" xml:"GatewayId"`
	RouteTableID string `ec2:"" xml:"RouteTableId"`
}

// DisassociateAddressRequest is undocumented.
type DisassociateAddressRequest struct {
	AssociationID string `ec2:"" xml:"AssociationId"`
	DryRun        bool   `ec2:"dryRun" xml:"dryRun"`
	PublicIP      string `ec2:"" xml:"PublicIp"`
}

// DisassociateRouteTableRequest is undocumented.
type DisassociateRouteTableRequest struct {
	AssociationID string `ec2:"associationId" xml:"associationId"`
	DryRun        bool   `ec2:"dryRun" xml:"dryRun"`
}

// DiskImage is undocumented.
type DiskImage struct {
	Description string          `ec2:"" xml:"Description"`
	Image       DiskImageDetail `ec2:"" xml:"Image"`
	Volume      VolumeDetail    `ec2:"" xml:"Volume"`
}

// DiskImageDescription is undocumented.
type DiskImageDescription struct {
	Checksum          string `ec2:"checksum" xml:"checksum"`
	Format            string `ec2:"format" xml:"format"`
	ImportManifestURL string `ec2:"importManifestUrl" xml:"importManifestUrl"`
	Size              int64  `ec2:"size" xml:"size"`
}

// DiskImageDetail is undocumented.
type DiskImageDetail struct {
	Bytes             int64  `ec2:"bytes" xml:"bytes"`
	Format            string `ec2:"format" xml:"format"`
	ImportManifestURL string `ec2:"importManifestUrl" xml:"importManifestUrl"`
}

// DiskImageVolumeDescription is undocumented.
type DiskImageVolumeDescription struct {
	ID   string `ec2:"id" xml:"id"`
	Size int64  `ec2:"size" xml:"size"`
}

// EbsBlockDevice is undocumented.
type EbsBlockDevice struct {
	DeleteOnTermination bool   `ec2:"deleteOnTermination" xml:"deleteOnTermination"`
	Encrypted           bool   `ec2:"encrypted" xml:"encrypted"`
	Iops                int    `ec2:"iops" xml:"iops"`
	SnapshotID          string `ec2:"snapshotId" xml:"snapshotId"`
	VolumeSize          int    `ec2:"volumeSize" xml:"volumeSize"`
	VolumeType          string `ec2:"volumeType" xml:"volumeType"`
}

// EbsInstanceBlockDevice is undocumented.
type EbsInstanceBlockDevice struct {
	AttachTime          time.Time `ec2:"attachTime" xml:"attachTime"`
	DeleteOnTermination bool      `ec2:"deleteOnTermination" xml:"deleteOnTermination"`
	Status              string    `ec2:"status" xml:"status"`
	VolumeID            string    `ec2:"volumeId" xml:"volumeId"`
}

// EbsInstanceBlockDeviceSpecification is undocumented.
type EbsInstanceBlockDeviceSpecification struct {
	DeleteOnTermination bool   `ec2:"deleteOnTermination" xml:"deleteOnTermination"`
	VolumeID            string `ec2:"volumeId" xml:"volumeId"`
}

// EnableVgwRoutePropagationRequest is undocumented.
type EnableVgwRoutePropagationRequest struct {
	GatewayID    string `ec2:"" xml:"GatewayId"`
	RouteTableID string `ec2:"" xml:"RouteTableId"`
}

// EnableVolumeIORequest is undocumented.
type EnableVolumeIORequest struct {
	DryRun   bool   `ec2:"dryRun" xml:"dryRun"`
	VolumeID string `ec2:"volumeId" xml:"volumeId"`
}

// ExportTask is undocumented.
type ExportTask struct {
	Description           string                `ec2:"description" xml:"description"`
	ExportTaskID          string                `ec2:"exportTaskId" xml:"exportTaskId"`
	ExportToS3Task        ExportToS3Task        `ec2:"exportToS3" xml:"exportToS3"`
	InstanceExportDetails InstanceExportDetails `ec2:"instanceExport" xml:"instanceExport"`
	State                 string                `ec2:"state" xml:"state"`
	StatusMessage         string                `ec2:"statusMessage" xml:"statusMessage"`
}

// ExportToS3Task is undocumented.
type ExportToS3Task struct {
	ContainerFormat string `ec2:"containerFormat" xml:"containerFormat"`
	DiskImageFormat string `ec2:"diskImageFormat" xml:"diskImageFormat"`
	S3Bucket        string `ec2:"s3Bucket" xml:"s3Bucket"`
	S3Key           string `ec2:"s3Key" xml:"s3Key"`
}

// ExportToS3TaskSpecification is undocumented.
type ExportToS3TaskSpecification struct {
	ContainerFormat string `ec2:"containerFormat" xml:"containerFormat"`
	DiskImageFormat string `ec2:"diskImageFormat" xml:"diskImageFormat"`
	S3Bucket        string `ec2:"s3Bucket" xml:"s3Bucket"`
	S3Prefix        string `ec2:"s3Prefix" xml:"s3Prefix"`
}

// Filter is undocumented.
type Filter struct {
	Name   string   `ec2:"" xml:"Name"`
	Values []string `ec2:"Value" xml:"Value>item"`
}

// GetConsoleOutputRequest is undocumented.
type GetConsoleOutputRequest struct {
	DryRun     bool   `ec2:"dryRun" xml:"dryRun"`
	InstanceID string `ec2:"" xml:"InstanceId"`
}

// GetConsoleOutputResult is undocumented.
type GetConsoleOutputResult struct {
	InstanceID string    `ec2:"instanceId" xml:"instanceId"`
	Output     string    `ec2:"output" xml:"output"`
	Timestamp  time.Time `ec2:"timestamp" xml:"timestamp"`
}

// GetPasswordDataRequest is undocumented.
type GetPasswordDataRequest struct {
	DryRun     bool   `ec2:"dryRun" xml:"dryRun"`
	InstanceID string `ec2:"" xml:"InstanceId"`
}

// GetPasswordDataResult is undocumented.
type GetPasswordDataResult struct {
	InstanceID   string    `ec2:"instanceId" xml:"instanceId"`
	PasswordData string    `ec2:"passwordData" xml:"passwordData"`
	Timestamp    time.Time `ec2:"timestamp" xml:"timestamp"`
}

// GroupIdentifier is undocumented.
type GroupIdentifier struct {
	GroupID   string `ec2:"groupId" xml:"groupId"`
	GroupName string `ec2:"groupName" xml:"groupName"`
}

// IamInstanceProfile is undocumented.
type IamInstanceProfile struct {
	ARN string `ec2:"arn" xml:"arn"`
	ID  string `ec2:"id" xml:"id"`
}

// IamInstanceProfileSpecification is undocumented.
type IamInstanceProfileSpecification struct {
	ARN  string `ec2:"arn" xml:"arn"`
	Name string `ec2:"name" xml:"name"`
}

// IcmpTypeCode is undocumented.
type IcmpTypeCode struct {
	Code int `ec2:"code" xml:"code"`
	Type int `ec2:"type" xml:"type"`
}

// Image is undocumented.
type Image struct {
	Architecture        string               `ec2:"architecture" xml:"architecture"`
	BlockDeviceMappings []BlockDeviceMapping `ec2:"blockDeviceMapping" xml:"blockDeviceMapping>item"`
	Description         string               `ec2:"description" xml:"description"`
	Hypervisor          string               `ec2:"hypervisor" xml:"hypervisor"`
	ImageID             string               `ec2:"imageId" xml:"imageId"`
	ImageLocation       string               `ec2:"imageLocation" xml:"imageLocation"`
	ImageOwnerAlias     string               `ec2:"imageOwnerAlias" xml:"imageOwnerAlias"`
	ImageType           string               `ec2:"imageType" xml:"imageType"`
	KernelID            string               `ec2:"kernelId" xml:"kernelId"`
	Name                string               `ec2:"name" xml:"name"`
	OwnerID             string               `ec2:"imageOwnerId" xml:"imageOwnerId"`
	Platform            string               `ec2:"platform" xml:"platform"`
	ProductCodes        []ProductCode        `ec2:"productCodes" xml:"productCodes>item"`
	Public              bool                 `ec2:"isPublic" xml:"isPublic"`
	RamdiskID           string               `ec2:"ramdiskId" xml:"ramdiskId"`
	RootDeviceName      string               `ec2:"rootDeviceName" xml:"rootDeviceName"`
	RootDeviceType      string               `ec2:"rootDeviceType" xml:"rootDeviceType"`
	SriovNetSupport     string               `ec2:"sriovNetSupport" xml:"sriovNetSupport"`
	State               string               `ec2:"imageState" xml:"imageState"`
	StateReason         StateReason          `ec2:"stateReason" xml:"stateReason"`
	Tags                []Tag                `ec2:"tagSet" xml:"tagSet>item"`
	VirtualizationType  string               `ec2:"virtualizationType" xml:"virtualizationType"`
}

// ImageAttribute is undocumented.
type ImageAttribute struct {
	BlockDeviceMappings []BlockDeviceMapping `ec2:"blockDeviceMapping" xml:"blockDeviceMapping>item"`
	Description         AttributeValue       `ec2:"description" xml:"description"`
	ImageID             string               `ec2:"imageId" xml:"imageId"`
	KernelID            AttributeValue       `ec2:"kernel" xml:"kernel"`
	LaunchPermissions   []LaunchPermission   `ec2:"launchPermission" xml:"launchPermission>item"`
	ProductCodes        []ProductCode        `ec2:"productCodes" xml:"productCodes>item"`
	RamdiskID           AttributeValue       `ec2:"ramdisk" xml:"ramdisk"`
	SriovNetSupport     AttributeValue       `ec2:"sriovNetSupport" xml:"sriovNetSupport"`
}

// ImportInstanceLaunchSpecification is undocumented.
type ImportInstanceLaunchSpecification struct {
	AdditionalInfo                    string    `ec2:"additionalInfo" xml:"additionalInfo"`
	Architecture                      string    `ec2:"architecture" xml:"architecture"`
	GroupIds                          []string  `ec2:"GroupId" xml:"GroupId>SecurityGroupId"`
	GroupNames                        []string  `ec2:"GroupName" xml:"GroupName>SecurityGroup"`
	InstanceInitiatedShutdownBehavior string    `ec2:"instanceInitiatedShutdownBehavior" xml:"instanceInitiatedShutdownBehavior"`
	InstanceType                      string    `ec2:"instanceType" xml:"instanceType"`
	Monitoring                        bool      `ec2:"monitoring" xml:"monitoring"`
	Placement                         Placement `ec2:"placement" xml:"placement"`
	PrivateIPAddress                  string    `ec2:"privateIpAddress" xml:"privateIpAddress"`
	SubnetID                          string    `ec2:"subnetId" xml:"subnetId"`
	UserData                          string    `ec2:"userData" xml:"userData"`
}

// ImportInstanceRequest is undocumented.
type ImportInstanceRequest struct {
	Description         string                            `ec2:"description" xml:"description"`
	DiskImages          []DiskImage                       `ec2:"diskImage" xml:"diskImage>member"`
	DryRun              bool                              `ec2:"dryRun" xml:"dryRun"`
	LaunchSpecification ImportInstanceLaunchSpecification `ec2:"launchSpecification" xml:"launchSpecification"`
	Platform            string                            `ec2:"platform" xml:"platform"`
}

// ImportInstanceResult is undocumented.
type ImportInstanceResult struct {
	ConversionTask ConversionTask `ec2:"conversionTask" xml:"conversionTask"`
}

// ImportInstanceTaskDetails is undocumented.
type ImportInstanceTaskDetails struct {
	Description string                           `ec2:"description" xml:"description"`
	InstanceID  string                           `ec2:"instanceId" xml:"instanceId"`
	Platform    string                           `ec2:"platform" xml:"platform"`
	Volumes     []ImportInstanceVolumeDetailItem `ec2:"volumes" xml:"volumes>item"`
}

// ImportInstanceVolumeDetailItem is undocumented.
type ImportInstanceVolumeDetailItem struct {
	AvailabilityZone string                     `ec2:"availabilityZone" xml:"availabilityZone"`
	BytesConverted   int64                      `ec2:"bytesConverted" xml:"bytesConverted"`
	Description      string                     `ec2:"description" xml:"description"`
	Image            DiskImageDescription       `ec2:"image" xml:"image"`
	Status           string                     `ec2:"status" xml:"status"`
	StatusMessage    string                     `ec2:"statusMessage" xml:"statusMessage"`
	Volume           DiskImageVolumeDescription `ec2:"volume" xml:"volume"`
}

// ImportKeyPairRequest is undocumented.
type ImportKeyPairRequest struct {
	DryRun            bool   `ec2:"dryRun" xml:"dryRun"`
	KeyName           string `ec2:"keyName" xml:"keyName"`
	PublicKeyMaterial []byte `ec2:"publicKeyMaterial" xml:"publicKeyMaterial"`
}

// ImportKeyPairResult is undocumented.
type ImportKeyPairResult struct {
	KeyFingerprint string `ec2:"keyFingerprint" xml:"keyFingerprint"`
	KeyName        string `ec2:"keyName" xml:"keyName"`
}

// ImportVolumeRequest is undocumented.
type ImportVolumeRequest struct {
	AvailabilityZone string          `ec2:"availabilityZone" xml:"availabilityZone"`
	Description      string          `ec2:"description" xml:"description"`
	DryRun           bool            `ec2:"dryRun" xml:"dryRun"`
	Image            DiskImageDetail `ec2:"image" xml:"image"`
	Volume           VolumeDetail    `ec2:"volume" xml:"volume"`
}

// ImportVolumeResult is undocumented.
type ImportVolumeResult struct {
	ConversionTask ConversionTask `ec2:"conversionTask" xml:"conversionTask"`
}

// ImportVolumeTaskDetails is undocumented.
type ImportVolumeTaskDetails struct {
	AvailabilityZone string                     `ec2:"availabilityZone" xml:"availabilityZone"`
	BytesConverted   int64                      `ec2:"bytesConverted" xml:"bytesConverted"`
	Description      string                     `ec2:"description" xml:"description"`
	Image            DiskImageDescription       `ec2:"image" xml:"image"`
	Volume           DiskImageVolumeDescription `ec2:"volume" xml:"volume"`
}

// Instance is undocumented.
type Instance struct {
	AmiLaunchIndex        int                          `ec2:"amiLaunchIndex" xml:"amiLaunchIndex"`
	Architecture          string                       `ec2:"architecture" xml:"architecture"`
	BlockDeviceMappings   []InstanceBlockDeviceMapping `ec2:"blockDeviceMapping" xml:"blockDeviceMapping>item"`
	ClientToken           string                       `ec2:"clientToken" xml:"clientToken"`
	EbsOptimized          bool                         `ec2:"ebsOptimized" xml:"ebsOptimized"`
	Hypervisor            string                       `ec2:"hypervisor" xml:"hypervisor"`
	IamInstanceProfile    IamInstanceProfile           `ec2:"iamInstanceProfile" xml:"iamInstanceProfile"`
	ImageID               string                       `ec2:"imageId" xml:"imageId"`
	InstanceID            string                       `ec2:"instanceId" xml:"instanceId"`
	InstanceLifecycle     string                       `ec2:"instanceLifecycle" xml:"instanceLifecycle"`
	InstanceType          string                       `ec2:"instanceType" xml:"instanceType"`
	KernelID              string                       `ec2:"kernelId" xml:"kernelId"`
	KeyName               string                       `ec2:"keyName" xml:"keyName"`
	LaunchTime            time.Time                    `ec2:"launchTime" xml:"launchTime"`
	Monitoring            Monitoring                   `ec2:"monitoring" xml:"monitoring"`
	NetworkInterfaces     []InstanceNetworkInterface   `ec2:"networkInterfaceSet" xml:"networkInterfaceSet>item"`
	Placement             Placement                    `ec2:"placement" xml:"placement"`
	Platform              string                       `ec2:"platform" xml:"platform"`
	PrivateDNSName        string                       `ec2:"privateDnsName" xml:"privateDnsName"`
	PrivateIPAddress      string                       `ec2:"privateIpAddress" xml:"privateIpAddress"`
	ProductCodes          []ProductCode                `ec2:"productCodes" xml:"productCodes>item"`
	PublicDNSName         string                       `ec2:"dnsName" xml:"dnsName"`
	PublicIPAddress       string                       `ec2:"ipAddress" xml:"ipAddress"`
	RamdiskID             string                       `ec2:"ramdiskId" xml:"ramdiskId"`
	RootDeviceName        string                       `ec2:"rootDeviceName" xml:"rootDeviceName"`
	RootDeviceType        string                       `ec2:"rootDeviceType" xml:"rootDeviceType"`
	SecurityGroups        []GroupIdentifier            `ec2:"groupSet" xml:"groupSet>item"`
	SourceDestCheck       bool                         `ec2:"sourceDestCheck" xml:"sourceDestCheck"`
	SpotInstanceRequestID string                       `ec2:"spotInstanceRequestId" xml:"spotInstanceRequestId"`
	SriovNetSupport       string                       `ec2:"sriovNetSupport" xml:"sriovNetSupport"`
	State                 InstanceState                `ec2:"instanceState" xml:"instanceState"`
	StateReason           StateReason                  `ec2:"stateReason" xml:"stateReason"`
	StateTransitionReason string                       `ec2:"reason" xml:"reason"`
	SubnetID              string                       `ec2:"subnetId" xml:"subnetId"`
	Tags                  []Tag                        `ec2:"tagSet" xml:"tagSet>item"`
	VirtualizationType    string                       `ec2:"virtualizationType" xml:"virtualizationType"`
	VpcID                 string                       `ec2:"vpcId" xml:"vpcId"`
}

// InstanceAttribute is undocumented.
type InstanceAttribute struct {
	BlockDeviceMappings               []InstanceBlockDeviceMapping `ec2:"blockDeviceMapping" xml:"blockDeviceMapping>item"`
	DisableAPITermination             AttributeBooleanValue        `ec2:"disableApiTermination" xml:"disableApiTermination"`
	EbsOptimized                      AttributeBooleanValue        `ec2:"ebsOptimized" xml:"ebsOptimized"`
	Groups                            []GroupIdentifier            `ec2:"groupSet" xml:"groupSet>item"`
	InstanceID                        string                       `ec2:"instanceId" xml:"instanceId"`
	InstanceInitiatedShutdownBehavior AttributeValue               `ec2:"instanceInitiatedShutdownBehavior" xml:"instanceInitiatedShutdownBehavior"`
	InstanceType                      AttributeValue               `ec2:"instanceType" xml:"instanceType"`
	KernelID                          AttributeValue               `ec2:"kernel" xml:"kernel"`
	ProductCodes                      []ProductCode                `ec2:"productCodes" xml:"productCodes>item"`
	RamdiskID                         AttributeValue               `ec2:"ramdisk" xml:"ramdisk"`
	RootDeviceName                    AttributeValue               `ec2:"rootDeviceName" xml:"rootDeviceName"`
	SourceDestCheck                   AttributeBooleanValue        `ec2:"sourceDestCheck" xml:"sourceDestCheck"`
	SriovNetSupport                   AttributeValue               `ec2:"sriovNetSupport" xml:"sriovNetSupport"`
	UserData                          AttributeValue               `ec2:"userData" xml:"userData"`
}

// InstanceBlockDeviceMapping is undocumented.
type InstanceBlockDeviceMapping struct {
	DeviceName string                 `ec2:"deviceName" xml:"deviceName"`
	Ebs        EbsInstanceBlockDevice `ec2:"ebs" xml:"ebs"`
}

// InstanceBlockDeviceMappingSpecification is undocumented.
type InstanceBlockDeviceMappingSpecification struct {
	DeviceName  string                              `ec2:"deviceName" xml:"deviceName"`
	Ebs         EbsInstanceBlockDeviceSpecification `ec2:"ebs" xml:"ebs"`
	NoDevice    string                              `ec2:"noDevice" xml:"noDevice"`
	VirtualName string                              `ec2:"virtualName" xml:"virtualName"`
}

// InstanceCount is undocumented.
type InstanceCount struct {
	InstanceCount int    `ec2:"instanceCount" xml:"instanceCount"`
	State         string `ec2:"state" xml:"state"`
}

// InstanceExportDetails is undocumented.
type InstanceExportDetails struct {
	InstanceID        string `ec2:"instanceId" xml:"instanceId"`
	TargetEnvironment string `ec2:"targetEnvironment" xml:"targetEnvironment"`
}

// InstanceMonitoring is undocumented.
type InstanceMonitoring struct {
	InstanceID string     `ec2:"instanceId" xml:"instanceId"`
	Monitoring Monitoring `ec2:"monitoring" xml:"monitoring"`
}

// InstanceNetworkInterface is undocumented.
type InstanceNetworkInterface struct {
	Association        InstanceNetworkInterfaceAssociation `ec2:"association" xml:"association"`
	Attachment         InstanceNetworkInterfaceAttachment  `ec2:"attachment" xml:"attachment"`
	Description        string                              `ec2:"description" xml:"description"`
	Groups             []GroupIdentifier                   `ec2:"groupSet" xml:"groupSet>item"`
	MacAddress         string                              `ec2:"macAddress" xml:"macAddress"`
	NetworkInterfaceID string                              `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	OwnerID            string                              `ec2:"ownerId" xml:"ownerId"`
	PrivateDNSName     string                              `ec2:"privateDnsName" xml:"privateDnsName"`
	PrivateIPAddress   string                              `ec2:"privateIpAddress" xml:"privateIpAddress"`
	PrivateIPAddresses []InstancePrivateIPAddress          `ec2:"privateIpAddressesSet" xml:"privateIpAddressesSet>item"`
	SourceDestCheck    bool                                `ec2:"sourceDestCheck" xml:"sourceDestCheck"`
	Status             string                              `ec2:"status" xml:"status"`
	SubnetID           string                              `ec2:"subnetId" xml:"subnetId"`
	VpcID              string                              `ec2:"vpcId" xml:"vpcId"`
}

// InstanceNetworkInterfaceAssociation is undocumented.
type InstanceNetworkInterfaceAssociation struct {
	IPOwnerID     string `ec2:"ipOwnerId" xml:"ipOwnerId"`
	PublicDNSName string `ec2:"publicDnsName" xml:"publicDnsName"`
	PublicIP      string `ec2:"publicIp" xml:"publicIp"`
}

// InstanceNetworkInterfaceAttachment is undocumented.
type InstanceNetworkInterfaceAttachment struct {
	AttachTime          time.Time `ec2:"attachTime" xml:"attachTime"`
	AttachmentID        string    `ec2:"attachmentId" xml:"attachmentId"`
	DeleteOnTermination bool      `ec2:"deleteOnTermination" xml:"deleteOnTermination"`
	DeviceIndex         int       `ec2:"deviceIndex" xml:"deviceIndex"`
	Status              string    `ec2:"status" xml:"status"`
}

// InstanceNetworkInterfaceSpecification is undocumented.
type InstanceNetworkInterfaceSpecification struct {
	AssociatePublicIPAddress       bool                            `ec2:"associatePublicIpAddress" xml:"associatePublicIpAddress"`
	DeleteOnTermination            bool                            `ec2:"deleteOnTermination" xml:"deleteOnTermination"`
	Description                    string                          `ec2:"description" xml:"description"`
	DeviceIndex                    int                             `ec2:"deviceIndex" xml:"deviceIndex"`
	Groups                         []string                        `ec2:"SecurityGroupId" xml:"SecurityGroupId>SecurityGroupId"`
	NetworkInterfaceID             string                          `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	PrivateIPAddress               string                          `ec2:"privateIpAddress" xml:"privateIpAddress"`
	PrivateIPAddresses             []PrivateIPAddressSpecification `ec2:"privateIpAddressesSet" xml:"privateIpAddressesSet>item"`
	SecondaryPrivateIPAddressCount int                             `ec2:"secondaryPrivateIpAddressCount" xml:"secondaryPrivateIpAddressCount"`
	SubnetID                       string                          `ec2:"subnetId" xml:"subnetId"`
}

// InstancePrivateIPAddress is undocumented.
type InstancePrivateIPAddress struct {
	Association      InstanceNetworkInterfaceAssociation `ec2:"association" xml:"association"`
	Primary          bool                                `ec2:"primary" xml:"primary"`
	PrivateDNSName   string                              `ec2:"privateDnsName" xml:"privateDnsName"`
	PrivateIPAddress string                              `ec2:"privateIpAddress" xml:"privateIpAddress"`
}

// InstanceState is undocumented.
type InstanceState struct {
	Code int    `ec2:"code" xml:"code"`
	Name string `ec2:"name" xml:"name"`
}

// InstanceStateChange is undocumented.
type InstanceStateChange struct {
	CurrentState  InstanceState `ec2:"currentState" xml:"currentState"`
	InstanceID    string        `ec2:"instanceId" xml:"instanceId"`
	PreviousState InstanceState `ec2:"previousState" xml:"previousState"`
}

// InstanceStatus is undocumented.
type InstanceStatus struct {
	AvailabilityZone string                `ec2:"availabilityZone" xml:"availabilityZone"`
	Events           []InstanceStatusEvent `ec2:"eventsSet" xml:"eventsSet>item"`
	InstanceID       string                `ec2:"instanceId" xml:"instanceId"`
	InstanceState    InstanceState         `ec2:"instanceState" xml:"instanceState"`
	InstanceStatus   InstanceStatusSummary `ec2:"instanceStatus" xml:"instanceStatus"`
	SystemStatus     InstanceStatusSummary `ec2:"systemStatus" xml:"systemStatus"`
}

// InstanceStatusDetails is undocumented.
type InstanceStatusDetails struct {
	ImpairedSince time.Time `ec2:"impairedSince" xml:"impairedSince"`
	Name          string    `ec2:"name" xml:"name"`
	Status        string    `ec2:"status" xml:"status"`
}

// InstanceStatusEvent is undocumented.
type InstanceStatusEvent struct {
	Code        string    `ec2:"code" xml:"code"`
	Description string    `ec2:"description" xml:"description"`
	NotAfter    time.Time `ec2:"notAfter" xml:"notAfter"`
	NotBefore   time.Time `ec2:"notBefore" xml:"notBefore"`
}

// InstanceStatusSummary is undocumented.
type InstanceStatusSummary struct {
	Details []InstanceStatusDetails `ec2:"details" xml:"details>item"`
	Status  string                  `ec2:"status" xml:"status"`
}

// InternetGateway is undocumented.
type InternetGateway struct {
	Attachments       []InternetGatewayAttachment `ec2:"attachmentSet" xml:"attachmentSet>item"`
	InternetGatewayID string                      `ec2:"internetGatewayId" xml:"internetGatewayId"`
	Tags              []Tag                       `ec2:"tagSet" xml:"tagSet>item"`
}

// InternetGatewayAttachment is undocumented.
type InternetGatewayAttachment struct {
	State string `ec2:"state" xml:"state"`
	VpcID string `ec2:"vpcId" xml:"vpcId"`
}

// IPPermission is undocumented.
type IPPermission struct {
	FromPort         int               `ec2:"fromPort" xml:"fromPort"`
	IPProtocol       string            `ec2:"ipProtocol" xml:"ipProtocol"`
	IPRanges         []IPRange         `ec2:"ipRanges" xml:"ipRanges>item"`
	ToPort           int               `ec2:"toPort" xml:"toPort"`
	UserIDGroupPairs []UserIDGroupPair `ec2:"groups" xml:"groups>item"`
}

// IPRange is undocumented.
type IPRange struct {
	CidrIP string `ec2:"cidrIp" xml:"cidrIp"`
}

// KeyPair is undocumented.
type KeyPair struct {
	KeyFingerprint string `ec2:"keyFingerprint" xml:"keyFingerprint"`
	KeyMaterial    string `ec2:"keyMaterial" xml:"keyMaterial"`
	KeyName        string `ec2:"keyName" xml:"keyName"`
}

// KeyPairInfo is undocumented.
type KeyPairInfo struct {
	KeyFingerprint string `ec2:"keyFingerprint" xml:"keyFingerprint"`
	KeyName        string `ec2:"keyName" xml:"keyName"`
}

// LaunchPermission is undocumented.
type LaunchPermission struct {
	Group  string `ec2:"group" xml:"group"`
	UserID string `ec2:"userId" xml:"userId"`
}

// LaunchPermissionModifications is undocumented.
type LaunchPermissionModifications struct {
	Add    []LaunchPermission `ec2:"" xml:"Add>item"`
	Remove []LaunchPermission `ec2:"" xml:"Remove>item"`
}

// LaunchSpecification is undocumented.
type LaunchSpecification struct {
	AddressingType      string                                  `ec2:"addressingType" xml:"addressingType"`
	BlockDeviceMappings []BlockDeviceMapping                    `ec2:"blockDeviceMapping" xml:"blockDeviceMapping>item"`
	EbsOptimized        bool                                    `ec2:"ebsOptimized" xml:"ebsOptimized"`
	IamInstanceProfile  IamInstanceProfileSpecification         `ec2:"iamInstanceProfile" xml:"iamInstanceProfile"`
	ImageID             string                                  `ec2:"imageId" xml:"imageId"`
	InstanceType        string                                  `ec2:"instanceType" xml:"instanceType"`
	KernelID            string                                  `ec2:"kernelId" xml:"kernelId"`
	KeyName             string                                  `ec2:"keyName" xml:"keyName"`
	Monitoring          RunInstancesMonitoringEnabled           `ec2:"monitoring" xml:"monitoring"`
	NetworkInterfaces   []InstanceNetworkInterfaceSpecification `ec2:"networkInterfaceSet" xml:"networkInterfaceSet>item"`
	Placement           SpotPlacement                           `ec2:"placement" xml:"placement"`
	RamdiskID           string                                  `ec2:"ramdiskId" xml:"ramdiskId"`
	SecurityGroups      []GroupIdentifier                       `ec2:"groupSet" xml:"groupSet>item"`
	SubnetID            string                                  `ec2:"subnetId" xml:"subnetId"`
	UserData            string                                  `ec2:"userData" xml:"userData"`
}

// ModifyImageAttributeRequest is undocumented.
type ModifyImageAttributeRequest struct {
	Attribute        string                        `ec2:"" xml:"Attribute"`
	Description      AttributeValue                `ec2:"" xml:"Description"`
	DryRun           bool                          `ec2:"dryRun" xml:"dryRun"`
	ImageID          string                        `ec2:"" xml:"ImageId"`
	LaunchPermission LaunchPermissionModifications `ec2:"" xml:"LaunchPermission"`
	OperationType    string                        `ec2:"" xml:"OperationType"`
	ProductCodes     []string                      `ec2:"ProductCode" xml:"ProductCode>ProductCode"`
	UserGroups       []string                      `ec2:"UserGroup" xml:"UserGroup>UserGroup"`
	UserIds          []string                      `ec2:"UserId" xml:"UserId>UserId"`
	Value            string                        `ec2:"" xml:"Value"`
}

// ModifyInstanceAttributeRequest is undocumented.
type ModifyInstanceAttributeRequest struct {
	Attribute                         string                                    `ec2:"attribute" xml:"attribute"`
	BlockDeviceMappings               []InstanceBlockDeviceMappingSpecification `ec2:"blockDeviceMapping" xml:"blockDeviceMapping>item"`
	DisableAPITermination             AttributeBooleanValue                     `ec2:"disableApiTermination" xml:"disableApiTermination"`
	DryRun                            bool                                      `ec2:"dryRun" xml:"dryRun"`
	EbsOptimized                      AttributeBooleanValue                     `ec2:"ebsOptimized" xml:"ebsOptimized"`
	Groups                            []string                                  `ec2:"GroupId" xml:"GroupId>groupId"`
	InstanceID                        string                                    `ec2:"instanceId" xml:"instanceId"`
	InstanceInitiatedShutdownBehavior AttributeValue                            `ec2:"instanceInitiatedShutdownBehavior" xml:"instanceInitiatedShutdownBehavior"`
	InstanceType                      AttributeValue                            `ec2:"instanceType" xml:"instanceType"`
	Kernel                            AttributeValue                            `ec2:"kernel" xml:"kernel"`
	Ramdisk                           AttributeValue                            `ec2:"ramdisk" xml:"ramdisk"`
	SourceDestCheck                   AttributeBooleanValue                     `ec2:"" xml:"SourceDestCheck"`
	SriovNetSupport                   AttributeValue                            `ec2:"sriovNetSupport" xml:"sriovNetSupport"`
	UserData                          BlobAttributeValue                        `ec2:"userData" xml:"userData"`
	Value                             string                                    `ec2:"value" xml:"value"`
}

// ModifyNetworkInterfaceAttributeRequest is undocumented.
type ModifyNetworkInterfaceAttributeRequest struct {
	Attachment         NetworkInterfaceAttachmentChanges `ec2:"attachment" xml:"attachment"`
	Description        AttributeValue                    `ec2:"description" xml:"description"`
	DryRun             bool                              `ec2:"dryRun" xml:"dryRun"`
	Groups             []string                          `ec2:"SecurityGroupId" xml:"SecurityGroupId>SecurityGroupId"`
	NetworkInterfaceID string                            `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	SourceDestCheck    AttributeBooleanValue             `ec2:"sourceDestCheck" xml:"sourceDestCheck"`
}

// ModifyReservedInstancesRequest is undocumented.
type ModifyReservedInstancesRequest struct {
	ClientToken          string                           `ec2:"clientToken" xml:"clientToken"`
	ReservedInstancesIds []string                         `ec2:"ReservedInstancesId" xml:"ReservedInstancesId>ReservedInstancesId"`
	TargetConfigurations []ReservedInstancesConfiguration `ec2:"ReservedInstancesConfigurationSetItemType" xml:"ReservedInstancesConfigurationSetItemType>item"`
}

// ModifyReservedInstancesResult is undocumented.
type ModifyReservedInstancesResult struct {
	ReservedInstancesModificationID string `ec2:"reservedInstancesModificationId" xml:"reservedInstancesModificationId"`
}

// ModifySnapshotAttributeRequest is undocumented.
type ModifySnapshotAttributeRequest struct {
	Attribute              string                              `ec2:"" xml:"Attribute"`
	CreateVolumePermission CreateVolumePermissionModifications `ec2:"" xml:"CreateVolumePermission"`
	DryRun                 bool                                `ec2:"dryRun" xml:"dryRun"`
	GroupNames             []string                            `ec2:"UserGroup" xml:"UserGroup>GroupName"`
	OperationType          string                              `ec2:"" xml:"OperationType"`
	SnapshotID             string                              `ec2:"" xml:"SnapshotId"`
	UserIds                []string                            `ec2:"UserId" xml:"UserId>UserId"`
}

// ModifySubnetAttributeRequest is undocumented.
type ModifySubnetAttributeRequest struct {
	MapPublicIPOnLaunch AttributeBooleanValue `ec2:"" xml:"MapPublicIpOnLaunch"`
	SubnetID            string                `ec2:"subnetId" xml:"subnetId"`
}

// ModifyVolumeAttributeRequest is undocumented.
type ModifyVolumeAttributeRequest struct {
	AutoEnableIO AttributeBooleanValue `ec2:"" xml:"AutoEnableIO"`
	DryRun       bool                  `ec2:"dryRun" xml:"dryRun"`
	VolumeID     string                `ec2:"" xml:"VolumeId"`
}

// ModifyVpcAttributeRequest is undocumented.
type ModifyVpcAttributeRequest struct {
	EnableDNSHostnames AttributeBooleanValue `ec2:"" xml:"EnableDnsHostnames"`
	EnableDNSSupport   AttributeBooleanValue `ec2:"" xml:"EnableDnsSupport"`
	VpcID              string                `ec2:"vpcId" xml:"vpcId"`
}

// MonitorInstancesRequest is undocumented.
type MonitorInstancesRequest struct {
	DryRun      bool     `ec2:"dryRun" xml:"dryRun"`
	InstanceIds []string `ec2:"InstanceId" xml:"InstanceId>InstanceId"`
}

// MonitorInstancesResult is undocumented.
type MonitorInstancesResult struct {
	InstanceMonitorings []InstanceMonitoring `ec2:"instancesSet" xml:"instancesSet>item"`
}

// Monitoring is undocumented.
type Monitoring struct {
	State string `ec2:"state" xml:"state"`
}

// NetworkAcl is undocumented.
type NetworkAcl struct {
	Associations []NetworkAclAssociation `ec2:"associationSet" xml:"associationSet>item"`
	Entries      []NetworkAclEntry       `ec2:"entrySet" xml:"entrySet>item"`
	IsDefault    bool                    `ec2:"default" xml:"default"`
	NetworkAclID string                  `ec2:"networkAclId" xml:"networkAclId"`
	Tags         []Tag                   `ec2:"tagSet" xml:"tagSet>item"`
	VpcID        string                  `ec2:"vpcId" xml:"vpcId"`
}

// NetworkAclAssociation is undocumented.
type NetworkAclAssociation struct {
	NetworkAclAssociationID string `ec2:"networkAclAssociationId" xml:"networkAclAssociationId"`
	NetworkAclID            string `ec2:"networkAclId" xml:"networkAclId"`
	SubnetID                string `ec2:"subnetId" xml:"subnetId"`
}

// NetworkAclEntry is undocumented.
type NetworkAclEntry struct {
	CidrBlock    string       `ec2:"cidrBlock" xml:"cidrBlock"`
	Egress       bool         `ec2:"egress" xml:"egress"`
	IcmpTypeCode IcmpTypeCode `ec2:"icmpTypeCode" xml:"icmpTypeCode"`
	PortRange    PortRange    `ec2:"portRange" xml:"portRange"`
	Protocol     string       `ec2:"protocol" xml:"protocol"`
	RuleAction   string       `ec2:"ruleAction" xml:"ruleAction"`
	RuleNumber   int          `ec2:"ruleNumber" xml:"ruleNumber"`
}

// NetworkInterface is undocumented.
type NetworkInterface struct {
	Association        NetworkInterfaceAssociation        `ec2:"association" xml:"association"`
	Attachment         NetworkInterfaceAttachment         `ec2:"attachment" xml:"attachment"`
	AvailabilityZone   string                             `ec2:"availabilityZone" xml:"availabilityZone"`
	Description        string                             `ec2:"description" xml:"description"`
	Groups             []GroupIdentifier                  `ec2:"groupSet" xml:"groupSet>item"`
	MacAddress         string                             `ec2:"macAddress" xml:"macAddress"`
	NetworkInterfaceID string                             `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	OwnerID            string                             `ec2:"ownerId" xml:"ownerId"`
	PrivateDNSName     string                             `ec2:"privateDnsName" xml:"privateDnsName"`
	PrivateIPAddress   string                             `ec2:"privateIpAddress" xml:"privateIpAddress"`
	PrivateIPAddresses []NetworkInterfacePrivateIPAddress `ec2:"privateIpAddressesSet" xml:"privateIpAddressesSet>item"`
	RequesterID        string                             `ec2:"requesterId" xml:"requesterId"`
	RequesterManaged   bool                               `ec2:"requesterManaged" xml:"requesterManaged"`
	SourceDestCheck    bool                               `ec2:"sourceDestCheck" xml:"sourceDestCheck"`
	Status             string                             `ec2:"status" xml:"status"`
	SubnetID           string                             `ec2:"subnetId" xml:"subnetId"`
	TagSet             []Tag                              `ec2:"tagSet" xml:"tagSet>item"`
	VpcID              string                             `ec2:"vpcId" xml:"vpcId"`
}

// NetworkInterfaceAssociation is undocumented.
type NetworkInterfaceAssociation struct {
	AllocationID  string `ec2:"allocationId" xml:"allocationId"`
	AssociationID string `ec2:"associationId" xml:"associationId"`
	IPOwnerID     string `ec2:"ipOwnerId" xml:"ipOwnerId"`
	PublicDNSName string `ec2:"publicDnsName" xml:"publicDnsName"`
	PublicIP      string `ec2:"publicIp" xml:"publicIp"`
}

// NetworkInterfaceAttachment is undocumented.
type NetworkInterfaceAttachment struct {
	AttachTime          time.Time `ec2:"attachTime" xml:"attachTime"`
	AttachmentID        string    `ec2:"attachmentId" xml:"attachmentId"`
	DeleteOnTermination bool      `ec2:"deleteOnTermination" xml:"deleteOnTermination"`
	DeviceIndex         int       `ec2:"deviceIndex" xml:"deviceIndex"`
	InstanceID          string    `ec2:"instanceId" xml:"instanceId"`
	InstanceOwnerID     string    `ec2:"instanceOwnerId" xml:"instanceOwnerId"`
	Status              string    `ec2:"status" xml:"status"`
}

// NetworkInterfaceAttachmentChanges is undocumented.
type NetworkInterfaceAttachmentChanges struct {
	AttachmentID        string `ec2:"attachmentId" xml:"attachmentId"`
	DeleteOnTermination bool   `ec2:"deleteOnTermination" xml:"deleteOnTermination"`
}

// NetworkInterfacePrivateIPAddress is undocumented.
type NetworkInterfacePrivateIPAddress struct {
	Association      NetworkInterfaceAssociation `ec2:"association" xml:"association"`
	Primary          bool                        `ec2:"primary" xml:"primary"`
	PrivateDNSName   string                      `ec2:"privateDnsName" xml:"privateDnsName"`
	PrivateIPAddress string                      `ec2:"privateIpAddress" xml:"privateIpAddress"`
}

// NewDhcpConfiguration is undocumented.
type NewDhcpConfiguration struct {
	Key    string   `ec2:"key" xml:"key"`
	Values []string `ec2:"Value" xml:"Value>item"`
}

// Placement is undocumented.
type Placement struct {
	AvailabilityZone string `ec2:"availabilityZone" xml:"availabilityZone"`
	GroupName        string `ec2:"groupName" xml:"groupName"`
	Tenancy          string `ec2:"tenancy" xml:"tenancy"`
}

// PlacementGroup is undocumented.
type PlacementGroup struct {
	GroupName string `ec2:"groupName" xml:"groupName"`
	State     string `ec2:"state" xml:"state"`
	Strategy  string `ec2:"strategy" xml:"strategy"`
}

// PortRange is undocumented.
type PortRange struct {
	From int `ec2:"from" xml:"from"`
	To   int `ec2:"to" xml:"to"`
}

// PriceSchedule is undocumented.
type PriceSchedule struct {
	Active       bool    `ec2:"active" xml:"active"`
	CurrencyCode string  `ec2:"currencyCode" xml:"currencyCode"`
	Price        float64 `ec2:"price" xml:"price"`
	Term         int64   `ec2:"term" xml:"term"`
}

// PriceScheduleSpecification is undocumented.
type PriceScheduleSpecification struct {
	CurrencyCode string  `ec2:"currencyCode" xml:"currencyCode"`
	Price        float64 `ec2:"price" xml:"price"`
	Term         int64   `ec2:"term" xml:"term"`
}

// PricingDetail is undocumented.
type PricingDetail struct {
	Count int     `ec2:"count" xml:"count"`
	Price float64 `ec2:"price" xml:"price"`
}

// PrivateIPAddressSpecification is undocumented.
type PrivateIPAddressSpecification struct {
	Primary          bool   `ec2:"primary" xml:"primary"`
	PrivateIPAddress string `ec2:"privateIpAddress" xml:"privateIpAddress"`
}

// ProductCode is undocumented.
type ProductCode struct {
	ProductCodeID   string `ec2:"productCode" xml:"productCode"`
	ProductCodeType string `ec2:"type" xml:"type"`
}

// PropagatingVgw is undocumented.
type PropagatingVgw struct {
	GatewayID string `ec2:"gatewayId" xml:"gatewayId"`
}

// PurchaseReservedInstancesOfferingRequest is undocumented.
type PurchaseReservedInstancesOfferingRequest struct {
	DryRun                      bool                       `ec2:"dryRun" xml:"dryRun"`
	InstanceCount               int                        `ec2:"" xml:"InstanceCount"`
	LimitPrice                  ReservedInstanceLimitPrice `ec2:"limitPrice" xml:"limitPrice"`
	ReservedInstancesOfferingID string                     `ec2:"" xml:"ReservedInstancesOfferingId"`
}

// PurchaseReservedInstancesOfferingResult is undocumented.
type PurchaseReservedInstancesOfferingResult struct {
	ReservedInstancesID string `ec2:"reservedInstancesId" xml:"reservedInstancesId"`
}

// RebootInstancesRequest is undocumented.
type RebootInstancesRequest struct {
	DryRun      bool     `ec2:"dryRun" xml:"dryRun"`
	InstanceIds []string `ec2:"InstanceId" xml:"InstanceId>InstanceId"`
}

// RecurringCharge is undocumented.
type RecurringCharge struct {
	Amount    float64 `ec2:"amount" xml:"amount"`
	Frequency string  `ec2:"frequency" xml:"frequency"`
}

// Region is undocumented.
type Region struct {
	Endpoint   string `ec2:"regionEndpoint" xml:"regionEndpoint"`
	RegionName string `ec2:"regionName" xml:"regionName"`
}

// RegisterImageRequest is undocumented.
type RegisterImageRequest struct {
	Architecture        string               `ec2:"architecture" xml:"architecture"`
	BlockDeviceMappings []BlockDeviceMapping `ec2:"BlockDeviceMapping" xml:"BlockDeviceMapping>BlockDeviceMapping"`
	Description         string               `ec2:"description" xml:"description"`
	DryRun              bool                 `ec2:"dryRun" xml:"dryRun"`
	ImageLocation       string               `ec2:"" xml:"ImageLocation"`
	KernelID            string               `ec2:"kernelId" xml:"kernelId"`
	Name                string               `ec2:"name" xml:"name"`
	RamdiskID           string               `ec2:"ramdiskId" xml:"ramdiskId"`
	RootDeviceName      string               `ec2:"rootDeviceName" xml:"rootDeviceName"`
	SriovNetSupport     string               `ec2:"sriovNetSupport" xml:"sriovNetSupport"`
	VirtualizationType  string               `ec2:"virtualizationType" xml:"virtualizationType"`
}

// RegisterImageResult is undocumented.
type RegisterImageResult struct {
	ImageID string `ec2:"imageId" xml:"imageId"`
}

// RejectVpcPeeringConnectionRequest is undocumented.
type RejectVpcPeeringConnectionRequest struct {
	DryRun                 bool   `ec2:"dryRun" xml:"dryRun"`
	VpcPeeringConnectionID string `ec2:"vpcPeeringConnectionId" xml:"vpcPeeringConnectionId"`
}

// RejectVpcPeeringConnectionResult is undocumented.
type RejectVpcPeeringConnectionResult struct {
	Return bool `ec2:"return" xml:"return"`
}

// ReleaseAddressRequest is undocumented.
type ReleaseAddressRequest struct {
	AllocationID string `ec2:"" xml:"AllocationId"`
	DryRun       bool   `ec2:"dryRun" xml:"dryRun"`
	PublicIP     string `ec2:"" xml:"PublicIp"`
}

// ReplaceNetworkAclAssociationRequest is undocumented.
type ReplaceNetworkAclAssociationRequest struct {
	AssociationID string `ec2:"associationId" xml:"associationId"`
	DryRun        bool   `ec2:"dryRun" xml:"dryRun"`
	NetworkAclID  string `ec2:"networkAclId" xml:"networkAclId"`
}

// ReplaceNetworkAclAssociationResult is undocumented.
type ReplaceNetworkAclAssociationResult struct {
	NewAssociationID string `ec2:"newAssociationId" xml:"newAssociationId"`
}

// ReplaceNetworkAclEntryRequest is undocumented.
type ReplaceNetworkAclEntryRequest struct {
	CidrBlock    string       `ec2:"cidrBlock" xml:"cidrBlock"`
	DryRun       bool         `ec2:"dryRun" xml:"dryRun"`
	Egress       bool         `ec2:"egress" xml:"egress"`
	IcmpTypeCode IcmpTypeCode `ec2:"Icmp" xml:"Icmp"`
	NetworkAclID string       `ec2:"networkAclId" xml:"networkAclId"`
	PortRange    PortRange    `ec2:"portRange" xml:"portRange"`
	Protocol     string       `ec2:"protocol" xml:"protocol"`
	RuleAction   string       `ec2:"ruleAction" xml:"ruleAction"`
	RuleNumber   int          `ec2:"ruleNumber" xml:"ruleNumber"`
}

// ReplaceRouteRequest is undocumented.
type ReplaceRouteRequest struct {
	DestinationCidrBlock   string `ec2:"destinationCidrBlock" xml:"destinationCidrBlock"`
	DryRun                 bool   `ec2:"dryRun" xml:"dryRun"`
	GatewayID              string `ec2:"gatewayId" xml:"gatewayId"`
	InstanceID             string `ec2:"instanceId" xml:"instanceId"`
	NetworkInterfaceID     string `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	RouteTableID           string `ec2:"routeTableId" xml:"routeTableId"`
	VpcPeeringConnectionID string `ec2:"vpcPeeringConnectionId" xml:"vpcPeeringConnectionId"`
}

// ReplaceRouteTableAssociationRequest is undocumented.
type ReplaceRouteTableAssociationRequest struct {
	AssociationID string `ec2:"associationId" xml:"associationId"`
	DryRun        bool   `ec2:"dryRun" xml:"dryRun"`
	RouteTableID  string `ec2:"routeTableId" xml:"routeTableId"`
}

// ReplaceRouteTableAssociationResult is undocumented.
type ReplaceRouteTableAssociationResult struct {
	NewAssociationID string `ec2:"newAssociationId" xml:"newAssociationId"`
}

// ReportInstanceStatusRequest is undocumented.
type ReportInstanceStatusRequest struct {
	Description string    `ec2:"description" xml:"description"`
	DryRun      bool      `ec2:"dryRun" xml:"dryRun"`
	EndTime     time.Time `ec2:"endTime" xml:"endTime"`
	Instances   []string  `ec2:"instanceId" xml:"instanceId>InstanceId"`
	ReasonCodes []string  `ec2:"reasonCode" xml:"reasonCode>item"`
	StartTime   time.Time `ec2:"startTime" xml:"startTime"`
	Status      string    `ec2:"status" xml:"status"`
}

// RequestSpotInstancesRequest is undocumented.
type RequestSpotInstancesRequest struct {
	AvailabilityZoneGroup string                         `ec2:"availabilityZoneGroup" xml:"availabilityZoneGroup"`
	DryRun                bool                           `ec2:"dryRun" xml:"dryRun"`
	InstanceCount         int                            `ec2:"instanceCount" xml:"instanceCount"`
	LaunchGroup           string                         `ec2:"launchGroup" xml:"launchGroup"`
	LaunchSpecification   RequestSpotLaunchSpecification `ec2:"" xml:"LaunchSpecification"`
	SpotPrice             string                         `ec2:"spotPrice" xml:"spotPrice"`
	Type                  string                         `ec2:"type" xml:"type"`
	ValidFrom             time.Time                      `ec2:"validFrom" xml:"validFrom"`
	ValidUntil            time.Time                      `ec2:"validUntil" xml:"validUntil"`
}

// RequestSpotInstancesResult is undocumented.
type RequestSpotInstancesResult struct {
	SpotInstanceRequests []SpotInstanceRequest `ec2:"spotInstanceRequestSet" xml:"spotInstanceRequestSet>item"`
}

// RequestSpotLaunchSpecification is undocumented.
type RequestSpotLaunchSpecification struct {
	AddressingType      string                                  `ec2:"addressingType" xml:"addressingType"`
	BlockDeviceMappings []BlockDeviceMapping                    `ec2:"blockDeviceMapping" xml:"blockDeviceMapping>item"`
	EbsOptimized        bool                                    `ec2:"ebsOptimized" xml:"ebsOptimized"`
	IamInstanceProfile  IamInstanceProfileSpecification         `ec2:"iamInstanceProfile" xml:"iamInstanceProfile"`
	ImageID             string                                  `ec2:"imageId" xml:"imageId"`
	InstanceType        string                                  `ec2:"instanceType" xml:"instanceType"`
	KernelID            string                                  `ec2:"kernelId" xml:"kernelId"`
	KeyName             string                                  `ec2:"keyName" xml:"keyName"`
	Monitoring          RunInstancesMonitoringEnabled           `ec2:"monitoring" xml:"monitoring"`
	NetworkInterfaces   []InstanceNetworkInterfaceSpecification `ec2:"NetworkInterface" xml:"NetworkInterface>item"`
	Placement           SpotPlacement                           `ec2:"placement" xml:"placement"`
	RamdiskID           string                                  `ec2:"ramdiskId" xml:"ramdiskId"`
	SecurityGroupIds    []string                                `ec2:"SecurityGroupId" xml:"SecurityGroupId>item"`
	SecurityGroups      []string                                `ec2:"SecurityGroup" xml:"SecurityGroup>item"`
	SubnetID            string                                  `ec2:"subnetId" xml:"subnetId"`
	UserData            string                                  `ec2:"userData" xml:"userData"`
}

// Reservation is undocumented.
type Reservation struct {
	Groups        []GroupIdentifier `ec2:"groupSet" xml:"groupSet>item"`
	Instances     []Instance        `ec2:"instancesSet" xml:"instancesSet>item"`
	OwnerID       string            `ec2:"ownerId" xml:"ownerId"`
	RequesterID   string            `ec2:"requesterId" xml:"requesterId"`
	ReservationID string            `ec2:"reservationId" xml:"reservationId"`
}

// ReservedInstanceLimitPrice is undocumented.
type ReservedInstanceLimitPrice struct {
	Amount       float64 `ec2:"amount" xml:"amount"`
	CurrencyCode string  `ec2:"currencyCode" xml:"currencyCode"`
}

// ReservedInstances is undocumented.
type ReservedInstances struct {
	AvailabilityZone    string            `ec2:"availabilityZone" xml:"availabilityZone"`
	CurrencyCode        string            `ec2:"currencyCode" xml:"currencyCode"`
	Duration            int64             `ec2:"duration" xml:"duration"`
	End                 time.Time         `ec2:"end" xml:"end"`
	FixedPrice          float32           `ec2:"fixedPrice" xml:"fixedPrice"`
	InstanceCount       int               `ec2:"instanceCount" xml:"instanceCount"`
	InstanceTenancy     string            `ec2:"instanceTenancy" xml:"instanceTenancy"`
	InstanceType        string            `ec2:"instanceType" xml:"instanceType"`
	OfferingType        string            `ec2:"offeringType" xml:"offeringType"`
	ProductDescription  string            `ec2:"productDescription" xml:"productDescription"`
	RecurringCharges    []RecurringCharge `ec2:"recurringCharges" xml:"recurringCharges>item"`
	ReservedInstancesID string            `ec2:"reservedInstancesId" xml:"reservedInstancesId"`
	Start               time.Time         `ec2:"start" xml:"start"`
	State               string            `ec2:"state" xml:"state"`
	Tags                []Tag             `ec2:"tagSet" xml:"tagSet>item"`
	UsagePrice          float32           `ec2:"usagePrice" xml:"usagePrice"`
}

// ReservedInstancesConfiguration is undocumented.
type ReservedInstancesConfiguration struct {
	AvailabilityZone string `ec2:"availabilityZone" xml:"availabilityZone"`
	InstanceCount    int    `ec2:"instanceCount" xml:"instanceCount"`
	InstanceType     string `ec2:"instanceType" xml:"instanceType"`
	Platform         string `ec2:"platform" xml:"platform"`
}

// ReservedInstancesID is undocumented.
type ReservedInstancesID struct {
	ReservedInstancesID string `ec2:"reservedInstancesId" xml:"reservedInstancesId"`
}

// ReservedInstancesListing is undocumented.
type ReservedInstancesListing struct {
	ClientToken                string          `ec2:"clientToken" xml:"clientToken"`
	CreateDate                 time.Time       `ec2:"createDate" xml:"createDate"`
	InstanceCounts             []InstanceCount `ec2:"instanceCounts" xml:"instanceCounts>item"`
	PriceSchedules             []PriceSchedule `ec2:"priceSchedules" xml:"priceSchedules>item"`
	ReservedInstancesID        string          `ec2:"reservedInstancesId" xml:"reservedInstancesId"`
	ReservedInstancesListingID string          `ec2:"reservedInstancesListingId" xml:"reservedInstancesListingId"`
	Status                     string          `ec2:"status" xml:"status"`
	StatusMessage              string          `ec2:"statusMessage" xml:"statusMessage"`
	Tags                       []Tag           `ec2:"tagSet" xml:"tagSet>item"`
	UpdateDate                 time.Time       `ec2:"updateDate" xml:"updateDate"`
}

// ReservedInstancesModification is undocumented.
type ReservedInstancesModification struct {
	ClientToken                     string                                `ec2:"clientToken" xml:"clientToken"`
	CreateDate                      time.Time                             `ec2:"createDate" xml:"createDate"`
	EffectiveDate                   time.Time                             `ec2:"effectiveDate" xml:"effectiveDate"`
	ModificationResults             []ReservedInstancesModificationResult `ec2:"modificationResultSet" xml:"modificationResultSet>item"`
	ReservedInstancesIds            []ReservedInstancesID                 `ec2:"reservedInstancesSet" xml:"reservedInstancesSet>item"`
	ReservedInstancesModificationID string                                `ec2:"reservedInstancesModificationId" xml:"reservedInstancesModificationId"`
	Status                          string                                `ec2:"status" xml:"status"`
	StatusMessage                   string                                `ec2:"statusMessage" xml:"statusMessage"`
	UpdateDate                      time.Time                             `ec2:"updateDate" xml:"updateDate"`
}

// ReservedInstancesModificationResult is undocumented.
type ReservedInstancesModificationResult struct {
	ReservedInstancesID string                         `ec2:"reservedInstancesId" xml:"reservedInstancesId"`
	TargetConfiguration ReservedInstancesConfiguration `ec2:"targetConfiguration" xml:"targetConfiguration"`
}

// ReservedInstancesOffering is undocumented.
type ReservedInstancesOffering struct {
	AvailabilityZone            string            `ec2:"availabilityZone" xml:"availabilityZone"`
	CurrencyCode                string            `ec2:"currencyCode" xml:"currencyCode"`
	Duration                    int64             `ec2:"duration" xml:"duration"`
	FixedPrice                  float32           `ec2:"fixedPrice" xml:"fixedPrice"`
	InstanceTenancy             string            `ec2:"instanceTenancy" xml:"instanceTenancy"`
	InstanceType                string            `ec2:"instanceType" xml:"instanceType"`
	Marketplace                 bool              `ec2:"marketplace" xml:"marketplace"`
	OfferingType                string            `ec2:"offeringType" xml:"offeringType"`
	PricingDetails              []PricingDetail   `ec2:"pricingDetailsSet" xml:"pricingDetailsSet>item"`
	ProductDescription          string            `ec2:"productDescription" xml:"productDescription"`
	RecurringCharges            []RecurringCharge `ec2:"recurringCharges" xml:"recurringCharges>item"`
	ReservedInstancesOfferingID string            `ec2:"reservedInstancesOfferingId" xml:"reservedInstancesOfferingId"`
	UsagePrice                  float32           `ec2:"usagePrice" xml:"usagePrice"`
}

// ResetImageAttributeRequest is undocumented.
type ResetImageAttributeRequest struct {
	Attribute string `ec2:"" xml:"Attribute"`
	DryRun    bool   `ec2:"dryRun" xml:"dryRun"`
	ImageID   string `ec2:"" xml:"ImageId"`
}

// ResetInstanceAttributeRequest is undocumented.
type ResetInstanceAttributeRequest struct {
	Attribute  string `ec2:"attribute" xml:"attribute"`
	DryRun     bool   `ec2:"dryRun" xml:"dryRun"`
	InstanceID string `ec2:"instanceId" xml:"instanceId"`
}

// ResetNetworkInterfaceAttributeRequest is undocumented.
type ResetNetworkInterfaceAttributeRequest struct {
	DryRun             bool   `ec2:"dryRun" xml:"dryRun"`
	NetworkInterfaceID string `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	SourceDestCheck    string `ec2:"sourceDestCheck" xml:"sourceDestCheck"`
}

// ResetSnapshotAttributeRequest is undocumented.
type ResetSnapshotAttributeRequest struct {
	Attribute  string `ec2:"" xml:"Attribute"`
	DryRun     bool   `ec2:"dryRun" xml:"dryRun"`
	SnapshotID string `ec2:"" xml:"SnapshotId"`
}

// RevokeSecurityGroupEgressRequest is undocumented.
type RevokeSecurityGroupEgressRequest struct {
	CidrIP                     string         `ec2:"cidrIp" xml:"cidrIp"`
	DryRun                     bool           `ec2:"dryRun" xml:"dryRun"`
	FromPort                   int            `ec2:"fromPort" xml:"fromPort"`
	GroupID                    string         `ec2:"groupId" xml:"groupId"`
	IPPermissions              []IPPermission `ec2:"ipPermissions" xml:"ipPermissions>item"`
	IPProtocol                 string         `ec2:"ipProtocol" xml:"ipProtocol"`
	SourceSecurityGroupName    string         `ec2:"sourceSecurityGroupName" xml:"sourceSecurityGroupName"`
	SourceSecurityGroupOwnerID string         `ec2:"sourceSecurityGroupOwnerId" xml:"sourceSecurityGroupOwnerId"`
	ToPort                     int            `ec2:"toPort" xml:"toPort"`
}

// RevokeSecurityGroupIngressRequest is undocumented.
type RevokeSecurityGroupIngressRequest struct {
	CidrIP                     string         `ec2:"" xml:"CidrIp"`
	DryRun                     bool           `ec2:"dryRun" xml:"dryRun"`
	FromPort                   int            `ec2:"" xml:"FromPort"`
	GroupID                    string         `ec2:"" xml:"GroupId"`
	GroupName                  string         `ec2:"" xml:"GroupName"`
	IPPermissions              []IPPermission `ec2:"" xml:"IpPermissions>item"`
	IPProtocol                 string         `ec2:"" xml:"IpProtocol"`
	SourceSecurityGroupName    string         `ec2:"" xml:"SourceSecurityGroupName"`
	SourceSecurityGroupOwnerID string         `ec2:"" xml:"SourceSecurityGroupOwnerId"`
	ToPort                     int            `ec2:"" xml:"ToPort"`
}

// Route is undocumented.
type Route struct {
	DestinationCidrBlock   string `ec2:"destinationCidrBlock" xml:"destinationCidrBlock"`
	GatewayID              string `ec2:"gatewayId" xml:"gatewayId"`
	InstanceID             string `ec2:"instanceId" xml:"instanceId"`
	InstanceOwnerID        string `ec2:"instanceOwnerId" xml:"instanceOwnerId"`
	NetworkInterfaceID     string `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	Origin                 string `ec2:"origin" xml:"origin"`
	State                  string `ec2:"state" xml:"state"`
	VpcPeeringConnectionID string `ec2:"vpcPeeringConnectionId" xml:"vpcPeeringConnectionId"`
}

// RouteTable is undocumented.
type RouteTable struct {
	Associations    []RouteTableAssociation `ec2:"associationSet" xml:"associationSet>item"`
	PropagatingVgws []PropagatingVgw        `ec2:"propagatingVgwSet" xml:"propagatingVgwSet>item"`
	RouteTableID    string                  `ec2:"routeTableId" xml:"routeTableId"`
	Routes          []Route                 `ec2:"routeSet" xml:"routeSet>item"`
	Tags            []Tag                   `ec2:"tagSet" xml:"tagSet>item"`
	VpcID           string                  `ec2:"vpcId" xml:"vpcId"`
}

// RouteTableAssociation is undocumented.
type RouteTableAssociation struct {
	Main                    bool   `ec2:"main" xml:"main"`
	RouteTableAssociationID string `ec2:"routeTableAssociationId" xml:"routeTableAssociationId"`
	RouteTableID            string `ec2:"routeTableId" xml:"routeTableId"`
	SubnetID                string `ec2:"subnetId" xml:"subnetId"`
}

// RunInstancesMonitoringEnabled is undocumented.
type RunInstancesMonitoringEnabled struct {
	Enabled bool `ec2:"enabled" xml:"enabled"`
}

// RunInstancesRequest is undocumented.
type RunInstancesRequest struct {
	AdditionalInfo                    string                                  `ec2:"additionalInfo" xml:"additionalInfo"`
	BlockDeviceMappings               []BlockDeviceMapping                    `ec2:"BlockDeviceMapping" xml:"BlockDeviceMapping>BlockDeviceMapping"`
	ClientToken                       string                                  `ec2:"clientToken" xml:"clientToken"`
	DisableAPITermination             bool                                    `ec2:"disableApiTermination" xml:"disableApiTermination"`
	DryRun                            bool                                    `ec2:"dryRun" xml:"dryRun"`
	EbsOptimized                      bool                                    `ec2:"ebsOptimized" xml:"ebsOptimized"`
	IamInstanceProfile                IamInstanceProfileSpecification         `ec2:"iamInstanceProfile" xml:"iamInstanceProfile"`
	ImageID                           string                                  `ec2:"" xml:"ImageId"`
	InstanceInitiatedShutdownBehavior string                                  `ec2:"instanceInitiatedShutdownBehavior" xml:"instanceInitiatedShutdownBehavior"`
	InstanceType                      string                                  `ec2:"" xml:"InstanceType"`
	KernelID                          string                                  `ec2:"" xml:"KernelId"`
	KeyName                           string                                  `ec2:"" xml:"KeyName"`
	MaxCount                          int                                     `ec2:"" xml:"MaxCount"`
	MinCount                          int                                     `ec2:"" xml:"MinCount"`
	Monitoring                        RunInstancesMonitoringEnabled           `ec2:"" xml:"Monitoring"`
	NetworkInterfaces                 []InstanceNetworkInterfaceSpecification `ec2:"networkInterface" xml:"networkInterface>item"`
	Placement                         Placement                               `ec2:"" xml:"Placement"`
	PrivateIPAddress                  string                                  `ec2:"privateIpAddress" xml:"privateIpAddress"`
	RamdiskID                         string                                  `ec2:"" xml:"RamdiskId"`
	SecurityGroupIds                  []string                                `ec2:"SecurityGroupId" xml:"SecurityGroupId>SecurityGroupId"`
	SecurityGroups                    []string                                `ec2:"SecurityGroup" xml:"SecurityGroup>SecurityGroup"`
	SubnetID                          string                                  `ec2:"" xml:"SubnetId"`
	UserData                          string                                  `ec2:"" xml:"UserData"`
}

// S3Storage is undocumented.
type S3Storage struct {
	AWSAccessKeyID        string `ec2:"" xml:"AWSAccessKeyId"`
	Bucket                string `ec2:"bucket" xml:"bucket"`
	Prefix                string `ec2:"prefix" xml:"prefix"`
	UploadPolicy          []byte `ec2:"uploadPolicy" xml:"uploadPolicy"`
	UploadPolicySignature string `ec2:"uploadPolicySignature" xml:"uploadPolicySignature"`
}

// SecurityGroup is undocumented.
type SecurityGroup struct {
	Description         string         `ec2:"groupDescription" xml:"groupDescription"`
	GroupID             string         `ec2:"groupId" xml:"groupId"`
	GroupName           string         `ec2:"groupName" xml:"groupName"`
	IPPermissions       []IPPermission `ec2:"ipPermissions" xml:"ipPermissions>item"`
	IPPermissionsEgress []IPPermission `ec2:"ipPermissionsEgress" xml:"ipPermissionsEgress>item"`
	OwnerID             string         `ec2:"ownerId" xml:"ownerId"`
	Tags                []Tag          `ec2:"tagSet" xml:"tagSet>item"`
	VpcID               string         `ec2:"vpcId" xml:"vpcId"`
}

// Snapshot is undocumented.
type Snapshot struct {
	Description string    `ec2:"description" xml:"description"`
	Encrypted   bool      `ec2:"encrypted" xml:"encrypted"`
	KmsKeyID    string    `ec2:"kmsKeyId" xml:"kmsKeyId"`
	OwnerAlias  string    `ec2:"ownerAlias" xml:"ownerAlias"`
	OwnerID     string    `ec2:"ownerId" xml:"ownerId"`
	Progress    string    `ec2:"progress" xml:"progress"`
	SnapshotID  string    `ec2:"snapshotId" xml:"snapshotId"`
	StartTime   time.Time `ec2:"startTime" xml:"startTime"`
	State       string    `ec2:"status" xml:"status"`
	Tags        []Tag     `ec2:"tagSet" xml:"tagSet>item"`
	VolumeID    string    `ec2:"volumeId" xml:"volumeId"`
	VolumeSize  int       `ec2:"volumeSize" xml:"volumeSize"`
}

// SpotDatafeedSubscription is undocumented.
type SpotDatafeedSubscription struct {
	Bucket  string                 `ec2:"bucket" xml:"bucket"`
	Fault   SpotInstanceStateFault `ec2:"fault" xml:"fault"`
	OwnerID string                 `ec2:"ownerId" xml:"ownerId"`
	Prefix  string                 `ec2:"prefix" xml:"prefix"`
	State   string                 `ec2:"state" xml:"state"`
}

// SpotInstanceRequest is undocumented.
type SpotInstanceRequest struct {
	AvailabilityZoneGroup    string                 `ec2:"availabilityZoneGroup" xml:"availabilityZoneGroup"`
	CreateTime               time.Time              `ec2:"createTime" xml:"createTime"`
	Fault                    SpotInstanceStateFault `ec2:"fault" xml:"fault"`
	InstanceID               string                 `ec2:"instanceId" xml:"instanceId"`
	LaunchGroup              string                 `ec2:"launchGroup" xml:"launchGroup"`
	LaunchSpecification      LaunchSpecification    `ec2:"launchSpecification" xml:"launchSpecification"`
	LaunchedAvailabilityZone string                 `ec2:"launchedAvailabilityZone" xml:"launchedAvailabilityZone"`
	ProductDescription       string                 `ec2:"productDescription" xml:"productDescription"`
	SpotInstanceRequestID    string                 `ec2:"spotInstanceRequestId" xml:"spotInstanceRequestId"`
	SpotPrice                string                 `ec2:"spotPrice" xml:"spotPrice"`
	State                    string                 `ec2:"state" xml:"state"`
	Status                   SpotInstanceStatus     `ec2:"status" xml:"status"`
	Tags                     []Tag                  `ec2:"tagSet" xml:"tagSet>item"`
	Type                     string                 `ec2:"type" xml:"type"`
	ValidFrom                time.Time              `ec2:"validFrom" xml:"validFrom"`
	ValidUntil               time.Time              `ec2:"validUntil" xml:"validUntil"`
}

// SpotInstanceStateFault is undocumented.
type SpotInstanceStateFault struct {
	Code    string `ec2:"code" xml:"code"`
	Message string `ec2:"message" xml:"message"`
}

// SpotInstanceStatus is undocumented.
type SpotInstanceStatus struct {
	Code       string    `ec2:"code" xml:"code"`
	Message    string    `ec2:"message" xml:"message"`
	UpdateTime time.Time `ec2:"updateTime" xml:"updateTime"`
}

// SpotPlacement is undocumented.
type SpotPlacement struct {
	AvailabilityZone string `ec2:"availabilityZone" xml:"availabilityZone"`
	GroupName        string `ec2:"groupName" xml:"groupName"`
}

// SpotPrice is undocumented.
type SpotPrice struct {
	AvailabilityZone   string    `ec2:"availabilityZone" xml:"availabilityZone"`
	InstanceType       string    `ec2:"instanceType" xml:"instanceType"`
	ProductDescription string    `ec2:"productDescription" xml:"productDescription"`
	SpotPrice          string    `ec2:"spotPrice" xml:"spotPrice"`
	Timestamp          time.Time `ec2:"timestamp" xml:"timestamp"`
}

// StartInstancesRequest is undocumented.
type StartInstancesRequest struct {
	AdditionalInfo string   `ec2:"additionalInfo" xml:"additionalInfo"`
	DryRun         bool     `ec2:"dryRun" xml:"dryRun"`
	InstanceIds    []string `ec2:"InstanceId" xml:"InstanceId>InstanceId"`
}

// StartInstancesResult is undocumented.
type StartInstancesResult struct {
	StartingInstances []InstanceStateChange `ec2:"instancesSet" xml:"instancesSet>item"`
}

// StateReason is undocumented.
type StateReason struct {
	Code    string `ec2:"code" xml:"code"`
	Message string `ec2:"message" xml:"message"`
}

// StopInstancesRequest is undocumented.
type StopInstancesRequest struct {
	DryRun      bool     `ec2:"dryRun" xml:"dryRun"`
	Force       bool     `ec2:"force" xml:"force"`
	InstanceIds []string `ec2:"InstanceId" xml:"InstanceId>InstanceId"`
}

// StopInstancesResult is undocumented.
type StopInstancesResult struct {
	StoppingInstances []InstanceStateChange `ec2:"instancesSet" xml:"instancesSet>item"`
}

// Storage is undocumented.
type Storage struct {
	S3 S3Storage `ec2:"" xml:"S3"`
}

// Subnet is undocumented.
type Subnet struct {
	AvailabilityZone        string `ec2:"availabilityZone" xml:"availabilityZone"`
	AvailableIPAddressCount int    `ec2:"availableIpAddressCount" xml:"availableIpAddressCount"`
	CidrBlock               string `ec2:"cidrBlock" xml:"cidrBlock"`
	DefaultForAz            bool   `ec2:"defaultForAz" xml:"defaultForAz"`
	MapPublicIPOnLaunch     bool   `ec2:"mapPublicIpOnLaunch" xml:"mapPublicIpOnLaunch"`
	State                   string `ec2:"state" xml:"state"`
	SubnetID                string `ec2:"subnetId" xml:"subnetId"`
	Tags                    []Tag  `ec2:"tagSet" xml:"tagSet>item"`
	VpcID                   string `ec2:"vpcId" xml:"vpcId"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `ec2:"key" xml:"key"`
	Value string `ec2:"value" xml:"value"`
}

// TagDescription is undocumented.
type TagDescription struct {
	Key          string `ec2:"key" xml:"key"`
	ResourceID   string `ec2:"resourceId" xml:"resourceId"`
	ResourceType string `ec2:"resourceType" xml:"resourceType"`
	Value        string `ec2:"value" xml:"value"`
}

// TerminateInstancesRequest is undocumented.
type TerminateInstancesRequest struct {
	DryRun      bool     `ec2:"dryRun" xml:"dryRun"`
	InstanceIds []string `ec2:"InstanceId" xml:"InstanceId>InstanceId"`
}

// TerminateInstancesResult is undocumented.
type TerminateInstancesResult struct {
	TerminatingInstances []InstanceStateChange `ec2:"instancesSet" xml:"instancesSet>item"`
}

// UnassignPrivateIPAddressesRequest is undocumented.
type UnassignPrivateIPAddressesRequest struct {
	NetworkInterfaceID string   `ec2:"networkInterfaceId" xml:"networkInterfaceId"`
	PrivateIPAddresses []string `ec2:"privateIpAddress" xml:"privateIpAddress>PrivateIpAddress"`
}

// UnmonitorInstancesRequest is undocumented.
type UnmonitorInstancesRequest struct {
	DryRun      bool     `ec2:"dryRun" xml:"dryRun"`
	InstanceIds []string `ec2:"InstanceId" xml:"InstanceId>InstanceId"`
}

// UnmonitorInstancesResult is undocumented.
type UnmonitorInstancesResult struct {
	InstanceMonitorings []InstanceMonitoring `ec2:"instancesSet" xml:"instancesSet>item"`
}

// UserIDGroupPair is undocumented.
type UserIDGroupPair struct {
	GroupID   string `ec2:"groupId" xml:"groupId"`
	GroupName string `ec2:"groupName" xml:"groupName"`
	UserID    string `ec2:"userId" xml:"userId"`
}

// VgwTelemetry is undocumented.
type VgwTelemetry struct {
	AcceptedRouteCount int       `ec2:"acceptedRouteCount" xml:"acceptedRouteCount"`
	LastStatusChange   time.Time `ec2:"lastStatusChange" xml:"lastStatusChange"`
	OutsideIPAddress   string    `ec2:"outsideIpAddress" xml:"outsideIpAddress"`
	Status             string    `ec2:"status" xml:"status"`
	StatusMessage      string    `ec2:"statusMessage" xml:"statusMessage"`
}

// Volume is undocumented.
type Volume struct {
	Attachments      []VolumeAttachment `ec2:"attachmentSet" xml:"attachmentSet>item"`
	AvailabilityZone string             `ec2:"availabilityZone" xml:"availabilityZone"`
	CreateTime       time.Time          `ec2:"createTime" xml:"createTime"`
	Encrypted        bool               `ec2:"encrypted" xml:"encrypted"`
	Iops             int                `ec2:"iops" xml:"iops"`
	KmsKeyID         string             `ec2:"kmsKeyId" xml:"kmsKeyId"`
	Size             int                `ec2:"size" xml:"size"`
	SnapshotID       string             `ec2:"snapshotId" xml:"snapshotId"`
	State            string             `ec2:"status" xml:"status"`
	Tags             []Tag              `ec2:"tagSet" xml:"tagSet>item"`
	VolumeID         string             `ec2:"volumeId" xml:"volumeId"`
	VolumeType       string             `ec2:"volumeType" xml:"volumeType"`
}

// VolumeAttachment is undocumented.
type VolumeAttachment struct {
	AttachTime          time.Time `ec2:"attachTime" xml:"attachTime"`
	DeleteOnTermination bool      `ec2:"deleteOnTermination" xml:"deleteOnTermination"`
	Device              string    `ec2:"device" xml:"device"`
	InstanceID          string    `ec2:"instanceId" xml:"instanceId"`
	State               string    `ec2:"status" xml:"status"`
	VolumeID            string    `ec2:"volumeId" xml:"volumeId"`
}

// VolumeDetail is undocumented.
type VolumeDetail struct {
	Size int64 `ec2:"size" xml:"size"`
}

// VolumeStatusAction is undocumented.
type VolumeStatusAction struct {
	Code        string `ec2:"code" xml:"code"`
	Description string `ec2:"description" xml:"description"`
	EventID     string `ec2:"eventId" xml:"eventId"`
	EventType   string `ec2:"eventType" xml:"eventType"`
}

// VolumeStatusDetails is undocumented.
type VolumeStatusDetails struct {
	Name   string `ec2:"name" xml:"name"`
	Status string `ec2:"status" xml:"status"`
}

// VolumeStatusEvent is undocumented.
type VolumeStatusEvent struct {
	Description string    `ec2:"description" xml:"description"`
	EventID     string    `ec2:"eventId" xml:"eventId"`
	EventType   string    `ec2:"eventType" xml:"eventType"`
	NotAfter    time.Time `ec2:"notAfter" xml:"notAfter"`
	NotBefore   time.Time `ec2:"notBefore" xml:"notBefore"`
}

// VolumeStatusInfo is undocumented.
type VolumeStatusInfo struct {
	Details []VolumeStatusDetails `ec2:"details" xml:"details>item"`
	Status  string                `ec2:"status" xml:"status"`
}

// VolumeStatusItem is undocumented.
type VolumeStatusItem struct {
	Actions          []VolumeStatusAction `ec2:"actionsSet" xml:"actionsSet>item"`
	AvailabilityZone string               `ec2:"availabilityZone" xml:"availabilityZone"`
	Events           []VolumeStatusEvent  `ec2:"eventsSet" xml:"eventsSet>item"`
	VolumeID         string               `ec2:"volumeId" xml:"volumeId"`
	VolumeStatus     VolumeStatusInfo     `ec2:"volumeStatus" xml:"volumeStatus"`
}

// Vpc is undocumented.
type Vpc struct {
	CidrBlock       string `ec2:"cidrBlock" xml:"cidrBlock"`
	DhcpOptionsID   string `ec2:"dhcpOptionsId" xml:"dhcpOptionsId"`
	InstanceTenancy string `ec2:"instanceTenancy" xml:"instanceTenancy"`
	IsDefault       bool   `ec2:"isDefault" xml:"isDefault"`
	State           string `ec2:"state" xml:"state"`
	Tags            []Tag  `ec2:"tagSet" xml:"tagSet>item"`
	VpcID           string `ec2:"vpcId" xml:"vpcId"`
}

// VpcAttachment is undocumented.
type VpcAttachment struct {
	State string `ec2:"state" xml:"state"`
	VpcID string `ec2:"vpcId" xml:"vpcId"`
}

// VpcPeeringConnection is undocumented.
type VpcPeeringConnection struct {
	AccepterVpcInfo        VpcPeeringConnectionVpcInfo     `ec2:"accepterVpcInfo" xml:"accepterVpcInfo"`
	ExpirationTime         time.Time                       `ec2:"expirationTime" xml:"expirationTime"`
	RequesterVpcInfo       VpcPeeringConnectionVpcInfo     `ec2:"requesterVpcInfo" xml:"requesterVpcInfo"`
	Status                 VpcPeeringConnectionStateReason `ec2:"status" xml:"status"`
	Tags                   []Tag                           `ec2:"tagSet" xml:"tagSet>item"`
	VpcPeeringConnectionID string                          `ec2:"vpcPeeringConnectionId" xml:"vpcPeeringConnectionId"`
}

// VpcPeeringConnectionStateReason is undocumented.
type VpcPeeringConnectionStateReason struct {
	Code    string `ec2:"code" xml:"code"`
	Message string `ec2:"message" xml:"message"`
}

// VpcPeeringConnectionVpcInfo is undocumented.
type VpcPeeringConnectionVpcInfo struct {
	CidrBlock string `ec2:"cidrBlock" xml:"cidrBlock"`
	OwnerID   string `ec2:"ownerId" xml:"ownerId"`
	VpcID     string `ec2:"vpcId" xml:"vpcId"`
}

// VpnConnection is undocumented.
type VpnConnection struct {
	CustomerGatewayConfiguration string               `ec2:"customerGatewayConfiguration" xml:"customerGatewayConfiguration"`
	CustomerGatewayID            string               `ec2:"customerGatewayId" xml:"customerGatewayId"`
	Options                      VpnConnectionOptions `ec2:"options" xml:"options"`
	Routes                       []VpnStaticRoute     `ec2:"routes" xml:"routes>item"`
	State                        string               `ec2:"state" xml:"state"`
	Tags                         []Tag                `ec2:"tagSet" xml:"tagSet>item"`
	Type                         string               `ec2:"type" xml:"type"`
	VgwTelemetry                 []VgwTelemetry       `ec2:"vgwTelemetry" xml:"vgwTelemetry>item"`
	VpnConnectionID              string               `ec2:"vpnConnectionId" xml:"vpnConnectionId"`
	VpnGatewayID                 string               `ec2:"vpnGatewayId" xml:"vpnGatewayId"`
}

// VpnConnectionOptions is undocumented.
type VpnConnectionOptions struct {
	StaticRoutesOnly bool `ec2:"staticRoutesOnly" xml:"staticRoutesOnly"`
}

// VpnConnectionOptionsSpecification is undocumented.
type VpnConnectionOptionsSpecification struct {
	StaticRoutesOnly bool `ec2:"staticRoutesOnly" xml:"staticRoutesOnly"`
}

// VpnGateway is undocumented.
type VpnGateway struct {
	AvailabilityZone string          `ec2:"availabilityZone" xml:"availabilityZone"`
	State            string          `ec2:"state" xml:"state"`
	Tags             []Tag           `ec2:"tagSet" xml:"tagSet>item"`
	Type             string          `ec2:"type" xml:"type"`
	VpcAttachments   []VpcAttachment `ec2:"attachments" xml:"attachments>item"`
	VpnGatewayID     string          `ec2:"vpnGatewayId" xml:"vpnGatewayId"`
}

// VpnStaticRoute is undocumented.
type VpnStaticRoute struct {
	DestinationCidrBlock string `ec2:"destinationCidrBlock" xml:"destinationCidrBlock"`
	Source               string `ec2:"source" xml:"source"`
	State                string `ec2:"state" xml:"state"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
