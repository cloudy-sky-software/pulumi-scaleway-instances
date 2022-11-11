// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Action
{
    public static class ListServerActions
    {
        public static Task<ListServerActionsResult> InvokeAsync(ListServerActionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListServerActionsResult>("scaleway-instances:action:listServerActions", args ?? new ListServerActionsArgs(), options.WithDefaults());

        public static Output<ListServerActionsResult> Invoke(ListServerActionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListServerActionsResult>("scaleway-instances:action:listServerActions", args ?? new ListServerActionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListServerActionsArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public ListServerActionsArgs()
        {
        }
        public static new ListServerActionsArgs Empty => new ListServerActionsArgs();
    }

    public sealed class ListServerActionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ListServerActionsInvokeArgs()
        {
        }
        public static new ListServerActionsInvokeArgs Empty => new ListServerActionsInvokeArgs();
    }


    [OutputType]
    public sealed class ListServerActionsResult
    {
        public readonly Outputs.ScalewayInstanceV1ListServerActionsResponse Items;

        [OutputConstructor]
        private ListServerActionsResult(Outputs.ScalewayInstanceV1ListServerActionsResponse items)
        {
            Items = items;
        }
    }
}
