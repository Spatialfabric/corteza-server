package types

type (
	Session struct {
		ID      uint64
		Invoker uint64
		Proxy   *Proxy

		Params    ParameterSet
		Templates TemplateSet
	}

	SessionSet []*Session
)

func (ss SessionSet) FindByID(id uint64) *Session {
	for _, s := range ss {
		if s.ID == id {
			return s
		}
	}
	return nil
}

func (s *Session) ToRequest() *Request {
	return &Request{
		Templates:  s.Templates,
		Parameters: s.Params,
	}
}
