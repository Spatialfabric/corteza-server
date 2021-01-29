package types

import "context"

type (
	// Parameter is one piece of information that the rendering driver is allowed to use.
	//
	// @note I think that it makes sense to keep these parameters as simple as possible.
	//       Would it make sense to use the same approach as we do with workflows?
	Parameter struct {
		Name  string
		Value interface{}
		Kind  string
	}
	ParameterSet []*Parameter

	// Request is what the driver can work with to produce a series of documents
	//
	// What document type to render should be determined by the caller and not left
	// up to the driver.
	Request struct {
		Templates  TemplateSet
		Parameters ParameterSet
	}

	// Response is what the driver returns as a rendering result
	Response struct {
		Documents DocumentSet
	}

	Driver interface {
		// Render asks the driver to render the document
		// @note we could have a RenderDeferred variant as well
		Render(ctx context.Context, r *Request) (rsp *Response, err error)

		// CanRender determines if this driver is able to use the given TemplateType
		CanRender(tt TemplateType) bool
		// CanProduce determines if this driver is able to produce a specific document type
		CanProduce(dt DocumentType) bool
	}
	DriverSet []Driver
)

// GetDrivers returns a set of drivers that can encode the given content type
func (ss DriverSet) GetDrivers(t DocumentType) DriverSet {
	rr := make(DriverSet, 0, len(ss))
	for _, s := range ss {
		if s.CanProduce(t) {
			rr = append(rr, s)
		}
	}

	return rr
}
