// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.ScalewayInstances.Rules
{
    [EnumType]
    public readonly struct Action : IEquatable<Action>
    {
        private readonly string _value;

        private Action(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Action Accept { get; } = new Action("accept");
        public static Action Drop { get; } = new Action("drop");

        public static bool operator ==(Action left, Action right) => left.Equals(right);
        public static bool operator !=(Action left, Action right) => !left.Equals(right);

        public static explicit operator string(Action value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Action other && Equals(other);
        public bool Equals(Action other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct Direction : IEquatable<Direction>
    {
        private readonly string _value;

        private Direction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Direction Inbound { get; } = new Direction("inbound");
        public static Direction Outbound { get; } = new Direction("outbound");

        public static bool operator ==(Direction left, Direction right) => left.Equals(right);
        public static bool operator !=(Direction left, Direction right) => !left.Equals(right);

        public static explicit operator string(Direction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Direction other && Equals(other);
        public bool Equals(Direction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct Protocol : IEquatable<Protocol>
    {
        private readonly string _value;

        private Protocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Protocol Tcp { get; } = new Protocol("TCP");
        public static Protocol Udp { get; } = new Protocol("UDP");
        public static Protocol Icmp { get; } = new Protocol("ICMP");
        public static Protocol Any { get; } = new Protocol("ANY");

        public static bool operator ==(Protocol left, Protocol right) => left.Equals(right);
        public static bool operator !=(Protocol left, Protocol right) => !left.Equals(right);

        public static explicit operator string(Protocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Protocol other && Equals(other);
        public bool Equals(Protocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

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

    /// <summary>
    /// Action to apply when the rule matches a packet
    /// </summary>
    [EnumType]
    public readonly struct ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction : IEquatable<ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction>
    {
        private readonly string _value;

        private ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction Accept { get; } = new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction("accept");
        public static ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction Drop { get; } = new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction("drop");

        public static bool operator ==(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction left, ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction left, ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction other && Equals(other);
        public bool Equals(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Direction the rule applies to
    /// </summary>
    [EnumType]
    public readonly struct ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection : IEquatable<ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection>
    {
        private readonly string _value;

        private ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection Inbound { get; } = new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection("inbound");
        public static ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection Outbound { get; } = new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection("outbound");

        public static bool operator ==(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection left, ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection left, ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection other && Equals(other);
        public bool Equals(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Protocol family this rule applies to
    /// </summary>
    [EnumType]
    public readonly struct ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol : IEquatable<ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol>
    {
        private readonly string _value;

        private ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol Tcp { get; } = new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol("TCP");
        public static ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol Udp { get; } = new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol("UDP");
        public static ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol Icmp { get; } = new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol("ICMP");
        public static ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol Any { get; } = new ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol("ANY");

        public static bool operator ==(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol left, ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol left, ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol other && Equals(other);
        public bool Equals(ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}