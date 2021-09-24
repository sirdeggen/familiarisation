package shortener

// Redirect Services can find and store short urls
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
