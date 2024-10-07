// Code generated by t1 - DO NOT EDIT.

package testcontext

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/senforsce/tndr"
import "context"
import "io"
import "bytes"

type contextKey string

var contextKeyName contextKey = "name"

func render() t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
		t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
		if !t1_7745c5c3_IsBuffer {
			t1_7745c5c3_Buffer = t1.GetBuffer()
			defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
		}
		ctx = t1.InitializeContext(ctx)
		t1_7745c5c3_Var1 := t1.GetChildren(ctx)
		if t1_7745c5c3_Var1 == nil {
			t1_7745c5c3_Var1 = t1.NopComponent
		}
		ctx = t1.ClearChildren(ctx)
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<ul><li>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var2 string
		t1_7745c5c3_Var2, t1_7745c5c3_Err = t1.JoinStringErrs(ctx.Value(contextKeyName).(string))
		if t1_7745c5c3_Err != nil {
			return t1.Error{Err: t1_7745c5c3_Err, FileName: `generator/test-context/template.t1`, Line: 8, Col: 42}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var2))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</li>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if ctx.Value(contextKeyName).(string) == "test" {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<li>the if passed</li>")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		}
		if ctx.Value(contextKeyName).(string) != "test" {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<li>the else if failed</li>")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		} else if ctx.Value(contextKeyName).(string) == "test" {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<li>the else if passed</li>")
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</ul>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}
