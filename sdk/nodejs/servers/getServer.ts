// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<outputs.servers.ScalewayInstanceV1GetServerResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:servers:getServer", {
        "id": args.id,
        "zone": args.zone,
    }, opts);
}

export interface GetServerArgs {
    /**
     * UUID of the server you want to get
     */
    id: string;
    /**
     * The zone you want to target
     */
    zone: string;
}
export function getServerOutput(args: GetServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.servers.ScalewayInstanceV1GetServerResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway-instances:servers:getServer", {
        "id": args.id,
        "zone": args.zone,
    }, opts);
}

export interface GetServerOutputArgs {
    /**
     * UUID of the server you want to get
     */
    id: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
