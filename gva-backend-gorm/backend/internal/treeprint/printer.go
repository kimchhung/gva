package treeprint

import (
	"fmt"
	"io"
	"os"
)

// Box-drawing symbols
const (
	EdgeEmpty = "    "
	EdgePipe  = "│   "
	EdgeItem  = "├── "
	EdgeLast  = "└── "
)

// DefaultPrinter prints the tree to the stdout with default settings.
var DefaultPrinter = NewPrinter(os.Stdout, nil)

// Printer prints the tree.
type Printer interface {
	Print(*Node) error
}

// FormaterFunc formats a single node.
type FormaterFunc func(*Node) string

type printer struct {
	writer io.Writer
	format FormaterFunc
}

func NewPrinter(writer io.Writer, format FormaterFunc) Printer {
	if format == nil {
		format = func(n *Node) string {
			return fmt.Sprint(n.Value)
		}
	}
	return &printer{
		writer: writer,
		format: format,
	}
}

func (p *printer) Print(root *Node) error {
	return p.print("", root)
}

func (p *printer) print(prefix string, n *Node) error {
	if _, err := fmt.Fprintln(p.writer, p.format(n)); err != nil {
		return err
	}
	size := len(n.Nodes)
	if size == 0 {
		return nil
	}

	prefixItem := prefix + EdgeItem
	prefixPipe := prefix + EdgePipe
	last := size - 1
	for _, n := range n.Nodes[:last] {
		if _, err := io.WriteString(p.writer, prefixItem); err != nil {
			return err
		}
		if err := p.print(prefixPipe, n); err != nil {
			return nil
		}
	}
	if _, err := io.WriteString(p.writer, prefix+EdgeLast); err != nil {
		return err
	}
	return p.print(prefix+EdgeEmpty, n.Nodes[last])
}

// PrinterFunc type is an adapter to allow the use of ordinary functions as a printer.
type PrinterFunc func(*Node) error

func (f PrinterFunc) Print(root *Node) error {
	return f(root)
}

// Print prints the tree using the default printer.
func Print(root *Node) error {
	return DefaultPrinter.Print(root)
}
