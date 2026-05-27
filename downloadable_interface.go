package notionapi

import "time"

// DownloadableFileBlock is an interface for blocks that can be downloaded
// such as Pdf, FileBlock, and Image
type DownloadableFileBlock interface {
	Block
	GetURL() string
	GetExpiryTime() *time.Time
}

// GetURL implements DownloadableFileBlock interface for PdfBlock
func (b *PdfBlock) GetURL() string { _ = "STUB: not implemented"; return "" }

// GetExpiryTime implements DownloadableFileBlock interface for PdfBlock
func (b *PdfBlock) GetExpiryTime() *time.Time { _ = "STUB: not implemented"; return nil }

// GetURL implements DownloadableFileBlock interface for FileBlock
func (b *FileBlock) GetURL() string { _ = "STUB: not implemented"; return "" }

// GetExpiryTime implements DownloadableFileBlock interface for FileBlock
func (b *FileBlock) GetExpiryTime() *time.Time { _ = "STUB: not implemented"; return nil }

// GetURL implements DownloadableFileBlock interface for ImageBlock
func (b *ImageBlock) GetURL() string { _ = "STUB: not implemented"; return "" }

// GetExpiryTime implements DownloadableFileBlock interface for ImageBlock
func (b *ImageBlock) GetExpiryTime() *time.Time { _ = "STUB: not implemented"; return nil }

// Verify that types implement DownloadableFileBlock interface
var (
	_ DownloadableFileBlock = (*PdfBlock)(nil)
	_ DownloadableFileBlock = (*FileBlock)(nil)
	_ DownloadableFileBlock = (*ImageBlock)(nil)
)
