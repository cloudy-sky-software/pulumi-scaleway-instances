// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * The volume on which the snapshot is based on
 */
export interface BaseVolumeProperties {
    /**
     * The volume ID on which the snapshot is based on
     */
    id?: string;
    /**
     * The volume name on which the snapshot is based on
     */
    name?: string;
}

export interface ScalewayInstanceV1GetSnapshotResponse {
    snapshot?: outputs.snapshots.ScalewayInstanceV1Snapshot;
}
/**
 * scalewayInstanceV1GetSnapshotResponseProvideDefaults sets the appropriate defaults for ScalewayInstanceV1GetSnapshotResponse
 */
export function scalewayInstanceV1GetSnapshotResponseProvideDefaults(val: ScalewayInstanceV1GetSnapshotResponse): ScalewayInstanceV1GetSnapshotResponse {
    return {
        ...val,
        snapshot: (val.snapshot ? outputs.snapshots.scalewayInstanceV1SnapshotProvideDefaults(val.snapshot) : undefined),
    };
}

export interface ScalewayInstanceV1ListSnapshotsResponse {
    /**
     * List of snapshots
     */
    snapshots?: outputs.snapshots.ScalewayInstanceV1Snapshot[];
}

export interface ScalewayInstanceV1Snapshot {
    /**
     * The volume on which the snapshot is based on
     */
    base_volume?: outputs.snapshots.ScalewayInstanceV1SnapshotBaseVolumeProperties;
    /**
     * The snapshot creation date (RFC 3339 format)
     */
    creation_date?: string;
    /**
     * The reason for the failed snapshot import
     */
    error_reason?: string;
    id?: string;
    /**
     * The snapshot modification date (RFC 3339 format)
     */
    modification_date?: string;
    /**
     * The snapshot name
     */
    name?: string;
    /**
     * The snapshot organization ID
     */
    organization?: string;
    /**
     * The snapshot project ID
     */
    project?: string;
    /**
     * The snapshot size (in bytes)
     */
    size?: number;
    state?: enums.snapshots.ScalewayInstanceV1SnapshotState;
    /**
     * The snapshot tags
     */
    tags?: string[];
    volume_type?: enums.snapshots.ScalewayInstanceV1SnapshotVolumeType;
    /**
     * The snapshot zone
     */
    zone?: string;
}
/**
 * scalewayInstanceV1SnapshotProvideDefaults sets the appropriate defaults for ScalewayInstanceV1Snapshot
 */
export function scalewayInstanceV1SnapshotProvideDefaults(val: ScalewayInstanceV1Snapshot): ScalewayInstanceV1Snapshot {
    return {
        ...val,
        state: (val.state) ?? "available",
        volume_type: (val.volume_type) ?? "l_ssd",
    };
}

/**
 * The volume on which the snapshot is based on
 */
export interface ScalewayInstanceV1SnapshotBaseVolumeProperties {
    /**
     * The volume ID on which the snapshot is based on
     */
    id?: string;
    /**
     * The volume name on which the snapshot is based on
     */
    name?: string;
}
