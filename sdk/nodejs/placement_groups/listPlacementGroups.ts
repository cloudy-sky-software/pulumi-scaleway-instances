// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listPlacementGroups(args: ListPlacementGroupsArgs, opts?: pulumi.InvokeOptions): Promise<ListPlacementGroupsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:placement_groups:listPlacementGroups", {
        "zone": args.zone,
    }, opts);
}

export interface ListPlacementGroupsArgs {
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface ListPlacementGroupsResult {
    readonly items: outputs.placement_groups.ScalewayInstanceV1ListPlacementGroupsResponse;
}

export function listPlacementGroupsOutput(args: ListPlacementGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListPlacementGroupsResult> {
    return pulumi.output(args).apply(a => listPlacementGroups(a, opts))
}

export interface ListPlacementGroupsOutputArgs {
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
