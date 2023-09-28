// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway-instances:servers:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * The boot type to use
     */
    public readonly boot_type!: pulumi.Output<enums.servers.BootType | undefined>;
    /**
     * The bootscript ID to use when `boot_type` is set to `bootscript`
     */
    public readonly bootscript!: pulumi.Output<string | undefined>;
    /**
     * Define the server commercial type (i.e. GP1-S)
     */
    public readonly commercial_type!: pulumi.Output<string>;
    /**
     * Define if a dynamic IP is required for the instance
     */
    public readonly dynamic_ip_required!: pulumi.Output<boolean | undefined>;
    /**
     * True if IPv6 is enabled on the server
     */
    public readonly enable_ipv6!: pulumi.Output<boolean | undefined>;
    /**
     * The server image ID
     */
    public readonly image!: pulumi.Output<string | undefined>;
    /**
     * The server name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The server organization ID
     */
    public readonly organization!: pulumi.Output<string | undefined>;
    /**
     * Placement group ID if server must be part of a placement group
     */
    public readonly placement_group!: pulumi.Output<string | undefined>;
    /**
     * The server project ID
     */
    public readonly project!: pulumi.Output<string | undefined>;
    /**
     * The ID of the reserved IP to attach to the server
     */
    public readonly public_ip!: pulumi.Output<string | undefined>;
    /**
     * The security group ID
     */
    public readonly security_group!: pulumi.Output<string | undefined>;
    public /*out*/ readonly server!: pulumi.Output<outputs.servers.ScalewayInstanceV1Server | undefined>;
    /**
     * The server tags
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly volumes!: pulumi.Output<{[key: string]: outputs.servers.ScalewayInstanceV1VolumeServerTemplate} | undefined>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.commercial_type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'commercial_type'");
            }
            resourceInputs["boot_type"] = (args ? args.boot_type : undefined) ?? "local";
            resourceInputs["bootscript"] = args ? args.bootscript : undefined;
            resourceInputs["commercial_type"] = args ? args.commercial_type : undefined;
            resourceInputs["dynamic_ip_required"] = args ? args.dynamic_ip_required : undefined;
            resourceInputs["enable_ipv6"] = args ? args.enable_ipv6 : undefined;
            resourceInputs["image"] = args ? args.image : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["placement_group"] = args ? args.placement_group : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["public_ip"] = args ? args.public_ip : undefined;
            resourceInputs["security_group"] = args ? args.security_group : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volumes"] = args ? args.volumes : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["server"] = undefined /*out*/;
        } else {
            resourceInputs["boot_type"] = undefined /*out*/;
            resourceInputs["bootscript"] = undefined /*out*/;
            resourceInputs["commercial_type"] = undefined /*out*/;
            resourceInputs["dynamic_ip_required"] = undefined /*out*/;
            resourceInputs["enable_ipv6"] = undefined /*out*/;
            resourceInputs["image"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["organization"] = undefined /*out*/;
            resourceInputs["placement_group"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["public_ip"] = undefined /*out*/;
            resourceInputs["security_group"] = undefined /*out*/;
            resourceInputs["server"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["volumes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Server.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * The boot type to use
     */
    boot_type?: pulumi.Input<enums.servers.BootType>;
    /**
     * The bootscript ID to use when `boot_type` is set to `bootscript`
     */
    bootscript?: pulumi.Input<string>;
    /**
     * Define the server commercial type (i.e. GP1-S)
     */
    commercial_type: pulumi.Input<string>;
    /**
     * Define if a dynamic IP is required for the instance
     */
    dynamic_ip_required?: pulumi.Input<boolean>;
    /**
     * True if IPv6 is enabled on the server
     */
    enable_ipv6?: pulumi.Input<boolean>;
    /**
     * The server image ID
     */
    image?: pulumi.Input<string>;
    /**
     * The server name
     */
    name?: pulumi.Input<string>;
    /**
     * The server organization ID
     */
    organization?: pulumi.Input<string>;
    /**
     * Placement group ID if server must be part of a placement group
     */
    placement_group?: pulumi.Input<string>;
    /**
     * The server project ID
     */
    project?: pulumi.Input<string>;
    /**
     * The ID of the reserved IP to attach to the server
     */
    public_ip?: pulumi.Input<string>;
    /**
     * The security group ID
     */
    security_group?: pulumi.Input<string>;
    /**
     * The server tags
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    volumes?: pulumi.Input<{[key: string]: pulumi.Input<inputs.servers.ScalewayInstanceV1VolumeServerTemplateArgs>}>;
    /**
     * The zone you want to target
     */
    zone?: pulumi.Input<string>;
}
