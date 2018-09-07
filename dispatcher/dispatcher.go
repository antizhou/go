package dispatcher

import "sync"

var (
	handlers   []Handler
	handlersMu sync.Mutex
)

// Handler is a handler with contains a constructor and destructor
type Handler struct {
	Constructor func()
	Destructor  func()
}

// Start the dispatcher, it will call all constructor functions
func Start() {
	handlersMu.Lock()
	defer handlersMu.Unlock()

	wg := sync.WaitGroup{}

	for _, handler := range handlers {
		wg.Add(1)
		constructor := handler.Constructor
		go func() {
			constructor()
			wg.Done()
		}()
	}

	wg.Wait()
}

// Stop the dispatcher, it will call all destructor functions
// Always it will be put after defer
func Stop() {
	handlersMu.Lock()
	defer handlersMu.Unlock()

	wg := sync.WaitGroup{}

	for _, handler := range handlers {
		wg.Add(1)
		destructor := handler.Destructor
		go func() {
			destructor()
			wg.Done()
		}()
	}

	wg.Wait()
}

// Register a constructor and destructor into dispatcher
func Register(constructor, destructor func()) {
	handlersMu.Lock()
	defer handlersMu.Unlock()

	handler := Handler{
		Constructor: constructor,
		Destructor:  destructor,
	}

	handlers = append(handlers, handler)
}
