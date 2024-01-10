package log

type contextKey int

const requestIDKey contextKey = iota


func FromContext(ctx context.Context) (string, error) {
	if v := ctx.Value(requestIDKey); v != nil {
		return v.(string), nil
	}

	return "", errors.New("requestID not found in context")

}

func ToContext(ctx context.Context, requestID string)context.Context {
	ctx := context.WithValue(parentCtx, requestIDKey, requestID)

	return ctx
}
