package qrcode

import "github.com/yeqown/go-qrcode/v2/matrix"

// Writer is the interface of a QR code writer, it defines the rule of how to
// `print` the code image from matrix. There's built-in writer to output into
// file, terminal.
type Writer interface {
	// Write writes the code image into itself stream, such as io.Writer,
	// terminal output stream, and etc
	Write(mat matrix.Matrix) error

	// Close the writer stream if it exists after QRCode.Save() is called.
	Close() error
}

var _ Writer = (*nonWriter)(nil)

type nonWriter struct{}

func (n nonWriter) Close() error {
	return nil
}

func (n nonWriter) Write(mat matrix.Matrix) error {
	return nil
}
