package provider

import (
	//"fmt"

	compute "github.com/pulumi/pulumi-azure-native/sdk/go/azure/compute"
	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
	resources "github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type virtualmachineArgs struct {
	Name          pulumi.StringInput `pulumi:"name"`
	AdminUsername pulumi.StringInput `pulumi:"adminUsername"`
	AdminPassword pulumi.StringInput `pulumi:"adminPassword"`
	//	ImageType pulumi.StringInput `pulumi:"imageType"`
	Location pulumi.StringInput `pulumi:"location"`
	VmSize   pulumi.StringInput `pulumi:"vmSize"`
}

type virtualmachine struct {
	pulumi.ResourceState
}

func Newvirtualmachine(ctx *pulumi.Context,
	name string, args *virtualmachineArgs, opts ...pulumi.ResourceOption) (*virtualmachine, error) {
	if args == nil {
		args = &virtualmachineArgs{}
	}

	component := &virtualmachine{}

	err := ctx.RegisterComponentResource("azure-quickstart-compute:index:virtualmachine", name, component, opts...)
	if err != nil {
		return nil, err
	}
	// required parameters
	VmNameParam := args.Name
	AdminUsernameParam := args.AdminUsername
	AdminPasswordParam := args.AdminPassword
	//	ImageTypeParam := args.ImageType
	LocationParam := args.Location
	VmSizeParam := args.VmSize

	// resources
	// resourcegroup
	resourceGroupVar, err := resources.NewResourceGroup(ctx, "resourceGroup", &resources.ResourceGroupArgs{
		Location: LocationParam,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	//networkSecurityGroupResource
	networkSecurityGroupResource, err := network.NewNetworkSecurityGroup(ctx, "networkSecurityGroupResource", &network.NetworkSecurityGroupArgs{
		Location:          resourceGroupVar.Location,
		ResourceGroupName: resourceGroupVar.Name,
		SecurityRules: network.SecurityRuleTypeArray{
			&network.SecurityRuleTypeArgs{
				Access:                   pulumi.String("Allow"),
				DestinationAddressPrefix: pulumi.String("*"),
				DestinationPortRange:     pulumi.String("22"),
				Direction:                pulumi.String("Inbound"),
				Name:                     pulumi.String("SSH"),
				Priority:                 pulumi.Int(1000),
				Protocol:                 pulumi.String("Tcp"),
				SourceAddressPrefix:      pulumi.String("*"),
				SourcePortRange:          pulumi.String("*"),
			},
		},
	}, pulumi.Parent(resourceGroupVar))
	if err != nil {
		return nil, err
	}

	//public IP address resource

	publicIPAddressResource, err := network.NewPublicIPAddress(ctx, "publicIPAddressResource", &network.PublicIPAddressArgs{
		IdleTimeoutInMinutes:     pulumi.Int(4),
		Location:                 resourceGroupVar.Location,
		PublicIPAddressVersion:   pulumi.String("IPv4"),
		PublicIPAllocationMethod: pulumi.String("Dynamic"),
		ResourceGroupName:        resourceGroupVar.Name,
		Sku: &network.PublicIPAddressSkuArgs{
			Name: pulumi.String("Basic"),
		},
	}, pulumi.Parent(resourceGroupVar))
	if err != nil {
		return nil, err
	}

	//virtualNetworkResource
	virtualNetworkResource, err := network.NewVirtualNetwork(ctx, "virtualNetworkResource", &network.VirtualNetworkArgs{
		AddressSpace: &network.AddressSpaceArgs{
			AddressPrefixes: pulumi.StringArray{
				pulumi.String("10.1.0.0/16"),
			},
		},
		Location:          resourceGroupVar.Location,
		ResourceGroupName: resourceGroupVar.Name,
	}, pulumi.Parent(resourceGroupVar))
	if err != nil {
		return nil, err
	}
	//subnetResource
	subnetResource, err := network.NewSubnet(ctx, "subnetResource", &network.SubnetArgs{
		AddressPrefix:                     pulumi.String("10.1.0.0/24"),
		PrivateEndpointNetworkPolicies:    pulumi.String("Enabled"),
		PrivateLinkServiceNetworkPolicies: pulumi.String("Enabled"),
		ResourceGroupName:                 resourceGroupVar.Name,
		VirtualNetworkName:                virtualNetworkResource.Name,
	}, pulumi.Parent(resourceGroupVar))
	if err != nil {
		return nil, err
	}
	//networkInterfaceResource
	networkInterfaceResource, err := network.NewNetworkInterface(ctx, "networkInterfaceResource", &network.NetworkInterfaceArgs{
		IpConfigurations: network.NetworkInterfaceIPConfigurationArray{
			&network.NetworkInterfaceIPConfigurationArgs{
				Name:                      pulumi.String("ipconfig1"),
				PrivateIPAllocationMethod: pulumi.String("Dynamic"),
				PublicIPAddress: &network.PublicIPAddressTypeArgs{
					Id: publicIPAddressResource.ID(),
				},
				Subnet: &network.SubnetTypeArgs{
					Id: subnetResource.ID(),
				},
			},
		},
		Location: resourceGroupVar.Location,
		NetworkSecurityGroup: &network.NetworkSecurityGroupTypeArgs{
			Id: networkSecurityGroupResource.ID(),
		},
		ResourceGroupName: resourceGroupVar.Name,
	}, pulumi.Parent(resourceGroupVar))
	if err != nil {
		return nil, err
	}

	_, err = compute.NewVirtualMachine(ctx, "virtualMachine", &compute.VirtualMachineArgs{
		HardwareProfile: &compute.HardwareProfileArgs{
			VmSize: VmSizeParam,
		},

		ResourceGroupName: resourceGroupVar.Name,
		Location:          resourceGroupVar.Location,
		NetworkProfile: &compute.NetworkProfileArgs{
			NetworkInterfaces: compute.NetworkInterfaceReferenceArray{
				&compute.NetworkInterfaceReferenceArgs{
					Id: networkInterfaceResource.ID(),
				},
			},
		},
		OsProfile: &compute.OSProfileArgs{
			AdminPassword: AdminPasswordParam,
			AdminUsername: AdminUsernameParam,
			ComputerName:  VmNameParam,
		},
		StorageProfile: &compute.StorageProfileArgs{
			ImageReference: &compute.ImageReferenceArgs{
				Offer:     pulumi.String("UbuntuServer"),
				Publisher: pulumi.String("Canonical"),
				Sku:       pulumi.String("18.04-LTS"),
				Version:   pulumi.String("latest"),
			},
			OsDisk: &compute.OSDiskArgs{
				CreateOption: pulumi.String("FromImage"),
				DiskSizeGB:   pulumi.Int(30),
			},
		},
	}, pulumi.Parent(resourceGroupVar))

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{}); err != nil {
		return nil, err
	}

	return component, nil
}
