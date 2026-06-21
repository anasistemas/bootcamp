# Answers

## Were you surprised by any of the responses?

Yes! I expected `/blog`, `/hello`, and `/world` to return a 404 error
since those routes are not registered in the mux. However, all three
returned "This is the home page."

This happens because in Go's ServeMux, the `/` pattern acts as a
catch-all. Any request that doesn't match a registered route falls
through to the `/` handler automatically.