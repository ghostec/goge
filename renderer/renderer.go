package renderer

import (
	"time"

	"github.com/ghostec/goge/camera"
	"github.com/ghostec/goge/scene"
)

type Renderer interface {
	// Attach, Dettach ? with origin
	// SetSize ?
	Render() error
	Camera() camera.Camera
	SetCamera(camera.Camera)
	SetScene(*scene.Scene)
	SetScreen(Screen)
	Update(elapsed time.Duration) error
}
