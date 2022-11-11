// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Rules.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1ListSecurityGroupRulesResponse
    {
        /// <summary>
        /// List of security rules
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalewayInstanceV1SecurityGroupRule> Rules;

        [OutputConstructor]
        private ScalewayInstanceV1ListSecurityGroupRulesResponse(ImmutableArray<Outputs.ScalewayInstanceV1SecurityGroupRule> rules)
        {
            Rules = rules;
        }
    }
}
