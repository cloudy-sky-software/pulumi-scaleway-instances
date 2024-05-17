// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway-instances:volumes:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    public /*out*/ readonly Location!: pulumi.Output<string | undefined>;
    /**
     * The volume creation date (RFC 3339 format)
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string | undefined>;
    /**
     * Show the volume NBD export URI
     */
    public /*out*/ readonly exportUri!: pulumi.Output<string | undefined>;
    /**
     * The volume modification date (RFC 3339 format)
     */
    public /*out*/ readonly modificationDate!: pulumi.Output<string | undefined>;
    /**
     * The volume name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The volume organization ID
     */
    public readonly organization!: pulumi.Output<string | undefined>;
    /**
     * The volume project ID
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The server attached to the volume
     */
    public /*out*/ readonly server!: pulumi.Output<outputs.volumes.ServerProperties | undefined>;
    /**
     * The volume disk size (in bytes)
     */
    public readonly size!: pulumi.Output<number | undefined>;
    public /*out*/ readonly state!: pulumi.Output<enums.volumes.VolumeState | undefined>;
    /**
     * The volume tags
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly volume!: pulumi.Output<outputs.volumes.ScalewayInstanceV1Volume | undefined>;
    public readonly volumeType!: pulumi.Output<enums.volumes.VolumeType | undefined>;
    /**
     * The zone in which is the volume
     */
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volumeType"] = (args ? args.volumeType : undefined) ?? "l_ssd";
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["Location"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["exportUri"] = undefined /*out*/;
            resourceInputs["modificationDate"] = undefined /*out*/;
            resourceInputs["server"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["volume"] = undefined /*out*/;
        } else {
            resourceInputs["Location"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["exportUri"] = undefined /*out*/;
            resourceInputs["modificationDate"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["organization"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["server"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["volume"] = undefined /*out*/;
            resourceInputs["volumeType"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Volume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The volume name
     */
    name?: pulumi.Input<string>;
    /**
     * The volume organization ID
     */
    organization?: pulumi.Input<string>;
    /**
     * The volume project ID
     */
    project: pulumi.Input<string>;
    /**
     * The volume disk size (in bytes)
     */
    size?: pulumi.Input<number>;
    /**
     * The volume tags
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    volumeType?: pulumi.Input<enums.volumes.VolumeType>;
    /**
     * The zone you want to target
     */
    zone?: pulumi.Input<string>;
}
