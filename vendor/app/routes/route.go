package route

import (
	"net/http"

	"app/common/session"
	"app/controller"
	"app/routes/middleware/acl"
	hr "app/routes/middleware/httprouterwrapper"
	"app/routes/middleware/logrequest"
	"app/routes/middleware/pprofhandler"

	"github.com/gorilla/context"
	"github.com/josephspurrier/csrfbanana"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Load returns the routes and middleware
func Load() http.Handler {
	return middleware(routes())
}

// LoadHTTPS returns the HTTP routes and middleware
func LoadHTTPS() http.Handler {
	return middleware(routes())
}

// LoadHTTP returns the HTTPS routes and middleware
func LoadHTTP() http.Handler {
	return middleware(routes())

	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
}

// Optional method to make it easy to redirect from HTTP to HTTPS
func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}

// *****************************************************************************
// Routes
// *****************************************************************************

func routes() *httprouter.Router {
	r := httprouter.New()

	// Set 404 handler
	r.NotFound = alice.
		New().
		ThenFunc(controller.Error404)

	// Serve static files, no directory browsing
	r.GET("/static/*filepath", hr.Handler(alice.
		New().
		ThenFunc(controller.Static)))

	// Home page
	r.GET("/", hr.Handler(alice.
		New().
		ThenFunc(controller.IndexGET)))

	r.POST("/login", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controller.LoginPOST)))
	r.GET("/logout", hr.Handler(alice.
		New().
		ThenFunc(controller.LogoutGET)))
	r.POST("/register", hr.Handler(alice.
		New().
		ThenFunc(controller.RegisterPOST)))

	// Test Page
	r.GET("/test", hr.Handler(alice.
		New().
		ThenFunc(controller.TestGET)))

	// Inventory Page
	r.GET("/inventory", hr.Handler(alice.
		New().
		ThenFunc(controller.InventoryGET)))

	// Inventory retriever
	r.GET("/api/inventory", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controller.GetInventory)))

	// Inventory retriever
	r.GET("/api/item", hr.Handler(alice.
		New().
		ThenFunc(controller.CreateItem)))

	r.GET("/api/item/delete", hr.Handler(alice.
		New().
		ThenFunc(controller.DeleteItem)))

	// Sales Page
	r.GET("/sales", hr.Handler(alice.
		New().
		ThenFunc(controller.SalesGET)))

	// Orders Page
	r.GET("/orders", hr.Handler(alice.
		New().
		ThenFunc(controller.OrdersGET)))

	r.GET("/api/orders", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controller.GetOrders)))

	// Enable Pprof
	r.GET("/debug/pprof/*pprof", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(pprofhandler.Handler)))

	return r
}

// *****************************************************************************
// Middleware
// *****************************************************************************

func middleware(h http.Handler) http.Handler {
	// Prevents CSRF and Double Submits
	cs := csrfbanana.New(h, session.Store, session.Name)
	cs.FailureHandler(http.HandlerFunc(controller.InvalidToken))
	cs.ClearAfterUsage(true)
	cs.ExcludeRegexPaths([]string{"/static(.*)"})
	csrfbanana.TokenLength = 32
	csrfbanana.TokenName = "token"
	csrfbanana.SingleToken = true
	h = cs

	// Log every request
	h = logrequest.Handler(h)

	// Clear handler for Gorilla Context
	h = context.ClearHandler(h)

	return h
}
