// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.ScalewayInstances.Rules
{
    [EnumType]
    public readonly struct ScalewayInstanceV1SecurityGroupRuleAction : IEquatable<ScalewayInstanceV1SecurityGroupRuleAction>
    {
        private readonly string _value;

        private ScalewayInstanceV1SecurityGroupRuleAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1SecurityGroupRuleAction Accept { get; } = new ScalewayInstanceV1SecurityGroupRuleAction("accept");
        public static ScalewayInstanceV1SecurityGroupRuleAction Drop { get; } = new ScalewayInstanceV1SecurityGroupRuleAction("drop");

        public static bool operator ==(ScalewayInstanceV1SecurityGroupRuleAction left, ScalewayInstanceV1SecurityGroupRuleAction right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1SecurityGroupRuleAction left, ScalewayInstanceV1SecurityGroupRuleAction right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1SecurityGroupRuleAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1SecurityGroupRuleAction other && Equals(other);
        public bool Equals(ScalewayInstanceV1SecurityGroupRuleAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScalewayInstanceV1SecurityGroupRuleDirection : IEquatable<ScalewayInstanceV1SecurityGroupRuleDirection>
    {
        private readonly string _value;

        private ScalewayInstanceV1SecurityGroupRuleDirection(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1SecurityGroupRuleDirection Inbound { get; } = new ScalewayInstanceV1SecurityGroupRuleDirection("inbound");
        public static ScalewayInstanceV1SecurityGroupRuleDirection Outbound { get; } = new ScalewayInstanceV1SecurityGroupRuleDirection("outbound");

        public static bool operator ==(ScalewayInstanceV1SecurityGroupRuleDirection left, ScalewayInstanceV1SecurityGroupRuleDirection right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1SecurityGroupRuleDirection left, ScalewayInstanceV1SecurityGroupRuleDirection right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1SecurityGroupRuleDirection value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1SecurityGroupRuleDirection other && Equals(other);
        public bool Equals(ScalewayInstanceV1SecurityGroupRuleDirection other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScalewayInstanceV1SecurityGroupRuleProtocol : IEquatable<ScalewayInstanceV1SecurityGroupRuleProtocol>
    {
        private readonly string _value;

        private ScalewayInstanceV1SecurityGroupRuleProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1SecurityGroupRuleProtocol Tcp { get; } = new ScalewayInstanceV1SecurityGroupRuleProtocol("TCP");
        public static ScalewayInstanceV1SecurityGroupRuleProtocol Udp { get; } = new ScalewayInstanceV1SecurityGroupRuleProtocol("UDP");
        public static ScalewayInstanceV1SecurityGroupRuleProtocol Icmp { get; } = new ScalewayInstanceV1SecurityGroupRuleProtocol("ICMP");
        public static ScalewayInstanceV1SecurityGroupRuleProtocol Any { get; } = new ScalewayInstanceV1SecurityGroupRuleProtocol("ANY");

        public static bool operator ==(ScalewayInstanceV1SecurityGroupRuleProtocol left, ScalewayInstanceV1SecurityGroupRuleProtocol right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1SecurityGroupRuleProtocol left, ScalewayInstanceV1SecurityGroupRuleProtocol right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1SecurityGroupRuleProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1SecurityGroupRuleProtocol other && Equals(other);
        public bool Equals(ScalewayInstanceV1SecurityGroupRuleProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct SecurityGroupRuleAction : IEquatable<SecurityGroupRuleAction>
    {
        private readonly string _value;

        private SecurityGroupRuleAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecurityGroupRuleAction Accept { get; } = new SecurityGroupRuleAction("accept");
        public static SecurityGroupRuleAction Drop { get; } = new SecurityGroupRuleAction("drop");

        public static bool operator ==(SecurityGroupRuleAction left, SecurityGroupRuleAction right) => left.Equals(right);
        public static bool operator !=(SecurityGroupRuleAction left, SecurityGroupRuleAction right) => !left.Equals(right);

        public static explicit operator string(SecurityGroupRuleAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecurityGroupRuleAction other && Equals(other);
        public bool Equals(SecurityGroupRuleAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct SecurityGroupRuleDirection : IEquatable<SecurityGroupRuleDirection>
    {
        private readonly string _value;

        private SecurityGroupRuleDirection(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecurityGroupRuleDirection Inbound { get; } = new SecurityGroupRuleDirection("inbound");
        public static SecurityGroupRuleDirection Outbound { get; } = new SecurityGroupRuleDirection("outbound");

        public static bool operator ==(SecurityGroupRuleDirection left, SecurityGroupRuleDirection right) => left.Equals(right);
        public static bool operator !=(SecurityGroupRuleDirection left, SecurityGroupRuleDirection right) => !left.Equals(right);

        public static explicit operator string(SecurityGroupRuleDirection value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecurityGroupRuleDirection other && Equals(other);
        public bool Equals(SecurityGroupRuleDirection other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct SecurityGroupRuleProtocol : IEquatable<SecurityGroupRuleProtocol>
    {
        private readonly string _value;

        private SecurityGroupRuleProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecurityGroupRuleProtocol Tcp { get; } = new SecurityGroupRuleProtocol("TCP");
        public static SecurityGroupRuleProtocol Udp { get; } = new SecurityGroupRuleProtocol("UDP");
        public static SecurityGroupRuleProtocol Icmp { get; } = new SecurityGroupRuleProtocol("ICMP");
        public static SecurityGroupRuleProtocol Any { get; } = new SecurityGroupRuleProtocol("ANY");

        public static bool operator ==(SecurityGroupRuleProtocol left, SecurityGroupRuleProtocol right) => left.Equals(right);
        public static bool operator !=(SecurityGroupRuleProtocol left, SecurityGroupRuleProtocol right) => !left.Equals(right);

        public static explicit operator string(SecurityGroupRuleProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecurityGroupRuleProtocol other && Equals(other);
        public bool Equals(SecurityGroupRuleProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
