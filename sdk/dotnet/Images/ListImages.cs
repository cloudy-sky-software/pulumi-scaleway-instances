// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Images
{
    public static class ListImages
    {
        public static Task<ListImagesResult> InvokeAsync(ListImagesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListImagesResult>("scaleway-instances:images:listImages", args ?? new ListImagesArgs(), options.WithDefaults());

        public static Output<ListImagesResult> Invoke(ListImagesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListImagesResult>("scaleway-instances:images:listImages", args ?? new ListImagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListImagesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public ListImagesArgs()
        {
        }
        public static new ListImagesArgs Empty => new ListImagesArgs();
    }

    public sealed class ListImagesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ListImagesInvokeArgs()
        {
        }
        public static new ListImagesInvokeArgs Empty => new ListImagesInvokeArgs();
    }


    [OutputType]
    public sealed class ListImagesResult
    {
        public readonly Outputs.ScalewayInstanceV1ListImagesResponse Items;

        [OutputConstructor]
        private ListImagesResult(Outputs.ScalewayInstanceV1ListImagesResponse items)
        {
            Items = items;
        }
    }
}
