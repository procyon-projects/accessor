package accessor

type RequestInterceptor interface {
	Apply(template RequestTemplate)
}
