package main

import (
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Create an AWS EC2 instance
        ami, err := aws.LookupAmi(ctx, &aws.LookupAmiArgs{
            Filters: []aws.GetAmiFilter{
                {
                    Name:   "name",
                    Values: []string{"amzn2-ami-hvm-*"},
                },
            },
            Owners: []string{"amazon"},
            MostRecent: pulumi.Bool(true),
        })
        if err != nil {
            return err
        }

        instance, err := ec2.NewInstance(ctx, "radio-instance", &ec2.InstanceArgs{
            Ami:           pulumi.String(ami.Id),
            InstanceType:  pulumi.String("t2.micro"),
            Tags:          pulumi.StringMap{"Name": "RadioInstance"},
        })
        if err != nil {
            return err
        }

        ctx.Export("instanceId", instance.ID())
        return nil
    })
}