// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listServerUserData(args: ListServerUserDataArgs, opts?: pulumi.InvokeOptions): Promise<outputs.user_data.ScalewayInstanceV1ListServerUserDataResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:user_data:listServerUserData", {
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

export interface ListServerUserDataArgs {
    /**
     * UUID of the server
     */
    serverId: string;
    /**
     * The zone you want to target
     */
    zone: string;
}
export function listServerUserDataOutput(args: ListServerUserDataOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.user_data.ScalewayInstanceV1ListServerUserDataResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway-instances:user_data:listServerUserData", {
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

export interface ListServerUserDataOutputArgs {
    /**
     * UUID of the server
     */
    serverId: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
