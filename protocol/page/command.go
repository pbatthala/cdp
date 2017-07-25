// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"github.com/mafredri/cdp/protocol/debugger"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/emulation"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/runtime"
)

// AddScriptToEvaluateOnLoadArgs represents the arguments for AddScriptToEvaluateOnLoad in the Page domain.
type AddScriptToEvaluateOnLoadArgs struct {
	ScriptSource string `json:"scriptSource"` //
}

// NewAddScriptToEvaluateOnLoadArgs initializes AddScriptToEvaluateOnLoadArgs with the required arguments.
func NewAddScriptToEvaluateOnLoadArgs(scriptSource string) *AddScriptToEvaluateOnLoadArgs {
	args := new(AddScriptToEvaluateOnLoadArgs)
	args.ScriptSource = scriptSource
	return args
}

// AddScriptToEvaluateOnLoadReply represents the return values for AddScriptToEvaluateOnLoad in the Page domain.
type AddScriptToEvaluateOnLoadReply struct {
	Identifier ScriptIdentifier `json:"identifier"` // Identifier of the added script.
}

// RemoveScriptToEvaluateOnLoadArgs represents the arguments for RemoveScriptToEvaluateOnLoad in the Page domain.
type RemoveScriptToEvaluateOnLoadArgs struct {
	Identifier ScriptIdentifier `json:"identifier"` //
}

// NewRemoveScriptToEvaluateOnLoadArgs initializes RemoveScriptToEvaluateOnLoadArgs with the required arguments.
func NewRemoveScriptToEvaluateOnLoadArgs(identifier ScriptIdentifier) *RemoveScriptToEvaluateOnLoadArgs {
	args := new(RemoveScriptToEvaluateOnLoadArgs)
	args.Identifier = identifier
	return args
}

// AddScriptToEvaluateOnNewDocumentArgs represents the arguments for AddScriptToEvaluateOnNewDocument in the Page domain.
type AddScriptToEvaluateOnNewDocumentArgs struct {
	Source string `json:"source"` //
}

// NewAddScriptToEvaluateOnNewDocumentArgs initializes AddScriptToEvaluateOnNewDocumentArgs with the required arguments.
func NewAddScriptToEvaluateOnNewDocumentArgs(source string) *AddScriptToEvaluateOnNewDocumentArgs {
	args := new(AddScriptToEvaluateOnNewDocumentArgs)
	args.Source = source
	return args
}

// AddScriptToEvaluateOnNewDocumentReply represents the return values for AddScriptToEvaluateOnNewDocument in the Page domain.
type AddScriptToEvaluateOnNewDocumentReply struct {
	Identifier ScriptIdentifier `json:"identifier"` // Identifier of the added script.
}

// RemoveScriptToEvaluateOnNewDocumentArgs represents the arguments for RemoveScriptToEvaluateOnNewDocument in the Page domain.
type RemoveScriptToEvaluateOnNewDocumentArgs struct {
	Identifier ScriptIdentifier `json:"identifier"` //
}

// NewRemoveScriptToEvaluateOnNewDocumentArgs initializes RemoveScriptToEvaluateOnNewDocumentArgs with the required arguments.
func NewRemoveScriptToEvaluateOnNewDocumentArgs(identifier ScriptIdentifier) *RemoveScriptToEvaluateOnNewDocumentArgs {
	args := new(RemoveScriptToEvaluateOnNewDocumentArgs)
	args.Identifier = identifier
	return args
}

// SetAutoAttachToCreatedPagesArgs represents the arguments for SetAutoAttachToCreatedPages in the Page domain.
type SetAutoAttachToCreatedPagesArgs struct {
	AutoAttach bool `json:"autoAttach"` // If true, browser will open a new inspector window for every page created from this one.
}

// NewSetAutoAttachToCreatedPagesArgs initializes SetAutoAttachToCreatedPagesArgs with the required arguments.
func NewSetAutoAttachToCreatedPagesArgs(autoAttach bool) *SetAutoAttachToCreatedPagesArgs {
	args := new(SetAutoAttachToCreatedPagesArgs)
	args.AutoAttach = autoAttach
	return args
}

// ReloadArgs represents the arguments for Reload in the Page domain.
type ReloadArgs struct {
	IgnoreCache            *bool   `json:"ignoreCache,omitempty"`            // If true, browser cache is ignored (as if the user pressed Shift+refresh).
	ScriptToEvaluateOnLoad *string `json:"scriptToEvaluateOnLoad,omitempty"` // If set, the script will be injected into all frames of the inspected page after reload.
}

// NewReloadArgs initializes ReloadArgs with the required arguments.
func NewReloadArgs() *ReloadArgs {
	args := new(ReloadArgs)

	return args
}

// SetIgnoreCache sets the IgnoreCache optional argument. If true, browser cache is ignored (as if the user pressed Shift+refresh).
func (a *ReloadArgs) SetIgnoreCache(ignoreCache bool) *ReloadArgs {
	a.IgnoreCache = &ignoreCache
	return a
}

// SetScriptToEvaluateOnLoad sets the ScriptToEvaluateOnLoad optional argument. If set, the script will be injected into all frames of the inspected page after reload.
func (a *ReloadArgs) SetScriptToEvaluateOnLoad(scriptToEvaluateOnLoad string) *ReloadArgs {
	a.ScriptToEvaluateOnLoad = &scriptToEvaluateOnLoad
	return a
}

// NavigateArgs represents the arguments for Navigate in the Page domain.
type NavigateArgs struct {
	URL            string         `json:"url"`                      // URL to navigate the page to.
	Referrer       *string        `json:"referrer,omitempty"`       // Referrer URL.
	TransitionType TransitionType `json:"transitionType,omitempty"` // Intended transition type.
}

// NewNavigateArgs initializes NavigateArgs with the required arguments.
func NewNavigateArgs(url string) *NavigateArgs {
	args := new(NavigateArgs)
	args.URL = url
	return args
}

// SetReferrer sets the Referrer optional argument. Referrer URL.
func (a *NavigateArgs) SetReferrer(referrer string) *NavigateArgs {
	a.Referrer = &referrer
	return a
}

// SetTransitionType sets the TransitionType optional argument. Intended transition type.
func (a *NavigateArgs) SetTransitionType(transitionType TransitionType) *NavigateArgs {
	a.TransitionType = transitionType
	return a
}

// GetNavigationHistoryReply represents the return values for GetNavigationHistory in the Page domain.
type GetNavigationHistoryReply struct {
	CurrentIndex int               `json:"currentIndex"` // Index of the current navigation history entry.
	Entries      []NavigationEntry `json:"entries"`      // Array of navigation history entries.
}

// NavigateToHistoryEntryArgs represents the arguments for NavigateToHistoryEntry in the Page domain.
type NavigateToHistoryEntryArgs struct {
	EntryID int `json:"entryId"` // Unique id of the entry to navigate to.
}

// NewNavigateToHistoryEntryArgs initializes NavigateToHistoryEntryArgs with the required arguments.
func NewNavigateToHistoryEntryArgs(entryID int) *NavigateToHistoryEntryArgs {
	args := new(NavigateToHistoryEntryArgs)
	args.EntryID = entryID
	return args
}

// GetCookiesReply represents the return values for GetCookies in the Page domain.
type GetCookiesReply struct {
	Cookies []network.Cookie `json:"cookies"` // Array of cookie objects.
}

// DeleteCookieArgs represents the arguments for DeleteCookie in the Page domain.
type DeleteCookieArgs struct {
	CookieName string `json:"cookieName"` // Name of the cookie to remove.
	URL        string `json:"url"`        // URL to match cooke domain and path.
}

// NewDeleteCookieArgs initializes DeleteCookieArgs with the required arguments.
func NewDeleteCookieArgs(cookieName string, url string) *DeleteCookieArgs {
	args := new(DeleteCookieArgs)
	args.CookieName = cookieName
	args.URL = url
	return args
}

// GetResourceTreeReply represents the return values for GetResourceTree in the Page domain.
type GetResourceTreeReply struct {
	FrameTree FrameResourceTree `json:"frameTree"` // Present frame / resource tree structure.
}

// GetResourceContentReply represents the return values for GetResourceContent in the Page domain.
type GetResourceContentReply struct {
	Content       string `json:"content"`       // Resource content.
	Base64Encoded bool   `json:"base64Encoded"` // True, if content was served as base64.
}

// SearchInResourceReply represents the return values for SearchInResource in the Page domain.
type SearchInResourceReply struct {
	Result []debugger.SearchMatch `json:"result"` // List of search matches.
}

// SetDeviceMetricsOverrideArgs represents the arguments for SetDeviceMetricsOverride in the Page domain.
type SetDeviceMetricsOverrideArgs struct {
	Width              int                          `json:"width"`                        // Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height             int                          `json:"height"`                       // Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	DeviceScaleFactor  float64                      `json:"deviceScaleFactor"`            // Overriding device scale factor value. 0 disables the override.
	Mobile             bool                         `json:"mobile"`                       // Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Scale              *float64                     `json:"scale,omitempty"`              // Scale to apply to resulting view image. Ignored in |fitWindow| mode.
	ScreenWidth        *int                         `json:"screenWidth,omitempty"`        // Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	ScreenHeight       *int                         `json:"screenHeight,omitempty"`       // Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	PositionX          *int                         `json:"positionX,omitempty"`          // Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	PositionY          *int                         `json:"positionY,omitempty"`          // Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	DontSetVisibleSize *bool                        `json:"dontSetVisibleSize,omitempty"` // Do not set visible view size, rely upon explicit setVisibleSize call.
	ScreenOrientation  *emulation.ScreenOrientation `json:"screenOrientation,omitempty"`  // Screen orientation override.
}

// NewSetDeviceMetricsOverrideArgs initializes SetDeviceMetricsOverrideArgs with the required arguments.
func NewSetDeviceMetricsOverrideArgs(width int, height int, deviceScaleFactor float64, mobile bool) *SetDeviceMetricsOverrideArgs {
	args := new(SetDeviceMetricsOverrideArgs)
	args.Width = width
	args.Height = height
	args.DeviceScaleFactor = deviceScaleFactor
	args.Mobile = mobile
	return args
}

// SetScale sets the Scale optional argument. Scale to apply to resulting view image. Ignored in |fitWindow| mode.
func (a *SetDeviceMetricsOverrideArgs) SetScale(scale float64) *SetDeviceMetricsOverrideArgs {
	a.Scale = &scale
	return a
}

// SetScreenWidth sets the ScreenWidth optional argument. Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
func (a *SetDeviceMetricsOverrideArgs) SetScreenWidth(screenWidth int) *SetDeviceMetricsOverrideArgs {
	a.ScreenWidth = &screenWidth
	return a
}

// SetScreenHeight sets the ScreenHeight optional argument. Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
func (a *SetDeviceMetricsOverrideArgs) SetScreenHeight(screenHeight int) *SetDeviceMetricsOverrideArgs {
	a.ScreenHeight = &screenHeight
	return a
}

// SetPositionX sets the PositionX optional argument. Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
func (a *SetDeviceMetricsOverrideArgs) SetPositionX(positionX int) *SetDeviceMetricsOverrideArgs {
	a.PositionX = &positionX
	return a
}

// SetPositionY sets the PositionY optional argument. Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
func (a *SetDeviceMetricsOverrideArgs) SetPositionY(positionY int) *SetDeviceMetricsOverrideArgs {
	a.PositionY = &positionY
	return a
}

// SetDontSetVisibleSize sets the DontSetVisibleSize optional argument. Do not set visible view size, rely upon explicit setVisibleSize call.
func (a *SetDeviceMetricsOverrideArgs) SetDontSetVisibleSize(dontSetVisibleSize bool) *SetDeviceMetricsOverrideArgs {
	a.DontSetVisibleSize = &dontSetVisibleSize
	return a
}

// SetScreenOrientation sets the ScreenOrientation optional argument. Screen orientation override.
func (a *SetDeviceMetricsOverrideArgs) SetScreenOrientation(screenOrientation emulation.ScreenOrientation) *SetDeviceMetricsOverrideArgs {
	a.ScreenOrientation = &screenOrientation
	return a
}

// SetGeolocationOverrideArgs represents the arguments for SetGeolocationOverride in the Page domain.
type SetGeolocationOverrideArgs struct {
	Latitude  *float64 `json:"latitude,omitempty"`  // Mock latitude
	Longitude *float64 `json:"longitude,omitempty"` // Mock longitude
	Accuracy  *float64 `json:"accuracy,omitempty"`  // Mock accuracy
}

// NewSetGeolocationOverrideArgs initializes SetGeolocationOverrideArgs with the required arguments.
func NewSetGeolocationOverrideArgs() *SetGeolocationOverrideArgs {
	args := new(SetGeolocationOverrideArgs)

	return args
}

// SetLatitude sets the Latitude optional argument. Mock latitude
func (a *SetGeolocationOverrideArgs) SetLatitude(latitude float64) *SetGeolocationOverrideArgs {
	a.Latitude = &latitude
	return a
}

// SetLongitude sets the Longitude optional argument. Mock longitude
func (a *SetGeolocationOverrideArgs) SetLongitude(longitude float64) *SetGeolocationOverrideArgs {
	a.Longitude = &longitude
	return a
}

// SetAccuracy sets the Accuracy optional argument. Mock accuracy
func (a *SetGeolocationOverrideArgs) SetAccuracy(accuracy float64) *SetGeolocationOverrideArgs {
	a.Accuracy = &accuracy
	return a
}

// SetDeviceOrientationOverrideArgs represents the arguments for SetDeviceOrientationOverride in the Page domain.
type SetDeviceOrientationOverrideArgs struct {
	Alpha float64 `json:"alpha"` // Mock alpha
	Beta  float64 `json:"beta"`  // Mock beta
	Gamma float64 `json:"gamma"` // Mock gamma
}

// NewSetDeviceOrientationOverrideArgs initializes SetDeviceOrientationOverrideArgs with the required arguments.
func NewSetDeviceOrientationOverrideArgs(alpha float64, beta float64, gamma float64) *SetDeviceOrientationOverrideArgs {
	args := new(SetDeviceOrientationOverrideArgs)
	args.Alpha = alpha
	args.Beta = beta
	args.Gamma = gamma
	return args
}

// SetTouchEmulationEnabledArgs represents the arguments for SetTouchEmulationEnabled in the Page domain.
type SetTouchEmulationEnabledArgs struct {
	Enabled       bool    `json:"enabled"`                 // Whether the touch event emulation should be enabled.
	Configuration *string `json:"configuration,omitempty"` // Touch/gesture events configuration. Default: current platform.
}

// NewSetTouchEmulationEnabledArgs initializes SetTouchEmulationEnabledArgs with the required arguments.
func NewSetTouchEmulationEnabledArgs(enabled bool) *SetTouchEmulationEnabledArgs {
	args := new(SetTouchEmulationEnabledArgs)
	args.Enabled = enabled
	return args
}

// SetConfiguration sets the Configuration optional argument. Touch/gesture events configuration. Default: current platform.
func (a *SetTouchEmulationEnabledArgs) SetConfiguration(configuration string) *SetTouchEmulationEnabledArgs {
	a.Configuration = &configuration
	return a
}

// CaptureScreenshotArgs represents the arguments for CaptureScreenshot in the Page domain.
type CaptureScreenshotArgs struct {
	Format      *string   `json:"format,omitempty"`      // Image compression format (defaults to png).
	Quality     *int      `json:"quality,omitempty"`     // Compression quality from range [0..100] (jpeg only).
	Clip        *Viewport `json:"clip,omitempty"`        // Capture the screenshot of a given region only.
	FromSurface *bool     `json:"fromSurface,omitempty"` // Capture the screenshot from the surface, rather than the view. Defaults to true.
}

// NewCaptureScreenshotArgs initializes CaptureScreenshotArgs with the required arguments.
func NewCaptureScreenshotArgs() *CaptureScreenshotArgs {
	args := new(CaptureScreenshotArgs)

	return args
}

// SetFormat sets the Format optional argument. Image compression format (defaults to png).
func (a *CaptureScreenshotArgs) SetFormat(format string) *CaptureScreenshotArgs {
	a.Format = &format
	return a
}

// SetQuality sets the Quality optional argument. Compression quality from range [0..100] (jpeg only).
func (a *CaptureScreenshotArgs) SetQuality(quality int) *CaptureScreenshotArgs {
	a.Quality = &quality
	return a
}

// SetClip sets the Clip optional argument. Capture the screenshot of a given region only.
func (a *CaptureScreenshotArgs) SetClip(clip Viewport) *CaptureScreenshotArgs {
	a.Clip = &clip
	return a
}

// SetFromSurface sets the FromSurface optional argument. Capture the screenshot from the surface, rather than the view. Defaults to true.
func (a *CaptureScreenshotArgs) SetFromSurface(fromSurface bool) *CaptureScreenshotArgs {
	a.FromSurface = &fromSurface
	return a
}

// CaptureScreenshotReply represents the return values for CaptureScreenshot in the Page domain.
type CaptureScreenshotReply struct {
	Data []byte `json:"data"` // Base64-encoded image data.
}

// PrintToPDFArgs represents the arguments for PrintToPDF in the Page domain.
type PrintToPDFArgs struct {
	Landscape               *bool    `json:"landscape,omitempty"`               // Paper orientation. Defaults to false.
	DisplayHeaderFooter     *bool    `json:"displayHeaderFooter,omitempty"`     // Display header and footer. Defaults to false.
	PrintBackground         *bool    `json:"printBackground,omitempty"`         // Print background graphics. Defaults to false.
	Scale                   *float64 `json:"scale,omitempty"`                   // Scale of the webpage rendering. Defaults to 1.
	PaperWidth              *float64 `json:"paperWidth,omitempty"`              // Paper width in inches. Defaults to 8.5 inches.
	PaperHeight             *float64 `json:"paperHeight,omitempty"`             // Paper height in inches. Defaults to 11 inches.
	MarginTop               *float64 `json:"marginTop,omitempty"`               // Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom            *float64 `json:"marginBottom,omitempty"`            // Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft              *float64 `json:"marginLeft,omitempty"`              // Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight             *float64 `json:"marginRight,omitempty"`             // Right margin in inches. Defaults to 1cm (~0.4 inches).
	PageRanges              *string  `json:"pageRanges,omitempty"`              // Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
	IgnoreInvalidPageRanges *bool    `json:"ignoreInvalidPageRanges,omitempty"` // Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
}

// NewPrintToPDFArgs initializes PrintToPDFArgs with the required arguments.
func NewPrintToPDFArgs() *PrintToPDFArgs {
	args := new(PrintToPDFArgs)

	return args
}

// SetLandscape sets the Landscape optional argument. Paper orientation. Defaults to false.
func (a *PrintToPDFArgs) SetLandscape(landscape bool) *PrintToPDFArgs {
	a.Landscape = &landscape
	return a
}

// SetDisplayHeaderFooter sets the DisplayHeaderFooter optional argument. Display header and footer. Defaults to false.
func (a *PrintToPDFArgs) SetDisplayHeaderFooter(displayHeaderFooter bool) *PrintToPDFArgs {
	a.DisplayHeaderFooter = &displayHeaderFooter
	return a
}

// SetPrintBackground sets the PrintBackground optional argument. Print background graphics. Defaults to false.
func (a *PrintToPDFArgs) SetPrintBackground(printBackground bool) *PrintToPDFArgs {
	a.PrintBackground = &printBackground
	return a
}

// SetScale sets the Scale optional argument. Scale of the webpage rendering. Defaults to 1.
func (a *PrintToPDFArgs) SetScale(scale float64) *PrintToPDFArgs {
	a.Scale = &scale
	return a
}

// SetPaperWidth sets the PaperWidth optional argument. Paper width in inches. Defaults to 8.5 inches.
func (a *PrintToPDFArgs) SetPaperWidth(paperWidth float64) *PrintToPDFArgs {
	a.PaperWidth = &paperWidth
	return a
}

// SetPaperHeight sets the PaperHeight optional argument. Paper height in inches. Defaults to 11 inches.
func (a *PrintToPDFArgs) SetPaperHeight(paperHeight float64) *PrintToPDFArgs {
	a.PaperHeight = &paperHeight
	return a
}

// SetMarginTop sets the MarginTop optional argument. Top margin in inches. Defaults to 1cm (~0.4 inches).
func (a *PrintToPDFArgs) SetMarginTop(marginTop float64) *PrintToPDFArgs {
	a.MarginTop = &marginTop
	return a
}

// SetMarginBottom sets the MarginBottom optional argument. Bottom margin in inches. Defaults to 1cm (~0.4 inches).
func (a *PrintToPDFArgs) SetMarginBottom(marginBottom float64) *PrintToPDFArgs {
	a.MarginBottom = &marginBottom
	return a
}

// SetMarginLeft sets the MarginLeft optional argument. Left margin in inches. Defaults to 1cm (~0.4 inches).
func (a *PrintToPDFArgs) SetMarginLeft(marginLeft float64) *PrintToPDFArgs {
	a.MarginLeft = &marginLeft
	return a
}

// SetMarginRight sets the MarginRight optional argument. Right margin in inches. Defaults to 1cm (~0.4 inches).
func (a *PrintToPDFArgs) SetMarginRight(marginRight float64) *PrintToPDFArgs {
	a.MarginRight = &marginRight
	return a
}

// SetPageRanges sets the PageRanges optional argument. Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
func (a *PrintToPDFArgs) SetPageRanges(pageRanges string) *PrintToPDFArgs {
	a.PageRanges = &pageRanges
	return a
}

// SetIgnoreInvalidPageRanges sets the IgnoreInvalidPageRanges optional argument. Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
func (a *PrintToPDFArgs) SetIgnoreInvalidPageRanges(ignoreInvalidPageRanges bool) *PrintToPDFArgs {
	a.IgnoreInvalidPageRanges = &ignoreInvalidPageRanges
	return a
}

// PrintToPDFReply represents the return values for PrintToPDF in the Page domain.
type PrintToPDFReply struct {
	Data []byte `json:"data"` // Base64-encoded pdf data.
}

// StartScreencastArgs represents the arguments for StartScreencast in the Page domain.
type StartScreencastArgs struct {
	Format        *string `json:"format,omitempty"`        // Image compression format.
	Quality       *int    `json:"quality,omitempty"`       // Compression quality from range [0..100].
	MaxWidth      *int    `json:"maxWidth,omitempty"`      // Maximum screenshot width.
	MaxHeight     *int    `json:"maxHeight,omitempty"`     // Maximum screenshot height.
	EveryNthFrame *int    `json:"everyNthFrame,omitempty"` // Send every n-th frame.
}

// NewStartScreencastArgs initializes StartScreencastArgs with the required arguments.
func NewStartScreencastArgs() *StartScreencastArgs {
	args := new(StartScreencastArgs)

	return args
}

// SetFormat sets the Format optional argument. Image compression format.
func (a *StartScreencastArgs) SetFormat(format string) *StartScreencastArgs {
	a.Format = &format
	return a
}

// SetQuality sets the Quality optional argument. Compression quality from range [0..100].
func (a *StartScreencastArgs) SetQuality(quality int) *StartScreencastArgs {
	a.Quality = &quality
	return a
}

// SetMaxWidth sets the MaxWidth optional argument. Maximum screenshot width.
func (a *StartScreencastArgs) SetMaxWidth(maxWidth int) *StartScreencastArgs {
	a.MaxWidth = &maxWidth
	return a
}

// SetMaxHeight sets the MaxHeight optional argument. Maximum screenshot height.
func (a *StartScreencastArgs) SetMaxHeight(maxHeight int) *StartScreencastArgs {
	a.MaxHeight = &maxHeight
	return a
}

// SetEveryNthFrame sets the EveryNthFrame optional argument. Send every n-th frame.
func (a *StartScreencastArgs) SetEveryNthFrame(everyNthFrame int) *StartScreencastArgs {
	a.EveryNthFrame = &everyNthFrame
	return a
}

// ScreencastFrameAckArgs represents the arguments for ScreencastFrameAck in the Page domain.
type ScreencastFrameAckArgs struct {
	SessionID int `json:"sessionId"` // Frame number.
}

// NewScreencastFrameAckArgs initializes ScreencastFrameAckArgs with the required arguments.
func NewScreencastFrameAckArgs(sessionID int) *ScreencastFrameAckArgs {
	args := new(ScreencastFrameAckArgs)
	args.SessionID = sessionID
	return args
}

// HandleJavaScriptDialogArgs represents the arguments for HandleJavaScriptDialog in the Page domain.
type HandleJavaScriptDialogArgs struct {
	Accept     bool    `json:"accept"`               // Whether to accept or dismiss the dialog.
	PromptText *string `json:"promptText,omitempty"` // The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
}

// NewHandleJavaScriptDialogArgs initializes HandleJavaScriptDialogArgs with the required arguments.
func NewHandleJavaScriptDialogArgs(accept bool) *HandleJavaScriptDialogArgs {
	args := new(HandleJavaScriptDialogArgs)
	args.Accept = accept
	return args
}

// SetPromptText sets the PromptText optional argument. The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
func (a *HandleJavaScriptDialogArgs) SetPromptText(promptText string) *HandleJavaScriptDialogArgs {
	a.PromptText = &promptText
	return a
}

// GetAppManifestReply represents the return values for GetAppManifest in the Page domain.
type GetAppManifestReply struct {
	URL    string             `json:"url"`            // Manifest location.
	Errors []AppManifestError `json:"errors"`         //
	Data   *string            `json:"data,omitempty"` // Manifest content.
}

// SetControlNavigationsArgs represents the arguments for SetControlNavigations in the Page domain.
type SetControlNavigationsArgs struct {
	Enabled bool `json:"enabled"` //
}

// NewSetControlNavigationsArgs initializes SetControlNavigationsArgs with the required arguments.
func NewSetControlNavigationsArgs(enabled bool) *SetControlNavigationsArgs {
	args := new(SetControlNavigationsArgs)
	args.Enabled = enabled
	return args
}

// ProcessNavigationArgs represents the arguments for ProcessNavigation in the Page domain.
type ProcessNavigationArgs struct {
	Response     NavigationResponse `json:"response"`     //
	NavigationID int                `json:"navigationId"` //
}

// NewProcessNavigationArgs initializes ProcessNavigationArgs with the required arguments.
func NewProcessNavigationArgs(response NavigationResponse, navigationID int) *ProcessNavigationArgs {
	args := new(ProcessNavigationArgs)
	args.Response = response
	args.NavigationID = navigationID
	return args
}

// GetLayoutMetricsReply represents the return values for GetLayoutMetrics in the Page domain.
type GetLayoutMetricsReply struct {
	LayoutViewport LayoutViewport `json:"layoutViewport"` // Metrics relating to the layout viewport.
	VisualViewport VisualViewport `json:"visualViewport"` // Metrics relating to the visual viewport.
	ContentSize    dom.Rect       `json:"contentSize"`    // Size of scrollable area.
}

// CreateIsolatedWorldReply represents the return values for CreateIsolatedWorld in the Page domain.
type CreateIsolatedWorldReply struct {
	ExecutionContextID runtime.ExecutionContextID `json:"executionContextId"` // Execution context of the isolated world.
}