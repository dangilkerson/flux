package tableutil

import (
	"reflect"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/plan"
)

// Constructor is the minimum interface for constructing a Table.
type Constructor interface {
	// Table will construct a Table from the existing contents.
	// Invoking this method should reset the builder and all allocated
	// memory will be owned by the returned flux.Table.
	Table() (flux.Table, error)

	// Release will release the buffered contents from the builder.
	// This method is unnecessary if Table is called.
	Release()
}

type BuilderContext struct {
	Builder Constructor
	Trigger execute.Trigger
}

// BuilderCache is an alternative table builder cache.
// It allows a pluggable implementation of the builder.
// The only requirement for the builder is that it must
// implement the above Constructor interface.
type BuilderCache struct {
	New func(key flux.GroupKey) Constructor

	tables      *execute.GroupLookup
	triggerSpec plan.TriggerSpec
}

// Get retrieves the TableBuilder for this group key.
// If one doesn't exist, it will invoke the New function and
// store it within the TableBuilder.
// If the builder was newly created, this method returns true
// for the second parameter.
// The interface must be a pointer to the type that is created
// from the New method. This method will use reflection to set
// the value of the pointer.
func (d *BuilderCache) Get(key flux.GroupKey, b interface{}) bool {
	ts, ok := d.lookupState(key)
	if !ok {
		if d.tables == nil {
			d.tables = execute.NewGroupLookup()
		}
		builder := d.New(key)
		triggerSpec := d.triggerSpec
		if triggerSpec == nil {
			triggerSpec = plan.DefaultTriggerSpec
		}
		t := execute.NewTriggerFromSpec(triggerSpec)
		ts = BuilderContext{
			Builder: builder,
			Trigger: t,
		}
		d.tables.Set(key, ts)
	}
	r := reflect.ValueOf(b)
	r.Elem().Set(reflect.ValueOf(ts.Builder))
	return !ok
}

// Table will remove a builder from the cache and construct a flux.Table
// from the buffered contents.
func (d *BuilderCache) Table(key flux.GroupKey) (flux.Table, error) {
	if d.tables == nil {
		return nil, errors.Newf(codes.Internal, "table not found with key %v", key)
	}
	ts, ok := d.tables.Delete(key)
	if !ok {
		return nil, errors.Newf(codes.Internal, "table not found with key %v", key)
	}
	builder := ts.(BuilderContext).Builder
	return builder.Table()
}

func (d *BuilderCache) ForEachBuilder(f func(key flux.GroupKey, builder Constructor) error) error {
	return d.ForEachContext(func(key flux.GroupKey, context BuilderContext) error {
		return f(key, context.Builder)
	})
}

func (d *BuilderCache) ForEachContext(f func(key flux.GroupKey, context BuilderContext) error) error {
	if d.tables == nil {
		return nil
	}
	var err error
	d.tables.Range(func(key flux.GroupKey, value interface{}) {
		if err != nil {
			return
		}
		ts := value.(BuilderContext)
		err = f(key, ts)
	})
	return err
}

func (d *BuilderCache) lookupState(key flux.GroupKey) (BuilderContext, bool) {
	if d.tables == nil {
		return BuilderContext{}, false
	}
	v, ok := d.tables.Lookup(key)
	if !ok {
		return BuilderContext{}, false
	}
	return v.(BuilderContext), true
}

func (d *BuilderCache) DiscardTable(key flux.GroupKey) {
	if d.tables == nil {
		return
	}

	if b, ok := d.lookupState(key); ok {
		// If the builder supports Clear, then call that method.
		if builder, ok := b.Builder.(interface {
			Clear()
		}); ok {
			builder.Clear()
		} else {
			// Release the table and construct a new one.
			b.Builder.Release()
			b.Builder = d.New(key)
		}
	}
}

func (d *BuilderCache) ExpireTable(key flux.GroupKey) {
	if d.tables == nil {
		return
	}
	ts, ok := d.tables.Delete(key)
	if ok {
		ts.(BuilderContext).Builder.Release()
	}
}

func (d *BuilderCache) SetTriggerSpec(t plan.TriggerSpec) {
	d.triggerSpec = t
}

type BuilderDataCache struct {
	*BuilderCache
}

func (d *BuilderDataCache) ForEach(f func(flux.GroupKey)) {
	_ = d.ForEachContext(func(key flux.GroupKey, _ BuilderContext) error {
		f(key)
		return nil
	})
}

func (d *BuilderDataCache) ForEachWithContext(f func(flux.GroupKey, execute.Trigger, execute.TableContext)) {
	_ = d.ForEachContext(func(key flux.GroupKey, context BuilderContext) error {
		f(key, context.Trigger, execute.TableContext{
			Key: key,
		})
		return nil
	})
}
