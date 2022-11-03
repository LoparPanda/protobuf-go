// <LOPAR>Stuff we need to access.
package proto

import "google.golang.org/protobuf/internal/detrand"

func DisableDetrand() {
	detrand.Disable()
}
