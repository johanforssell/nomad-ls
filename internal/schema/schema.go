package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
)

const (
	variablesLabel = "variables"
	variableLabel  = "variable"
	localsLabel    = "locals"
	vaultLabel     = "vault"
	taskLabel      = "task"
	secretLabel    = "secret"

	inputVariablesAccessor = "var"
	localsAccessor         = "local"
)

var SchemaMapBetter map[string]*hcl.BodySchema = map[string]*hcl.BodySchema{
	"root": RootBodySchema.Copy().ToHCLSchema(),

	"affinity":         AffinitySchema.Copy().ToHCLSchema(),
	"artifact":         ArtifactSchema.Copy().ToHCLSchema(),
	"constraint":       ConstraintSchema.Copy().ToHCLSchema(),
	"consul":           ConsulSchema.Copy().ToHCLSchema(),
	"dispatch_payload": DispatchPayloadSchema.Copy().ToHCLSchema(),
	"env":              EnvSchema.Copy().ToHCLSchema(),
	"ephemeral_disk":   EphemeralDiskSchema.Copy().ToHCLSchema(),
	"group":            GroupSchema.Copy().ToHCLSchema(),
	"identity":         IdentitySchema.Copy().ToHCLSchema(),
	"job":              JobSchemaBetter.Copy().ToHCLSchema(),
	"lifecycle":        LifecycleSchema.Copy().ToHCLSchema(),
	"logs":             LogsSchema.Copy().ToHCLSchema(),
	"meta":             MetaSchema.Copy().ToHCLSchema(),
	"migrate":          MigrateSchema.Copy().ToHCLSchema(),
	"parameterized":    ParameterizedSchema.Copy().ToHCLSchema(),
	"periodic":         PeriodicSchema.Copy().ToHCLSchema(),
	"reschedule":       RescheduleSchema.Copy().ToHCLSchema(),
	"resources":        ResourcesSchema.Copy().ToHCLSchema(),
	"restart":          RestartSchema.Copy().ToHCLSchema(),
	"spread":           SpreadSchema.Copy().ToHCLSchema(),
	"target":           TargetSchema.Copy().ToHCLSchema(),
	"task":             TaskSchema.Copy().ToHCLSchema(),
	"update":           UpdateSchema.Copy().ToHCLSchema(),
	"volume":           VolumeSchema.Copy().ToHCLSchema(),
}

var RootBodySchema = schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		variablesLabel: {},
		variableLabel:  {},
		localsLabel:    {},
	},
	Blocks: map[string]*schema.BlockSchema{
		"job": {
			Description: lang.Markdown("## h2\ntest"),
			Labels: []*schema.LabelSchema{
				{Name: "name"},
			},
			Body: JobSchemaBetter,
		},
	},
}
