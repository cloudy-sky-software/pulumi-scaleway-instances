// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class SecurityGroupRule extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroupRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecurityGroupRule {
        return new SecurityGroupRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway-instances:rules:SecurityGroupRule';

    /**
     * Returns true if the given object is an instance of SecurityGroupRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroupRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroupRule.__pulumiType;
    }

    public readonly action!: pulumi.Output<enums.rules.Action>;
    /**
     * The beginning of the range of ports to apply this rule to (inclusive)
     */
    public readonly dest_port_from!: pulumi.Output<number | undefined>;
    /**
     * The end of the range of ports to apply this rule to (inclusive)
     */
    public readonly dest_port_to!: pulumi.Output<number | undefined>;
    public readonly direction!: pulumi.Output<enums.rules.Direction>;
    /**
     * Indicates if this rule is editable (will be ignored)
     */
    public readonly editable!: pulumi.Output<boolean | undefined>;
    /**
     * (IP network)
     */
    public readonly ip_range!: pulumi.Output<string>;
    /**
     * The position of this rule in the security group rules list
     */
    public readonly position!: pulumi.Output<number | undefined>;
    public readonly protocol!: pulumi.Output<enums.rules.Protocol>;

    /**
     * Create a SecurityGroupRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.ip_range === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip_range'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["action"] = (args ? args.action : undefined) ?? "accept";
            resourceInputs["dest_port_from"] = args ? args.dest_port_from : undefined;
            resourceInputs["dest_port_to"] = args ? args.dest_port_to : undefined;
            resourceInputs["direction"] = (args ? args.direction : undefined) ?? "inbound";
            resourceInputs["editable"] = args ? args.editable : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["ip_range"] = args ? args.ip_range : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["protocol"] = (args ? args.protocol : undefined) ?? "TCP";
            resourceInputs["zone"] = args ? args.zone : undefined;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["dest_port_from"] = undefined /*out*/;
            resourceInputs["dest_port_to"] = undefined /*out*/;
            resourceInputs["direction"] = undefined /*out*/;
            resourceInputs["editable"] = undefined /*out*/;
            resourceInputs["ip_range"] = undefined /*out*/;
            resourceInputs["position"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroupRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecurityGroupRule resource.
 */
export interface SecurityGroupRuleArgs {
    action: pulumi.Input<enums.rules.Action>;
    /**
     * The beginning of the range of ports to apply this rule to (inclusive)
     */
    dest_port_from?: pulumi.Input<number>;
    /**
     * The end of the range of ports to apply this rule to (inclusive)
     */
    dest_port_to?: pulumi.Input<number>;
    direction: pulumi.Input<enums.rules.Direction>;
    /**
     * Indicates if this rule is editable (will be ignored)
     */
    editable?: pulumi.Input<boolean>;
    /**
     * UUID of the security group
     */
    id?: pulumi.Input<string>;
    /**
     * (IP network)
     */
    ip_range: pulumi.Input<string>;
    /**
     * The position of this rule in the security group rules list
     */
    position?: pulumi.Input<number>;
    protocol: pulumi.Input<enums.rules.Protocol>;
    /**
     * The zone you want to target
     */
    zone?: pulumi.Input<string>;
}
