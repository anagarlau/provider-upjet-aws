package cloudfront

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for ebs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cloudfront_distribution", func(r *config.Resource) {
		r.UseAsync = true
		delete(r.References, "origin.domain_name")
	})

	// Setting the field as sensitive to be able to pass the content from a k8s secret
	p.AddResourceConfigurator("aws_cloudfront_function", func(r *config.Resource) {
		r.TerraformResource.Schema["code"].Sensitive = true
	})

	// Setting the field as sensitive to be able to pass the content from a k8s secret
	p.AddResourceConfigurator("aws_cloudfront_public_key", func(r *config.Resource) {
		r.TerraformResource.Schema["encoded_key"].Sensitive = true
	})

}