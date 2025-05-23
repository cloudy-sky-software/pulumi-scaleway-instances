// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getPlacementGroup(args: GetPlacementGroupArgs, opts?: pulumi.InvokeOptions): Promise<outputs.placement_groups.ScalewayInstanceV1GetPlacementGroupResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:placement_groups:getPlacementGroup", {
        "id": args.id,
        "zone": args.zone,
    }, opts);
}

export interface GetPlacementGroupArgs {
    /**
     * UUID of the placement group you want to get
     */
    id: string;
    /**
     * The zone you want to target
     */
    zone: string;
}
export function getPlacementGroupOutput(args: GetPlacementGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.placement_groups.ScalewayInstanceV1GetPlacementGroupResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway-instances:placement_groups:getPlacementGroup", {
        "id": args.id,
        "zone": args.zone,
    }, opts);
}

export interface GetPlacementGroupOutputArgs {
    /**
     * UUID of the placement group you want to get
     */
    id: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
