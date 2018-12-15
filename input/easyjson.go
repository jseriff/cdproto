// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package input

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput(in *jlexer.Lexer, out *TouchPoint) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "radiusX":
			out.RadiusX = float64(in.Float64())
		case "radiusY":
			out.RadiusY = float64(in.Float64())
		case "rotationAngle":
			out.RotationAngle = float64(in.Float64())
		case "force":
			out.Force = float64(in.Float64())
		case "id":
			out.ID = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput(out *jwriter.Writer, in TouchPoint) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.X))
	}
	{
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	if in.RadiusX != 0 {
		const prefix string = ",\"radiusX\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.RadiusX))
	}
	if in.RadiusY != 0 {
		const prefix string = ",\"radiusY\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.RadiusY))
	}
	if in.RotationAngle != 0 {
		const prefix string = ",\"rotationAngle\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.RotationAngle))
	}
	if in.Force != 0 {
		const prefix string = ",\"force\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Force))
	}
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TouchPoint) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TouchPoint) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TouchPoint) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TouchPoint) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput1(in *jlexer.Lexer, out *SynthesizeTapGestureParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "duration":
			out.Duration = int64(in.Int64())
		case "tapCount":
			out.TapCount = int64(in.Int64())
		case "gestureSourceType":
			(out.GestureSourceType).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput1(out *jwriter.Writer, in SynthesizeTapGestureParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.X))
	}
	{
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	if in.Duration != 0 {
		const prefix string = ",\"duration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Duration))
	}
	if in.TapCount != 0 {
		const prefix string = ",\"tapCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.TapCount))
	}
	if in.GestureSourceType != "" {
		const prefix string = ",\"gestureSourceType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.GestureSourceType).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SynthesizeTapGestureParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SynthesizeTapGestureParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SynthesizeTapGestureParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SynthesizeTapGestureParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput2(in *jlexer.Lexer, out *SynthesizeScrollGestureParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "xDistance":
			out.XDistance = float64(in.Float64())
		case "yDistance":
			out.YDistance = float64(in.Float64())
		case "xOverscroll":
			out.XOverscroll = float64(in.Float64())
		case "yOverscroll":
			out.YOverscroll = float64(in.Float64())
		case "preventFling":
			out.PreventFling = bool(in.Bool())
		case "speed":
			out.Speed = int64(in.Int64())
		case "gestureSourceType":
			(out.GestureSourceType).UnmarshalEasyJSON(in)
		case "repeatCount":
			out.RepeatCount = int64(in.Int64())
		case "repeatDelayMs":
			out.RepeatDelayMs = int64(in.Int64())
		case "interactionMarkerName":
			out.InteractionMarkerName = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput2(out *jwriter.Writer, in SynthesizeScrollGestureParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.X))
	}
	{
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	if in.XDistance != 0 {
		const prefix string = ",\"xDistance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.XDistance))
	}
	if in.YDistance != 0 {
		const prefix string = ",\"yDistance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.YDistance))
	}
	if in.XOverscroll != 0 {
		const prefix string = ",\"xOverscroll\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.XOverscroll))
	}
	if in.YOverscroll != 0 {
		const prefix string = ",\"yOverscroll\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.YOverscroll))
	}
	if in.PreventFling {
		const prefix string = ",\"preventFling\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.PreventFling))
	}
	if in.Speed != 0 {
		const prefix string = ",\"speed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Speed))
	}
	if in.GestureSourceType != "" {
		const prefix string = ",\"gestureSourceType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.GestureSourceType).MarshalEasyJSON(out)
	}
	if in.RepeatCount != 0 {
		const prefix string = ",\"repeatCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RepeatCount))
	}
	if in.RepeatDelayMs != 0 {
		const prefix string = ",\"repeatDelayMs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RepeatDelayMs))
	}
	if in.InteractionMarkerName != "" {
		const prefix string = ",\"interactionMarkerName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InteractionMarkerName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SynthesizeScrollGestureParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SynthesizeScrollGestureParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SynthesizeScrollGestureParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SynthesizeScrollGestureParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput3(in *jlexer.Lexer, out *SynthesizePinchGestureParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "scaleFactor":
			out.ScaleFactor = float64(in.Float64())
		case "relativeSpeed":
			out.RelativeSpeed = int64(in.Int64())
		case "gestureSourceType":
			(out.GestureSourceType).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput3(out *jwriter.Writer, in SynthesizePinchGestureParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.X))
	}
	{
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	{
		const prefix string = ",\"scaleFactor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.ScaleFactor))
	}
	if in.RelativeSpeed != 0 {
		const prefix string = ",\"relativeSpeed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RelativeSpeed))
	}
	if in.GestureSourceType != "" {
		const prefix string = ",\"gestureSourceType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.GestureSourceType).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SynthesizePinchGestureParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SynthesizePinchGestureParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SynthesizePinchGestureParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SynthesizePinchGestureParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput4(in *jlexer.Lexer, out *SetIgnoreInputEventsParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ignore":
			out.Ignore = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput4(out *jwriter.Writer, in SetIgnoreInputEventsParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ignore\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Ignore))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetIgnoreInputEventsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetIgnoreInputEventsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetIgnoreInputEventsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetIgnoreInputEventsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput5(in *jlexer.Lexer, out *InsertTextParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "text":
			out.Text = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput5(out *jwriter.Writer, in InsertTextParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InsertTextParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InsertTextParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InsertTextParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InsertTextParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput6(in *jlexer.Lexer, out *EmulateTouchFromMouseEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "x":
			out.X = int64(in.Int64())
		case "y":
			out.Y = int64(in.Int64())
		case "button":
			(out.Button).UnmarshalEasyJSON(in)
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(TimeSinceEpoch)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		case "deltaX":
			out.DeltaX = float64(in.Float64())
		case "deltaY":
			out.DeltaY = float64(in.Float64())
		case "modifiers":
			(out.Modifiers).UnmarshalEasyJSON(in)
		case "clickCount":
			out.ClickCount = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput6(out *jwriter.Writer, in EmulateTouchFromMouseEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Type).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.X))
	}
	{
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Y))
	}
	{
		const prefix string = ",\"button\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Button).MarshalEasyJSON(out)
	}
	if in.Timestamp != nil {
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Timestamp).MarshalEasyJSON(out)
	}
	if in.DeltaX != 0 {
		const prefix string = ",\"deltaX\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.DeltaX))
	}
	if in.DeltaY != 0 {
		const prefix string = ",\"deltaY\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.DeltaY))
	}
	{
		const prefix string = ",\"modifiers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Modifiers).MarshalEasyJSON(out)
	}
	if in.ClickCount != 0 {
		const prefix string = ",\"clickCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ClickCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EmulateTouchFromMouseEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EmulateTouchFromMouseEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EmulateTouchFromMouseEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EmulateTouchFromMouseEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput7(in *jlexer.Lexer, out *DispatchTouchEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "touchPoints":
			if in.IsNull() {
				in.Skip()
				out.TouchPoints = nil
			} else {
				in.Delim('[')
				if out.TouchPoints == nil {
					if !in.IsDelim(']') {
						out.TouchPoints = make([]*TouchPoint, 0, 8)
					} else {
						out.TouchPoints = []*TouchPoint{}
					}
				} else {
					out.TouchPoints = (out.TouchPoints)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *TouchPoint
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(TouchPoint)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.TouchPoints = append(out.TouchPoints, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "modifiers":
			(out.Modifiers).UnmarshalEasyJSON(in)
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(TimeSinceEpoch)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput7(out *jwriter.Writer, in DispatchTouchEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Type).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"touchPoints\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.TouchPoints == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.TouchPoints {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"modifiers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Modifiers).MarshalEasyJSON(out)
	}
	if in.Timestamp != nil {
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Timestamp).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DispatchTouchEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DispatchTouchEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DispatchTouchEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DispatchTouchEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput7(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput8(in *jlexer.Lexer, out *DispatchMouseEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "modifiers":
			(out.Modifiers).UnmarshalEasyJSON(in)
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(TimeSinceEpoch)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		case "button":
			(out.Button).UnmarshalEasyJSON(in)
		case "buttons":
			out.Buttons = int64(in.Int64())
		case "clickCount":
			out.ClickCount = int64(in.Int64())
		case "deltaX":
			out.DeltaX = float64(in.Float64())
		case "deltaY":
			out.DeltaY = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput8(out *jwriter.Writer, in DispatchMouseEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Type).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.X))
	}
	{
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	{
		const prefix string = ",\"modifiers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Modifiers).MarshalEasyJSON(out)
	}
	if in.Timestamp != nil {
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Timestamp).MarshalEasyJSON(out)
	}
	if in.Button != "" {
		const prefix string = ",\"button\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Button).MarshalEasyJSON(out)
	}
	if in.Buttons != 0 {
		const prefix string = ",\"buttons\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Buttons))
	}
	if in.ClickCount != 0 {
		const prefix string = ",\"clickCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ClickCount))
	}
	if in.DeltaX != 0 {
		const prefix string = ",\"deltaX\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.DeltaX))
	}
	if in.DeltaY != 0 {
		const prefix string = ",\"deltaY\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.DeltaY))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DispatchMouseEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DispatchMouseEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DispatchMouseEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DispatchMouseEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput8(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput9(in *jlexer.Lexer, out *DispatchKeyEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "modifiers":
			(out.Modifiers).UnmarshalEasyJSON(in)
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(TimeSinceEpoch)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		case "text":
			out.Text = string(in.String())
		case "unmodifiedText":
			out.UnmodifiedText = string(in.String())
		case "keyIdentifier":
			out.KeyIdentifier = string(in.String())
		case "code":
			out.Code = string(in.String())
		case "key":
			out.Key = string(in.String())
		case "windowsVirtualKeyCode":
			out.WindowsVirtualKeyCode = int64(in.Int64())
		case "nativeVirtualKeyCode":
			out.NativeVirtualKeyCode = int64(in.Int64())
		case "autoRepeat":
			out.AutoRepeat = bool(in.Bool())
		case "isKeypad":
			out.IsKeypad = bool(in.Bool())
		case "isSystemKey":
			out.IsSystemKey = bool(in.Bool())
		case "location":
			out.Location = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput9(out *jwriter.Writer, in DispatchKeyEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Type).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"modifiers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Modifiers).MarshalEasyJSON(out)
	}
	if in.Timestamp != nil {
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Timestamp).MarshalEasyJSON(out)
	}
	if in.Text != "" {
		const prefix string = ",\"text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	if in.UnmodifiedText != "" {
		const prefix string = ",\"unmodifiedText\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UnmodifiedText))
	}
	if in.KeyIdentifier != "" {
		const prefix string = ",\"keyIdentifier\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.KeyIdentifier))
	}
	if in.Code != "" {
		const prefix string = ",\"code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Code))
	}
	if in.Key != "" {
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Key))
	}
	if in.WindowsVirtualKeyCode != 0 {
		const prefix string = ",\"windowsVirtualKeyCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WindowsVirtualKeyCode))
	}
	if in.NativeVirtualKeyCode != 0 {
		const prefix string = ",\"nativeVirtualKeyCode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NativeVirtualKeyCode))
	}
	{
		const prefix string = ",\"autoRepeat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.AutoRepeat))
	}
	{
		const prefix string = ",\"isKeypad\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsKeypad))
	}
	{
		const prefix string = ",\"isSystemKey\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsSystemKey))
	}
	if in.Location != 0 {
		const prefix string = ",\"location\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Location))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DispatchKeyEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DispatchKeyEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoInput9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DispatchKeyEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DispatchKeyEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoInput9(l, v)
}
