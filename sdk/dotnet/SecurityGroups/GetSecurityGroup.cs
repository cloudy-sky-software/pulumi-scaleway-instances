// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.SecurityGroups
{
    public static class GetSecurityGroup
    {
        public static Task<GetSecurityGroupResult> InvokeAsync(GetSecurityGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupResult>("scaleway-instances:security_groups:getSecurityGroup", args ?? new GetSecurityGroupArgs(), options.WithDefaults());

        public static Output<GetSecurityGroupResult> Invoke(GetSecurityGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityGroupResult>("scaleway-instances:security_groups:getSecurityGroup", args ?? new GetSecurityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the security group you want to get
        /// </summary>
        [Input("securityGroupId", required: true)]
        public string SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetSecurityGroupArgs()
        {
        }
        public static new GetSecurityGroupArgs Empty => new GetSecurityGroupArgs();
    }

    public sealed class GetSecurityGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the security group you want to get
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetSecurityGroupInvokeArgs()
        {
        }
        public static new GetSecurityGroupInvokeArgs Empty => new GetSecurityGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityGroupResult
    {
        public readonly Outputs.ScalewayInstanceV1GetSecurityGroupResponse Items;

        [OutputConstructor]
        private GetSecurityGroupResult(Outputs.ScalewayInstanceV1GetSecurityGroupResponse items)
        {
            Items = items;
        }
    }
}
