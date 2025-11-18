package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

var ResourcesSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"cpu": {
			Description: lang.Markdown("Specifies the CPU required to run this task in MHz."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(100),
			},
		},
		"cores": {
			Description: lang.Markdown("Specifies the number of CPU cores to reserve specifically for the task. This may not be used with `cpu`. The behavior of setting `cores` is specific to each task driver (e.g. [docker](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/docker#cpu), [exec](https://developer.hashicorp.com/nomad/docs/deploy/task-driver/exec#cpu))."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(0),
			},
			IsOptional: true,
		},
		"memory": {
			Description: lang.Markdown("Specifies the memory required in MB."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(300),
			},
		},
		"memory_max": {
			Description: lang.Markdown("Optionally, specifies the maximum memory the task may use, if the client has excess memory capacity, in MB. See [Memory Oversubscription](https://developer.hashicorp.com/nomad/docs/job-specification/resources#memory-oversubscription) for more details."),
			DefaultValue: &schema.DefaultValue{
				Value: cty.NumberIntVal(300),
			},
			IsOptional: true,
		},
		// TODO: add numa, device, secrets
	},
}
