// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ScalewayInstanceV1TaskStatus = {
    Pending: "pending",
    Started: "started",
    Success: "success",
    Failure: "failure",
    Retry: "retry",
} as const;

/**
 * The task status
 */
export type ScalewayInstanceV1TaskStatus = (typeof ScalewayInstanceV1TaskStatus)[keyof typeof ScalewayInstanceV1TaskStatus];
