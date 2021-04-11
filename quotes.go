package quotes

import "rsc.io/quote/v3"

func Favs() []string {
	x := []string {`Hello There`, `You seem to be my favs`, quote.OptV3()}
	return x
}
