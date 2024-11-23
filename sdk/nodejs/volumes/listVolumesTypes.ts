// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listVolumesTypes(args: ListVolumesTypesArgs, opts?: pulumi.InvokeOptions): Promise<outputs.volumes.ScalewayInstanceV1ListVolumesTypesResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:volumes:listVolumesTypes", {
        "zone": args.zone,
    }, opts);
}

export interface ListVolumesTypesArgs {
    /**
     * The zone you want to target
     */
    zone: string;
}
export function listVolumesTypesOutput(args: ListVolumesTypesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.volumes.ScalewayInstanceV1ListVolumesTypesResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway-instances:volumes:listVolumesTypes", {
        "zone": args.zone,
    }, opts);
}

export interface ListVolumesTypesOutputArgs {
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
