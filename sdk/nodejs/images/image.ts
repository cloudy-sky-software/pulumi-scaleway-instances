// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway-instances:images:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    public /*out*/ readonly Location!: pulumi.Output<string | undefined>;
    public readonly arch!: pulumi.Output<enums.images.Arch | undefined>;
    /**
     * (RFC 3339 format)
     */
    public /*out*/ readonly creation_date!: pulumi.Output<string | undefined>;
    public readonly default_bootscript!: pulumi.Output<outputs.images.ScalewayInstanceV1Bootscript | undefined>;
    public readonly extra_volumes!: pulumi.Output<{[key: string]: outputs.images.ScalewayInstanceV1Volume} | undefined>;
    public /*out*/ readonly from_server!: pulumi.Output<string | undefined>;
    public /*out*/ readonly image!: pulumi.Output<outputs.images.ScalewayInstanceV1Image | undefined>;
    /**
     * (RFC 3339 format)
     */
    public /*out*/ readonly modification_date!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly organization!: pulumi.Output<string | undefined>;
    public readonly project!: pulumi.Output<string>;
    public readonly public!: pulumi.Output<boolean | undefined>;
    public readonly root_volume!: pulumi.Output<outputs.images.ScalewayInstanceV1VolumeSummary>;
    public readonly state!: pulumi.Output<enums.images.State | undefined>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.root_volume === undefined) && !opts.urn) {
                throw new Error("Missing required property 'root_volume'");
            }
            resourceInputs["arch"] = (args ? args.arch : undefined) ?? "x86_64";
            resourceInputs["default_bootscript"] = args ? (args.default_bootscript ? pulumi.output(args.default_bootscript).apply(inputs.images.scalewayInstanceV1BootscriptArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["extra_volumes"] = args ? args.extra_volumes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["public"] = args ? args.public : undefined;
            resourceInputs["root_volume"] = args ? (args.root_volume ? pulumi.output(args.root_volume).apply(inputs.images.scalewayInstanceV1VolumeSummaryArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["state"] = (args ? args.state : undefined) ?? "available";
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["Location"] = undefined /*out*/;
            resourceInputs["creation_date"] = undefined /*out*/;
            resourceInputs["from_server"] = undefined /*out*/;
            resourceInputs["image"] = undefined /*out*/;
            resourceInputs["modification_date"] = undefined /*out*/;
        } else {
            resourceInputs["Location"] = undefined /*out*/;
            resourceInputs["arch"] = undefined /*out*/;
            resourceInputs["creation_date"] = undefined /*out*/;
            resourceInputs["default_bootscript"] = undefined /*out*/;
            resourceInputs["extra_volumes"] = undefined /*out*/;
            resourceInputs["from_server"] = undefined /*out*/;
            resourceInputs["image"] = undefined /*out*/;
            resourceInputs["modification_date"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["organization"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["public"] = undefined /*out*/;
            resourceInputs["root_volume"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    arch?: pulumi.Input<enums.images.Arch>;
    default_bootscript?: pulumi.Input<inputs.images.ScalewayInstanceV1BootscriptArgs>;
    extra_volumes?: pulumi.Input<{[key: string]: pulumi.Input<inputs.images.ScalewayInstanceV1VolumeArgs>}>;
    name?: pulumi.Input<string>;
    organization?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    public?: pulumi.Input<boolean>;
    root_volume: pulumi.Input<inputs.images.ScalewayInstanceV1VolumeSummaryArgs>;
    state?: pulumi.Input<enums.images.State>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone you want to target
     */
    zone?: pulumi.Input<string>;
}
