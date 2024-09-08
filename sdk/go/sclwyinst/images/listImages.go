// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListImages(ctx *pulumi.Context, args *ListImagesArgs, opts ...pulumi.InvokeOption) (*ListImagesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListImagesResult
	err := ctx.Invoke("scaleway-instances:images:listImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListImagesArgs struct {
	// The zone you want to target
	Zone string `pulumi:"zone"`
}

type ListImagesResult struct {
	// List of images
	Images []ScalewayInstanceV1Image `pulumi:"images"`
}

func ListImagesOutput(ctx *pulumi.Context, args ListImagesOutputArgs, opts ...pulumi.InvokeOption) ListImagesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListImagesResult, error) {
			args := v.(ListImagesArgs)
			r, err := ListImages(ctx, &args, opts...)
			var s ListImagesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListImagesResultOutput)
}

type ListImagesOutputArgs struct {
	// The zone you want to target
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (ListImagesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListImagesArgs)(nil)).Elem()
}

type ListImagesResultOutput struct{ *pulumi.OutputState }

func (ListImagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListImagesResult)(nil)).Elem()
}

func (o ListImagesResultOutput) ToListImagesResultOutput() ListImagesResultOutput {
	return o
}

func (o ListImagesResultOutput) ToListImagesResultOutputWithContext(ctx context.Context) ListImagesResultOutput {
	return o
}

// List of images
func (o ListImagesResultOutput) Images() ScalewayInstanceV1ImageArrayOutput {
	return o.ApplyT(func(v ListImagesResult) []ScalewayInstanceV1Image { return v.Images }).(ScalewayInstanceV1ImageArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListImagesResultOutput{})
}
