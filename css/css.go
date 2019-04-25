// Package css provides the Chrome DevTools Protocol
// commands, types, and events for the CSS domain.
//
// This domain exposes CSS read/write operations. All CSS objects
// (stylesheets, rules, and styles) have an associated id used in subsequent
// operations on the related object. Each object type has a specific id
// structure, and those are not interchangeable between objects of different
// kinds. CSS objects can be loaded using the get*ForNode() calls (which accept
// a DOM node id). A client can also keep track of stylesheets via the
// styleSheetAdded/styleSheetRemoved events and subsequently load the required
// stylesheet contents using the getStyleSheet[Text]() methods.
//
// Generated by the cdproto-gen command.
package css

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// AddRuleParams inserts a new rule with the given ruleText in a stylesheet
// with given styleSheetId, at the position specified by location.
type AddRuleParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // The css style sheet identifier where a new rule should be inserted.
	RuleText     string       `json:"ruleText"`     // The text of a new rule.
	Location     *SourceRange `json:"location"`     // Text position of a new rule in the target style sheet.
}

// AddRule inserts a new rule with the given ruleText in a stylesheet with
// given styleSheetId, at the position specified by location.
//
// parameters:
//   styleSheetID - The css style sheet identifier where a new rule should be inserted.
//   ruleText - The text of a new rule.
//   location - Text position of a new rule in the target style sheet.
func AddRule(styleSheetID StyleSheetID, ruleText string, location *SourceRange) *AddRuleParams {
	return &AddRuleParams{
		StyleSheetID: styleSheetID,
		RuleText:     ruleText,
		Location:     location,
	}
}

// AddRuleReturns return values.
type AddRuleReturns struct {
	Rule *Rule `json:"rule,omitempty"` // The newly created rule.
}

// Do executes CSS.addRule against the provided context.
//
// returns:
//   rule - The newly created rule.
func (p *AddRuleParams) Do(ctxt context.Context) (rule *Rule, err error) {
	// execute
	var res AddRuleReturns
	err = cdp.Execute(ctxt, CommandAddRule, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Rule, nil
}

// CollectClassNamesParams returns all class names from specified stylesheet.
type CollectClassNamesParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

// CollectClassNames returns all class names from specified stylesheet.
//
// parameters:
//   styleSheetID
func CollectClassNames(styleSheetID StyleSheetID) *CollectClassNamesParams {
	return &CollectClassNamesParams{
		StyleSheetID: styleSheetID,
	}
}

// CollectClassNamesReturns return values.
type CollectClassNamesReturns struct {
	ClassNames []string `json:"classNames,omitempty"` // Class name list.
}

// Do executes CSS.collectClassNames against the provided context.
//
// returns:
//   classNames - Class name list.
func (p *CollectClassNamesParams) Do(ctxt context.Context) (classNames []string, err error) {
	// execute
	var res CollectClassNamesReturns
	err = cdp.Execute(ctxt, CommandCollectClassNames, p, &res)
	if err != nil {
		return nil, err
	}

	return res.ClassNames, nil
}

// CreateStyleSheetParams creates a new special "via-inspector" stylesheet in
// the frame with given frameId.
type CreateStyleSheetParams struct {
	FrameID cdp.FrameID `json:"frameId"` // Identifier of the frame where "via-inspector" stylesheet should be created.
}

// CreateStyleSheet creates a new special "via-inspector" stylesheet in the
// frame with given frameId.
//
// parameters:
//   frameID - Identifier of the frame where "via-inspector" stylesheet should be created.
func CreateStyleSheet(frameID cdp.FrameID) *CreateStyleSheetParams {
	return &CreateStyleSheetParams{
		FrameID: frameID,
	}
}

// CreateStyleSheetReturns return values.
type CreateStyleSheetReturns struct {
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"` // Identifier of the created "via-inspector" stylesheet.
}

// Do executes CSS.createStyleSheet against the provided context.
//
// returns:
//   styleSheetID - Identifier of the created "via-inspector" stylesheet.
func (p *CreateStyleSheetParams) Do(ctxt context.Context) (styleSheetID StyleSheetID, err error) {
	// execute
	var res CreateStyleSheetReturns
	err = cdp.Execute(ctxt, CommandCreateStyleSheet, p, &res)
	if err != nil {
		return "", err
	}

	return res.StyleSheetID, nil
}

// DisableParams disables the CSS agent for the given page.
type DisableParams struct{}

// Disable disables the CSS agent for the given page.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes CSS.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables the CSS agent for the given page. Clients should not
// assume that the CSS agent has been enabled until the result of this command
// is received.
type EnableParams struct{}

// Enable enables the CSS agent for the given page. Clients should not assume
// that the CSS agent has been enabled until the result of this command is
// received.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes CSS.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandEnable, nil, nil)
}

// ForcePseudoStateParams ensures that the given node will have specified
// pseudo-classes whenever its style is computed by the browser.
type ForcePseudoStateParams struct {
	NodeID              cdp.NodeID `json:"nodeId"`              // The element id for which to force the pseudo state.
	ForcedPseudoClasses []string   `json:"forcedPseudoClasses"` // Element pseudo classes to force when computing the element's style.
}

// ForcePseudoState ensures that the given node will have specified
// pseudo-classes whenever its style is computed by the browser.
//
// parameters:
//   nodeID - The element id for which to force the pseudo state.
//   forcedPseudoClasses - Element pseudo classes to force when computing the element's style.
func ForcePseudoState(nodeID cdp.NodeID, forcedPseudoClasses []string) *ForcePseudoStateParams {
	return &ForcePseudoStateParams{
		NodeID:              nodeID,
		ForcedPseudoClasses: forcedPseudoClasses,
	}
}

// Do executes CSS.forcePseudoState against the provided context.
func (p *ForcePseudoStateParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandForcePseudoState, p, nil)
}

// GetBackgroundColorsParams [no description].
type GetBackgroundColorsParams struct {
	NodeID cdp.NodeID `json:"nodeId"` // Id of the node to get background colors for.
}

// GetBackgroundColors [no description].
//
// parameters:
//   nodeID - Id of the node to get background colors for.
func GetBackgroundColors(nodeID cdp.NodeID) *GetBackgroundColorsParams {
	return &GetBackgroundColorsParams{
		NodeID: nodeID,
	}
}

// GetBackgroundColorsReturns return values.
type GetBackgroundColorsReturns struct {
	BackgroundColors   []string `json:"backgroundColors,omitempty"`   // The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
	ComputedFontSize   string   `json:"computedFontSize,omitempty"`   // The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontWeight string   `json:"computedFontWeight,omitempty"` // The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
}

// Do executes CSS.getBackgroundColors against the provided context.
//
// returns:
//   backgroundColors - The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
//   computedFontSize - The computed font size for this node, as a CSS computed value string (e.g. '12px').
//   computedFontWeight - The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
func (p *GetBackgroundColorsParams) Do(ctxt context.Context) (backgroundColors []string, computedFontSize string, computedFontWeight string, err error) {
	// execute
	var res GetBackgroundColorsReturns
	err = cdp.Execute(ctxt, CommandGetBackgroundColors, p, &res)
	if err != nil {
		return nil, "", "", err
	}

	return res.BackgroundColors, res.ComputedFontSize, res.ComputedFontWeight, nil
}

// GetComputedStyleForNodeParams returns the computed style for a DOM node
// identified by nodeId.
type GetComputedStyleForNodeParams struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// GetComputedStyleForNode returns the computed style for a DOM node
// identified by nodeId.
//
// parameters:
//   nodeID
func GetComputedStyleForNode(nodeID cdp.NodeID) *GetComputedStyleForNodeParams {
	return &GetComputedStyleForNodeParams{
		NodeID: nodeID,
	}
}

// GetComputedStyleForNodeReturns return values.
type GetComputedStyleForNodeReturns struct {
	ComputedStyle []*ComputedProperty `json:"computedStyle,omitempty"` // Computed style for the specified DOM node.
}

// Do executes CSS.getComputedStyleForNode against the provided context.
//
// returns:
//   computedStyle - Computed style for the specified DOM node.
func (p *GetComputedStyleForNodeParams) Do(ctxt context.Context) (computedStyle []*ComputedProperty, err error) {
	// execute
	var res GetComputedStyleForNodeReturns
	err = cdp.Execute(ctxt, CommandGetComputedStyleForNode, p, &res)
	if err != nil {
		return nil, err
	}

	return res.ComputedStyle, nil
}

// GetInlineStylesForNodeParams returns the styles defined inline (explicitly
// in the "style" attribute and implicitly, using DOM attributes) for a DOM node
// identified by nodeId.
type GetInlineStylesForNodeParams struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// GetInlineStylesForNode returns the styles defined inline (explicitly in
// the "style" attribute and implicitly, using DOM attributes) for a DOM node
// identified by nodeId.
//
// parameters:
//   nodeID
func GetInlineStylesForNode(nodeID cdp.NodeID) *GetInlineStylesForNodeParams {
	return &GetInlineStylesForNodeParams{
		NodeID: nodeID,
	}
}

// GetInlineStylesForNodeReturns return values.
type GetInlineStylesForNodeReturns struct {
	InlineStyle     *Style `json:"inlineStyle,omitempty"`     // Inline style for the specified DOM node.
	AttributesStyle *Style `json:"attributesStyle,omitempty"` // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
}

// Do executes CSS.getInlineStylesForNode against the provided context.
//
// returns:
//   inlineStyle - Inline style for the specified DOM node.
//   attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (p *GetInlineStylesForNodeParams) Do(ctxt context.Context) (inlineStyle *Style, attributesStyle *Style, err error) {
	// execute
	var res GetInlineStylesForNodeReturns
	err = cdp.Execute(ctxt, CommandGetInlineStylesForNode, p, &res)
	if err != nil {
		return nil, nil, err
	}

	return res.InlineStyle, res.AttributesStyle, nil
}

// GetMatchedStylesForNodeParams returns requested styles for a DOM node
// identified by nodeId.
type GetMatchedStylesForNodeParams struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// GetMatchedStylesForNode returns requested styles for a DOM node identified
// by nodeId.
//
// parameters:
//   nodeID
func GetMatchedStylesForNode(nodeID cdp.NodeID) *GetMatchedStylesForNodeParams {
	return &GetMatchedStylesForNodeParams{
		NodeID: nodeID,
	}
}

// GetMatchedStylesForNodeReturns return values.
type GetMatchedStylesForNodeReturns struct {
	InlineStyle       *Style                  `json:"inlineStyle,omitempty"`       // Inline style for the specified DOM node.
	AttributesStyle   *Style                  `json:"attributesStyle,omitempty"`   // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	MatchedCSSRules   []*RuleMatch            `json:"matchedCSSRules,omitempty"`   // CSS rules matching this node, from all applicable stylesheets.
	PseudoElements    []*PseudoElementMatches `json:"pseudoElements,omitempty"`    // Pseudo style matches for this node.
	Inherited         []*InheritedStyleEntry  `json:"inherited,omitempty"`         // A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	CSSKeyframesRules []*KeyframesRule        `json:"cssKeyframesRules,omitempty"` // A list of CSS keyframed animations matching this node.
}

// Do executes CSS.getMatchedStylesForNode against the provided context.
//
// returns:
//   inlineStyle - Inline style for the specified DOM node.
//   attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
//   matchedCSSRules - CSS rules matching this node, from all applicable stylesheets.
//   pseudoElements - Pseudo style matches for this node.
//   inherited - A chain of inherited styles (from the immediate node parent up to the DOM tree root).
//   cssKeyframesRules - A list of CSS keyframed animations matching this node.
func (p *GetMatchedStylesForNodeParams) Do(ctxt context.Context) (inlineStyle *Style, attributesStyle *Style, matchedCSSRules []*RuleMatch, pseudoElements []*PseudoElementMatches, inherited []*InheritedStyleEntry, cssKeyframesRules []*KeyframesRule, err error) {
	// execute
	var res GetMatchedStylesForNodeReturns
	err = cdp.Execute(ctxt, CommandGetMatchedStylesForNode, p, &res)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return res.InlineStyle, res.AttributesStyle, res.MatchedCSSRules, res.PseudoElements, res.Inherited, res.CSSKeyframesRules, nil
}

// GetMediaQueriesParams returns all media queries parsed by the rendering
// engine.
type GetMediaQueriesParams struct{}

// GetMediaQueries returns all media queries parsed by the rendering engine.
func GetMediaQueries() *GetMediaQueriesParams {
	return &GetMediaQueriesParams{}
}

// GetMediaQueriesReturns return values.
type GetMediaQueriesReturns struct {
	Medias []*Media `json:"medias,omitempty"`
}

// Do executes CSS.getMediaQueries against the provided context.
//
// returns:
//   medias
func (p *GetMediaQueriesParams) Do(ctxt context.Context) (medias []*Media, err error) {
	// execute
	var res GetMediaQueriesReturns
	err = cdp.Execute(ctxt, CommandGetMediaQueries, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Medias, nil
}

// GetPlatformFontsForNodeParams requests information about platform fonts
// which we used to render child TextNodes in the given node.
type GetPlatformFontsForNodeParams struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// GetPlatformFontsForNode requests information about platform fonts which we
// used to render child TextNodes in the given node.
//
// parameters:
//   nodeID
func GetPlatformFontsForNode(nodeID cdp.NodeID) *GetPlatformFontsForNodeParams {
	return &GetPlatformFontsForNodeParams{
		NodeID: nodeID,
	}
}

// GetPlatformFontsForNodeReturns return values.
type GetPlatformFontsForNodeReturns struct {
	Fonts []*PlatformFontUsage `json:"fonts,omitempty"` // Usage statistics for every employed platform font.
}

// Do executes CSS.getPlatformFontsForNode against the provided context.
//
// returns:
//   fonts - Usage statistics for every employed platform font.
func (p *GetPlatformFontsForNodeParams) Do(ctxt context.Context) (fonts []*PlatformFontUsage, err error) {
	// execute
	var res GetPlatformFontsForNodeReturns
	err = cdp.Execute(ctxt, CommandGetPlatformFontsForNode, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Fonts, nil
}

// GetStyleSheetTextParams returns the current textual content for a
// stylesheet.
type GetStyleSheetTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

// GetStyleSheetText returns the current textual content for a stylesheet.
//
// parameters:
//   styleSheetID
func GetStyleSheetText(styleSheetID StyleSheetID) *GetStyleSheetTextParams {
	return &GetStyleSheetTextParams{
		StyleSheetID: styleSheetID,
	}
}

// GetStyleSheetTextReturns return values.
type GetStyleSheetTextReturns struct {
	Text string `json:"text,omitempty"` // The stylesheet text.
}

// Do executes CSS.getStyleSheetText against the provided context.
//
// returns:
//   text - The stylesheet text.
func (p *GetStyleSheetTextParams) Do(ctxt context.Context) (text string, err error) {
	// execute
	var res GetStyleSheetTextReturns
	err = cdp.Execute(ctxt, CommandGetStyleSheetText, p, &res)
	if err != nil {
		return "", err
	}

	return res.Text, nil
}

// SetEffectivePropertyValueForNodeParams find a rule with the given active
// property for the given node and set the new value for this property.
type SetEffectivePropertyValueForNodeParams struct {
	NodeID       cdp.NodeID `json:"nodeId"` // The element id for which to set property.
	PropertyName string     `json:"propertyName"`
	Value        string     `json:"value"`
}

// SetEffectivePropertyValueForNode find a rule with the given active
// property for the given node and set the new value for this property.
//
// parameters:
//   nodeID - The element id for which to set property.
//   propertyName
//   value
func SetEffectivePropertyValueForNode(nodeID cdp.NodeID, propertyName string, value string) *SetEffectivePropertyValueForNodeParams {
	return &SetEffectivePropertyValueForNodeParams{
		NodeID:       nodeID,
		PropertyName: propertyName,
		Value:        value,
	}
}

// Do executes CSS.setEffectivePropertyValueForNode against the provided context.
func (p *SetEffectivePropertyValueForNodeParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandSetEffectivePropertyValueForNode, p, nil)
}

// SetKeyframeKeyParams modifies the keyframe rule key text.
type SetKeyframeKeyParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	KeyText      string       `json:"keyText"`
}

// SetKeyframeKey modifies the keyframe rule key text.
//
// parameters:
//   styleSheetID
//   range
//   keyText
func SetKeyframeKey(styleSheetID StyleSheetID, rangeVal *SourceRange, keyText string) *SetKeyframeKeyParams {
	return &SetKeyframeKeyParams{
		StyleSheetID: styleSheetID,
		Range:        rangeVal,
		KeyText:      keyText,
	}
}

// SetKeyframeKeyReturns return values.
type SetKeyframeKeyReturns struct {
	KeyText *Value `json:"keyText,omitempty"` // The resulting key text after modification.
}

// Do executes CSS.setKeyframeKey against the provided context.
//
// returns:
//   keyText - The resulting key text after modification.
func (p *SetKeyframeKeyParams) Do(ctxt context.Context) (keyText *Value, err error) {
	// execute
	var res SetKeyframeKeyReturns
	err = cdp.Execute(ctxt, CommandSetKeyframeKey, p, &res)
	if err != nil {
		return nil, err
	}

	return res.KeyText, nil
}

// SetMediaTextParams modifies the rule selector.
type SetMediaTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Text         string       `json:"text"`
}

// SetMediaText modifies the rule selector.
//
// parameters:
//   styleSheetID
//   range
//   text
func SetMediaText(styleSheetID StyleSheetID, rangeVal *SourceRange, text string) *SetMediaTextParams {
	return &SetMediaTextParams{
		StyleSheetID: styleSheetID,
		Range:        rangeVal,
		Text:         text,
	}
}

// SetMediaTextReturns return values.
type SetMediaTextReturns struct {
	Media *Media `json:"media,omitempty"` // The resulting CSS media rule after modification.
}

// Do executes CSS.setMediaText against the provided context.
//
// returns:
//   media - The resulting CSS media rule after modification.
func (p *SetMediaTextParams) Do(ctxt context.Context) (media *Media, err error) {
	// execute
	var res SetMediaTextReturns
	err = cdp.Execute(ctxt, CommandSetMediaText, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Media, nil
}

// SetRuleSelectorParams modifies the rule selector.
type SetRuleSelectorParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Selector     string       `json:"selector"`
}

// SetRuleSelector modifies the rule selector.
//
// parameters:
//   styleSheetID
//   range
//   selector
func SetRuleSelector(styleSheetID StyleSheetID, rangeVal *SourceRange, selector string) *SetRuleSelectorParams {
	return &SetRuleSelectorParams{
		StyleSheetID: styleSheetID,
		Range:        rangeVal,
		Selector:     selector,
	}
}

// SetRuleSelectorReturns return values.
type SetRuleSelectorReturns struct {
	SelectorList *SelectorList `json:"selectorList,omitempty"` // The resulting selector list after modification.
}

// Do executes CSS.setRuleSelector against the provided context.
//
// returns:
//   selectorList - The resulting selector list after modification.
func (p *SetRuleSelectorParams) Do(ctxt context.Context) (selectorList *SelectorList, err error) {
	// execute
	var res SetRuleSelectorReturns
	err = cdp.Execute(ctxt, CommandSetRuleSelector, p, &res)
	if err != nil {
		return nil, err
	}

	return res.SelectorList, nil
}

// SetStyleSheetTextParams sets the new stylesheet text.
type SetStyleSheetTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Text         string       `json:"text"`
}

// SetStyleSheetText sets the new stylesheet text.
//
// parameters:
//   styleSheetID
//   text
func SetStyleSheetText(styleSheetID StyleSheetID, text string) *SetStyleSheetTextParams {
	return &SetStyleSheetTextParams{
		StyleSheetID: styleSheetID,
		Text:         text,
	}
}

// SetStyleSheetTextReturns return values.
type SetStyleSheetTextReturns struct {
	SourceMapURL string `json:"sourceMapURL,omitempty"` // URL of source map associated with script (if any).
}

// Do executes CSS.setStyleSheetText against the provided context.
//
// returns:
//   sourceMapURL - URL of source map associated with script (if any).
func (p *SetStyleSheetTextParams) Do(ctxt context.Context) (sourceMapURL string, err error) {
	// execute
	var res SetStyleSheetTextReturns
	err = cdp.Execute(ctxt, CommandSetStyleSheetText, p, &res)
	if err != nil {
		return "", err
	}

	return res.SourceMapURL, nil
}

// SetStyleTextsParams applies specified style edits one after another in the
// given order.
type SetStyleTextsParams struct {
	Edits []*StyleDeclarationEdit `json:"edits"`
}

// SetStyleTexts applies specified style edits one after another in the given
// order.
//
// parameters:
//   edits
func SetStyleTexts(edits []*StyleDeclarationEdit) *SetStyleTextsParams {
	return &SetStyleTextsParams{
		Edits: edits,
	}
}

// SetStyleTextsReturns return values.
type SetStyleTextsReturns struct {
	Styles []*Style `json:"styles,omitempty"` // The resulting styles after modification.
}

// Do executes CSS.setStyleTexts against the provided context.
//
// returns:
//   styles - The resulting styles after modification.
func (p *SetStyleTextsParams) Do(ctxt context.Context) (styles []*Style, err error) {
	// execute
	var res SetStyleTextsReturns
	err = cdp.Execute(ctxt, CommandSetStyleTexts, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Styles, nil
}

// StartRuleUsageTrackingParams enables the selector recording.
type StartRuleUsageTrackingParams struct{}

// StartRuleUsageTracking enables the selector recording.
func StartRuleUsageTracking() *StartRuleUsageTrackingParams {
	return &StartRuleUsageTrackingParams{}
}

// Do executes CSS.startRuleUsageTracking against the provided context.
func (p *StartRuleUsageTrackingParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandStartRuleUsageTracking, nil, nil)
}

// StopRuleUsageTrackingParams stop tracking rule usage and return the list
// of rules that were used since last call to takeCoverageDelta (or since start
// of coverage instrumentation).
type StopRuleUsageTrackingParams struct{}

// StopRuleUsageTracking stop tracking rule usage and return the list of
// rules that were used since last call to takeCoverageDelta (or since start of
// coverage instrumentation).
func StopRuleUsageTracking() *StopRuleUsageTrackingParams {
	return &StopRuleUsageTrackingParams{}
}

// StopRuleUsageTrackingReturns return values.
type StopRuleUsageTrackingReturns struct {
	RuleUsage []*RuleUsage `json:"ruleUsage,omitempty"`
}

// Do executes CSS.stopRuleUsageTracking against the provided context.
//
// returns:
//   ruleUsage
func (p *StopRuleUsageTrackingParams) Do(ctxt context.Context) (ruleUsage []*RuleUsage, err error) {
	// execute
	var res StopRuleUsageTrackingReturns
	err = cdp.Execute(ctxt, CommandStopRuleUsageTracking, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.RuleUsage, nil
}

// TakeCoverageDeltaParams obtain list of rules that became used since last
// call to this method (or since start of coverage instrumentation).
type TakeCoverageDeltaParams struct{}

// TakeCoverageDelta obtain list of rules that became used since last call to
// this method (or since start of coverage instrumentation).
func TakeCoverageDelta() *TakeCoverageDeltaParams {
	return &TakeCoverageDeltaParams{}
}

// TakeCoverageDeltaReturns return values.
type TakeCoverageDeltaReturns struct {
	Coverage []*RuleUsage `json:"coverage,omitempty"`
}

// Do executes CSS.takeCoverageDelta against the provided context.
//
// returns:
//   coverage
func (p *TakeCoverageDeltaParams) Do(ctxt context.Context) (coverage []*RuleUsage, err error) {
	// execute
	var res TakeCoverageDeltaReturns
	err = cdp.Execute(ctxt, CommandTakeCoverageDelta, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Coverage, nil
}

// Command names.
const (
	CommandAddRule                          = "CSS.addRule"
	CommandCollectClassNames                = "CSS.collectClassNames"
	CommandCreateStyleSheet                 = "CSS.createStyleSheet"
	CommandDisable                          = "CSS.disable"
	CommandEnable                           = "CSS.enable"
	CommandForcePseudoState                 = "CSS.forcePseudoState"
	CommandGetBackgroundColors              = "CSS.getBackgroundColors"
	CommandGetComputedStyleForNode          = "CSS.getComputedStyleForNode"
	CommandGetInlineStylesForNode           = "CSS.getInlineStylesForNode"
	CommandGetMatchedStylesForNode          = "CSS.getMatchedStylesForNode"
	CommandGetMediaQueries                  = "CSS.getMediaQueries"
	CommandGetPlatformFontsForNode          = "CSS.getPlatformFontsForNode"
	CommandGetStyleSheetText                = "CSS.getStyleSheetText"
	CommandSetEffectivePropertyValueForNode = "CSS.setEffectivePropertyValueForNode"
	CommandSetKeyframeKey                   = "CSS.setKeyframeKey"
	CommandSetMediaText                     = "CSS.setMediaText"
	CommandSetRuleSelector                  = "CSS.setRuleSelector"
	CommandSetStyleSheetText                = "CSS.setStyleSheetText"
	CommandSetStyleTexts                    = "CSS.setStyleTexts"
	CommandStartRuleUsageTracking           = "CSS.startRuleUsageTracking"
	CommandStopRuleUsageTracking            = "CSS.stopRuleUsageTracking"
	CommandTakeCoverageDelta                = "CSS.takeCoverageDelta"
)
