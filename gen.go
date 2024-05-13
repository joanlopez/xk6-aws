//go:build gen

package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const tmpl = `// Code generated by go generate; DO NOT EDIT.
package aws

import (
	"context"

	"github.com/dop251/goja"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

{{ range . }}
func (a *AWS) {{ .Name }}({{ .FunctionCall }}) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &{{.InputType}}{}
	if err := fromGojaObject(obj, in); err != nil {
		panic(err)
	}

	out, err := s3.NewFromConfig(cfg, func(o *s3.Options) {
        o.UsePathStyle = true  // Ensure path-style is used
    }).{{ .InnerFunctionCall }}
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}
{{ end }}
`

type S3Client interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	ListObjects(ctx context.Context, params *s3.ListObjectsInput, optFns ...func(*s3.Options)) (*s3.ListObjectsOutput, error)
}

func main() {
	t := reflect.TypeOf((*S3Client)(nil)).Elem()
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

	file, err := os.Create("s3_gen.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	if err := tmplParsed.Execute(file, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
	}
}
