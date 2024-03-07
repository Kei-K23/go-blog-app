// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package edit

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Kei-K23/go-blog-app/views/layout"
import "github.com/Kei-K23/go-blog-app/services"

func isCheck(isPublic bool) string {
	if isPublic {
		return "checked"
	}
	return ""
}

func ShowEditPage(blog services.Blog) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"https://cdn.tiny.cloud/1/5sg0q5e97evp2i5a97nt8q3i5zkjot0hxqhboodm81slvbk7/tinymce/6/tinymce.min.js\" referrerpolicy=\"origin\"></script> <!-- Place the following <script> and <textarea> tags your HTML's <body> --> <script>\n  tinymce.init({\n    selector: '#body',\n     plugins: 'image',\n    toolbar: 'undo redo | blocks | image | ' +\n        'bold italic backcolor | alignleft aligncenter ' +\n        'alignright alignjustify | bullist numlist outdent indent | ' +\n        'removeformat | help',\n    images_upload_url: \"/upload\",\n    relative_urls : false,\n    remove_script_host : false,\n    convert_urls : true,\n  })\n</script> <div class=\"container\"><div class=\"row justify-content-center\"><div class=\"col-md-6\"><h4 class=\"mt-3 text-center\">Edit your blog here</h4><form action=\"/blogs/edit\" class=\"mt-3\" method=\"post\"><div class=\"mb-3\"><label for=\"title\" class=\"form-label\">Title</label> <input value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(blog.Title))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" id=\"title\" type=\"text\" name=\"title\" class=\"form-control\" placeholder=\"Title\" required></div><div class=\"mb-3\"><label for=\"description\" class=\"form-label\">Description</label> <input value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(blog.Description))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" id=\"description\" type=\"text\" name=\"description\" class=\"form-control\" placeholder=\"Description\"></div><div class=\"mb-3\"><label for=\"body\" class=\"form-label\">Body</label> <textarea id=\"body\" class=\"form-control\" name=\"body\" rows=\"3\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(string(blog.Body))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/edit/show.templ`, Line: 46, Col: 102}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</textarea></div><div class=\"form-check\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if blog.IsPublic {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input name=\"isPublic\" class=\"form-check-input\" type=\"checkbox\" id=\"isPublic\" checked> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input name=\"isPublic\" class=\"form-check-input\" type=\"checkbox\" id=\"isPublic\"> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label class=\"form-check-label\" for=\"isPublic\">Public</label></div><input type=\"hidden\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(blog.ID))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"id\"> <button type=\"submit\" class=\"btn btn-primary\">Submit</button></form></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}