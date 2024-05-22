//go:build gen

package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

const tmpl = `// Code generated by go generate; DO NOT EDIT.
package aws

import (
	"context"

	"github.com/dop251/goja"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

{{ range . }}
func (a *AWS) {{ .Name }}({{ .FunctionCall }}) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &{{.InputType}}{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := eventbridge.NewFromConfig(cfg).{{ .InnerFunctionCall }}
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}
{{ end }}
`

type EventBridgeClient interface {
	ActivateEventSource(context.Context, *eventbridge.ActivateEventSourceInput, ...func(*eventbridge.Options)) (*eventbridge.ActivateEventSourceOutput, error)
	CancelReplay(context.Context, *eventbridge.CancelReplayInput, ...func(*eventbridge.Options)) (*eventbridge.CancelReplayOutput, error)
	CreateApiDestination(context.Context, *eventbridge.CreateApiDestinationInput, ...func(*eventbridge.Options)) (*eventbridge.CreateApiDestinationOutput, error)
	CreateArchive(context.Context, *eventbridge.CreateArchiveInput, ...func(*eventbridge.Options)) (*eventbridge.CreateArchiveOutput, error)
	CreateConnection(context.Context, *eventbridge.CreateConnectionInput, ...func(*eventbridge.Options)) (*eventbridge.CreateConnectionOutput, error)
	CreateEndpoint(context.Context, *eventbridge.CreateEndpointInput, ...func(*eventbridge.Options)) (*eventbridge.CreateEndpointOutput, error)
	CreateEventBus(context.Context, *eventbridge.CreateEventBusInput, ...func(*eventbridge.Options)) (*eventbridge.CreateEventBusOutput, error)
	CreatePartnerEventSource(context.Context, *eventbridge.CreatePartnerEventSourceInput, ...func(*eventbridge.Options)) (*eventbridge.CreatePartnerEventSourceOutput, error)
	DeactivateEventSource(context.Context, *eventbridge.DeactivateEventSourceInput, ...func(*eventbridge.Options)) (*eventbridge.DeactivateEventSourceOutput, error)
	DeauthorizeConnection(context.Context, *eventbridge.DeauthorizeConnectionInput, ...func(*eventbridge.Options)) (*eventbridge.DeauthorizeConnectionOutput, error)
	DeleteApiDestination(context.Context, *eventbridge.DeleteApiDestinationInput, ...func(*eventbridge.Options)) (*eventbridge.DeleteApiDestinationOutput, error)
	DeleteArchive(context.Context, *eventbridge.DeleteArchiveInput, ...func(*eventbridge.Options)) (*eventbridge.DeleteArchiveOutput, error)
	DeleteConnection(context.Context, *eventbridge.DeleteConnectionInput, ...func(*eventbridge.Options)) (*eventbridge.DeleteConnectionOutput, error)
	DeleteEndpoint(context.Context, *eventbridge.DeleteEndpointInput, ...func(*eventbridge.Options)) (*eventbridge.DeleteEndpointOutput, error)
	DeleteEventBus(context.Context, *eventbridge.DeleteEventBusInput, ...func(*eventbridge.Options)) (*eventbridge.DeleteEventBusOutput, error)
	DeletePartnerEventSource(context.Context, *eventbridge.DeletePartnerEventSourceInput, ...func(*eventbridge.Options)) (*eventbridge.DeletePartnerEventSourceOutput, error)
	DeleteRule(context.Context, *eventbridge.DeleteRuleInput, ...func(*eventbridge.Options)) (*eventbridge.DeleteRuleOutput, error)
	DescribeApiDestination(context.Context, *eventbridge.DescribeApiDestinationInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeApiDestinationOutput, error)
	DescribeArchive(context.Context, *eventbridge.DescribeArchiveInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeArchiveOutput, error)
	DescribeConnection(context.Context, *eventbridge.DescribeConnectionInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeConnectionOutput, error)
	DescribeEndpoint(context.Context, *eventbridge.DescribeEndpointInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeEndpointOutput, error)
	DescribeEventBus(context.Context, *eventbridge.DescribeEventBusInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeEventBusOutput, error)
	DescribeEventSource(context.Context, *eventbridge.DescribeEventSourceInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeEventSourceOutput, error)
	DescribePartnerEventSource(context.Context, *eventbridge.DescribePartnerEventSourceInput, ...func(*eventbridge.Options)) (*eventbridge.DescribePartnerEventSourceOutput, error)
	DescribeReplay(context.Context, *eventbridge.DescribeReplayInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeReplayOutput, error)
	DescribeRule(context.Context, *eventbridge.DescribeRuleInput, ...func(*eventbridge.Options)) (*eventbridge.DescribeRuleOutput, error)
	DisableRule(context.Context, *eventbridge.DisableRuleInput, ...func(*eventbridge.Options)) (*eventbridge.DisableRuleOutput, error)
	EnableRule(context.Context, *eventbridge.EnableRuleInput, ...func(*eventbridge.Options)) (*eventbridge.EnableRuleOutput, error)
	ListApiDestinations(context.Context, *eventbridge.ListApiDestinationsInput, ...func(*eventbridge.Options)) (*eventbridge.ListApiDestinationsOutput, error)
	ListArchives(context.Context, *eventbridge.ListArchivesInput, ...func(*eventbridge.Options)) (*eventbridge.ListArchivesOutput, error)
	ListConnections(context.Context, *eventbridge.ListConnectionsInput, ...func(*eventbridge.Options)) (*eventbridge.ListConnectionsOutput, error)
	ListEndpoints(context.Context, *eventbridge.ListEndpointsInput, ...func(*eventbridge.Options)) (*eventbridge.ListEndpointsOutput, error)
	ListEventBuses(context.Context, *eventbridge.ListEventBusesInput, ...func(*eventbridge.Options)) (*eventbridge.ListEventBusesOutput, error)
	ListEventSources(context.Context, *eventbridge.ListEventSourcesInput, ...func(*eventbridge.Options)) (*eventbridge.ListEventSourcesOutput, error)
	ListPartnerEventSourceAccounts(context.Context, *eventbridge.ListPartnerEventSourceAccountsInput, ...func(*eventbridge.Options)) (*eventbridge.ListPartnerEventSourceAccountsOutput, error)
	ListPartnerEventSources(context.Context, *eventbridge.ListPartnerEventSourcesInput, ...func(*eventbridge.Options)) (*eventbridge.ListPartnerEventSourcesOutput, error)
	ListReplays(context.Context, *eventbridge.ListReplaysInput, ...func(*eventbridge.Options)) (*eventbridge.ListReplaysOutput, error)
	ListRuleNamesByTarget(context.Context, *eventbridge.ListRuleNamesByTargetInput, ...func(*eventbridge.Options)) (*eventbridge.ListRuleNamesByTargetOutput, error)
	ListRules(context.Context, *eventbridge.ListRulesInput, ...func(*eventbridge.Options)) (*eventbridge.ListRulesOutput, error)
	ListTagsForResource(context.Context, *eventbridge.ListTagsForResourceInput, ...func(*eventbridge.Options)) (*eventbridge.ListTagsForResourceOutput, error)
	ListTargetsByRule(context.Context, *eventbridge.ListTargetsByRuleInput, ...func(*eventbridge.Options)) (*eventbridge.ListTargetsByRuleOutput, error)
	PutEvents(context.Context, *eventbridge.PutEventsInput, ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)
	PutPartnerEvents(context.Context, *eventbridge.PutPartnerEventsInput, ...func(*eventbridge.Options)) (*eventbridge.PutPartnerEventsOutput, error)
	PutPermission(context.Context, *eventbridge.PutPermissionInput, ...func(*eventbridge.Options)) (*eventbridge.PutPermissionOutput, error)
	PutRule(context.Context, *eventbridge.PutRuleInput, ...func(*eventbridge.Options)) (*eventbridge.PutRuleOutput, error)
	PutTargets(context.Context, *eventbridge.PutTargetsInput, ...func(*eventbridge.Options)) (*eventbridge.PutTargetsOutput, error)
	RemovePermission(context.Context, *eventbridge.RemovePermissionInput, ...func(*eventbridge.Options)) (*eventbridge.RemovePermissionOutput, error)
	RemoveTargets(context.Context, *eventbridge.RemoveTargetsInput, ...func(*eventbridge.Options)) (*eventbridge.RemoveTargetsOutput, error)
	StartReplay(context.Context, *eventbridge.StartReplayInput, ...func(*eventbridge.Options)) (*eventbridge.StartReplayOutput, error)
	TagResource(context.Context, *eventbridge.TagResourceInput, ...func(*eventbridge.Options)) (*eventbridge.TagResourceOutput, error)
	TestEventPattern(context.Context, *eventbridge.TestEventPatternInput, ...func(*eventbridge.Options)) (*eventbridge.TestEventPatternOutput, error)
	UntagResource(context.Context, *eventbridge.UntagResourceInput, ...func(*eventbridge.Options)) (*eventbridge.UntagResourceOutput, error)
	UpdateApiDestination(context.Context, *eventbridge.UpdateApiDestinationInput, ...func(*eventbridge.Options)) (*eventbridge.UpdateApiDestinationOutput, error)
	UpdateArchive(context.Context, *eventbridge.UpdateArchiveInput, ...func(*eventbridge.Options)) (*eventbridge.UpdateArchiveOutput, error)
	UpdateConnection(context.Context, *eventbridge.UpdateConnectionInput, ...func(*eventbridge.Options)) (*eventbridge.UpdateConnectionOutput, error)
	UpdateEndpoint(context.Context, *eventbridge.UpdateEndpointInput, ...func(*eventbridge.Options)) (*eventbridge.UpdateEndpointOutput, error)
	UpdateEventBus(context.Context, *eventbridge.UpdateEventBusInput, ...func(*eventbridge.Options)) (*eventbridge.UpdateEventBusOutput, error)
}

func main() {
	t := reflect.TypeOf((*EventBridgeClient)(nil)).Elem()
	data := make([]map[string]string, 0)

	for i := 0; i < t.NumMethod(); i++ {
		// Parsing the method signature
		var (
			method   = t.Method(i)
			params   []string
			callArgs []string

			functionCall      string
			inputType         string
			innerFunctionCall = fmt.Sprintf("%s(", method.Name)
		)

		for j := 0; j < method.Type.NumIn(); j++ {
			isTheLastArgument := j == method.Type.NumIn()-1

			p := method.Type.In(j)
			paramName := fmt.Sprintf("param%d", j)
			params = append(params, fmt.Sprintf("%s %s", paramName, p))
			if j > 0 { // skip receiver
				callArgs = append(callArgs, paramName)
			}

			switch {
			// Is context.Context
			case p.Implements(reflect.TypeOf((*context.Context)(nil)).Elem()):
				innerFunctionCall += "context.Background(), "
			// Variadic argument
			case method.Type.IsVariadic() && isTheLastArgument:
				// explicitly skip variadic parameters
			// Pointer to
			case p.Kind() == reflect.Ptr:
				functionCall += "obj *goja.Object,"
				inputType = strings.ReplaceAll(p.String(), "*", "")
				innerFunctionCall += "in, "
			}
		}
		innerFunctionCall += ")" // close the function call

		results := []string{}
		for j := 0; j < method.Type.NumOut(); j++ {
			r := method.Type.Out(j)
			results = append(results, r.String())
		}
		data = append(data, map[string]string{
			"Name":              method.Name,
			"Params":            fmt.Sprintf("%s", strings.Join(params, ", ")),
			"Results":           fmt.Sprintf("%s", strings.Join(results, ", ")),
			"CallArgs":          fmt.Sprintf("%s", strings.Join(callArgs, ", ")),
			"FunctionCall":      functionCall,
			"InnerFunctionCall": innerFunctionCall,
			"InputType":         inputType,
		})
	}

	tmplParsed, err := template.New("wrapper").Parse(tmpl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
		return
	}

	file, err := os.Create("eventbridge_gen.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	if err := tmplParsed.Execute(file, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
	}
}
