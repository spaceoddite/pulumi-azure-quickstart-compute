// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Virtualmachine extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'azure-quickstart-compute:index:virtualmachine';

    /**
     * Returns true if the given object is an instance of Virtualmachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Virtualmachine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Virtualmachine.__pulumiType;
    }


    /**
     * Create a Virtualmachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualmachineArgs, opts?: pulumi.ComponentResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.adminPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminPassword'");
            }
            if ((!args || args.adminUsername === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminUsername'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.vmSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmSize'");
            }
            inputs["adminPassword"] = args ? args.adminPassword : undefined;
            inputs["adminUsername"] = args ? args.adminUsername : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vmSize"] = args ? args.vmSize : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Virtualmachine.__pulumiType, name, inputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Virtualmachine resource.
 */
export interface VirtualmachineArgs {
    /**
     * admin password
     */
    adminPassword: pulumi.Input<string>;
    /**
     * admin username
     */
    adminUsername: pulumi.Input<string>;
    /**
     * location
     */
    location: pulumi.Input<string>;
    /**
     * Name of your virtual machine
     */
    name: pulumi.Input<string>;
    /**
     * vmsize
     */
    vmSize: pulumi.Input<string>;
}
