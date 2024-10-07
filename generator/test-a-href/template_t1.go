// Code generated by t1 - DO NOT EDIT.

package testahref

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/senforsce/tndr"
import "context"
import "io"
import "bytes"

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
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<a href=\"javascript:alert(&#39;unaffected&#39;);\">Ignored</a> <a href=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var2 t1.SafeURL = t1.URL("javascript:alert('should be sanitized')")
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(string(t1_7745c5c3_Var2)))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\">Sanitized</a> <a href=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var3 t1.SafeURL = t1.SafeURL("javascript:alert('should not be sanitized')")
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(string(t1_7745c5c3_Var3)))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\">Unsanitized</a>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}
