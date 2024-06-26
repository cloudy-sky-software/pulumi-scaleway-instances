// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class ServerAction extends pulumi.CustomResource {
    /**
     * Get an existing ServerAction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerAction {
        return new ServerAction(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway-instances:action:ServerAction';

    /**
     * Returns true if the given object is an instance of ServerAction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerAction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerAction.__pulumiType;
    }

    /**
     * The action to perform on the server
     */
    public readonly action!: pulumi.Output<enums.action.Action | undefined>;
    /**
     * The name of the backup you want to create.
     * This field should only be specified when performing a backup action.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public /*out*/ readonly task!: pulumi.Output<outputs.action.ScalewayInstanceV1Task | undefined>;
    public readonly volumes!: pulumi.Output<{[key: string]: outputs.action.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplate} | undefined>;

    /**
     * Create a ServerAction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServerActionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["action"] = (args ? args.action : undefined) ?? "poweron";
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serverId"] = args ? args.serverId : undefined;
            resourceInputs["volumes"] = args ? args.volumes : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["task"] = undefined /*out*/;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["task"] = undefined /*out*/;
            resourceInputs["volumes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerAction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServerAction resource.
 */
export interface ServerActionArgs {
    /**
     * The action to perform on the server
     */
    action?: pulumi.Input<enums.action.Action>;
    /**
     * The name of the backup you want to create.
     * This field should only be specified when performing a backup action.
     */
    name?: pulumi.Input<string>;
    /**
     * UUID of the server
     */
    serverId?: pulumi.Input<string>;
    volumes?: pulumi.Input<{[key: string]: pulumi.Input<inputs.action.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs>}>;
    /**
     * The zone you want to target
     */
    zone?: pulumi.Input<string>;
}
