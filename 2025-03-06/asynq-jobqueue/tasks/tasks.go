package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/cshum/imagor"
	"github.com/cshum/imagor/imagorpath"
	"github.com/cshum/imagor/loader/httploader"
	"github.com/cshum/imagor/vips"
	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeImageResize = "image:resize"
)

type ImageResizePayload struct {
	SourceURL  string
	TargetPath string
	Width      int
	Height     int
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewImageResizeTask(src, targetPath string, width, height int) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{
		SourceURL:  src,
		TargetPath: targetPath,
		Width:      width,
		Height:     height,
	})
	if err != nil {
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(TypeImageResize, payload, asynq.MaxRetry(5)), nil
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.
//---------------------------------------------------------------

// ImageProcessor implements asynq.Handler interface.
type ImageProcessor struct {
	img *imagor.Imagor
}

func NewImageProcessor() *ImageProcessor {
	img := imagor.New(
		imagor.WithLoaders(httploader.New()),
		imagor.WithProcessors(vips.NewProcessor()),
	)

	return &ImageProcessor{
		img: img,
	}
}

func (p *ImageProcessor) Startup(ctx context.Context) error {
	return p.img.Startup(ctx)
}

func (p *ImageProcessor) Shutdown(ctx context.Context) error {
	return p.img.Shutdown(ctx)
}

func (p *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("decoding payload: %v: %w", err, asynq.SkipRetry)
	}
	slog.Debug("Resizing image", "src", payload.SourceURL)

	blob, err := p.img.Serve(ctx, imagorpath.Params{
		Image:  payload.SourceURL,
		Width:  payload.Width,
		Height: payload.Height,
		Smart:  true,
		Filters: []imagorpath.Filter{
			{"fill", "white"},
			{"format", "jpg"},
		},
	})
	if err != nil {
		return fmt.Errorf("resizing with imagor: %w", err)
	}
	reader, _, err := blob.NewReader()
	if err != nil {
		return fmt.Errorf("getting reader: %w", err)
	}
	defer reader.Close()
	file, err := os.Create(payload.TargetPath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()
	if _, err := io.Copy(file, reader); err != nil {
		return fmt.Errorf("copying file: %w", err)
	}

	slog.Info("Resized image", "src", payload.SourceURL, "target", payload.TargetPath)

	return nil
}
