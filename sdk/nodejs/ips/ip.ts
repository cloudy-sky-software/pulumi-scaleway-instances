// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Ip extends pulumi.CustomResource {
    /**
     * Get an existing Ip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Ip {
        return new Ip(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway-instances:ips:Ip';

    /**
     * Returns true if the given object is an instance of Ip.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ip {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ip.__pulumiType;
    }

    /**
     * (IPv4 address)
     */
    public /*out*/ readonly address!: pulumi.Output<string | undefined>;
    public readonly organization!: pulumi.Output<string | undefined>;
    public readonly project!: pulumi.Output<string | undefined>;
    public readonly reverse!: pulumi.Output<outputs.ips.GoogleProtobufStringValue | undefined>;
    public readonly server!: pulumi.Output<outputs.ips.ScalewayInstanceV1ServerSummary | undefined>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a Ip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["reverse"] = args ? args.reverse : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["address"] = undefined /*out*/;
        } else {
            resourceInputs["address"] = undefined /*out*/;
            resourceInputs["organization"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["reverse"] = undefined /*out*/;
            resourceInputs["server"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ip.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Ip resource.
 */
export interface IpArgs {
    organization?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    reverse?: pulumi.Input<inputs.ips.GoogleProtobufStringValueArgs>;
    server?: pulumi.Input<inputs.ips.ScalewayInstanceV1ServerSummaryArgs>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone you want to target
     */
    zone?: pulumi.Input<string>;
}