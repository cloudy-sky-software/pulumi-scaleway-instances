// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetVolumeArgs, GetVolumeResult, GetVolumeOutputArgs } from "./getVolume";
export const getVolume: typeof import("./getVolume").getVolume = null as any;
export const getVolumeOutput: typeof import("./getVolume").getVolumeOutput = null as any;
utilities.lazyLoad(exports, ["getVolume","getVolumeOutput"], () => require("./getVolume"));

export { ListVolumesArgs, ListVolumesResult, ListVolumesOutputArgs } from "./listVolumes";
export const listVolumes: typeof import("./listVolumes").listVolumes = null as any;
export const listVolumesOutput: typeof import("./listVolumes").listVolumesOutput = null as any;
utilities.lazyLoad(exports, ["listVolumes","listVolumesOutput"], () => require("./listVolumes"));

export { ListVolumesTypesArgs, ListVolumesTypesResult, ListVolumesTypesOutputArgs } from "./listVolumesTypes";
export const listVolumesTypes: typeof import("./listVolumesTypes").listVolumesTypes = null as any;
export const listVolumesTypesOutput: typeof import("./listVolumesTypes").listVolumesTypesOutput = null as any;
utilities.lazyLoad(exports, ["listVolumesTypes","listVolumesTypesOutput"], () => require("./listVolumesTypes"));

export { VolumeArgs } from "./volume";
export type Volume = import("./volume").Volume;
export const Volume: typeof import("./volume").Volume = null as any;
utilities.lazyLoad(exports, ["Volume"], () => require("./volume"));


// Export enums:
export * from "../types/enums/volumes";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway-instances:volumes:Volume":
                return new Volume(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway-instances", "volumes", _module)