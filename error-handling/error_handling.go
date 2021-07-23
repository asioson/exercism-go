// package erratum implements a routine to demonstrate
// various kinds of error handling and resource management
package erratum

// Use takes in an object meant to open a resource and
// an input string to test various kinds of error handling
// and resource management 
func Use(o ResourceOpener, input string) (err error) {
    var res Resource

    res, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
        res, err = o()
	}
	defer res.Close()
	defer func() {
        r := recover()
		if r != nil {
            switch e := r.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()

	res.Frob(input)
	return err
}
