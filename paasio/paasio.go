// package paasio implements opCounter, readCounter, writeCounter,
// and rwCounter
package paasio

import "io" 
import "sync"


// opCounter object and methods
type opCounter struct {
    mux   sync.Mutex
    bytes int64
    ops   int32
    f     func([]byte) (int, error)
}

func (oc *opCounter) Do(p []byte) (int, error) {
    n, err := oc.f(p)
    if n != 0 {
        oc.mux.Lock()
        oc.bytes += int64(n)
        oc.ops++
        oc.mux.Unlock()
    }
    return n, err
}

func (oc *opCounter) Count() (int64, int) {
    oc.mux.Lock()
    defer oc.mux.Unlock()
    return oc.bytes, int(oc.ops)
}


// readCounter object and methods
type readCounter struct {
    counter opCounter
}

func NewReadCounter(r io.Reader) ReadCounter {
    return &readCounter{counter: opCounter{f: r.Read}}
}

func (r *readCounter) Read(p []byte) (int, error) { 
    return r.counter.Do(p) 
}

func (r *readCounter) ReadCount() (int64, int) {
    return r.counter.Count() 
}


// writeCounter object and methods
type writeCounter struct {
    counter opCounter
}

func NewWriteCounter(w io.Writer) WriteCounter {
    return &writeCounter{counter: opCounter{f: w.Write}}
}

func (w *writeCounter) Write(p []byte) (int, error) { 
    return w.counter.Do(p) 
}

func (w *writeCounter) WriteCount() (int64, int) { 
    return w.counter.Count() 
}


// rwCounter object and methods
type rwCounter struct {
    readCounter
    writeCounter
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
    return &rwCounter{
        readCounter:  readCounter{counter: opCounter{f: rw.Read}},
        writeCounter: writeCounter{counter: opCounter{f: rw.Write}},
    }
}
