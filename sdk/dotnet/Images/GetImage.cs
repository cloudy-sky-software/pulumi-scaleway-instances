// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Images
{
    public static class GetImage
    {
        public static Task<GetImageResult> InvokeAsync(GetImageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("scaleway-instances:images:getImage", args ?? new GetImageArgs(), options.WithDefaults());

        public static Output<GetImageResult> Invoke(GetImageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageResult>("scaleway-instances:images:getImage", args ?? new GetImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the image you want to get
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetImageArgs()
        {
        }
        public static new GetImageArgs Empty => new GetImageArgs();
    }

    public sealed class GetImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the image you want to get
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetImageInvokeArgs()
        {
        }
        public static new GetImageInvokeArgs Empty => new GetImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetImageResult
    {
        public readonly Outputs.ScalewayInstanceV1GetImageResponse Items;

        [OutputConstructor]
        private GetImageResult(Outputs.ScalewayInstanceV1GetImageResponse items)
        {
            Items = items;
        }
    }
}