package rule

// Monitor provides methods to monitoring evaluation.
type Monitor interface {
	ConditionError(*Context, *Rule, error)
	ConditionResult(*Context, *Rule, bool)
	ActionIgnore(*Context, *Rule)
	ActionCompileError(*Context, *Rule, error)
	ActionError(*Context, *Rule, error)
	ActionResult(*Context, *Rule, interface{})
}

type dummyMonitor struct{}

func (*dummyMonitor) ConditionError(*Context, *Rule, error)     {}
func (*dummyMonitor) ConditionResult(*Context, *Rule, bool)     {}
func (*dummyMonitor) ActionIgnore(*Context, *Rule)              {}
func (*dummyMonitor) ActionCompileError(*Context, *Rule, error) {}
func (*dummyMonitor) ActionError(*Context, *Rule, error)        {}
func (*dummyMonitor) ActionResult(*Context, *Rule, interface{}) {}

type MonitorHooks struct {
	OnConditionError     func(*Context, *Rule, error)
	OnConditionResult    func(*Context, *Rule, bool)
	OnActionIgnore       func(*Context, *Rule)
	OnActionCompileError func(*Context, *Rule, error)
	OnActionError        func(*Context, *Rule, error)
	OnActionResult       func(*Context, *Rule, interface{})
}

func (m *MonitorHooks) ConditionError(ctx *Context, r *Rule, err error) {
	if m.OnConditionError != nil {
		m.OnConditionError(ctx, r, err)
	}
}

func (m *MonitorHooks) ConditionResult(ctx *Context, r *Rule, f bool) {
	if m.OnConditionResult != nil {
		m.OnConditionResult(ctx, r, f)
	}
}

func (m *MonitorHooks) ActionIgnore(ctx *Context, r *Rule) {
	if m.OnActionIgnore != nil {
		m.OnActionIgnore(ctx, r)
	}
}

func (m *MonitorHooks) ActionCompileError(ctx *Context, r *Rule, err error) {
	if m.OnActionCompileError != nil {
		m.OnActionCompileError(ctx, r, err)
	}
}

func (m *MonitorHooks) ActionError(ctx *Context, r *Rule, err error) {
	if m.OnActionError != nil {
		m.OnActionError(ctx, r, err)
	}
}

func (m *MonitorHooks) ActionResult(ctx *Context, r *Rule, v interface{}) {
	if m.OnActionResult != nil {
		m.OnActionResult(ctx, r, v)
	}
}
