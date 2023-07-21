// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package user_data

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ScalewayInstanceV1ListServerUserDataResponse struct {
	User_data []string `pulumi:"user_data"`
}

type ScalewayInstanceV1ListServerUserDataResponseOutput struct{ *pulumi.OutputState }

func (ScalewayInstanceV1ListServerUserDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayInstanceV1ListServerUserDataResponse)(nil)).Elem()
}

func (o ScalewayInstanceV1ListServerUserDataResponseOutput) ToScalewayInstanceV1ListServerUserDataResponseOutput() ScalewayInstanceV1ListServerUserDataResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListServerUserDataResponseOutput) ToScalewayInstanceV1ListServerUserDataResponseOutputWithContext(ctx context.Context) ScalewayInstanceV1ListServerUserDataResponseOutput {
	return o
}

func (o ScalewayInstanceV1ListServerUserDataResponseOutput) User_data() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScalewayInstanceV1ListServerUserDataResponse) []string { return v.User_data }).(pulumi.StringArrayOutput)
}

type ScalewayStdFile struct {
	Content      *string `pulumi:"content"`
	Content_type *string `pulumi:"content_type"`
	Name         *string `pulumi:"name"`
}

type ScalewayStdFileOutput struct{ *pulumi.OutputState }

func (ScalewayStdFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalewayStdFile)(nil)).Elem()
}

func (o ScalewayStdFileOutput) ToScalewayStdFileOutput() ScalewayStdFileOutput {
	return o
}

func (o ScalewayStdFileOutput) ToScalewayStdFileOutputWithContext(ctx context.Context) ScalewayStdFileOutput {
	return o
}

func (o ScalewayStdFileOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayStdFile) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o ScalewayStdFileOutput) Content_type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayStdFile) *string { return v.Content_type }).(pulumi.StringPtrOutput)
}

func (o ScalewayStdFileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalewayStdFile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalewayInstanceV1ListServerUserDataResponseOutput{})
	pulumi.RegisterOutputType(ScalewayStdFileOutput{})
}
