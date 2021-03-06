package html

import (
	"context"

	"github.com/MontFerret/ferret/pkg/html/dynamic"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

// WaitClass waits for a class to appear on a given element.
// Stops the execution until the navigation ends or operation times out.
// @param docOrEl (HTMLDocument|HTMLElement) - Target document or element.
// @param selectorOrClass (String) - If document is passed, this param must represent an element selector.
// Otherwise target class.
// @param classOrTimeout (String|Int, optional) - If document is passed, this param must represent target class name.
// Otherwise timeout.
// @param timeout (Int, optional) - If document is passed, this param must represent timeout.
// Otherwise not passed.
func WaitClass(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 4)

	if err != nil {
		return values.None, err
	}

	// document or element
	err = core.ValidateType(args[0], core.HTMLDocumentType, core.HTMLElementType)

	if err != nil {
		return values.None, err
	}

	// selector or class
	err = core.ValidateType(args[1], core.StringType)

	if err != nil {
		return values.None, err
	}

	timeout := values.NewInt(defaultTimeout)

	// lets figure out what is passed as 1st argument
	switch args[0].(type) {
	case *dynamic.HTMLDocument:
		// class
		err = core.ValidateType(args[2], core.StringType)

		if err != nil {
			return values.None, err
		}

		doc, ok := args[0].(*dynamic.HTMLDocument)

		if !ok {
			return values.None, core.Errors(core.ErrInvalidType, ErrNotDynamic)
		}

		selector := args[1].(values.String)
		class := args[2].(values.String)

		if len(args) == 4 {
			err = core.ValidateType(args[3], core.IntType)

			if err != nil {
				return values.None, err
			}

			timeout = args[3].(values.Int)
		}

		return values.None, doc.WaitForClass(selector, class, timeout)
	case *dynamic.HTMLElement:
		el, ok := args[0].(*dynamic.HTMLElement)

		if !ok {
			return values.None, core.Errors(core.ErrInvalidType, ErrNotDynamic)
		}

		class := args[1].(values.String)

		if len(args) == 3 {
			err = core.ValidateType(args[2], core.IntType)

			if err != nil {
				return values.None, err
			}

			timeout = args[3].(values.Int)
		}

		return values.None, el.WaitForClass(class, timeout)
	default:
		return values.None, core.Errors(core.ErrInvalidType, ErrNotDynamic)
	}
}
