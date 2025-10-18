package bootstrap

import (
	"backend/internal/treeprint"
	"backend/utils/color"
	"context"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
)

// notify when server started
func (b *Bootstrap) Started(done chan struct{}) <-chan struct{} {
	b.startedListeners = append(b.startedListeners, done)
	return done
}

func (b *Bootstrap) OnStarted(ctx context.Context, runFn func()) {
	go func() {
		done := make(chan struct{})
		defer func() {
			done <- struct{}{}
		}()

		for {
			select {
			case <-b.Started(done):
				runFn()
				return
			case <-ctx.Done():
				return
			}
		}
	}()
}

// notify when server is shuting down
func (b *Bootstrap) ShuttingDown(done chan struct{}) <-chan struct{} {
	b.shutdownListerners = append(b.shutdownListerners, done)
	return done
}

func (b *Bootstrap) OnShuttingDown(ctx context.Context, closeFn func()) {
	go func() {
		done := make(chan struct{})
		defer func() {
			done <- struct{}{}
		}()

		for {
			select {
			case <-b.ShuttingDown(done):
				closeFn()
				return
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (b *Bootstrap) notifyShuttingDown() {
	resps := make([]chan struct{}, len(b.startedListeners))
	for i, req := range b.shutdownListerners {
		// notify shutdown
		req <- struct{}{}
		resps[i] = req
	}

	for _, resp := range resps {
		<-resp
		close(resp)
	}
}

func (b *Bootstrap) notifyStarted() {
	for _, req := range b.startedListeners {
		// notify server started
		req <- struct{}{}

		// wait process is done
		<-req
		close(req)
	}
}

func printRoutes(routes []*echo.Route) {
	tree := treeprint.New("api")
	N := 4

	sort.Slice(routes, func(i, j int) bool {
		return len(strings.Split(routes[i].Path, "")) > len(strings.Split(routes[j].Path, ""))
	})

	maxLenth := calculateMaxLength(routes, N)
	for _, r := range routes {
		if r.Method == "echo_route_not_found" {
			continue
		}

		paths := []any{}
		for _, str := range strings.SplitAfterN(r.Path, "/", N) {
			str := strings.TrimSuffix(str, "/")
			if str == "" {
				continue
			}
			paths = append(paths, strings.TrimSuffix(str, "/"))
		}
		if len(paths) > N-2 {
			paths[N-2] = strings.ReplaceAll(strings.Split(paths[N-2].(string), "/")[0], "/", "")
		}
		httpPath := color.MethodColor(r.Method) + " " + r.Path
		space := calculateDynamicSpace(httpPath, maxLenth)
		paths = append(paths, httpPath+space+color.Cyan(r.Name))
		tree.AddPath(paths...)
	}

	treeprint.Print(tree)
}

func calculateMaxLength(routes []*echo.Route, N int) int {
	maxLength := 0

	for _, r := range routes {
		paths := []any{}
		for _, str := range strings.SplitAfterN(r.Path, "/", N) {
			paths = append(paths, strings.TrimSuffix(str, "/"))
		}

		length := 0
		for _, str := range paths {
			strs := strings.ReplaceAll(str.(string), "/", "")
			length += len(strs)
		}

		length += len(color.MethodColor(r.Method))
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func calculateDynamicSpace(path string, maxLength int) string {
	spaceNeeded := maxLength - len(path)
	if spaceNeeded <= 0 {
		spaceNeeded = 1
	}
	return strings.Repeat(" ", spaceNeeded)
}
