package cases

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func RenderBadtag(_buffer bytes.Buffer, w *gorazor.Widget) {
	if w.ErrorMsg != "" {

		_buffer.WriteString("<div class=\"form-group has-error\">\n	<div class=\"alert alert-danger\">")
		_buffer.WriteString(gorazor.HTMLEscape(w.ErrorMsg))
		_buffer.WriteString("</div>")
	} else {

		_buffer.WriteString("<div class=\"form-group\">")
	}
	_buffer.WriteString("\n\n	<label for=\"")
	_buffer.WriteString(gorazor.HTMLEscape(w.Name))
	_buffer.WriteString("\">")
	_buffer.WriteString(gorazor.HTMLEscape(w.Label))
	_buffer.WriteString("</label>\n	<input type=\"text\" name=\"")
	_buffer.WriteString(gorazor.HTMLEscape(w.Name))
	_buffer.WriteString("\" class=\"form-control\" id=\"")
	_buffer.WriteString(gorazor.HTMLEscape(w.Name))
	_buffer.WriteString("\" placeholder=\"")
	_buffer.WriteString(gorazor.HTMLEscape(w.PlaceHolder))
	_buffer.WriteString("\" value=\"")
	_buffer.WriteString(gorazor.HTMLEscape(w.Value))
	_buffer.WriteString("\">\n</div>")

}

func Badtag(w *gorazor.Widget) string {
	var _buffer bytes.Buffer
	RenderBadtag(_buffer, w)
	return _buffer.String()
}
