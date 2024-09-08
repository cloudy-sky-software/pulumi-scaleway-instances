// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Volumes
{
    public static class ListVolumesTypes
    {
        public static Task<Outputs.ScalewayInstanceV1ListVolumesTypesResponse> InvokeAsync(ListVolumesTypesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ScalewayInstanceV1ListVolumesTypesResponse>("scaleway-instances:volumes:listVolumesTypes", args ?? new ListVolumesTypesArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayInstanceV1ListVolumesTypesResponse> Invoke(ListVolumesTypesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayInstanceV1ListVolumesTypesResponse>("scaleway-instances:volumes:listVolumesTypes", args ?? new ListVolumesTypesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListVolumesTypesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public ListVolumesTypesArgs()
        {
        }
        public static new ListVolumesTypesArgs Empty => new ListVolumesTypesArgs();
    }

    public sealed class ListVolumesTypesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ListVolumesTypesInvokeArgs()
        {
        }
        public static new ListVolumesTypesInvokeArgs Empty => new ListVolumesTypesInvokeArgs();
    }
}
