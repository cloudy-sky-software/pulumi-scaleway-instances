// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listServersTypes(args: ListServersTypesArgs, opts?: pulumi.InvokeOptions): Promise<ListServersTypesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:servers:listServersTypes", {
        "zone": args.zone,
    }, opts);
}

export interface ListServersTypesArgs {
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface ListServersTypesResult {
    readonly items: outputs.servers.ScalewayInstanceV1ListServersTypesResponse;
}

export function listServersTypesOutput(args: ListServersTypesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListServersTypesResult> {
    return pulumi.output(args).apply(a => listServersTypes(a, opts))
}

export interface ListServersTypesOutputArgs {
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
